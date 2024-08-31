package base

type ObjectLike interface {
	ClassName() string
	Children() []Node
	NodeType() NodeType
	IsTerminationConfirmed() bool
	ConfirmTermination() ObjectLike
	ReplaceChild(i int, n Node) ObjectLike
}

func UnsafeNodeToObjectLike(x Node) ObjectLike {
	return x.(ObjectLike)
}
