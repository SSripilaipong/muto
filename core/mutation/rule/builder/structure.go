package builder

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type structureBuilderFactory struct {
	node nodeBuilderFactory
}

func newStructureBuilderFactory(nodeFactory nodeBuilderFactory) structureBuilderFactory {
	return structureBuilderFactory{node: nodeFactory}
}

func (f structureBuilderFactory) NewBuilder(structure stResult.Structure) mutator.Builder {
	return structureBuilder{recordsBuilder: f.buildStructureRecords(structure)}
}

func (f structureBuilderFactory) buildStructureRecords(structure stResult.Structure) func(mutation *parameter.Parameter) optional.Of[[]base.StructureRecord] {
	recordBuilders := slc.Map(f.buildStructureRecord)(structure.Records())
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

func (f structureBuilderFactory) buildStructureRecord(record stResult.StructureRecord) func(mutation *parameter.Parameter) optional.Of[base.StructureRecord] {
	keyBuilder := f.node.NewBuilder(record.Key())
	valueBuilder := f.node.NewBuilder(record.Value())

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

type structureBuilder struct {
	recordsBuilder func(mutation *parameter.Parameter) optional.Of[[]base.StructureRecord]
}

func (b structureBuilder) Build(param *parameter.Parameter) optional.Of[base.Node] {
	records := b.recordsBuilder(param)
	return optional.Fmap(fn.Compose(base.ToNode, base.NewStructureFromRecords))(records)
}
