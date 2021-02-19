package tower

import (
	"fmt"
	"strconv"
	"sync"
)

type Router interface {
	AddRoute(msgId uint, handleFunc func(ctx *Context))
}

type route struct {
	routes map[uint]func(ctx *Context)
	sync.Mutex
}

func newRoute() *route {
	return &route{routes: make(map[uint]func(ctx *Context))}
}

func (r *route) AddRoute(msgId uint, handleFunc func(ctx *Context)) {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.routes[msgId]; ok {
		panic("repeated api , msgId = " + strconv.FormatUint(uint64(msgId), 10))
	}
	r.routes[msgId] = handleFunc
}

func (r *route) doHandler(ctx *Context) {
	handler, ok := r.routes[ctx.GetMsgId()]
	if !ok {
		fmt.Println("api msgId = ", ctx.GetMsgId(), " is not FOUND!")
		return
	}
	handler(ctx)
}
