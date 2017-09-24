package reflect

import (
	"testing"

	"github.com/astaxie/beego/utils"
	"github.com/mojo-zd/go-library/reflect"
)

type Person struct {
	Name string
}

func Test_New(t *testing.T) {
	p := reflect.NewInstance(&Person{})
	utils.Display("===", p)
}
