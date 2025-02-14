package base

import "github.com/SSripilaipong/muto/common/optional"

type Object interface {
	MutableNode
	ObjectType() ObjectType
	ParamChain() ParamChain
	AppendChildren(children []Node) Object
	ChainParams(params ParamChain) Object
	Children() []Node
	Head() Node
	Equals(x Object) bool
	BubbleUp() optional.Of[Node]
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
