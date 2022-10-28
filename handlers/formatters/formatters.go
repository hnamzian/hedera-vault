package formatters

import (
	"reflect"
)

func FormatResponse(s interface{}) map[string]interface{} {
	v := reflect.TypeOf(s)
	rv := reflect.ValueOf(s)
	rv = reflect.Indirect(rv)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	m := make(map[string]interface{})
	for i := 0; i < rv.NumField(); i++ {
		t := v.Field(i).Tag.Get("json")
		if t != "" && t != "-" {
			m[t] = rv.Field(i).Interface()
		}
	}

	return m
}
