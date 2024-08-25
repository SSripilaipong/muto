package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"phi-lang/common/tuple"
	tk "phi-lang/parser/tokenizer"
	"phi-lang/syntaxtree"
)

func TestParseToken(t *testing.T) {
	tokens := []tk.Token{
		tk.NewToken("hello", tk.Identifier),
		tk.NewToken("X", tk.Identifier),
		tk.NewToken("Y", tk.Identifier),
		tk.NewToken("Z", tk.Identifier),
		tk.NewToken("=", tk.Symbol),
		tk.NewToken("Y", tk.Identifier),
		tk.NewToken("\\n", tk.LineBreak),
		tk.NewToken("main", tk.Identifier),
		tk.NewToken("X", tk.Identifier),
		tk.NewToken("=", tk.Symbol),
		tk.NewToken("hello", tk.Identifier),
		tk.NewToken(`"world"`, tk.String),
		tk.NewToken(`123`, tk.Number),
		tk.NewToken("X", tk.Identifier),
	}
	expectedParsedTree := syntaxtree.NewPackage([]syntaxtree.File{syntaxtree.NewFile([]syntaxtree.Statement{
		syntaxtree.NewRule(syntaxtree.NewRulePattern("hello", []syntaxtree.RuleParamPattern{syntaxtree.NewVariable("X"), syntaxtree.NewVariable("Y"), syntaxtree.NewVariable("Z")}), syntaxtree.NewVariable("Y")),
		syntaxtree.NewRule(syntaxtree.NewRulePattern("main", []syntaxtree.RuleParamPattern{syntaxtree.NewVariable("X")}), syntaxtree.NewRuleResultObject("hello", []syntaxtree.ObjectParam{syntaxtree.NewString("world"), syntaxtree.NewNumber("123"), syntaxtree.NewVariable("X")})),
	})})
	assert.Equal(t,
		[]tuple.Of2[syntaxtree.Package, []tk.Token]{tuple.New2(expectedParsedTree, []tk.Token{})},
		ParseToken(tokens),
	)
}

func TestParseString(t *testing.T) {
	s := `main A = + 1 "abc"`
	expectedParsedTree := syntaxtree.NewPackage([]syntaxtree.File{syntaxtree.NewFile([]syntaxtree.Statement{
		syntaxtree.NewRule(syntaxtree.NewRulePattern("main", []syntaxtree.RuleParamPattern{syntaxtree.NewVariable("A")}), syntaxtree.NewRuleResultObject("+", []syntaxtree.ObjectParam{syntaxtree.NewNumber("1"), syntaxtree.NewString("abc")})),
	})})
	assert.Equal(t,
		[]tuple.Of2[syntaxtree.Package, []tk.Token]{tuple.New2(expectedParsedTree, []tk.Token{})},
		ParseString(s),
	)
}
