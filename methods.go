package blot

import (
	"encoding/base64"
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

	// Keyset does not exist
	if _, ok := b.Data["egJSON"]; !ok {
		eg := everglade.New()
		// Create method parameters
		if b.Data == nil {
			b.Data = make(Parameters)
		}
		b.Data["egJSON"] = eg.Export()
	}

	return NewMethod(encryptMethodFunc, b.Data)
}

func (b *Blot) Decrypt() Method {
	// Assert that config exists
	if _, ok := b.Data["egJSON"]; !ok {
		log.Fatalln("Data not set for decryption")
	}

	return NewMethod(decryptMethodFunc, b.Data)
}

func (b *Blot) Encode() Method {
	return NewMethod(encodeBase64, b.Data)
}

func (b *Blot) Decode() Method {
	return NewMethod(decodeBase64, b.Data)
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

// TODO b64 encode/decode
var encodeBase64 MethodFunc = func(in []byte, parameters Parameters) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(in)))
	base64.StdEncoding.Encode(dst, in)
	return dst
}

var decodeBase64 MethodFunc = func(in []byte, parameters Parameters) []byte {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(in)))
	n, _ := base64.StdEncoding.Decode(dst, in)

	dst = dst[:n]
	return dst
}
