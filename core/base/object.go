package base

type Object interface {
	MutableNode
	ObjectType() ObjectType
	ParamChain() ParamChain
	AppendChildren(children []Node) Object
	ChainParams(params ParamChain) Object
	Children() []Node
	Head() Node
	Equals(x Object) bool
	AppendParams(params ParamChain) Object
}

type ObjectType string

const (
	ObjectTypePrimitive ObjectType = "PRIMITIVE"
	ObjectTypeCompound  ObjectType = "COMPOUND"
)

func IsPrimitiveObject(x Object) bool {
	return x.ObjectType() == ObjectTypePrimitive
}

func IsCompoundObject(x Object) bool {
	return x.ObjectType() == ObjectTypeCompound
}

func AppendParams(params ParamChain) func(Object) Object {
	return func(x Object) Object { return x.AppendParams(params) }
}
