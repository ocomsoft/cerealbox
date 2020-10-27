package cerealbox

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
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

func (this SerializerFromMap) DoInt(keyName string, fieldName string, required bool, min int, max int) ISerializer {

	if val, ok := this.jsonmap[keyName]; ok {
		fv, err := this.getFieldValue(fieldName)
		if err != nil {
			this.errors[fieldName] = err
		} else {
			switch val.(type) {
			default:
				this.errors[fieldName] = errors.New("Unknown type")
			case string:
				val, err := strconv.ParseInt(val.(string), 10, 64)
				if err != nil {
					this.errors[fieldName] = err
				} else {
					fv.SetInt(val)
				}
			case int:
				fv.SetInt(int64(val.(int)))
			case int64:
				fv.SetInt(val.(int64))
			}
		}
	} else {
		if required {
			this.errors[fieldName] = errors.New("required")
		}
	}

	return this
}

func (this SerializerFromMap) DoFloat64(keyName string, fieldName string, required bool, min float64, max float64) ISerializer {
	if val, ok := this.jsonmap[keyName]; ok {
		fv, err := this.getFieldValue(fieldName)
		if err != nil {
			this.errors[fieldName] = err
		} else {
			fv.SetFloat(val.(float64))
		}
	} else {
		if required {
			this.errors[fieldName] = errors.New("required")
		}
	}

	return this
}

func (this SerializerFromMap) DoString(keyName string, fieldName string, required bool, minLength int, maxLength int) ISerializer {
	if val, ok := this.jsonmap[keyName]; ok {
		fv, err := this.getFieldValue(fieldName)
		if err != nil {
			this.errors[fieldName] = err
		} else {
			fv.SetString(val.(string))
		}
	} else {
		if required {
			this.errors[fieldName] = errors.New("required")
		}
	}

	return this
}

func (this SerializerFromMap) DoBool(keyName string, fieldName string, required bool) ISerializer {
	if val, ok := this.jsonmap[keyName]; ok {
		fv, err := this.getFieldValue(fieldName)
		if err != nil {
			this.errors[fieldName] = err
		} else {
			fv.SetBool(val.(bool))
		}
	} else {
		if required {
			this.errors[fieldName] = errors.New("required")
		}
	}

	return this
}

func (this SerializerFromMap) DoTime(keyName string, fieldName string, required bool, min *time.Time, max *time.Time) ISerializer {
	if val, ok := this.jsonmap[keyName]; ok {
		fv, err := this.getFieldValue(fieldName)
		if err != nil {
			this.errors[fieldName] = err
		} else {
			fv.Set(reflect.ValueOf(val))
		}
	} else {
		if required {
			this.errors[fieldName] = errors.New("required")
		}
	}

	return this
}

func FromMapWithFunc(item interface{}, jsonmap map[string]interface{}, serializerFunc SerializerFunc) interface{} {
	serialier := SerializerFromMap{jsonmap: jsonmap, errors: make(map[string]error), item: item}

	serializerFunc(serialier)

	return item
}

func FromMap(serializable ISerializable, jsonmap map[string]interface{}) interface{} {
	return FromMapWithFunc(serializable, jsonmap, serializable.Serialize)
}
