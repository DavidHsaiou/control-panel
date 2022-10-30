package lib

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"control-panel/server/lib/extension"
	"control-panel/server/lib/inf"
)

type serverGin struct {
	engine *gin.Engine
}

func NewDefaultGin() inf.IServer {
	engine := gin.Default()

	return &serverGin{
		engine: engine,
	}
}

func NewDefaultGinWithFx(lc fx.Lifecycle, handlers extension.ServerParams) inf.IServer {
	engine := NewDefaultGin()

	engine.Register(handlers.Handlers)

	lc.Append(fx.Hook{
		OnStart: engine.Start,
	})

	return engine
}

func (s *serverGin) Register(handlers []inf.IHandler) {
	for _, handler := range handlers {
		switch handler.GetMethod() {
		case http.MethodGet:
			ginHandler := newHandlerGin(handler)
			s.engine.GET(ginHandler.GetRouter(), ginHandler.Execute)
		}
	}
}

func (s *serverGin) Start(_ context.Context) error {
	go func() {
		err := s.engine.Run()
		if err != nil {
			panic(err)
			// return err
		}
	}()

	return nil
}

func (s *serverGin) Stop() {
	// TODO implement me
	panic("implement me")
}
