package file

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/rslt"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/common/tuple"
	tk "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
	"github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestParseString(t *testing.T) {
	t.Run("should parse hello world", func(t *testing.T) {
		s := `main = "hello world"`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("main"), stPattern.PatternsToFixedParamPart([]base.Pattern{})),
				base.NewString(`"hello world"`),
			),
		})
		assert.Equal(t, expected, parsing.FilterResult(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse with pattern param", func(t *testing.T) {
		s := `main A = + 1 "abc"`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("main"), stPattern.PatternsToFixedParamPart([]base.Pattern{base.NewVariable("A")})),
				stResult.NewObject(base.NewClass("+"), stResult.FixedParamPart([]stResult.Param{base.NewNumber("1"), base.NewString("\"abc\"")})),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})
}

func TestParseVariable(t *testing.T) {
	t.Run("should parse rule with variables with same name", func(t *testing.T) {
		s := `f X X = 1`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("f"), stPattern.PatternsToParamPart([]base.Pattern{base.NewVariable("X"), base.NewVariable("X")})),
				base.NewNumber("1"),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})
}

func TestParseObjectParam(t *testing.T) {
	t.Run("should parse boolean as object head", func(t *testing.T) {
		s := `f (true 456) = 789`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("f"), stPattern.PatternsToFixedParamPart([]base.Pattern{
					stPattern.NewNonDeterminantObject(base.NewBoolean("true"), stPattern.PatternsToFixedParamPart([]base.Pattern{base.NewNumber("456")})),
				})),
				base.NewNumber("789"),
			),
		})
		assert.Equal(t, expected, parsing.FilterResult(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse string as anonymous head", func(t *testing.T) {
		s := `f ("a" 456) = 789`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("f"), stPattern.PatternsToFixedParamPart([]base.Pattern{
					stPattern.NewNonDeterminantObject(base.NewString(`"a"`), stPattern.PatternsToFixedParamPart([]base.Pattern{base.NewNumber("456")})),
				})),
				base.NewNumber("789"),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse number as anonymous head", func(t *testing.T) {
		s := `f (123 456) = 789`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("f"), stPattern.PatternsToFixedParamPart([]base.Pattern{
					stPattern.NewNonDeterminantObject(base.NewNumber("123"), stPattern.PatternsToFixedParamPart([]base.Pattern{base.NewNumber("456")})),
				})),
				base.NewNumber("789"),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse number as anonymous head without children", func(t *testing.T) {
		s := `f (123) = 789`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("f"), stPattern.PatternsToFixedParamPart([]base.Pattern{
					stPattern.NewNonDeterminantObject(base.NewNumber("123"), stPattern.PatternsToFixedParamPart([]base.Pattern{})),
				})),
				base.NewNumber("789"),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})
}

func TestParseObject(t *testing.T) {
	t.Run("should parse object name in object param", func(t *testing.T) {
		s := `main = a b`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("main"), stPattern.PatternsToParamPart([]base.Pattern{})),
				stResult.NewObject(base.NewClass("a"), stResult.FixedParamPart([]stResult.Param{base.NewClass("b")})),
			),
		})
		assert.Equal(t, expected, parsing.FilterResult(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse nested object", func(t *testing.T) {
		s := `h ((g 1) X) = 999`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("h"), stPattern.PatternsToParamPart([]base.Pattern{
					stPattern.NewNonDeterminantObject(
						stPattern.NewNonDeterminantObject(base.NewClass("g"), stPattern.PatternsToParamPart([]base.Pattern{base.NewNumber("1")})),
						stPattern.FixedParamPart([]base.Pattern{base.NewVariable("X")}),
					),
				})),
				base.NewNumber("999"),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})
	t.Run("should parse object", func(t *testing.T) {
		s := `main A = (+ 1 2) 3 4`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("main"), stPattern.PatternsToParamPart([]base.Pattern{base.NewVariable("A")})),
				stResult.NewObject(
					stResult.NewObject(base.NewClass("+"), stResult.FixedParamPart([]stResult.Param{base.NewNumber("1"), base.NewNumber("2")})),
					stResult.ParamsToFixedParamPart([]stResult.Param{base.NewNumber("3"), base.NewNumber("4")}),
				),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse nested anonymous object", func(t *testing.T) {
		s := `main A = ((+ 1 2) 3 4) 5 6`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("main"), stPattern.PatternsToParamPart([]base.Pattern{base.NewVariable("A")})),
				stResult.NewObject(
					stResult.NewObject(stResult.NewObject(base.NewClass("+"), stResult.FixedParamPart([]stResult.Param{base.NewNumber("1"), base.NewNumber("2")})), stResult.ParamsToFixedParamPart([]stResult.Param{base.NewNumber("3"), base.NewNumber("4")})),
					stResult.ParamsToFixedParamPart([]stResult.Param{base.NewNumber("5"), base.NewNumber("6")}),
				),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse number as a head", func(t *testing.T) {
		s := `main = 123 456`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("main"), stPattern.PatternsToParamPart([]base.Pattern{})),
				stResult.NewObject(
					base.NewNumber("123"),
					stResult.ParamsToFixedParamPart([]stResult.Param{base.NewNumber("456")}),
				),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})
}

func TestParseVariadicVarPattern(t *testing.T) {
	t.Run("should parse left variadic var", func(t *testing.T) {
		s := `f Xs... X = g X`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("f"), stPattern.NewLeftVariadicParamPart("Xs", []base.Pattern{base.NewVariable("X")})),
				stResult.NewObject(base.NewClass("g"), stResult.FixedParamPart([]stResult.Param{base.NewVariable("X")})),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse right variadic var", func(t *testing.T) {
		s := `f X Xs... = g X`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("f"), stPattern.NewRightVariadicParamPart("Xs", []base.Pattern{base.NewVariable("X")})),
				stResult.NewObject(base.NewClass("g"), stResult.FixedParamPart([]stResult.Param{base.NewVariable("X")})),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse nested variadic var", func(t *testing.T) {
		s := `f X (g Y Ys...) = g Y`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(
					base.NewClass("f"),
					stPattern.PatternsToParamPart([]base.Pattern{
						base.NewVariable("X"),
						stPattern.NewNonDeterminantObject(base.NewClass("g"), stPattern.NewRightVariadicParamPart("Ys", []base.Pattern{base.NewVariable("Y")})),
					}),
				),
				stResult.NewObject(base.NewClass("g"), stResult.FixedParamPart([]stResult.Param{base.NewVariable("Y")})),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})
}

func TestParseVariadicParamPart(t *testing.T) {
	t.Run("should parse left variadic param part", func(t *testing.T) {
		s := `f Xs... = g Xs... X`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("f"), stPattern.NewLeftVariadicParamPart("Xs", stPattern.PatternsToFixedParamPart([]base.Pattern{}))),
				stResult.NewObject(base.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{stResult.NewVariadicVariable("Xs"), base.NewVariable("X")})),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse right variadic param part", func(t *testing.T) {
		s := `f Xs... X = g X Xs...`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("f"), stPattern.NewLeftVariadicParamPart("Xs", []base.Pattern{base.NewVariable("X")})),
				stResult.NewObject(base.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{base.NewVariable("X"), stResult.NewVariadicVariable("Xs")})),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse result with multiple variadic variables in param part", func(t *testing.T) {
		s := `f Xs... X = g Xs... X Xs...`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("f"), stPattern.NewLeftVariadicParamPart("Xs", []base.Pattern{base.NewVariable("X")})),
				stResult.NewObject(base.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{stResult.NewVariadicVariable("Xs"), base.NewVariable("X"), stResult.NewVariadicVariable("Xs")})),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})
}

func TestVariableRulePattern(t *testing.T) {
	t.Run("should parse variable rule pattern", func(t *testing.T) {
		s := `f (G X) = 1`
		expectedNestedPattern := stPattern.NewNonDeterminantObject(
			base.NewVariable("G"),
			stPattern.PatternsToFixedParamPart([]base.Pattern{base.NewVariable("X")}),
		)
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(
					base.NewClass("f"),
					stPattern.PatternsToFixedParamPart([]base.Pattern{expectedNestedPattern}),
				),
				base.NewNumber("1"),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})
}

func TestActiveRule(t *testing.T) {
	t.Run("should parse active rule", func(t *testing.T) {
		s := `@ f Xs... X = g X Xs...`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewActiveRule(
				stPattern.NewDeterminantObject(base.NewClass("f"), stPattern.NewLeftVariadicParamPart("Xs", []base.Pattern{base.NewVariable("X")})),
				stResult.NewObject(base.NewClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{base.NewVariable("X"), stResult.NewVariadicVariable("Xs")})),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})
}

func TestBoolean(t *testing.T) {
	t.Run("should parse boolean as a rule result", func(t *testing.T) {
		s := `main = true`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("main"), stPattern.PatternsToFixedParamPart([]base.Pattern{})),
				base.NewBoolean("true"),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse boolean as a rule param", func(t *testing.T) {
		s := `main = f true`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("main"), stPattern.PatternsToFixedParamPart([]base.Pattern{})),
				stResult.NewObject(base.NewClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{base.NewBoolean("true")})),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse boolean as a rule pattern param", func(t *testing.T) {
		s := `f true = "a"`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("f"), stPattern.PatternsToFixedParamPart([]base.Pattern{base.NewBoolean("true")})),
				base.NewString("\"a\""),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})

	t.Run("should parse boolean as an object head", func(t *testing.T) {
		s := `main = true "a"`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("main"), stPattern.PatternsToFixedParamPart([]base.Pattern{})),
				stResult.NewObject(base.NewBoolean("true"), stResult.ParamsToFixedParamPart([]stResult.Param{base.NewString("\"a\"")})),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})
}

func TestNestedResult(t *testing.T) {
	t.Run("should parse nested result with class head", func(t *testing.T) {
		s := `main = (p) "a"`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(base.NewClass("main"), stPattern.PatternsToFixedParamPart([]base.Pattern{})),
				stResult.NewObject(
					stResult.NewObject(base.NewClass("p"), stResult.ParamsToFixedParamPart([]stResult.Param{})),
					stResult.ParamsToFixedParamPart([]stResult.Param{base.NewString("\"a\"")}),
				),
			),
		})
		assert.Equal(t, expected, parsing.FilterSuccess(ParsePackageCombinationFromString(s)))
	})
}

func expectedStatements(sts []base.Statement) []tuple.Of2[rslt.Of[base.Package], []tk.Character] {
	pkg := base.NewPackage([]base.File{base.NewFile(sts)})
	return slc.Pure(tuple.New2(rslt.Value(pkg), []tk.Character{}))
}
