package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	digit := ConsumeIf(func(s rune) bool { return s == '1' || s == '2' })
	two := ConsumeIf(func(s rune) bool { return s == '2' })
	p := First(ToParser(digit), ToParser(two))
	r, k := p([]rune("2")).Return()
	assert.True(t, r.IsOk())
	assert.Empty(t, k)
}
