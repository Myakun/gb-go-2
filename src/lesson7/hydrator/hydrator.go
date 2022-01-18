package hydrator

import (
	"errors"
	"fmt"
	"reflect"
)

func Hydrate(in interface{}, values *map[string]interface{}) error {
	if nil == in {
		return errors.New("param 'in' is nil")
	}

	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return errors.New("param 'in' not a struct")
	}

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		valueFromMap, ok := (*values)[typeField.Name]
		if !ok {
			continue
		}

		fieldVal := val.Field(i)
		if !fieldVal.CanSet() {
			return fmt.Errorf("can't set unexported field %s of struct", typeField.Name)
		}

		valueFromMapVal := reflect.ValueOf(valueFromMap)

		if typeField.Type.String() != valueFromMapVal.Kind().String() {
			return fmt.Errorf("missmatched types for field %s: struct is %s, values map is %s", typeField.Name, typeField.Type.String(), valueFromMapVal.Kind().String())
		}

		fieldVal.Set(valueFromMapVal)
	}

	return nil
}
