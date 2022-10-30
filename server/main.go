package main

import (
	"go.uber.org/fx"

	"control-panel/server/handlers"
	"control-panel/server/lib"
	"control-panel/server/lib/extension"
	"control-panel/server/lib/inf"
)

func main() {
	fx.New(
		fx.Provide(
			extension.AsHandler(handlers.NewHandlerPing),
		),
		fx.Provide(lib.NewDefaultGinWithFx),
		fx.Invoke(func(server inf.IServer) {}),
	).Run()
}
