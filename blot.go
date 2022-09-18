package blot

import (
	"encoding/json"
	"log"
)

// Blot represents a new Blot instance
type Blot struct {
	Methods []Method
	Data    Parameters
}

// NewBlot creates a new blot instance
func NewBlot() Blot {
	return Blot{}
}

func (b *Blot) addMethod(method Method) {
	b.Methods = append(b.Methods, method)
}

// Run processes data through each method in order
func (b *Blot) Run(data []byte) []byte {
	for _, method := range b.Methods {
		parameters := method.Parameters
		data = method.Method(data, parameters)
	}
	return data
}

// Export Blot configuration to JSON
func (b *Blot) Export() []byte {
	o, err := json.Marshal(b.Data)
	if err != nil {
		log.Fatal(err)
	}

	return o
}

// Import Blot configuration from JSON
func Import(j []byte) Blot {
	var p Parameters
	err := json.Unmarshal(j, &p)

	if err != nil {
		log.Fatal(err)
	}

	o := Blot{
		Methods: nil,
		Data:    p,
	}
	return o
}
