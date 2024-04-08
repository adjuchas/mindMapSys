package tools

import (
	"fmt"
	"reflect"
)

func CopyStructValue(source interface{}, target interface{}) error {
	sValueOf := reflect.ValueOf(source)
	sKind := sValueOf.Kind()
	sEKind := sValueOf.Elem().Kind()

	tValueOf := reflect.ValueOf(target)
	tKind := tValueOf.Kind()
	tEKind := tValueOf.Elem().Kind()
	if sKind != reflect.Ptr || sEKind != reflect.Struct {
		fmt.Errorf("source is err type")
	}

	if tKind != reflect.Ptr || tEKind != reflect.Struct {
		fmt.Errorf("target is err type")
	}

	sTypes := sValueOf.Elem().Type()
	tTypes := tValueOf.Elem().Type()
	for i := 0; i < tTypes.NumField(); i++ {
		targetField := tTypes.Field(i)
		sourceField, ok := sTypes.FieldByName(targetField.Name)
		if !ok {
			continue
		}

		sourceFieldValue := sValueOf.Elem().FieldByName(targetField.Name)
		targetFieldValue := tValueOf.Elem().Field(i)

		if targetFieldValue.CanSet() && sourceField.Type == targetField.Type {
			targetFieldValue.Set(sourceFieldValue)
		}
	}

	return nil
}
