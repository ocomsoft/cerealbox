package tests

import (
	"github.com/guregu/null"
	"github.com/ocomsoft/cerealbox"
	"github.com/ocomsoft/cerealbox/validation"
	"testing"
)

type FloatExample struct {
	Float64Val     float64
	Float32Val     float32
	RequiredVal    float64
	NullFloat64Val null.Float
	NullFloat32Val null.Float
	NullVal        null.Float
}

func TestFloatToMap(t *testing.T) {
	example := FloatExample{}

	serializerFunc := func(builder cerealbox.ISerializer) cerealbox.ISerializer {
		return builder.
			DoFloat64("Float64Val", "Float64Val", true, validation.FloatVal()).
			DoFloat32("Float32Val", "Float32Val", true, validation.FloatVal()).
			DoFloat64("RequiredVal", "RequiredVal", true, validation.FloatVal()).
			DoFloat64("NullFloat64Val", "NullFloat64Val", false, validation.FloatVal()).
			DoFloat32("NullFloat32Val", "NullFloat32Val", false, validation.FloatVal()).
			DoFloat64("NullVal", "NullVal", false, validation.FloatVal())
	}

	example.Float32Val = 32
	example.Float64Val = 64
	example.NullFloat32Val = null.FloatFrom(32)
	example.NullFloat64Val = null.FloatFrom(64)
	example.NullVal = null.NewFloat(0, false)

	json := cerealbox.ToMapWithFunc(&example, serializerFunc)

	if len(json) == 0 {
		t.Fail()
	}

	if json["Float32Val"] != float32(32.0) {
		t.Fail()
	}

	if json["Float64Val"] != 64.0 {
		t.Fail()
	}

	if (json["NullFloat64Val"].(null.Float)).Valid != true {
		t.Fail()
	}
	if (json["NullFloat64Val"].(null.Float)).Float64 != 64.0 {
		t.Fail()
	}

	if (json["NullFloat32Val"].(null.Float)).Valid != true {
		t.Fail()
	}
	if (json["NullFloat32Val"].(null.Float)).Float64 != 32.0 {
		t.Fail()
	}

	if (json["NullVal"].(null.Float)).Valid != false {
		t.Fail()
	}
}

func TestFloatFromMap(t *testing.T) {
	example := FloatExample{}

	serializerFunc := func(builder cerealbox.ISerializer) cerealbox.ISerializer {
		return builder.
			DoFloat64("Float64Val", "Float64Val", true, validation.FloatVal()).
			DoFloat32("Float32Val", "Float32Val", true, validation.FloatVal()).
			DoFloat64("RequiredVal", "RequiredVal", true, validation.FloatVal()).
			DoFloat64("NullFloat64Val", "NullFloat64Val", false, validation.FloatVal()).
			DoFloat32("NullFloat32Val", "NullFloat32Val", false, validation.FloatVal()).
			DoFloat64("NullVal", "NullVal", false, validation.FloatVal())
	}

	json := map[string]interface{}{
		"Float64Val":     64.0,
		"Float32Val":     32.0,
		"NullFloat64Val": 64.0,
		"NullFloat32Val": 32.0,
		"NullVal":        nil,
	}

	_, errs := cerealbox.FromMapWithFunc(&example, json, serializerFunc)

	if len(errs) == 0 {
		t.Fail()
	}

	if len(errs["RequiredVal"]) != 1 {
		t.Fail()
	}

	if errs["RequiredVal"][0].Error() != "required" {
		t.Fail()
	}

	delete(errs, "RequiredVal")

	if len(errs) != 0 {
		t.Fail()
	}

	if example.Float64Val != 64 {
		t.Fail()
	}

	if example.Float32Val != 32 {
		t.Fail()
	}

	if example.NullVal.Valid != false {
		t.Fail()
	}

	if example.NullFloat32Val.Valid != true && example.NullFloat32Val.Float64 == 32 {
		t.Fail()
	}

	if example.NullFloat64Val.Valid != true && example.NullFloat64Val.Float64 == 64 {
		t.Fail()
	}
}
