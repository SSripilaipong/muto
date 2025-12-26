package command

import (
	"fmt"
	"strings"

	"github.com/SSripilaipong/go-common/optional"

	"github.com/SSripilaipong/muto/core/base"
	ruleMutator "github.com/SSripilaipong/muto/core/mutation/rule/mutator"
	"github.com/SSripilaipong/muto/parser/repl"
	"github.com/SSripilaipong/muto/syntaxtree"
	replSt "github.com/SSripilaipong/muto/syntaxtree/repl"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

type Builder interface {
	BuildRule(syntaxtree.Rule) ruleMutator.NamedUnit
	BuildNode(stResult.Object) optional.Of[base.Node]
}

type Parser struct {
	builder Builder
}

func NewParser(builder Builder) Parser {
	return Parser{builder: builder}
}

func (p Parser) Parse(text string) optional.Of[Command] {
	return p.parse(strings.TrimSpace(text))
}

func (p Parser) parse(text string) optional.Of[Command] {
	statement := repl.ParseStatement(text)
	if statement.IsErr() {
		fmt.Println(statement.Error().Error())
		return optional.Empty[Command]()
	}

	s := statement.Value()
	switch {
	case replSt.IsReplCommand(s):
		return newReplCommand(replSt.UnsafeStatementToReplCommand(s))
	case replSt.IsRuleStatement(s):
		return p.newAddRuleCommand(replSt.UnsafeStatementToRule(s))
	case replSt.IsNodeStatement(s):
		return p.newMutateNodeCommand(replSt.UnsafeStatementToNode(s))
	default:
		return optional.Empty[Command]()
	}
}

func (p Parser) newAddRuleCommand(s replSt.Rule) optional.Of[Command] {
	return optional.Value[Command](NewAddRule(p.builder.BuildRule(s.Rule())))
}

func (p Parser) newMutateNodeCommand(x replSt.Node) optional.Of[Command] {
	node := p.builder.BuildNode(x.Node().AsObject())
	if node.IsEmpty() {
		return optional.Empty[Command]()
	}
	return optional.Value[Command](NewMutateNode(node.Value()))
}

func newReplCommand(st replSt.Command) optional.Of[Command] {
	switch {
	case replSt.IsQuitCommand(st):
		return optional.Value[Command](NewQuit())
	case replSt.IsImportCommand(st):
		cmd := replSt.UnsafeCommandToImportCommand(st)
		return optional.Value[Command](NewImport(cmd.JoinedPath()))
	}
	fmt.Println("unknown command:", st)
	return optional.Empty[Command]()
}
