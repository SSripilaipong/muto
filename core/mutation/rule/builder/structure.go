package builder

import (
	"fmt"
	"strings"

	"github.com/SSripilaipong/go-common/optional"

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
	return structureBuilder{recordBuilders: f.buildStructureRecords(structure)}
}

func (f structureBuilderFactory) buildStructureRecords(structure stResult.Structure) []recordBuilder {
	return slc.Map(f.buildStructureRecord)(structure.Records())
}

func (f structureBuilderFactory) buildStructureRecord(record stResult.StructureRecord) recordBuilder {
	keyBuilder := f.node.NewBuilder(record.Key())
	valueBuilder := f.node.NewBuilder(record.Value())

	return recordBuilder{keyBuilder: keyBuilder, valueBuilder: valueBuilder}
}

type structureBuilder struct {
	recordBuilders []recordBuilder
}

func (b structureBuilder) Build(param *parameter.Parameter) optional.Of[base.Node] {
	var records []base.StructureRecord
	for _, builder := range b.recordBuilders {
		record, ok := builder.Build(param).Return()
		if !ok {
			return optional.Empty[base.Node]()
		}
		records = append(records, record)
	}
	return optional.Value[base.Node](base.NewStructureFromRecords(records))
}

func (x structureBuilder) VisitClass(f func(base.Class)) {
	for _, record := range x.recordBuilders {
		mutator.VisitClass(f, record)
	}
}

func (b structureBuilder) DisplayString() string {
	var records []string
	for _, x := range b.recordBuilders {
		records = append(records, DisplayString(x)+",")
	}
	return fmt.Sprintf("{%s}", strings.Trim(strings.Join(records, " "), ","))
}

type recordBuilder struct {
	keyBuilder   mutator.Builder
	valueBuilder mutator.Builder
}

func (b recordBuilder) Build(param *parameter.Parameter) optional.Of[base.StructureRecord] {
	key := b.keyBuilder.Build(param)
	if key.IsEmpty() {
		return optional.Empty[base.StructureRecord]()
	}
	value := b.valueBuilder.Build(param)
	if value.IsEmpty() {
		return optional.Empty[base.StructureRecord]()
	}
	return optional.Value(base.NewStructureRecord(key.Value(), value.Value()))
}

func (b recordBuilder) VisitClass(f func(base.Class)) {
	mutator.VisitClass(f, b.keyBuilder)
	mutator.VisitClass(f, b.valueBuilder)
}

func (b recordBuilder) DisplayString() string {
	return fmt.Sprintf("%s: %s", DisplayString(b.keyBuilder), DisplayString(b.valueBuilder))
}
