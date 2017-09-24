package traverse

import (
	"fmt"
	"testing"

	"github.com/mojo-zd/go-library/debug"
	"github.com/mojo-zd/go-library/traverse"
)

type Person struct {
	Name    string
	Sex     int
	Country *Country
}

type Country struct {
	Location string
}

func Test_Iterator(t *testing.T) {
	persons := []Person{{Name: "mojo"}, {Name: "mt"}}
	traverse.Iterator(persons, func(index int, value interface{}) (flag traverse.CYCLE_FLAG) {
		if value.(Person).Name == "mojo" {
			flag = traverse.BREAK_FLAT
		}
		debug.Display("========", "")
		return
	})
}
func Test_Map_Iterator(t *testing.T) {
	m := map[string]interface{}{"A": Person{Name: "mojo"}, "B": Person{Name: "mt"}}
	traverse.MapIterator(m, func(key, value interface{}, index int) {
		debug.Display("==k==", key)
		debug.Display("==v=", value)
	})
}

func Test_Contains(t *testing.T) {
	country := &Country{Location: "china"}
	person1 := &Person{Name: "mojo", Sex: 5, Country: country}
	person2 := &Person{Name: "mt", Sex: 4}
	person3 := &Person{Name: "mojo", Sex: 5, Country: country}
	persons := []*Person{person1, person2}
	fmt.Println(traverse.Contains(persons, person3))
}
