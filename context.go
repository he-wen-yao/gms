package gms

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// 框架中请求上下文
type context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
}

func NewContext(w http.ResponseWriter, req *http.Request) *context {
	return &context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Params: map[string]string{},
		Method: req.Method,
	}
}

func (c *context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 获取路径参数
func (c *context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 返回 String 数据
func (c *context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 返回 JSON 数据
func (c *context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 返回 字节流数据
func (c *context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 返回 HTML 数据
func (c *context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
