package base

type ObjectLike interface {
	ClassName() string
	Children() []Node
	NodeType() NodeType
	IsTerminated() bool
	Terminate() ObjectLike
}

func UnsafeNodeToObjectLike(x Node) ObjectLike {
	return x.(ObjectLike)
}
