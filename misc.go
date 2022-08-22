package helper

import "reflect"

func Type(x interface{}) reflect.Type {
	return reflect.TypeOf(x)
}

func RemoveElementFromSlice(s interface{}, sl []interface{}) []interface{} {
	ind := 0
	for a := range sl {
		if sl[a] == s {
			ind = a
			break
		}
	}
	sl[ind] = sl[len(sl)-1]
	return sl[:len(sl)-1]
}
