package extractor

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func New(rule stPattern.NamedRule) func(obj base.Object) optional.Of[*data.Mutation] {
	return newForParamPart(rule.ParamPart(), nonStrictlyMatchChildren)
}

func newForParamPart(paramPart stPattern.ParamPart, nChildrenMatch func(nP int, nC int) bool) func(obj base.Object) optional.Of[*data.Mutation] {
	switch {
	case stPattern.IsParamPartTypeFixed(paramPart):
		return newForFixedParamPart(stPattern.UnsafeParamPartToParams(paramPart), nChildrenMatch)
	case stPattern.IsParamPartTypeVariadic(paramPart):
		return newForVariadicParamPart(stPattern.UnsafeParamPartToVariadicParamPart(paramPart), nChildrenMatch)
	}
	panic("not implemented")
}

func newForVariadicParamPart(paramPart stPattern.VariadicParamPart, nChildrenMatch func(nP int, nC int) bool) func(obj base.Object) optional.Of[*data.Mutation] {
	switch {
	case stPattern.IsVariadicParamPartTypeRight(paramPart):
		return newForRightVariadicParamPart(stPattern.UnsafeVariadicParamPartToRightVariadicParamPart(paramPart), nChildrenMatch)
	case stPattern.IsVariadicParamPartTypeLeft(paramPart):
		return newForLeftVariadicParamPart(stPattern.UnsafeVariadicParamPartToLeftVariadicParamPart(paramPart), nChildrenMatch)
	}
	panic("not implemented")
}

func newForFixedParamPart(params []stPattern.Param, nChildrenMatch func(nP int, nC int) bool) func(obj base.Object) optional.Of[*data.Mutation] {
	return fn.Compose(extractChildrenNodes(params, nChildrenMatch), base.ObjectToChildren)
}
