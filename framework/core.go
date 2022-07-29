package framework

import "net/http"

type Core struct {
}

func NewCore() *Core {
	return &Core{}
}
func (c *Core) ServeHTTP(reponse http.ResponseWriter, requeset *http.Request) {
}
