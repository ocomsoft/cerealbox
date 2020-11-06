package cerealbox

import (
	"errors"
	"fmt"
	"strconv"
)

func (this SerializerToMap) DoInt(keyName string, fieldName string, required bool, validator IValidator) ISerializer {

	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.addError(keyName, err)
	} else {
		switch v := fv.Interface().(type) {
		default:
			this.addError(keyName, fmt.Errorf("%s is not an int field", fieldName))
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
			this.addError(keyName, err)
		} else {
			switch x := val.(type) {
			default:
				this.addError(keyName, errors.New("Unknown type"))
			case string:
				val, err := strconv.ParseInt(val.(string), 10, 64)
				if err != nil {
					this.addError(keyName, err)
				} else {
					fv.SetInt(val)
				}
			case int:
				fv.SetInt(int64(x))
			case int64:
				fv.SetInt(x)
			case float32:
				fv.SetInt(int64(int(x)))
			case float64:
				fv.SetInt(int64(int(x)))
			}
		}
	} else {
		if required {
			this.addError(keyName, errors.New("required"))
		}
	}

	return this
}

func (this SerializerToMap) DoUint(keyName string, fieldName string, required bool, validator IValidator) ISerializer {

	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.addError(keyName, err)
	} else {
		switch v := fv.Interface().(type) {
		default:
			this.addError(keyName, fmt.Errorf("%s is not an int field", fieldName))
		case int, int8, int16, int32, int64:
			this.result[keyName] = v
		case uint, uint8, uint16, uint32, uint64:
			this.result[keyName] = v
		}
	}

	return this
}

func (this SerializerFromMap) DoUint(keyName string, fieldName string, required bool, validator IValidator) ISerializer {

	if val, ok := this.jsonmap[keyName]; ok {
		fv, err := this.getFieldValue(fieldName)
		if err != nil {
			this.addError(keyName, err)
		} else {
			switch x := val.(type) {
			default:
				this.addError(keyName, errors.New("Unknown type"))
			case string:
				val, err := strconv.ParseInt(val.(string), 10, 64)
				if err != nil {
					this.addError(keyName, err)
				} else {
					fv.SetInt(val)
				}
			case uint:
				fv.SetUint(uint64(x))
			case uint8:
				fv.SetUint(uint64(x))
			case uint16:
				fv.SetUint(uint64(x))
			case uint32:
				fv.SetUint(uint64(x))
			case uint64:
				fv.SetUint(x)
			case int:
				fv.SetUint(uint64(x))
			case int64:
				fv.SetUint(uint64(x))
			case float32:
				fv.SetUint(uint64(int(x)))
			case float64:
				fv.SetUint(uint64(int(x)))
			}
		}
	} else {
		if required {
			this.addError(keyName, errors.New("required"))
		}
	}

	return this
}
