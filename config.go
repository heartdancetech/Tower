package tower

import "github.com/go-tower/tower/logger"

type Config struct {
	Name string
	IP   string
	Port int

	MaxPacketSize    uint32 //都需数据包的最大值
	MaxConn          int    //当前服务器主机允许的最大链接个数
	WorkerPoolSize   uint32 //业务工作Worker池的数量
	MaxWorkerTaskLen uint32 //业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32 //SendBuffMsg发送消息的缓冲最大长度

	Logging logger.Logger
}

func (c *Config) setDefault() {
	if c.Logging == nil {
		c.Logging = logger.DefaultLogging
	}

	if c.IP == "" {
		c.IP = "0.0.0.0"
	}

	if c.Port == 0 {
		c.Port = 8999
	}

	if c.MaxPacketSize == 0 {
		c.MaxPacketSize = 4096
	}

	if c.MaxConn == 0 {
		c.MaxConn = 1024
	}

	if c.MaxMsgChanLen == 0 {
		c.MaxMsgChanLen = 1024
	}
}
