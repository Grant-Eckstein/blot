package blot

import (
	"bytes"
	"testing"
)

func TestBlot_Compression(t *testing.T) {
	//var methods []Method

	blot := NewBlot()

	blot.Add(blot.Compress())
	blot.Add(blot.Decompress())

	input := []byte("Hello,world")
	output := blot.Run(input)

	if bytes.Compare(input, output) != 0 {
		t.Errorf("Input %v does not match output %v", input, output)
	}
}

func TestBlot_Encrypt(t *testing.T) {
	blot := NewBlot()

	blot.Add(blot.Encrypt())
	blot.Add(blot.Decrypt())

	input := []byte("Hello, world")
	output := blot.Run(input)

	if bytes.Compare(input, output) != 0 {
		t.Errorf("Input %v does not match output %v", input, output)
	}
}

func TestBlot_Export(t *testing.T) {
	// Create test blot
	blot := NewBlot()

	blot.Add(blot.Encrypt())

	input := []byte("Hello, world")
	output := blot.Run(input)

	// Export blot to JSON
	j := blot.Export()

	// Create new blot from export
	blot2 := Import(j)
	blot2.Add(blot2.Decrypt())
	input2 := blot2.Run(output)

	if bytes.Compare(input, input2) != 0 {
		t.Errorf("Input %v does not match input2 %v", input, input2)
	}
}

func TestBlot_Encode(t *testing.T) {
	blot := NewBlot()

	blot.Add(blot.Encode())
	blot.Add(blot.Decode())

	input := []byte("Hello, world")
	output := blot.Run(input)

	if bytes.Compare(input, output) != 0 {
		t.Errorf("Input %v does not match output %v", input, output)
	}
}

func TestBlot_Layering(t *testing.T) {
	blot := NewBlot()

	blot.Add(blot.Encode())
	blot.Add(blot.Compress())
	blot.Add(blot.Encrypt())

	blot.Add(blot.Decrypt())
	blot.Add(blot.Decompress())
	blot.Add(blot.Decode())

	input := []byte("Hello, world")
	output := blot.Run(input)

	if bytes.Compare(input, output) != 0 {
		t.Errorf("Input %v does not match output %v", input, output)
	}
}
