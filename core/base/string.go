package base

type String struct {
	value string
}

func NewString(value string) String {
	return String{value: value}
}

func (s String) NodeType() NodeType {
	return NodeTypeString
}
