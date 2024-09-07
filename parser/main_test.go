package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"muto/common/tuple"
	tk "muto/parser/tokenizer"
	st "muto/syntaxtree"
)

func TestParseString(t *testing.T) {
	s := `main A = + 1 "abc"`
	expectedParsedTree := st.NewPackage([]st.File{st.NewFile([]st.Statement{
		st.NewRule(st.NewRulePattern("main", []st.RuleParamPattern{st.NewVariable("A")}), st.NewRuleResultNamedObject("+", []st.ObjectParam{st.NewNumber("1"), st.NewString("abc")})),
	})})
	assert.Equal(t,
		[]tuple.Of2[st.Package, []tk.Token]{tuple.New2(expectedParsedTree, []tk.Token{})},
		ParseString(s),
	)
}

func TestParseAnonymousObject(t *testing.T) {
	t.Run("should parse anonymous object", func(t *testing.T) {

		s := `main A = (+ 1 2) 3 4`
		expectedParsedTree := st.NewPackage([]st.File{st.NewFile([]st.Statement{
			st.NewRule(st.NewRulePattern("main", []st.RuleParamPattern{st.NewVariable("A")}),
				st.NewRuleResultAnonymousObject(st.NewRuleResultNamedObject("+", []st.ObjectParam{st.NewNumber("1"), st.NewNumber("2")}), []st.ObjectParam{st.NewNumber("3"), st.NewNumber("4")}),
			),
		})})
		assert.Equal(t,
			[]tuple.Of2[st.Package, []tk.Token]{tuple.New2(expectedParsedTree, []tk.Token{})},
			ParseString(s),
		)
	})

	t.Run("should parse nested anonymous object", func(t *testing.T) {
		s := `main A = ((+ 1 2) 3 4) 5 6`
		expectedParsedTree := st.NewPackage([]st.File{st.NewFile([]st.Statement{
			st.NewRule(st.NewRulePattern("main", []st.RuleParamPattern{st.NewVariable("A")}),
				st.NewRuleResultAnonymousObject(
					st.NewRuleResultAnonymousObject(st.NewRuleResultNamedObject("+", []st.ObjectParam{st.NewNumber("1"), st.NewNumber("2")}), []st.ObjectParam{st.NewNumber("3"), st.NewNumber("4")}),
					[]st.ObjectParam{st.NewNumber("5"), st.NewNumber("6")},
				),
			),
		})})
		assert.Equal(t,
			[]tuple.Of2[st.Package, []tk.Token]{tuple.New2(expectedParsedTree, []tk.Token{})},
			ParseString(s),
		)
	})
}

func TestParseObject(t *testing.T) {
	t.Run("should parse object name in object param", func(t *testing.T) {

		s := `main = a b`
		expectedParsedTree := st.NewPackage([]st.File{st.NewFile([]st.Statement{
			st.NewRule(st.NewRulePattern("main", []st.RuleParamPattern{}),
				st.NewRuleResultNamedObject("a", []st.ObjectParam{st.NewRuleResultNamedObject("b", nil)}),
			),
		})})
		assert.Equal(t,
			[]tuple.Of2[st.Package, []tk.Token]{tuple.New2(expectedParsedTree, []tk.Token{})},
			ParseString(s),
		)
	})
}
