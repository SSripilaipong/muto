package result

type Object struct {
	head      Node
	paramPart ParamPart
}

func NewObject(head Node, paramPart ParamPart) Object {
	return Object{head: head, paramPart: paramPart}
}

func (Object) RuleResultNodeType() NodeType {
	return NodeTypeObject
}

func (Object) ObjectParamType() ParamType { return ParamTypeSingle }

func (obj Object) Head() Node {
	return obj.head
}

func (obj Object) ParamPart() ParamPart {
	return obj.paramPart
}

func UnsafeNodeToObject(r Node) Object {
	return r.(Object)
}
