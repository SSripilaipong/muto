package result

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var RsSimplifiedNode = ps.Map(rslt.Value, SimplifiedNode())

func SimplifiedNode() func([]psBase.Character) []tuple.Of2[stResult.SimplifiedNode, []psBase.Character] {
	return ps.Or(
		ps.Map(wrapNodeWithNakedObject, nonNestedNode),
		ps.Map(castObjectNodeSimplified, psBase.InParenthesesWhiteSpaceAllowed(NakedObjectMultilines)),
		ps.Map(fn.Compose(castObjectNodeNaked, mergeObject), psBase.SpaceSeparated2(objectHead, objectParamPart)),
		ps.Map(fn.Compose(wrapNodeWithNakedObject, stResult.ToNode[stResult.Structure]), structure),
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

func castNakedObjectNode(obj objectNode) stResult.SimplifiedNode {
	return stResult.NewNakedObject(obj.Head(), obj.ParamPart())
}

func castObjectNode(obj objectNode) stResult.Node {
	return stResult.NewObject(obj.Head(), obj.ParamPart())
}

func castObjectNodeSimplified(obj objectNode) stResult.SimplifiedNode {
	return stResult.NewObject(obj.Head(), obj.ParamPart())
}

func castObjectNodeNaked(obj objectNode) stResult.SimplifiedNode {
	return stResult.NewNakedObject(obj.Head(), obj.ParamPart())
}

var wrapNodeWithNakedObject = fn.Compose(stResult.ToSimplifiedNode, stResult.SingleNodeToNakedObject)
