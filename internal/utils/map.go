package utils

import (
	"reflect"
)

// Create new map from struct by comparing the source struct with specified model
func StructToMap(model, source interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	modelFields := reflect.ValueOf(model)
	srcFields := reflect.ValueOf(source)

	for i := 0; i < modelFields.NumField(); i++ {
		modelF := modelFields.Type().Field(i)
		sourceF := srcFields.FieldByName(modelF.Name)

		if sourceF.IsValid() {
			// When current field is a nested struct
			if modelF.Type.Kind() == reflect.Struct && sourceF.Kind() == reflect.Struct {
				result[modelF.Name] = StructToMap(sourceF.Interface(), modelFields.FieldByName(modelF.Name).Interface())
			} else {
				// If the source field is not empty and
				// not required then assign result
				val := sourceF.Interface()
				if modelF.Tag.Get("binding") != "required" &&
					//modelF.Tag.Get("json") != "-" &&
					val != reflect.Zero(sourceF.Type()).Interface() {
					result[modelF.Name] = val
				}
			}
		}
	}

	return result
}
