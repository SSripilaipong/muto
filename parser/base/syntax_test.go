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
