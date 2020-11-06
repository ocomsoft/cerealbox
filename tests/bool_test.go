package tests

import (
	"github.com/guregu/null"
	"github.com/ocomsoft/cerealbox"
	"testing"
)

type BoolExample struct {
	TrueVal      bool
	FalseVal     bool
	RequiredVal  bool
	NullTrueVal  null.Bool
	NullFalseVal null.Bool
	NullVal      null.Bool
}

func TestBoolToMap(t *testing.T) {
	example := BoolExample{}

	serializerFunc := func(builder cerealbox.ISerializer) cerealbox.ISerializer {
		return builder.
			DoBool("TrueVal", "TrueVal", true).
			DoBool("FalseVal", "FalseVal", true).
			DoBool("RequiredVal", "RequiredVal", true).
			DoBool("NullTrueVal", "NullTrueVal", false).
			DoBool("NullFalseVal", "NullFalseVal", false).
			DoBool("NullVal", "NullVal", false)
	}

	example.FalseVal = false
	example.TrueVal = true
	example.NullFalseVal = null.BoolFrom(false)
	example.NullTrueVal = null.BoolFrom(true)
	example.NullVal = null.NewBool(false, false)

	json := cerealbox.ToMapWithFunc(&example, serializerFunc)

	if len(json) == 0 {
		t.Fail()
	}

	if json["FalseVal"] != false {
		t.Fail()
	}

	if json["TrueVal"] != true {
		t.Fail()
	}

	if (json["NullTrueVal"].(null.Bool)).Valid != true {
		t.Fail()
	}
	if (json["NullTrueVal"].(null.Bool)).Bool != true {
		t.Fail()
	}

	if (json["NullFalseVal"].(null.Bool)).Valid != true {
		t.Fail()
	}
	if (json["NullFalseVal"].(null.Bool)).Bool != false {
		t.Fail()
	}

	if (json["NullFalseVal"].(null.Bool)).Valid != false {
		t.Fail()
	}
}

func TestBoolFromMap(t *testing.T) {
	example := BoolExample{}

	serializerFunc := func(builder cerealbox.ISerializer) cerealbox.ISerializer {
		return builder.
			DoBool("TrueVal", "TrueVal", true).
			DoBool("FalseVal", "FalseVal", true).
			DoBool("RequiredVal", "RequiredVal", true).
			DoBool("NullTrueVal", "NullTrueVal", false).
			DoBool("NullFalseVal", "NullFalseVal", false).
			DoBool("NullVal", "NullVal", false)
	}

	json := map[string]interface{}{
		"TrueVal":      true,
		"FalseVal":     false,
		"NullTrueVal":  true,
		"NullFalseVal": false,
		"NullVal":      nil,
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

	if example.TrueVal != true {
		t.Fail()
	}

	if example.FalseVal != false {
		t.Fail()
	}

	if example.NullVal.Valid != false {
		t.Fail()
	}

	if example.NullFalseVal.Valid != true && example.NullFalseVal.Bool == false {
		t.Fail()
	}

	if example.NullTrueVal.Valid != true && example.NullTrueVal.Bool == true {
		t.Fail()
	}
}
