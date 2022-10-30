package extension

import (
	"go.uber.org/fx"

	"control-panel/server/lib/inf"
)

var (
	_handlerTags = `group:"handlers"`
)

type ServerParams struct {
	fx.In

	Handlers []inf.IHandler `group:"handlers"`
}

func AsHandler(handler any) any {
	return fx.Annotate(handler, fx.ResultTags(_handlerTags))
}

// func AsServer(server any) any {
// 	return fx.Annotate(server, fx.ParamTags(_handlerTags))
// }
