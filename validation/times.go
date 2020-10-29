package validation

import (
	"errors"
	"time"
)

type TimeValFunc func(time.Time) error

type TimeValidation struct {
	rules []TimeValFunc
}

func TimeVal() TimeValidation {
	result := TimeValidation{
		rules: make([]TimeValFunc, 0, 0),
	}

	return result
}

func (this TimeValidation) Validate(val interface{}) []error {
	results := make([]error, 0, 0)

	for _, f := range this.rules {
		v := val.(time.Time)

		result := f(v)
		if result != nil {
			results = append(results, result)
		}
	}

	return results
}

func (this TimeValidation) Min(min time.Time) TimeValidation {
	this.rules = append(this.rules, func(s time.Time) error {
		if min.Before(s) {
			return errors.New("Min")
		}

		return nil
	})

	return this // so we can chain them
}

func (this TimeValidation) Max(max time.Time) TimeValidation {
	this.rules = append(this.rules, func(s time.Time) error {
		if s.After(max) {
			return errors.New("max")
		}

		return nil
	})

	return this // so we can chain them
}
func (this TimeValidation) Func(valFunc TimeValFunc) TimeValidation {
	this.rules = append(this.rules, valFunc)

	return this // so we can chain them
}
