package tower

import (
	"context"
	"errors"
	"github.com/go-tower/tower/logger"
	"io"
	"net"
	"sync"
)

type Connectioner interface {
	Start()                                      //启动连接，让当前连接开始工作
	Stop()                                       //停止连接，结束当前连接状态M
	GetTCPConnection() *net.TCPConn              //从当前连接获取原始的socket TCPConn
	GetConnID() uint                             //获取当前连接ID
	RemoteAddr() net.Addr                        //获取远程客户端地址信息
	SendMsg(msgId uint, data []byte) error       //直接将Message数据发送数据给远程的TCP客户端(无缓冲)
	SendBuffMsg(msgId uint, data []byte) error   //直接将Message数据发送给远程的TCP客户端(有缓冲)
	SetProperty(key string, value interface{})   //设置链接属性
	GetProperty(key string) (interface{}, error) //获取链接属性
	RemoveProperty(key string)                   //移除链接属性
}

type Connection struct {
	Server    BootStraper
	Conn      *net.TCPConn
	ConnID    uint
	ctx       context.Context
	ctxCancel context.CancelFunc
	//无缓冲管道，用于读、写两个goroutine之间的消息通信
	msgChan chan []byte
	//有缓冲管道，用于读、写两个goroutine之间的消息通信
	msgBuffChan chan []byte
	logging     logger.Logger

	sync.RWMutex
	property     map[string]interface{} //链接属性
	propertyLock sync.Mutex             //保护当前property的锁
	isClosed     bool                   ///当前连接的关闭状态
}

func NewConnection(server BootStraper, conn *net.TCPConn, connID uint) *Connection {
	c := &Connection{
		Server:      server,
		Conn:        conn,
		ConnID:      connID,
		msgChan:     make(chan []byte),
		msgBuffChan: make(chan []byte, server.getConfig().MaxMsgChanLen),
		logging:     server.getConfig().Logging,
		property:    make(map[string]interface{}),
		isClosed:    false,
	}
	c.Server.GetConnMgr().Add(c)
	return c
}

func (c *Connection) startWrite() {
	c.logging.Debug("[Writer Goroutine is running]")
	defer c.logging.Debug("%s [conn Writer exit!]", c.RemoteAddr().String())
	for {
		select {
		case data := <-c.msgChan: //有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				c.logging.Error("Send Buff Data error:, %v Conn Writer exit", err)
				return
			}
		case data, ok := <-c.msgBuffChan: //有数据要写给客户端
			if ok {
				if _, err := c.Conn.Write(data); err != nil {
					c.logging.Error("Send Buff Data error:, %v Conn Writer exit", err)
					return
				}
			} else {
				c.logging.Debug("msgBuffChan is Closed")
				break
			}
		case <-c.ctx.Done():
			return
		}
	}
}
func (c *Connection) startRead() {
	c.logging.Debug("[Reader Goroutine is running]")
	defer c.logging.Debug("%s [conn Writer exit!]", c.RemoteAddr().String())
	defer c.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			// 创建拆包解包的对象
			dp := NewDataPack()

			//读取客户端的Msg head
			headData := make([]byte, dp.GetHeadLen())
			if _, err := io.ReadFull(c.Conn, headData); err != nil {
				c.logging.Error("read msg head error: %v", err)
				return
			}

			//拆包，得到msgid 和 datalen 放在msg中
			msg, err := dp.Unpack(headData)
			if err != nil {
				c.logging.Error("unpack error: %v", err)
				return
			}

			//根据 dataLen 读取 data，放在msg.Data中
			var data []byte
			if msg.GetDataLen() > 0 {
				data = make([]byte, msg.GetDataLen())
				if _, err := io.ReadFull(c.Conn, data); err != nil {
					c.logging.Error("read msg data error: %v", err)
					return
				}
			}
			msg.SetData(data)
		}
	}
}

func (c *Connection) Start() {
	c.ctx, c.ctxCancel = context.WithCancel(context.Background())
	go c.startRead()
	go c.startWrite()
	c.Server.CallOnConnStart(c)
}

func (c *Connection) Stop() {
	c.Lock()
	defer c.Unlock()

	//如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用
	c.Server.CallOnConnClose(c)

	//如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// 关闭socket链接
	_ = c.Conn.Close()
	//关闭Writer
	c.ctxCancel()

	//将链接从连接管理器中删除
	c.Server.GetConnMgr().Remove(c)

	//关闭该链接全部管道
	close(c.msgBuffChan)
	close(c.msgChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetConnID() uint {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.RemoteAddr()
}

func (c *Connection) SendMsg(msgId uint, data []byte) error {
	c.RLock()
	if c.isClosed == true {
		c.RUnlock()
		return errors.New("connection closed when send msg")
	}
	c.RUnlock()

	//将data封包，并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		c.logging.Error("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}

	//写回客户端
	c.msgChan <- msg

	return nil
}

func (c *Connection) SendBuffMsg(msgId uint, data []byte) error {
	c.RLock()
	if c.isClosed == true {
		c.RUnlock()
		return errors.New("Connection closed when send buff msg")
	}
	c.RUnlock()

	//将data封包，并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		c.logging.Error("Pack error msg id = %v", msgId)
		return errors.New("Pack error msg ")
	}

	//写回客户端
	c.msgBuffChan <- msg

	return nil
}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Unlock()
	defer c.propertyLock.Unlock()
	c.property[key] = value
}

func (c *Connection) GetProperty(key string) (interface{}, error) {
	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found")
	}
}

func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
