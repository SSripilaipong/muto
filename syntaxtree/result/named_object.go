package result

type NamedObject struct {
	objectName string
	paramPart  ParamPart
}

func NewNamedObject(objectName string, paramPart ParamPart) NamedObject {
	return NamedObject{objectName: objectName, paramPart: paramPart}
}

func (NamedObject) RuleResultNodeType() NodeType { return NodeTypeNamedObject }

func (NamedObject) ObjectParamType() ParamType { return ParamTypeSingle }

func (obj NamedObject) ObjectName() string {
	return obj.objectName
}

func (obj NamedObject) ParamPart() ParamPart {
	return obj.paramPart
}

func UnsafeNodeToNamedObject(r Node) NamedObject {
	return r.(NamedObject)
}
