package result

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Node = ps.Or(
	nonNestedNode,
	ps.Map(castObjectNode, ps.Filter(objectWithChildren, object)),
)

var nonNestedNode = ps.Or(
	psBase.BooleanResultNode,
	psBase.StringResultNode,
	psBase.NumberResultNode,
	psBase.ClassResultNode,
	psBase.FixedVarResultNode,
)

func objectWithChildren(obj objectNode) bool {
	param := obj.ParamPart()
	switch {
	case stResult.IsParamPartTypeFixed(param):
		return stResult.UnsafeParamPartToFixedParamPart(param).Size() > 0
	}
	return false
}

func castObjectNode(obj objectNode) stResult.Node {
	return stResult.NewObject(obj.Head(), obj.ParamPart())
}
