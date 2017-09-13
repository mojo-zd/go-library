package traverse

import (
	"testing"
	"wise-catalog/helper"

	"fmt"

	"github.com/mojo-zd/go-library/debug"
	"github.com/mojo-zd/go-library/traverse"
)

type Person struct {
	Name string
	Sex  int
}

func Test_Iterator_(t *testing.T) {
	persons := []Person{{Name: "mojo"}}
	traverse.Iterator(persons, func(index int, value interface{}) {
		debug.Display("", value)
	})
}

func Test_Contains(t *testing.T) {
	person1 := Person{Name: "mojo", Sex: 5}
	person2 := Person{Name: "mt", Sex: 4}
	person3 := Person{Name: "mojo", Sex: 5}
	persons := []Person{person1, person2}
	fmt.Println(helper.Contains(persons, person3))
}
