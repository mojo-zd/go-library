package traverse

import (
	"fmt"
	"testing"

	"github.com/mojo-zd/go-library/debug"
	"github.com/mojo-zd/go-library/traverse"
)

type Person struct {
	Name string
	Sex  int
}

func Test_Iterator_(t *testing.T) {
	persons := []Person{{Name: "mojo"}, {Name: "mt"}}
	traverse.Iterator(persons, func(index int, value interface{}) (flag traverse.CYCLE_FLAG) {
		if value.(Person).Name == "mojo" {
			flag = traverse.BREAK_FLAT
		}
		debug.Display("========", "")
		return
	})
}

func Test_Contains(t *testing.T) {
	person1 := &Person{Name: "mojo", Sex: 5}
	person2 := &Person{Name: "mt", Sex: 4}
	person3 := &Person{Name: "mojo", Sex: 5}
	persons := []*Person{person1, person2}
	fmt.Println(traverse.Contains(persons, person3))
	//Eq(person1, person3)
}

func Eq(value1, value2 interface{}) {
	debug.Display("", value1 == value2)
}
