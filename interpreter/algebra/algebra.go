package algebra

import (
	"github.com/egoholic/ra/tuple"
)

const (
	LET_OP     = 'L'
	JOIN_OP    = 'J'
	PROJECT_OP = 'P'
	UNION_OP   = 'U'
	SUB_OP     = 'S'
	GROUP_OP   = 'G'
)

type Expr struct {
	Operation int
	Left      *Expr
	Right     *Expr
	Param1    interface{}
	Param2    interface{}
}

type Pattern struct{}

func LET(name string, r *Expr) *Expr {
	return &Expr{
		Operation: LET_OP,
		Left:      nil,
		Right:     r,
		Param1:    name,
		Param2:    nil,
	}
}
func JOIN(match func(*tuple.Tuple, *tuple.Tuple) bool, r1, r2 *Expr) *Expr {
	return &Expr{
		Operation: JOIN_OP,
		Left:      r1,
		Right:     r2,
		Param1:    match,
		Param2:    nil,
	}
}
func PROJECT(transformers map[string]func(interface{}) interface{}, r *Expr) *Expr {
	return &Expr{
		Operation: PROJECT_OP,
		Left:      nil,
		Right:     r,
		Param1:    transformers,
		Param2:    nil,
	}
}
func UNION(r1, r2 *Expr) *Expr {
	return &Expr{
		Operation: UNION_OP,
		Left:      r1,
		Right:     r2,
		Param1:    nil,
		Param2:    nil,
	}
}
func SUB(match func(*tuple.Tuple) bool, r *Expr) *Expr {
	return &Expr{
		Operation: SUB_OP,
		Left:      nil,
		Right:     r,
		Param1:    match,
		Param2:    nil,
	}
}

func GROUP(attrNames []string, aggregators map[string]func(*tuple.Tuple) interface{}, r *Expr) *Expr {
	return &Expr{
		Operation: GROUP_OP,
		Left:      nil,
		Right:     nil,
		Param1:    attrNames,
		Param2:    aggregators,
	}
}
