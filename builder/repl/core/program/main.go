package program

import (
	"github.com/SSripilaipong/muto/program"
)

type Wrapper struct {
	program program.Program
	printer Printer
}

type Printer interface {
	Print(x string)
}

func New(prog program.Program, printer Printer) Wrapper {
	return Wrapper{program: prog, printer: printer}
}
