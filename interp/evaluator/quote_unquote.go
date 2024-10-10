package evaluator

import (
	"github.com/egriff89/monkey/interp/ast"
	"github.com/egriff89/monkey/interp/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
