package tower

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
}

func (bs *bootStrap) Stop() {

}
