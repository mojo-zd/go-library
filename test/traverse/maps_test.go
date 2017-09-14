package traverse

import (
	"testing"

	"github.com/mojo-zd/go-library/debug"
	"github.com/mojo-zd/go-library/traverse"
)

func Test_GetStructByName(t *testing.T) {
	person := Person{Name: "mojo", Sex: 5}
	debug.Display("", traverse.GetValueByName(person, "Name"))
}

func Test_StructToMap(t *testing.T) {
	persons := []*Person{}
	person := &Person{Name: "mojo", Sex: 1}
	person1 := &Person{Name: "mt", Sex: 2}
	persons = append(persons, person, person1)
	result := traverse.StructsToMap(persons, "Name", func(value interface{}) (v interface{}) {
		v = value.(*Person).Sex
		return
	})
	debug.Display("", result)
}

func Test_Map_Contains(t *testing.T) {
	m := map[string]string{"a": "A", "b": "B"}
	debug.Display("", traverse.ContainsKey(m, "a"))
	//fmt.Sprintf()
}
