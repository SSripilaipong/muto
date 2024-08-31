package base

import "fmt"

type String struct {
	value string
}

func NewString(value string) String {
	return String{value: value}
}

func (String) NodeType() NodeType { return NodeTypeString }

func (String) IsTerminationConfirmed() bool { return true }

func (s String) Value() string {
	return s.value
}

func (s String) String() string {
	return fmt.Sprintf("%#v", s.value)
}

func UnsafeNodeToString(n Node) String {
	return n.(String)
}

func StringToNode(s String) Node {
	return s
}
