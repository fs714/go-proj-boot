package pgsql

import (
	"reflect"
)

func GetStructFields(i interface{}) (fields []string) {
	v := reflect.ValueOf(i)
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Tag.Get("db")
		fields = append(fields, key)
	}

	return
}
