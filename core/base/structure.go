package base

import (
	"fmt"
	"strings"

	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/strutil"
)

type Structure struct {
	records []StructureRecord
}

func NewStructureFromRecords(records []StructureRecord) Structure {
	return Structure{records: records}
}

func (Structure) NodeType() NodeType { return NodeTypeStructure }

func (s Structure) MutateAsHead(children []Node, mutation Mutation) optional.Of[Node] {
	return optional.Empty[Node]() // TODO implement
}

func (s Structure) Mutate(mutation Mutation) optional.Of[Node] {
	var records []StructureRecord
	isMutated := false
	for _, record := range s.records {
		newRecord := record.Mutate(mutation)
		if newRecord.IsNotEmpty() {
			isMutated = true
			record = newRecord.Value()
		}
		records = append(records, record)
	}
	if !isMutated {
		return optional.Empty[Node]()
	}
	return optional.Value[Node](NewStructureFromRecords(records))
}

func (s Structure) TopLevelString() string {
	return s.String()
}

func (s Structure) recordSlices() []StructureRecord {
	return s.records
}

func (s Structure) String() string {
	return fmt.Sprintf("{%s}", strings.Join(slc.Map(strutil.ToString[StructureRecord])(s.recordSlices()), ", "))
}

var _ MutableNode = Structure{}

type StructureRecord struct {
	key   Node
	value Node
}

func NewStructureRecord(key Node, value Node) StructureRecord {
	return StructureRecord{key: key, value: value}
}

func (r StructureRecord) Mutate(mutation Mutation) optional.Of[StructureRecord] {
	value := r.Value()
	if IsMutableNode(value) {
		newValue := UnsafeNodeToMutable(value).Mutate(mutation)
		if newValue.IsNotEmpty() {
			return optional.Value(r.replaceValue(newValue.Value()))
		}
	}
	return optional.Empty[StructureRecord]()
}

func (r StructureRecord) Key() Node { return r.key }

func (r StructureRecord) Value() Node { return r.value }

func (r StructureRecord) String() string {
	return fmt.Sprintf("%s: %s", r.Key(), r.Value())
}

func (r StructureRecord) replaceValue(value Node) StructureRecord {
	return NewStructureRecord(r.Key(), value)
}
