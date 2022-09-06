package gms

import (
	"net/http"
	"strings"
)

// HandlerFunc 定义处理请求的函数类型
type HandlerFunc func(*Context)

// Gms implement the interface of ServeHTTP
type Gms struct {
	RouterGroup                // 通过内嵌的方式使得 Gms 继承 RouterGroup
	groups      []*RouterGroup // 存储分组信息
}

// NewGms 返回 Gms 实例
func NewGms() (g *Gms) {
	g = &Gms{}
	g.RouterGroup = RouterGroup{gms: g}
	g.groups = []*RouterGroup{}
	return
}

// ServeHTTP http.Handler 接口
func (g *Gms) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range g.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := NewContext(w, req)
	c.handlers = middlewares
	g.router.Handle(c)
}

// Run 运行 Gms
func (g *Gms) Run(addr string) (err error) {
	return http.ListenAndServe(addr, g)
}
