package tower

import (
	"errors"
	"sync"
)

type ConnManager interface {
	Add(conn Connectioner)                   // add connection
	Remove(conn Connectioner)                // delete connection
	Get(connID uint32) (Connectioner, error) // get connection by connection id
	Len() int                                // get connections' count
	ClearConn()                              // stop all connections, then delete them
}

type ConnManage struct {
	connections map[uint32]Connectioner // connection map
	connLock    sync.RWMutex            // connection map RWLock
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

// Len get connections' count
func (c *ConnManage) Len() int {
	return len(c.connections)
}

// ClearConn stop all connections, then delete them
func (c *ConnManage) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	// stop and clear all conn
	for connID, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connID)
	}
}
