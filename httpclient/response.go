package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Raw     *http.Response
	Headers map[string]string
	Body    string
}

func NewResponse() *Response {
	return &Response{}
}

// 获取状态码
func (r *Response) StatusCode() int {
	return r.Raw.StatusCode
}

// 解析并设置响应头
func (r *Response) parseHeaders() error {
	headers := map[string]string{}
	for k, v := range r.Raw.Header {
		headers[k] = v[0]
	}
	r.Headers = headers
	return nil
}

// 解析并设置响应信息
func (r *Response) parseBody() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if body, err := ioutil.ReadAll(r.Raw.Body); err != nil {
		panic(err)
	} else {
		r.Body = string(body)
	}
	return nil
}
