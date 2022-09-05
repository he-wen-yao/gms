// 路由组
package gms

import "log"

type RouterGroup struct {
	prefix      string
	router      *Router
	middlewares []HandlerFunc
	gms         *Gms
}

// Group 创建新的路由组实例，所有派生的路由共享一个 Gms 实例
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	g := group.gms
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		gms:    g,
	}
	// 添加新的路由分组
	g.groups = append(g.groups, newGroup)
	return newGroup
}

// addRoute 向 gms 路由树中添加新的路由
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.gms.router.AddRoute(method, pattern, handler)
}

// GET 添加 GET 请求
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 添加 POST 请求
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Use 路由组应用中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
