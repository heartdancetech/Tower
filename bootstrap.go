package tower

import (
	"fmt"
	"net"
)

type BootStraper interface {
	Listen()                                // 启动服务
	Stop()                                  // 关闭服务
	GetConnMgr() ConnManager                //得到链接管理
	SetOnConnStart(func(conn Connectioner)) //设置该Server的连接创建时Hook函数
	SetOnConnClose(func(conn Connectioner)) //设置该Server的连接断开时的Hook函数
	CallOnConnStart(conn Connectioner)      //调用连接OnConnStart Hook函数
	CallOnConnClose(conn Connectioner)      //调用连接OnConnStop Hook函数
	getConfig() *Config
}

type bootStrap struct {
	*Config
	ConnMgr     ConnManager
	OnConnStart func(conn Connectioner)
	OnConnClose func(conn Connectioner)
}

func NewBootStrap(config *Config) BootStraper {
	if config == nil {
		config = &Config{}
	}
	config.check()

	return &bootStrap{
		Config:      config,
		ConnMgr:     NewConnManage(),
		OnConnStart: nil,
		OnConnClose: nil,
	}
}

func (bs *bootStrap) Listen() {
	bs.Logging.Debug("Server listener at IP: %v, Port %v, is starting\n", bs.IP, bs.Port)
	addr, err := net.ResolveTCPAddr("", fmt.Sprintf("%s:%d", bs.IP, bs.Port))
	if err != nil {
		bs.Logging.Error("resolve tcp addr err: %v", err)
		return
	}

	// 监听服务器地址
	listener, err := net.ListenTCP("", addr)
	if err != nil {
		bs.Logging.Error("listen %s error: %v", bs.Port, err)
		return
	}
	bs.Logging.Debug("start server %s success, now listening...", bs.Name)
	var cid uint = 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			bs.Logging.Error("Accept err: %v", err)
			return
		}
		bs.Logging.Debug("Get conn remote addr = %v", conn.RemoteAddr().String())

		//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
		if bs.ConnMgr.Len() >= bs.Config.MaxConn {
			_ = conn.Close()
			return
		}

		//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
		dealConn := NewConnection(bs, conn, cid)

		//3.4 启动当前链接的处理业务
		go dealConn.Start()
	}
}

func (bs *bootStrap) Stop() {
	bs.ConnMgr.ClearConn()
	bs.Logging.Info("server stop")
}

func (bs *bootStrap) GetConnMgr() ConnManager {
	return bs.ConnMgr
}

func (bs *bootStrap) SetOnConnStart(hookFunc func(conn Connectioner)) {
	bs.OnConnStart = hookFunc
}

func (bs *bootStrap) SetOnConnClose(hookFunc func(conn Connectioner)) {
	bs.OnConnClose = hookFunc
}

func (bs *bootStrap) CallOnConnStart(conn Connectioner) {
	if bs.OnConnStart != nil {
		bs.OnConnStart(conn)
	}
}

func (bs *bootStrap) CallOnConnClose(conn Connectioner) {
	if bs.OnConnClose != nil {
		bs.OnConnClose(conn)
	}
}

func (bs *bootStrap) getConfig() *Config {
	return bs.Config
}
