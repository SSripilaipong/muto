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
	structureValue := ps.First(ps.ToParser(nonNestedNode), ps.ToParser(nestedNode())).Legacy
	recordTuple := ps.Map(
		mergeStructureRecord,
		ps.ToParser(psBase.IgnoreSpaceBetween3(structureKey, psBase.Colon, structureValue)),
	).Legacy

	structureRecordWithComma := psBase.EndingWithCommaSpaceAllowed(recordTuple)
	optionalCommaSeparatedRecords := psBase.OptionalGreedyRepeatIgnoreWhiteSpaceBetween(structureRecordWithComma)

	optionalTrailingRecord := ps.GreedyOptional(ps.ToParser(psBase.OptionalLeadingWhiteSpace(recordTuple)))
	parser := ps.Map(
		mergeStructure,
		ps.ToParser(psBase.InBracesWhiteSpacesAllowed(ps.Sequence2(
			ps.ToParser(optionalCommaSeparatedRecords),
			optionalTrailingRecord,
		).Legacy)),
	)
	return ps.Filter(noRepeatStructureKeys, parser).Legacy(xs)
}

var structureKey = ps.First(
	ps.ToParser(psBase.BooleanResultNode),
	ps.ToParser(psBase.NumberResultNode),
	ps.ToParser(psBase.StringResultNode),
	ps.ToParser(psBase.RuneResultNode),
	ps.ToParser(psBase.ClassResultNode),
	ps.ToParser(psBase.TagResultNode),
).Legacy

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
