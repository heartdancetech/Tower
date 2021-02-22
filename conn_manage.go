package tower

import (
	"errors"
	"sync"
)

type ConnManager interface {
	Add(conn Connectioner)                   // add connection
	Remove(conn Connectioner)                // delete connection
	Get(connID uint32) (Connectioner, error) // get connection by connection id
	Len() int                                // get connections's count
	ClearConn()                              // stop all connections, then delete them
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

// Add add connection
func (c *ConnManage) Add(conn Connectioner) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn
}

// Remove delete connection
func (c *ConnManage) Remove(conn Connectioner) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	delete(c.connections, conn.GetConnID())
}

// Get get connection by connection id
func (c *ConnManage) Get(connID uint32) (Connectioner, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	if conn, ok := c.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

// Len get connections's count
func (c *ConnManage) Len() int {
	return len(c.connections)
}

// ClearConn stop all connections, then delete them
func (c *ConnManage) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	//停止并删除全部的连接信息
	for connID, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connID)
	}
}
