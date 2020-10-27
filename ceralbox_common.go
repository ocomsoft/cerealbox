package cerealbox

import "time"

type SerializerFunc func(builder ISerializer) ISerializer

type ISerializable interface {
	Serialize(builder ISerializer) ISerializer
}

type ISerializer interface {
	DoInt(string, string, bool, int, int) ISerializer
	DoFloat64(string, string, bool, float64, float64) ISerializer
	DoString(string, string, bool, int, int) ISerializer
	DoBool(string, string, bool) ISerializer
	DoTime(string, string, bool, *time.Time, *time.Time) ISerializer
}
