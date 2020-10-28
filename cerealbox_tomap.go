package cerealbox

import (
	"errors"
	"fmt"
	"reflect"
)

type SerializerToMap struct {
	result map[string]interface{}
	errors map[string]error
	item   interface{}
}

func (this SerializerToMap) getFieldValue(fieldName string) (reflect.Value, error) {
	/* Get the Field value from the Struct by Name */

	rv := reflect.ValueOf(this.item)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return reflect.Value{}, errors.New("item must be pointer to struct")
	}

	// Dereference pointer
	rv = rv.Elem()

	// Lookup field by name
	fv := rv.FieldByName(fieldName)
	if !fv.IsValid() {
		return reflect.Value{}, fmt.Errorf("not a field name: %s", fieldName)
	}

	// Field must be exported
	if !fv.CanSet() {
		return reflect.Value{}, fmt.Errorf("cannot get field %s", fieldName)
	}

	return fv, nil
}

func ToMapWithFunc(item interface{}, serializerFunc SerializerFunc) map[string]interface{} {
	result := make(map[string]interface{})

	serialier := SerializerToMap{result: result, errors: make(map[string]error), item: item}

	serializerFunc(serialier)

	return result
}

func ToMap(serializable ISerializable) map[string]interface{} {
	return ToMapWithFunc(serializable, serializable.Serialize)
}
