
package lib;

type LabelledNumber struct {
	value float64
	label *string
}

type Callable interface {
	arity(args *[]LabelledNumber) uint16
	call(args *[]LabelledNumber) LabelledNumber
}

type Expr interface {
	Type() ExprType
}
type ExprType int8;
const (
	BINARY ExprType = iota
	GROUPING
	LITERAL
	UNARY
)

type BinaryExpr struct {
	left *Expr
	operator Token
	right *Expr
}
func (b *BinaryExpr) Type() ExprType {
	return BINARY;
}

type GroupingExpr struct {
	operator Token
	expression *Expr
}
func (b *GroupingExpr) Type() ExprType {
	return GROUPING;
}

type LiteralExpr struct {
	value float64
	label *string
}
func (b *LiteralExpr) Type() ExprType {
	return LITERAL;
}

type UnaryExpr struct {
	operator Token
	right *Expr
}
func (b *UnaryExpr) Type() ExprType {
	return UNARY;
}


type Stmt interface {
	Type() StmtType
}
type StmtType int8;
const (
	EXPRESSION StmtType = iota
	PRINTER
)

type ExpressionStmt struct {
	expression *Expr
}
func (b *ExpressionStmt) Type() StmtType {
	return EXPRESSION;
}
type PrintStmt struct {
	expression *Expr
}
func (b *PrintStmt) Type() StmtType {
	return PRINTER;
}
