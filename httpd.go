package ant

import (
	"fmt"
	"net/http"
)

//var e *Engine

//统一上下文
type Context struct {
	Request  *http.Request
	Write    http.ResponseWriter
	index    int8
	handlers []HandlerFunc
}

//执行下一个handler
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

//定义自己的 HandlerFunc
type HandlerFunc func(*Context)

//实现自己的 HandlerFunc 的 ServeHTTP 方法
func (f HandlerFunc) ServeHTTP(c *Context) {
	f(c)
}

//路由
type Router struct {
	method  string
	path    string
	handles []HandlerFunc
}

type Engine struct {
	routers map[string]*Router //Router的指针类型
}

//AddRouter 方法属于 AntEngine 类型对象中的方法
func (e *Engine) AddRouter(method string, path string, h []HandlerFunc) {
	e.routers[method+"_"+path] = &Router{
		method:  method,
		path:    path,
		handles: h,
	}
}

//Get 方法属于 Engine 类型对象中的方法
func (e *Engine) Get(path string, h ...HandlerFunc) {
	e.AddRouter("GET", path, h)
}

//POST 方法属于 Engine 类型对象中的方法
func (e *Engine) Post(path string, h ...HandlerFunc) {
	e.AddRouter("POST", path, h)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("HTTP/1.0 200 OK \r\n Hello ant-go"))

	//c := &AntContext{
	//	Request: r,
	//	Write: w,
	//}
	//
	//testHandlerFunc(c)

	method := r.Method
	path := r.RequestURI

	router := e.routers[method+"_"+path]

	if router == nil {
		w.Write([]byte("404 你访问的页面不存在 uri:" + path))
		return
	}

	fmt.Printf("=== e %v router: %v \n", e, router)

	c := &Context{
		Request:  r,
		Write:    w,
		index:    -1,
		handlers: router.handles,
	}

	c.Next()

}

func New() *Engine {
	//e := &Engine{}
	//e.routers = make(map[string]*Router)
	//return e

	checkConfig()

	return &Engine{
		routers: make(map[string]*Router),
	}
}

func (e *Engine) Run() {
	http.ListenAndServe(":8080", e)
}
