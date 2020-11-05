package validation

import (
	"errors"
	"regexp"
)

type StringValFunc func(string) error

type StringValidation struct {
	rules []StringValFunc
}

func StringVal() StringValidation {
	result := StringValidation{
		rules: make([]StringValFunc, 0, 0),
	}

	return result
}

func (this StringValidation) Validate(val interface{}) []error {
	results := make([]error, 0, 0)

	for _, f := range this.rules {
		str := val.(string)

		result := f(str)
		if result != nil {
			results = append(results, result)
		}
	}

	return results
}

func (this StringValidation) MinLength(minLength int) StringValidation {
	this.rules = append(this.rules, func(s string) error {
		if len(s) < minLength {
			return errors.New("minlength")
		}

		return nil
	})

	return this // so we can chain them
}

func (this StringValidation) MaxLength(maxLength int) StringValidation {
	this.rules = append(this.rules, func(s string) error {
		if len(s) > maxLength {
			return errors.New("maxlength")
		}

		return nil
	})

	return this // so we can chain them
}

func (this StringValidation) Email() StringValidation {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	this.rules = append(this.rules, func(s string) error {
		if len(s) < 3 && len(s) > 254 {
			return errors.New("email")
		}
		if !emailRegex.MatchString(s) {
			return errors.New("email")
		}

		return nil
	})

	return this // so we can chain them
}

func (this StringValidation) Func(valFunc StringValFunc) StringValidation {
	this.rules = append(this.rules, valFunc)

	return this // so we can chain them
}
