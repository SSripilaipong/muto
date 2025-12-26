package repl

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	st "github.com/SSripilaipong/muto/syntaxtree"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

func TestReplImportBuildNode(t *testing.T) {
	printer := &capturePrinter{}
	repl := New(noopReader{}, printer)
	repl.program.ImportBuiltin("time")

	obj := stResult.NewObject(
		st.NewImportedClass("time", "sleep"),
		stResult.ParamsToFixedParamPart([]stResult.Param{st.NewNumber("0")}),
	)

	node := repl.program.BuildNode(obj)
	if node.IsEmpty() {
		t.Fatal("expected node to build")
	}

	repl.program.MutateNode(node.Value())
	assert.Equal(t, "$", printer.output)
}

type noopReader struct{}

func (noopReader) ReadLine() (string, error) {
	return "", io.EOF
}

type noopPrinter struct{}

func (noopPrinter) Print(string) {}

type capturePrinter struct {
	output string
}

func (p *capturePrinter) Print(x string) {
	p.output = x
}
