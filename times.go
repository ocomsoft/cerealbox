package cerealbox

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func (this SerializerToMap) DoTime(keyName string, fieldName string, required bool, validator IValidator) ISerializer {
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

func (this SerializerFromMap) DoTime(keyName string, fieldName string, required bool, validator IValidator) ISerializer {
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
