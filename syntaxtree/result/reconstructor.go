package result

import "github.com/SSripilaipong/muto/syntaxtree/pattern"

type Reconstructor struct {
	extractor pattern.ParamPart
	builder   Object
}

func NewReconstructor(extractor pattern.ParamPart, builder Object) Reconstructor {
	return Reconstructor{
		extractor: extractor,
		builder:   builder,
	}
}

func (Reconstructor) RuleResultNodeType() NodeType { return NodeTypeReconstructor }

func (Reconstructor) ObjectParamType() ParamType { return ParamTypeSingle }

func (r Reconstructor) Extractor() pattern.ParamPart {
	return r.extractor
}

func (r Reconstructor) Builder() Object {
	return r.builder
}

func UnsafeNodeToReconstructor(x Node) Reconstructor { return x.(Reconstructor) }
