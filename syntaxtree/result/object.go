package result

type Object struct {
	head      Node
	paramPart ParamPart
}

func NewObject(head Node, paramPart ParamPart) Object {
	return Object{head: head, paramPart: paramPart}
}

func WrapNodeWithObject(head Node) Object {
	return NewObject(head, ParamsToFixedParamPart([]Param{}))
}

func (Object) RuleResultNodeType() NodeType {
	return NodeTypeObject
}

func (Object) ObjectParamType() ParamType { return ParamTypeSingle }

func (Object) SimplifiedNodeType() SimplifiedNodeType {
	return SimplifiedNodeTypeObject
}
func (obj Object) AsObject() Object {
	return obj
}

func (obj Object) Head() Node {
	return obj.head
}

func (obj Object) ParamPart() ParamPart {
	return obj.paramPart
}

func UnsafeNodeToObject(r Node) Object {
	return r.(Object)
}

func UnsafeSimplifiedNodeToObject(r SimplifiedNode) Object {
	return r.(Object)
}
