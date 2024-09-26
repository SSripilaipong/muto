package base

func NewDataObject(children []Node) Object {
	return UnsafeNodeToObject(NewNamedObject("$", children))
}
