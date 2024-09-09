package base

type Boolean struct {
	NamedObject
}

func NewBoolean(x bool) Boolean {
	var name string
	if x {
		name = "true"
	} else {
		name = "false"
	}
	return Boolean{NewNamedObject(name, nil)}
}

func IsBooleanNode(x Node) bool {
	if !IsNamedObjectNode(x) {
		return false
	}
	name := UnsafeNodeToNamedObject(x).Name()
	return name == "true" || name == "false"
}
