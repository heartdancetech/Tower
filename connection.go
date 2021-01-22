package tower

import "context"

type Connectioner interface {
}

type Connection struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
}
