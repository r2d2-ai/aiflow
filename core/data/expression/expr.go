package expression

import (
	"github.com/r2d2-ai/ai-box/core/data"
	"github.com/r2d2-ai/ai-box/core/data/resolve"
)

type Factory interface {
	NewExpr(exprStr string) (Expr, error)
}

type Expr interface {
	Eval(scope data.Scope) (interface{}, error)
}

type FactoryCreatorFunc func(resolve.CompositeResolver) Factory
