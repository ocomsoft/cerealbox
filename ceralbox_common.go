package cerealbox

type SerializerFunc func(builder ISerializer) ISerializer

type ISerializable interface {
	Serialize(builder ISerializer) ISerializer
}

type IValidator interface {
	Validate(interface{}) []error
}

type ISerializer interface {
	DoBool(string, string, bool) ISerializer
	DoFloat64(string, string, bool, IValidator) ISerializer
	DoFloat32(string, string, bool, IValidator) ISerializer
	DoString(string, string, bool, IValidator) ISerializer
	DoInt(string, string, bool, IValidator) ISerializer
	DoTime(string, string, bool, IValidator) ISerializer
	DoSlice(string, string) ISerializer
}
