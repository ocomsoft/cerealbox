package cerealbox

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func (this SerializerToMap) DoInt(keyName string, fieldName string, required bool, validator IValidator) ISerializer {

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

type IntValFunc func(int) error

type IntValidation struct {
	rules []IntValFunc
}

func IntVal() IntValidation {
	result := IntValidation{
		rules: make([]IntValFunc, 0, 0),
	}

	return result
}

func (this IntValidation) Validate(val interface{}) []error {
	results := make([]error, 0, 0)

	for _, f := range this.rules {
		v := val.(int)

		result := f(v)
		if result != nil {
			results = append(results, result)
		}
	}

	return results
}

func (this IntValidation) Min(min int) IntValidation {
	this.rules = append(this.rules, func(s int) error {
		if s < min {
			return errors.New("Min")
		}

		return nil
	})

	return this // so we can chain them
}

func (this IntValidation) Max(max int) IntValidation {
	this.rules = append(this.rules, func(s int) error {
		if s > max {
			return errors.New("max")
		}

		return nil
	})

	return this // so we can chain them
}

func (this IntValidation) Func(valFunc IntValFunc) IntValidation {
	this.rules = append(this.rules, valFunc)

	return this // so we can chain them
}
