package cerealbox

import (
	"fmt"
	"reflect"
)

func (this SerializerToMap) DoSlice(keyName string, fieldName string) ISerializer {
	fv, err := this.getFieldValue(fieldName)
	if err != nil {
		this.addError(keyName, err)
	} else {
		if fv.Kind() != reflect.Slice {
			this.addError(keyName, fmt.Errorf("%s is not a Slice field", fieldName))
		} else {
			this.result[keyName] = ToSlice(fv.Interface())
		}
	}

	return this
}

func (this SerializerFromMap) DoSlice(keyName string, fieldName string) ISerializer {
	if val, ok := this.jsonmap[keyName]; ok {
		fv, err := this.getFieldValue(fieldName)
		if err != nil {
			this.addError(keyName, err)
		} else {
			fv.SetString(val.(string))
		}
	}

	return this
}
