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

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAgtypeListener is a complete listener for a parse tree produced by AgtypeParser.
type BaseAgtypeListener struct{}

var _ AgtypeListener = &BaseAgtypeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAgtypeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAgtypeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAgtypeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAgtypeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAgType is called when production agType is entered.
func (s *BaseAgtypeListener) EnterAgType(ctx *AgTypeContext) {}

// ExitAgType is called when production agType is exited.
func (s *BaseAgtypeListener) ExitAgType(ctx *AgTypeContext) {}

// EnterAgValue is called when production agValue is entered.
func (s *BaseAgtypeListener) EnterAgValue(ctx *AgValueContext) {}

// ExitAgValue is called when production agValue is exited.
func (s *BaseAgtypeListener) ExitAgValue(ctx *AgValueContext) {}

// EnterStringValue is called when production StringValue is entered.
func (s *BaseAgtypeListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production StringValue is exited.
func (s *BaseAgtypeListener) ExitStringValue(ctx *StringValueContext) {}

// EnterIntegerValue is called when production IntegerValue is entered.
func (s *BaseAgtypeListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production IntegerValue is exited.
func (s *BaseAgtypeListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterFloatValue is called when production FloatValue is entered.
func (s *BaseAgtypeListener) EnterFloatValue(ctx *FloatValueContext) {}

// ExitFloatValue is called when production FloatValue is exited.
func (s *BaseAgtypeListener) ExitFloatValue(ctx *FloatValueContext) {}

// EnterTrueBoolean is called when production TrueBoolean is entered.
func (s *BaseAgtypeListener) EnterTrueBoolean(ctx *TrueBooleanContext) {}

// ExitTrueBoolean is called when production TrueBoolean is exited.
func (s *BaseAgtypeListener) ExitTrueBoolean(ctx *TrueBooleanContext) {}

// EnterFalseBoolean is called when production FalseBoolean is entered.
func (s *BaseAgtypeListener) EnterFalseBoolean(ctx *FalseBooleanContext) {}

// ExitFalseBoolean is called when production FalseBoolean is exited.
func (s *BaseAgtypeListener) ExitFalseBoolean(ctx *FalseBooleanContext) {}

// EnterNullValue is called when production NullValue is entered.
func (s *BaseAgtypeListener) EnterNullValue(ctx *NullValueContext) {}

// ExitNullValue is called when production NullValue is exited.
func (s *BaseAgtypeListener) ExitNullValue(ctx *NullValueContext) {}

// EnterObjectValue is called when production ObjectValue is entered.
func (s *BaseAgtypeListener) EnterObjectValue(ctx *ObjectValueContext) {}

// ExitObjectValue is called when production ObjectValue is exited.
func (s *BaseAgtypeListener) ExitObjectValue(ctx *ObjectValueContext) {}

// EnterArrayValue is called when production ArrayValue is entered.
func (s *BaseAgtypeListener) EnterArrayValue(ctx *ArrayValueContext) {}

// ExitArrayValue is called when production ArrayValue is exited.
func (s *BaseAgtypeListener) ExitArrayValue(ctx *ArrayValueContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseAgtypeListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseAgtypeListener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseAgtypeListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseAgtypeListener) ExitPair(ctx *PairContext) {}

// EnterArray is called when production array is entered.
func (s *BaseAgtypeListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseAgtypeListener) ExitArray(ctx *ArrayContext) {}

// EnterTypeAnnotation is called when production typeAnnotation is entered.
func (s *BaseAgtypeListener) EnterTypeAnnotation(ctx *TypeAnnotationContext) {}

// ExitTypeAnnotation is called when production typeAnnotation is exited.
func (s *BaseAgtypeListener) ExitTypeAnnotation(ctx *TypeAnnotationContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseAgtypeListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseAgtypeListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}
