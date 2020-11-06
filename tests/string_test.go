package tests

import (
	"github.com/guregu/null"
	"github.com/ocomsoft/cerealbox"
	"github.com/ocomsoft/cerealbox/validation"
	"testing"
)

type StringExample struct {
	StringVal      string
	StringRequired string
	NullStringVal  null.String
	NullVal        null.String
}

func TestStringsToMap(t *testing.T) {
	example := StringExample{}

	serializerFunc := func(builder cerealbox.ISerializer) cerealbox.ISerializer {
		return builder.
			DoString("StringVal", "StringVal", true, validation.StringVal()).
			DoString("StringRequired", "StringRequired", true, validation.StringVal()).
			DoString("NullStringVal", "NullStringVal", false, validation.StringVal()).
			DoString("NullVal", "NullVal", false, validation.StringVal())

	}

	const STRING_VALUE1 = "HEY"
	const STRING_VALUE2 = "YOU"
	const STRING_VALUE3 = "THERE!"

	example.StringVal = STRING_VALUE1
	example.StringRequired = STRING_VALUE2
	example.NullStringVal = null.StringFrom(STRING_VALUE3)
	example.NullVal = null.NewString("", false)

	json := cerealbox.ToMapWithFunc(&example, serializerFunc)

	if len(json) == 0 {
		t.Fail()
	}

	if json["StringVal"] != STRING_VALUE1 {
		t.Fail()
	}

	if json["StringRequired"] != STRING_VALUE2 {
		t.Fail()
	}

	if (json["NullStringVal"].(null.String)).Valid != true {
		t.Fail()
	}
	if (json["NullStringVal"].(null.String)).String != STRING_VALUE3 {
		t.Fail()
	}

	if (json["NullVal"].(null.String)).Valid != false {
		t.Fail()
	}
}

func TestStringsFromMap(t *testing.T) {
	example := StringExample{}

	serializerFunc := func(builder cerealbox.ISerializer) cerealbox.ISerializer {
		return builder.
			DoString("StringVal", "StringVal", true, validation.StringVal()).
			DoString("StringRequired", "StringRequired", true, validation.StringVal()).
			DoString("NullStringVal", "NullStringVal", false, validation.StringVal()).
			DoString("NullVal", "NullVal", false, validation.StringVal())

	}

	const StringVal1 = "HELLO"
	const StringVal2 = "WORLD"

	json := map[string]interface{}{
		"StringVal":     StringVal1,
		"NullStringVal": StringVal2,
		"NullVal":       nil,
	}

	_, errs := cerealbox.FromMapWithFunc(&example, json, serializerFunc)

	if len(errs) == 0 {
		t.Fail()
	}

	if len(errs["StringRequired"]) != 1 {
		t.Fail()
	}

	if errs["StringRequired"][0].Error() != "required" {
		t.Fail()
	}

	delete(errs, "StringRequired")

	if len(errs) != 0 {
		t.Fail()
	}

	if example.StringVal != StringVal1 {
		t.Fail()
	}

	if example.NullVal.Valid != false {
		t.Fail()
	}

	if example.NullStringVal.Valid != true && example.NullStringVal.String == StringVal2 {
		t.Fail()
	}
}
