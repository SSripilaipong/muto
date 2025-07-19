package global

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
)

const mapCode = `
map F A                        = (map' F ($)) A
((map' _ B) (_))               = B
(map' F ($ Ys...)) (T X Xs...) = (map' F ($ Ys... (F X))) (T Xs...)
`

const filterCode = `
filter P A                        = (filter' P ($)) A
((filter' _ B) (_))               = B
(filter' P ($ Ys...)) (T X Xs...) = (match
  \true  [(filter' P ($ Ys... X)) (T Xs...)]
  \false [(filter' P ($ Ys...)  ) (T Xs...)]
) (P X)
`

var parseRunesToStringMutator = NewRuleBasedMutatorFromFunctions("parse-runes-to-string", slc.Pure(objectStrictUnaryOp(func(obj base.Object) optional.Of[base.Node] {
	nodes := obj.ParamChain().MostOuter()

	var result []rune
	lastRuneIndex := -1
	for i, x := range nodes {
		if !base.IsRuneNode(x) {
			break
		}
		result = append(result, base.UnsafeNodeToRune(x).Value())
		lastRuneIndex = i
	}

	var remainder []base.Node
	for _, x := range nodes[lastRuneIndex+1:] {
		remainder = append(remainder, x)
	}

	return optional.Value[base.Node](base.NewStructureFromRecords([]base.StructureRecord{
		base.NewStructureRecord(base.ResultTag, base.NewString(string(result))),
		base.NewStructureRecord(base.RemainderTag, base.NewConventionalList(remainder...)),
	}))
})))
