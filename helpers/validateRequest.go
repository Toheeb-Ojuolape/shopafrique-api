package helpers

import (
	"fmt"
	"reflect"
	"strings"
)

func ValidateRequest(input interface{}) string {
	var missingProps []string

	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()

		if isEmpty(value) {
			missingProps = append(missingProps, field.Name)
		}
	}

	return strings.Join(missingProps, ",")
}

func isEmpty(value interface{}) bool {
	switch v := value.(type) {
	case string:
		return v == ""
	case int:
		return v == 0
	// Add more cases for other types if needed
	default:
		fmt.Printf("Validation for type %T not implemented\n", v)
		return true // Consider unknown types as empty
	}
}
