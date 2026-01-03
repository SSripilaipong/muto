package extractor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
)

func TestNew_ObjectWithValueHead(t *testing.T) {
	t.Run("should match boolean as a nested object head", func(t *testing.T) {
		// pattern: f (true 456)
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewNonDeterminantObject(syntaxtree.NewBoolean("true"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewNumber("456")})),
		}))
		// obj: f (true 456)
		obj := base.NewNamedOneLayerObject("f", base.NewOneLayerObject(base.NewBoolean(true), base.NewNumberFromString("456")))

		assert.True(t, NewNamedRule(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should match string as a nested object head", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewNonDeterminantObject(syntaxtree.NewString(`"a"`), stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewNumber("456")})),
		}))
		obj := base.NewNamedOneLayerObject("f", base.NewOneLayerObject(base.NewString(`a`), base.NewNumberFromString("456")))

		assert.True(t, NewNamedRule(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should match number as a nested object head", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewNonDeterminantObject(syntaxtree.NewNumber("123"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewNumber("456")})),
		}))
		obj := base.NewNamedOneLayerObject("f", base.NewOneLayerObject(base.NewNumberFromString("123"), base.NewNumberFromString("456")))

		assert.True(t, NewNamedRule(pattern).Extract(obj).IsNotEmpty())
	})
}

func TestTag(t *testing.T) {
	t.Run("should match a tag child", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewTag("abc")}))
		obj := base.NewNamedOneLayerObject("f", base.NewTag("abc"))

		assert.True(t, NewNamedRule(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should not match a non-tag child", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewTag("abc")}))
		obj := base.NewNamedOneLayerObject("f", base.NewUnlinkedRuleBasedClass("abc"))

		assert.False(t, NewNamedRule(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should match a tag in a nested head", func(t *testing.T) {
		// pattern: f (.abc 1 2)
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewNonDeterminantObject(syntaxtree.NewTag("abc"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
				syntaxtree.NewNumber("1"), syntaxtree.NewNumber("2"),
			})),
		}))
		// obj: f (.abc 1 2)
		obj := base.NewNamedOneLayerObject("f", base.NewOneLayerObject(base.NewTag("abc"), base.NewNumberFromString("1"), base.NewNumberFromString("2")))

		assert.True(t, NewNamedRule(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should match a tag in a nested param", func(t *testing.T) {
		// pattern: f (g 1 .abc 2)
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("g"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
				syntaxtree.NewNumber("1"), syntaxtree.NewTag("abc"), syntaxtree.NewNumber("2"),
			})),
		}))
		// obj: f (g 1 .abc 2)
		obj := base.NewNamedOneLayerObject("f", base.NewOneLayerObject(base.NewUnlinkedRuleBasedClass("g"), base.NewNumberFromString("1"), base.NewTag("abc"), base.NewNumberFromString("2")))

		assert.True(t, NewNamedRule(pattern).Extract(obj).IsNotEmpty())
	})
}

func TestNestedObject(t *testing.T) {
	t.Run("should not match leaf object with simple node", func(t *testing.T) {
		// pattern: f (g)
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("g"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{})),
		}))
		// obj: f g
		obj := base.NewNamedOneLayerObject("f", base.NewUnlinkedRuleBasedClass("g"))

		assert.True(t, NewNamedRule(pattern).Extract(obj).IsEmpty())
	})
}

func TestVariadicParam(t *testing.T) {
	t.Run("should match nested variadic param with size 0", func(t *testing.T) {
		// pattern: g (f Xs...)
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("g"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"),
				stPattern.NewLeftVariadicParamPart("Xs", stPattern.FixedParamPart{})),
		}))
		// obj: g (f)
		obj := base.NewNamedOneLayerObject("g", base.NewNamedOneLayerObject("f", nil))

		assert.True(t, NewNamedRule(pattern).Extract(obj).IsNotEmpty())
	})

	t.Run("should ignore underscore variadic variable", func(t *testing.T) {
		// pattern: g (f _Bc...)
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("g"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"),
				stPattern.NewLeftVariadicParamPart("_Bc", stPattern.FixedParamPart{})),
		}))
		// obj: g (f x y)
		obj := base.NewNamedOneLayerObject("g",
			base.NewNamedOneLayerObject("f", base.NewUnlinkedRuleBasedClass("x"), base.NewUnlinkedRuleBasedClass("y")),
		)

		p := NewNamedRule(pattern).Extract(obj)
		assert.True(t, p.IsNotEmpty() && p.Value().VariadicVarValue("_Bc").IsEmpty())
	})
}

func TestConjunctionPattern(t *testing.T) {
	t.Run("should match conjunction on same param", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewConjunction(
				syntaxtree.NewVariable("X"),
				stPattern.NewNonDeterminantObject(
					syntaxtree.NewLocalClass("g"),
					stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("Y")}),
				),
			),
		}))
		obj := base.NewNamedOneLayerObject("f", base.NewNamedOneLayerObject("g", base.NewNumberFromString("1")))

		params := NewNamedRule(pattern).Extract(obj)
		if assert.True(t, params.IsNotEmpty()) {
			p := params.Value()
			assert.Equal(t, base.NewNamedOneLayerObject("g", base.NewNumberFromString("1")), p.VariableValue("X").Value())
			assert.Equal(t, base.NewNumberFromString("1"), p.VariableValue("Y").Value())
		}
	})

	t.Run("should not match when conjunction fails", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewConjunction(
				syntaxtree.NewVariable("X"),
				stPattern.NewNonDeterminantObject(
					syntaxtree.NewLocalClass("g"),
					stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("Y")}),
				),
			),
		}))
		obj := base.NewNamedOneLayerObject("f", base.NewNumberFromString("123"))

		assert.True(t, NewNamedRule(pattern).Extract(obj).IsEmpty())
	})

	t.Run("should match object-head conjunction without params", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewConjunction(
				syntaxtree.NewVariable("X"),
				stPattern.NewNonDeterminantObject(
					syntaxtree.NewVariable("Y"),
					stPattern.PatternsToFixedParamPart([]stBase.Pattern{}),
				),
			),
		}))
		obj := base.NewNamedOneLayerObject("f", base.NewNamedOneLayerObject("g"))

		params := NewNamedRule(pattern).Extract(obj)
		if assert.True(t, params.IsNotEmpty()) {
			p := params.Value()
			assert.Equal(t, base.NewNamedOneLayerObject("g"), p.VariableValue("X").Value())
			assert.Equal(t, base.NewUnlinkedRuleBasedClass("g"), p.VariableValue("Y").Value())
		}
	})

	t.Run("should not match object-head conjunction against class node", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewConjunction(
				syntaxtree.NewVariable("X"),
				stPattern.NewNonDeterminantObject(
					syntaxtree.NewVariable("Y"),
					stPattern.PatternsToFixedParamPart([]stBase.Pattern{}),
				),
			),
		}))
		obj := base.NewNamedOneLayerObject("f", base.NewUnlinkedRuleBasedClass("g"))

		assert.True(t, NewNamedRule(pattern).Extract(obj).IsEmpty())
	})

	t.Run("should require all conjunctions to match", func(t *testing.T) {
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewConjunction(
				stPattern.NewConjunction(
					syntaxtree.NewVariable("X"),
					stPattern.NewNonDeterminantObject(
						syntaxtree.NewLocalClass("g"),
						stPattern.PatternsToFixedParamPart([]stBase.Pattern{}),
					),
				),
				stPattern.NewNonDeterminantObject(
					syntaxtree.NewVariable("Y"),
					stPattern.PatternsToFixedParamPart([]stBase.Pattern{}),
				),
			),
		}))
		obj := base.NewNamedOneLayerObject("f", base.NewNamedOneLayerObject("g"))

		params := NewNamedRule(pattern).Extract(obj)
		if assert.True(t, params.IsNotEmpty()) {
			p := params.Value()
			assert.Equal(t, base.NewNamedOneLayerObject("g"), p.VariableValue("X").Value())
			assert.Equal(t, base.NewUnlinkedRuleBasedClass("g"), p.VariableValue("Y").Value())
		}
	})
}

func TestVariableParam(t *testing.T) {
	t.Run("should ignore underscore variable", func(t *testing.T) {
		// pattern: g (f _Bc)
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("g"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"),
				stPattern.NewLeftVariadicParamPart("_Bc", stPattern.FixedParamPart{})),
		}))
		// obj: g (f x)
		obj := base.NewNamedOneLayerObject("g",
			base.NewNamedOneLayerObject("f", base.NewUnlinkedRuleBasedClass("x"), base.NewUnlinkedRuleBasedClass("y")),
		)

		p := NewNamedRule(pattern).Extract(obj)
		assert.True(t, p.IsNotEmpty() && p.Value().VariableValue("_Bc").IsEmpty())
	})
}

func TestRemainingChildren(t *testing.T) {
	t.Run("should extract remaining children", func(t *testing.T) {
		// pattern: f 1
		pattern := stPattern.NewDeterminantObject(syntaxtree.NewLocalClass("f"), stPattern.PatternsToFixedParamPart([]stBase.Pattern{
			syntaxtree.NewNumber("1"),
		}))
		// obj: f 1 2 3
		obj := base.NewNamedOneLayerObject("f", base.NewNumberFromString("1"), base.NewNumberFromString("2"), base.NewNumberFromString("3"))

		assert.Equal(t, []base.Node{base.NewNumberFromString("2"), base.NewNumberFromString("3")}, NewNamedRule(pattern).Extract(obj).Value().RemainingParamChain().All()[0])
	})
}

func TestDeterminantConjunction(t *testing.T) {
	t.Run("should extract determinant conjunction binding", func(t *testing.T) {
		// Pattern: (f S)^P Y
		// Object: (f 1) 2
		// Expected: S=1, P=(f 1), Y=2
		pattern := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantConjunction(
				stPattern.NewDeterminantObject(
					syntaxtree.NewLocalClass("f"),
					stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("S")}),
				),
				[]stBase.Pattern{syntaxtree.NewVariable("P")},
			),
			stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("Y")}),
		)
		// obj: (f 1) 2 - represented as f with params [[1], [2]]
		obj := base.NewCompoundObject(
			base.NewUnlinkedRuleBasedClass("f"),
			base.NewParamChain([][]base.Node{
				{base.NewNumberFromString("1")},
				{base.NewNumberFromString("2")},
			}),
		)

		params := NewNamedRule(pattern).Extract(obj)
		if assert.True(t, params.IsNotEmpty()) {
			p := params.Value()
			assert.Equal(t, base.NewNumberFromString("1"), p.VariableValue("S").Value())
			assert.Equal(t, base.NewNamedOneLayerObject("f", base.NewNumberFromString("1")), p.VariableValue("P").Value())
			assert.Equal(t, base.NewNumberFromString("2"), p.VariableValue("Y").Value())
		}
	})

	t.Run("should extract multiple conjunctions at same level", func(t *testing.T) {
		// Pattern: (f S)^P^Q Y
		// Object: (f 1) 2
		// Expected: S=1, P=(f 1), Q=(f 1), Y=2
		pattern := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantConjunction(
				stPattern.NewDeterminantObject(
					syntaxtree.NewLocalClass("f"),
					stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("S")}),
				),
				[]stBase.Pattern{syntaxtree.NewVariable("P"), syntaxtree.NewVariable("Q")},
			),
			stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("Y")}),
		)
		obj := base.NewCompoundObject(
			base.NewUnlinkedRuleBasedClass("f"),
			base.NewParamChain([][]base.Node{
				{base.NewNumberFromString("1")},
				{base.NewNumberFromString("2")},
			}),
		)

		params := NewNamedRule(pattern).Extract(obj)
		if assert.True(t, params.IsNotEmpty()) {
			p := params.Value()
			assert.Equal(t, base.NewNumberFromString("1"), p.VariableValue("S").Value())
			assert.Equal(t, base.NewNamedOneLayerObject("f", base.NewNumberFromString("1")), p.VariableValue("P").Value())
			assert.Equal(t, base.NewNamedOneLayerObject("f", base.NewNumberFromString("1")), p.VariableValue("Q").Value())
			assert.Equal(t, base.NewNumberFromString("2"), p.VariableValue("Y").Value())
		}
	})

	t.Run("should fail when conjunction pattern does not match", func(t *testing.T) {
		// Pattern: (f S)^(g X) Y
		// Object: (f 1) 2 - head is (f 1), not (g ...)
		// Expected: no match
		pattern := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantConjunction(
				stPattern.NewDeterminantObject(
					syntaxtree.NewLocalClass("f"),
					stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("S")}),
				),
				[]stBase.Pattern{
					stPattern.NewNonDeterminantObject(
						syntaxtree.NewLocalClass("g"),
						stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("X")}),
					),
				},
			),
			stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("Y")}),
		)
		obj := base.NewCompoundObject(
			base.NewUnlinkedRuleBasedClass("f"),
			base.NewParamChain([][]base.Node{
				{base.NewNumberFromString("1")},
				{base.NewNumberFromString("2")},
			}),
		)

		assert.True(t, NewNamedRule(pattern).Extract(obj).IsEmpty())
	})

	t.Run("should extract bare class with conjunction", func(t *testing.T) {
		// Pattern: f^P X
		// Object: f 1
		// Expected: P=f, X=1
		pattern := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantConjunction(
				syntaxtree.NewLocalClass("f"),
				[]stBase.Pattern{syntaxtree.NewVariable("P")},
			),
			stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("X")}),
		)
		obj := base.NewNamedOneLayerObject("f", base.NewNumberFromString("1"))

		params := NewNamedRule(pattern).Extract(obj)
		if assert.True(t, params.IsNotEmpty()) {
			p := params.Value()
			assert.Equal(t, base.NewUnlinkedRuleBasedClass("f"), p.VariableValue("P").Value())
			assert.Equal(t, base.NewNumberFromString("1"), p.VariableValue("X").Value())
		}
	})

	t.Run("should extract nested conjunctions at different levels", func(t *testing.T) {
		// Pattern: ((f S)^P Y)^Q Z
		// Object: ((f 1) 2) 3
		// Expected: S=1, P=(f 1), Y=2, Q=((f 1) 2), Z=3
		innerConj := stPattern.NewDeterminantConjunction(
			stPattern.NewDeterminantObject(
				syntaxtree.NewLocalClass("f"),
				stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("S")}),
			),
			[]stBase.Pattern{syntaxtree.NewVariable("P")},
		)
		middleObj := stPattern.NewDeterminantObject(
			innerConj,
			stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("Y")}),
		)
		outerConj := stPattern.NewDeterminantConjunction(
			middleObj,
			[]stBase.Pattern{syntaxtree.NewVariable("Q")},
		)
		pattern := stPattern.NewDeterminantObject(
			outerConj,
			stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("Z")}),
		)

		// obj: ((f 1) 2) 3 - represented as f with params [[1], [2], [3]]
		obj := base.NewCompoundObject(
			base.NewUnlinkedRuleBasedClass("f"),
			base.NewParamChain([][]base.Node{
				{base.NewNumberFromString("1")},
				{base.NewNumberFromString("2")},
				{base.NewNumberFromString("3")},
			}),
		)

		params := NewNamedRule(pattern).Extract(obj)
		if assert.True(t, params.IsNotEmpty()) {
			p := params.Value()
			assert.Equal(t, base.NewNumberFromString("1"), p.VariableValue("S").Value())
			assert.Equal(t, base.NewNamedOneLayerObject("f", base.NewNumberFromString("1")), p.VariableValue("P").Value())
			assert.Equal(t, base.NewNumberFromString("2"), p.VariableValue("Y").Value())
			// Q should be ((f 1) 2) = f with params [[1], [2]]
			expectedQ := base.NewCompoundObject(
				base.NewUnlinkedRuleBasedClass("f"),
				base.NewParamChain([][]base.Node{
					{base.NewNumberFromString("1")},
					{base.NewNumberFromString("2")},
				}),
			)
			assert.Equal(t, expectedQ, p.VariableValue("Q").Value())
			assert.Equal(t, base.NewNumberFromString("3"), p.VariableValue("Z").Value())
		}
	})

	t.Run("should extract conjunction with variable head pattern", func(t *testing.T) {
		// Pattern: (f S)^(G X) Y
		// Object: (f 1) 2
		// Expected: S=1, G=f, X=1, Y=2
		pattern := stPattern.NewDeterminantObject(
			stPattern.NewDeterminantConjunction(
				stPattern.NewDeterminantObject(
					syntaxtree.NewLocalClass("f"),
					stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("S")}),
				),
				[]stBase.Pattern{
					stPattern.NewNonDeterminantObject(
						syntaxtree.NewVariable("G"),
						stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("X")}),
					),
				},
			),
			stPattern.PatternsToFixedParamPart([]stBase.Pattern{syntaxtree.NewVariable("Y")}),
		)
		obj := base.NewCompoundObject(
			base.NewUnlinkedRuleBasedClass("f"),
			base.NewParamChain([][]base.Node{
				{base.NewNumberFromString("1")},
				{base.NewNumberFromString("2")},
			}),
		)

		params := NewNamedRule(pattern).Extract(obj)
		if assert.True(t, params.IsNotEmpty()) {
			p := params.Value()
			assert.Equal(t, base.NewNumberFromString("1"), p.VariableValue("S").Value())
			assert.Equal(t, base.NewUnlinkedRuleBasedClass("f"), p.VariableValue("G").Value())
			assert.Equal(t, base.NewNumberFromString("1"), p.VariableValue("X").Value())
			assert.Equal(t, base.NewNumberFromString("2"), p.VariableValue("Y").Value())
		}
	})
}
