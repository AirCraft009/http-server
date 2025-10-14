package handlers

import (
	"http-server/parser"
	"http-server/server"
	"net/http"
)

func Homehandler(req *parser.Request) (res *server.Response) {
	res = server.NewResponse(req)
	res.StatusCode = http.StatusOK
	res.AddHeader("Content-Type", "text/html; charset=utf-8")
	res.AddHeader("Connection", "keep-alive")
	res.SetBody([]byte("Hello World"))
	return res
}
