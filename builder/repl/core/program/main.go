package program

import (
	"github.com/SSripilaipong/muto/program"
)

type Wrapper struct {
	program program.Program
}

func New(prog program.Program) Wrapper {
	return Wrapper{program: prog}
}
