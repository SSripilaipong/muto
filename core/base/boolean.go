package base

type Boolean struct {
	Object
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
	if !IsObjectNode(x) {
		return false
	}
	obj := UnsafeNodeToObject(x)
	return obj.Equals(NewNamedObject("true", nil)) || obj.Equals(NewNamedObject("false", nil))
}
