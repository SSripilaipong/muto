package file

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	psBase "github.com/SSripilaipong/muto/parser/base"
	"github.com/SSripilaipong/muto/syntaxtree/base"
)

var ParseFileFromString = fn.Compose3(FilterResult, File, psBase.StringToCharTokens)

var File = ps.RsMap(base.NewFile, psBase.IgnoreLeadingLineBreak(psBase.IgnoreTrailingLineBreak(rsStatements)))
