package result

type NakedObject struct {
	Object
}

func NewNakedObject(head Node, paramPart ParamPart) NakedObject {
	return NakedObject{Object: NewObject(head, paramPart)}
}

func SingleNodeToNakedObject(node Node) NakedObject {
	return NewNakedObject(node, ParamsToFixedParamPart([]Param{}))
}

func (obj NakedObject) AsObject() Object {
	return obj.Object
}

func (NakedObject) SimplifiedNodeType() SimplifiedNodeType {
	return SimplifiedNodeTypeNakedObject
}

func UnsafeSimplifiedNodeToNakedObject(r SimplifiedNode) NakedObject {
	return r.(NakedObject)
}
