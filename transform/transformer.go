package transform

import (
	"reflect"

	"fmt"

	"github.com/json-iterator/go"
	r "github.com/mojo-zd/go-library/reflect"
)

func Convert(source, target interface{}) (result interface{}) {
	ty := r.GetType(reflect.TypeOf(source))
	v := r.GetValue(reflect.ValueOf(source))
	switch ty.Kind() {
	case reflect.Slice:
		targets := []interface{}{}
		for i := 0; i < v.Len(); i++ {
			targets = append(targets, r.NewInstance(target))
		}
		b, err := jsoniter.Marshal(source)
		if err != nil {
			panic(fmt.Sprintf("source marshal failed, error info is %s", err.Error()))
		}
		jsoniter.Unmarshal(b, &targets)
		result = targets
	default:

	}

	return
}
