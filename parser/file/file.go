package file

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree"
)

var ParseFileFromString = fn.Compose3(FilterResult, File, psBase.StringToCharTokens)

var File = ps.RsMap(syntaxtree.NewFile, psBase.IgnoreLeadingLineBreak(psBase.IgnoreTrailingLineBreak(rsStatements)))
