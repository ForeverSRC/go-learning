package eval

import (
	"fmt"
	"math"
)
// An Expr is an arithmetic expression
type Expr interface {
	// Eval returns the value of this Expr in the environment env
	Eval(env Env) float64
}

// A Var identifies a variable
type Var string

// A literal is a numeric constant
type literal float64

//A unary represents a unary operator expression
type unary struct {
	op rune
	x  Expr
}

// A binary represents a binary operator expression
type binary struct {
	op   rune
	x, y Expr
}

// A call represents a function call expression
type call struct {
	fn   string
	args []Expr
}

type Env map[Var]float64

func (v Var) Eval(env Env) float64{
	return env[v]
}

func (l literal)Eval(_ Env) float64{
	return float64(l)
}

func (u unary) Eval(env Env)float64{
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator:%q",u.op))
}

func (b binary)Eval(env Env)float64{
	switch b.op {
	case '+':
		return b.x.Eval(env)+b.y.Eval(env)
	case '-':
		return b.x.Eval(env)-b.y.Eval(env)
	case '*':
		return b.x.Eval(env)*b.y.Eval(env)
	case '/':
		return b.x.Eval(env)/b.y.Eval(env)

	}
	panic(fmt.Sprintf("unsupported binary operator:%q",b.op))
}

func (c call)Eval(env Env)float64{
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env),c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s",c.fn))
}