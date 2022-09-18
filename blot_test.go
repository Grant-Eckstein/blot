package blot

import (
	"bytes"
	"testing"
)

func TestBlot_Compression(t *testing.T) {
	//var methods []Method

	blot := NewBlot()

	blot.addMethod(blot.Compress())
	blot.addMethod(blot.Decompress())

	input := []byte("Hello,world")
	output := blot.Run(input)

	if bytes.Compare(input, output) != 0 {
		t.Errorf("Input %v does not match output %v", input, output)
	}
}

func TestBlot_Encrypt(t *testing.T) {
	blot := NewBlot()

	blot.addMethod(blot.Encrypt())
	blot.addMethod(blot.Decrypt())

	input := []byte("Hello, world")
	output := blot.Run(input)

	if bytes.Compare(input, output) != 0 {
		t.Errorf("Input %v does not match output %v", input, output)
	}
}

func TestBlot_Export(t *testing.T) {
	// Create test blot
	blot := NewBlot()

	blot.addMethod(blot.Encrypt())

	input := []byte("Hello, world")
	output := blot.Run(input)

	// Export blot to JSON
	j := blot.Export()

	// Create new blot from export
	blot2 := Import(j)
	blot2.addMethod(blot2.Decrypt())
	input2 := blot2.Run(output)

	if bytes.Compare(input, input2) != 0 {
		t.Errorf("Input %v does not match input2 %v", input, input2)
	}
}
