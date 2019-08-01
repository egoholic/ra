package algebra

import (
	attr "github.com/egoholic/ra/attribute"
	"github.com/egoholic/ra/tuple"
)

type Expr struct {
	left    *Expr
	right   *Expr
	pattern *Pattern
}
type Pattern struct{}

func VAR(name string, r *Expr) *Expr {

}

type Matcher func(*tuple.Tuple, *tuple.Tuple) bool

func JOIN(match Matcher, r1, r2 *Expr) *Expr {

}

func GROUP(aggregators map[string]func(interface{}, *attr.Attr) interface{}, r *Expr) *Expr {

}

func PROJECT(transformers map[string]func(interface{}) interface{}) {

}
func UNION(r1, r2 *Expr) *Expr {

}

func SUB(match func(*tuple.Tuple) bool, r *Expr) *Expr {

}
