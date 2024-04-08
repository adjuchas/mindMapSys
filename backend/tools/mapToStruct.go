package tools

import (
	"errors"
	"reflect"
)

func MapToStruct(m map[string]string, s interface{}) error {
	// 获取结构体的反射值
	structValue := reflect.ValueOf(s)
	if structValue.Kind() != reflect.Ptr || structValue.IsNil() {
		return errors.New("s must be a non-nil pointer to struct")
	}
	// 获取结构体的类型
	structType := structValue.Elem().Type()

	// 确保结构体是指针指向的结构体类型
	if structType.Kind() != reflect.Struct {
		return errors.New("s must be a pointer to struct")
	}

	// 遍历 map 中的键值对
	for k, v := range m {
		// 获取结构体字段
		field, ok := structType.FieldByName(k)
		if !ok {
			continue
		}
		// 如果字段存在且可设置，则设置其值
		if structValue.Elem().FieldByName(k).CanSet() {
			// 将字符串类型的值转换为字段的类型
			fieldValue := reflect.ValueOf(v)
			fieldType := field.Type
			if fieldValue.Type().ConvertibleTo(fieldType) {
				structValue.Elem().FieldByName(k).Set(fieldValue.Convert(fieldType))
			} else {
				return errors.New("value type does not match field type")
			}
		}
	}

	return nil
}
