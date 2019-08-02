package algebra

import (
	"github.com/egoholic/ra/tuple"
)

const (
	DUMB_OP    = 'D'
	LET_OP     = 'L'
	JOIN_OP    = 'J'
	PROJECT_OP = 'P'
	UNION_OP   = 'U'
	SUB_OP     = 'S'
)

type Expr struct {
	Operation int
	Left      *Expr
	Right     *Expr
	Param     interface{}
}

type Pattern struct{}

func DUMB() *Expr {
	return &Expr{
		Operation: DUMB_OP,
		Left:      nil,
		Right:     nil,
		Param:     nil,
	}
}
func LET(name string, r *Expr) *Expr {
	return &Expr{
		Operation: LET_OP,
		Left:      nil,
		Right:     r,
		Param:     name,
	}
}
func JOIN(match func(*tuple.Tuple, *tuple.Tuple) bool, r1, r2 *Expr) *Expr {
	return &Expr{
		Operation: JOIN_OP,
		Left:      r1,
		Right:     r2,
		Param:     match,
	}
}
func PROJECT(transformers map[string]func(interface{}) interface{}, r *Expr) *Expr {
	return &Expr{
		Operation: PROJECT_OP,
		Left:      nil,
		Right:     r,
		Param:     transformers,
	}
}
func UNION(r1, r2 *Expr) *Expr {
	return &Expr{
		Operation: UNION_OP,
		Left:      r1,
		Right:     r2,
		Param:     nil,
	}
}
func SUB(match func(*tuple.Tuple) bool, r *Expr) *Expr {
	return &Expr{
		Operation: SUB_OP,
		Left:      nil,
		Right:     r,
		Param:     match,
	}
}
