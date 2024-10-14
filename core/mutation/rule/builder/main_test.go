package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestBuildBoolean(t *testing.T) {
	t.Run("should build true", func(t *testing.T) {
		assert.Equal(t, base.NewBoolean(true), New(st.NewBoolean("true"))(nil).Value())
	})

	t.Run("should build false", func(t *testing.T) {
		assert.Equal(t, base.NewBoolean(false), New(st.NewBoolean("false"))(nil).Value())
	})
}

func TestBuildTag(t *testing.T) {
	t.Run("should build tag", func(t *testing.T) {
		assert.Equal(t, base.NewTag("abc"), New(st.NewTag(".abc"))(nil).Value())
	})
}

func TestNew_Structure(t *testing.T) {
	t.Run("should build structure", func(t *testing.T) {
		assert.Equal(t, base.NewStructureFromRecords(nil), New(stResult.NewStructure([]stResult.StructureRecord{}))(nil).Value())
	})
}
