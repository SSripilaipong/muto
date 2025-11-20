package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequence3(t *testing.T) {
	t.Run("should accept sequence", func(t *testing.T) {
		p := Sequence3(char('a'), char('b'), char('c'))
		r := p([]rune("abc"))
		assert.True(t, IsResultOk(r))
		assert.Empty(t, r.X2())
	})
}
