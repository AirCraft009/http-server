package server

import "strconv"

func (res *Response) AddHeader(key, value string) {
	res.Headers[key] = value
}

func (res *Response) SetBody(body []byte) {
	res.Body = body
	res.Headers["Content-Length"] = strconv.Itoa(len(body))
}
