package extractor

import (
	"phi-lang/core/base"
)

type ObjectLike interface {
	ClassName() string
	Children() []base.Node
	NodeType() base.NodeType
}

func UnsafeNodeToObjectLike(x base.Node) ObjectLike {
	return x.(ObjectLike)
}
