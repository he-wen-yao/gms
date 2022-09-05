package gms

import (
	"net/http"
	"strings"

	"github.com/he-wen-yao/gms/ds"
)

type Router struct {
	roots    map[string]*ds.Node
	handlers map[string]HandlerFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

func NewRouter() *Router {
	return &Router{
		roots:    make(map[string]*ds.Node),
		handlers: make(map[string]HandlerFunc),
	}
}

// Only one * is allowed
func ParsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *Router) AddRoute(method string, pattern string, handler HandlerFunc) {
	parts := ParsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &ds.Node{}
	}
	r.roots[method].Insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) GetRoute(method string, path string) (*ds.Node, map[string]string) {
	searchParts := ParsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.Search(searchParts, 0)

	if n != nil {
		parts := ParsePattern(n.Pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *Router) Handle(c *Context) {
	n, params := r.GetRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.Part
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

func (r *Router) handle(c *Context) {
	n, params := r.GetRoute(c.Method, c.Path)
	if n != nil {
		key := c.Method + "-" + n.Pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
