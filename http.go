package egb

import (
	"net/http"
	"io"
	"fmt"
	"bytes"
	"os"
	"path"
	"io/ioutil"
)

var UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1541.0 Safari/537.36"

// HttpCall makes HTTP method call.
func HttpCall(client *http.Client, method, url string, header http.Header, body io.Reader) (io.ReadCloser, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", UserAgent)
	for k, v := range header {
		request.Header[k] = v
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 200 {
		return response.Body, nil
	}
	response.Body.Close()
	if response.StatusCode == 404 {
		err = fmt.Errorf("resource not found: %s", url)
	} else {
		err = fmt.Errorf("%s %s -> %d", method, url, response.StatusCode)
	}
	return nil, err
}

// HttpGet gets the specified resource.
// ErrNotFound is returned if the server responds with status 404.
// eg:
// rc, err := HttpGet(&http.Client{}, "http://example.com", nil)
// p, err := ioutil.ReadAll(rc)
// s := string(p)
func HttpGet(client *http.Client, url string, header http.Header) (io.ReadCloser, error) {
	return HttpCall(client, "GET", url, header, nil)
}

// HttpPost posts the specified resource.
// ErrNotFound is returned if the server responds with status 404.
// var params = url.Values{}
// param.Add("userName", user)
// param.Add("pwd", pwd)
// data := param.Encode()
// rc, err := HttpGet(&http.Client{}, "http://example.com", nil,[]byte(data))
// p, err := ioutil.ReadAll(rc)
// s := string(p)
func HttpPost(client *http.Client, url string, header http.Header, body []byte) (io.ReadCloser, error) {
	return HttpCall(client, "POST", url, header, bytes.NewBuffer(body))
}

// HttpGetToFile gets the specified resource and writes to file.
// ErrNotFound is returned if the server responds with status 404.
func HttpGetToFile(client *http.Client, url string, header http.Header, fileName string) error {
	rc, err := HttpGet(client, url, header)
	if err != nil {
		return err
	}
	defer rc.Close()

	os.MkdirAll(path.Dir(fileName), os.ModePerm)
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, rc)
	return err
}

// HttpGetBytes gets the specified resource. ErrNotFound is returned if the server
// responds with status 404.
func HttpGetBytes(client *http.Client, url string, header http.Header) ([]byte, error) {
	rc, err := HttpGet(client, url, header)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}
