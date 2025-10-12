package pustruct

import (
	"fmt"
	"reflect"
)

func UpdateFieldByName[T any](obj *T, name string, value any) error {
	if err := isStructPointer(obj); err != nil {
		return err
	}
	return updateFieldByName(reflect.ValueOf(obj).Elem(), name, value)
}

func UpdateFieldsByName[T any](obj *T, updates map[string]any) error {
	if err := isStructPointer(obj); err != nil {
		return err
	}
	for name, value := range updates {
		if err := updateFieldByName(reflect.ValueOf(obj).Elem(), name, value); err != nil {
			return fmt.Errorf("cannot update %s, %w", name, err)
		}
	}
	return nil
}
