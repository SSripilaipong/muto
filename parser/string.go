package parser

import (
	"fmt"

	fileParser "github.com/SSripilaipong/muto/parser/file"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"
)

func StringsToStatements(codes []string) ([]stBase.Statement, error) {
	var statements []stBase.Statement
	for _, code := range codes {
		file, err := fileParser.ParseFileFromString(code).Return()
		if err != nil {
			return nil, fmt.Errorf("parse code string: %w", err)
		}
		statements = append(statements, file.Statements()...)
	}
	return statements, nil
}

func StringsToStatementsOrPanic(codes []string) []stBase.Statement {
	statements, err := StringsToStatements(codes)
	if err != nil {
		panic(err)
	}
	return statements
}
