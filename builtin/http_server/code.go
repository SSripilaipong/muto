package httpserver

import "github.com/SSripilaipong/muto/parser"

var rawStatements = parser.StringsToStatementsOrPanic([]string{
	startCode,
	requestCode,
})
