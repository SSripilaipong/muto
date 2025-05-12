package syntaxtree

import (
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Class struct {
	name string
}

func (Class) PatternType() base.PatternType { return base.PatternTypeClass }

func (Class) DeterminantType() base.DeterminantType { return base.DeterminantTypeClass }

func (Class) RuleResultNodeType() stResult.NodeType { return stResult.NodeTypeClass }

func (Class) ObjectParamType() stResult.ParamType { return stResult.ParamTypeSingle }

func (Class) NonObjectNode() {}

func (c Class) Value() string {
	return c.Name()
}

func (c Class) Name() string {
	return c.name
}

func (c Class) DeterminantName() string {
	return c.Name()
}

func NewClass(name string) Class {
	return Class{name: name}
}

func ClassToName(c Class) string {
	return c.Name()
}

func UnsafeRuleResultToClass(p stResult.Node) Class {
	return p.(Class)
}

func UnsafePatternToClass(p base.Pattern) Class {
	return p.(Class)
}
