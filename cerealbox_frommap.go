package cerealbox

import (
	"errors"
	"fmt"
	"github.com/ocomsoft/cerealbox/validation"
	"reflect"
)

type SerializerFromMap struct {
	jsonmap map[string]interface{}
	errors  validation.ValidationErrors
	item    interface{}
}

func (this SerializerFromMap) addError(keyName string, err error) {
	if this.errors == nil {
		this.errors = make(validation.ValidationErrors)
	}

	keyErrors, exist := this.errors[keyName]
	if !exist {
		keyErrors = make([]error, 0, 0)
	}

	this.errors[keyName] = append(keyErrors, err)
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

func FromMapWithFunc(item interface{}, jsonmap map[string]interface{}, serializerFunc SerializerFunc) (interface{}, validation.ValidationErrors) {
	serialier := SerializerFromMap{jsonmap: jsonmap, errors: make(validation.ValidationErrors), item: item}

	serializerFunc(serialier)

	return item, serialier.errors
}

func FromMap(item interface{}, jsonmap map[string]interface{}) (interface{}, validation.ValidationErrors) {
	serializable, ok := item.(ISerializable)
	if ok {
		return FromMapWithFunc(item, jsonmap, serializable.Serialize)
	} else {
		val, ok := item.(reflect.Value)

		serializable, ok := val.Interface().(ISerializable)

		if ok {
			return FromMapWithFunc(item, jsonmap, serializable.Serialize)
		}
	}

	return nil, nil
}
