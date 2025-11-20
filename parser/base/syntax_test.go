package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ps "github.com/SSripilaipong/muto/common/parsing"
)

func TestIgnoreSpaceBetween2(t *testing.T) {
	t.Run("should accept without space", func(t *testing.T) {
		p := IgnoreSpaceBetween2(chRune('a'), chRune('b'))
		x := p(StringToCharTokens("ab"))
		assert.True(t, ps.IsResultOk(x))
	})

	t.Run("should accept with space", func(t *testing.T) {
		p := IgnoreSpaceBetween2(chRune('a'), chRune('b'))
		x := p(StringToCharTokens("a b"))
		assert.True(t, ps.IsResultOk(x))
	})
}

func TestWhiteSpaceSeparated2(t *testing.T) {
	t.Run("should not accept without space", func(t *testing.T) {
		p := WhiteSpaceSeparated2(chRune('a'), chRune('b'))
		x := p(StringToCharTokens("ab"))
		assert.True(t, ps.IsResultErr(x))
	})

	t.Run("should accept with space", func(t *testing.T) {
		p := WhiteSpaceSeparated2(chRune('a'), chRune('b'))
		x := p(StringToCharTokens("a b"))
		assert.True(t, ps.IsResultOk(x))
		assert.Len(t, x.X2(), 0)
	})

	t.Run("should accept with linebreak", func(t *testing.T) {
		p := WhiteSpaceSeparated2(chRune('a'), chRune('b'))
		x := p(StringToCharTokens("a\nb"))
		assert.True(t, ps.IsResultOk(x))
		assert.Len(t, x.X2(), 0)
	})

	t.Run("should accept with space followed by linebreak", func(t *testing.T) {
		p := WhiteSpaceSeparated2(chRune('a'), chRune('b'))
		x := p(StringToCharTokens("a \nb"))
		assert.True(t, ps.IsResultOk(x))
		assert.Len(t, x.X2(), 0)
	})
}

func TestOptionalGreedyRepeatWhiteSpaceSeparated(t *testing.T) {
	t.Run("should accept empty string", func(t *testing.T) {
		p := OptionalGreedyRepeatWhiteSpaceSeparated(chRune('a'))
		x := p(StringToCharTokens(""))
		assert.True(t, ps.IsResultOk(x))
		assert.Len(t, ps.ResultValue(x), 0)
		assert.Len(t, x.X2(), 0)
	})

	t.Run("should accept one pattern", func(t *testing.T) {
		p := OptionalGreedyRepeatWhiteSpaceSeparated(chRune('a'))
		x := p(StringToCharTokens("a"))
		assert.True(t, ps.IsResultOk(x))
		assert.Len(t, ps.ResultValue(x), 1)
		assert.Len(t, x.X2(), 0)
	})

	t.Run("should accept two pattern separated by space", func(t *testing.T) {
		p := OptionalGreedyRepeatWhiteSpaceSeparated(chRune('a'))
		x := p(StringToCharTokens("a a"))
		assert.True(t, ps.IsResultOk(x))
		assert.Len(t, ps.ResultValue(x), 2)
		assert.Len(t, x.X2(), 0)
	})

	t.Run("should accept two pattern separated by space and linebreak", func(t *testing.T) {
		p := OptionalGreedyRepeatWhiteSpaceSeparated(chRune('a'))
		x := p(StringToCharTokens("a \na"))
		assert.True(t, ps.IsResultOk(x))
		assert.Len(t, ps.ResultValue(x), 2)
		assert.Len(t, x.X2(), 0)
	})

	t.Run("should not accept two pattern without space/linebreak", func(t *testing.T) {
		p := OptionalGreedyRepeatWhiteSpaceSeparated(chRune('a'))
		x := p(StringToCharTokens("aa"))
		assert.True(t, ps.IsResultOk(x))
		assert.Equal(t, IgnoreLineAndColumn(StringToCharTokens("a")), IgnoreLineAndColumn(x.X2()))
	})
}
