package global

import (
	"github.com/SSripilaipong/muto/parser"
)

var rawStatements = parser.StringsToStatementsOrPanic([]string{
	doCode,
	matchCode,
	retCode,
	okCode,
	errorCode,
	composeCode,
	curryCode,
	withCode,
	useCode,
	mapCode,
	filterCode,
	printCode,
	inputCode,
	spawnCode,
	newChannelCode,
})
