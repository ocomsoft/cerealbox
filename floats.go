package cerealbox

import (
	"errors"
	"fmt"
	"reflect"
)

func (this SerializerFromMap) DoFloat64(keyName string, fieldName string, required bool, validator IValidator) ISerializer {
	if val, ok := this.jsonmap[keyName]; ok {
		fv, err := this.getFieldValue(fieldName)
		if err != nil {
			this.addError(keyName, err)
		} else {
			fv.SetFloat(val.(float64))
		}
	} else {
		if required {
			this.addError(keyName, errors.New("required"))
		}
	}

	return this
}

func (this SerializerFromMap) DoFloat32(keyName string, fieldName string, required bool, validator IValidator) ISerializer {
	if val, ok := this.jsonmap[keyName]; ok {
		fv, err := this.getFieldValue(fieldName)
		if err != nil {
			this.addError(keyName, err)
		} else {
			fv.SetFloat(val.(float64))
		}
	} else {
		if required {
			this.addError(keyName, errors.New("required"))
		}
	}

	return this
}

func (this SerializerToMap) DoFloat64(keyName string, fieldName string, required bool, validator IValidator) ISerializer {
	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.addError(keyName, err)
	} else {
		switch v := fv.Interface().(type) {
		default:
			this.addError(keyName, fmt.Errorf("%s is not an Float field", fieldName))
		case float32, float64:
			this.result[keyName] = v
		}
	}

	return this
}

func (this SerializerToMap) DoFloat32(keyName string, fieldName string, required bool, validator IValidator) ISerializer {
	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.addError(keyName, err)
	} else {
		if fv.Kind() != reflect.Float32 {
			this.addError(keyName, fmt.Errorf("%s is not a Float64 field", fieldName))
		} else {
			this.result[keyName] = fv.Float()
		}
	}

	return this
}
