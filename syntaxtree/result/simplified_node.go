package result

type SimplifiedNode interface {
	Node
	SimplifiedNodeType() SimplifiedNodeType
	AsObject() Object
}

type SimplifiedNodeType string

const (
	SimplifiedNodeTypeObject      SimplifiedNodeType = "OBJECT"
	SimplifiedNodeTypeNakedObject SimplifiedNodeType = "NAKED_OBJECT"
)

func IsSimplifiedNodeTypeObject(x SimplifiedNode) bool {
	return x.SimplifiedNodeType() == SimplifiedNodeTypeObject
}

func IsSimplifiedNodeTypeNakedObject(x SimplifiedNode) bool {
	return x.SimplifiedNodeType() == SimplifiedNodeTypeNakedObject
}

func ToSimplifiedNode[T SimplifiedNode](x T) SimplifiedNode { return x }
