package httpd

import (
	"fmt"
	"net/http"
)

//https://www.cnblogs.com/ham-731/p/12637656.html

//统一上下文
type AntContext struct {
	Request *http.Request
	Write   http.ResponseWriter

	index    int8
	handlers []AntHandlerFunc
}

//执行下一个handler
func (c *AntContext) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

//定义自己的 AntHandlerFunc
type AntHandlerFunc func(*AntContext)

//实现自己的 AntHandlerFunc 的 ServeHTTP 方法
func (f AntHandlerFunc) ServeHTTP(c *AntContext) {
	f(c)
}

//路由
type Router struct {
	method  string
	path    string
	handles []AntHandlerFunc
}

//给 AntEngine 的  routers 里面添加路由
//AddRouter 方法属于 AntEngine 类型对象中的方法
func (e *AntEngine) AddRouter(method string, path string, h []AntHandlerFunc) {
	e.routers[method+"_"+path] = &Router{
		method:  method,
		path:    path,
		handles: h,
	}
}

//Get 方法属于 AntEngine 类型对象中的方法
func (e *AntEngine) Get(path string, h ...AntHandlerFunc) {
	e.AddRouter("GET", path, h)
}

//POST 方法属于 AntEngine 类型对象中的方法
func (e *AntEngine) Post(path string, h ...AntHandlerFunc) {
	e.AddRouter("POST", path, h)
}

//自己的引擎
type AntEngine struct {
	routers map[string]*Router
}

func New() *AntEngine {
	return &AntEngine{
		routers: make(map[string]*Router),
	}
}

func (e *AntEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	fmt.Printf("=== e %v router: %v \n", e, router)

	c := &AntContext{
		Request:  r,
		Write:    w,
		index:    -1,
		handlers: router.handles,
	}

	c.Next()

}

func (e *AntEngine) Run() {
	http.ListenAndServe(":8080", e)
}
