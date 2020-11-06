package cerealbox

import (
	"errors"
	"fmt"
	"github.com/guregu/null"
	"reflect"
)

func (this SerializerToMap) DoBool(keyName string, fieldName string, required bool) ISerializer {
	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.addError(keyName, err)
	} else {
		if fv.Kind() != reflect.Bool && required {
			this.addError(keyName, fmt.Errorf("%s is not a Bool field", fieldName))
		} else {
			if required {
				this.result[keyName] = fv.Bool()
			} else {
				this.result[keyName] = fv.Interface() //null.Bool
			}
		}
	}

	return this
}

func (this SerializerFromMap) DoBool(keyName string, fieldName string, required bool) ISerializer {
	if val, ok := this.jsonmap[keyName]; ok {
		fv, err := this.getFieldValue(fieldName)
		if err != nil {
			this.addError(keyName, err)
		} else {
			if fv.CanSet() {
				if required {
					fv.SetBool(val.(bool))
				} else {
					if val == nil {
						fv.Set(reflect.ValueOf(null.NewBool(false, false)))
					} else {
						fv.Set(reflect.ValueOf(null.BoolFrom(val.(bool))))
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
