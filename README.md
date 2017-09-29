### 安装
```
go get github.com/mojo-zd/go-library
```

### 常用工具类
```
示例对象:
type Person struct {
	Name string
	Sex  int
}

type Tag struct {
    ID string
    Name string
}
```
#### 数组遍历
```
func Iterator(collection interface{}, handlerFunc IteratorFunc)
IteratorFunc 返回值用于决定是否退出循环

示例(如果遇到Person.Name == "mojo"就退出循环):
persons := []Person{{Name: "mojo"}, {Name: "mt"}}
traverse.Iterator(persons, func(index int, value interface{}) (flag traverse.CYCLE_FLAG) {
	if value.(Person).Name == "mojo" {
		flag = traverse.BREAK_FLAT
	}
	return
})
```
#### 数组包含
```
traverse.Contains 判断数组中是否存在给定值, 如果是struct类型只比较struct中各个属性的值是否相等

示例:
person1 := &Person{Name: "mojo", Sex: 1}
person2 := &Person{Name: "mt", Sex: 0}
person3 := &Person{Name: "mojo", Sex: 1}
persons := []*Person{person1, person2}
fmt.Println(traverse.Contains(persons, person3))
```
#### slice转换为map
```
func StructsToMap(slice interface{}, key string, handleFunc MapHandleFunc) (m interface{})
MapHandleFunc可以指定map的value, 如果设置为nil则以slice中的对象为map的value

示例(返回的map.Key为 Person.Name的值, map.Value为value.(Person).Sex):
persons := []*Person{}
person := &Person{Name: "mojo", Sex: 1}
person1 := &Person{Name: "mt", Sex: 2}
persons = append(persons, person, person1)
result := traverse.StructsToMap(persons, "Name", func(value interface{}) (v interface{}) {
	v = value.(*Person).Sex
	return
})
debug.Display("", result)
```
#### map是否包含指定key值
```
示例:
m := map[string]string{"a": "A", "b": "B"}
debug.Display("", traverse.ContainsKey(m, "a"))
```

#### map是否包含指定value
```
示例(指针类型结构体只比较各属性值):
person1 := &Person{Name: "mojo"}
person2 := &Person{Name: "mt"}
person3 := &Person{Name: "mojo"}
m := map[string]*Person{"a": person1, "b": person2}
debug.Display("", traverse.ContainsValue(m, person3))
```
### http模块
```
结构说明:
type RequestInfo struct {
	URL           string //请求地址
	Params        map[string]interface{} //url参数
	Data          interface{} //body数据
	Header        map[string]string //header信息
	DefaultHeader bool //如果指定为true  header请求时将添加content-type:application/json头
}
client.BuildRequestInfo用于构建请求的相关参数
```
#### http Get请求
```
示例:
var client = http.NewHttpClient()
response := client.BuildRequestInfo(
	&http.RequestInfo{
		URL:           "http://127.0.0.1:8001/api/catalogs",
		DefaultHeader: true},
).Get()

debug.Display("===返回数据==", response)
```
#### http Post请求
```
var client = http.NewHttpClient()
response := client.BuildRequestInfo(&http.RequestInfo{
	URL:           "http://127.0.0.1:8001/api/tags",
	DefaultHeader: true,
	Data:          Tag{Name: "mt"},
}).Post()

debug.Display("==创建成功的返回数据==", response)
```

#### http Put请求
```
var client = http.NewHttpClient()
response := client.BuildRequestInfo(&http.RequestInfo{
	URL:           "http://127.0.0.1:8001/api/tags/2",
	DefaultHeader: true,
	Data:          Tag{Name: "mojo-v2", ID: 2},
}).Put()

debug.Display("==修改成功的返回数据==", response)
```

#### json转struct
推荐使用 https://github.com/json-iterator/go