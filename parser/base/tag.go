package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/strutil"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Tag = ps.Map(classToTag, Prefix(Dot, Class))

var TagResultNode = ps.Map(stResult.ToNode, Tag)

var TagPattern = ps.Map(stBase.ToPattern, Tag)

var classToTag = fn.Compose3(stBase.NewTag, strutil.WithPrefix("."), stBase.ClassToName)
