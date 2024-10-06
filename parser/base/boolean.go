package base

import (
	ps "github.com/SSripilaipong/muto/common/parsing"
	psPred "github.com/SSripilaipong/muto/parser/predicate"
	tk "github.com/SSripilaipong/muto/parser/tokenizer"
)

var Boolean = ps.Or(
	consumeId(psPred.IsBooleanValue),
	ps.Map(tk.NewIdentifier, fixedChars("true")),
	ps.Map(tk.NewIdentifier, fixedChars("false")),
)
