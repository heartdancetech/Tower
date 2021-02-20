package tower

import (
	"errors"
	"sync"
)

type ConnManager interface {
	Add(conn Connectioner)                   //添加链接
	Remove(conn Connectioner)                //删除连接
	Get(connID uint32) (Connectioner, error) //利用ConnID获取链接
	Len() int                                //获取当前连接
	ClearConn()                              //删除并停止所有链接
}

type ConnManage struct {
	connections map[uint32]Connectioner // 连接管理
	connLock    sync.RWMutex            //读写连接的读写锁
}

func NewConnManage() *ConnManage {
	return &ConnManage{
		connections: make(map[uint32]Connectioner),
	}
}

func (c *ConnManage) Add(conn Connectioner) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn
}

func (c *ConnManage) Remove(conn Connectioner) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	delete(c.connections, conn.GetConnID())
}

func (c *ConnManage) Get(connID uint32) (Connectioner, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	if conn, ok := c.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

func (c *ConnManage) Len() int {
	return len(c.connections)
}

func (c *ConnManage) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	//停止并删除全部的连接信息
	for connID, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connID)
	}
}
