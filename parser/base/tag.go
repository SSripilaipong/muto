package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/strutil"
	"github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Tag = ps.Map(classToTag, ps.Prefix(ps.ToParser(Dot), ps.ToParser(Class)))

var TagResultNode = ps.Map(stResult.ToNode, Tag).Legacy

var TagPattern = ps.Map(stBase.ToPattern, Tag).Legacy

var classToTag = fn.Compose3(syntaxtree.NewTag, strutil.WithPrefix("."), syntaxtree.LocalClassToName)
