package tower

import "sync"

type ConnManager interface {
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
