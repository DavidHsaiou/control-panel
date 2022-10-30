package inf

import "context"

type IServer interface {
	Register(handlers []IHandler)
	Start(ctx context.Context) error
	Stop()
}
