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
