package utils

import (
	"fmt"
	"reflect"
)

func CopyStructFields(partial interface{}, dest interface{}) error {
	partialValue := reflect.ValueOf(partial)
	destValue := reflect.ValueOf(dest)

	if partialValue.Type().Kind() != reflect.Struct {
		return fmt.Errorf("partial argument must be a struct")
	}

	if destValue.Type().Kind() != reflect.Ptr || destValue.Elem().Type().Kind() != reflect.Struct {
		return fmt.Errorf("dest argument must be a pointer to a struct")
	}

	for i := 0; i < partialValue.NumField(); i++ {
		partialField := partialValue.Type().Field(i)
		fieldName := partialField.Name
		destField := destValue.Elem().FieldByName(fieldName)

		if destField.IsValid() {
			destField.Set(partialValue.Field(i))
		}
	}
	return nil
}
