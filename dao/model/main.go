package model

import (
	"reflect"
)

func GetStructs() []reflect.Type {
	typesArr := []reflect.Type{
		reflect.TypeOf(Product{}),
		reflect.TypeOf(Sale{}),
		reflect.TypeOf(Order{}),
		reflect.TypeOf(Employee{}),
	}
	return typesArr
}

