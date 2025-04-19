package reader

import (
	"fmt"

	"github.com/SSripilaipong/muto/builder/repl/core/command"
	"github.com/SSripilaipong/muto/common/optional"
	ruleMutation "github.com/SSripilaipong/muto/core/mutation/rule"
	mutationRuleBuilder "github.com/SSripilaipong/muto/core/mutation/rule/builder"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	"github.com/SSripilaipong/muto/parser/repl"
	replSt "github.com/SSripilaipong/muto/syntaxtree/repl"
)

func TextToCommand(text string) optional.Of[command.Command] {
	statement := repl.ParseStatement(text)
	if statement.IsErr() {
		fmt.Println(statement.Error().Error())
		return optional.Empty[command.Command]()
	}

	s := statement.Value()
	switch {
	case replSt.IsReplCommand(s):
		return newReplCommand(replSt.UnsafeStatementToReplCommand(s))
	case replSt.IsRuleStatement(s):
		return newAddRuleCommand(replSt.UnsafeStatementToRule(s))
	case replSt.IsNodeStatement(s):
		return newMutateNodeCommand(replSt.UnsafeStatementToNode(s))
	default:
		return optional.Empty[command.Command]()
	}
}

func newAddRuleCommand(s replSt.Rule) optional.Of[command.Command] {
	return optional.Value[command.Command](command.NewAddRule(ruleMutation.New(s.Rule())))
}

func newMutateNodeCommand(x replSt.Node) optional.Of[command.Command] {
	builder := mutationRuleBuilder.NewObject(x.Node().AsObject())
	node := builder.Build(parameter.New())
	if node.IsEmpty() {
		return optional.Empty[command.Command]()
	}
	return optional.Value[command.Command](command.NewMutateNode(node.Value()))
}
