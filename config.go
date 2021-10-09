package tower

type Config struct {
	Name             string // server name
	IP               string // server listen ip
	IPVersion        string // ip version
	Port             int    // server listen port
	MaxPacketSize    uint32 // server accept max packet size
	MaxConn          int    // server accept max connection count
	WorkerPoolSize   uint32 // work pool
	MaxWorkerTaskLen uint32 // 业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32 // SendBuffMsg发送消息的缓冲最大长度
}

// NewConfig new config instance use default config option
func NewConfig() *Config {
	return &Config{
		Name:             "Tower",
		IP:               "0.0.0.0",
		IPVersion:        "tcp4",
		Port:             8999,
		MaxPacketSize:    0,
		MaxConn:          1024,
		WorkerPoolSize:   0,
		MaxWorkerTaskLen: 0,
		MaxMsgChanLen:    1024,
	}
}

// check if bootstrap init without custom config
// set default config option
func (c *Config) check() {
	if c.IP == "" {
		c.IP = "0.0.0.0"
	}
	if c.IPVersion == "" {
		c.IPVersion = "tcp4"
	}

	if c.Port == 0 {
		c.Port = 8999
	}

	if c.MaxConn == 0 {
		c.MaxConn = 1024
	}

	if c.MaxMsgChanLen == 0 {
		c.MaxMsgChanLen = 1024
	}
}
