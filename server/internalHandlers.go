package server

import (
	"fmt"
	"http-server/parser"
	"http-server/reader"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func Http404Handler(req *parser.Request) (res *Response) {
	res = NewResponse(req)
	res.StatusCode = http.StatusNotFound
	res.Headers = make(map[string]string)
	res.Headers["Content-Type"] = "text/html; charset=utf-8"
	res.Headers["Connection"] = "keep-alive"
	data, err := reader.ReadFile("frontend/404.html")
	if err != nil {
		res.StatusCode = http.StatusBadRequest
		res.Body = []byte("So bad that even 404 doesn't work")
	} else {
		res.Body = data
	}
	res.Headers["Content-Lenght"] = strconv.Itoa(len(res.Body))
	return res
}

func StreamHandler(req *parser.Request) (res *Response) {
	//the req.Path is cleaned up in router useRoute and now only contains the relevant info
	pathParts := strings.Split(req.Path, "/")
	ext := filepath.Ext(pathParts[len(pathParts)-1])
	ctype := mime.TypeByExtension(ext)
	if ctype == "" {
		//default for saftey shouldn't happen
		ctype = "application/octet-stream"
	}
	res = NewResponse(req)
	fmt.Println(ctype)
	res.Headers["Content-Type"] = ctype
	res.Headers["Connection"] = "close"
	data, err := reader.ReadFile(req.SourceFolder + req.Path)
	if err != nil {
		res.StatusCode = http.StatusBadRequest
		res.Body = []byte("File doesn't exist")
		res.Headers["Content-Type"] = "text/html; charset=utf-8"
	} else {
		res.Body = data
	}
	res.Headers["Content-Lenght"] = strconv.Itoa(len(res.Body))
	return res
}
