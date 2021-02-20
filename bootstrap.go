package tower

import (
	"fmt"
	"github.com/go-tower/tower/logger"
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
	getConfig() *Config                                   // get server global config
	Logging() logger.Logger                               // get logging
	AddRoute(msgId uint32, handleFunc func(ctx *Context)) // add route
}

type bootStrap struct {
	*Config
	logging     logger.Logger
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
		logging:     logger.DefaultLogging,
		Router:      newRoute(),
		OnConnStart: nil,
		OnConnClose: nil,
	}
}

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

func (bs *bootStrap) Stop() {
	bs.ConnMgr.ClearConn()
	bs.logging.Info("server stop")
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

func (bs *bootStrap) Logging() logger.Logger {
	return bs.logging
}

func (bs *bootStrap) getConfig() *Config {
	return bs.Config
}

func (bs *bootStrap) AddRoute(msgId uint32, handleFunc func(ctx *Context)) {
	bs.Router.AddRoute(msgId, handleFunc)
}
