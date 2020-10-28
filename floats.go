package cerealbox

import (
	"errors"
	"fmt"
	"reflect"
)

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

func (this SerializerFromMap) DoFloat32(keyName string, fieldName string, required bool, min float64, max float64) ISerializer {
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

func (this SerializerToMap) DoFloat32(keyName string, fieldName string, required bool, min float64, max float64) ISerializer {
	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.errors[fieldName] = err
	} else {
		if fv.Kind() != reflect.Float32 {
			this.errors[fieldName] = fmt.Errorf("%s is not a Float64 field", fieldName)
		} else {
			this.result[keyName] = fv.Float()
		}
	}

	return this
}
