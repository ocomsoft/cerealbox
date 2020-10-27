# cerealbox
Custom JSON Serialisation for Go

# Documentation
For now.. See ceralbox_tomap_test for examples of usage.

# Why?
I needed more flexiablity on how I serialised my Structs into JSON for our Rest API
Coming from Django Rest Framework I missed the power of the Serializers so I have attempted to add the same idea in Go.

I found https://github.com/danhper/structomap but it didn't do all I wanted. It missed parsing Mapt to struct and I also 
figured here is a good place to add in vaidation at the same time.

# Example
Create your struct
```go
type Example struct {
	Name        string
	Age         int
	DateOfBirth time.Time
	Hide        bool
}
``` 

Add a Serialize method to implement ISerializable
```go
func (this Example) Serialize(builder ISerializer) ISerializer {
	return builder.DoString("name", "Name", true, 0, 255).
		DoInt("age", "Age", true, 0, 100).
		DoTime("date_of_birth", "DateOfBirth", true, nil, nil).
		DoBool("hidden", "Hide", true)
}
``` 

This is builder function that handles converting you structure into Json. Note there are validation rules here but they are not used when converting to a Map 

Usage
```go
    example := Example{
		Name:        "Jack Benny",
		Age:         21,
		DateOfBirth: time.Now(),
		Hide:        false}

	map := ToMap(&example)
    //TODO use your favourite JSON library to generate the JSON
```