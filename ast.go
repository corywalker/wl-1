// Code generated by yy. DO NOT EDIT.

// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"go/token"
)

// CommaOpt represents data reduced by productions:
//
//	CommaOpt:
//	        /* empty */  // Case 0
//	|       ','          // Case 1
type CommaOpt struct {
	Token Token
}

func (n *CommaOpt) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *CommaOpt) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *CommaOpt) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Token.Pos()
}

// ExprList represents data reduced by productions:
//
//	ExprList:
//	        Expression               // Case 0
//	|       ExprList ',' Expression  // Case 1
type ExprList struct {
	Case       int
	ExprList   *ExprList
	Expression *Expression
	Token      Token
}

func (n *ExprList) reverse() *ExprList {
	if n == nil {
		return nil
	}

	na := n
	nb := na.ExprList
	for nb != nil {
		nc := nb.ExprList
		nb.ExprList = na
		na = nb
		nb = nc
	}
	n.ExprList = nil
	return na
}

func (n *ExprList) fragment() interface{} { return n.reverse() }

// String implements fmt.Stringer.
func (n *ExprList) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *ExprList) Pos() token.Pos {
	if n == nil {
		return 0
	}

	switch n.Case {
	case 1:
		return n.ExprList.Pos()
	case 0:
		return n.Expression.Pos()
	default:
		panic("internal error")
	}
}

// ExpressionCase represents case numbers of production Expression
type ExpressionCase int

// Values of type ExpressionCase
const (
	ExpressionPreInc ExpressionCase = iota
	ExpressionPreDec
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	ExpressionParenExpr
	ExpressionUnaryPlus
	ExpressionUnaryMinus
	_
	_
	ExpressionNe
	ExpressionLAnd
	_
	ExpressionMulAssign
	ExpressionPostInc
	ExpressionAddAssign
	ExpressionPostDec
	ExpressionSubAssign
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	ExpressionLe
	_
	_
	ExpressionEq
	_
	_
	ExpressionRsh
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	ExpressionLOr
	_
	_
	_
	_
	ExpressionMul
	ExpressionAdd
	ExpressionSub
	_
	ExpressionDiv
	_
	_
	_
	ExpressionLt
	_
	ExpressionAssign
	ExpressionGt
	_
	_
	_
	_
	_
	ExpressionOr
	_
	_
	ExpressionFloat
	ExpressionIdent
	_
	_
	ExpressionInteger
	_
	ExpressionPattern
	ExpressionSlot
	ExpressionString
)

// Expression represents data reduced by productions:
//
//	Expression:
//	        "++" Expression                                            // Case ExpressionPreInc
//	|       "--" Expression                                            // Case ExpressionPreDec
//	|       ";;"                                                       // Case 2
//	|       ";;" Expression                                            // Case 3
//	|       "<<" FileName                                              // Case 4
//	|       "\\@" Expression                                           // Case 5
//	|       "\\@" Expression "\\%" Expression                          // Case 6
//	|       "\\[Del]" Expression                                       // Case 7
//	|       "\\[Integrate]" Expression "\\[DifferentialD]" Expression  // Case 8
//	|       "\\[MinusPlus]" Expression                                 // Case 9
//	|       "\\[PlusMinus]" Expression                                 // Case 10
//	|       "\\[Sqrt]" Expression                                      // Case 11
//	|       "\\[Square]" Expression                                    // Case 12
//	|       '!' Expression                                             // Case 13
//	|       '(' Expression ')'                                         // Case ExpressionParenExpr
//	|       '+' Expression                                             // Case ExpressionUnaryPlus
//	|       '-' Expression                                             // Case ExpressionUnaryMinus
//	|       '{' '}'                                                    // Case 17
//	|       '{' ExprList CommaOpt '}'                                  // Case 18
//	|       Expression "!=" Expression                                 // Case ExpressionNe
//	|       Expression "&&" Expression                                 // Case ExpressionLAnd
//	|       Expression "**" Expression                                 // Case 21
//	|       Expression "*=" Expression                                 // Case ExpressionMulAssign
//	|       Expression "++"                                            // Case ExpressionPostInc
//	|       Expression "+=" Expression                                 // Case ExpressionAddAssign
//	|       Expression "--"                                            // Case ExpressionPostDec
//	|       Expression "-=" Expression                                 // Case ExpressionSubAssign
//	|       Expression "->" Expression                                 // Case 27
//	|       Expression ".."                                            // Case 28
//	|       Expression "..."                                           // Case 29
//	|       Expression "/*" Expression                                 // Case 30
//	|       Expression "/." Expression                                 // Case 31
//	|       Expression "//" Expression                                 // Case 32
//	|       Expression "//." Expression                                // Case 33
//	|       Expression "//@" Expression                                // Case 34
//	|       Expression "/:" Expression                                 // Case 35
//	|       Expression "/;" Expression                                 // Case 36
//	|       Expression "/=" Expression                                 // Case 37
//	|       Expression "/@" Expression                                 // Case 38
//	|       Expression ":=" Expression                                 // Case 39
//	|       Expression ":>" Expression                                 // Case 40
//	|       Expression ";;"                                            // Case 41
//	|       Expression ";;" Expression                                 // Case 42
//	|       Expression "<=" Expression                                 // Case ExpressionLe
//	|       Expression "<>" Expression                                 // Case 44
//	|       Expression "=!=" Expression                                // Case 45
//	|       Expression "==" Expression                                 // Case ExpressionEq
//	|       Expression "===" Expression                                // Case 47
//	|       Expression ">=" Expression                                 // Case 48
//	|       Expression ">>" FileName                                   // Case ExpressionRsh
//	|       Expression ">>>" FileName                                  // Case 50
//	|       Expression "@*" Expression                                 // Case 51
//	|       Expression "@@" Expression                                 // Case 52
//	|       Expression "@@@" Expression                                // Case 53
//	|       Expression "[[" ExprList CommaOpt "]]"                     // Case 54
//	|       Expression "\\&" Expression                                // Case 55
//	|       Expression "\\+" Expression                                // Case 56
//	|       Expression "\\[Backslash]" Expression                      // Case 57
//	|       Expression "\\[Because]" Expression                        // Case 58
//	|       Expression "\\[Cap]" Expression                            // Case 59
//	|       Expression "\\[CenterDot]" Expression                      // Case 60
//	|       Expression "\\[CircleDot]" Expression                      // Case 61
//	|       Expression "\\[CircleMinus]" Expression                    // Case 62
//	|       Expression "\\[CirclePlus]" Expression                     // Case 63
//	|       Expression "\\[CircleTimes]" Expression                    // Case 64
//	|       Expression "\\[ConjugateTranspose]"                        // Case 65
//	|       Expression "\\[Conjugate]"                                 // Case 66
//	|       Expression "\\[Coproduct]" Expression                      // Case 67
//	|       Expression "\\[Cross]" Expression                          // Case 68
//	|       Expression "\\[Cup]" Expression                            // Case 69
//	|       Expression "\\[Diamond]" Expression                        // Case 70
//	|       Expression "\\[DifferenceDelta]" Expression                // Case 71
//	|       Expression "\\[DiscreteRatio]" Expression                  // Case 72
//	|       Expression "\\[DiscreteShift]" Expression                  // Case 73
//	|       Expression "\\[DoubleLeftTee]" Expression                  // Case 74
//	|       Expression "\\[DoubleRightTee]" Expression                 // Case 75
//	|       Expression "\\[DoubleVerticalBar]" Expression              // Case 76
//	|       Expression "\\[DownTee]" Expression                        // Case 77
//	|       Expression "\\[Element]" Expression                        // Case 78
//	|       Expression "\\[Equivalent]" Expression                     // Case 79
//	|       Expression "\\[Function]" Expression                       // Case 80
//	|       Expression "\\[HermitianConjugate]"                        // Case 81
//	|       Expression "\\[Implies]" Expression                        // Case 82
//	|       Expression "\\[Intersection]" Expression                   // Case 83
//	|       Expression "\\[LeftTee]" Expression                        // Case 84
//	|       Expression "\\[Nand]" Expression                           // Case 85
//	|       Expression "\\[Nor]" Expression                            // Case 86
//	|       Expression "\\[NotDoubleVerticalBar]" Expression           // Case 87
//	|       Expression "\\[NotElement]" Expression                     // Case 88
//	|       Expression "\\[NotVerticalBar]" Expression                 // Case 89
//	|       Expression "\\[PartialD]" Expression                       // Case 90
//	|       Expression "\\[RightTee]" Expression                       // Case 91
//	|       Expression "\\[SmallCircle]" Expression                    // Case 92
//	|       Expression "\\[Star]" Expression                           // Case 93
//	|       Expression "\\[Subset]" Expression                         // Case 94
//	|       Expression "\\[SuchThat]" Expression                       // Case 95
//	|       Expression "\\[Superset]" Expression                       // Case 96
//	|       Expression "\\[Therefore]" Expression                      // Case 97
//	|       Expression "\\[Transpose]"                                 // Case 98
//	|       Expression "\\[Union]" Expression                          // Case 99
//	|       Expression "\\[UpTee]" Expression                          // Case 100
//	|       Expression "\\[Vee]" Expression                            // Case 101
//	|       Expression "\\[VerticalBar]" Expression                    // Case 102
//	|       Expression "\\[VerticalSeparator]" Expression              // Case 103
//	|       Expression "\\[VerticalTilde]" Expression                  // Case 104
//	|       Expression "\\[Wedge]" Expression                          // Case 105
//	|       Expression "\\[Xnor]" Expression                           // Case 106
//	|       Expression "\\[Xor]" Expression                            // Case 107
//	|       Expression "\\^" Expression "\\%" Expression               // Case 108
//	|       Expression "\\_" Expression                                // Case 109
//	|       Expression "\\`" STRING                                    // Case 110
//	|       Expression "^:=" Expression                                // Case 111
//	|       Expression "^=" Expression                                 // Case 112
//	|       Expression "||" Expression                                 // Case ExpressionLOr
//	|       Expression "~~" Expression                                 // Case 114
//	|       Expression '!'                                             // Case 115
//	|       Expression '!' '!'                                         // Case 116
//	|       Expression '&'                                             // Case 117
//	|       Expression '*' Expression                                  // Case ExpressionMul
//	|       Expression '+' Expression                                  // Case ExpressionAdd
//	|       Expression '-' Expression                                  // Case ExpressionSub
//	|       Expression '.' Expression                                  // Case 121
//	|       Expression '/' Expression                                  // Case ExpressionDiv
//	|       Expression ':' Expression                                  // Case 123
//	|       Expression ';'                                             // Case 124
//	|       Expression ';' Expression                                  // Case 125
//	|       Expression '<' Expression                                  // Case ExpressionLt
//	|       Expression '=' '.'                                         // Case 127
//	|       Expression '=' Expression                                  // Case ExpressionAssign
//	|       Expression '>' Expression                                  // Case ExpressionGt
//	|       Expression '?' Expression                                  // Case 130
//	|       Expression '@' Expression                                  // Case 131
//	|       Expression '[' ']'                                         // Case 132
//	|       Expression '[' ExprList CommaOpt ']'                       // Case 133
//	|       Expression '^' Expression                                  // Case 134
//	|       Expression '|' Expression                                  // Case ExpressionOr
//	|       Expression '~' Expression                                  // Case 136
//	|       Expression QUOTE                                           // Case 137
//	|       FLOAT                                                      // Case ExpressionFloat
//	|       IDENT                                                      // Case ExpressionIdent
//	|       IDENT "::" Tag                                             // Case 140
//	|       IDENT "::" Tag "::" Tag                                    // Case 141
//	|       INT                                                        // Case ExpressionInteger
//	|       OUT                                                        // Case 143
//	|       PATTERN                                                    // Case ExpressionPattern
//	|       SLOT                                                       // Case ExpressionSlot
//	|       STRING                                                     // Case ExpressionString
type Expression struct {
	Case        ExpressionCase
	CommaOpt    *CommaOpt
	ExprList    *ExprList
	Expression  *Expression
	Expression2 *Expression
	Expression3 *Expression
	FileName    *FileName
	Tag         *Tag
	Tag2        *Tag
	Token       Token
	Token2      Token
	Token3      Token
}

func (n *Expression) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Expression) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Expression) Pos() token.Pos {
	if n == nil {
		return 0
	}

	switch n.Case {
	case 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137:
		return n.Expression.Pos()
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 138, 139, 140, 141, 142, 143, 144, 145, 146:
		return n.Token.Pos()
	default:
		panic("internal error")
	}
}

// FileNameCase represents case numbers of production FileName
type FileNameCase int

// Values of type FileNameCase
const (
	FileNameIdent FileNameCase = iota
	FileNameString
)

// FileName represents data reduced by productions:
//
//	FileName:
//	        IDENT   // Case FileNameIdent
//	|       STRING  // Case FileNameString
type FileName struct {
	Case  FileNameCase
	Token Token
}

func (n *FileName) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *FileName) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *FileName) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Token.Pos()
}

// TagCase represents case numbers of production Tag
type TagCase int

// Values of type TagCase
const (
	TagIdent TagCase = iota
	TagString
)

// Tag represents data reduced by productions:
//
//	Tag:
//	        IDENT   // Case TagIdent
//	|       STRING  // Case TagString
type Tag struct {
	Case  TagCase
	Token Token
}

func (n *Tag) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *Tag) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *Tag) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Token.Pos()
}

// start represents data reduced by production:
//
//	start:
//	        Expression  // Case 0
type start struct {
	Expression *Expression
}

func (n *start) fragment() interface{} { return n }

// String implements fmt.Stringer.
func (n *start) String() string {
	return prettyString(n)
}

// Pos reports the position of the first component of n or zero if it's empty.
func (n *start) Pos() token.Pos {
	if n == nil {
		return 0
	}

	return n.Expression.Pos()
}
