package parser

import (
	"errors"
	"fmt"
	"strings"
)

type Request struct {
	Method   string
	Path     string
	HTTPType string
	Headers  map[string]string
	Params   []string
	Id       string
	Body     []byte
}

func NewRequest(method, path, id, HTTPType string, headers map[string]string, params []string, body []byte) *Request {
	return &Request{method, path, HTTPType, headers, params, id, body}
}

func ParseRequest(req string) (*Request, error) {
	fmt.Println(req)
	if len(req) == 0 {
		return nil, errors.New("empty request")
	}
	lines := strings.Split(strings.ReplaceAll(req, "\n", ""), "\r")
	word := strings.Split(lines[0], " ")
	//GET /request HTTP/1.1
	//this parses the first line correctly
	fmt.Println(word)
	Method := word[0]
	Path := word[1]
	Http := word[2]
	headers := parseHeaders(lines[1:])
	return &Request{Method: Method, Path: Path, HTTPType: Http, Headers: headers}, nil
}

func parseHeaders(headers []string) (out map[string]string) {
	out = make(map[string]string)
	for _, header := range headers {
		Pair := strings.Split(header, ": ")
		if len(Pair) != 2 {
			continue
		}
		out[Pair[0]] = Pair[1]
	}
	return out
}
