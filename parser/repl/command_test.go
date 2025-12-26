package repl

import (
	"testing"

	"github.com/stretchr/testify/assert"

	replSt "github.com/SSripilaipong/muto/syntaxtree/repl"
)

func TestParseStatement_importCommand(t *testing.T) {
	result := ParseStatement(":import time")

	importCmd := replSt.UnsafeCommandToImportCommand(replSt.UnsafeStatementToReplCommand(result.Value()))
	assert.Equal(t, []string{"time"}, importCmd.Path())
}
