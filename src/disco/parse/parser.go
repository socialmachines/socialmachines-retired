// Copyright 2013 Mark Stahl. All rights reserved.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the BSD-LICENSE file.

package parse

import (
	"disco/ast"
	"disco/file"
	"disco/scan"
	"fmt"
)

type Parser struct {
	file    *file.File
	scanner scan.Scanner

	Exprs []ast.Expression

	pos file.Pos
	tok scan.Token
	lit string
}

func (p *Parser) Init(f *file.File, src []byte) {
	p.file = f
	p.scanner.Init(p.file, src, nil)

	p.next()
}

func (p *Parser) Parse() {
	var exprs []ast.Expression

	for p.tok != scan.EOF {
		var expr ast.Expression
		switch {
		case p.tok == scan.COMMENT:
			expr = &ast.Comment{Start: p.pos, Text: p.lit}
		case p.isDefine():
			d := &ast.Define{Start: p.pos, Type: ast.EXT}
			expr = p.parseExtDefine(d)
		case p.tok == scan.NAME || p.tok == scan.IDENT:
			r := &ast.Literal{Start: p.pos, Name: p.lit}
			expr = p.parseExpr(r)
		case p.tok == scan.LBRACK:
			b := &ast.Block{Start: p.pos}
			expr = p.parseBlock(b)
		}
		p.next()

		if expr != nil {
			exprs = append(exprs, expr)
		}
	}
	p.Exprs = exprs
}

func (p *Parser) next() {
	p.pos, p.tok, p.lit = p.scanner.Scan()
}

func (p *Parser) expect(tok scan.Token) (lit string) {
	lit = p.lit
	if p.tok != tok {
		fmt.Printf("expected '%s', received '%s'\n", tok, p.tok)
		return ""
	}
	p.next()
	return
}

func (p *Parser) isDefine() bool {
	return p.tok == scan.BINARY && p.lit == "+"
}