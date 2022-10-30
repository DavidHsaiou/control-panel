package lib

import (
	"github.com/gin-gonic/gin"

	"control-panel/server/lib/inf"
)

type handlerGin struct {
	handler inf.IHandler
}

func newHandlerGin(handler inf.IHandler) *handlerGin {
	return &handlerGin{
		handler: handler,
	}
}

func (h handlerGin) GetMethod() string {
	return h.handler.GetMethod()
}

func (h handlerGin) GetRouter() string {
	return h.handler.GetRouter()
}

func (h handlerGin) Execute(ctx *gin.Context) {
	h.handler.Execute(NewContextGin(ctx))
}
