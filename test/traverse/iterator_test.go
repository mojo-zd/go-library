package traverse

import (
	"testing"

	"github.com/mojo-zd/go-library/debug"
	"github.com/mojo-zd/go-library/traverse"
)

type Person struct {
	Name string
}

func Test_Iterator_(t *testing.T) {
	persons := []Person{Person{Name: "mojo"}}
	traverse.Iterator(persons, func(index int, value interface{}) {
		debug.Display("", value)
	})
}
