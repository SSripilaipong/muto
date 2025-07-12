package syntaxtree

import (
	"fmt"
	"strconv"

	"github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Rune struct {
	value string
}

func NewRune(value string) Rune {
	return Rune{value: value}
}

func (Rune) PatternType() base.PatternType { return base.PatternTypeRune }

func (Rune) RuleResultNodeType() stResult.NodeType { return stResult.NodeTypeRune }

func (Rune) ObjectParamType() stResult.ParamType { return stResult.ParamTypeSingle }

func (Rune) NonObjectNode() {}

func (r Rune) Value() string {
	return r.value
}

func (r Rune) RuneValue() rune {
	s, err := strconv.Unquote(r.value)
	if err != nil {
		panic(fmt.Errorf("unexpected error: %w", err))
	}
	x := []rune(s)
	if len(x) != 1 {
		panic(fmt.Errorf("expected 1 rune, got %d", len(x)))
	}
	return x[0]
}

func UnsafeRuleResultToRune(r stResult.Node) Rune {
	return r.(Rune)
}

func UnsafePatternToRune(p base.Pattern) Rune {
	return p.(Rune)
}
