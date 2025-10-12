package pustruct

import (
	"fmt"
	"reflect"
)

func isStructPointer[T any](obj *T) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("must pass a pointer to a struct")
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("must point to a struct")
	}
	return nil
}

func updateFieldByName(v reflect.Value, name string, value any) error {
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
