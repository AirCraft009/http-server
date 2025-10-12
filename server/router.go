package server

import (
	"http-server/parser"
	//http wird nur für constants wie status codes verwendet
	"net/http"
	"strconv"
	"strings"
)

type Router struct {
	// Routes Method -> route -> Handler
	Routes       map[string]map[string]Handler
	sourceFolder string
}

func NewRouter() *Router {
	return &Router{Routes: make(map[string]map[string]Handler)}
}

func NewResponse(req *parser.Request) *Response {
	return &Response{HTTPType: req.HTTPType, Headers: make(map[string]string), StatusCode: -1}
}

type Handler func(request *parser.Request) *Response

// Handle
//
//sets a route in the router but doesn't support spaces in the path
//it takes the Method as in a string like GET POST PUT
// and a path like(localhost:8080/front/gopher.jpg) aswell as a handler and routes any traffic on that route to this handler
func (r *Router) Handle(Method, path string, handler Handler) {
	if r.Routes[Method] == nil {
		r.Routes[Method] = make(map[string]Handler)
	}
	r.Routes[Method][path] = handler
}

func (r *Router) useRoute(route string, req *parser.Request) *Response {
	rawPath, queryMap := parseQuery(route)
	req.Querys = queryMap
	req.Path = rawPath
	req.SourceFolder = r.sourceFolder
	handler, ok := r.Routes[req.Method][rawPath]
	if !ok {
		return Http404Handler(req)
	}
	return setBasicResponseHeaders(handler(req))
}

func setBasicResponseHeaders(response *Response) *Response {
	// sets all params to a base if they haven't been set
	if response.StatusCode == -1 {
		response.StatusCode = 200
	}
	if _, ok := response.Headers["Content-Type"]; !ok {
		response.Headers["Content-Type"] = "text/html; charset=utf-8"
	}
	response.Headers["Content-Lenght"] = strconv.Itoa(len(response.Body))
	return response
}

func parseQuery(path string) (string, map[string]string) {
	querys := strings.Split(path, "?")
	path = querys[0]
	if len(querys) != 2 {
		return path, nil
	}
	querys = strings.Split(querys[1], "&")
	queryMap := make(map[string]string)
	for _, query := range querys {
		keyValue := strings.Split(query, "=")
		if len(keyValue) != 2 {
			continue
		}
		queryMap[keyValue[0]] = keyValue[1]
	}
	return path, queryMap
}

func formatResponse(response *Response) []byte {
	responseBuilder := strings.Builder{}
	if response.HTTPType == "" {
		response.StatusCode = http.StatusInternalServerError
		response.HTTPType = httpType0
	}
	responseBuilder.Write(formatStartLine(response))
	responseBuilder.Write(formatResponseHeaders(response))
	//empty line before body
	responseBuilder.WriteString("\r\n")
	responseBuilder.Write(response.Body)
	return []byte(responseBuilder.String())
}

func formatStartLine(response *Response) []byte {
	lineBuilder := strings.Builder{}
	lineBuilder.WriteString(response.HTTPType)
	lineBuilder.WriteString(" ")
	lineBuilder.WriteString(strconv.Itoa(response.StatusCode))
	lineBuilder.WriteString(" ")
	lineBuilder.WriteString(http.StatusText(response.StatusCode))
	lineBuilder.WriteString("\r\n")
	return []byte(lineBuilder.String())
}

func formatResponseHeaders(response *Response) []byte {
	headerBuilder := strings.Builder{}
	for key, value := range response.Headers {
		headerBuilder.WriteString(key)
		headerBuilder.WriteString(": ")
		headerBuilder.WriteString(value)
		headerBuilder.WriteString("\r\n")
	}
	return []byte(headerBuilder.String())
}
