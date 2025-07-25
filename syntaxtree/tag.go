package syntaxtree

import (
	"strings"

	"github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Tag struct {
	value string
}

func NewTag(value string) Tag {
	return Tag{value: value}
}

func (Tag) PatternType() base.PatternType { return base.PatternTypeTag }

func (Tag) RuleResultNodeType() stResult.NodeType { return stResult.NodeTypeTag }

func (Tag) ObjectParamType() stResult.ParamType { return stResult.ParamTypeSingle }

func (t Tag) Value() string {
	return t.value
}

func (t Tag) Name() string {
	name, _ := strings.CutPrefix(t.Value(), ".")
	return name
}

func UnsafeRuleResultToTag(x stResult.Node) Tag { return x.(Tag) }

func UnsafePatternToTag(p base.Pattern) Tag {
	return p.(Tag)
}
