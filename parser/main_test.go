package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
)

func TestParseString(t *testing.T) {
	s := `main A = + 1 "abc"`
	expected := expectedStatements([]st.Statement{
		st.NewRule(
			st.NewNamedRulePattern("main", st.RulePatternFixedParamPart([]st.RuleParamPattern{st.NewVariable("A")})),
			st.NewRuleResultNamedObject("+", st.ObjectFixedParamPart([]st.ObjectParam{st.NewNumber("1"), st.NewString("abc")})),
		),
	})
	assert.Equal(t, expected, FilterSuccess(ParseString(s)))
}

func TestParseVariable(t *testing.T) {
	t.Run("should parse rule with variables with same name", func(t *testing.T) {
		s := `f X X = 1`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				st.NewNamedRulePattern("f", st.RulePatternFixedParamPart([]st.RuleParamPattern{st.NewVariable("X"), st.NewVariable("X")})),
				st.NewNumber("1"),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
}

func TestParseAnonymousObject(t *testing.T) {
	t.Run("should parse anonymous object", func(t *testing.T) {
		s := `main A = (+ 1 2) 3 4`
		expectedParsedTree := st.NewPackage([]st.File{st.NewFile([]st.Statement{
			st.NewRule(
				st.NewNamedRulePattern("main", st.RulePatternFixedParamPart([]st.RuleParamPattern{st.NewVariable("A")})),
				st.NewRuleResultAnonymousObject(
					st.NewRuleResultNamedObject("+", st.ObjectFixedParamPart([]st.ObjectParam{st.NewNumber("1"), st.NewNumber("2")})),
					st.ObjectParamsToObjectFixedParamPart([]st.ObjectParam{st.NewNumber("3"), st.NewNumber("4")}),
				),
			),
		})})
		assert.Equal(t,
			[]tuple.Of2[st.Package, []tk.Token]{tuple.New2(expectedParsedTree, []tk.Token{})},
			FilterSuccess(ParseString(s)),
		)
	})

	t.Run("should parse nested anonymous object", func(t *testing.T) {
		s := `main A = ((+ 1 2) 3 4) 5 6`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				st.NewNamedRulePattern("main", st.RulePatternFixedParamPart([]st.RuleParamPattern{st.NewVariable("A")})),
				st.NewRuleResultAnonymousObject(
					st.NewRuleResultAnonymousObject(st.NewRuleResultNamedObject("+", st.ObjectFixedParamPart([]st.ObjectParam{st.NewNumber("1"), st.NewNumber("2")})), st.ObjectParamsToObjectFixedParamPart([]st.ObjectParam{st.NewNumber("3"), st.NewNumber("4")})),
					st.ObjectParamsToObjectFixedParamPart([]st.ObjectParam{st.NewNumber("5"), st.NewNumber("6")}),
				),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
}

func TestParseObject(t *testing.T) {
	t.Run("should parse object name in object param", func(t *testing.T) {
		s := `main = a b`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				st.NewNamedRulePattern("main", st.RulePatternFixedParamPart([]st.RuleParamPattern{})),
				st.NewRuleResultNamedObject("a", st.ObjectFixedParamPart([]st.ObjectParam{st.NewRuleResultNamedObject("b", nil)})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse nested anonymous object", func(t *testing.T) {
		s := `h ((g 1) X) = 999`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				st.NewNamedRulePattern("h", st.RulePatternFixedParamPart([]st.RuleParamPattern{
					st.NewAnonymousRulePattern(
						st.NewNamedRulePattern("g", st.RulePatternFixedParamPart([]st.RuleParamPattern{st.NewNumber("1")})),
						st.RulePatternFixedParamPart([]st.RuleParamPattern{st.NewVariable("X")}),
					),
				})),
				st.NewNumber("999"),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
}

func TestParseVariadicVarPattern(t *testing.T) {
	t.Run("should parse left variadic var", func(t *testing.T) {
		s := `f Xs... X = g X`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				st.NewNamedRulePattern("f", st.NewRulePatternLeftVariadicParamPart("Xs", []st.RuleParamPattern{st.NewVariable("X")})),
				st.NewRuleResultNamedObject("g", st.ObjectFixedParamPart([]st.ObjectParam{st.NewVariable("X")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse right variadic var", func(t *testing.T) {
		s := `f X Xs... = g X`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				st.NewNamedRulePattern("f", st.NewRulePatternRightVariadicParamPart("Xs", []st.RuleParamPattern{st.NewVariable("X")})),
				st.NewRuleResultNamedObject("g", st.ObjectFixedParamPart([]st.ObjectParam{st.NewVariable("X")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse nested variadic var", func(t *testing.T) {
		s := `f X (g Y Ys...) = g Y`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				st.NewNamedRulePattern("f", st.RulePatternFixedParamPart([]st.RuleParamPattern{st.NewVariable("X"), st.NewNamedRulePattern("g", st.NewRulePatternRightVariadicParamPart("Ys", []st.RuleParamPattern{st.NewVariable("Y")}))})),
				st.NewRuleResultNamedObject("g", st.ObjectFixedParamPart([]st.ObjectParam{st.NewVariable("Y")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
}

func TestParseVariadicParamPart(t *testing.T) {
	t.Run("should parse left variadic param part", func(t *testing.T) {
		s := `f Xs... = g Xs... X`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				st.NewNamedRulePattern("f", st.NewRulePatternLeftVariadicParamPart("Xs", st.RulePatternFixedParamPart{})),
				st.NewRuleResultNamedObject("g", st.ObjectParamsToObjectFixedParamPart([]st.ObjectParam{st.NewVariadicVariable("Xs"), st.NewVariable("X")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse right variadic param part", func(t *testing.T) {
		s := `f Xs... X = g X Xs...`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				st.NewNamedRulePattern("f", st.NewRulePatternLeftVariadicParamPart("Xs", []st.RuleParamPattern{st.NewVariable("X")})),
				st.NewRuleResultNamedObject("g", st.ObjectParamsToObjectFixedParamPart([]st.ObjectParam{st.NewVariable("X"), st.NewVariadicVariable("Xs")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse result with multiple variadic variables in param part", func(t *testing.T) {
		s := `f Xs... X = g Xs... X Xs...`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				st.NewNamedRulePattern("f", st.NewRulePatternLeftVariadicParamPart("Xs", []st.RuleParamPattern{st.NewVariable("X")})),
				st.NewRuleResultNamedObject("g", st.ObjectParamsToObjectFixedParamPart([]st.ObjectParam{st.NewVariadicVariable("Xs"), st.NewVariable("X"), st.NewVariadicVariable("Xs")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
}

func TestVariableRulePattern(t *testing.T) {
	t.Run("should parse variable rule pattern", func(t *testing.T) {
		s := `f (G X) = 1`
		expectedNestedPattern := st.NewVariableRulePattern(
			"G",
			st.RuleParamPatternsToRulePatternFixedParamPart([]st.RuleParamPattern{st.NewVariable("X")}),
		)
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				st.NewNamedRulePattern(
					"f",
					st.RuleParamPatternsToRulePatternFixedParamPart([]st.RuleParamPattern{expectedNestedPattern}),
				),
				st.NewNumber("1"),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
}

func TestActiveRule(t *testing.T) {
	t.Run("should parse active rule", func(t *testing.T) {
		s := `@ f Xs... X = g X Xs...`
		expected := expectedStatements([]st.Statement{
			st.NewActiveRule(
				st.NewNamedRulePattern("f", st.NewRulePatternLeftVariadicParamPart("Xs", []st.RuleParamPattern{st.NewVariable("X")})),
				st.NewRuleResultNamedObject("g", st.ObjectParamsToObjectFixedParamPart([]st.ObjectParam{st.NewVariable("X"), st.NewVariadicVariable("Xs")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
}

func expectedStatements(sts []st.Statement) []tuple.Of2[st.Package, []tk.Token] {
	pkg := st.NewPackage([]st.File{st.NewFile(sts)})
	return []tuple.Of2[st.Package, []tk.Token]{tuple.New2(pkg, []tk.Token{})}
}
