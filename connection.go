package tower

import (
	"context"
	"net"
	"sync"
)

type Connectioner interface {
	GetConnID() uint

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

func (c *Connection) SetProperty(key string, value interface{}) {
	panic("implement me")
}

func (c *Connection) GetProperty(key string) (interface{}, error) {
	panic("implement me")
}

func (c *Connection) RemoveProperty(key string) {
	panic("implement me")
}

func (c *Connection) GetConnID() uint {
	return c.ConnID
}
