package mathinterval

import (
	"math"
	"strconv"
	"strings"
)

// Expr is object parse from string expression
type Expr struct {
	expr       string
	max        float64
	min        float64
	minInclude bool
	maxInclude bool
}

// Get expression from string input
func Get(sExpression string) Expr {
	var res Expr
	res.expr = sExpression
	a := strings.Split(sExpression, ",")
	if len(a) != 2 {
		panic("wrong expression format")
	}
	left := a[0]
	right := a[1]

	switch left[0:1] {
	case "[":
		res.minInclude = true
	case "(":
		res.minInclude = false
	default:
		panic("wrong expression format")
	}

	switch left[1:] {
	case "-inf":
		res.min = math.Inf(-1)
	default:
		minnum, err := strconv.ParseFloat(left[1:], 64)
		if err != nil {
			panic(err)
		}
		res.min = minnum
	}

	switch right[len(right)-1 : len(right)] {
	case "]":
		res.maxInclude = true
	case ")":
		res.maxInclude = false
	default:
		panic("wrong expression format")
	}

	switch right[:len(right)-1] {
	case "inf":
		res.max = math.Inf(1)
	default:
		maxnum, err := strconv.ParseFloat(right[:len(right)-1], 64)
		if err != nil {
			panic(err)
		}
		res.max = maxnum
	}

	return res
}

// Hit ...
func (exp Expr) Hit(v float64) bool {
	var a, b bool
	if exp.maxInclude == true && v <= exp.max {
		a = true
	}

	if exp.maxInclude == false && v < exp.max {
		a = true
	}

	if exp.minInclude == true && v >= exp.min {
		b = true
	}

	if exp.minInclude == false && v > exp.min {
		b = true
	}

	return a && b
}
