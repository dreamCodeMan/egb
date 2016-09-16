package egb

import (
	"net/http"
	"strings"
	"io/ioutil"
	"io"
	"os"
)

//封装的http请求
//请求客户端
//Request
//Response
//error(如果有错误发生)
type RequestWrapper struct {
	client   *http.Client
	request  *http.Request
	response *http.Response
	err      error
}

func NewRequest(method, urlStr string) *RequestWrapper {
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil
	}
	return &RequestWrapper{request: req}
}

func Get(urlStr string) *RequestWrapper {
	return NewRequest("GET", urlStr)
}

func Post(urlStr string) *RequestWrapper {
	return NewRequest("POST", urlStr)
}

func Put(urlStr string) *RequestWrapper {
	return NewRequest("PUT", urlStr)
}

func Delete(urlStr string) *RequestWrapper {
	return NewRequest("DELETE", urlStr)
}

func (r *RequestWrapper) Query(query string) *RequestWrapper {
	r.request.URL.RawQuery = query
	return r
}

//添加url参数
func (r *RequestWrapper) Param(key, value string) *RequestWrapper {
	query := r.request.URL.Query()
	query.Add(key, value)
	return r.Query(query.Encode())
}

//设置header
func (r *RequestWrapper) Set(key, value string) *RequestWrapper {
	r.request.Header.Set(key, value)
	return r
}

//设置client
func (r *RequestWrapper) Use(client *http.Client) *RequestWrapper {
	if client != nil {
		r.client = client
	}
	return r
}

//设置post请求的参数 body主体
//`{"greeting":"hello world"}`
func (r *RequestWrapper) Json(data string) *RequestWrapper {
	reader := strings.NewReader(data)
	r.request.Body = ioutil.NopCloser(reader)
	r.request.ContentLength = int64(reader.Len())
	r.Set("Content-Type", "application/json")
	return r
}

//执行http请求
func (r *RequestWrapper) Exec() *RequestWrapper {
	client := http.DefaultClient
	if r.client != nil {
		client = r.client
	}
	r.response, r.err = client.Do(r.request)
	return r
}

//结果输出为[]byte
func (r *RequestWrapper) ToBytes() ([]byte, error) {
	if r.err != nil {
		return nil, r.err
	}
	defer r.response.Body.Close()
	return ioutil.ReadAll(r.response.Body)
}

//结果输出为string
func (r *RequestWrapper) ToString() (string, error) {
	data, err := r.ToBytes()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//结果输出为io.Writer
//返回length 以及error
func (r *RequestWrapper) Pipe(w io.Writer) (written int64, err error) {
	if r.err != nil {
		return 0, r.err
	}
	defer r.response.Body.Close()
	written, err = io.Copy(w, r.response.Body)
	return
}

//结果输出到文件
func (r *RequestWrapper) ToFile(filename string) (size int64, err error) {
	file, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	size, err = r.Pipe(file)
	return
}

//下载文件
func Download(urlStr, filename string) (size int64, err error) {
	size, err = Get(urlStr).Exec().ToFile(filename)
	return
}




