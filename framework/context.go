package framework

import (
	"context"
	"net/http"
)

type Context struct {
	request        http.Request
	responseWriter http.ResponseWriter
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}
func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}
