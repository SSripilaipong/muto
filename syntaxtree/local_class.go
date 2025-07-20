package syntaxtree

import (
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type LocalClass struct {
	name string
}

func NewLocalClass(name string) LocalClass {
	return LocalClass{name: name}
}

func (c LocalClass) ClassType() ClassType    { return ClassTypeLocal }
func (c LocalClass) Name() string            { return c.name }
func (c LocalClass) DeterminantName() string { return c.Name() }

func (LocalClass) PatternType() base.PatternType         { return base.PatternTypeClass }
func (LocalClass) DeterminantType() base.DeterminantType { return base.DeterminantTypeClass }
func (LocalClass) RuleResultNodeType() stResult.NodeType { return stResult.NodeTypeClass }
func (LocalClass) ObjectParamType() stResult.ParamType   { return stResult.ParamTypeSingle }

var _ Class = LocalClass{}

func LocalClassToName(c LocalClass) string { return c.Name() }

func UnsafeClassToLocalClass(p Class) LocalClass { return p.(LocalClass) }
