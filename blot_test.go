package blot

import (
	"bytes"
	"testing"
)

func TestBlot_Obfuscate(t *testing.T) {
	var methods []Method

	methods = append(methods, Compress())
	methods = append(methods, Encrypt())

	blot := NewBlot(methods)

	input := []byte("Hello, world")
	output := blot.Run(input)

	if bytes.Compare(input, output) != 0 {
		t.Errorf("Input %v does not match output %v", input, output)
	}
}
