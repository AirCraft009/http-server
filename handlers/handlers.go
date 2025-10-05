package handlers

import (
	"http-server/parser"
	"http-server/server"
	"net/http"
	"strconv"
)

func Homehandler(req *parser.Request) (res *server.Response) {
	res = server.NewResponse(req)
	res.HTTPType = req.HTTPType
	res.StatusCode = http.StatusOK
	res.Headers = make(map[string]string)
	res.Headers["Content-Type"] = "text/html; charset=utf-8"
	res.Headers["Connection"] = "keep-alive"
	res.Body = []byte("Hello World")
	res.Headers["Content-Lenght"] = strconv.Itoa(len(res.Body))
	return res
}

func Http404Handler(req *parser.Request) (res *server.Response) {
	res = &server.Response{}
	res.HTTPType = req.HTTPType
	res.StatusCode = http.StatusNotFound
	res.Headers = make(map[string]string)
	res.Headers["Content-Type"] = "text/html; charset=utf-8"
	res.Headers["Connection"] = "close"
	res.Body = []byte()
}
