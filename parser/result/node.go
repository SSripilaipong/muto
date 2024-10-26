package result

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var RsNode = ps.Map(rslt.Value, Node())

func Node() func([]psBase.Character) []tuple.Of2[stResult.Node, []psBase.Character] {
	return ps.Or(
		nonNestedNode,
		ps.Map(castObjectNode, ps.Filter(objectWithChildren, object)),
		ps.Map(castObjectNode, psBase.InParenthesesWhiteSpaceAllowed(objectMultilines)),
		ps.Map(stResult.ToNode, structure),
	)
}

var nonNestedNode = ps.Or(
	psBase.BooleanResultNode,
	psBase.StringResultNode,
	psBase.NumberResultNode,
	psBase.ClassResultNode,
	psBase.TagResultNode,
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
