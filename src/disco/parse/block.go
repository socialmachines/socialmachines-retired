// Copyright 2013 Mark Stahl. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can found in the BSD-LICENSE file.

package parse

import (
	"disco/ast"
	"disco/rt"
	"disco/scan"
)

// block := 
//	'{' [statements] '}'
//
func (p *Parser) parseBlock() (b *ast.Block) {
	p.expect(scan.LBRACE)

	b = &ast.Block{Statements: p.parseStatements()}

	p.expect(scan.RBRACE)

	return
}

// statements :=
//	[expression ['.' statements]]
//
func (p *Parser) parseStatements() []rt.Expr {
	var stmts []rt.Expr
	for p.tok != scan.RBRACE {
		stmts = append(stmts, p.parseExpr())

		if p.tok == scan.PERIOD {
			p.next()
		}
	}

	return stmts
}
