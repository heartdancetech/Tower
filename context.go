package tower

type Context struct {
	conn    Connectioner
	msgId   uint
	message []byte
}

func (ctx *Context) GetConnection() Connectioner {
	return ctx.conn
}

func (ctx *Context) GetMsgId() uint {
	return ctx.msgId
}

func (ctx *Context) GetMsgData() []byte {
	return ctx.message
}
