package tower

type Config struct {
	Name             string // server name
	IP               string // server listen ip
	IPVersion        string // ip version
	Port             int    // server listen port
	MaxPacketSize    uint32 // server accpect max packet size
	MaxConn          int    // server accpect max connection count
	WorkerPoolSize   uint32 // work pool
	MaxWorkerTaskLen uint32 // 业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32 // SendBuffMsg发送消息的缓冲最大长度
}

func (c *Config) setDefault() {
	if c.IP == "" {
		c.IP = "0.0.0.0"
	}
	if c.IPVersion == "" {
		c.IPVersion = "tcp4"
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
