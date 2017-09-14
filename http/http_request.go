package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mojo-zd/go-library/traverse"
)

var (
	defaultHeader = map[string]string{"Content-Type": "application/json"}
)

type RequestInfo struct {
	URL    string
	Params map[string]interface{}
	Data   interface{}
	Header map[string]string
}

type HttpClient struct {
	*RequestInfo
}

func NewHttpClient() *HttpClient {
	return &HttpClient{&RequestInfo{}}
}

func (httpClient *HttpClient) Get() (bytes []byte, err error) {
	if err = validate(httpClient); err != nil {
		return
	}
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, httpClient.URL, strings.NewReader(toString(httpClient.Data)))
	if err != nil {
		return
	}
	response, err := client.Do(request)
	bytes, err = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return
}

func (client *HttpClient) Post() (v interface{}, err error) {
	return
}

func (client *HttpClient) Put() (v interface{}, err error) {
	return
}

func (client *HttpClient) BuildRequestInfo(requestInfo *RequestInfo) *HttpClient {
	client.RequestInfo = requestInfo
	return client
}

func (client *HttpClient) DefaultHeader(request *http.Request) *HttpClient {
	traverse.MapIterator(defaultHeader, func(key, value interface{}, index int) {
		request.Header.Add(key.(string), value.(string))
	})
	return client
}

func validate(client *HttpClient) (err error) {
	if len(client.URL) == 0 {
		err = errors.New("必须执行请求地址!")
		return
	}
	return
}

func (client *HttpClient) buildHeader(request *http.Request) *HttpClient {
	traverse.MapIterator(client.Header, func(key, value interface{}, index int) {
		request.Header.Add(key.(string), value.(string))
	})
	return client
}

func (client *HttpClient) BuildURL() (url string) {
	url = client.URL
	traverse.MapIterator(client.Params, func(key, value interface{}, index int) {
		if index > 0 {
			url += fmt.Sprintf("&%s=%v", key.(string), value)
		} else {
			url += fmt.Sprintf("?%s=%v", key.(string), value)
		}
		index++
	})
	return
}

func toString(data interface{}) (str string) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	str = string(bytes)
	return
}
