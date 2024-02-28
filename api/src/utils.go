package src

import (
	"fmt"
	"reflect"
)

func SetField(obj interface{}, fieldName string, value interface{}) {
	val := reflect.ValueOf(obj).Elem()
	field := val.FieldByName(fieldName)

	if field.IsValid() {
		field.Set(reflect.ValueOf(value))
	}
}

func GetFieldValue(obj interface{}, fieldName string) (interface{}, error) {
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, fmt.Errorf("field '%s' not found in the struct", fieldName)
	}

	return field.Interface(), nil
}
