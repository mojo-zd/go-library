package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mojo-zd/go-library/debug"
	"github.com/mojo-zd/go-library/traverse"
)

var (
	defaultHeader = map[string]string{"Content-Type": "application/json"}
)

type RequestInfo struct {
	URL           string
	Params        map[string]interface{}
	Data          interface{}
	Header        map[string]string
	DefaultHeader bool
}

type HttpClient struct {
	*RequestInfo
}

func NewHttpClient() *HttpClient {
	return &HttpClient{&RequestInfo{}}
}

func (client *HttpClient) Get() (bytes []byte, err error) {
	bytes, err = doRequest(client, http.MethodGet)
	return
}

func (client *HttpClient) Post() (bytes []byte, err error) {
	bytes, err = doRequest(client, http.MethodPost)
	return
}

func (client *HttpClient) Put() (bytes []byte, err error) {
	bytes, err = doRequest(client, http.MethodPut)
	return
}

func doRequest(httpClient *HttpClient, method string) (bytes []byte, err error) {
	if err = validate(httpClient); err != nil {
		return
	}

	client := &http.Client{}
	request, err := http.NewRequest(method, httpClient.URL, strings.NewReader(toString(httpClient.Data)))
	if httpClient.DefaultHeader {
		httpClient.defaultHeader(request)
	}
	response, err := client.Do(request)
	bytes, err = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	debug.Display("client is ", toString(httpClient.Data))
	return
}

func (client *HttpClient) BuildRequestInfo(requestInfo *RequestInfo) *HttpClient {
	client.RequestInfo = requestInfo
	return client
}

func (client *HttpClient) defaultHeader(request *http.Request) *HttpClient {
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
