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

// A complete Visitor for a parse tree produced by AgtypeParser.
type AgtypeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AgtypeParser#agType.
	VisitAgType(ctx *AgTypeContext) interface{}

	// Visit a parse tree produced by AgtypeParser#agValue.
	VisitAgValue(ctx *AgValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#StringValue.
	VisitStringValue(ctx *StringValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#IntegerValue.
	VisitIntegerValue(ctx *IntegerValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#FloatValue.
	VisitFloatValue(ctx *FloatValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#TrueBoolean.
	VisitTrueBoolean(ctx *TrueBooleanContext) interface{}

	// Visit a parse tree produced by AgtypeParser#FalseBoolean.
	VisitFalseBoolean(ctx *FalseBooleanContext) interface{}

	// Visit a parse tree produced by AgtypeParser#NullValue.
	VisitNullValue(ctx *NullValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#ObjectValue.
	VisitObjectValue(ctx *ObjectValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#ArrayValue.
	VisitArrayValue(ctx *ArrayValueContext) interface{}

	// Visit a parse tree produced by AgtypeParser#obj.
	VisitObj(ctx *ObjContext) interface{}

	// Visit a parse tree produced by AgtypeParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by AgtypeParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by AgtypeParser#typeAnnotation.
	VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{}

	// Visit a parse tree produced by AgtypeParser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}
}
