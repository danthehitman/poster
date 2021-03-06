package utilities

import (
	"reflect"
	"fmt"
	"errors"
)

func FillStructFromMap(s interface{}, m map[string]interface{}, ignoreMissingFields bool) (interface{}, error){
	for k, v := range m {
		err := setField(s, k, v, ignoreMissingFields)
		if err != nil {
			return &s, err
		}
	}
	return s, nil
}

func setField(obj interface{}, name string, value interface{}, ignoreMissingfields bool) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		if ignoreMissingfields{return nil}
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() && !ignoreMissingfields {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}
