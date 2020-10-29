package validation

import (
	"errors"
)

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
