package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	router map[string]*Tree
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
func (c *Core) Get(url string, handler ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	if err := c.router["GET"].addRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Post(url string, handler ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	if err := c.router["POST"].addRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Put(url string, handler ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	if err := c.router["PUT"].addRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Delete(url string, handler ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	if err := c.router["DELETE"].addRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)

	router := c.FindRouterByRequest(request)
	if router == nil {
		ctx.Json(404, "not found")
		return
	}
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
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
func (c *Core) FindRouterByRequest(request *http.Request) ControllerHandler {
	uri := request.URL.Path
	method := request.Method
	upperUri := strings.ToUpper(uri)
	upperMethod := strings.ToUpper(method)
	if methodHandlerTree, ok := c.router[upperMethod]; ok {
		return methodHandlerTree.FindHandler(upperUri)
	}
	return nil
}
