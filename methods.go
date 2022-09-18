package blot

import (
	"github.com/Grant-Eckstein/everglade"
	"log"
)

// MethodFunc represents a function to run against the data
type MethodFunc func(in []byte, parameters Parameters) []byte

type Parameters map[string][]byte

type Method struct {
	Method     MethodFunc
	Parameters Parameters
}

func NewMethod(method MethodFunc, parameters Parameters) Method {
	return Method{
		Method:     method,
		Parameters: parameters,
	}
}

/* Method exports go here */

func (b *Blot) Compress() Method {
	return NewMethod(compressMethodFunc, b.Data)
}

func (b *Blot) Decompress() Method {
	return NewMethod(decompressMethodFunc, b.Data)
}

func (b *Blot) Encrypt() Method {
	eg := everglade.New()

	// Create method parameters
	if b.Data == nil {
		b.Data = make(Parameters)
	}
	b.Data["egJSON"] = eg.Export()

	return NewMethod(encryptMethodFunc, b.Data)
}

func (b *Blot) Decrypt() Method {
	// Assert that config exists
	if b.Data == nil {
		log.Fatalln("Data not set for decryption")
	}

	return NewMethod(decryptMethodFunc, b.Data)
}

/* Actual methods go here */
var compressMethodFunc MethodFunc = func(in []byte, parameters Parameters) []byte {
	return in
}

var decompressMethodFunc MethodFunc = func(in []byte, parameters Parameters) []byte {
	return in
}

var encryptMethodFunc MethodFunc = func(in []byte, parameters Parameters) []byte {

	if parameters == nil {
		parameters = make(Parameters)
	}

	eg := everglade.Import(parameters["egJSON"])
	iv, ct := eg.EncryptCBC(in)
	parameters["iv"] = iv
	return ct
}

var decryptMethodFunc MethodFunc = func(in []byte, parameters Parameters) []byte {
	eg := everglade.Import(parameters["egJSON"])
	pt := eg.DecryptCBC(parameters["iv"], in)
	return pt
}
