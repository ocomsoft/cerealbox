package cerealbox

import (
	"testing"
)

func TestFromMapWithFunc(t *testing.T) {
	example := Example{}

	serializerFunc := func(builder ISerializer) ISerializer {
		return builder.DoString("name", "Name", true, 0, 255).
			DoInt("age", "Age", true, 0, 100).
			DoBool("hidden", "Hide", true)
	}

	json := map[string]interface{}{
		"age":    10,
		"name":   "Rochester Van Jones",
		"hidden": true,
	}

	FromMapWithFunc(&example, json, serializerFunc)

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
