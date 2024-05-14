package ex01

import (
	"fmt"
	"reflect"
	"strings"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(plant interface{}) string {
	v := reflect.ValueOf(plant)
	t := v.Type()

	var result string

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)

		fieldName := field.Name
		fieldTag := string(field.Tag)
		fieldValue := v.Field(i)

		if fieldTag != "" {
			fieldTag = strings.ReplaceAll(fieldTag, ":", "=")
			fieldTag = strings.ReplaceAll(fieldTag, "\"", "")
			result += fmt.Sprintf("%s(%s):%v\n", fieldName, fieldTag, fieldValue)
		} else {
			result += fmt.Sprintf("%s:%v\n", fieldName, fieldValue)
		}
	}

	return result
}
