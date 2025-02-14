package object

import (
	"github.com/SSripilaipong/muto/common/fn"
	"github.com/SSripilaipong/muto/common/slc"
	ruleMutation "github.com/SSripilaipong/muto/core/mutation/rule"
	"github.com/SSripilaipong/muto/syntaxtree"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

var NewMutatorsFromStatements = fn.Compose(slc.Map(ruleMutation.New), mapFilterRuleFromStatement)

var mapFilterRuleFromStatement = fn.Compose(slc.Map(syntaxtree.UnsafeStatementToRule), slc.Filter(base.IsRuleStatement))
