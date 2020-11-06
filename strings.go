package cerealbox

import (
	"errors"
	"fmt"
	"reflect"
)

func (this SerializerToMap) DoString(keyName string, fieldName string, required bool, validator IValidator) ISerializer {
	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.addError(keyName, err)
	} else {
		if fv.Kind() != reflect.String && required {
			this.addError(keyName, fmt.Errorf("%s is not a String field", fieldName))
		} else {
			this.result[keyName] = fv.String()
		}
	}

	return this
}

func (this SerializerFromMap) DoString(keyName string, fieldName string, required bool, validator IValidator) ISerializer {
	if val, ok := this.jsonmap[keyName]; ok {
		fv, err := this.getFieldValue(fieldName)
		if err != nil {
			this.addError(keyName, err)
		} else {
			fv.SetString(val.(string))
		}
	} else {
		if required {
			this.addError(keyName, errors.New("required"))
		}
	}

	return this
}
