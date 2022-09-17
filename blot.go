package blot

import (
	"encoding/json"
	"log"
)

// Blot represents a new Blot instance
type Blot struct {
	Methods []Method
}

// NewBlot creates a new blot instance
func NewBlot(methods []Method) Blot {
	return Blot{
		Methods: methods,
	}
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
	o, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}

	return o
}

// Import Blot configuration from JSON
func Import(j []byte) Blot {
	var o Blot
	err := json.Unmarshal(j, &o)

	if err != nil {
		log.Fatal(err)
	}

	return o
}
