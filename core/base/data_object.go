package base

func NewDataObject(children []Node) NamedObject {
	return UnsafeObjectToNamedObject(NewNamedObject("$", children).ConfirmTermination())
}
