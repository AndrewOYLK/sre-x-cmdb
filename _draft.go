package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/AndrewOYLK/ou-cmdb/model"
)

// Enum 枚举结构
type Enum struct {
	Key   string
	Value string
}

// [{"key":"k1","value":"v1"},{"key":"k2","value":"v2"},{"key":"k3","value":"v3"}]
func ParaseEnumOptions(options string) (*[]Enum, error) {
	var enums []Enum

	if err := json.Unmarshal([]byte(options), &enums); err != nil {
		return nil, err
	}

	return &enums, nil
}

/*
	type Host struct {
		p1 string
		p2 int64
		p3 bool
		p4 string 	// date
		p5 string 	// enum
		p6 float64
	}
*/

// 根据模型的属性，组装模型结构体
func GenDynamicType(model model.Model, attrs []model.Attribute) reflect.Type {
	var structFields = make([]reflect.StructField, 0)

	fmt.Println(model)
	for _, attr := range attrs {
		structFields = append(structFields, GenStuctField(attr))
	}
	typ := reflect.StructOf(structFields)
	return typ
}

func GenStuctField(attr model.Attribute) reflect.StructField {
	var sf = reflect.StructField{}

	sf.Name = strings.ToUpper(string(attr.Key[0])) + attr.Key[1:]
	sf.Tag = reflect.StructTag(fmt.Sprintf(`json:"%s"`, strings.ToLower(attr.Key)))
	switch attr.Type {
	case "string":
		sf.Type = reflect.TypeOf(string(""))
	case "number":
		sf.Type = reflect.TypeOf(float64(0))
	case "bool":
		sf.Type = reflect.TypeOf(bool(false))
	case "date":
		sf.Type = reflect.TypeOf(string(""))
	case "enum":
		sf.Type = reflect.TypeOf(string("")) // key值
	}

	return sf
}
