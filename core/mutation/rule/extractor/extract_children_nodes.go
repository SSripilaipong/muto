package extractor

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/data"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func extractChildrenNodes(params []stPattern.Param, nChildrenMatch func(int, int) bool) func(children []base.Node) optional.Of[*data.Mutation] {
	nParams := len(params)
	extract := extractWithParamExtractors(params)

	return func(children []base.Node) optional.Of[*data.Mutation] {
		if !nChildrenMatch(nParams, len(children)) {
			return optional.Empty[*data.Mutation]()
		}

		mutation := extract(children)
		if mutation.IsEmpty() {
			return optional.Empty[*data.Mutation]()
		}
		return optional.Value(mutation.Value().WithRemainingChildren(children[nParams:]))
	}
}

func extractWithParamExtractors(params []stPattern.Param) func(children []base.Node) optional.Of[*data.Mutation] {
	paramExtractors := newParamExtractors(params)
	nConsumed := len(paramExtractors)

	return func(children []base.Node) optional.Of[*data.Mutation] {
		mutation := data.NewMutation()
		for i, child := range children[:nConsumed] {
			e := paramExtractors[i](child)
			if e.IsEmpty() {
				return optional.Empty[*data.Mutation]()
			}
			m := mutation.Merge(e.Value())
			if m.IsEmpty() {
				return optional.Empty[*data.Mutation]()
			}
			mutation = m.Value()
		}
		return optional.Value(mutation)
	}
}

func nonStrictlyMatchChildren(nP int, nC int) bool {
	return nP <= nC
}

func strictlyMatchChildren(nP int, nC int) bool {
	return nP == nC
}
