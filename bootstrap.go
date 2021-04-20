package tower

import (
	"fmt"
	"net"
)

type BootStraper interface {
	Listen()                                              // start server
	Stop()                                                // stop server
	GetConnMgr() ConnManager                              // get connection manager
	SetOnConnStart(func(conn Connectioner))               // set hook func when client connect server
	SetOnConnClose(func(conn Connectioner))               // set hook func when client disconnect server
	CallOnConnStart(conn Connectioner)                    // call OnConnStart hook func
	CallOnConnClose(conn Connectioner)                    // call OnConnStop hook func
	SetLogging(Logger)                                    // set logging
	AddRoute(msgId uint32, handleFunc func(ctx *Context)) // add route
	Logging() Logger                                      // logging
	GetConfig() *Config                                   // get server global config
}

type bootStrap struct {
	*Config
	logging     Logger
	ConnMgr     ConnManager
	router      router
	OnConnStart func(conn Connectioner)
	OnConnClose func(conn Connectioner)
}

func NewBootStrap(config *Config) BootStraper {
	if config == nil {
		config = NewConfig()
	}
	config.check()

	return &bootStrap{
		Config:      config,
		ConnMgr:     NewConnManage(),
		logging:     defaultLogging,
		router:      newRoute(),
		OnConnStart: nil,
		OnConnClose: nil,
	}
}

// Listen start server,listen port
func (bs *bootStrap) Listen() {
	bs.logging.Debug("Server listener at IP: %v, Port %v, is starting\n", bs.IP, bs.Port)
	addr, err := net.ResolveTCPAddr(bs.IPVersion, fmt.Sprintf("%s:%d", bs.IP, bs.Port))
	if err != nil {
		bs.logging.Error("resolve tcp addr err: %v", err)
		return
	}

	// listen server addr and port
	listener, err := net.ListenTCP(bs.IPVersion, addr)
	if err != nil {
		bs.logging.Error("listen %s error: %v", bs.Port, err)
		return
	}
	bs.logging.Debug("start server %s success, now listening...", bs.Name)
	var cid uint32 = 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			bs.logging.Error("Accept err: %v", err)
			return
		}
		bs.logging.Debug("Get conn remote addr = %v", conn.RemoteAddr().String())

		// set server's max conn accept number, if greater than config's value then close this conn
		if bs.ConnMgr.Len() >= bs.Config.MaxConn {
			_ = conn.Close()
			return
		}

		//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
		dealConn := NewConnection(bs, conn, cid, bs.router)
		cid++

		//3.4 启动当前链接的处理业务
		go dealConn.Start()
	}
}

// Stop stop server
func (bs *bootStrap) Stop() {
	bs.ConnMgr.ClearConn()
	bs.logging.Info("server stop")
}

// GetConnMgr get connection manager
func (bs *bootStrap) GetConnMgr() ConnManager {
	return bs.ConnMgr
}

// SetOnConnStart set hook func when client connect server
func (bs *bootStrap) SetOnConnStart(hookFunc func(conn Connectioner)) {
	bs.OnConnStart = hookFunc
}

// SetOnConnClose set hook func when client disconnect server
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

// SetLogging set customer logging
func (bs *bootStrap) SetLogging(logging Logger) {
	bs.logging = logging
}

// Logging get bootstrap's logging
func (bs *bootStrap) Logging() Logger {
	return bs.logging
}

// GetConfig get config pointer
func (bs *bootStrap) GetConfig() *Config {
	return bs.Config
}

// AddRoute add route
func (bs *bootStrap) AddRoute(msgId uint32, handleFunc func(ctx *Context)) {
	bs.router.addRoute(msgId, handleFunc)
}
