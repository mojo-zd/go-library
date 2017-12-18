package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mojo-zd/go-library/traverse"
)

var (
	defaultHeader = map[string]string{"Content-Type": "application/json", "Accept": "application/json"}
)

type RequestInfo struct {
	URL     string
	Params  map[string]interface{}
	Data    interface{}
	Header  map[string]string
	Timeout int //以s为单位
}

type ResponseInfo struct {
	Code   int
	Result []byte
	Error  error
}

type HttpClient struct {
	*RequestInfo
}

func NewHttpClient() *HttpClient {
	return &HttpClient{&RequestInfo{}}
}

func (client *HttpClient) Get() (responseInfo *ResponseInfo) {
	responseInfo = doRequest(client, http.MethodGet)
	return
}

func (client *HttpClient) Post() (responseInfo *ResponseInfo) {
	responseInfo = doRequest(client, http.MethodPost)
	return
}

func (client *HttpClient) Put() (responseInfo *ResponseInfo) {
	responseInfo = doRequest(client, http.MethodPut)
	return
}

func (client *HttpClient) Delete() (responseInfo *ResponseInfo) {
	responseInfo = doRequest(client, http.MethodDelete)
	return
}

func doRequest(httpClient *HttpClient, method string) (responseInfo *ResponseInfo) {
	responseInfo = &ResponseInfo{}

	if err := validate(httpClient); err != nil {
		responseInfo.Error = err
		return
	}

	bytes := []byte{}
	client := &http.Client{}

	request, err := http.NewRequest(method, httpClient.BuildURL(), strings.NewReader(toString(httpClient.Data)))
	setClientInfo(httpClient.RequestInfo, client, httpClient, request)

	response, err := client.Do(request)
	if response != nil {
		responseInfo.Code = response.StatusCode
	}

	if err != nil {
		responseInfo.Error = err
		return
	}

	bytes, err = ioutil.ReadAll(response.Body)
	responseInfo.Result = bytes

	if responseInfo.Code == http.StatusNotFound {
		responseInfo.Error = errors.New(string(bytes))
	}
	defer response.Body.Close()

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

func setClientInfo(requestInfo *RequestInfo, client *http.Client, httpClient *HttpClient, request *http.Request) {
	if requestInfo.Timeout > 0 {
		client.Timeout = time.Duration(requestInfo.Timeout) * time.Second
	}

	httpClient.defaultHeader(request)

	if len(requestInfo.Header) > 0 {
		httpClient.buildHeader(request)
	}
}

func toString(data interface{}) (str string) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	str = string(bytes)
	return
}
