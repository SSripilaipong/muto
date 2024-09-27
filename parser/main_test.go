package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestParseString(t *testing.T) {
	s := `main A = + 1 "abc"`
	expected := expectedStatements([]st.Statement{
		st.NewRule(
			stPattern.NewNamedRule("main", stPattern.FixedParamPart([]stPattern.Param{st.NewVariable("A")})),
			stResult.NewObject(st.NewClass("+"), stResult.FixedParamPart([]stResult.Param{st.NewNumber("1"), st.NewString("abc")})),
		),
	})
	assert.Equal(t, expected, FilterSuccess(ParseString(s)))
}

func TestParseVariable(t *testing.T) {
	t.Run("should parse rule with variables with same name", func(t *testing.T) {
		s := `f X X = 1`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("f", stPattern.FixedParamPart([]stPattern.Param{st.NewVariable("X"), st.NewVariable("X")})),
				st.NewNumber("1"),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
}

func TestParseAnonymousObjectPattern(t *testing.T) {
	t.Run("should parse boolean as anonymous head", func(t *testing.T) {
		s := `f (true 456) = 789`
		expectedParsedTree := st.NewPackage([]st.File{st.NewFile([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
					stPattern.NewAnonymousRule(st.NewBoolean("true"), stPattern.ParamsToFixedParamPart([]stPattern.Param{st.NewNumber("456")})),
				})),
				st.NewNumber("789"),
			),
		})})
		assert.Equal(t,
			[]tuple.Of2[st.Package, []tk.Token]{tuple.New2(expectedParsedTree, []tk.Token{})},
			FilterSuccess(ParseString(s)),
		)
	})

	t.Run("should parse string as anonymous head", func(t *testing.T) {
		s := `f ("a" 456) = 789`
		expectedParsedTree := st.NewPackage([]st.File{st.NewFile([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
					stPattern.NewAnonymousRule(st.NewString(`"a"`), stPattern.ParamsToFixedParamPart([]stPattern.Param{st.NewNumber("456")})),
				})),
				st.NewNumber("789"),
			),
		})})
		assert.Equal(t,
			[]tuple.Of2[st.Package, []tk.Token]{tuple.New2(expectedParsedTree, []tk.Token{})},
			FilterSuccess(ParseString(s)),
		)
	})

	t.Run("should parse number as anonymous head", func(t *testing.T) {
		s := `f (123 456) = 789`
		expectedParsedTree := st.NewPackage([]st.File{st.NewFile([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
					stPattern.NewAnonymousRule(st.NewNumber("123"), stPattern.ParamsToFixedParamPart([]stPattern.Param{st.NewNumber("456")})),
				})),
				st.NewNumber("789"),
			),
		})})
		assert.Equal(t,
			[]tuple.Of2[st.Package, []tk.Token]{tuple.New2(expectedParsedTree, []tk.Token{})},
			FilterSuccess(ParseString(s)),
		)
	})

	t.Run("should parse number as anonymous head without children", func(t *testing.T) {
		s := `f (123) = 789`
		expectedParsedTree := st.NewPackage([]st.File{st.NewFile([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{
					stPattern.NewAnonymousRule(st.NewNumber("123"), stPattern.ParamsToFixedParamPart([]stPattern.Param{})),
				})),
				st.NewNumber("789"),
			),
		})})
		assert.Equal(t,
			[]tuple.Of2[st.Package, []tk.Token]{tuple.New2(expectedParsedTree, []tk.Token{})},
			FilterSuccess(ParseString(s)),
		)
	})
}

func TestParseObject(t *testing.T) {
	t.Run("should parse object name in object param", func(t *testing.T) {
		s := `main = a b`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("main", stPattern.FixedParamPart([]stPattern.Param{})),
				stResult.NewObject(st.NewClass("a"), stResult.FixedParamPart([]stResult.Param{st.NewClass("b")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse nested object", func(t *testing.T) {
		s := `h ((g 1) X) = 999`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("h", stPattern.FixedParamPart([]stPattern.Param{
					stPattern.NewAnonymousRule(
						stPattern.NewNamedRule("g", stPattern.FixedParamPart([]stPattern.Param{st.NewNumber("1")})),
						stPattern.FixedParamPart([]stPattern.Param{st.NewVariable("X")}),
					),
				})),
				st.NewNumber("999"),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
	t.Run("should parse object", func(t *testing.T) {
		s := `main A = (+ 1 2) 3 4`
		expectedParsedTree := st.NewPackage([]st.File{st.NewFile([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("main", stPattern.FixedParamPart([]stPattern.Param{st.NewVariable("A")})),
				stResult.NewObject(
					stResult.NewObject(st.NewClass("+"), stResult.FixedParamPart([]stResult.Param{st.NewNumber("1"), st.NewNumber("2")})),
					stResult.ParamsToFixedParamPart([]stResult.Param{st.NewNumber("3"), st.NewNumber("4")}),
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
				stPattern.NewNamedRule("main", stPattern.FixedParamPart([]stPattern.Param{st.NewVariable("A")})),
				stResult.NewObject(
					stResult.NewObject(stResult.NewObject(st.NewClass("+"), stResult.FixedParamPart([]stResult.Param{st.NewNumber("1"), st.NewNumber("2")})), stResult.ParamsToFixedParamPart([]stResult.Param{st.NewNumber("3"), st.NewNumber("4")})),
					stResult.ParamsToFixedParamPart([]stResult.Param{st.NewNumber("5"), st.NewNumber("6")}),
				),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse number as a head", func(t *testing.T) {
		s := `main = 123 456`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("main", stPattern.FixedParamPart([]stPattern.Param{})),
				stResult.NewObject(
					st.NewNumber("123"),
					stResult.ParamsToFixedParamPart([]stResult.Param{st.NewNumber("456")}),
				),
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
				stPattern.NewNamedRule("f", stPattern.NewLeftVariadicParamPart("Xs", []stPattern.Param{st.NewVariable("X")})),
				stResult.NewObject(st.NewClass("g"), stResult.FixedParamPart([]stResult.Param{st.NewVariable("X")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse right variadic var", func(t *testing.T) {
		s := `f X Xs... = g X`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("f", stPattern.NewRightVariadicParamPart("Xs", []stPattern.Param{st.NewVariable("X")})),
				stResult.NewObject(st.NewClass("g"), stResult.FixedParamPart([]stResult.Param{st.NewVariable("X")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse nested variadic var", func(t *testing.T) {
		s := `f X (g Y Ys...) = g Y`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("f", stPattern.FixedParamPart([]stPattern.Param{st.NewVariable("X"), stPattern.NewNamedRule("g", stPattern.NewRightVariadicParamPart("Ys", []stPattern.Param{st.NewVariable("Y")}))})),
				stResult.NewObject(st.NewClass("g"), stResult.FixedParamPart([]stResult.Param{st.NewVariable("Y")})),
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
				stPattern.NewNamedRule("f", stPattern.NewLeftVariadicParamPart("Xs", stPattern.FixedParamPart{})),
				stResult.NewObject(st.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{stResult.NewVariadicVariable("Xs"), st.NewVariable("X")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse right variadic param part", func(t *testing.T) {
		s := `f Xs... X = g X Xs...`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("f", stPattern.NewLeftVariadicParamPart("Xs", []stPattern.Param{st.NewVariable("X")})),
				stResult.NewObject(st.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{st.NewVariable("X"), stResult.NewVariadicVariable("Xs")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse result with multiple variadic variables in param part", func(t *testing.T) {
		s := `f Xs... X = g Xs... X Xs...`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("f", stPattern.NewLeftVariadicParamPart("Xs", []stPattern.Param{st.NewVariable("X")})),
				stResult.NewObject(st.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{stResult.NewVariadicVariable("Xs"), st.NewVariable("X"), stResult.NewVariadicVariable("Xs")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
}

func TestVariableRulePattern(t *testing.T) {
	t.Run("should parse variable rule pattern", func(t *testing.T) {
		s := `f (G X) = 1`
		expectedNestedPattern := stPattern.NewVariableRulePattern(
			"G",
			stPattern.ParamsToFixedParamPart([]stPattern.Param{st.NewVariable("X")}),
		)
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule(
					"f",
					stPattern.ParamsToFixedParamPart([]stPattern.Param{expectedNestedPattern}),
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
				stPattern.NewNamedRule("f", stPattern.NewLeftVariadicParamPart("Xs", []stPattern.Param{st.NewVariable("X")})),
				stResult.NewObject(st.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{st.NewVariable("X"), stResult.NewVariadicVariable("Xs")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
}

func TestBoolean(t *testing.T) {
	t.Run("should parse boolean as a rule result", func(t *testing.T) {
		s := `main = true`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("main", stPattern.ParamsToFixedParamPart([]stPattern.Param{})),
				st.NewBoolean("true"),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse boolean as a rule param", func(t *testing.T) {
		s := `main = f true`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("main", stPattern.ParamsToFixedParamPart([]stPattern.Param{})),
				stResult.NewObject(st.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{st.NewBoolean("true")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse boolean as a rule pattern param", func(t *testing.T) {
		s := `f true = "a"`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("f", stPattern.ParamsToFixedParamPart([]stPattern.Param{st.NewBoolean("true")})),
				st.NewString("a"),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})

	t.Run("should parse boolean as an object head", func(t *testing.T) {
		s := `main = true "a"`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("main", stPattern.ParamsToFixedParamPart([]stPattern.Param{})),
				stResult.NewObject(st.NewBoolean("true"), stResult.ParamsToFixedParamPart([]stResult.Param{st.NewString("a")})),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
}

func TestNestedResult(t *testing.T) {
	t.Run("should parse nested result with class head", func(t *testing.T) {
		s := `main = (p) "a"`
		expected := expectedStatements([]st.Statement{
			st.NewRule(
				stPattern.NewNamedRule("main", stPattern.ParamsToFixedParamPart([]stPattern.Param{})),
				stResult.NewObject(
					stResult.NewObject(st.NewClass("p"), stResult.ParamsToFixedParamPart([]stResult.Param{})),
					stResult.ParamsToFixedParamPart([]stResult.Param{st.NewString("a")}),
				),
			),
		})
		assert.Equal(t, expected, FilterSuccess(ParseString(s)))
	})
}

func expectedStatements(sts []st.Statement) []tuple.Of2[st.Package, []tk.Token] {
	pkg := st.NewPackage([]st.File{st.NewFile(sts)})
	return []tuple.Of2[st.Package, []tk.Token]{tuple.New2(pkg, []tk.Token{})}
}
