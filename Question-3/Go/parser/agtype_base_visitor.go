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

type BaseAgtypeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAgtypeVisitor) VisitAgType(ctx *AgTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitAgValue(ctx *AgValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitStringValue(ctx *StringValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitIntegerValue(ctx *IntegerValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitFloatValue(ctx *FloatValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitTrueBoolean(ctx *TrueBooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitFalseBoolean(ctx *FalseBooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitNullValue(ctx *NullValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitObjectValue(ctx *ObjectValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitArrayValue(ctx *ArrayValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgtypeVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
