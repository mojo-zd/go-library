package traverse

import (
	"fmt"
	"reflect"
)

type CYCLE_FLAG int

const (
	_ CYCLE_FLAG = iota
	BREAK_FLAT
	CONTINUE_FLAT
	NOTHING_TODO
)

type IteratorFunc func(index int, value interface{}) CYCLE_FLAG

func Iterator(collection interface{}, handlerFunc IteratorFunc) {
	v := reflect.ValueOf(collection)
	if !isSlice(collection) {
		panic("collection must be slice!")
		return
	}

	for index := 0; index < v.Len(); index++ {
		flag := handlerFunc(index, v.Index(index).Interface())
		if flag == BREAK_FLAT {
			break
		} else if flag == CONTINUE_FLAT {
			continue
		}
	}
}

//非指针结构直接比较 指针类型比较每一个属性
func Contains(collection interface{}, target interface{}) (contains bool) {

	if !isSlice(collection) {
		panic("collection must be slice!")
		return
	}

	v := reflect.ValueOf(collection)
	for index := 0; index < v.Len(); index++ {
		value := v.Index(index).Interface()
		if compare(value, target) {
			contains = true
			break
		}
	}
	return
}

func compare(v1, v2 interface{}) (isEqual bool) {
	isEqual = true
	v1Ty := reflect.TypeOf(v1)
	v2Ty := reflect.TypeOf(v2)
	if v1Ty.Kind() != v2Ty.Kind() {
		panic("collection item type is not same as target!!!")
	}

	if v1Ty.Kind() == reflect.Ptr {
		v1Ty = getType(v1Ty)
		v2Ty = getType(v2Ty)
		v1FieldNum := v1Ty.NumField()
		v2FiledNum := v2Ty.NumField()

		if v1FieldNum != v2FiledNum {
			isEqual = false
			return
		}

		v1Val := getValue(reflect.ValueOf(v1))
		v2Val := getValue(reflect.ValueOf(v2))

		for i := 0; i < v1FieldNum; i++ {
			fieldName := v1Ty.Field(i).Name
			if v1Val.FieldByName(fieldName).Interface() != v2Val.FieldByName(fieldName).Interface() {
				isEqual = false
				break
			}
		}
		return
	}
	return v1 == v2
}

func isSlice(value interface{}) (slice bool) {
	v := reflect.ValueOf(value)
	return v.Kind() == reflect.Slice
}

func isMap(value interface{}) bool {
	ty := reflect.ValueOf(value)
	fmt.Println(ty.Kind() == reflect.Map)
	return ty.Kind() == reflect.Map
}
