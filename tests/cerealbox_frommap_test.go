package tests

import (
	"github.com/ocomsoft/cerealbox"
	validation "github.com/ocomsoft/cerealbox/validation"
	"testing"
)

func TestFromMapWithFunc(t *testing.T) {
	example := Example{}

	serializerFunc := func(builder cerealbox.ISerializer) cerealbox.ISerializer {
		return builder.DoString("name", "Name", true, validation.StringVal().MinLength(0).MaxLength(255)).
			DoInt("age", "Age", true, validation.IntVal().Min(0).Max(255)).
			DoBool("hidden", "Hide", true)
	}

	json := map[string]interface{}{
		"age":    10,
		"name":   "Rochester Van Jones",
		"hidden": true,
	}

	cerealbox.FromMapWithFunc(&example, json, serializerFunc)

	if example.Age != 10 {
		t.Fail()
	}
	if example.Name != "Rochester Van Jones" {
		t.Fail()
	}

	if !example.Hide {
		t.Fail()
	}
}
