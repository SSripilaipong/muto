package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func newForRightVariadicParamPart(pp stPattern.RightVariadicParamPart, nChildrenMatch func(nP int, nC int) bool) func(obj base.Object) optional.Of[*data.Mutation] {
	fixedPart := pp.OtherPart()
	nFixed := len(fixedPart)
	extract := variadicExtractor(pp.Name(), fixedPart, nChildrenMatch)

	return func(obj base.Object) optional.Of[*data.Mutation] {
		children := obj.Children()
		if nFixed < 0 || nFixed > len(children) {
			return optional.Empty[*data.Mutation]()
		}
		return extract(children[:nFixed], children[nFixed:])
	}
}

func newForLeftVariadicParamPart(pp stPattern.LeftVariadicParamPart, nChildrenMatch func(nP int, nC int) bool) func(obj base.Object) optional.Of[*data.Mutation] {
	fixedPart := pp.OtherPart()
	nFixed := len(fixedPart)
	extract := variadicExtractor(pp.Name(), fixedPart, nChildrenMatch)

	return func(obj base.Object) optional.Of[*data.Mutation] {
		children := obj.Children()
		nVariadic := len(children) - nFixed
		if nVariadic < 0 {
			return optional.Empty[*data.Mutation]()
		}
		return extract(children[nVariadic:], children[:nVariadic])
	}
}

func variadicExtractor(name string, fixedParam stPattern.FixedParamPart, nChildrenMatch func(nP int, nC int) bool) func(fixedPart []base.Node, variadicPart []base.Node) optional.Of[*data.Mutation] {
	extractChildren := extractChildrenNodes(fixedParam, nChildrenMatch)

	return func(fixed []base.Node, variadic []base.Node) optional.Of[*data.Mutation] {
		if len(variadic) < 0 {
			return optional.Empty[*data.Mutation]()
		}

		mFixed, ok := extractChildren(fixed).Return()
		if !ok {
			return optional.Empty[*data.Mutation]()
		}

		mVar, ok := data.NewMutation().WithVariadicVarMappings(data.NewVariadicVarMapping(name, variadic)).Return()
		if !ok {
			return optional.Empty[*data.Mutation]()
		}
		return mFixed.Merge(mVar)
	}
}
