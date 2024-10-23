package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRsSequence3(t *testing.T) {
	t.Run("should accept sequence", func(t *testing.T) {
		p := RsSequence3(char('a'), char('b'), char('c'))
		r, k := p([]rune("abc"))[0].Return()
		assert.True(t, r.IsOk())
		assert.Empty(t, k)
	})
}
