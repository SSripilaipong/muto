package base

import (
	"strings"

	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/strutil"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var VariadicVarWithUnderscore = ps.Map(newVariadicVar,
	ps.Map(tuple.Fn2(strutil.Concat), ps.Sequence2(identifierStartingWithUpperCaseAndUnderscore, ThreeDots)),
)

var VariadicVar = ps.Map(newVariadicVar,
	ps.Map(tuple.Fn2(strutil.Concat), ps.Sequence2(identifierStartingWithUpperCase, ThreeDots)),
)

var VariadicVarResultNode = ps.Map(variadicVarToResultNode, VariadicVar)

type VariadicVarNode struct {
	name string
}

func (v VariadicVarNode) Name() string {
	return v.name
}

func newVariadicVar(name string) VariadicVarNode {
	return VariadicVarNode{name: strings.Trim(name, ".")}
}

func variadicVarToResultNode(x VariadicVarNode) stResult.Param {
	return stResult.NewVariadicVariable(x.Name())
}
