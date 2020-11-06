package cerealbox

import (
	"errors"
	"fmt"
	"github.com/guregu/null"
	"github.com/ocomsoft/cerealbox/helper"
	"reflect"
)

func (this SerializerToMap) DoFloat64(keyName string, fieldName string, required bool, validator IValidator) ISerializer {
	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.addError(keyName, err)
	} else {
		switch v := fv.Interface().(type) {
		default:
			if required {
				this.addError(keyName, fmt.Errorf("%s is not an Float64 field", fieldName))
			} else {
				this.result[keyName] = fv.Interface() // null.Float64
			}
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
		switch v := fv.Interface().(type) {
		default:
			if required {
				this.addError(keyName, fmt.Errorf("%s is not an Float32 field", fieldName))
			} else {
				this.result[keyName] = fv.Interface() // null.Float32
			}
		case float32, float64:
			this.result[keyName] = v
		}
	}

	return this
}

func (this SerializerFromMap) DoFloat64(keyName string, fieldName string, required bool, validator IValidator) ISerializer {
	if val, ok := this.jsonmap[keyName]; ok {
		fv, err := this.getFieldValue(fieldName)
		if err != nil {
			this.addError(keyName, err)
		} else {
			if fv.CanSet() {
				convertedValue, isNull, err := helper.ConvertToFloat64(val, !required)
				if err != nil {
					this.addError(keyName, err)
				}

				if required && isNull {
					this.addError(keyName, errors.New("required"))
				}

				if required {
					fv.SetFloat(convertedValue)
				} else {
					fv.Set(reflect.ValueOf(null.NewFloat(convertedValue, isNull)))
				}

			}
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
			if fv.CanSet() {
				convertedValue, isNull, err := helper.ConvertToFloat32(val, !required)
				if err != nil {
					this.addError(keyName, err)
				}

				if required && isNull {
					this.addError(keyName, errors.New("required"))
				}

				if required {
					fv.SetFloat(float64(convertedValue))
				} else {
					fv.Set(reflect.ValueOf(null.NewFloat(float64(convertedValue), isNull)))
				}

			}
		}
	} else {
		if required {
			this.addError(keyName, errors.New("required"))
		}
	}

	return this
}
