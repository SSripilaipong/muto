package base

type Variable struct{}

func (v Variable) NodeType() NodeType {
	return NodeTypeVariable
}
