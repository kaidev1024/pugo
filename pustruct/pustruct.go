package pustruct

import (
	"fmt"
	"reflect"
)

func SetFieldByName(obj interface{}, name string, value interface{}) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("must pass a pointer to struct")
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("must point to a struct")
	}

	field := v.FieldByName(name)
	if !field.IsValid() {
		return fmt.Errorf("no such field: %s", name)
	}
	if !field.CanSet() {
		return fmt.Errorf("cannot set field %s", name)
	}

	val := reflect.ValueOf(value)

	// Check type compatibility
	if field.Type() != val.Type() {
		return fmt.Errorf("provided value type (%s) didn't match object field type (%s)",
			val.Type(), field.Type())
	}

	field.Set(val)
	return nil
}
