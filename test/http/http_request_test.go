package http

import (
	"testing"

	"fmt"

	"github.com/mojo-zd/go-library/http"
)

var client = http.NewHttpClient()

func Test_Request_URL(t *testing.T) {
	client.URL = "http://127.0.0.1"
	client.Params = map[string]interface{}{"1": "xxx", "key": 5}
	fmt.Println(client.BuildURL())
}

func Test_Request_Get(t *testing.T) {
	client.URL = "http://121.40.232.51:30320/api/catalogs"
	client.Get()
}
