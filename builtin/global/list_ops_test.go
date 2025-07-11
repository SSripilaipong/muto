package global

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
)

func TestControl_map(t *testing.T) {
	module := NewModuleForStdio()
	class := module.GetOrCreateClass("map")
	stringClass := module.GetOrCreateClass("string")

	t.Run("should return empty list", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewOneLayerObject(class, []base.Node{
			stringClass, base.WrapWithObject(base.NewTag("my-data")),
		})) // map string ($)
		assert.Equal(t, base.WrapWithObject(base.NewUnlinkedClass("$")), result)
	})

	t.Run("should apply to each data", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewOneLayerObject(class, []base.Node{
			stringClass, base.NewOneLayerObject(base.NewTag("my-data"), []base.Node{base.NewNumberFromString("1"), base.NewNumberFromString("2")}),
		})) // map string ($ 1 2)
		assert.Equal(t, base.NewNamedOneLayerObject("$", []base.Node{base.NewString("1"), base.NewString("2")}), result)
	})
}

func TestControl_filter(t *testing.T) {
	module := NewModuleForStdio()
	class := module.GetOrCreateClass("filter")
	isStringClass := module.GetOrCreateClass("string?")

	t.Run("should return empty list", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewOneLayerObject(class, []base.Node{
			isStringClass, base.WrapWithObject(base.NewTag("my-data")),
		})) // filter string? ($)
		assert.Equal(t, base.WrapWithObject(base.NewUnlinkedClass("$")), result)
	})

	t.Run("should apply to each data", func(t *testing.T) {
		result := mutateUntilTerminated(base.NewOneLayerObject(class, []base.Node{
			isStringClass, base.NewOneLayerObject(base.NewTag("my-data"), []base.Node{
				base.NewNumberFromString("1"), base.NewString("2"), base.NewNumberFromString("3"), base.NewString("4"),
			}),
		})) // filter string? ($ 1 "2" 3 "4")
		assert.Equal(t, base.NewNamedOneLayerObject("$", []base.Node{base.NewString("2"), base.NewString("4")}), result)
	})
}
