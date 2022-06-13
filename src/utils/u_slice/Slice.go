package u_slice

import "reflect"

func Contains(slice interface{}, val interface{}) bool {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(slice)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					return true
				}
			}
		}
	}
	return false
}

func StringsContains(slice []string, val string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
