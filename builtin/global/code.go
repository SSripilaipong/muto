package global

import (
	"fmt"

	fileParser "github.com/SSripilaipong/muto/parser/file"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
)

var rawStatements = codesToStatements([]string{
	doCode,
	matchCode,
	retCode,
	composeCode,
	curryCode,
	withCode,
	useCode,
	mapCode,
	filterCode,
	printCode,
	inputCode,
})

func codesToStatements(codes []string) []stBase.Statement {
	var statements []stBase.Statement
	for _, code := range codes {
		file, err := fileParser.ParseFileFromString(code).Return()
		if err != nil {
			panic(fmt.Errorf("cannot build global module: %w", err))
		}
		statements = append(statements, file.Statements()...)
	}
	return statements
}
