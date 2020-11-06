package cerealbox

import (
	"errors"
	"fmt"
	"github.com/ocomsoft/cerealbox/validation"
	"reflect"
)

type SerializerToMap struct {
	result map[string]interface{}
	errors validation.ValidationErrors
	item   interface{}
}

func (this SerializerToMap) addError(keyName string, err error) {
	if this.errors == nil {
		this.errors = make(validation.ValidationErrors)
	}

	keyErrors, exist := this.errors[keyName]
	if !exist {
		keyErrors = make([]error, 0, 0)
	}

	this.errors[keyName] = append(keyErrors, err)
}

func (this SerializerToMap) getFieldValue(fieldName string) (reflect.Value, error) {
	/* Get the Field value from the Struct by Name */

	var rv reflect.Value

	switch v := this.item.(type) {
	default:
		rv = reflect.ValueOf(this.item)
	case reflect.Value:
		rv = v
	}

	if rv.Kind() == reflect.Ptr {
		// Dereference pointer
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return reflect.Value{}, errors.New("item must be pointer to struct")
	}

	// Lookup field by name
	fv := rv.FieldByName(fieldName)
	if !fv.IsValid() {
		return reflect.Value{}, fmt.Errorf("not a field name: %s", fieldName)
	}

	// Field must be exported
	if !fv.CanSet() {
		return reflect.Value{}, fmt.Errorf("cannot set field %s", fieldName)
	}

	return fv, nil
}

func ToMapWithFunc(item interface{}, serializerFunc SerializerFunc) map[string]interface{} {
	result := make(map[string]interface{})

	serialier := SerializerToMap{result: result, errors: make(validation.ValidationErrors), item: item}

	serializerFunc(serialier)

	return result
}

func ToMap(item interface{}) map[string]interface{} {
	serializable, ok := item.(ISerializable)
	if ok {
		return ToMapWithFunc(item, serializable.Serialize)
	} else {
		val, ok := item.(reflect.Value)

		serializable, ok := val.Interface().(ISerializable)

		if ok {
			return ToMapWithFunc(item, serializable.Serialize)
		}
	}

	return nil
}

func ToSliceWithFunc(items interface{}, serializerFunc SerializerFunc) []map[string]interface{} {
	s := reflect.ValueOf(items)

	result := make([]map[string]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		item := s.Index(i)
		result[i] = ToMapWithFunc(item, serializerFunc)
	}

	return result
}

func ToSlice(items interface{}) []map[string]interface{} {
	s := reflect.ValueOf(items)

	result := make([]map[string]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		item := s.Index(i)
		result[i] = ToMap(item)
	}

	return result
}
