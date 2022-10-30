package inf

type IContext interface {
	JSONResult(status int, msg map[string]any)
}
