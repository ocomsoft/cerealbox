package cerealbox

import (
	"errors"
	"fmt"
	"github.com/guregu/null"
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
			if required {
				this.result[keyName] = fv.String()
			} else {
				this.result[keyName] = fv.Interface() //null.String
			}
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
			if fv.CanSet() {
				if required {
					fv.SetString(val.(string))
				} else {
					if val == nil {
						fv.Set(reflect.ValueOf(null.NewString("", false)))
					} else {
						fv.Set(reflect.ValueOf(null.StringFrom(val.(string))))
					}
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
