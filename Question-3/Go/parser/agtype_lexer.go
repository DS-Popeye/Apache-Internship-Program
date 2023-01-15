/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
// Code generated from Agtype.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 21, 185,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	15, 3, 15, 7, 15, 104, 10, 15, 12, 15, 14, 15, 107, 11, 15, 3, 16, 3, 16,
	3, 16, 7, 16, 112, 10, 16, 12, 16, 14, 16, 115, 11, 16, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 5, 17, 122, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 5, 21, 135, 10, 21, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 7, 22, 142, 10, 22, 12, 22, 14, 22, 145, 11, 22,
	5, 22, 147, 10, 22, 3, 23, 5, 23, 150, 10, 23, 3, 23, 3, 23, 3, 23, 3,
	24, 5, 24, 156, 10, 24, 3, 24, 3, 24, 5, 24, 160, 10, 24, 3, 24, 3, 24,
	3, 25, 3, 25, 6, 25, 166, 10, 25, 13, 25, 14, 25, 167, 3, 26, 3, 26, 5,
	26, 172, 10, 26, 3, 26, 6, 26, 175, 10, 26, 13, 26, 14, 26, 176, 3, 27,
	6, 27, 180, 10, 27, 13, 27, 14, 27, 181, 3, 27, 3, 27, 2, 2, 28, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 2, 35, 2, 37, 2, 39, 2, 41, 18, 43,
	2, 45, 19, 47, 20, 49, 2, 51, 2, 53, 21, 3, 2, 12, 5, 2, 67, 92, 97, 97,
	99, 124, 7, 2, 38, 38, 50, 59, 67, 92, 97, 97, 99, 124, 10, 2, 36, 36,
	49, 49, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 5, 2,
	50, 59, 67, 72, 99, 104, 5, 2, 2, 33, 36, 36, 94, 94, 3, 2, 51, 59, 3,
	2, 50, 59, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 5, 2, 11, 12,
	15, 15, 34, 34, 2, 191, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23,
	3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2,
	31, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2,
	2, 53, 3, 2, 2, 2, 3, 55, 3, 2, 2, 2, 5, 60, 3, 2, 2, 2, 7, 66, 3, 2, 2,
	2, 9, 71, 3, 2, 2, 2, 11, 73, 3, 2, 2, 2, 13, 75, 3, 2, 2, 2, 15, 77, 3,
	2, 2, 2, 17, 79, 3, 2, 2, 2, 19, 81, 3, 2, 2, 2, 21, 83, 3, 2, 2, 2, 23,
	86, 3, 2, 2, 2, 25, 88, 3, 2, 2, 2, 27, 97, 3, 2, 2, 2, 29, 101, 3, 2,
	2, 2, 31, 108, 3, 2, 2, 2, 33, 118, 3, 2, 2, 2, 35, 123, 3, 2, 2, 2, 37,
	129, 3, 2, 2, 2, 39, 131, 3, 2, 2, 2, 41, 134, 3, 2, 2, 2, 43, 146, 3,
	2, 2, 2, 45, 149, 3, 2, 2, 2, 47, 155, 3, 2, 2, 2, 49, 163, 3, 2, 2, 2,
	51, 169, 3, 2, 2, 2, 53, 179, 3, 2, 2, 2, 55, 56, 7, 118, 2, 2, 56, 57,
	7, 116, 2, 2, 57, 58, 7, 119, 2, 2, 58, 59, 7, 103, 2, 2, 59, 4, 3, 2,
	2, 2, 60, 61, 7, 104, 2, 2, 61, 62, 7, 99, 2, 2, 62, 63, 7, 110, 2, 2,
	63, 64, 7, 117, 2, 2, 64, 65, 7, 103, 2, 2, 65, 6, 3, 2, 2, 2, 66, 67,
	7, 112, 2, 2, 67, 68, 7, 119, 2, 2, 68, 69, 7, 110, 2, 2, 69, 70, 7, 110,
	2, 2, 70, 8, 3, 2, 2, 2, 71, 72, 7, 125, 2, 2, 72, 10, 3, 2, 2, 2, 73,
	74, 7, 46, 2, 2, 74, 12, 3, 2, 2, 2, 75, 76, 7, 127, 2, 2, 76, 14, 3, 2,
	2, 2, 77, 78, 7, 60, 2, 2, 78, 16, 3, 2, 2, 2, 79, 80, 7, 93, 2, 2, 80,
	18, 3, 2, 2, 2, 81, 82, 7, 95, 2, 2, 82, 20, 3, 2, 2, 2, 83, 84, 7, 60,
	2, 2, 84, 85, 7, 60, 2, 2, 85, 22, 3, 2, 2, 2, 86, 87, 7, 47, 2, 2, 87,
	24, 3, 2, 2, 2, 88, 89, 7, 75, 2, 2, 89, 90, 7, 112, 2, 2, 90, 91, 7, 104,
	2, 2, 91, 92, 7, 107, 2, 2, 92, 93, 7, 112, 2, 2, 93, 94, 7, 107, 2, 2,
	94, 95, 7, 118, 2, 2, 95, 96, 7, 123, 2, 2, 96, 26, 3, 2, 2, 2, 97, 98,
	7, 80, 2, 2, 98, 99, 7, 99, 2, 2, 99, 100, 7, 80, 2, 2, 100, 28, 3, 2,
	2, 2, 101, 105, 9, 2, 2, 2, 102, 104, 9, 3, 2, 2, 103, 102, 3, 2, 2, 2,
	104, 107, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106,
	30, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 108, 113, 7, 36, 2, 2, 109, 112,
	5, 33, 17, 2, 110, 112, 5, 39, 20, 2, 111, 109, 3, 2, 2, 2, 111, 110, 3,
	2, 2, 2, 112, 115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2,
	2, 114, 116, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 116, 117, 7, 36, 2, 2, 117,
	32, 3, 2, 2, 2, 118, 121, 7, 94, 2, 2, 119, 122, 9, 4, 2, 2, 120, 122,
	5, 35, 18, 2, 121, 119, 3, 2, 2, 2, 121, 120, 3, 2, 2, 2, 122, 34, 3, 2,
	2, 2, 123, 124, 7, 119, 2, 2, 124, 125, 5, 37, 19, 2, 125, 126, 5, 37,
	19, 2, 126, 127, 5, 37, 19, 2, 127, 128, 5, 37, 19, 2, 128, 36, 3, 2, 2,
	2, 129, 130, 9, 5, 2, 2, 130, 38, 3, 2, 2, 2, 131, 132, 10, 6, 2, 2, 132,
	40, 3, 2, 2, 2, 133, 135, 7, 47, 2, 2, 134, 133, 3, 2, 2, 2, 134, 135,
	3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 137, 5, 43, 22, 2, 137, 42, 3, 2,
	2, 2, 138, 147, 7, 50, 2, 2, 139, 143, 9, 7, 2, 2, 140, 142, 9, 8, 2, 2,
	141, 140, 3, 2, 2, 2, 142, 145, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 143,
	144, 3, 2, 2, 2, 144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 146, 138,
	3, 2, 2, 2, 146, 139, 3, 2, 2, 2, 147, 44, 3, 2, 2, 2, 148, 150, 7, 47,
	2, 2, 149, 148, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2,
	151, 152, 5, 43, 22, 2, 152, 153, 5, 49, 25, 2, 153, 46, 3, 2, 2, 2, 154,
	156, 7, 47, 2, 2, 155, 154, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 157,
	3, 2, 2, 2, 157, 159, 5, 43, 22, 2, 158, 160, 5, 49, 25, 2, 159, 158, 3,
	2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 162, 5, 51, 26,
	2, 162, 48, 3, 2, 2, 2, 163, 165, 7, 48, 2, 2, 164, 166, 9, 8, 2, 2, 165,
	164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 168,
	3, 2, 2, 2, 168, 50, 3, 2, 2, 2, 169, 171, 9, 9, 2, 2, 170, 172, 9, 10,
	2, 2, 171, 170, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 174, 3, 2, 2, 2,
	173, 175, 9, 8, 2, 2, 174, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176,
	174, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 52, 3, 2, 2, 2, 178, 180, 9,
	11, 2, 2, 179, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 179, 3, 2, 2,
	2, 181, 182, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 184, 8, 27, 2, 2, 184,
	54, 3, 2, 2, 2, 17, 2, 105, 111, 113, 121, 134, 143, 146, 149, 155, 159,
	167, 171, 176, 181, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'true'", "'false'", "'null'", "'{'", "','", "'}'", "':'", "'['", "']'",
	"'::'", "'-'", "'Infinity'", "'NaN'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "IDENT", "STRING",
	"INTEGER", "RegularFloat", "ExponentFloat", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "IDENT", "STRING", "ESC", "UNICODE",
	"HEX", "SAFECODEPOINT", "INTEGER", "INT", "RegularFloat", "ExponentFloat",
	"DECIMAL", "SCIENTIFIC", "WS",
}

type AgtypeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewAgtypeLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *AgtypeLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewAgtypeLexer(input antlr.CharStream) *AgtypeLexer {
	l := new(AgtypeLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Agtype.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AgtypeLexer tokens.
const (
	AgtypeLexerT__0          = 1
	AgtypeLexerT__1          = 2
	AgtypeLexerT__2          = 3
	AgtypeLexerT__3          = 4
	AgtypeLexerT__4          = 5
	AgtypeLexerT__5          = 6
	AgtypeLexerT__6          = 7
	AgtypeLexerT__7          = 8
	AgtypeLexerT__8          = 9
	AgtypeLexerT__9          = 10
	AgtypeLexerT__10         = 11
	AgtypeLexerT__11         = 12
	AgtypeLexerT__12         = 13
	AgtypeLexerIDENT         = 14
	AgtypeLexerSTRING        = 15
	AgtypeLexerINTEGER       = 16
	AgtypeLexerRegularFloat  = 17
	AgtypeLexerExponentFloat = 18
	AgtypeLexerWS            = 19
)
