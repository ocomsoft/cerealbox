package helper

import (
	"errors"
	"github.com/guregu/null"
	"strconv"
)

// Convert the val int a float64, return value and null/nill status
func ConvertToFloat64(val interface{}, nullable bool) (float64, bool, error) {
	if nullable && val == nil {
		return 0, true, nil
	}

	switch v := val.(type) {
	case null.String:
		if nullable && !v.Valid {
			return 0, !v.Valid, nil // null
		}

		return parseStringToFloat64(v.String)
	case string:
		return parseStringToFloat64(v)
	case bool:
		if v {
			return 1.0, false, nil
		} else {
			return 0.0, false, nil
		}
	case int:
		return float64(v), false, nil
	case int8:
		return float64(v), false, nil
	case int16:
		return float64(v), false, nil
	case int32:
		return float64(v), false, nil
	case int64:
		return float64(v), false, nil
	case uint:
		return float64(v), false, nil
	case uint8:
		return float64(v), false, nil
	case uint16:
		return float64(v), false, nil
	case uint32:
		return float64(v), false, nil
	case uint64:
		return float64(v), false, nil
	case float32:
		return float64(v), false, nil
	case float64:
		return v, false, nil
	case null.Float:
		return v.Float64, !v.Valid, nil
	case null.Int:
		return float64(v.Int64), !v.Valid, nil
	default:
		return 0, true, errors.New("Cannot convert value to float64") // null!!
	}
}

func parseStringToFloat64(v string) (float64, bool, error) {
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, true, err
	}
	return f, false, nil
}

// Convert the val int a float32, return value and null/nill status
func ConvertToFloat32(val interface{}, nullable bool) (float32, bool, error) {
	if nullable && val == nil {
		return 0, true, nil
	}

	switch v := val.(type) {
	case null.String:
		if nullable && !v.Valid {
			return 0, !v.Valid, nil // null
		}

		return parseStringToFloat32(v.String)
	case string:
		return parseStringToFloat32(v)
	case bool:
		if v {
			return 1.0, false, nil
		} else {
			return 0.0, false, nil
		}
	case int:
		return float32(v), false, nil
	case int8:
		return float32(v), false, nil
	case int16:
		return float32(v), false, nil
	case int32:
		return float32(v), false, nil
	case int64:
		return float32(v), false, nil
	case uint:
		return float32(v), false, nil
	case uint8:
		return float32(v), false, nil
	case uint16:
		return float32(v), false, nil
	case uint32:
		return float32(v), false, nil
	case uint64:
		return float32(v), false, nil
	case float32:
		return v, false, nil
	case float64:
		return float32(v), false, nil
	case null.Float:
		return float32(v.Float64), !v.Valid, nil
	case null.Int:
		return float32(v.Int64), !v.Valid, nil
	default:
		return 0, true, errors.New("Cannot convert value to float32") // null!!
	}
}

func parseStringToFloat32(v string) (float32, bool, error) {
	f, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return 0, true, err
	}
	return float32(f), false, nil
}
