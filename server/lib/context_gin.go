package lib

import "github.com/gin-gonic/gin"

type ContextGin struct {
	context *gin.Context
}

func (c ContextGin) JSONResult(status int, msg map[string]any) {
	c.context.JSON(status, msg)
}

func NewContextGin(ctx *gin.Context) *ContextGin {
	return &ContextGin{
		context: ctx,
	}
}
