package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRsFirst(t *testing.T) {
	digit := RsConsumeIf(func(s rune) bool { return s == '1' || s == '2' })
	two := RsConsumeIf(func(s rune) bool { return s == '2' })
	p := RsFirst(digit, two)
	x := p([]rune("2"))
	assert.Len(t, x, 1)

	r, k := x[0].Return()
	assert.True(t, r.IsOk())
	assert.Empty(t, k)
}
