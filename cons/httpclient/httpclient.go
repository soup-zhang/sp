/**
 * @Time: 2023/1/6 11:12
 * @Author: soupzhb@gmail.com
 * @File: httpclient.go
 * @Software: GoLand
 */

package httpclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//HTTPGet get 请求
func HTTPGet(uri string, proxyAddr string) ([]byte, error) {
	//代理
	var proxy func(r *http.Request) (*url.URL, error)
	var proxyUrl string

	if proxyAddr != "" {
		proxyUrl = proxyAddr
	}

	if proxyUrl != "" {
		proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
	}

	dialer := &net.Dialer{
		Timeout: time.Duration(10 * int64(time.Second)),
		// KeepAlive: time.Duration(1 * int64(time.Second)),
	}

	var isHttps bool
	if strings.Index(uri, "https") != -1 {
		isHttps = true
	}

	transport := &http.Transport{
		Proxy: proxy, DialContext: dialer.DialContext,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: isHttps,
		},
	}
	client := &http.Client{
		Transport: transport,
	}

	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//PostJSON post json 数据请求
func PostJSON(uri string, obj interface{}, proxyAddr string) ([]byte, error) {
	//代理
	var proxy func(r *http.Request) (*url.URL, error)
	var proxyUrl string

	if proxyAddr != "" {
		proxyUrl = proxyAddr
	}

	if proxyUrl != "" {
		proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
	}

	dialer := &net.Dialer{
		Timeout: time.Duration(10 * int64(time.Second)),
		// KeepAlive: time.Duration(1 * int64(time.Second)),
	}

	var isHttps bool
	if strings.Index(uri, "https") != -1 {
		isHttps = true
	}

	transport := &http.Transport{
		Proxy: proxy, DialContext: dialer.DialContext,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: isHttps,
		},
	}
	client := &http.Client{
		Transport: transport,
	}

	//参数处理
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	jsonData = bytes.Replace(jsonData, []byte("\\u003c"), []byte("<"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u003e"), []byte(">"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u0026"), []byte("&"), -1)

	body := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

//MultipartFormField 保存文件或其他字段信息
type MultipartFormField struct {
	IsFile    bool
	FieldName string
	Value     []byte
	Filename  string
}

//PostFile 上传文件
func PostFile(fieldName, filename, uri, proxyAddr string) (respBody []byte, err error) {
	fields := []MultipartFormField{
		{
			IsFile:    true,
			FieldName: fieldName,
			Filename:  filename,
		},
	}
	return PostMultipartForm(fields, uri, proxyAddr)
}

//PostMultipartForm 上传文件或其他多个字段
func PostMultipartForm(fields []MultipartFormField, uri string, proxyAddr string) (respBody []byte, err error) {
	//代理
	var proxy func(r *http.Request) (*url.URL, error)
	var proxyUrl string

	if proxyAddr != "" {
		proxyUrl = proxyAddr
	}

	if proxyUrl != "" {
		proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
	}

	dialer := &net.Dialer{
		Timeout: time.Duration(10 * int64(time.Second)),
		// KeepAlive: time.Duration(1 * int64(time.Second)),
	}

	var isHttps bool
	if strings.Index(uri, "https") != -1 {
		isHttps = true
	}

	transport := &http.Transport{
		Proxy: proxy, DialContext: dialer.DialContext,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: isHttps,
		},
	}
	client := &http.Client{
		Transport: transport,
	}

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for _, field := range fields {
		if field.IsFile {
			fileWriter, e := bodyWriter.CreateFormFile(field.FieldName, field.Filename)
			if e != nil {
				err = fmt.Errorf("error writing to buffer , err=%v", e)
				return
			}

			fh, e := os.Open(field.Filename)
			if e != nil {
				err = fmt.Errorf("error opening file , err=%v", e)
				return
			}
			defer fh.Close()

			if _, err = io.Copy(fileWriter, fh); err != nil {
				return
			}
		} else {
			partWriter, e := bodyWriter.CreateFormField(field.FieldName)
			if e != nil {
				err = e
				return
			}
			valueReader := bytes.NewReader(field.Value)
			if _, err = io.Copy(partWriter, valueReader); err != nil {
				return
			}
		}
	}

	bodyWriter.Close()

	req, err := http.NewRequest("POST", uri, bodyBuf)
	req.Header.Set("Content-Type", "multipart/form-data")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http PostMultipartForm error : uri=%v , statusCode=%v", uri, resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

// HttpHandle http请求 可设置cookie, header
// todo 扩展
func HttpHandle(method, uri, data string, proxyAddr string, cookieVar []http.Cookie, headerVal map[string]string) (respBody []byte, err error) {
	//代理
	var proxy func(r *http.Request) (*url.URL, error)
	var proxyUrl string

	if proxyAddr != "" {
		proxyUrl = proxyAddr
	}

	if proxyUrl != "" {
		proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
	}

	dialer := &net.Dialer{
		Timeout: time.Duration(10 * int64(time.Second)),
		// KeepAlive: time.Duration(1 * int64(time.Second)),
	}

	var isHttps bool
	if strings.Index(uri, "https") != -1 {
		isHttps = true
	}

	transport := &http.Transport{
		Proxy: proxy, DialContext: dialer.DialContext,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: isHttps,
		},
	}
	client := &http.Client{
		Transport: transport,
	}

	//请求参数处理
	var req *http.Request

	if data == "" {
		urlArr := strings.Split(uri, "?")
		if len(urlArr) == 2 {
			uri = urlArr[0] + "?" + getParseParam(urlArr[1])
		}
		req, _ = http.NewRequest(method, uri, nil)
	} else {
		req, _ = http.NewRequest(method, uri, strings.NewReader(data))
	}

	//设置cookie
	for _, v := range cookieVar {
		cookie := &http.Cookie{Name: v.Name, Value: v.Value, HttpOnly: v.HttpOnly, Path: v.Path, Domain: v.Domain}
		req.AddCookie(cookie)
	}

	//设置header
	for k, v := range headerVal {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)

	if err != nil {
		return
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//getParseParam 请求参数进行转义
func getParseParam(param string) string {
	return url.PathEscape(param)
}

