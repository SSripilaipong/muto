package base

func NewDataObject(children []Node) Object {
	return NewNamedObject("$", children)
}
