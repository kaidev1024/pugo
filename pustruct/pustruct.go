package pustruct

import (
	"fmt"
	"reflect"
	"strconv"
)

func UpdateField[T any](obj *T, name string, value any) error {
	if err := isStructPointer(obj); err != nil {
		return err
	}
	return updateField(reflect.ValueOf(obj).Elem(), name, value)
}

func UpdateFields[T any](obj *T, updates map[string]any) error {
	if err := isStructPointer(obj); err != nil {
		return err
	}
	for name, value := range updates {
		if err := updateField(reflect.ValueOf(obj).Elem(), name, value); err != nil {
			return fmt.Errorf("cannot update %s, %w", name, err)
		}
	}
	return nil
}

func UpdateFieldsWithStrings[T any](target *T, data map[string]string) error {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("target must be a non-nil pointer to a struct")
	}

	v = v.Elem()
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("target must point to a struct")
	}

	for key, strVal := range data {
		field := v.FieldByName(key)
		if !field.IsValid() || !field.CanSet() {
			continue
		}

		fieldType := field.Type()
		kind := fieldType.Kind()

		switch kind {
		case reflect.String:
			field.Set(reflect.ValueOf(strVal).Convert(fieldType))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
				field.Set(reflect.ValueOf(intVal).Convert(fieldType))
			}
		case reflect.Float32, reflect.Float64:
			if floatVal, err := strconv.ParseFloat(strVal, 64); err == nil {
				field.Set(reflect.ValueOf(floatVal).Convert(fieldType))
			}
		case reflect.Bool:
			if boolVal, err := strconv.ParseBool(strVal); err == nil {
				field.Set(reflect.ValueOf(boolVal).Convert(fieldType))
			}
		}
	}

	return nil
}

func GetFields[T any](v *T) []any {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	// Handle pointer to struct
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		panic("GetFields: input is not a struct")
	}

	var result []any
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()
		result = append(result, field.Name, value)
	}

	return result
}
