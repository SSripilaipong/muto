package builder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SSripilaipong/muto/core/base"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestBuildBoolean(t *testing.T) {
	t.Run("should build true", func(t *testing.T) {
		fmt.Println(New(stBase.NewBoolean("true")).Build(nil))
		assert.Equal(t, base.NewBoolean(true), New(stBase.NewBoolean("true")).Build(nil).Value())
	})

	t.Run("should build false", func(t *testing.T) {
		assert.Equal(t, base.NewBoolean(false), New(stBase.NewBoolean("false")).Build(nil).Value())
	})
}

func TestBuildTag(t *testing.T) {
	t.Run("should build tag", func(t *testing.T) {
		assert.Equal(t, base.NewTag("abc"), New(stBase.NewTag(".abc")).Build(nil).Value())
	})
}

func TestNew_Structure(t *testing.T) {
	t.Run("should build structure", func(t *testing.T) {
		assert.Equal(t, base.NewStructureFromRecords(nil), New(stResult.NewStructure([]stResult.StructureRecord{})).Build(nil).Value())
	})
}
