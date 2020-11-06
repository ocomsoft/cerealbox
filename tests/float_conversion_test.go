package tests

import (
	"github.com/guregu/null"
	"github.com/ocomsoft/cerealbox/helper"
	"testing"
)

func TestConvertToFloat64(t *testing.T) {
	if v, isNull, err := helper.ConvertToFloat64(nil, true); !(isNull == true && v == 0.0 && err == nil) {
		t.Fail() // expect isNull == true
	}

	if v, isNull, err := helper.ConvertToFloat64(true, true); !(isNull == false && v == 1 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64(false, true); !(isNull == false && v == 0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64("10.05", true); !(isNull == false && v == 10.05 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64("ten hundred", true); !(isNull == true && v == 0 && err != nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64(int8(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64(int16(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}
	if v, isNull, err := helper.ConvertToFloat64(int32(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64(int64(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64(uint8(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64(uint16(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64(uint32(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64(uint64(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64(float32(10.07), true); !(isNull == false && v == 10.07 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64(float64(10.07), true); !(isNull == false && v == 10.07 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64(null.FloatFrom(10.07), true); !(isNull == false && v == 10.07 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat64(null.IntFrom(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

}

func TestConvertToFloat32(t *testing.T) {
	if v, isNull, err := helper.ConvertToFloat32(nil, true); !(isNull == true && v == 0.0 && err == nil) {
		t.Fail() // expect isNull == true
	}

	if v, isNull, err := helper.ConvertToFloat32(true, true); !(isNull == false && v == 1 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32(false, true); !(isNull == false && v == 0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32("10.05", true); !(isNull == false && v == 10.05 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32("ten hundred", true); !(isNull == true && v == 0 && err != nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32(int8(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32(int16(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}
	if v, isNull, err := helper.ConvertToFloat32(int32(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32(int64(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32(uint8(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32(uint16(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32(uint32(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32(uint64(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32(float32(10.07), true); !(isNull == false && v == 10.07 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32(float64(10.07), true); !(isNull == false && v == 10.07 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32(null.FloatFrom(10.07), true); !(isNull == false && v == 10.07 && err == nil) {
		t.Fail()
	}

	if v, isNull, err := helper.ConvertToFloat32(null.IntFrom(10), true); !(isNull == false && v == 10.0 && err == nil) {
		t.Fail()
	}

}
