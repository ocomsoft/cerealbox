package cerealbox

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
)

type Example struct {
	Name        string
	Age         int
	DateOfBirth time.Time
	Hide        bool
}

func (this Example) Serialize(builder ISerializer) ISerializer {
	return builder.DoString("name", "Name", true, StringVal().MinLength(0).MaxLength(255)).
		DoInt("age", "Age", true, IntVal().Min(0).Max(255)).
		DoTime("date_of_birth", "DateOfBirth", true, TimeVal()).
		DoBool("hidden", "Hide", true)
}

func TestISerializableBasic(t *testing.T) {

	example := Example{
		Name:        "Jack Benny",
		Age:         21,
		DateOfBirth: time.Now(),
		Hide:        false}

	enc := json.NewEncoder(os.Stdout)
	err := enc.Encode(ToMap(&example))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestWithFunc(t *testing.T) {
	example := Example{
		Name:        "Jack Benny",
		Age:         21,
		DateOfBirth: time.Now(),
		Hide:        false}

	serializerFunc := func(builder ISerializer) ISerializer {
		return builder.DoString("name", "Name", true, StringVal().MinLength(0).MaxLength(255)).
			DoInt("age", "Age", true, nil)
	}

	enc := json.NewEncoder(os.Stdout)
	err := enc.Encode(ToMapWithFunc(&example, serializerFunc))
	if err != nil {
		fmt.Println(err.Error())
	}
}
