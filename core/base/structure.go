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

func (s Structure) MutateAsHead(params ParamChain) optional.Of[Node] {
	newChildren := mutateParamChain(params)
	if newChildren.IsNotEmpty() {
		return optional.Value[Node](NewCompoundObject(s, newChildren.Value()))
	}

	return strictUnaryOp(func(x Node) optional.Of[Node] {
		if !IsObjectNode(x) {
			return optional.Empty[Node]()
		}
		obj := UnsafeNodeToObject(x)
		head := obj.Head()
		if !IsTagNode(head) {
			return optional.Empty[Node]()
		}

		return s.processTag(UnsafeNodeToTag(head), obj.Children())
	})(params)
}

func (s Structure) processTag(tag Tag, params []Node) optional.Of[Node] {
	switch tag.Name() {
	case GetTag.Name():
		return s.processGetCommand(params)
	case SetTag.Name():
		return s.processSetCommand(params)
	}
	return optional.Empty[Node]()
}

func (s Structure) Mutate() optional.Of[Node] {
	var records []StructureRecord
	isMutated := false
	for _, record := range s.records {
		newRecord := record.Mutate()
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

func (s Structure) processGetCommand(children []Node) optional.Of[Node] {
	if len(children) != 1 {
		return optional.Empty[Node]()
	}
	return s.get(children[0])
}

func (s Structure) processSetCommand(children []Node) optional.Of[Node] {
	if len(children) != 2 {
		return optional.Empty[Node]()
	}
	return s.set(children[0], children[1])
}

func (s Structure) get(key Node) optional.Of[Node] {
	for _, record := range s.recordSlices() {
		if NodeEqual(record.Key(), key) {
			return optional.Value(record.Value())
		}
	}
	return optional.Empty[Node]()
}

func (s Structure) set(key Node, value Node) optional.Of[Node] {
	var records []StructureRecord
	isMutated := false
	for _, record := range s.recordSlices() {
		if NodeEqual(record.Key(), key) {
			isMutated = true
			record = record.replaceValue(value)
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

func (r StructureRecord) Mutate() optional.Of[StructureRecord] {
	value := r.Value()
	if IsMutableNode(value) {
		newValue := UnsafeNodeToMutable(value).Mutate()
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
