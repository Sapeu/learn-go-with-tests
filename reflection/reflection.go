package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
	val := getValue(x)

	numberOfValues := 0
	var getFeild func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getFeild = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getFeild = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getFeild(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
