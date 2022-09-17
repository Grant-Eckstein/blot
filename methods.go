package blot

import "github.com/Grant-Eckstein/everglade"

// MethodFunc represents a function to run against the data
type MethodFunc func(in []byte, parameters MethodParameters) []byte

type Parameters map[string][]byte

type Method struct {
	Method     MethodFunc
	Parameters MethodParameters
}

type MethodParameters struct {
	Needed bool
	Data   Parameters
}

func NewMethodParameters(needed bool, data map[string][]byte) MethodParameters {
	return MethodParameters{
		Needed: needed,
		Data:   data,
	}
}

func NewMethod(method MethodFunc, parameters MethodParameters) Method {
	return Method{
		Method:     method,
		Parameters: parameters,
	}
}

/* Method exports go here */

func Compress() Method {
	parameters := NewMethodParameters(false, Parameters{})
	return NewMethod(compressMethodFunc, parameters)
}

func Encrypt() Method {
	eg := everglade.New()

	// Create method parameters
	parameters := map[string][]byte{
		"egJSON": eg.Export(),
	}
	//parameters["egJSON"] = eg.Export()

	return NewMethod(encryptMethodFunc, NewMethodParameters(true, parameters))
}

/* Actual methods go here */
var compressMethodFunc MethodFunc = func(in []byte, parameters MethodParameters) []byte {
	return in
}

var encryptMethodFunc MethodFunc = func(in []byte, parameters MethodParameters) []byte {
	eg := everglade.Import(parameters.Data["egJSON"])
	eg.EncryptCBC(in)
	return in
}
