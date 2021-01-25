package tower

import (
	"context"
	"net"
	"sync"
)

type Connectioner interface {
	Start()
	Stop()

	//从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	//获取当前连接ID
	GetConnID() uint
	//获取远程客户端地址信息
	RemoteAddr() net.Addr

	//直接将Message数据发送数据给远程的TCP客户端(无缓冲)
	SendMsg(msgId uint, data []byte) error
	//直接将Message数据发送给远程的TCP客户端(有缓冲)
	SendBuffMsg(msgId uint, data []byte) error

	//设置链接属性
	SetProperty(key string, value interface{})
	//获取链接属性
	GetProperty(key string) (interface{}, error)
	//移除链接属性
	RemoveProperty(key string)
}

type Connection struct {
	Server    BootStraper
	Conn      *net.TCPConn
	ConnID    uint
	ctx       context.Context
	ctxCancel context.CancelFunc

	sync.RWMutex
	//链接属性
	property map[string]interface{}
	////保护当前property的锁
	propertyLock sync.Mutex
	//当前连接的关闭状态
	isClosed bool
}

func NewConnection(server BootStraper, conn *net.TCPConn, connID uint) *Connection {
	return nil
}

func (c *Connection) Start() {
	panic("implement me")
}

func (c *Connection) Stop() {
	panic("implement me")
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	panic("implement me")
}
func (c *Connection) GetConnID() uint {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	panic("implement me")
}

func (c *Connection) SendMsg(msgId uint, data []byte) error {
	panic("implement me")
}

func (c *Connection) SendBuffMsg(msgId uint, data []byte) error {
	panic("implement me")
}

func (c *Connection) SetProperty(key string, value interface{}) {
	panic("implement me")
}

func (c *Connection) GetProperty(key string) (interface{}, error) {
	panic("implement me")
}

func (c *Connection) RemoveProperty(key string) {
	panic("implement me")
}
