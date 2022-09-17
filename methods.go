package blot

// MethodFunc represents a function to run against the data
type MethodFunc func(in []byte) []byte

type Method struct {
	method MethodFunc
}

func NewMethod(method MethodFunc) Method {
	return Method{
		method: method,
	}
}

/* Method exports go here */
func Compress() Method {
	return NewMethod(compressMethodFunc)
}

func Encrypt() Method {
	return NewMethod(encryptMethodFunc)
}

/* Actual methods go here */
var compressMethodFunc MethodFunc = func(in []byte) []byte {
	return in
}

var encryptMethodFunc MethodFunc = func(in []byte) []byte {
	return in
}
