package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
)

func TestAppendRemainingChildren(t *testing.T) {
	t.Run("should just add all params to primitive node", func(t *testing.T) {
		value := base.NewString("abc")
		b := wrapAppendRemainingChildren(newConstantWrapper(value))
		params := base.NewParamChain([][]base.Node{
			{base.NewString("def")},
			{base.NewString("xyz"), base.NewString("123")},
		})
		result, _ := b.Build(parameter.New().AppendAllRemainingParamChain(params)).Return()
		assert.Equal(t, base.NewCompoundObject(value, params), result)
	})
}
