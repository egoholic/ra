package interpreter

import (
	"github.com/egoholic/ra/interpreter/algebra"
	"github.com/egoholic/ra/space"
)

type Interpreter struct {
	space *space.Space
}

func New(s *space.Space) *Interpreter {
	return &Interpreter{s}
}

func (i *Interpreter) Exec(*algebra.Expr) {

}

func execLET()     {}
func execJOIN()    {}
func execPROJECT() {}
func execUNION()   {}
func execSUB()     {}
func execGROUP()   {}
