package go_library

import (
	"strings"
	"unicode"
)

//根据驼峰命名生成下划线字符串 ex:TestName --> test_name
func ToUnderLineLower(value string) string {
	r := ""
	for i, v := range value {
		if unicode.IsUpper(v) && i != 0 {
			r += "_"
		}
		r += string(v)
	}
	return strings.ToLower(r)
}
