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

func describePlant(plant interface{}) (string, error) {
	t := reflect.TypeOf(plant)

	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("Error: input is not a struct, it's a %T\n", plant)
	}

	var result string
	v := reflect.ValueOf(plant)

	for i := 0; i < t.NumField(); i++ {
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

	return result, nil
}
