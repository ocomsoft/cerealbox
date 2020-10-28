package cerealbox

import (
	"errors"
	"fmt"
	"reflect"
)

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
