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
	Body     []byte
}

func NewRequest(method, path, HTTPType string, headers map[string]string, body []byte) *Request {
	return &Request{method, path, HTTPType, headers, body}
}

func ParseRequest(byteReq []byte) (*Request, error) {
	req := string(byteReq)
	if len(req) == 0 {
		return nil, errors.New("empty request")
	}
	/**\r\n is an empty line and marks the body
	because the header only uses ascii it perfectly matches up with a byte arrray.
	Now I can pinpoint the start of the body which should stay in []byte format because
	it could be a file like an image or pdf.
	*/
	headerlen := strings.Index(req, "\r\n\r\n")
	fmt.Println(headerlen)
	lines := strings.Split(strings.ReplaceAll(req, "\n", ""), "\r")
	word := strings.Split(lines[0], " ")
	//GET /request HTTP/1.1
	//this parses the first line correctly
	Method := word[0]
	Path := word[1]
	Http := word[2]
	headers := parseHeaders(lines[1:])

	return NewRequest(Method, Path, Http, headers, byteReq[headerlen:]), nil
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
