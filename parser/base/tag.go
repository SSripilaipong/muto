package base

import (
	"github.com/SSripilaipong/muto/common/fn"
	ps "github.com/SSripilaipong/muto/common/parsing"
	"github.com/SSripilaipong/muto/common/strutil"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stPattern "github.com/SSripilaipong/muto/syntaxtree/pattern"
	stResult "github.com/SSripilaipong/muto/syntaxtree/result"
)

var Tag = ps.Map(classToTag, Prefix(Dot, Class))

var TagResultNode = ps.Map(stResult.ToNode, Tag)

var TagPatternParam = ps.Map(tagToPatternParam, Tag)

func tagToPatternParam(x st.Tag) stPattern.Param { return x }

var classToTag = fn.Compose3(st.NewTag, strutil.WithPrefix("."), st.ClassToName)
