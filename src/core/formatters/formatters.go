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
		tag_json := v.Field(i).Tag.Get("json")
		tag_vault := v.Field(i).Tag.Get("vault")
		if tag_vault != "-" {
			m[tag_json] = rv.Field(i).Interface()
		}
	}

	return m
}
