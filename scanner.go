// Code generated by golex. DO NOT EDIT.

// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"fmt"

	"github.com/cznic/golex/lex"
)

func (lx *lexer) scan() (r int) {
	// fmt.Printf("%T.scan\n", lx)
	// defer func() { fmt.Printf("\t%T.scan: %U %s\n", lx, r, yySymName(r)) }()
	c := lx.Enter()

	/* classes */

yystate0:
	yyrule := -1
	_ = yyrule
	c = lx.Rule0()

	switch yyt := lx.sc; yyt {
	default:
		panic(fmt.Errorf(`invalid start condition %d`, yyt))
	case 0: // start condition: INITIAL
		goto yystart1
	case 1: // start condition: S
		goto yystart118
	}

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
	case 29:
		goto yyrule29
	case 30:
		goto yyrule30
	case 31:
		goto yyrule31
	case 32:
		goto yyrule32
	case 33:
		goto yyrule33
	case 34:
		goto yyrule34
	case 35:
		goto yyrule35
	case 36:
		goto yyrule36
	case 37:
		goto yyrule37
	case 38:
		goto yyrule38
	case 39:
		goto yyrule39
	case 40:
		goto yyrule40
	case 41:
		goto yyrule41
	case 42:
		goto yyrule42
	case 43:
		goto yyrule43
	case 44:
		goto yyrule44
	case 45:
		goto yyrule45
	case 46:
		goto yyrule46
	case 47:
		goto yyrule47
	case 48:
		goto yyrule48
	case 49:
		goto yyrule49
	case 50:
		goto yyrule50
	case 51:
		goto yyrule51
	case 52:
		goto yyrule52
	case 53:
		goto yyrule53
	case 54:
		goto yyrule54
	case 55:
		goto yyrule55
	case 56:
		goto yyrule56
	case 57:
		goto yyrule57
	case 58:
		goto yyrule58
	case 59:
		goto yyrule59
	case 60:
		goto yyrule60
	case 61:
		goto yyrule61
	case 62:
		goto yyrule62
	case 63:
		goto yyrule63
	case 64:
		goto yyrule64
	case 65:
		goto yyrule65
	case 66:
		goto yyrule66
	case 67:
		goto yyrule67
	case 68:
		goto yyrule68
	case 69:
		goto yyrule69
	case 70:
		goto yyrule70
	case 71:
		goto yyrule71
	case 72:
		goto yyrule72
	case 73:
		goto yyrule73
	case 74:
		goto yyrule74
	case 75:
		goto yyrule75
	case 76:
		goto yyrule76
	case 77:
		goto yyrule77
	case 78:
		goto yyrule78
	case 79:
		goto yyrule79
	case 80:
		goto yyrule80
	case 81:
		goto yyrule81
	case 82:
		goto yyrule82
	case 83:
		goto yyrule83
	case 84:
		goto yyrule84
	case 85:
		goto yyrule85
	case 86:
		goto yyrule86
	case 87:
		goto yyrule87
	case 88:
		goto yyrule88
	case 89:
		goto yyrule89
	}
	goto yystate1 // silence unused label error
yystate1:
	c = lx.Next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate5
	case c == '#':
		goto yystate6
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082' || c == '\u0083':
		goto yystate13
	case c == '&':
		goto yystate22
	case c == '(':
		goto yystate25
	case c == '*':
		goto yystate29
	case c == '+':
		goto yystate31
	case c == '-':
		goto yystate33
	case c == '.':
		goto yystate36
	case c == '/':
		goto yystate41
	case c == ':':
		goto yystate53
	case c == '<':
		goto yystate57
	case c == '=':
		goto yystate61
	case c == '>':
		goto yystate66
	case c == '@':
		goto yystate68
	case c == '[':
		goto yystate72
	case c == '\'':
		goto yystate24
	case c == '\\':
		goto yystate74
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate2
	case c == '\u0084':
		goto yystate87
	case c == '\u0085':
		goto yystate88
	case c == '\u0086':
		goto yystate89
	case c == '\u0087':
		goto yystate90
	case c == '\u0088':
		goto yystate91
	case c == '\u0089':
		goto yystate92
	case c == '\u008a':
		goto yystate93
	case c == '\u008b':
		goto yystate94
	case c == '\u008c':
		goto yystate95
	case c == '\u008d':
		goto yystate96
	case c == '\u008e':
		goto yystate97
	case c == '\u008f':
		goto yystate98
	case c == '\u0090':
		goto yystate99
	case c == '\u0091':
		goto yystate100
	case c == '\u0092':
		goto yystate101
	case c == '\u0093':
		goto yystate102
	case c == '\u0094':
		goto yystate103
	case c == '\u0095':
		goto yystate104
	case c == '\u0096':
		goto yystate105
	case c == '\u0097':
		goto yystate106
	case c == '\u0098':
		goto yystate107
	case c == '\u0099':
		goto yystate108
	case c == '\u009a':
		goto yystate109
	case c == '\u009b':
		goto yystate110
	case c == '\u009c':
		goto yystate111
	case c == '\u009d':
		goto yystate112
	case c == '\u009e':
		goto yystate113
	case c == '\u009f':
		goto yystate114
	case c == '\u00a0':
		goto yystate115
	case c == ']':
		goto yystate82
	case c == '_':
		goto yystate14
	case c == '`':
		goto yystate84
	case c == '|':
		goto yystate85
	case c == '¡':
		goto yystate116
	case c == '¢':
		goto yystate117
	case c >= '0' && c <= '9':
		goto yystate49
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
	case c == '=':
		goto yystate4
	}

yystate4:
	c = lx.Next()
	yyrule = 39
	lx.Mark()
	goto yyrule39

yystate5:
	c = lx.Next()
	yyrule = 3
	lx.Mark()
	goto yyrule3

yystate6:
	c = lx.Next()
	yyrule = 89
	lx.Mark()
	switch {
	default:
		goto yyrule89
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082' || c == '\u0083':
		goto yystate7
	case c == '`':
		goto yystate12
	case c >= '0' && c <= '9':
		goto yystate11
	}

yystate7:
	c = lx.Next()
	yyrule = 89
	lx.Mark()
	switch {
	default:
		goto yyrule89
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= '\u0081' && c <= '\u0083':
		goto yystate7
	case c == '`':
		goto yystate8
	}

yystate8:
	c = lx.Next()
	yyrule = 89
	lx.Mark()
	switch {
	default:
		goto yyrule89
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082' || c == '\u0083':
		goto yystate9
	}

yystate9:
	c = lx.Next()
	yyrule = 89
	lx.Mark()
	switch {
	default:
		goto yyrule89
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= '\u0081' && c <= '\u0083':
		goto yystate9
	case c == '`':
		goto yystate10
	}

yystate10:
	c = lx.Next()
	yyrule = 89
	lx.Mark()
	goto yyrule89

yystate11:
	c = lx.Next()
	yyrule = 89
	lx.Mark()
	switch {
	default:
		goto yyrule89
	case c >= '0' && c <= '9':
		goto yystate11
	}

yystate12:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082' || c == '\u0083':
		goto yystate9
	}

yystate13:
	c = lx.Next()
	yyrule = 86
	lx.Mark()
	switch {
	default:
		goto yyrule86
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= '\u0081' && c <= '\u0083':
		goto yystate13
	case c == '_':
		goto yystate14
	case c == '`':
		goto yystate19
	}

yystate14:
	c = lx.Next()
	yyrule = 88
	lx.Mark()
	switch {
	default:
		goto yyrule88
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082' || c == '\u0083':
		goto yystate15
	case c == '.':
		goto yystate16
	case c == '_':
		goto yystate17
	}

yystate15:
	c = lx.Next()
	yyrule = 88
	lx.Mark()
	switch {
	default:
		goto yyrule88
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= '\u0081' && c <= '\u0083':
		goto yystate15
	case c == '.':
		goto yystate16
	}

yystate16:
	c = lx.Next()
	yyrule = 88
	lx.Mark()
	goto yyrule88

yystate17:
	c = lx.Next()
	yyrule = 88
	lx.Mark()
	switch {
	default:
		goto yyrule88
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082' || c == '\u0083':
		goto yystate15
	case c == '.':
		goto yystate16
	case c == '_':
		goto yystate18
	}

yystate18:
	c = lx.Next()
	yyrule = 88
	lx.Mark()
	switch {
	default:
		goto yyrule88
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082' || c == '\u0083':
		goto yystate15
	case c == '.':
		goto yystate16
	}

yystate19:
	c = lx.Next()
	yyrule = 86
	lx.Mark()
	switch {
	default:
		goto yyrule86
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082' || c == '\u0083':
		goto yystate20
	}

yystate20:
	c = lx.Next()
	yyrule = 86
	lx.Mark()
	switch {
	default:
		goto yyrule86
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= '\u0081' && c <= '\u0083':
		goto yystate20
	case c == '`':
		goto yystate21
	}

yystate21:
	c = lx.Next()
	yyrule = 86
	lx.Mark()
	goto yyrule86

yystate22:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '&':
		goto yystate23
	}

yystate23:
	c = lx.Next()
	yyrule = 18
	lx.Mark()
	goto yyrule18

yystate24:
	c = lx.Next()
	yyrule = 19
	lx.Mark()
	goto yyrule19

yystate25:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate26
	}

yystate26:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate27
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate26
	}

yystate27:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == ')':
		goto yystate28
	case c == '*':
		goto yystate27
	case c >= '\x01' && c <= '(' || c >= '+' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate26
	}

yystate28:
	c = lx.Next()
	yyrule = 2
	lx.Mark()
	goto yyrule2

yystate29:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate30
	}

yystate30:
	c = lx.Next()
	yyrule = 20
	lx.Mark()
	goto yyrule20

yystate31:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '+':
		goto yystate32
	}

yystate32:
	c = lx.Next()
	yyrule = 21
	lx.Mark()
	goto yyrule21

yystate33:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '-':
		goto yystate34
	case c == '>':
		goto yystate35
	}

yystate34:
	c = lx.Next()
	yyrule = 22
	lx.Mark()
	goto yyrule22

yystate35:
	c = lx.Next()
	yyrule = 23
	lx.Mark()
	goto yyrule23

yystate36:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate37
	}

yystate37:
	c = lx.Next()
	yyrule = 85
	lx.Mark()
	switch {
	default:
		goto yyrule85
	case c == '`':
		goto yystate38
	case c >= '0' && c <= '9':
		goto yystate37
	}

yystate38:
	c = lx.Next()
	yyrule = 85
	lx.Mark()
	switch {
	default:
		goto yyrule85
	case c == '`':
		goto yystate40
	case c >= '0' && c <= '9':
		goto yystate39
	}

yystate39:
	c = lx.Next()
	yyrule = 85
	lx.Mark()
	switch {
	default:
		goto yyrule85
	case c >= '0' && c <= '9':
		goto yystate39
	}

yystate40:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate39
	}

yystate41:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate42
	case c == '.':
		goto yystate43
	case c == '/':
		goto yystate44
	case c == ';':
		goto yystate47
	case c == '@':
		goto yystate48
	}

yystate42:
	c = lx.Next()
	yyrule = 24
	lx.Mark()
	goto yyrule24

yystate43:
	c = lx.Next()
	yyrule = 25
	lx.Mark()
	goto yyrule25

yystate44:
	c = lx.Next()
	yyrule = 26
	lx.Mark()
	switch {
	default:
		goto yyrule26
	case c == '.':
		goto yystate45
	case c == '@':
		goto yystate46
	}

yystate45:
	c = lx.Next()
	yyrule = 27
	lx.Mark()
	goto yyrule27

yystate46:
	c = lx.Next()
	yyrule = 28
	lx.Mark()
	goto yyrule28

yystate47:
	c = lx.Next()
	yyrule = 29
	lx.Mark()
	goto yyrule29

yystate48:
	c = lx.Next()
	yyrule = 30
	lx.Mark()
	goto yyrule30

yystate49:
	c = lx.Next()
	yyrule = 87
	lx.Mark()
	switch {
	default:
		goto yyrule87
	case c == '.':
		goto yystate37
	case c == '^':
		goto yystate50
	case c >= '0' && c <= '9':
		goto yystate49
	}

yystate50:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '^':
		goto yystate51
	}

yystate51:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate36
	case c >= '0' && c <= '9':
		goto yystate52
	}

yystate52:
	c = lx.Next()
	yyrule = 87
	lx.Mark()
	switch {
	default:
		goto yyrule87
	case c == '.':
		goto yystate37
	case c >= '0' && c <= '9':
		goto yystate52
	}

yystate53:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == ':':
		goto yystate54
	case c == '=':
		goto yystate55
	case c == '>':
		goto yystate56
	}

yystate54:
	c = lx.Next()
	yyrule = 31
	lx.Mark()
	goto yyrule31

yystate55:
	c = lx.Next()
	yyrule = 32
	lx.Mark()
	goto yyrule32

yystate56:
	c = lx.Next()
	yyrule = 33
	lx.Mark()
	goto yyrule33

yystate57:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '<':
		goto yystate58
	case c == '=':
		goto yystate59
	case c == '>':
		goto yystate60
	}

yystate58:
	c = lx.Next()
	yyrule = 34
	lx.Mark()
	goto yyrule34

yystate59:
	c = lx.Next()
	yyrule = 35
	lx.Mark()
	goto yyrule35

yystate60:
	c = lx.Next()
	yyrule = 36
	lx.Mark()
	goto yyrule36

yystate61:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate62
	case c == '=':
		goto yystate64
	}

yystate62:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate63
	}

yystate63:
	c = lx.Next()
	yyrule = 37
	lx.Mark()
	goto yyrule37

yystate64:
	c = lx.Next()
	yyrule = 38
	lx.Mark()
	switch {
	default:
		goto yyrule38
	case c == '=':
		goto yystate65
	}

yystate65:
	c = lx.Next()
	yyrule = 40
	lx.Mark()
	goto yyrule40

yystate66:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate67
	}

yystate67:
	c = lx.Next()
	yyrule = 41
	lx.Mark()
	goto yyrule41

yystate68:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate69
	case c == '@':
		goto yystate70
	}

yystate69:
	c = lx.Next()
	yyrule = 42
	lx.Mark()
	goto yyrule42

yystate70:
	c = lx.Next()
	yyrule = 43
	lx.Mark()
	switch {
	default:
		goto yyrule43
	case c == '@':
		goto yystate71
	}

yystate71:
	c = lx.Next()
	yyrule = 44
	lx.Mark()
	goto yyrule44

yystate72:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '[':
		goto yystate73
	}

yystate73:
	c = lx.Next()
	yyrule = 45
	lx.Mark()
	goto yyrule45

yystate74:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '%':
		goto yystate75
	case c == '&':
		goto yystate76
	case c == '+':
		goto yystate77
	case c == '/':
		goto yystate78
	case c == '@':
		goto yystate79
	case c == '^':
		goto yystate80
	case c == '_':
		goto yystate81
	}

yystate75:
	c = lx.Next()
	yyrule = 46
	lx.Mark()
	goto yyrule46

yystate76:
	c = lx.Next()
	yyrule = 47
	lx.Mark()
	goto yyrule47

yystate77:
	c = lx.Next()
	yyrule = 48
	lx.Mark()
	goto yyrule48

yystate78:
	c = lx.Next()
	yyrule = 49
	lx.Mark()
	goto yyrule49

yystate79:
	c = lx.Next()
	yyrule = 50
	lx.Mark()
	goto yyrule50

yystate80:
	c = lx.Next()
	yyrule = 51
	lx.Mark()
	goto yyrule51

yystate81:
	c = lx.Next()
	yyrule = 52
	lx.Mark()
	goto yyrule52

yystate82:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == ']':
		goto yystate83
	}

yystate83:
	c = lx.Next()
	yyrule = 54
	lx.Mark()
	goto yyrule54

yystate84:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '$' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c == '\u0082' || c == '\u0083':
		goto yystate20
	}

yystate85:
	c = lx.Next()
	switch {
	default:
		goto yyabort
	case c == '|':
		goto yystate86
	}

yystate86:
	c = lx.Next()
	yyrule = 53
	lx.Mark()
	goto yyrule53

yystate87:
	c = lx.Next()
	yyrule = 17
	lx.Mark()
	goto yyrule17

yystate88:
	c = lx.Next()
	yyrule = 60
	lx.Mark()
	goto yyrule60

yystate89:
	c = lx.Next()
	yyrule = 81
	lx.Mark()
	goto yyrule81

yystate90:
	c = lx.Next()
	yyrule = 59
	lx.Mark()
	goto yyrule59

yystate91:
	c = lx.Next()
	yyrule = 70
	lx.Mark()
	goto yyrule70

yystate92:
	c = lx.Next()
	yyrule = 77
	lx.Mark()
	goto yyrule77

yystate93:
	c = lx.Next()
	yyrule = 71
	lx.Mark()
	goto yyrule71

yystate94:
	c = lx.Next()
	yyrule = 66
	lx.Mark()
	goto yyrule66

yystate95:
	c = lx.Next()
	yyrule = 73
	lx.Mark()
	goto yyrule73

yystate96:
	c = lx.Next()
	yyrule = 63
	lx.Mark()
	goto yyrule63

yystate97:
	c = lx.Next()
	yyrule = 68
	lx.Mark()
	goto yyrule68

yystate98:
	c = lx.Next()
	yyrule = 67
	lx.Mark()
	goto yyrule67

yystate99:
	c = lx.Next()
	yyrule = 65
	lx.Mark()
	goto yyrule65

yystate100:
	c = lx.Next()
	yyrule = 78
	lx.Mark()
	goto yyrule78

yystate101:
	c = lx.Next()
	yyrule = 76
	lx.Mark()
	goto yyrule76

yystate102:
	c = lx.Next()
	yyrule = 57
	lx.Mark()
	goto yyrule57

yystate103:
	c = lx.Next()
	yyrule = 62
	lx.Mark()
	goto yyrule62

yystate104:
	c = lx.Next()
	yyrule = 74
	lx.Mark()
	goto yyrule74

yystate105:
	c = lx.Next()
	yyrule = 72
	lx.Mark()
	goto yyrule72

yystate106:
	c = lx.Next()
	yyrule = 69
	lx.Mark()
	goto yyrule69

yystate107:
	c = lx.Next()
	yyrule = 55
	lx.Mark()
	goto yyrule55

yystate108:
	c = lx.Next()
	yyrule = 64
	lx.Mark()
	goto yyrule64

yystate109:
	c = lx.Next()
	yyrule = 84
	lx.Mark()
	goto yyrule84

yystate110:
	c = lx.Next()
	yyrule = 82
	lx.Mark()
	goto yyrule82

yystate111:
	c = lx.Next()
	yyrule = 58
	lx.Mark()
	goto yyrule58

yystate112:
	c = lx.Next()
	yyrule = 56
	lx.Mark()
	goto yyrule56

yystate113:
	c = lx.Next()
	yyrule = 80
	lx.Mark()
	goto yyrule80

yystate114:
	c = lx.Next()
	yyrule = 79
	lx.Mark()
	goto yyrule79

yystate115:
	c = lx.Next()
	yyrule = 75
	lx.Mark()
	goto yyrule75

yystate116:
	c = lx.Next()
	yyrule = 83
	lx.Mark()
	goto yyrule83

yystate117:
	c = lx.Next()
	yyrule = 61
	lx.Mark()
	goto yyrule61

	goto yystate118 // silence unused label error
yystate118:
	c = lx.Next()
yystart118:
	switch {
	default:
		goto yystate120 // c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u0083' || c >= '\u0085' && c <= 'ÿ'
	case c == '"':
		goto yystate121
	case c == '\\':
		goto yystate122
	case c == '\u0084':
		goto yystate132
	case c == '\x00':
		goto yystate119
	}

yystate119:
	c = lx.Next()
	yyrule = 4
	lx.Mark()
	goto yyrule4

yystate120:
	c = lx.Next()
	yyrule = 15
	lx.Mark()
	goto yyrule15

yystate121:
	c = lx.Next()
	yyrule = 16
	lx.Mark()
	goto yyrule16

yystate122:
	c = lx.Next()
	yyrule = 15
	lx.Mark()
	switch {
	default:
		goto yyrule15
	case c == ' ':
		goto yystate124
	case c == '"':
		goto yystate125
	case c == '\\':
		goto yystate126
	case c == '\n':
		goto yystate123
	case c == 'b':
		goto yystate127
	case c == 'f':
		goto yystate128
	case c == 'n':
		goto yystate129
	case c == 'r':
		goto yystate130
	case c == 't':
		goto yystate131
	}

yystate123:
	c = lx.Next()
	yyrule = 13
	lx.Mark()
	goto yyrule13

yystate124:
	c = lx.Next()
	yyrule = 6
	lx.Mark()
	goto yyrule6

yystate125:
	c = lx.Next()
	yyrule = 7
	lx.Mark()
	goto yyrule7

yystate126:
	c = lx.Next()
	yyrule = 5
	lx.Mark()
	goto yyrule5

yystate127:
	c = lx.Next()
	yyrule = 8
	lx.Mark()
	goto yyrule8

yystate128:
	c = lx.Next()
	yyrule = 11
	lx.Mark()
	goto yyrule11

yystate129:
	c = lx.Next()
	yyrule = 10
	lx.Mark()
	goto yyrule10

yystate130:
	c = lx.Next()
	yyrule = 12
	lx.Mark()
	goto yyrule12

yystate131:
	c = lx.Next()
	yyrule = 9
	lx.Mark()
	goto yyrule9

yystate132:
	c = lx.Next()
	yyrule = 14
	lx.Mark()
	goto yyrule14

yyrule1: // [ \t\r]+

	goto yystate0
yyrule2: // "(*"([^*\x80]|\*+[^*)\x80])*\*+\)

	goto yystate0
yyrule3: // \"
	{
		lx.sc = 1
		lx.str = lx.str[:0]
		goto yystate0
	}
yyrule4: // \x00
	{
		lx.str = append(lx.str, 0)
		goto yystate0
	}
yyrule5: // \\\\
	{
		lx.str = append(lx.str, '\\')
		goto yystate0
	}
yyrule6: // \\\x20
	{
		lx.str = append(lx.str, ' ')
		goto yystate0
	}
yyrule7: // \\\"
	{
		lx.str = append(lx.str, '"')
		goto yystate0
	}
yyrule8: // \\b
	{
		lx.str = append(lx.str, '\b')
		goto yystate0
	}
yyrule9: // \\t
	{
		lx.str = append(lx.str, '\t')
		goto yystate0
	}
yyrule10: // \\n
	{
		lx.str = append(lx.str, '\n')
		goto yystate0
	}
yyrule11: // \\f
	{
		lx.str = append(lx.str, '\f')
		goto yystate0
	}
yyrule12: // \\r
	{
		lx.str = append(lx.str, '\r')
		goto yystate0
	}
yyrule13: // \\\n

	goto yystate0
yyrule14: // {ignore}

	goto yystate0
yyrule15: // {schar}
	{

		lx.str = append(lx.str, lx.TokenBytes(nil)...)
		goto yystate0
	}
yyrule16: // \"
	{
		lx.sc = 0
		return STRING
	}
yyrule17: // {ignore}
	{

		return IGNORE
	}
yyrule18: // "&&"
	{
		return AND
	}
yyrule19: // "'"
	{
		return QUOTE
	}
yyrule20: // "**"
	{
		return NON_COMMUTATIVE_MULTIPLY
	}
yyrule21: // "++"
	{
		return INC
	}
yyrule22: // "--"
	{
		return DEC
	}
yyrule23: // "->"
	{
		return RULE
	}
yyrule24: // "/*"
	{
		return RIGHT_COMPOSITION
	}
yyrule25: // "/."
	{
		return REPLACEALL
	}
yyrule26: // "//"
	{
		return POSTFIX
	}
yyrule27: // "//."
	{
		return REPLACEREP
	}
yyrule28: // "//@"
	{
		return MAP_ALL
	}
yyrule29: // "/;"
	{
		return CONDITION
	}
yyrule30: // "/@"
	{
		return MAP
	}
yyrule31: // "::"
	{
		return MESSAGE_NAME
	}
yyrule32: // ":="
	{
		return SET_DELAYED
	}
yyrule33: // ":>"
	{
		return RULEDELAYED
	}
yyrule34: // "<<"
	{
		return GET
	}
yyrule35: // "<="
	{
		return LEQ
	}
yyrule36: // "<>"
	{
		return STRINGJOIN
	}
yyrule37: // "=!="
	{
		return UNSAME
	}
yyrule38: // "=="
	{
		return EQUAL
	}
yyrule39: // "!="
	{
		return UNEQUAL
	}
yyrule40: // "==="
	{
		return SAME
	}
yyrule41: // ">="
	{
		return GEQ
	}
yyrule42: // "@*"
	{
		return COMPOSITION
	}
yyrule43: // "@@"
	{
		return APPLY
	}
yyrule44: // "@@@"
	{
		return APPLY_ALL
	}
yyrule45: // "[["
	{
		return lx.push(LPART)
	}
yyrule46: // "\%"
	{
		return POWER_SUBSCRIPT2
	}
yyrule47: // "\&"
	{
		return OVERSCRIPT
	}
yyrule48: // "\+"
	{
		return UNDERSCRIPT
	}
yyrule49: // "\/"
	{
		return '/'
	}
yyrule50: // "\@"
	{
		return SQRT2
	}
yyrule51: // "\^"
	{
		return POWER_SUBSCRIPT1
	}
yyrule52: // "\_"
	{
		return SUBSCRIPT
	}
yyrule53: // "||"
	{
		return OR
	}
yyrule54: // "]]"
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
yyrule55: // {Backslash}
	{
		return BACKSLASH
	}
yyrule56: // {CenterDot}
	{
		return CENTER_DOT
	}
yyrule57: // {CircleDot}
	{
		return CIRCLE_DOT
	}
yyrule58: // {CircleTimes}
	{
		return CIRCLE_TIMES
	}
yyrule59: // {ConjugateTranspose}
	{
		return CONJUGATE_TRANSPOSE
	}
yyrule60: // {Conjugate}
	{
		return CONJUGATE
	}
yyrule61: // {Coproduct}
	{
		return COPRODUCT
	}
yyrule62: // {Cross}
	{
		return CROSS
	}
yyrule63: // {Del}
	{
		return DEL
	}
yyrule64: // {Diamond}
	{
		return DIAMOND
	}
yyrule65: // {DifferenceDelta}
	{
		return DIFFERENCE_DELTA
	}
yyrule66: // {DifferentialD}
	{
		return DIFFERENTIAL_D
	}
yyrule67: // {DiscreteRatio}
	{
		return DISCRETE_RATIO
	}
yyrule68: // {DiscreteShift}
	{
		return DISCRETE_SHIFT
	}
yyrule69: // {Divide}
	{
		return '/'
	}
yyrule70: // {HermitianConjugate}
	{
		return HERMITIAN_CONJUGATE
	}
yyrule71: // {Integrate}
	{
		return INTEGRATE
	}
yyrule72: // {MinusPlus}
	{
		return MINUS_PLUS
	}
yyrule73: // {PartialD}
	{
		return PARTIAL_D
	}
yyrule74: // {PlusMinus}
	{
		return PLUS_MINUS
	}
yyrule75: // {Product}
	{
		return PRODUCT
	}
yyrule76: // {SmallCircle}
	{
		return SMALL_CIRCLE
	}
yyrule77: // {Sqrt}
	{
		return SQRT
	}
yyrule78: // {Square}
	{
		return SQUARE
	}
yyrule79: // {Star}
	{
		return STAR
	}
yyrule80: // {Times}
	{
		return '*'
	}
yyrule81: // {Transpose}
	{
		return TRANSPOSE
	}
yyrule82: // {Vee}
	{
		return VEE
	}
yyrule83: // {VerticalTilde}
	{
		return VERTICAL_TILDE
	}
yyrule84: // {Wedge}
	{
		return WEDGE
	}
yyrule85: // {float}
	{
		return FLOAT
	}
yyrule86: // {ident}
	{
		return IDENT
	}
yyrule87: // {int}
	{
		return INT
	}
yyrule88: // {pattern}
	{
		return PATTERN
	}
yyrule89: // {slot}
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
