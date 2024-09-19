package base

func NewDataObject(children []Node) NamedObject {
	return UnsafeNodeToNamedObject(NewNamedObject("$", children))
}
