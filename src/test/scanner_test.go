// Copyright (C) 2013 Mark Stahl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package test

import (
	"disco/file"
	"disco/scan"
	"testing"
)

func TestUnary(t *testing.T) {
	expected := []scan.Token{
		scan.BINARY,
		scan.NAME,
		scan.IDENT,
		scan.DEFINE,
		scan.NAME,
		scan.PERIOD,
	}

	testTokens(t, "+ True not => False.", expected...)
}

func TestBinary(t *testing.T) {
	expected := []scan.Token{
		scan.BINARY,
		scan.NAME,
		scan.BINARY,
		scan.IDENT,
		scan.DEFINE,
		scan.IDENT,
		scan.KEYWORD,
		scan.LBRACK,
		scan.NAME,
		scan.RBRACK,
		scan.KEYWORD,
		scan.LBRACK,
		scan.NAME,
		scan.RBRACK,
		scan.PERIOD,
	}

	received := "+ True ^ aBool => aBool ifTrue: { False } ifFalse: { True }."
	testTokens(t, received, expected...)
}

func TestKeyword(t *testing.T) {
	expected := []scan.Token {
		scan.BINARY,
		scan.NAME,
		scan.KEYWORD,
		scan.IDENT,
		scan.KEYWORD,
		scan.IDENT,
		scan.DEFINE,
		scan.IDENT,
		scan.IDENT,
		scan.PERIOD,
	}
	
	received := "+ True ifTrue: tBlock ifFalse: fBlock => tBlock value."
	testTokens(t, received, expected...)
}

func initScanner(expr string) scan.Scanner {
	src := []byte(expr)

	fset := file.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))

	var s scan.Scanner
	s.Init(file, src, nil)

	return s
}

var expRecv = "Expected (%s) -- Received (%s)\n"

func testTokens(t *testing.T, expr string, tokens ...scan.Token) {
	s := initScanner(expr)

	for _, token := range tokens {
		_, tok, _ := s.Scan()
		if tok != token {
			t.Fatalf(expRecv, token, tok)
		}
	}
}