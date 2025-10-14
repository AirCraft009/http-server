// Package parser
//
// used by the Package server to parse Requests
// for that purposes it uses the Request struct
// aswell as the ParseRequest Method
package parser

import (
	"bytes"
	"errors"
	"strings"
)

type Request struct {
	Method   string
	Path     string
	HTTPType string
	Headers  map[string]string
	Body     []byte
	//field that will only be filled later
	Queries      map[string]string
	SourceFolder string
}

func NewRequest(method, path, HTTPType string, headers map[string]string, body []byte) *Request {
	//currently all queries are in the path variable of the request they will be parsed later
	return &Request{method, path, HTTPType, headers, body, nil, ""}
}

func (request *Request) GetQuery(key string) (value string, err error) {
	if data, ok := request.Queries[key]; ok {
		return data, nil
	}
	return "", errors.New(key + " not found")
}

func (request *Request) GetHeader(key string) (value string, err error) {
	if data, ok := request.Headers[key]; ok {
		return data, nil
	}
	return "", errors.New(key + " not found")
}

func (request *Request) GetBody() (body []byte, err error) {
	if request.Body != nil {
		return request.Body, nil
	}
	return nil, errors.New("no body")
}

// ParseRequest
// parses a Request in byteform and returns *Request, err
// it also checks the request for any malformed data (no end of header \r\n line)
// this is used for handleConnection in server.AcceptConnections()
// it doesn't handle chunked incoding or packets ariving out of sync.
// to big for the scope of this project
func ParseRequest(byteReq []byte) (*Request, error) {
	req := string(byteReq)
	if len(req) == 0 {
		return nil, errors.New("empty request")
	}
	/**\r\n is an empty line and marks the body
	TODO: !! switched to using bytes.Index because even though the header should be ascii only
	TODO: !! it could not be to not have issues I'll use bytes.Index
	(because the header only uses ascii it perfectly matches up with a byte arrray.
		Now I can pinpoint the start of the body which should stay in []byte format because
		it could be a file like an image or pdf.)
	*/
	headerlen := bytes.Index(byteReq, []byte("\r\n\r\n"))
	if headerlen == -1 {
		return nil, errors.New("invalid syntax no header end")
	}
	lines := strings.Split(req, "\r\n")
	word := strings.Split(lines[0], " ")
	if len(word) != 3 {
		return nil, errors.New("bad request")
	}
	//GET /request HTTP/1.1
	//this parses the first line correctly
	Method, Path, Http := word[0], word[1], word[2]
	headers := parseHeaders(lines[1:])
	return NewRequest(Method, Path, Http, headers, byteReq[headerlen:]), nil
}

func parseHeaders(headers []string) (out map[string]string) {
	if len(headers) == 0 {
		return map[string]string{}
	}
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
