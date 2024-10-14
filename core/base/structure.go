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
	return optional.Empty[Node]() // TODO implement
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

func (r StructureRecord) String() string {
	return fmt.Sprintf("%s: %s", r.key, r.value)
}
