package traverse

import (
	"testing"

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
	person1 := &Person{Name: "mojo", Sex: 5}
	//person2 := &Person{Name: "mt", Sex: 4}
	person3 := &Person{Name: "mojo", Sex: 5}
	//persons := []*Person{person1, person2}
	//debug.Display("", *person1 == *person3)
	//fmt.Println(helper.Contains(persons, person3))
	Eq(person1, person3)
}

func Eq(value1, value2 interface{}) {
	debug.Display("", value1 == value2)
}
