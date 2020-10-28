package cerealbox

import (
	"errors"
	"fmt"
	"reflect"
)

type SerializerFromMap struct {
	jsonmap map[string]interface{}
	errors  map[string]error
	item    interface{}
}

func (this SerializerFromMap) getFieldValue(fieldName string) (reflect.Value, error) {
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

func FromMapWithFunc(item interface{}, jsonmap map[string]interface{}, serializerFunc SerializerFunc) interface{} {
	serialier := SerializerFromMap{jsonmap: jsonmap, errors: make(map[string]error), item: item}

	serializerFunc(serialier)

	return item
}

func FromMap(serializable ISerializable, jsonmap map[string]interface{}) interface{} {
	return FromMapWithFunc(serializable, jsonmap, serializable.Serialize)
}
