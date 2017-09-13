package traverse

import "reflect"

type MapHandleFunc func(value interface{}) interface{}

//指定struct中Key的value作为map的key map的value可以由MapHandleFunc的返回值决定, 如果想以struct作为value回调函数设置为nil即可
//ex: Person{
// Name string
// Sex int
// }
// structs := []Person{{Name: "mojo", Sex:1}}
// StructsToMap(structs, "Name")
func StructsToMap(slice interface{}, key string, handleFunc MapHandleFunc) (result map[interface{}]interface{}) {
	if !isSlice(slice) {
		panic("collection must be slice!")
		return
	}
	result = map[interface{}]interface{}{}
	v := reflect.ValueOf(slice)
	for index := 0; index < v.Len(); index++ {
		value := v.Index(index).Interface()
		keyValue := GetValueByName(value, key)
		if handleFunc != nil {
			result[keyValue] = handleFunc(value)
		} else {
			result[keyValue] = value
		}
	}
	return
}

func GetValueByName(i interface{}, key string) (value interface{}) {
	ty := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	for index := 0; index < ty.NumField(); index++ {
		if ty.Field(index).Name == key {
			value = v.FieldByName(ty.Field(index).Name).Interface()
			break
		}
	}
	return
}
