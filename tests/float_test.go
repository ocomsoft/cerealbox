package tests

import (
	"github.com/guregu/null"
	"github.com/ocomsoft/cerealbox"
	"github.com/ocomsoft/cerealbox/validation"
	"testing"
)

type Float32Example struct {
	FloatVal       float32
	RequiredVal    float32
	NullFloatVal   null.Float
	NullVal        null.Float
	FromIntVal     float32    // converts int64 value to float
	FromNullInt    null.Float // converts int64 value to null.Float
	FromString     float32    // Convert from String
	NullFromString null.Float // Convert from String
}

func Float32ExampleSerial(builder cerealbox.ISerializer) cerealbox.ISerializer {
	return builder.
		DoFloat32("FloatVal", "FloatVal", true, validation.FloatVal()).
		DoFloat32("RequiredVal", "RequiredVal", true, validation.FloatVal()).
		DoFloat32("NullFloatVal", "NullFloatVal", false, validation.FloatVal()).
		DoFloat32("NullVal", "NullVal", false, validation.FloatVal()).
		DoFloat32("FromIntVal", "FromIntVal", true, validation.FloatVal()).
		DoFloat32("FromNullInt", "FromNullInt", false, validation.FloatVal()).
		DoFloat32("FromString", "FromString", true, validation.FloatVal()).
		DoFloat32("NullFromString", "NullFromString", false, validation.FloatVal())
}

type Float64Example struct {
	FloatVal       float64
	RequiredVal    float64
	NullFloatVal   null.Float
	NullVal        null.Float
	FromIntVal     float64    // converts int64 value to float
	FromNullInt    null.Float // converts int64 value to null.Float
	FromString     float64    // Convert from String
	NullFromString null.Float // Convert from String
}

func Float64ExampleSerial(builder cerealbox.ISerializer) cerealbox.ISerializer {
	return builder.
		DoFloat64("FloatVal", "FloatVal", true, validation.FloatVal()).
		DoFloat64("RequiredVal", "RequiredVal", true, validation.FloatVal()).
		DoFloat64("NullFloatVal", "NullFloatVal", false, validation.FloatVal()).
		DoFloat64("NullVal", "NullVal", false, validation.FloatVal()).
		DoFloat64("FromIntVal", "FromIntVal", true, validation.FloatVal()).
		DoFloat64("FromNullInt", "FromNullInt", true, validation.FloatVal()).
		DoFloat64("FromString", "FromString", true, validation.FloatVal()).
		DoFloat64("NullFromString", "NullFromString", true, validation.FloatVal())
}
func TestFloat32ToMap(t *testing.T) {
	example := Float32Example{}

	example.FloatVal = 32
	example.NullFloatVal = null.FloatFrom(32)
	example.NullVal = null.NewFloat(0, false)
	example.FromIntVal = 0                       // not used in this test
	example.FromNullInt = null.NewFloat(0, true) // not used in this test

	json := cerealbox.ToMapWithFunc(&example, Float32ExampleSerial)

	if len(json) == 0 {
		t.Fail()
	}

	if json["FloatVal"] != float32(32.0) {
		t.Fail()
	}

	if (json["NullFloatVal"].(null.Float)).Valid != true {
		t.Fail()
	}
	if (json["NullFloatVal"].(null.Float)).Float64 != 32.0 {
		t.Fail()
	}

	if (json["NullVal"].(null.Float)).Valid != false {
		t.Fail()
	}

	if (json["NullVal"].(null.Float)).Valid != false {
		t.Fail()
	}
}

func TestFloat64ToMap(t *testing.T) {
	example := Float64Example{}

	example.FloatVal = 32
	example.NullFloatVal = null.FloatFrom(32)
	example.NullVal = null.NewFloat(0, false)
	example.FromIntVal = 0                       // not used in this test
	example.FromNullInt = null.NewFloat(0, true) // not used in this test

	json := cerealbox.ToMapWithFunc(&example, Float64ExampleSerial)

	if len(json) == 0 {
		t.Fail()
	}

	if json["FloatVal"] != float32(32.0) {
		t.Fail()
	}

	if (json["NullFloatVal"].(null.Float)).Valid != true {
		t.Fail()
	}
	if (json["NullFloatVal"].(null.Float)).Float64 != 32.0 {
		t.Fail()
	}

	if (json["NullVal"].(null.Float)).Valid != false {
		t.Fail()
	}

	if (json["NullVal"].(null.Float)).Valid != false {
		t.Fail()
	}
}

func TestFloat32FromMap(t *testing.T) {
	example := Float32Example{}

	json := map[string]interface{}{
		"FloatVal":       32.0,
		"NullFloatVal":   32.0,
		"NullVal":        nil,
		"FromIntVal":     128,
		"FromNullInt":    256,
		"FromString":     "1024",
		"NullFromString": null.StringFrom("100"),
	}

	_, errs := cerealbox.FromMapWithFunc(&example, json, Float32ExampleSerial)

	if len(errs) == 0 {
		t.Fail()
	}

	//RequiredVal is missing!
	if len(errs["RequiredVal"]) != 1 {
		t.Fail()
	}

	if errs["RequiredVal"][0].Error() != "required" {
		t.Fail()
	}

	delete(errs, "RequiredVal")

	// RequiredVal is the only field with an error
	if len(errs) != 0 {
		t.Fail()
	}

	if example.FloatVal != 32.0 {
		t.Fail()
	}

	if example.NullVal.Valid != false {
		t.Fail()
	}

	if example.NullFloatVal.Valid != true && example.NullFloatVal.Float64 == 32 {
		t.Fail()
	}

	// Was the integer converted to float here
	if example.FromIntVal != 128.0 {
		t.Fail()
	}

	if example.FromNullInt.Valid != true && example.FromNullInt.Float64 == 256 {
		t.Fail()
	}

	if example.FromString == 1024 {
		t.Fail()
	}

	if example.NullFromString.Valid != true && example.FromNullInt.Float64 == 100 {
		t.Fail()
	}

}
