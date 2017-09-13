package traverse

import "reflect"

type IteratorFunc func(index int, value interface{})

func Iterator(collection interface{}, handlerFunc IteratorFunc) {
	v := reflect.ValueOf(collection)
	if !isSlice(collection) {
		panic("collection must be slice!")
		return
	}

	for index := 0; index < v.Len(); index++ {
		handlerFunc(index, v.Index(index).Interface())
	}
}

func Contains(collection interface{}, target interface{}) (contains bool) {

	if !isSlice(collection) {
		panic("collection must be slice!")
		return
	}
	v := reflect.ValueOf(collection)
	for index := 0; index < v.Len(); index++ {
		if v.Index(index).Interface() == target {
			contains = true
			break
		}
	}
	return
}

func isSlice(value interface{}) (slice bool) {
	v := reflect.ValueOf(value)
	slice = true
	if v.Kind() != reflect.Slice {
		slice = false
		return
	}
	return
}
