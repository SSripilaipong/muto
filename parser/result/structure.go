package result

import (
	"github.com/SSripilaipong/go-common/optional"
	"github.com/SSripilaipong/go-common/rslt"
	"github.com/SSripilaipong/go-common/tuple"

	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func structure(xs []psBase.Character) tuple.Of2[rslt.Of[stResult.Structure], []psBase.Character] {
	structureValue := ps.First(nonNestedNode, nestedNode())
	recordTuple := ps.Map(
		mergeStructureRecord,
		psBase.IgnoreSpaceBetween3(structureKey, psBase.Colon, structureValue),
	)

	structureRecordWithComma := psBase.EndingWithCommaSpaceAllowed(recordTuple)
	optionalCommaSeparatedRecords := psBase.OptionalGreedyRepeatIgnoreWhiteSpaceBetween(structureRecordWithComma)

	optionalTrailingRecord := ps.GreedyOptional(psBase.OptionalLeadingWhiteSpace(recordTuple))
	parser := ps.Map(
		mergeStructure,
		psBase.InBracesWhiteSpacesAllowed(ps.Sequence2(
			optionalCommaSeparatedRecords,
			optionalTrailingRecord,
		)),
	)
	return ps.Filter(noRepeatStructureKeys, parser)(xs)
}

var structureKey = ps.First(
	psBase.BooleanResultNode,
	psBase.NumberResultNode,
	psBase.StringResultNode,
	psBase.RuneResultNode,
	psBase.ClassResultNode,
	psBase.TagResultNode,
)

func noRepeatStructureKeys(x stResult.Structure) bool {
	mem := make(map[stResult.Node]bool) // Node is too loose as a map key
	for _, record := range x.Records() {
		key := record.Key()
		if mem[key] {
			return false
		}
		mem[key] = true
	}
	return true
}

var mergeStructure = tuple.Fn2(func(xs []stResult.StructureRecord, x optional.Of[stResult.StructureRecord]) stResult.Structure {
	if x.IsNotEmpty() {
		xs = append(xs, x.Value())
	}
	return stResult.NewStructure(xs)
})

var mergeStructureRecord = tuple.Fn3(func(key stResult.Node, _ psBase.Character, value stResult.Node) stResult.StructureRecord {
	return stResult.NewStructureRecord(key, value)
})
