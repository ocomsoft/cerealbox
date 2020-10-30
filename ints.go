package cerealbox

import (
	"errors"
	"fmt"
	"strconv"
)

func (this SerializerToMap) DoInt(keyName string, fieldName string, required bool, validator IValidator) ISerializer {

	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.errors[fieldName] = err
	} else {
		switch v := fv.Interface().(type) {
		default:
			this.errors[fieldName] = fmt.Errorf("%s is not an int field", fieldName)
		case int, int8, int16, int32, int64:
			this.result[keyName] = v
		case uint, uint8, uint16, uint32, uint64:
			this.result[keyName] = v
		}
	}

	return this
}

func (this SerializerFromMap) DoInt(keyName string, fieldName string, required bool, validator IValidator) ISerializer {

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
