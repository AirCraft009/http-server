package server

import (
	"http-server/parser"
	"http-server/reader"
	"net/http"
	"strconv"
)

func Http404Handler(req *parser.Request) (res *Response) {
	res = &Response{}
	res.HTTPType = req.HTTPType
	res.StatusCode = http.StatusNotFound
	res.Headers = make(map[string]string)
	res.Headers["Content-Type"] = "text/html; charset=utf-8"
	res.Headers["Connection"] = "close"
	data := reader.ReadFile("frontend/4044.html")
	res.Body = data
	res.Headers["Content-Lenght"] = strconv.Itoa(len(res.Body))
	return res
}
