package transform

import (
	"testing"

	"github.com/mojo-zd/go-library/debug"
	"github.com/mojo-zd/go-library/transform"
)

type Source struct {
	Name  string
	Value interface{}
}

type Inline struct {
	TenantId int
	Type     string
}

func Test_Transform(t *testing.T) {
	inlines := []Inline{{TenantId: 1, Type: "png"}, {TenantId: 2, Type: "jpeg"}}
	source := Source{Name: "images", Value: inlines}
	result := transform.Convert(source.Value, Inline{})
	debug.Display("result is ", result)
}
