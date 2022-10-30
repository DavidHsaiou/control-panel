package inf

type IHandler interface {
	GetMethod() string
	GetRouter() string
	Execute(ctx IContext)
}
