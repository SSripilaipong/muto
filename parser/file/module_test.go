package file

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/go-common/rslt"

	psBase "github.com/SSripilaipong/muto/parser/base"
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
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("main"), stPattern.PatternsToFixedParamPart([]base.Pattern{})),
				stResult.SingleNodeToNakedObject(syntaxtree.NewString(`"hello world"`)),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse with pattern param", func(t *testing.T) {
		s := `main A = + 1 "abc"`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("main"), stPattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("A")})),
				stResult.NewNakedObject(syntaxtree.NewLocalClass("+"), stResult.FixedParamPart([]stResult.Param{syntaxtree.NewNumber("1"), syntaxtree.NewString("\"abc\"")})),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})
}

func TestParseVariable(t *testing.T) {
	t.Run("should parse rule with variables with same name", func(t *testing.T) {
		s := `f X X = 1`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToParamPart([]base.Pattern{syntaxtree.NewVariable("X"), syntaxtree.NewVariable("X")})),
				stResult.SingleNodeToNakedObject(syntaxtree.NewNumber("1")),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})
}

func TestParseConjunctionPattern(t *testing.T) {
	t.Run("should parse conjunction pattern in param", func(t *testing.T) {
		s := `f X^(g Y) = h Y X`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(
					syntaxtree.NewLocalClass("f"),
					stPattern.PatternsToFixedParamPart([]base.Pattern{
						stPattern.NewConjunction(
							syntaxtree.NewVariable("X"),
							stPattern.NewNonDeterminantObject(
								syntaxtree.NewLocalClass("g"),
								stPattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("Y")}),
							),
						),
					}),
				),
				stResult.NewNakedObject(
					syntaxtree.NewLocalClass("h"),
					stResult.FixedParamPart([]stResult.Param{syntaxtree.NewVariable("Y"), syntaxtree.NewVariable("X")}),
				),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})
}

func TestParseObjectParam(t *testing.T) {
	t.Run("should parse boolean as object head", func(t *testing.T) {
		s := `f (true 456) = 789`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]base.Pattern{
					stPattern.NewNonDeterminantObject(syntaxtree.NewBoolean("true"), stPattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewNumber("456")})),
				})),
				stResult.SingleNodeToNakedObject(syntaxtree.NewNumber("789")),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse string as anonymous head", func(t *testing.T) {
		s := `f ("a" 456) = 789`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]base.Pattern{
					stPattern.NewNonDeterminantObject(syntaxtree.NewString(`"a"`), stPattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewNumber("456")})),
				})),
				stResult.SingleNodeToNakedObject(syntaxtree.NewNumber("789")),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse number as anonymous head", func(t *testing.T) {
		s := `f (123 456) = 789`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]base.Pattern{
					stPattern.NewNonDeterminantObject(syntaxtree.NewNumber("123"), stPattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewNumber("456")})),
				})),
				stResult.SingleNodeToNakedObject(syntaxtree.NewNumber("789")),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse number as anonymous head without children", func(t *testing.T) {
		s := `f (123) = 789`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]base.Pattern{
					stPattern.NewNonDeterminantObject(syntaxtree.NewNumber("123"), stPattern.PatternsToFixedParamPart([]base.Pattern{})),
				})),
				stResult.SingleNodeToNakedObject(syntaxtree.NewNumber("789")),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})
}

func TestParseObject(t *testing.T) {
	t.Run("should parse object name in object param", func(t *testing.T) {
		s := `main = a b`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("main"), stPattern.PatternsToParamPart([]base.Pattern{})),
				stResult.NewNakedObject(syntaxtree.NewLocalClass("a"), stResult.FixedParamPart([]stResult.Param{syntaxtree.NewLocalClass("b")})),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse imported class as object head", func(t *testing.T) {
		s := `main = time.sleep 1`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("main"), stPattern.PatternsToParamPart([]base.Pattern{})),
				stResult.NewNakedObject(
					syntaxtree.NewImportedClass("time", "sleep"),
					stResult.FixedParamPart([]stResult.Param{syntaxtree.NewNumber("1")}),
				),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse nested object", func(t *testing.T) {
		s := `h ((g 1) X) = 999`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("h"), stPattern.PatternsToParamPart([]base.Pattern{
					stPattern.NewNonDeterminantObject(
						stPattern.NewNonDeterminantObject(syntaxtree.NewLocalClass("g"), stPattern.PatternsToParamPart([]base.Pattern{syntaxtree.NewNumber("1")})),
						stPattern.FixedParamPart([]base.Pattern{syntaxtree.NewVariable("X")}),
					),
				})),
				stResult.SingleNodeToNakedObject(syntaxtree.NewNumber("999")),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})
	t.Run("should parse object", func(t *testing.T) {
		s := `main A = (+ 1 2) 3 4`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("main"), stPattern.PatternsToParamPart([]base.Pattern{syntaxtree.NewVariable("A")})),
				stResult.NewNakedObject(
					stResult.NewObject(syntaxtree.NewLocalClass("+"), stResult.FixedParamPart([]stResult.Param{syntaxtree.NewNumber("1"), syntaxtree.NewNumber("2")})),
					stResult.ParamsToFixedParamPart([]stResult.Param{syntaxtree.NewNumber("3"), syntaxtree.NewNumber("4")}),
				),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse nested anonymous object", func(t *testing.T) {
		s := `main A = ((+ 1 2) 3 4) 5 6`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("main"), stPattern.PatternsToParamPart([]base.Pattern{syntaxtree.NewVariable("A")})),
				stResult.NewNakedObject(
					stResult.NewObject(stResult.NewObject(syntaxtree.NewLocalClass("+"), stResult.FixedParamPart([]stResult.Param{syntaxtree.NewNumber("1"), syntaxtree.NewNumber("2")})), stResult.ParamsToFixedParamPart([]stResult.Param{syntaxtree.NewNumber("3"), syntaxtree.NewNumber("4")})),
					stResult.ParamsToFixedParamPart([]stResult.Param{syntaxtree.NewNumber("5"), syntaxtree.NewNumber("6")}),
				),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse number as a head", func(t *testing.T) {
		s := `main = 123 456`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("main"), stPattern.PatternsToParamPart([]base.Pattern{})),
				stResult.NewNakedObject(
					syntaxtree.NewNumber("123"),
					stResult.ParamsToFixedParamPart([]stResult.Param{syntaxtree.NewNumber("456")}),
				),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})
}

func TestParseVariadicVarPattern(t *testing.T) {
	t.Run("should parse left variadic var", func(t *testing.T) {
		s := `f Xs... X = g X`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.NewLeftVariadicParamPart("Xs", []base.Pattern{syntaxtree.NewVariable("X")})),
				stResult.NewNakedObject(syntaxtree.NewLocalClass("g"), stResult.FixedParamPart([]stResult.Param{syntaxtree.NewVariable("X")})),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse right variadic var", func(t *testing.T) {
		s := `f X Xs... = g X`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.NewRightVariadicParamPart("Xs", []base.Pattern{syntaxtree.NewVariable("X")})),
				stResult.NewNakedObject(syntaxtree.NewLocalClass("g"), stResult.FixedParamPart([]stResult.Param{syntaxtree.NewVariable("X")})),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse nested variadic var", func(t *testing.T) {
		s := `f X (g Y Ys...) = g Y`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(
					syntaxtree.NewLocalClass("f"),
					stPattern.PatternsToParamPart([]base.Pattern{
						syntaxtree.NewVariable("X"),
						stPattern.NewNonDeterminantObject(syntaxtree.NewLocalClass("g"), stPattern.NewRightVariadicParamPart("Ys", []base.Pattern{syntaxtree.NewVariable("Y")})),
					}),
				),
				stResult.NewNakedObject(syntaxtree.NewLocalClass("g"), stResult.FixedParamPart([]stResult.Param{syntaxtree.NewVariable("Y")})),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})
}

func TestParseVariadicParamPart(t *testing.T) {
	t.Run("should parse left variadic param part", func(t *testing.T) {
		s := `f Xs... = g Xs... X`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.NewLeftVariadicParamPart("Xs", stPattern.PatternsToFixedParamPart([]base.Pattern{}))),
				stResult.NewNakedObject(syntaxtree.NewLocalClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{stResult.NewVariadicVariable("Xs"), syntaxtree.NewVariable("X")})),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse right variadic param part", func(t *testing.T) {
		s := `f Xs... X = g X Xs...`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.NewLeftVariadicParamPart("Xs", []base.Pattern{syntaxtree.NewVariable("X")})),
				stResult.NewNakedObject(syntaxtree.NewLocalClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{syntaxtree.NewVariable("X"), stResult.NewVariadicVariable("Xs")})),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse result with multiple variadic variables in param part", func(t *testing.T) {
		s := `f Xs... X = g Xs... X Xs...`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.NewLeftVariadicParamPart("Xs", []base.Pattern{syntaxtree.NewVariable("X")})),
				stResult.NewNakedObject(syntaxtree.NewLocalClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{stResult.NewVariadicVariable("Xs"), syntaxtree.NewVariable("X"), stResult.NewVariadicVariable("Xs")})),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})
}

func TestVariableRulePattern(t *testing.T) {
	t.Run("should parse variable rule pattern", func(t *testing.T) {
		s := `f (G X) = 1`
		expectedNestedPattern := stPattern.NewNonDeterminantObject(
			syntaxtree.NewVariable("G"),
			stPattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewVariable("X")}),
		)
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(
					syntaxtree.NewLocalClass("f"),
					stPattern.PatternsToFixedParamPart([]base.Pattern{expectedNestedPattern}),
				),
				stResult.SingleNodeToNakedObject(syntaxtree.NewNumber("1")),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})
}

func TestActiveRule(t *testing.T) {
	t.Run("should parse active rule", func(t *testing.T) {
		s := `@ f Xs... X = g X Xs...`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewActiveRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.NewLeftVariadicParamPart("Xs", []base.Pattern{syntaxtree.NewVariable("X")})),
				stResult.NewNakedObject(syntaxtree.NewLocalClass("g"), stResult.ParamsToFixedParamPart([]stResult.Param{syntaxtree.NewVariable("X"), stResult.NewVariadicVariable("Xs")})),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})
}

func TestBoolean(t *testing.T) {
	t.Run("should parse boolean as a rule result", func(t *testing.T) {
		s := `main = true`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("main"), stPattern.PatternsToFixedParamPart([]base.Pattern{})),
				stResult.SingleNodeToNakedObject(syntaxtree.NewBoolean("true")),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse boolean as a rule param", func(t *testing.T) {
		s := `main = f true`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("main"), stPattern.PatternsToFixedParamPart([]base.Pattern{})),
				stResult.NewNakedObject(syntaxtree.NewLocalClass("f"), stResult.ParamsToFixedParamPart([]stResult.Param{syntaxtree.NewBoolean("true")})),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse boolean as a rule pattern param", func(t *testing.T) {
		s := `f true = "a"`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]base.Pattern{syntaxtree.NewBoolean("true")})),
				stResult.SingleNodeToNakedObject(syntaxtree.NewString("\"a\"")),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})

	t.Run("should parse boolean as an object head", func(t *testing.T) {
		s := `main = true "a"`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("main"), stPattern.PatternsToFixedParamPart([]base.Pattern{})),
				stResult.NewNakedObject(syntaxtree.NewBoolean("true"), stResult.ParamsToFixedParamPart([]stResult.Param{syntaxtree.NewString("\"a\"")})),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})
}

func TestNestedResult(t *testing.T) {
	t.Run("should parse nested result with class head", func(t *testing.T) {
		s := `main = (p) "a"`
		expected := expectedStatements([]base.Statement{
			syntaxtree.NewRule(
				stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("main"), stPattern.PatternsToFixedParamPart([]base.Pattern{})),
				stResult.NewNakedObject(
					stResult.NewObject(syntaxtree.NewLocalClass("p"), stResult.ParamsToFixedParamPart([]stResult.Param{})),
					stResult.ParamsToFixedParamPart([]stResult.Param{syntaxtree.NewString("\"a\"")}),
				),
			),
		})
		assert.Equal(t, expected, psBase.FilterResult(ParseModuleCombinationFromString(s)))
	})
}

func expectedStatements(sts []base.Statement) rslt.Of[syntaxtree.Module] {
	pkg := syntaxtree.NewModule([]syntaxtree.File{syntaxtree.NewFile(sts)})
	return rslt.Value(pkg)
}
