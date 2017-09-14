package http

import (
	"testing"

	"fmt"

	"github.com/mojo-zd/go-library/debug"
	"github.com/mojo-zd/go-library/http"
)

var client = http.NewHttpClient()

func init() {
	client.URL = "http://127.0.0.1:8001"
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func Test_Request_Get(t *testing.T) {
	client.BuildRequestInfo(&http.RequestInfo{URL: fmt.Sprintf("%s%s", client.URL, "/api/catalogs"), DefaultHeader: true})
	bytes, err := client.Get()
	debug.Display("===错误信息===", err)
	debug.Display("===返回数据==", string(bytes))
}

func Test_Request_Post(t *testing.T) {
	client.BuildRequestInfo(&http.RequestInfo{URL: fmt.Sprintf("%s%s", client.URL, "/api/tags"), DefaultHeader: true})
	client.Data = Tag{Name: "mt"}
	bytes, err := client.Post()
	debug.Display("==创建成功的返回数据==", string(bytes))
	debug.Display("==错误信息==", err)
}

func Test_Request_Put(t *testing.T) {
	client.BuildRequestInfo(&http.RequestInfo{
		URL:           fmt.Sprintf("%s%s", client.URL, "/api/tags/2"),
		DefaultHeader: true,
		Data:          Tag{Name: "mojo-v2", ID: 2},
	})
	bytes, err := client.Put()
	debug.Display("==修改成功的返回数据==", string(bytes))
	debug.Display("==错误信息==", err)
}
