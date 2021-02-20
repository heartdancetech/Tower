package tower

import (
	"strconv"
	"sync"
)

type Router interface {
	AddRoute(msgId uint32, handleFunc func(ctx *Context))
	doHandler(ctx *Context)
}

type route struct {
	routes map[uint32]func(ctx *Context)
	sync.Mutex
}

func newRoute() *route {
	return &route{routes: make(map[uint32]func(ctx *Context))}
}

func (r *route) AddRoute(msgId uint32, handleFunc func(ctx *Context)) {
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
		return
	}
	handler(ctx)
}
