package framework

type IGroup interface {
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)
}

type Group struct {
	core        *Core
	parent      *Group
	prefix      string
	middlewares []ControllerHandler
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		prefix: prefix,
	}

}
func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	uri = g.prefix + uri
	g.core.Get(uri, handlers...)
}
func (g *Group) Post(uri string, handlers ...ControllerHandler) {
	uri = g.prefix + uri
	g.core.Post(uri, handlers...)
}
func (g *Group) Put(uri string, handlers ...ControllerHandler) {
	uri = g.prefix + uri
	g.core.Put(uri, handlers...)
}
func (g *Group) Delete(uri string, handlers ...ControllerHandler) {
	uri = g.prefix + uri
	g.core.Delete(uri, handlers...)
}
func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}
	return append(g.parent.getMiddlewares(), g.middlewares...)
}

// 实现 Group 方法
func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}

// 注册中间件
func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}
