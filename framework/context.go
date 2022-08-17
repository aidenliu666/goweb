package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handlers       []ControllerHandler
	index          int //记录当前请求调用到调用链的哪个节点
	hasTimeout     bool
	writeMux       *sync.Mutex
	params         map[string]string // url路由匹配的参数
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writeMux:       &sync.Mutex{},
		index:          -1,
	}

}

// Next https://static001.geekbang.org/resource/image/73/3c/73a80752cf6d94b90febd2e23e80bc3c.jpg?wh=1920x915
// Next 函数是整个链路执行的重点，要好好理解，它通过维护 Context 中的一个下标，来控制链路移动，这个下标表示当前调用 Next 要执行的控制器序列。
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}
func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writeMux
}
func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}
func (ctx *Context) GetReponseWriter() http.ResponseWriter {
	return ctx.responseWriter
}
func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}
func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// 设置参数
func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}
func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}
func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}
func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}
func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}
func (ctx *Context) Value(key any) any {
	return ctx.BaseContext().Value(key)
}
