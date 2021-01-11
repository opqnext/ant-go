package httpd

import "net/http"

//https://www.cnblogs.com/ham-731/p/12637656.html

type Engine struct {
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP/1.0 200 OK \r\n Hello ant-go"))
}

func Run() {
	http.ListenAndServe(":8080", &Engine{})
}
