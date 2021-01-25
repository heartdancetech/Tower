package tower

import "sync"

type ConnManager interface {
	Add(conn Connectioner)                 //添加链接
	Remove(conn Connectioner)              //删除连接
	Get(connID uint) (Connectioner, error) //利用ConnID获取链接
	Len() int                              //获取当前连接
	ClearConn()                            //删除并停止所有链接
}

type ConnManage struct {
	connections map[uint]Connectioner // 连接管理
	connLock    sync.RWMutex          //读写连接的读写锁
}

func NewConnManage() *ConnManage {
	return &ConnManage{
		connections: make(map[uint]Connectioner),
	}
}

func (c *ConnManage) Add(conn Connectioner) {
	panic("implement me")
}

func (c *ConnManage) Remove(conn Connectioner) {
	panic("implement me")
}

func (c *ConnManage) Get(connID uint) (Connectioner, error) {
	panic("implement me")
}

func (c *ConnManage) Len() int {
	panic("implement me")
}

func (c *ConnManage) ClearConn() {
	panic("implement me")
}
