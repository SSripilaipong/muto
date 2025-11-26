package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequence3(t *testing.T) {
	t.Run("should accept sequence", func(t *testing.T) {
		p := Sequence3(ToParser(char('a')), ToParser(char('b')), ToParser(char('c')))
		r := p([]rune("abc"))
		assert.True(t, r.IsOk())
		assert.Empty(t, r.Remaining())
	})
}
