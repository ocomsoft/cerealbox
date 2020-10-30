package validation

import "errors"

type FloatValFunc func(float64) error

type FloatValidation struct {
	rules []FloatValFunc
}

func FloatVal() FloatValidation {
	result := FloatValidation{
		rules: make([]FloatValFunc, 0, 0),
	}

	return result
}

func (this FloatValidation) Validate(val interface{}) []error {
	results := make([]error, 0, 0)

	for _, f := range this.rules {
		v := val.(float64)

		result := f(v)
		if result != nil {
			results = append(results, result)
		}
	}

	return results
}

func (this FloatValidation) Min(min float64) FloatValidation {
	this.rules = append(this.rules, func(s float64) error {
		if s < min {
			return errors.New("Min")
		}

		return nil
	})

	return this // so we can chain them
}

func (this FloatValidation) Max(max float64) FloatValidation {
	this.rules = append(this.rules, func(s float64) error {
		if s > max {
			return errors.New("max")
		}

		return nil
	})

	return this // so we can chain them
}

func (this FloatValidation) Func(valFunc FloatValFunc) FloatValidation {
	this.rules = append(this.rules, valFunc)

	return this // so we can chain them
}
