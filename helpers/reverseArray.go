package helpers

import "reflect"

func ReverseArray(slice interface{}) interface{} {
	sliceValue := reflect.ValueOf(slice)
	reversedSlice := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len(), sliceValue.Cap())

	for i, j := 0, sliceValue.Len()-1; i < sliceValue.Len(); i, j = i+1, j-1 {
		reversedSlice.Index(j).Set(sliceValue.Index(i))
	}

	return reversedSlice.Interface()
}
