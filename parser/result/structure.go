package result

import (
	"github.com/SSripilaipong/muto/common/optional"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/tuple"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func structure(xs []psBase.Character) []tuple.Of2[stResult.Structure, []psBase.Character] {
	structureValue := ps.Or(
		nonNestedNode,
		ps.Map(castObjectNode, psBase.InParentheses(object)),
		ps.Map(stResult.ToNode, structure),
	)
	record := ps.Map(mergeStructureRecord, psBase.IgnoreSpaceBetween3(structureKey, psBase.Colon, structureValue))

	structureRecordWithComma := psBase.EndingWithCommaSpaceAllowed(record)
	optionalCommaSeparatedRecords := psBase.OptionalGreedyRepeatIgnoreWhiteSpaceBetween(structureRecordWithComma)

	return ps.Map(mergeStructure, psBase.InBracesWhiteSpacesAllowed(ps.Sequence2(
		optionalCommaSeparatedRecords, ps.GreedyOptional(psBase.OptionalLeadingWhiteSpace(record)),
	)))(xs)
}

var structureKey = ps.Or(
	psBase.BooleanResultNode,
	psBase.NumberResultNode,
	psBase.StringResultNode,
	psBase.ClassResultNode,
	psBase.TagResultNode,
)

var mergeStructure = tuple.Fn2(func(xs []stResult.StructureRecord, x optional.Of[stResult.StructureRecord]) stResult.Structure {
	if x.IsNotEmpty() {
		xs = append(xs, x.Value())
	}
	return stResult.NewStructure(xs)
})

var mergeStructureRecord = tuple.Fn3(func(key stResult.Node, _ psBase.Character, value stResult.Node) stResult.StructureRecord {
	return stResult.NewStructureRecord(key, value)
})
