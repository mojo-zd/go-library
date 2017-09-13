package traverse

import "reflect"

type IteratorFunc func(index int, value interface{})

func Iterator(collection interface{}, handlerFunc IteratorFunc) {
	v := reflect.ValueOf(collection)
	if v.Kind() != reflect.Slice {
		panic("collection must be slice!")
		return
	}

	for index := 0; index < v.Len(); index++ {
		handlerFunc(index, v.Index(index).Interface())
	}
}
