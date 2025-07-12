package base

func NewConventionalList(nodes []Node) Object {
	return NewNamedOneLayerObject("$", nodes)
}
