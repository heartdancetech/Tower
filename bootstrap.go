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
	SetLogging(logger)                                    // set logging
	AddRoute(msgId uint32, handleFunc func(ctx *Context)) // add route
	Logging() logger                                      // get logging
	getConfig() *Config                                   // get server global config
}

type bootStrap struct {
	*Config
	logging     logger
	ConnMgr     ConnManager
	Router      Router
	OnConnStart func(conn Connectioner)
	OnConnClose func(conn Connectioner)
}

func NewBootStrap(config *Config) BootStraper {
	if config == nil {
		config = &Config{}
	}
	config.setDefault()

	return &bootStrap{
		Config:      config,
		ConnMgr:     NewConnManage(),
		logging:     defaultLogging,
		Router:      newRoute(),
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

	// 监听服务器地址
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

		//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
		if bs.ConnMgr.Len() >= bs.Config.MaxConn {
			_ = conn.Close()
			return
		}

		//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
		dealConn := NewConnection(bs, conn, cid, bs.Router)

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

// SetOnConnStart set func on client start connect
func (bs *bootStrap) SetOnConnStart(hookFunc func(conn Connectioner)) {
	bs.OnConnStart = hookFunc
}

// SetOnConnStart set func on client close connect
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

// SetLogging set custmer logging
func (bs *bootStrap) SetLogging(logging logger) {
	bs.logging = logging
}

func (bs *bootStrap) Logging() logger {
	return bs.logging
}

func (bs *bootStrap) getConfig() *Config {
	return bs.Config
}

// AddRoute add route
func (bs *bootStrap) AddRoute(msgId uint32, handleFunc func(ctx *Context)) {
	bs.Router.AddRoute(msgId, handleFunc)
}
