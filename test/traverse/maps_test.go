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

func Test_Map_Contains_Key(t *testing.T) {
	m := map[string]string{"a": "A", "b": "B"}
	debug.Display("", traverse.ContainsKey(m, "a"))
	//fmt.Sprintf()
}

func Test_Map_Conatins_Value(t *testing.T) {
	person1 := &Person{Name: "mojo"}
	person2 := &Person{Name: "mt"}
	person3 := &Person{Name: "mojo"}
	m := map[string]*Person{"a": person1, "b": person2}
	debug.Display("", traverse.ContainsValue(m, person3))
}
