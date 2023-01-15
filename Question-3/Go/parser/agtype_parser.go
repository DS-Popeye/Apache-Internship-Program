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

package parser // Agtype

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 21, 82, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 24, 10, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 34, 10, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 7, 5, 40, 10, 5, 12, 5, 14, 5, 43, 11, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 5, 5, 49, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 7, 7, 59, 10, 7, 12, 7, 14, 7, 62, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5,
	7, 68, 10, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 76, 10, 9, 3, 9,
	3, 9, 5, 9, 80, 10, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 2,
	2, 89, 2, 18, 3, 2, 2, 2, 4, 21, 3, 2, 2, 2, 6, 33, 3, 2, 2, 2, 8, 48,
	3, 2, 2, 2, 10, 50, 3, 2, 2, 2, 12, 67, 3, 2, 2, 2, 14, 69, 3, 2, 2, 2,
	16, 79, 3, 2, 2, 2, 18, 19, 5, 4, 3, 2, 19, 20, 7, 2, 2, 3, 20, 3, 3, 2,
	2, 2, 21, 23, 5, 6, 4, 2, 22, 24, 5, 14, 8, 2, 23, 22, 3, 2, 2, 2, 23,
	24, 3, 2, 2, 2, 24, 5, 3, 2, 2, 2, 25, 34, 7, 17, 2, 2, 26, 34, 7, 18,
	2, 2, 27, 34, 5, 16, 9, 2, 28, 34, 7, 3, 2, 2, 29, 34, 7, 4, 2, 2, 30,
	34, 7, 5, 2, 2, 31, 34, 5, 8, 5, 2, 32, 34, 5, 12, 7, 2, 33, 25, 3, 2,
	2, 2, 33, 26, 3, 2, 2, 2, 33, 27, 3, 2, 2, 2, 33, 28, 3, 2, 2, 2, 33, 29,
	3, 2, 2, 2, 33, 30, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 32, 3, 2, 2, 2,
	34, 7, 3, 2, 2, 2, 35, 36, 7, 6, 2, 2, 36, 41, 5, 10, 6, 2, 37, 38, 7,
	7, 2, 2, 38, 40, 5, 10, 6, 2, 39, 37, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41,
	39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 44, 3, 2, 2, 2, 43, 41, 3, 2, 2,
	2, 44, 45, 7, 8, 2, 2, 45, 49, 3, 2, 2, 2, 46, 47, 7, 6, 2, 2, 47, 49,
	7, 8, 2, 2, 48, 35, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 9, 3, 2, 2, 2,
	50, 51, 7, 17, 2, 2, 51, 52, 7, 9, 2, 2, 52, 53, 5, 4, 3, 2, 53, 11, 3,
	2, 2, 2, 54, 55, 7, 10, 2, 2, 55, 60, 5, 4, 3, 2, 56, 57, 7, 7, 2, 2, 57,
	59, 5, 4, 3, 2, 58, 56, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2,
	2, 60, 61, 3, 2, 2, 2, 61, 63, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 64,
	7, 11, 2, 2, 64, 68, 3, 2, 2, 2, 65, 66, 7, 10, 2, 2, 66, 68, 7, 11, 2,
	2, 67, 54, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 68, 13, 3, 2, 2, 2, 69, 70,
	7, 12, 2, 2, 70, 71, 7, 16, 2, 2, 71, 15, 3, 2, 2, 2, 72, 80, 7, 19, 2,
	2, 73, 80, 7, 20, 2, 2, 74, 76, 7, 13, 2, 2, 75, 74, 3, 2, 2, 2, 75, 76,
	3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 80, 7, 14, 2, 2, 78, 80, 7, 15, 2,
	2, 79, 72, 3, 2, 2, 2, 79, 73, 3, 2, 2, 2, 79, 75, 3, 2, 2, 2, 79, 78,
	3, 2, 2, 2, 80, 17, 3, 2, 2, 2, 10, 23, 33, 41, 48, 60, 67, 75, 79,
}
var literalNames = []string{
	"", "'true'", "'false'", "'null'", "'{'", "','", "'}'", "':'", "'['", "']'",
	"'::'", "'-'", "'Infinity'", "'NaN'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "IDENT", "STRING",
	"INTEGER", "RegularFloat", "ExponentFloat", "WS",
}

var ruleNames = []string{
	"agType", "agValue", "value", "obj", "pair", "array", "typeAnnotation",
	"floatLiteral",
}

type AgtypeParser struct {
	*antlr.BaseParser
}

// NewAgtypeParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *AgtypeParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewAgtypeParser(input antlr.TokenStream) *AgtypeParser {
	this := new(AgtypeParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Agtype.g4"

	return this
}

// AgtypeParser tokens.
const (
	AgtypeParserEOF           = antlr.TokenEOF
	AgtypeParserT__0          = 1
	AgtypeParserT__1          = 2
	AgtypeParserT__2          = 3
	AgtypeParserT__3          = 4
	AgtypeParserT__4          = 5
	AgtypeParserT__5          = 6
	AgtypeParserT__6          = 7
	AgtypeParserT__7          = 8
	AgtypeParserT__8          = 9
	AgtypeParserT__9          = 10
	AgtypeParserT__10         = 11
	AgtypeParserT__11         = 12
	AgtypeParserT__12         = 13
	AgtypeParserIDENT         = 14
	AgtypeParserSTRING        = 15
	AgtypeParserINTEGER       = 16
	AgtypeParserRegularFloat  = 17
	AgtypeParserExponentFloat = 18
	AgtypeParserWS            = 19
)

// AgtypeParser rules.
const (
	AgtypeParserRULE_agType         = 0
	AgtypeParserRULE_agValue        = 1
	AgtypeParserRULE_value          = 2
	AgtypeParserRULE_obj            = 3
	AgtypeParserRULE_pair           = 4
	AgtypeParserRULE_array          = 5
	AgtypeParserRULE_typeAnnotation = 6
	AgtypeParserRULE_floatLiteral   = 7
)

// IAgTypeContext is an interface to support dynamic dispatch.
type IAgTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAgTypeContext differentiates from other interfaces.
	IsAgTypeContext()
}

type AgTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAgTypeContext() *AgTypeContext {
	var p = new(AgTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgtypeParserRULE_agType
	return p
}

func (*AgTypeContext) IsAgTypeContext() {}

func NewAgTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AgTypeContext {
	var p = new(AgTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_agType

	return p
}

func (s *AgTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *AgTypeContext) AgValue() IAgValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAgValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAgValueContext)
}

func (s *AgTypeContext) EOF() antlr.TerminalNode {
	return s.GetToken(AgtypeParserEOF, 0)
}

func (s *AgTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AgTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AgTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterAgType(s)
	}
}

func (s *AgTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitAgType(s)
	}
}

func (s *AgTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitAgType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) AgType() (localctx IAgTypeContext) {
	localctx = NewAgTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AgtypeParserRULE_agType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.AgValue()
	}
	{
		p.SetState(17)
		p.Match(AgtypeParserEOF)
	}

	return localctx
}

// IAgValueContext is an interface to support dynamic dispatch.
type IAgValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAgValueContext differentiates from other interfaces.
	IsAgValueContext()
}

type AgValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAgValueContext() *AgValueContext {
	var p = new(AgValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgtypeParserRULE_agValue
	return p
}

func (*AgValueContext) IsAgValueContext() {}

func NewAgValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AgValueContext {
	var p = new(AgValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_agValue

	return p
}

func (s *AgValueContext) GetParser() antlr.Parser { return s.parser }

func (s *AgValueContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AgValueContext) TypeAnnotation() ITypeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeAnnotationContext)
}

func (s *AgValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AgValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AgValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterAgValue(s)
	}
}

func (s *AgValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitAgValue(s)
	}
}

func (s *AgValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitAgValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) AgValue() (localctx IAgValueContext) {
	localctx = NewAgValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AgtypeParserRULE_agValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Value()
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AgtypeParserT__9 {
		{
			p.SetState(20)
			p.TypeAnnotation()
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgtypeParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullValueContext struct {
	*ValueContext
}

func NewNullValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullValueContext {
	var p = new(NullValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NullValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterNullValue(s)
	}
}

func (s *NullValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitNullValue(s)
	}
}

func (s *NullValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitNullValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ObjectValueContext struct {
	*ValueContext
}

func NewObjectValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectValueContext {
	var p = new(ObjectValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ObjectValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectValueContext) Obj() IObjContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *ObjectValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterObjectValue(s)
	}
}

func (s *ObjectValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitObjectValue(s)
	}
}

func (s *ObjectValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitObjectValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegerValueContext struct {
	*ValueContext
}

func NewIntegerValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerValueContext {
	var p = new(IntegerValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *IntegerValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerValueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(AgtypeParserINTEGER, 0)
}

func (s *IntegerValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterIntegerValue(s)
	}
}

func (s *IntegerValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitIntegerValue(s)
	}
}

func (s *IntegerValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitIntegerValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueBooleanContext struct {
	*ValueContext
}

func NewTrueBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueBooleanContext {
	var p = new(TrueBooleanContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *TrueBooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueBooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterTrueBoolean(s)
	}
}

func (s *TrueBooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitTrueBoolean(s)
	}
}

func (s *TrueBooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitTrueBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseBooleanContext struct {
	*ValueContext
}

func NewFalseBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseBooleanContext {
	var p = new(FalseBooleanContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *FalseBooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseBooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterFalseBoolean(s)
	}
}

func (s *FalseBooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitFalseBoolean(s)
	}
}

func (s *FalseBooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitFalseBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatValueContext struct {
	*ValueContext
}

func NewFloatValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatValueContext {
	var p = new(FloatValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *FloatValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatValueContext) FloatLiteral() IFloatLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *FloatValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterFloatValue(s)
	}
}

func (s *FloatValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitFloatValue(s)
	}
}

func (s *FloatValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitFloatValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringValueContext struct {
	*ValueContext
}

func NewStringValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringValueContext {
	var p = new(StringValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(AgtypeParserSTRING, 0)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (s *StringValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitStringValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayValueContext struct {
	*ValueContext
}

func NewArrayValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayValueContext {
	var p = new(ArrayValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterArrayValue(s)
	}
}

func (s *ArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitArrayValue(s)
	}
}

func (s *ArrayValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitArrayValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AgtypeParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(31)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AgtypeParserSTRING:
		localctx = NewStringValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Match(AgtypeParserSTRING)
		}

	case AgtypeParserINTEGER:
		localctx = NewIntegerValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.Match(AgtypeParserINTEGER)
		}

	case AgtypeParserT__10, AgtypeParserT__11, AgtypeParserT__12, AgtypeParserRegularFloat, AgtypeParserExponentFloat:
		localctx = NewFloatValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(25)
			p.FloatLiteral()
		}

	case AgtypeParserT__0:
		localctx = NewTrueBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(26)
			p.Match(AgtypeParserT__0)
		}

	case AgtypeParserT__1:
		localctx = NewFalseBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(27)
			p.Match(AgtypeParserT__1)
		}

	case AgtypeParserT__2:
		localctx = NewNullValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(28)
			p.Match(AgtypeParserT__2)
		}

	case AgtypeParserT__3:
		localctx = NewObjectValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(29)
			p.Obj()
		}

	case AgtypeParserT__7:
		localctx = NewArrayValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(30)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgtypeParserRULE_obj
	return p
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *ObjContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterObj(s)
	}
}

func (s *ObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitObj(s)
	}
}

func (s *ObjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitObj(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) Obj() (localctx IObjContext) {
	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AgtypeParserRULE_obj)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Match(AgtypeParserT__3)
		}
		{
			p.SetState(34)
			p.Pair()
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AgtypeParserT__4 {
			{
				p.SetState(35)
				p.Match(AgtypeParserT__4)
			}
			{
				p.SetState(36)
				p.Pair()
			}

			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(42)
			p.Match(AgtypeParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Match(AgtypeParserT__3)
		}
		{
			p.SetState(45)
			p.Match(AgtypeParserT__5)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgtypeParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(AgtypeParserSTRING, 0)
}

func (s *PairContext) AgValue() IAgValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAgValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAgValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitPair(s)
	}
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AgtypeParserRULE_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(AgtypeParserSTRING)
	}
	{
		p.SetState(49)
		p.Match(AgtypeParserT__6)
	}
	{
		p.SetState(50)
		p.AgValue()
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgtypeParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllAgValue() []IAgValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAgValueContext)(nil)).Elem())
	var tst = make([]IAgValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAgValueContext)
		}
	}

	return tst
}

func (s *ArrayContext) AgValue(i int) IAgValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAgValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAgValueContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AgtypeParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(AgtypeParserT__7)
		}
		{
			p.SetState(53)
			p.AgValue()
		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AgtypeParserT__4 {
			{
				p.SetState(54)
				p.Match(AgtypeParserT__4)
			}
			{
				p.SetState(55)
				p.AgValue()
			}

			p.SetState(60)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(61)
			p.Match(AgtypeParserT__8)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Match(AgtypeParserT__7)
		}
		{
			p.SetState(64)
			p.Match(AgtypeParserT__8)
		}

	}

	return localctx
}

// ITypeAnnotationContext is an interface to support dynamic dispatch.
type ITypeAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeAnnotationContext differentiates from other interfaces.
	IsTypeAnnotationContext()
}

type TypeAnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeAnnotationContext() *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgtypeParserRULE_typeAnnotation
	return p
}

func (*TypeAnnotationContext) IsTypeAnnotationContext() {}

func NewTypeAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAnnotationContext {
	var p = new(TypeAnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_typeAnnotation

	return p
}

func (s *TypeAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAnnotationContext) IDENT() antlr.TerminalNode {
	return s.GetToken(AgtypeParserIDENT, 0)
}

func (s *TypeAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitTypeAnnotation(s)
	}
}

func (s *TypeAnnotationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitTypeAnnotation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) TypeAnnotation() (localctx ITypeAnnotationContext) {
	localctx = NewTypeAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AgtypeParserRULE_typeAnnotation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(AgtypeParserT__9)
	}
	{
		p.SetState(68)
		p.Match(AgtypeParserIDENT)
	}

	return localctx
}

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgtypeParserRULE_floatLiteral
	return p
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgtypeParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) RegularFloat() antlr.TerminalNode {
	return s.GetToken(AgtypeParserRegularFloat, 0)
}

func (s *FloatLiteralContext) ExponentFloat() antlr.TerminalNode {
	return s.GetToken(AgtypeParserExponentFloat, 0)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgtypeListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgtypeVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgtypeParser) FloatLiteral() (localctx IFloatLiteralContext) {
	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AgtypeParserRULE_floatLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AgtypeParserRegularFloat:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(AgtypeParserRegularFloat)
		}

	case AgtypeParserExponentFloat:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(AgtypeParserExponentFloat)
		}

	case AgtypeParserT__10, AgtypeParserT__11:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AgtypeParserT__10 {
			{
				p.SetState(72)
				p.Match(AgtypeParserT__10)
			}

		}
		{
			p.SetState(75)
			p.Match(AgtypeParserT__11)
		}

	case AgtypeParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(76)
			p.Match(AgtypeParserT__12)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
