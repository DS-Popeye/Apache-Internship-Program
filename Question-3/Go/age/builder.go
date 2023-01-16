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
package age

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/DS-Popeye/Apache-Internship-Program/Question-3/Go/parser"

)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinUint = 0
const MinInt = -MaxInt - 1

type Unmarshaller interface {
	unmarshal(text string) (Entity, error)
}

type AGUnmarshaler struct {
	Unmarshaller
	ageParser   *parser.AgtypeParser
	visitor     parser.AgtypeVisitor
	errListener *AGErrorListener
	vcache      map[int64]interface{}
}

func NewAGUnmarshaler() *AGUnmarshaler {
	vcache := make(map[int64]interface{})

	m := &AGUnmarshaler{ageParser: parser.NewAgtypeParser(nil),
		visitor:     &UnmarshalVisitor{vcache: vcache},
		errListener: NewAGErrorListener(),
		vcache:      vcache,
	}
	m.ageParser.AddErrorListener(m.errListener)

	return m
}

func (p *AGUnmarshaler) unmarshal(text string) (Entity, error) {
	if len(text) == 0 {
		return NewSimpleEntity(nil), nil
	}
	input := antlr.NewInputStream(text)
	lexer := parser.NewAgtypeLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p.ageParser.SetInputStream(stream)
	tree := p.ageParser.AgType()
	rst := tree.Accept(p.visitor)

	if len(p.errListener.errList) > 0 {
		var ape *AgeParseError = nil
		errs := make([]string, len(p.errListener.errList))
		for idx, re := range p.errListener.errList {
			errs[idx] = re.GetMessage()
			fmt.Println(re)
		}
		p.errListener.ClearErrs()

		ape = &AgeParseError{msg: "Cypher query:" + text, errors: errs}

		return nil, ape
	}

	if !IsEntity(rst) {
		rst = NewSimpleEntity(rst)
	}

	return rst.(Entity), nil
}

type AGErrorListener struct {
	*antlr.DefaultErrorListener
	errList []antlr.RecognitionException
}

func NewAGErrorListener() *AGErrorListener {
	return &AGErrorListener{DefaultErrorListener: &antlr.DefaultErrorListener{}, errList: []antlr.RecognitionException{}}
}

func (el *AGErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.errList = append(el.errList, e)
}

func (el *AGErrorListener) GetErrs() []antlr.RecognitionException {
	return el.errList
}
func (el *AGErrorListener) ClearErrs() {
	el.errList = []antlr.RecognitionException{}
}

type UnmarshalVisitor struct {
	parser.AgtypeVisitor
	vcache map[int64]interface{}
}

func (v *UnmarshalVisitor) Visit(tree antlr.ParseTree) interface{} { return nil }

func (v *UnmarshalVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	var rtn interface{}
	for _, c := range node.GetChildren() {
		pt := c.(antlr.ParseTree)
		rtn = pt.Accept(v)
	}
	return rtn
}

func (v *UnmarshalVisitor) VisitTerminal(node antlr.TerminalNode) interface{} { return nil }
func (v *UnmarshalVisitor) VisitErrorNode(node antlr.ErrorNode) interface{}   { return nil }
func (v *UnmarshalVisitor) VisitAgType(ctx *parser.AgTypeContext) interface{} {
	agv := ctx.AgValue()
	if agv != nil {
		return agv.Accept(v)
	}
	return nil
}

func (v *UnmarshalVisitor) VisitAgValue(ctx *parser.AgValueContext) interface{} {
	annoCtx := ctx.TypeAnnotation()
	valueCtx := ctx.Value()

	if annoCtx != nil {
		annoStr := annoCtx.(*parser.TypeAnnotationContext).IDENT().GetText()
		value, err := v.handleAnnotatedValue(annoStr, valueCtx)
		if err != nil {
			panic(err)
		}

		return value
	} else {
		return valueCtx.Accept(v)
	}
}

// Visit a parse tree produced by AgtypeParser#value.
func (v *UnmarshalVisitor) VisitValue(ctx *parser.ValueContext) interface{} {
	rtn := v.VisitChildren(ctx)
	return rtn
}

func (v *UnmarshalVisitor) handleAnnotatedValue(anno string, ctx antlr.ParserRuleContext) (interface{}, error) {
	if anno == "numeric" {
		numStr := ctx.GetText()
		if strings.Contains(numStr, ".") {
			bi := new(big.Float)
			bi, ok := bi.SetString(numStr)
			if !ok {
				return nil, &AgeParseError{msg: "Parse big float " + numStr}
			}
			return bi, nil
		} else {
			bi := new(big.Int)
			bi, ok := bi.SetString(numStr, 10)
			if !ok {
				return nil, &AgeParseError{msg: "Parse big int " + numStr}
			}
			return bi, nil
		}
	} else if anno == "vertex" {
		dict := ctx.Accept(v).(map[string]interface{})
		vid, ok := dict["id"].(int64)
		if !ok {
			return nil, &AgeParseError{msg: "Vertex id is nil. "}
		}
		vertex, ok := v.vcache[vid].(*Vertex)

		if !ok {
			vertex = NewVertex(vid, dict["label"].(string), dict["properties"].(map[string]interface{}))
			v.vcache[vid] = vertex
		}

		return vertex, nil
	} else if anno == "edge" {
		dict := ctx.Accept(v).(map[string]interface{})
		edge := NewEdge(int64(dict["id"].(int64)), dict["label"].(string),
			int64(dict["start_id"].(int64)), int64(dict["end_id"].(int64)),
			dict["properties"].(map[string]interface{}))

		return edge, nil
	} else if anno == "path" {
		arr := ctx.Accept(v).([]interface{})
		entities := []Entity{}

		for _, child := range arr {
			switch child.(type) {
			case *Vertex:
				entities = append(entities, child.(Entity))
			case *Edge:
				entities = append(entities, child.(Entity))
			default:
			}
		}

		path := NewPath(entities)
		return path, nil
	}
	return ctx.Accept(v), nil
}

func (v *UnmarshalVisitor) VisitStringValue(ctx *parser.StringValueContext) interface{} {
	txt := ctx.GetText()
	return strings.Trim(txt, "\"")
}

func (v *UnmarshalVisitor) VisitIntegerValue(ctx *parser.IntegerValueContext) interface{} {
	val, err := strconv.ParseInt(ctx.INTEGER().GetText(), 10, 64)
	if err != nil {
		panic(err)
	}
	return val
}

// Visit a parse tree produced by AgtypeParser#FloatValue.
func (v *UnmarshalVisitor) VisitFloatValue(ctx *parser.FloatValueContext) interface{} {
	fCtx := ctx.GetChild(0).(*parser.FloatLiteralContext)
	return fCtx.Accept(v)
}

func (v *UnmarshalVisitor) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {
	c := ctx.GetChild(0)
	text := ctx.GetText()
	tokenType := c.(antlr.TerminalNode).GetSymbol().GetTokenType()

	switch tokenType {
	case parser.AgtypeParserRegularFloat:
		val, err := strconv.ParseFloat(text, 64)
		if err != nil {
			panic(err)
		}
		return val
	case parser.AgtypeParserExponentFloat:
		val, err := strconv.ParseFloat(text, 64)
		if err != nil {
			panic(err)
		}
		return val
	default:
		switch text {
		case "NaN":
			return math.NaN()
		case "-Infinity":
			return math.Inf(-1)
		case "Infinity":
			return math.Inf(1)
		default:
			panic(&AgeParseError{msg: "Unknown float expression:" + text})
		}
	}
}

func (v *UnmarshalVisitor) VisitTrueBoolean(ctx *parser.TrueBooleanContext) interface{} {
	return true
}

func (v *UnmarshalVisitor) VisitFalseBoolean(ctx *parser.FalseBooleanContext) interface{} {
	return false
}

func (v *UnmarshalVisitor) VisitNullValue(ctx *parser.NullValueContext) interface{} {
	return nil
}

// Visit a parse tree produced by AgtypeParser#ObjectValue.
func (v *UnmarshalVisitor) VisitObjectValue(ctx *parser.ObjectValueContext) interface{} {
	objCtx := ctx.GetChild(0).(*parser.ObjContext)
	return objCtx.Accept(v)
}

func (v *UnmarshalVisitor) VisitObj(ctx *parser.ObjContext) interface{} {
	props := make(map[string]interface{})
	for _, pairCtx := range ctx.AllPair() {
		pairCtx.Accept(v)
		pair := pairCtx.(*parser.PairContext)
		key := strings.Trim(pair.STRING().GetText(), "\"")
		value := pair.AgValue().Accept(v)
		props[key] = value
	}

	return props
}

func (v *UnmarshalVisitor) VisitPair(ctx *parser.PairContext) interface{} {
	return nil
}

func (v *UnmarshalVisitor) VisitArrayValue(ctx *parser.ArrayValueContext) interface{} {
	arrCtx := ctx.GetChild(0).(*parser.ArrayContext)
	return arrCtx.Accept(v)
}

// Visit a parse tree produced by AgtypeParser#array.
func (v *UnmarshalVisitor) VisitArray(ctx *parser.ArrayContext) interface{} {
	var arr []interface{}
	for _, child := range ctx.GetChildren() {
		switch child.(type) {
		case *antlr.TerminalNodeImpl:
			// skip
			break
		default:
			el := child.(antlr.ParserRuleContext).Accept(v)
			arr = append(arr, el)
		}
	}
	return arr
}

/*

   # Visit a parse tree produced by AgtypeParser#TrueBoolean.
   def visitTrueBoolean(self, ctx:AgtypeParser.TrueBooleanContext):
       return True


   # Visit a parse tree produced by AgtypeParser#FalseBoolean.
   def visitFalseBoolean(self, ctx:AgtypeParser.FalseBooleanContext):
       return False


   # Visit a parse tree produced by AgtypeParser#NullValue.
   def visitNullValue(self, ctx:AgtypeParser.NullValueContext):
       return None


   # Visit a parse tree produced by AgtypeParser#obj.
   def visitObj(self, ctx:AgtypeParser.ObjContext):
       obj = dict()
       for c in ctx.getChildren():
           if isinstance(c, AgtypeParser.PairContext):
               namVal = self.visitPair(c)
               name = namVal[0]
               valCtx = namVal[1]
               val = valCtx.accept(self)
               obj[name] = val
       return obj


   # Visit a parse tree produced by AgtypeParser#pair.
   def visitPair(self, ctx:AgtypeParser.PairContext):
       self.visitChildren(ctx)
       return (ctx.STRING().getText().strip('"') , ctx.agValue())


   # Visit a parse tree produced by AgtypeParser#array.
   def visitArray(self, ctx:AgtypeParser.ArrayContext):
       li = list()
       for c in ctx.getChildren():
           if not isinstance(c, TerminalNode):
               val = c.accept(self)
               li.append(val)
       return li

*/
