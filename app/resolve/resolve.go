package resolve

import "github.com/r2d2-ai/aiflow/data"

var r simpleResolve

func SetAppResolver(sr simpleResolve) {
	r = sr
}

type simpleResolve interface {
	Resolve(resolveDirective string, scope data.Scope) (value interface{}, err error)
}

func Resolve(resolveDirective string, scope data.Scope) (value interface{}, err error) {
	return r.Resolve(resolveDirective, scope)
}
