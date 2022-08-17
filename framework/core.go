package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router      map[string]*Tree
	middlewares []ControllerHandler
}

func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{
		router: router,
	}
}
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}
func (c *Core) Get(url string, handlers ...ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	if err := c.router["GET"].addRouter(url, handlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Post(url string, handlers ...ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	if err := c.router["POST"].addRouter(url, handlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Put(url string, handlers ...ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	if err := c.router["PUT"].addRouter(url, handlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	if err := c.router["DELETE"].addRouter(url, handlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)

	node := c.FindRouteNodeByRequest(request)
	if node == nil {
		ctx.SetStatus(404).Json("not found")
		return
	}
	// 设置context中的handlers字段
	ctx.SetHandlers(node.handlers)
	params := node.parseParamsFromEndNode(request.URL.Path)
	ctx.SetParams(params)
	if err := ctx.Next(); err != nil {
		ctx.SetStatus(500).Json("inner error")
		return
	}

	// 一个简单的路由选择器，这里直接写死为测试路由foo
	//for _, v := range c.router {
	//	if v == nil {
	//		break
	//	}
	//	v(ctx)
	//}
	//router := c.router["foo"]
	//if router == nil {
	//	return
	//}
	log.Println("core.router")

	//router(ctx)
}

func (c *Core) FindRouteNodeByRequest(request *http.Request) *node {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)
	if methodHandlerTree, ok := c.router[upperMethod]; ok {
		return methodHandlerTree.root.matchNode(uri)
	}
	return nil
}
