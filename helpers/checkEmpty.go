package helpers

import "reflect"

func HasEmptyValues(input interface{}) bool {
	val := reflect.ValueOf(input)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// Check if the field has a zero value
		if reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			return true
		}
	}

	return false
}
