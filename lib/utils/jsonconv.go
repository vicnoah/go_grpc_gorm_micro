package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// 使用
//1、JsonSnakeCase统一转下划线json
//使用jsonconv.JsonSnakeCase包裹一下要输出json的对象即可
//func main() {
//	type Person struct {
//		HelloWold       string
//		LightWeightBaby string
//	}
//	var a = Person{HelloWold: "chenqionghe", LightWeightBaby: "muscle"}
//	res, _ := json.Marshal(jsonconv.JsonSnakeCase{a})
//	fmt.Printf("%s", res)
//}
//输出如下
//{"hello_wold":"chenqionghe","light_weight_baby":"muscle"}


//2、JsonCamelCase统一转驼峰json
//已经指定了下划线标签的结构体，我们也可以统一转为驼峰的json
//func main() {
//	type Person struct {
//		HelloWold       string `json:"hello_wold"`
//		LightWeightBaby string `json:"light_weight_baby"`
//	}
//	var a = Person{HelloWold: "chenqionghe", LightWeightBaby: "muscle"}
//	res, _ := json.Marshal(jsonconv.JsonCamelCase{a})
//	fmt.Printf("%s", res)
//}
//输出如下
//{"helloWold":"chenqionghe","lightWeightBaby":"muscle"}


/*************************************** 下划线json ***************************************/
type JsonSnakeCase struct {
	Value interface{}
}

func (c JsonSnakeCase) MarshalJSON() ([]byte, error) {
	// Regexp definitions
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)
	marshalled, err := json.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)
	return converted, err
}

/*************************************** 驼峰json ***************************************/
type JsonCamelCase struct {
	Value interface{}
}

func (c JsonCamelCase) MarshalJSON() ([]byte, error) {
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	marshalled, err := json.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matchStr := string(match)
			key := matchStr[1 : len(matchStr)-2]
			resKey := Lcfirst(Case2Camel(key))
			return []byte(`"` + resKey + `":`)
		},
	)
	return converted, err
}

/*************************************** 其他方法 ***************************************/
// 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 下划线写法转为驼峰写法并且首字母大写
func Case2CamelAndUcfirst(name string) string {
	return Ucfirst(Case2Camel(name))
}

// 下划线写法转为驼峰写法并且首字母小写
func Case2CamelAndLcfirst(name string) string {
	return Lcfirst(Case2Camel(name))
}

// 首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}