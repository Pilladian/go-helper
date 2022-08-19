package helper

import "reflect"

func Type(x interface{}) reflect.Type {
	return reflect.TypeOf(x)
}
