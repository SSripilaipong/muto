package builder

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type structureBuilder struct {
	recordsBuilder func(mutation *parameter.Parameter) optional.Of[[]base.StructureRecord]
}

func newStructureBuilder(structure stResult.Structure) structureBuilder {
	return structureBuilder{recordsBuilder: buildStructureRecords(structure)}
}

func (b structureBuilder) Build(param *parameter.Parameter) optional.Of[base.Node] {
	records := b.recordsBuilder(param)
	return optional.Fmap(fn.Compose(base.ToNode, base.NewStructureFromRecords))(records)
}

func buildStructureRecords(structure stResult.Structure) func(mutation *parameter.Parameter) optional.Of[[]base.StructureRecord] {
	recordBuilders := slc.Map(buildStructureRecord)(structure.Records())
	return func(mutation *parameter.Parameter) optional.Of[[]base.StructureRecord] {
		var records []base.StructureRecord
		for _, build := range recordBuilders {
			record := build(mutation)
			if record.IsEmpty() {
				return optional.Empty[[]base.StructureRecord]()
			}
			records = append(records, record.Value())
		}
		return optional.Value(records)
	}
}

func buildStructureRecord(record stResult.StructureRecord) func(mutation *parameter.Parameter) optional.Of[base.StructureRecord] {
	keyBuilder := NewLiteralWithoutCarry(record.Key())
	valueBuilder := NewLiteralWithoutCarry(record.Value())

	return func(mutation *parameter.Parameter) optional.Of[base.StructureRecord] {
		key := keyBuilder.Build(mutation)
		if key.IsEmpty() {
			return optional.Empty[base.StructureRecord]()
		}
		value := valueBuilder.Build(mutation)
		if value.IsEmpty() {
			return optional.Empty[base.StructureRecord]()
		}
		return optional.Value(base.NewStructureRecord(key.Value(), value.Value()))
	}
}
