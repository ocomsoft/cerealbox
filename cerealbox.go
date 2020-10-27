package cerealbox

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type SerializerFunc func(builder ISerializer) ISerializer

type ISerializable interface {
	Serialize(builder ISerializer) ISerializer
}

type ISerializer interface {
	DoInt(string, string, bool, int, int) ISerializer
	DoFloat64(string, string, bool, float64, float64) ISerializer
	DoString(string, string, bool, int, int) ISerializer
	DoBool(string, string, bool) ISerializer
	DoTime(string, string, bool, *time.Time, *time.Time) ISerializer
}

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

func (this SerializerToMap) DoInt(keyName string, fieldName string, required bool, min int, max int) ISerializer {

	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.errors[fieldName] = err
	} else {
		if fv.Kind() != reflect.Int {
			this.errors[fieldName] = fmt.Errorf("%s is not an int field", fieldName)
		} else {
			this.result[keyName] = fv.Int()
		}
	}

	return this
}

func (this SerializerToMap) DoFloat64(keyName string, fieldName string, required bool, min float64, max float64) ISerializer {
	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.errors[fieldName] = err
	} else {
		if fv.Kind() != reflect.Float64 {
			this.errors[fieldName] = fmt.Errorf("%s is not a Float64 field", fieldName)
		} else {
			this.result[keyName] = fv.Float()
		}
	}

	return this
}

func (this SerializerToMap) DoString(keyName string, fieldName string, required bool, minLength int, maxLength int) ISerializer {
	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.errors[fieldName] = err
	} else {
		if fv.Kind() != reflect.String {
			this.errors[fieldName] = fmt.Errorf("%s is not a String field", fieldName)
		} else {
			this.result[keyName] = fv.String()
		}
	}

	return this
}

func (this SerializerToMap) DoBool(keyName string, fieldName string, required bool) ISerializer {
	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.errors[fieldName] = err
	} else {
		if fv.Kind() != reflect.Bool {
			this.errors[fieldName] = fmt.Errorf("%s is not a Bool field", fieldName)
		} else {
			this.result[keyName] = fv.Bool()
		}
	}

	return this
}

func (this SerializerToMap) DoTime(keyName string, fieldName string, required bool, min *time.Time, max *time.Time) ISerializer {
	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.errors[fieldName] = err
	} else {
		if fv.Kind() != reflect.Struct && fv.Type() == reflect.TypeOf(time.Time{}) {
			this.errors[fieldName] = fmt.Errorf("%s is not a Time field", fieldName)
		} else {
			this.result[keyName] = fv.Interface().(time.Time)
		}
	}

	return this
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
