package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRsIgnoreSpaceBetween2(t *testing.T) {

	t.Run("should accept without space", func(t *testing.T) {
		p := RsIgnoreSpaceBetween2(rsChRune('a'), rsChRune('b'))
		x := p(StringToCharTokens("ab"))
		assert.Len(t, x, 1)
		r, k := x[0].Return()
		assert.True(t, r.IsOk())
		assert.Empty(t, k)
	})

	t.Run("should accept with space", func(t *testing.T) {
		p := RsIgnoreSpaceBetween2(rsChRune('a'), rsChRune('b'))
		x := p(StringToCharTokens("a b"))
		assert.Len(t, x, 1)
		r, k := x[0].Return()
		assert.True(t, r.IsOk())
		assert.Empty(t, k)
	})
}

func TestWhiteSpaceSeparated2(t *testing.T) {
	t.Run("should not accept without space", func(t *testing.T) {
		p := WhiteSpaceSeparated2(chRune('a'), chRune('b'))
		x := p(StringToCharTokens("ab"))
		assert.Len(t, x, 0)
	})

	t.Run("should accept with space", func(t *testing.T) {
		p := WhiteSpaceSeparated2(chRune('a'), chRune('b'))
		x := p(StringToCharTokens("a b"))
		assert.Len(t, x, 1)
	})

	t.Run("should accept with linebreak", func(t *testing.T) {
		p := WhiteSpaceSeparated2(chRune('a'), chRune('b'))
		x := p(StringToCharTokens("a\nb"))
		assert.Len(t, x, 1)
	})

	t.Run("should accept with space followed by linebreak", func(t *testing.T) {
		p := WhiteSpaceSeparated2(chRune('a'), chRune('b'))
		x := p(StringToCharTokens("a \nb"))
		assert.Len(t, x, 1)
	})
}

func TestOptionalGreedyRepeatWhiteSpaceSeparated(t *testing.T) {
	t.Run("should accept empty string", func(t *testing.T) {
		p := OptionalGreedyRepeatWhiteSpaceSeparated(chRune('a'))
		x := p(StringToCharTokens(""))
		assert.Len(t, x, 1)
		assert.Len(t, x[0].X2(), 0)
	})

	t.Run("should accept one pattern", func(t *testing.T) {
		p := OptionalGreedyRepeatWhiteSpaceSeparated(chRune('a'))
		x := p(StringToCharTokens("a"))
		assert.Len(t, x, 1)
		assert.Len(t, x[0].X2(), 0)
	})

	t.Run("should accept two pattern separated by space", func(t *testing.T) {
		p := OptionalGreedyRepeatWhiteSpaceSeparated(chRune('a'))
		x := p(StringToCharTokens("a a"))
		assert.Len(t, x, 1)
		assert.Len(t, x[0].X2(), 0)
	})

	t.Run("should accept two pattern separated by space and linebreak", func(t *testing.T) {
		p := OptionalGreedyRepeatWhiteSpaceSeparated(chRune('a'))
		x := p(StringToCharTokens("a \na"))
		assert.Len(t, x, 1)
		assert.Len(t, x[0].X2(), 0)
	})

	t.Run("should not accept two pattern without space/linebreak", func(t *testing.T) {
		p := OptionalGreedyRepeatWhiteSpaceSeparated(chRune('a'))
		x := p(StringToCharTokens("aa"))
		assert.Len(t, x, 1)
		assert.Equal(t, IgnoreLineAndColumn(StringToCharTokens("a")), IgnoreLineAndColumn(x[0].X2()))
	})
}
