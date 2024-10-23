package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRsOptionalGreedyRepeat(t *testing.T) {
	t.Run("should accept empty", func(t *testing.T) {
		p := RsOptionalGreedyRepeat(RsConsumeIf(func(s rune) bool { return s == 'a' }))
		r, k := p([]rune(""))[0].Return()
		assert.True(t, r.IsOk())
		assert.Empty(t, k)
	})

	t.Run("should not consume other token", func(t *testing.T) {
		p := RsOptionalGreedyRepeat(RsConsumeIf(func(s rune) bool { return s == 'a' }))
		r, k := p([]rune("b"))[0].Return()
		assert.True(t, r.IsOk())
		assert.Equal(t, k, []rune("b"))
	})
}
