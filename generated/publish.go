// Start fileheader template
// Code generated by go generate; DO NOT EDIT.
// This file was generated by FactomGenerate robots

// Start Generated Code

package generated

// End fileheader template

// Start publisher generated go code

// Publish_Base_IMgs publisher has the basic necessary function implementations.
type Publish_Base_IMgs_type struct {
	*Base
}

// Receive the object of type and call the generic so the compiler can check the passed in type
func (p *Publish_Base_IMgs_type) Write(o IMgs) {
	p.Base.Write(o)
}

func Publish_Base_IMgs(p *Base) Publish_Base_IMgs_type {
	return Publish_Base_IMgs_type{p}
}

// End publisher generated go code
//
// Start filetail template
// Code generated by go generate; DO NOT EDIT.
// End filetail template
// End Generated Code
