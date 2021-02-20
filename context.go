package tower

type Context struct {
	conn    Connectioner
	msgId   uint32
	message []byte
}

func (ctx *Context) GetConnection() Connectioner {
	return ctx.conn
}

func (ctx *Context) GetMsgId() uint32 {
	return ctx.msgId
}

func (ctx *Context) GetMsgData() []byte {
	return ctx.message
}
