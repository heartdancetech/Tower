[![GoDoc][1]][2]
[![Coverage Status][3]][4]

[1]: https://godoc.org/github.com/go-tower/tower?status.svg

[2]: https://pkg.go.dev/github.com/go-tower/tower

[3]: https://coveralls.io/repos/github/go-tower/tower/badge.svg?branch=master

[4]: https://coveralls.io/github/go-tower/tower?branch=master

# Tower

`Tower` is a tcp connection manager server.

## Quick Start

### Get Dep

```shell
go get -u github.com/go-tower/tower
```

### Start Server

```go
package main

import (
	"fmt"
	"github.com/go-tower/tower"
)

func main() {
	SockServer := tower.NewBootStrap(&tower.Config{})
	SockServer.SetOnConnStart(func(conn tower.Connectioner) {

	})
	SockServer.SetOnConnClose(func(conn tower.Connectioner) {

	})
	SockServer.AddRoute(0, func(ctx *tower.Context) {
		fmt.Println("hello world")
	})
	SockServer.Listen()
}
```

## Tower API

### Tower Config

```go
package tower

type Config struct {
	Name          string // server name
	IP            string // server listen ip
	IPVersion     string // ip version
	Port          int    // server listen port
	MaxPacketSize uint32 // server accept max packet size
	MaxConn       int    // server accept max connection count
}
```

### BootStrap mod

```go
func NewBootStrap(config *Config) BootStraper {}
```

#### start server

```go 
func (bs *bootStrap) Listen() {}
```

#### stop server

```go 
func (bs *bootStrap) Stop() {}
```
