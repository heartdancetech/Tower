package tower

import (
	"fmt"
	"net"
)

type BootStraper interface {
	Listen()
	Stop()
}

type bootStrap struct {
	*Config
}

func NewBootStrap(config *Config) *bootStrap {
	if config.Logger == nil {
		config.Logger = DefaultLogging
	}

	if config.IP == "" {
		config.IP = "0.0.0.0"
	}

	if config.Port == 0 {
		config.Port = 8999
	}

	return &bootStrap{
		config,
	}
}

func (bs *bootStrap) Listen() {
	bs.Logger.Debug("Server listenner at IP: %v, Port %v, is starting\n", bs.IP, bs.Port)
	go func() {

		addr, err := net.ResolveTCPAddr("", fmt.Sprintf("%s:%d", bs.IP, bs.Port))
		if err != nil {
			bs.Logger.Debug("resolve tcp addr err: %v", err)
			return
		}

		//2 监听服务器地址
		listenner, err := net.ListenTCP("", addr)
		if err != nil {
			bs.Logger.Debug("listen %s error: %v", bs.Port, err)
			return
		}

		//已经监听成功
		bs.Logger.Debug("start Zinx server %s succ, now listenning...", bs.Name)
		//3 启动server网络连接业务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			//3.2 TODO Server.Start() 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			//3.3 TODO Server.Start() 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
		}
	}
}
