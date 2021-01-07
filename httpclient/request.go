package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/yuyu888/golibs/utils"
)

// Request构造类
type Request struct {
	Method          string
	Url             string
	DialTimeout     time.Duration
	ResponseTimeOut time.Duration
	Headers         map[string]string
	Body            io.Reader
}

// 创建一个Request实例
func NewRequest() *Request {
	r := &Request{}
	r.Method = "GET"
	r.DialTimeout = 5
	r.ResponseTimeOut = 5
	r.Body = nil
	r.Headers = map[string]string{}
	return r
}

//SetDialTimeOut
func (r *Request) SetDialTimeOut(timeOutSecond int) *Request {
	r.DialTimeout = time.Duration(timeOutSecond)
	return r
}

//SetResponseTimeOut
func (r *Request) SetResponseTimeOut(timeOutSecond int) *Request {
	r.ResponseTimeOut = time.Duration(timeOutSecond)
	return r
}

// 设置请求方法
func (r *Request) SetMethod(method string) *Request {
	r.Method = method
	return r
}

// 设置请求地址
func (r *Request) SetUrl(url string) *Request {
	r.Url = url
	return r
}

// 设置请求头
func (r *Request) SetHeaders(headers map[string]string) *Request {
	r.Headers = headers
	return r
}

// 设置body
func (r *Request) SetBody(body io.Reader) *Request {
	r.Body = body
	return r
}

func (r *Request) SetStringPostdata(postData string) *Request {
	r.Body = strings.NewReader(postData)
	return r
}

func (r *Request) SetBytePostdata(postData []byte) *Request {
	r.Body = bytes.NewReader(postData)
	return r
}

func (r *Request) Get(url string) (*Response, error) {
	r.Url = url
	r.Method = "GET"
	return r.Send()
}

// 普通的post提交 "Content-Type": "application/x-www-form-urlencoded"
func (r *Request) Post(url string, postData url.Values) (*Response, error) {
	r.Url = url
	r.Method = "POST"
	r.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	if postData != nil {
		r.Body = strings.NewReader(postData.Encode())
	}
	return r.Send()
}

// 模拟Form表单的方式提交
// postData 中 @upload_param_name 表示上传的字段名， @upload_file_path 表示上传的文件地址 可以是 本地地址或者远程地址
func (r *Request) PostForm(url string, postData map[string]string) (*Response, error) {
	r.Url = url
	r.Method = "POST"
	r.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	if postData != nil {
		buf := new(bytes.Buffer)
		writer := multipart.NewWriter(buf)
		for k, v := range postData {
			if k != "@upload_param_name" && k != "@upload_file_path" {
				writer.WriteField(k, v)
			}
		}

		_, ok1 := postData["@upload_param_name"]
		_, ok2 := postData["@upload_file_path"]

		// 上传文件
		if ok1 && ok2 {
			fileData, err := utils.FileGetContents(postData["@upload_file_path"]) // 此处内容可以来自本地文件读取或远程文件
			if err == nil {
				part, err := writer.CreateFormFile(postData["@upload_param_name"], "tmpfile")
				if err != nil {
					return nil, err
				}
				part.Write(fileData)
			}
		}

		if err := writer.Close(); err != nil {
			return nil, err
		}
		r.Headers["Content-Type"] = writer.FormDataContentType()
		r.Body = buf
	}
	return r.Send()
}

func (r *Request) Put(url string, body string) (*Response, error) {
	r.Url = url
	r.Method = "PUT"
	r.Body = strings.NewReader(body)
	return r.Send()
}

// delete 方法 参数传递一般通过URL_Query 传递
func (r *Request) Delete(url string) (*Response, error) {
	r.Url = url
	r.Method = "DELETE"
	return r.Send()
}

func (r *Request) Send() (*Response, error) {

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*r.DialTimeout)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * r.DialTimeout))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * r.ResponseTimeOut,
		},
	}

	req, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		return nil, err
	}

	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	// 初始化Response对象
	response := NewResponse()

	if resp, err := client.Do(req); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		response.Raw = resp
	}
	response.parseHeaders()
	response.parseBody()

	defer response.Raw.Body.Close()

	return response, nil
}
