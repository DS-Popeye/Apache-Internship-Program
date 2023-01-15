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

// AgtypeListener is a complete listener for a parse tree produced by AgtypeParser.
type AgtypeListener interface {
	antlr.ParseTreeListener

	// EnterAgType is called when entering the agType production.
	EnterAgType(c *AgTypeContext)

	// EnterAgValue is called when entering the agValue production.
	EnterAgValue(c *AgValueContext)

	// EnterStringValue is called when entering the StringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterIntegerValue is called when entering the IntegerValue production.
	EnterIntegerValue(c *IntegerValueContext)

	// EnterFloatValue is called when entering the FloatValue production.
	EnterFloatValue(c *FloatValueContext)

	// EnterTrueBoolean is called when entering the TrueBoolean production.
	EnterTrueBoolean(c *TrueBooleanContext)

	// EnterFalseBoolean is called when entering the FalseBoolean production.
	EnterFalseBoolean(c *FalseBooleanContext)

	// EnterNullValue is called when entering the NullValue production.
	EnterNullValue(c *NullValueContext)

	// EnterObjectValue is called when entering the ObjectValue production.
	EnterObjectValue(c *ObjectValueContext)

	// EnterArrayValue is called when entering the ArrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterTypeAnnotation is called when entering the typeAnnotation production.
	EnterTypeAnnotation(c *TypeAnnotationContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// ExitAgType is called when exiting the agType production.
	ExitAgType(c *AgTypeContext)

	// ExitAgValue is called when exiting the agValue production.
	ExitAgValue(c *AgValueContext)

	// ExitStringValue is called when exiting the StringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitIntegerValue is called when exiting the IntegerValue production.
	ExitIntegerValue(c *IntegerValueContext)

	// ExitFloatValue is called when exiting the FloatValue production.
	ExitFloatValue(c *FloatValueContext)

	// ExitTrueBoolean is called when exiting the TrueBoolean production.
	ExitTrueBoolean(c *TrueBooleanContext)

	// ExitFalseBoolean is called when exiting the FalseBoolean production.
	ExitFalseBoolean(c *FalseBooleanContext)

	// ExitNullValue is called when exiting the NullValue production.
	ExitNullValue(c *NullValueContext)

	// ExitObjectValue is called when exiting the ObjectValue production.
	ExitObjectValue(c *ObjectValueContext)

	// ExitArrayValue is called when exiting the ArrayValue production.
	ExitArrayValue(c *ArrayValueContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitTypeAnnotation is called when exiting the typeAnnotation production.
	ExitTypeAnnotation(c *TypeAnnotationContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)
}
