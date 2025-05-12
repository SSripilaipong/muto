package result

type Structure struct {
	records []StructureRecord
}

func NewStructure(records []StructureRecord) Structure {
	return Structure{records: records}
}

func (Structure) RuleResultNodeType() NodeType { return NodeTypeStructure }

func (Structure) ObjectParamType() ParamType { return ParamTypeSingle }

func (s Structure) Records() []StructureRecord {
	return s.records
}

var _ Node = Structure{}

type StructureRecord struct {
	key   Node
	value Node
}

func (r StructureRecord) Key() Node {
	return r.key
}

func (r StructureRecord) Value() Node {
	return r.value
}

func NewStructureRecord(key Node, value Node) StructureRecord {
	return StructureRecord{key: key, value: value}
}

func UnsafeNodeToStructure(r Node) Structure {
	return r.(Structure)
}
