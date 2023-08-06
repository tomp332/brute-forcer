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

	copyFields(partialValue, destValue.Elem(), map[string]bool{})

	return nil
}

func copyFields(source, dest reflect.Value, visited map[string]bool) {
	sourceType := source.Type()

	for i := 0; i < sourceType.NumField(); i++ {
		partialField := sourceType.Field(i)
		fieldName := partialField.Name
		destField := dest.FieldByName(fieldName)

		if destField.IsValid() {
			if source.Field(i).Kind() == reflect.Struct && destField.Kind() == reflect.Struct {
				// Check if the field is an embedded struct
				if partialField.Anonymous {
					// Avoid infinite recursion with embedded structs
					if !visited[partialField.Type.Name()] {
						visited[partialField.Type.Name()] = true
						copyFields(source.Field(i), dest, visited)
					}
				} else {
					copyFields(source.Field(i), destField, map[string]bool{})
				}
			} else {
				// Handle unexported fields
				if !destField.CanSet() {
					// Try to access unexported fields using reflection
					unexportedField := dest.Field(i)
					if unexportedField.CanSet() {
						unexportedField.Set(source.Field(i))
					}
				} else {
					destField.Set(source.Field(i))
				}
			}
		} else {
			// This means that the source struct does not have the field that the destination struct has
			// We need to iterate over the entire source and see if any fields match the destination
			// If they do, we need to copy them over
			for j := 0; j < sourceType.NumField(); j++ {
				sourceField := sourceType.Field(j)
				if sourceField.Name == fieldName && sourceField.Type == destField.Type() {
					destField.Set(source.Field(j))
					break
				}
			}
		}
	}
}
