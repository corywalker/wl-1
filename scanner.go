// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"github.com/cznic/golex/lex"
)

func (lx *lexer) scan() int {
	c := lx.Enter()

	/* classes */

yystate0:
	yyrule := -1
	_ = yyrule
	c = lx.Rule0()
	if lx.err != nil {
		return -1
	}

	goto yystart1

	goto yystate0 // silence unused label error
	goto yyAction // silence unused label error
yyAction:
	switch yyrule {
	case 1:
		goto yyrule1
	case 2:
		goto yyrule2
	case 3:
		goto yyrule3
	case 4:
		goto yyrule4
	case 5:
		goto yyrule5
	case 6:
		goto yyrule6
	case 7:
		goto yyrule7
	case 8:
		goto yyrule8
	case 9:
		goto yyrule9
	case 10:
		goto yyrule10
	case 11:
		goto yyrule11
	case 12:
		goto yyrule12
	case 13:
		goto yyrule13
	case 14:
		goto yyrule14
	case 15:
		goto yyrule15
	case 16:
		goto yyrule16
	case 17:
		goto yyrule17
	case 18:
		goto yyrule18
	case 19:
		goto yyrule19
	case 20:
		goto yyrule20
	case 21:
		goto yyrule21
	case 22:
		goto yyrule22
	case 23:
		goto yyrule23
	case 24:
		goto yyrule24
	case 25:
		goto yyrule25
	case 26:
		goto yyrule26
	case 27:
		goto yyrule27
	case 28:
		goto yyrule28
	}
	goto yystate1 // silence unused label error
yystate1:
	c = lx.Next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate3
	case c == '#':
		goto yystate7
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082':
		goto yystate10
	case c == '&':
		goto yystate18
	case c == '(':
		goto yystate20
	case c == '-':
		goto yystate24
	case c == '.':
		goto yystate26
	case c == '/':
		goto yystate30
	case c == ':':
		goto yystate37
	case c == '<':
		goto yystate41
	case c == '=':
		goto yystate44
	case c == '>':
		goto yystate49
	case c == '@':
		goto yystate51
	case c == '[':
		goto yystate53
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	case c == ']':
		goto yystate55
	case c == '_':
		goto yystate11
	case c == '|':
		goto yystate57
	case c >= '0' && c <= '9':
		goto yystate36
	}

yystate2:
	c = lx.Next()
	yyrule = 1
	lx.Mark()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate4
	case c == '\\':
		goto yystate5
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate3
	}

yystate4:
	c = lx.Next()
	yyrule = 23
	lx.Mark()
	goto yyrule23

yystate5:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate6
	case c == '\\':
		goto yystate5
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate3
	}

yystate6:
	c = lx.Next()
	yyrule = 23
	lx.Mark()
	switch {
	default:
		goto yyrule23
	case c == '"':
		goto yystate4
	case c == '\\':
		goto yystate5
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate3
	}

yystate7:
	c = lx.Next()
	yyrule = 28
	lx.Mark()
	switch {
	default:
		goto yyrule28
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082':
		goto yystate8
	case c >= '0' && c <= '9':
		goto yystate9
	}

yystate8:
	c = lx.Next()
	yyrule = 28
	lx.Mark()
	switch {
	default:
		goto yyrule28
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate8
	}

yystate9:
	c = lx.Next()
	yyrule = 28
	lx.Mark()
	switch {
	default:
		goto yyrule28
	case c >= '0' && c <= '9':
		goto yystate9
	}

yystate10:
	c = lx.Next()
	yyrule = 25
	lx.Mark()
	switch {
	default:
		goto yyrule25
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate10
	case c == '_':
		goto yystate11
	case c == '`':
		goto yystate16
	}

yystate11:
	c = lx.Next()
	yyrule = 27
	lx.Mark()
	switch {
	default:
		goto yyrule27
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082':
		goto yystate12
	case c == '.':
		goto yystate13
	case c == '_':
		goto yystate14
	}

yystate12:
	c = lx.Next()
	yyrule = 27
	lx.Mark()
	switch {
	default:
		goto yyrule27
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate12
	case c == '.':
		goto yystate13
	}

yystate13:
	c = lx.Next()
	yyrule = 27
	lx.Mark()
	goto yyrule27

yystate14:
	c = lx.Next()
	yyrule = 27
	lx.Mark()
	switch {
	default:
		goto yyrule27
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082':
		goto yystate12
	case c == '.':
		goto yystate13
	case c == '_':
		goto yystate15
	}

yystate15:
	c = lx.Next()
	yyrule = 27
	lx.Mark()
	switch {
	default:
		goto yyrule27
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082':
		goto yystate12
	case c == '.':
		goto yystate13
	}

yystate16:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082':
		goto yystate17
	}

yystate17:
	c = lx.Next()
	yyrule = 25
	lx.Mark()
	switch {
	default:
		goto yyrule25
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate17
	case c == '`':
		goto yystate16
	}

yystate18:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '&':
		goto yystate19
	}

yystate19:
	c = lx.Next()
	yyrule = 3
	lx.Mark()
	goto yyrule3

yystate20:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate21
	}

yystate21:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate22
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate21
	}

yystate22:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == ')':
		goto yystate23
	case c == '*':
		goto yystate22
	case c >= '\x01' && c <= '(' || c >= '+' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate21
	}

yystate23:
	c = lx.Next()
	yyrule = 2
	lx.Mark()
	goto yyrule2

yystate24:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '>':
		goto yystate25
	}

yystate25:
	c = lx.Next()
	yyrule = 4
	lx.Mark()
	goto yyrule4

yystate26:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate27
	}

yystate27:
	c = lx.Next()
	yyrule = 24
	lx.Mark()
	switch {
	default:
		goto yyrule24
	case c == 'E' || c == 'e':
		goto yystate28
	case c >= '0' && c <= '9':
		goto yystate27
	}

yystate28:
	c = lx.Next()
	yyrule = 24
	lx.Mark()
	switch {
	default:
		goto yyrule24
	case c == '+' || c == '-' || c >= '0' && c <= '9':
		goto yystate29
	}

yystate29:
	c = lx.Next()
	yyrule = 24
	lx.Mark()
	switch {
	default:
		goto yyrule24
	case c >= '0' && c <= '9':
		goto yystate29
	}

yystate30:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate31
	case c == '/':
		goto yystate32
	case c == ';':
		goto yystate34
	case c == '@':
		goto yystate35
	}

yystate31:
	c = lx.Next()
	yyrule = 5
	lx.Mark()
	goto yyrule5

yystate32:
	c = lx.Next()
	yyrule = 6
	lx.Mark()
	switch {
	default:
		goto yyrule6
	case c == '.':
		goto yystate33
	}

yystate33:
	c = lx.Next()
	yyrule = 7
	lx.Mark()
	goto yyrule7

yystate34:
	c = lx.Next()
	yyrule = 8
	lx.Mark()
	goto yyrule8

yystate35:
	c = lx.Next()
	yyrule = 9
	lx.Mark()
	goto yyrule9

yystate36:
	c = lx.Next()
	yyrule = 26
	lx.Mark()
	switch {
	default:
		goto yyrule26
	case c == '.':
		goto yystate27
	case c == 'E' || c == 'e':
		goto yystate28
	case c >= '0' && c <= '9':
		goto yystate36
	}

yystate37:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == ':':
		goto yystate38
	case c == '=':
		goto yystate39
	case c == '>':
		goto yystate40
	}

yystate38:
	c = lx.Next()
	yyrule = 10
	lx.Mark()
	goto yyrule10

yystate39:
	c = lx.Next()
	yyrule = 11
	lx.Mark()
	goto yyrule11

yystate40:
	c = lx.Next()
	yyrule = 12
	lx.Mark()
	goto yyrule12

yystate41:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate42
	case c == '>':
		goto yystate43
	}

yystate42:
	c = lx.Next()
	yyrule = 13
	lx.Mark()
	goto yyrule13

yystate43:
	c = lx.Next()
	yyrule = 14
	lx.Mark()
	goto yyrule14

yystate44:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate45
	case c == '=':
		goto yystate47
	}

yystate45:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate46
	}

yystate46:
	c = lx.Next()
	yyrule = 15
	lx.Mark()
	goto yyrule15

yystate47:
	c = lx.Next()
	yyrule = 16
	lx.Mark()
	switch {
	default:
		goto yyrule16
	case c == '=':
		goto yystate48
	}

yystate48:
	c = lx.Next()
	yyrule = 17
	lx.Mark()
	goto yyrule17

yystate49:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate50
	}

yystate50:
	c = lx.Next()
	yyrule = 18
	lx.Mark()
	goto yyrule18

yystate51:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '@':
		goto yystate52
	}

yystate52:
	c = lx.Next()
	yyrule = 19
	lx.Mark()
	goto yyrule19

yystate53:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '[':
		goto yystate54
	}

yystate54:
	c = lx.Next()
	yyrule = 20
	lx.Mark()
	goto yyrule20

yystate55:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == ']':
		goto yystate56
	}

yystate56:
	c = lx.Next()
	yyrule = 22
	lx.Mark()
	goto yyrule22

yystate57:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '|':
		goto yystate58
	}

yystate58:
	c = lx.Next()
	yyrule = 21
	lx.Mark()
	goto yyrule21

yyrule1: // [ \t\r]+

	goto yystate0
yyrule2: // "(*"([^*\x80]|\*+[^*)\x80])*\*+\)

	goto yystate0
yyrule3: // "&&"
	{
		return AND
	}
yyrule4: // "->"
	{
		return RULE
	}
yyrule5: // "/."
	{
		return REPLACEALL
	}
yyrule6: // "//"
	{
		return POSTFIX
	}
yyrule7: // "//."
	{
		return REPLACEREP
	}
yyrule8: // "/;"
	{
		return CONDITION
	}
yyrule9: // "/@"
	{
		return MAP
	}
yyrule10: // "::"
	{
		return MESSAGE
	}
yyrule11: // ":="
	{
		return SET_DELAYED
	}
yyrule12: // ":>"
	{
		return RULEDELAYED
	}
yyrule13: // "<="
	{
		return LEQ
	}
yyrule14: // "<>"
	{
		return STRINGJOIN
	}
yyrule15: // "=!="
	{
		return UNSAME
	}
yyrule16: // "=="
	{
		return EQUAL
	}
yyrule17: // "==="
	{
		return SAME
	}
yyrule18: // ">="
	{
		return GEQ
	}
yyrule19: // "@@"
	{
		return APPLY
	}
yyrule20: // "[["
	{
		return lx.push(LPART)
	}
yyrule21: // "||"
	{
		return OR
	}
yyrule22: // "]]"
	{
		if lx.pop() == LPART {
			return RPART
		}
		if la := lx.Lookahead(); la.Rune != 0 {
			lx.Unget(la)
		}
		lx.Unget(lex.NewChar(lx.First.Pos()+1, ']'))
		return ']'
	}
yyrule23: // \"([^"]|\\.)*\"
	{
		return STRING
	}
yyrule24: // {float}
	{
		return FLOAT
	}
yyrule25: // {ident}(`{ident})*
	{
		return IDENT
	}
yyrule26: // {int}
	{
		return INT
	}
yyrule27: // {pattern}
	{
		return PATTERN
	}
yyrule28: // {slot}
	{
		return SLOT
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	if c, ok := lx.Abort(); ok {
		switch c {
		case '(', '[', '{':
			lx.push(c)
		case ')', ']', '}':
			lx.pop()
		}
		return c
	}

	goto yyAction
}
