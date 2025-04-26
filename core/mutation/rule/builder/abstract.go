package builder

import (
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type nodeBuilderFactory interface {
	NewBuilder(stResult.Node) mutator.Builder
}

type ClassCollection interface {
	GetClass(name string) base.Class
}
