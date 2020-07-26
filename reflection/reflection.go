package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}

		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}
