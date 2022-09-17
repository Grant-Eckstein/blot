package blot

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
		data = method.method(data)
	}
	return data
}
