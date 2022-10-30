package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"control-panel/server/lib/inf"
)

type Ping struct {
}

func NewHandlerPing() inf.IHandler {
	return &Ping{}
}

func (p Ping) GetRouter() string {
	return "/ping"
}

func (p Ping) GetMethod() string {
	return http.MethodGet
}

func (p Ping) Execute(ctx inf.IContext) {
	ctx.JSONResult(http.StatusOK, gin.H{
		"message": "pong",
	})
}
