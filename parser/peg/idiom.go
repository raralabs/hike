
package peg

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"os"

	"strconv"
	"strings"

	"unicode"
	"unicode/utf8"

	"github.com/Knetic/govaluate"
	"github.com/raralabs/canal/utils/cast"
)

var g = &grammar {
	rules: []*rule{
{
	name: "Main",
	pos: position{line: 5, col: 1, offset: 21},
	expr: &actionExpr{
	pos: position{line: 5, col: 9, offset: 29},
	run: (*parser).callonMain1,
	expr: &seqExpr{
	pos: position{line: 5, col: 9, offset: 29},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 5, col: 9, offset: 29},
	label: "agg",
	expr: &ruleRefExpr{
	pos: position{line: 5, col: 13, offset: 33},
	name: "Job",
},
},
&labeledExpr{
	pos: position{line: 5, col: 17, offset: 37},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 5, col: 22, offset: 42},
	expr: &seqExpr{
	pos: position{line: 5, col: 23, offset: 43},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 5, col: 23, offset: 43},
	name: "_",
},
&litMatcher{
	pos: position{line: 5, col: 25, offset: 45},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 5, col: 29, offset: 49},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 5, col: 31, offset: 51},
	name: "Job",
},
	},
},
},
},
	},
},
},
},
{
	name: "MExpr",
	pos: position{line: 16, col: 1, offset: 242},
	expr: &actionExpr{
	pos: position{line: 16, col: 10, offset: 251},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 16, col: 10, offset: 251},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 16, col: 10, offset: 251},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 16, col: 16, offset: 257},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 16, col: 23, offset: 264},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 16, col: 28, offset: 269},
	expr: &seqExpr{
	pos: position{line: 16, col: 29, offset: 270},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 16, col: 29, offset: 270},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 16, col: 31, offset: 272},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 16, col: 36, offset: 277},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 16, col: 38, offset: 279},
	name: "MValue",
},
	},
},
},
},
	},
},
},
},
{
	name: "MValue",
	pos: position{line: 20, col: 1, offset: 344},
	expr: &actionExpr{
	pos: position{line: 20, col: 11, offset: 354},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 20, col: 11, offset: 354},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 20, col: 16, offset: 359},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 20, col: 16, offset: 359},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 20, col: 24, offset: 367},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 20, col: 33, offset: 376},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 20, col: 48, offset: 391},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 20, col: 59, offset: 402},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 24, col: 1, offset: 444},
	expr: &choiceExpr{
	pos: position{line: 24, col: 9, offset: 452},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 24, col: 9, offset: 452},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 24, col: 15, offset: 458},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 24, col: 21, offset: 464},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 24, col: 27, offset: 470},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 24, col: 33, offset: 476},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 24, col: 33, offset: 476},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 28, col: 1, offset: 524},
	expr: &choiceExpr{
	pos: position{line: 28, col: 12, offset: 535},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 28, col: 12, offset: 535},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 28, col: 12, offset: 535},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 28, col: 12, offset: 535},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 28, col: 16, offset: 539},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 28, col: 19, offset: 542},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 28, col: 25, offset: 548},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 32, col: 5, offset: 624},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 32, col: 5, offset: 624},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 32, col: 8, offset: 627},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 39, col: 1, offset: 683},
	expr: &actionExpr{
	pos: position{line: 39, col: 9, offset: 691},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 39, col: 9, offset: 691},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 39, col: 9, offset: 691},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 39, col: 15, offset: 697},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 39, col: 21, offset: 703},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 39, col: 26, offset: 708},
	expr: &seqExpr{
	pos: position{line: 39, col: 27, offset: 709},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 39, col: 27, offset: 709},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 39, col: 29, offset: 711},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 39, col: 34, offset: 716},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 39, col: 36, offset: 718},
	name: "Value",
},
	},
},
},
},
	},
},
},
},
{
	name: "Value",
	pos: position{line: 43, col: 1, offset: 782},
	expr: &actionExpr{
	pos: position{line: 43, col: 10, offset: 791},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 43, col: 10, offset: 791},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 43, col: 15, offset: 796},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 43, col: 15, offset: 796},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 43, col: 23, offset: 804},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 43, col: 32, offset: 813},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 43, col: 47, offset: 828},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 47, col: 1, offset: 871},
	expr: &actionExpr{
	pos: position{line: 47, col: 13, offset: 883},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 47, col: 13, offset: 883},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 47, col: 13, offset: 883},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 47, col: 20, offset: 890},
	expr: &choiceExpr{
	pos: position{line: 47, col: 21, offset: 891},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 47, col: 21, offset: 891},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 47, col: 31, offset: 901},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 47, col: 39, offset: 909},
	val: "_",
	ignoreCase: false,
},
	},
},
},
	},
},
},
},
{
	name: "String_Type1",
	pos: position{line: 51, col: 1, offset: 959},
	expr: &actionExpr{
	pos: position{line: 51, col: 17, offset: 975},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 51, col: 17, offset: 975},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 51, col: 17, offset: 975},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 51, col: 21, offset: 979},
	expr: &choiceExpr{
	pos: position{line: 51, col: 23, offset: 981},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 51, col: 23, offset: 981},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 51, col: 23, offset: 981},
	expr: &ruleRefExpr{
	pos: position{line: 51, col: 24, offset: 982},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 51, col: 36, offset: 994,
},
	},
},
&seqExpr{
	pos: position{line: 51, col: 40, offset: 998},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 51, col: 40, offset: 998},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 51, col: 45, offset: 1003},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 51, col: 63, offset: 1021},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 58, col: 1, offset: 1202},
	expr: &charClassMatcher{
	pos: position{line: 58, col: 16, offset: 1217},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 60, col: 1, offset: 1235},
	expr: &ruleRefExpr{
	pos: position{line: 60, col: 19, offset: 1253},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 62, col: 1, offset: 1273},
	expr: &charClassMatcher{
	pos: position{line: 62, col: 21, offset: 1293},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 64, col: 1, offset: 1308},
	expr: &actionExpr{
	pos: position{line: 64, col: 9, offset: 1316},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 64, col: 9, offset: 1316},
	expr: &choiceExpr{
	pos: position{line: 64, col: 10, offset: 1317},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 64, col: 10, offset: 1317},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 64, col: 19, offset: 1326},
	val: "_",
	ignoreCase: false,
},
	},
},
},
},
},
{
	name: "Letter",
	pos: position{line: 68, col: 1, offset: 1376},
	expr: &choiceExpr{
	pos: position{line: 68, col: 11, offset: 1386},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 68, col: 11, offset: 1386},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 68, col: 19, offset: 1394},
	val: "[A-Z]",
	ranges: []rune{'A','Z',},
	ignoreCase: false,
	inverted: false,
},
	},
},
},
{
	name: "Number",
	pos: position{line: 70, col: 1, offset: 1403},
	expr: &actionExpr{
	pos: position{line: 70, col: 11, offset: 1413},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 70, col: 11, offset: 1413},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 70, col: 11, offset: 1413},
	expr: &litMatcher{
	pos: position{line: 70, col: 11, offset: 1413},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 70, col: 16, offset: 1418},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 74, col: 1, offset: 1491},
	expr: &actionExpr{
	pos: position{line: 74, col: 19, offset: 1509},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 74, col: 19, offset: 1509},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 78, col: 1, offset: 1582},
	expr: &actionExpr{
	pos: position{line: 78, col: 10, offset: 1591},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 78, col: 10, offset: 1591},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 78, col: 10, offset: 1591},
	expr: &litMatcher{
	pos: position{line: 78, col: 10, offset: 1591},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 78, col: 15, offset: 1596},
	expr: &ruleRefExpr{
	pos: position{line: 78, col: 15, offset: 1596},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 78, col: 24, offset: 1605},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 78, col: 28, offset: 1609},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 82, col: 1, offset: 1676},
	expr: &choiceExpr{
	pos: position{line: 82, col: 12, offset: 1687},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 82, col: 12, offset: 1687},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 82, col: 18, offset: 1693},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 82, col: 18, offset: 1693},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 82, col: 38, offset: 1713},
	expr: &ruleRefExpr{
	pos: position{line: 82, col: 38, offset: 1713},
	name: "DecimalDigit",
},
},
	},
},
	},
},
},
{
	name: "DecimalDigit",
	pos: position{line: 84, col: 1, offset: 1730},
	expr: &charClassMatcher{
	pos: position{line: 84, col: 17, offset: 1746},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 86, col: 1, offset: 1755},
	expr: &charClassMatcher{
	pos: position{line: 86, col: 24, offset: 1778},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 88, col: 1, offset: 1787},
	expr: &zeroOrMoreExpr{
	pos: position{line: 88, col: 19, offset: 1805},
	expr: &charClassMatcher{
	pos: position{line: 88, col: 19, offset: 1805},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 90, col: 1, offset: 1819},
	expr: &charClassMatcher{
	pos: position{line: 90, col: 18, offset: 1836},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EOF",
	pos: position{line: 92, col: 1, offset: 1849},
	expr: &notExpr{
	pos: position{line: 92, col: 8, offset: 1856},
	expr: &anyMatcher{
	line: 92, col: 9, offset: 1857,
},
},
},
{
	name: "Job",
	pos: position{line: 96, col: 1, offset: 1873},
	expr: &choiceExpr{
	pos: position{line: 96, col: 8, offset: 1880},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 96, col: 8, offset: 1880},
	name: "FileJobG",
},
&ruleRefExpr{
	pos: position{line: 96, col: 19, offset: 1891},
	name: "FakeJobG",
},
&ruleRefExpr{
	pos: position{line: 96, col: 30, offset: 1902},
	name: "DoJobG",
},
&ruleRefExpr{
	pos: position{line: 96, col: 39, offset: 1911},
	name: "MapJobG",
},
&ruleRefExpr{
	pos: position{line: 96, col: 49, offset: 1921},
	name: "AggJobG",
},
&ruleRefExpr{
	pos: position{line: 96, col: 59, offset: 1931},
	name: "SinkJobG",
},
	},
},
},
{
	name: "FileJobG",
	pos: position{line: 98, col: 1, offset: 1943},
	expr: &actionExpr{
	pos: position{line: 98, col: 13, offset: 1955},
	run: (*parser).callonFileJobG1,
	expr: &seqExpr{
	pos: position{line: 98, col: 13, offset: 1955},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 98, col: 13, offset: 1955},
	name: "_",
},
&litMatcher{
	pos: position{line: 98, col: 15, offset: 1957},
	val: "file(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 98, col: 23, offset: 1965},
	name: "_",
},
&labeledExpr{
	pos: position{line: 98, col: 25, offset: 1967},
	label: "filename",
	expr: &ruleRefExpr{
	pos: position{line: 98, col: 34, offset: 1976},
	name: "FileName",
},
},
&ruleRefExpr{
	pos: position{line: 98, col: 43, offset: 1985},
	name: "_",
},
&litMatcher{
	pos: position{line: 98, col: 45, offset: 1987},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FakeJobG",
	pos: position{line: 105, col: 1, offset: 2169},
	expr: &actionExpr{
	pos: position{line: 105, col: 13, offset: 2181},
	run: (*parser).callonFakeJobG1,
	expr: &seqExpr{
	pos: position{line: 105, col: 13, offset: 2181},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 105, col: 13, offset: 2181},
	name: "_",
},
&litMatcher{
	pos: position{line: 105, col: 15, offset: 2183},
	val: "fake(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 105, col: 23, offset: 2191},
	name: "_",
},
&labeledExpr{
	pos: position{line: 105, col: 25, offset: 2193},
	label: "numData",
	expr: &ruleRefExpr{
	pos: position{line: 105, col: 33, offset: 2201},
	name: "PositiveNumber",
},
},
&ruleRefExpr{
	pos: position{line: 105, col: 48, offset: 2216},
	name: "_",
},
&litMatcher{
	pos: position{line: 105, col: 50, offset: 2218},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FileName",
	pos: position{line: 112, col: 1, offset: 2387},
	expr: &actionExpr{
	pos: position{line: 112, col: 13, offset: 2399},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 112, col: 13, offset: 2399},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 112, col: 13, offset: 2399},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 112, col: 22, offset: 2408},
	expr: &seqExpr{
	pos: position{line: 112, col: 23, offset: 2409},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 112, col: 23, offset: 2409},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 112, col: 27, offset: 2413},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 112, col: 31, offset: 2417},
	name: "Variable",
},
},
	},
},
},
	},
},
},
},
{
	name: "DoJobG",
	pos: position{line: 118, col: 1, offset: 2503},
	expr: &actionExpr{
	pos: position{line: 118, col: 11, offset: 2513},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 118, col: 11, offset: 2513},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 118, col: 16, offset: 2518},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 118, col: 16, offset: 2518},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 118, col: 27, offset: 2529},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 118, col: 38, offset: 2540},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 118, col: 49, offset: 2551},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 118, col: 58, offset: 2560},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 118, col: 67, offset: 2569},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "MapJobG",
	pos: position{line: 122, col: 1, offset: 2632},
	expr: &actionExpr{
	pos: position{line: 122, col: 12, offset: 2643},
	run: (*parser).callonMapJobG1,
	expr: &seqExpr{
	pos: position{line: 122, col: 12, offset: 2643},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 122, col: 12, offset: 2643},
	val: "map",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 122, col: 18, offset: 2649},
	name: "_",
},
&labeledExpr{
	pos: position{line: 122, col: 20, offset: 2651},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 122, col: 25, offset: 2656},
	name: "EnrichDo",
},
},
	},
},
},
},
{
	name: "AggJobG",
	pos: position{line: 126, col: 1, offset: 2720},
	expr: &actionExpr{
	pos: position{line: 126, col: 12, offset: 2731},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 126, col: 12, offset: 2731},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 126, col: 12, offset: 2731},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 126, col: 18, offset: 2737},
	name: "_",
},
&labeledExpr{
	pos: position{line: 126, col: 20, offset: 2739},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 126, col: 27, offset: 2746},
	expr: &ruleRefExpr{
	pos: position{line: 126, col: 27, offset: 2746},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 126, col: 39, offset: 2758},
	name: "_",
},
&labeledExpr{
	pos: position{line: 126, col: 41, offset: 2760},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 126, col: 45, offset: 2764},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 126, col: 53, offset: 2772},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 126, col: 58, offset: 2777},
	expr: &seqExpr{
	pos: position{line: 126, col: 59, offset: 2778},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 126, col: 59, offset: 2778},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 126, col: 63, offset: 2782},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 126, col: 65, offset: 2784},
	name: "AggJobs",
},
	},
},
},
},
	},
},
},
},
{
	name: "AggJobs",
	pos: position{line: 140, col: 1, offset: 3124},
	expr: &choiceExpr{
	pos: position{line: 140, col: 12, offset: 3135},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 140, col: 12, offset: 3135},
	name: "CountAgg",
},
&ruleRefExpr{
	pos: position{line: 140, col: 23, offset: 3146},
	name: "MaxAgg",
},
&ruleRefExpr{
	pos: position{line: 140, col: 32, offset: 3155},
	name: "MinAgg",
},
&ruleRefExpr{
	pos: position{line: 140, col: 41, offset: 3164},
	name: "AvgAgg",
},
&ruleRefExpr{
	pos: position{line: 140, col: 50, offset: 3173},
	name: "SumAgg",
},
&ruleRefExpr{
	pos: position{line: 140, col: 59, offset: 3182},
	name: "ModeAgg",
},
&ruleRefExpr{
	pos: position{line: 141, col: 14, offset: 3204},
	name: "VarAgg",
},
&ruleRefExpr{
	pos: position{line: 141, col: 23, offset: 3213},
	name: "DistinctCountAgg",
},
&ruleRefExpr{
	pos: position{line: 141, col: 42, offset: 3232},
	name: "QuantileAgg",
},
&ruleRefExpr{
	pos: position{line: 141, col: 56, offset: 3246},
	name: "MedianAgg",
},
&ruleRefExpr{
	pos: position{line: 141, col: 68, offset: 3258},
	name: "QuartileAgg",
},
&ruleRefExpr{
	pos: position{line: 142, col: 14, offset: 3284},
	name: "CovAgg",
},
&ruleRefExpr{
	pos: position{line: 142, col: 23, offset: 3293},
	name: "CorrAgg",
},
&ruleRefExpr{
	pos: position{line: 142, col: 33, offset: 3303},
	name: "PValueAgg",
},
	},
},
},
{
	name: "AggGroupBy",
	pos: position{line: 144, col: 1, offset: 3316},
	expr: &actionExpr{
	pos: position{line: 144, col: 15, offset: 3330},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 144, col: 15, offset: 3330},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 144, col: 15, offset: 3330},
	name: "_",
},
&litMatcher{
	pos: position{line: 144, col: 17, offset: 3332},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 144, col: 22, offset: 3337},
	name: "_",
},
&labeledExpr{
	pos: position{line: 144, col: 24, offset: 3339},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 144, col: 30, offset: 3345},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 144, col: 39, offset: 3354},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 144, col: 44, offset: 3359},
	expr: &seqExpr{
	pos: position{line: 144, col: 45, offset: 3360},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 144, col: 45, offset: 3360},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 144, col: 49, offset: 3364},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 144, col: 51, offset: 3366},
	name: "Variable",
},
	},
},
},
},
	},
},
},
},
{
	name: "SinkJobG",
	pos: position{line: 158, col: 1, offset: 3686},
	expr: &choiceExpr{
	pos: position{line: 158, col: 13, offset: 3698},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 158, col: 13, offset: 3698},
	name: "StdOutG",
},
&ruleRefExpr{
	pos: position{line: 158, col: 23, offset: 3708},
	name: "BlackHoleG",
},
&ruleRefExpr{
	pos: position{line: 158, col: 36, offset: 3721},
	name: "PlotG",
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 160, col: 1, offset: 3730},
	expr: &actionExpr{
	pos: position{line: 160, col: 12, offset: 3741},
	run: (*parser).callonStdOutG1,
	expr: &seqExpr{
	pos: position{line: 160, col: 12, offset: 3741},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 160, col: 12, offset: 3741},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 160, col: 22, offset: 3751},
	name: "_",
},
&labeledExpr{
	pos: position{line: 160, col: 24, offset: 3753},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 160, col: 31, offset: 3760},
	expr: &ruleRefExpr{
	pos: position{line: 160, col: 32, offset: 3761},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 160, col: 48, offset: 3777},
	name: "_",
},
&litMatcher{
	pos: position{line: 160, col: 50, offset: 3779},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BlackHoleG",
	pos: position{line: 174, col: 1, offset: 4025},
	expr: &actionExpr{
	pos: position{line: 174, col: 15, offset: 4039},
	run: (*parser).callonBlackHoleG1,
	expr: &seqExpr{
	pos: position{line: 174, col: 15, offset: 4039},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 174, col: 15, offset: 4039},
	val: "blackhole(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 174, col: 28, offset: 4052},
	name: "_",
},
&labeledExpr{
	pos: position{line: 174, col: 30, offset: 4054},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 174, col: 37, offset: 4061},
	expr: &ruleRefExpr{
	pos: position{line: 174, col: 38, offset: 4062},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 174, col: 54, offset: 4078},
	name: "_",
},
&litMatcher{
	pos: position{line: 174, col: 56, offset: 4080},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PlotG",
	pos: position{line: 178, col: 1, offset: 4140},
	expr: &actionExpr{
	pos: position{line: 178, col: 10, offset: 4149},
	run: (*parser).callonPlotG1,
	expr: &seqExpr{
	pos: position{line: 178, col: 10, offset: 4149},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 178, col: 10, offset: 4149},
	val: "plot",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 178, col: 17, offset: 4156},
	name: "_",
},
&labeledExpr{
	pos: position{line: 178, col: 19, offset: 4158},
	label: "widget",
	expr: &ruleRefExpr{
	pos: position{line: 178, col: 27, offset: 4166},
	name: "BarWidget",
},
},
	},
},
},
},
{
	name: "BarWidget",
	pos: position{line: 182, col: 1, offset: 4238},
	expr: &actionExpr{
	pos: position{line: 182, col: 14, offset: 4251},
	run: (*parser).callonBarWidget1,
	expr: &seqExpr{
	pos: position{line: 182, col: 14, offset: 4251},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 182, col: 14, offset: 4251},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 182, col: 20, offset: 4257},
	expr: &ruleRefExpr{
	pos: position{line: 182, col: 20, offset: 4257},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 182, col: 30, offset: 4267},
	name: "_",
},
&litMatcher{
	pos: position{line: 182, col: 32, offset: 4269},
	val: "bar(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 182, col: 39, offset: 4276},
	name: "_",
},
&labeledExpr{
	pos: position{line: 182, col: 41, offset: 4278},
	label: "xField",
	expr: &ruleRefExpr{
	pos: position{line: 182, col: 48, offset: 4285},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 182, col: 57, offset: 4294},
	name: "_",
},
&litMatcher{
	pos: position{line: 182, col: 59, offset: 4296},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 182, col: 63, offset: 4300},
	name: "_",
},
&labeledExpr{
	pos: position{line: 182, col: 65, offset: 4302},
	label: "yField",
	expr: &ruleRefExpr{
	pos: position{line: 182, col: 72, offset: 4309},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 182, col: 81, offset: 4318},
	name: "_",
},
&labeledExpr{
	pos: position{line: 182, col: 83, offset: 4320},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 182, col: 88, offset: 4325},
	expr: &seqExpr{
	pos: position{line: 182, col: 89, offset: 4326},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 182, col: 89, offset: 4326},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 182, col: 93, offset: 4330},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 182, col: 95, offset: 4332},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 182, col: 102, offset: 4339},
	name: "_",
},
&zeroOrOneExpr{
	pos: position{line: 182, col: 104, offset: 4341},
	expr: &seqExpr{
	pos: position{line: 182, col: 105, offset: 4342},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 182, col: 105, offset: 4342},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 182, col: 109, offset: 4346},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 182, col: 111, offset: 4348},
	name: "Number",
},
	},
},
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 182, col: 122, offset: 4359},
	name: "_",
},
&litMatcher{
	pos: position{line: 182, col: 124, offset: 4361},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 215, col: 1, offset: 5048},
	expr: &actionExpr{
	pos: position{line: 215, col: 13, offset: 5060},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 215, col: 13, offset: 5060},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 215, col: 13, offset: 5060},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 215, col: 23, offset: 5070},
	name: "_",
},
&labeledExpr{
	pos: position{line: 215, col: 25, offset: 5072},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 215, col: 30, offset: 5077},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 215, col: 42, offset: 5089},
	name: "_",
},
&litMatcher{
	pos: position{line: 215, col: 44, offset: 5091},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 224, col: 1, offset: 5377},
	expr: &actionExpr{
	pos: position{line: 224, col: 16, offset: 5392},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 224, col: 16, offset: 5392},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 224, col: 16, offset: 5392},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 224, col: 22, offset: 5398},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 224, col: 35, offset: 5411},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 224, col: 40, offset: 5416},
	expr: &seqExpr{
	pos: position{line: 224, col: 41, offset: 5417},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 224, col: 41, offset: 5417},
	name: "_",
},
&labeledExpr{
	pos: position{line: 224, col: 43, offset: 5419},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 224, col: 46, offset: 5422},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 224, col: 58, offset: 5434},
	name: "_",
},
&labeledExpr{
	pos: position{line: 224, col: 60, offset: 5436},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 224, col: 63, offset: 5439},
	name: "FilterFactor",
},
},
	},
},
},
},
	},
},
},
},
{
	name: "FilterFactor",
	pos: position{line: 229, col: 1, offset: 5597},
	expr: &choiceExpr{
	pos: position{line: 229, col: 17, offset: 5613},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 229, col: 17, offset: 5613},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 229, col: 17, offset: 5613},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 229, col: 17, offset: 5613},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 229, col: 21, offset: 5617},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 229, col: 24, offset: 5620},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 229, col: 36, offset: 5632},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 231, col: 5, offset: 5663},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 231, col: 5, offset: 5663},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 231, col: 8, offset: 5666},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 235, col: 1, offset: 5708},
	expr: &choiceExpr{
	pos: position{line: 235, col: 16, offset: 5723},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 235, col: 16, offset: 5723},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 235, col: 24, offset: 5731},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 235, col: 24, offset: 5731},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 242, col: 1, offset: 5910},
	expr: &actionExpr{
	pos: position{line: 242, col: 18, offset: 5927},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 242, col: 18, offset: 5927},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 242, col: 18, offset: 5927},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 242, col: 23, offset: 5932},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 242, col: 28, offset: 5937},
	name: "_",
},
&labeledExpr{
	pos: position{line: 242, col: 30, offset: 5939},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 242, col: 33, offset: 5942},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 242, col: 45, offset: 5954},
	name: "_",
},
&labeledExpr{
	pos: position{line: 242, col: 47, offset: 5956},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 242, col: 53, offset: 5962},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 242, col: 59, offset: 5968},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 246, col: 1, offset: 6042},
	expr: &choiceExpr{
	pos: position{line: 246, col: 16, offset: 6057},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 246, col: 16, offset: 6057},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 246, col: 23, offset: 6064},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 246, col: 30, offset: 6071},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 246, col: 36, offset: 6077},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 246, col: 42, offset: 6083},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 246, col: 49, offset: 6090},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 246, col: 49, offset: 6090},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "BranchDo",
	pos: position{line: 253, col: 1, offset: 6162},
	expr: &actionExpr{
	pos: position{line: 253, col: 13, offset: 6174},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 253, col: 13, offset: 6174},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 253, col: 13, offset: 6174},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 253, col: 23, offset: 6184},
	name: "_",
},
&labeledExpr{
	pos: position{line: 253, col: 25, offset: 6186},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 253, col: 28, offset: 6189},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 253, col: 42, offset: 6203},
	name: "_",
},
&litMatcher{
	pos: position{line: 253, col: 44, offset: 6205},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 260, col: 1, offset: 6400},
	expr: &actionExpr{
	pos: position{line: 260, col: 18, offset: 6417},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 260, col: 18, offset: 6417},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 260, col: 18, offset: 6417},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 260, col: 23, offset: 6422},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 260, col: 28, offset: 6427},
	name: "_",
},
&labeledExpr{
	pos: position{line: 260, col: 30, offset: 6429},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 260, col: 33, offset: 6432},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 260, col: 45, offset: 6444},
	name: "_",
},
&labeledExpr{
	pos: position{line: 260, col: 47, offset: 6446},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 260, col: 54, offset: 6453},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 264, col: 1, offset: 6504},
	expr: &ruleRefExpr{
	pos: position{line: 264, col: 16, offset: 6519},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 269, col: 1, offset: 6557},
	expr: &actionExpr{
	pos: position{line: 269, col: 13, offset: 6569},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 269, col: 13, offset: 6569},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 269, col: 13, offset: 6569},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 269, col: 23, offset: 6579},
	name: "_",
},
&labeledExpr{
	pos: position{line: 269, col: 25, offset: 6581},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 269, col: 31, offset: 6587},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 269, col: 40, offset: 6596},
	name: "_",
},
&labeledExpr{
	pos: position{line: 269, col: 42, offset: 6598},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 269, col: 47, offset: 6603},
	expr: &seqExpr{
	pos: position{line: 269, col: 48, offset: 6604},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 269, col: 48, offset: 6604},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 269, col: 52, offset: 6608},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 269, col: 54, offset: 6610},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 269, col: 65, offset: 6621},
	name: "_",
},
&litMatcher{
	pos: position{line: 269, col: 67, offset: 6623},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 284, col: 1, offset: 6960},
	expr: &actionExpr{
	pos: position{line: 284, col: 11, offset: 6970},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 284, col: 11, offset: 6970},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 284, col: 11, offset: 6970},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 284, col: 18, offset: 6977},
	name: "_",
},
&labeledExpr{
	pos: position{line: 284, col: 20, offset: 6979},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 284, col: 25, offset: 6984},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 284, col: 34, offset: 6993},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 284, col: 38, offset: 6997},
	name: "_",
},
&labeledExpr{
	pos: position{line: 284, col: 40, offset: 6999},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 284, col: 44, offset: 7003},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 284, col: 51, offset: 7010},
	name: "_",
},
&litMatcher{
	pos: position{line: 284, col: 53, offset: 7012},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 296, col: 1, offset: 7234},
	expr: &actionExpr{
	pos: position{line: 296, col: 11, offset: 7244},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 296, col: 11, offset: 7244},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 296, col: 11, offset: 7244},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 296, col: 20, offset: 7253},
	name: "_",
},
&litMatcher{
	pos: position{line: 296, col: 22, offset: 7255},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 296, col: 27, offset: 7260},
	name: "_",
},
&labeledExpr{
	pos: position{line: 296, col: 29, offset: 7262},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 296, col: 35, offset: 7268},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 302, col: 1, offset: 7374},
	expr: &actionExpr{
	pos: position{line: 302, col: 12, offset: 7385},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 302, col: 12, offset: 7385},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 302, col: 12, offset: 7385},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 302, col: 21, offset: 7394},
	name: "_",
},
&labeledExpr{
	pos: position{line: 302, col: 23, offset: 7396},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 302, col: 27, offset: 7400},
	expr: &ruleRefExpr{
	pos: position{line: 302, col: 27, offset: 7400},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 302, col: 35, offset: 7408},
	name: "_",
},
&litMatcher{
	pos: position{line: 302, col: 37, offset: 7410},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EnrichDo",
	pos: position{line: 312, col: 1, offset: 7555},
	expr: &actionExpr{
	pos: position{line: 312, col: 13, offset: 7567},
	run: (*parser).callonEnrichDo1,
	expr: &seqExpr{
	pos: position{line: 312, col: 13, offset: 7567},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 312, col: 13, offset: 7567},
	label: "fld",
	expr: &ruleRefExpr{
	pos: position{line: 312, col: 17, offset: 7571},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 312, col: 26, offset: 7580},
	name: "_",
},
&litMatcher{
	pos: position{line: 312, col: 28, offset: 7582},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 312, col: 32, offset: 7586},
	name: "_",
},
&labeledExpr{
	pos: position{line: 312, col: 34, offset: 7588},
	label: "expr",
	expr: &ruleRefExpr{
	pos: position{line: 312, col: 39, offset: 7593},
	name: "MExpr",
},
},
	},
},
},
},
{
	name: "CountAgg",
	pos: position{line: 326, col: 1, offset: 7883},
	expr: &actionExpr{
	pos: position{line: 326, col: 13, offset: 7895},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 326, col: 13, offset: 7895},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 326, col: 13, offset: 7895},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 326, col: 22, offset: 7904},
	name: "_",
},
&labeledExpr{
	pos: position{line: 326, col: 24, offset: 7906},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 326, col: 29, offset: 7911},
	expr: &ruleRefExpr{
	pos: position{line: 326, col: 29, offset: 7911},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 326, col: 42, offset: 7924},
	name: "_",
},
&litMatcher{
	pos: position{line: 326, col: 44, offset: 7926},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MaxAgg",
	pos: position{line: 336, col: 1, offset: 8100},
	expr: &actionExpr{
	pos: position{line: 336, col: 11, offset: 8110},
	run: (*parser).callonMaxAgg1,
	expr: &seqExpr{
	pos: position{line: 336, col: 11, offset: 8110},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 336, col: 11, offset: 8110},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 336, col: 17, offset: 8116},
	expr: &ruleRefExpr{
	pos: position{line: 336, col: 17, offset: 8116},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 336, col: 27, offset: 8126},
	name: "_",
},
&litMatcher{
	pos: position{line: 336, col: 29, offset: 8128},
	val: "max(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 336, col: 36, offset: 8135},
	name: "_",
},
&labeledExpr{
	pos: position{line: 336, col: 38, offset: 8137},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 336, col: 44, offset: 8143},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 336, col: 53, offset: 8152},
	name: "_",
},
&labeledExpr{
	pos: position{line: 336, col: 55, offset: 8154},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 336, col: 60, offset: 8159},
	expr: &seqExpr{
	pos: position{line: 336, col: 61, offset: 8160},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 336, col: 61, offset: 8160},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 336, col: 65, offset: 8164},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 336, col: 67, offset: 8166},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 336, col: 81, offset: 8180},
	name: "_",
},
&litMatcher{
	pos: position{line: 336, col: 83, offset: 8182},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MinAgg",
	pos: position{line: 358, col: 1, offset: 8643},
	expr: &actionExpr{
	pos: position{line: 358, col: 11, offset: 8653},
	run: (*parser).callonMinAgg1,
	expr: &seqExpr{
	pos: position{line: 358, col: 11, offset: 8653},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 358, col: 11, offset: 8653},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 358, col: 17, offset: 8659},
	expr: &ruleRefExpr{
	pos: position{line: 358, col: 17, offset: 8659},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 358, col: 27, offset: 8669},
	name: "_",
},
&litMatcher{
	pos: position{line: 358, col: 29, offset: 8671},
	val: "min(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 358, col: 36, offset: 8678},
	name: "_",
},
&labeledExpr{
	pos: position{line: 358, col: 38, offset: 8680},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 358, col: 44, offset: 8686},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 358, col: 53, offset: 8695},
	name: "_",
},
&labeledExpr{
	pos: position{line: 358, col: 55, offset: 8697},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 358, col: 60, offset: 8702},
	expr: &seqExpr{
	pos: position{line: 358, col: 61, offset: 8703},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 358, col: 61, offset: 8703},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 358, col: 65, offset: 8707},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 358, col: 67, offset: 8709},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 358, col: 81, offset: 8723},
	name: "_",
},
&litMatcher{
	pos: position{line: 358, col: 83, offset: 8725},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AvgAgg",
	pos: position{line: 380, col: 1, offset: 9186},
	expr: &actionExpr{
	pos: position{line: 380, col: 11, offset: 9196},
	run: (*parser).callonAvgAgg1,
	expr: &seqExpr{
	pos: position{line: 380, col: 11, offset: 9196},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 380, col: 11, offset: 9196},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 380, col: 17, offset: 9202},
	expr: &ruleRefExpr{
	pos: position{line: 380, col: 17, offset: 9202},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 380, col: 27, offset: 9212},
	name: "_",
},
&litMatcher{
	pos: position{line: 380, col: 29, offset: 9214},
	val: "avg(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 380, col: 36, offset: 9221},
	name: "_",
},
&labeledExpr{
	pos: position{line: 380, col: 38, offset: 9223},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 380, col: 44, offset: 9229},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 380, col: 53, offset: 9238},
	name: "_",
},
&labeledExpr{
	pos: position{line: 380, col: 55, offset: 9240},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 380, col: 60, offset: 9245},
	expr: &seqExpr{
	pos: position{line: 380, col: 61, offset: 9246},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 380, col: 61, offset: 9246},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 380, col: 65, offset: 9250},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 380, col: 67, offset: 9252},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 380, col: 81, offset: 9266},
	name: "_",
},
&litMatcher{
	pos: position{line: 380, col: 83, offset: 9268},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SumAgg",
	pos: position{line: 402, col: 1, offset: 9729},
	expr: &actionExpr{
	pos: position{line: 402, col: 11, offset: 9739},
	run: (*parser).callonSumAgg1,
	expr: &seqExpr{
	pos: position{line: 402, col: 11, offset: 9739},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 402, col: 11, offset: 9739},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 402, col: 17, offset: 9745},
	expr: &ruleRefExpr{
	pos: position{line: 402, col: 17, offset: 9745},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 402, col: 27, offset: 9755},
	name: "_",
},
&litMatcher{
	pos: position{line: 402, col: 29, offset: 9757},
	val: "sum(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 402, col: 36, offset: 9764},
	name: "_",
},
&labeledExpr{
	pos: position{line: 402, col: 38, offset: 9766},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 402, col: 44, offset: 9772},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 402, col: 53, offset: 9781},
	name: "_",
},
&labeledExpr{
	pos: position{line: 402, col: 55, offset: 9783},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 402, col: 60, offset: 9788},
	expr: &seqExpr{
	pos: position{line: 402, col: 61, offset: 9789},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 402, col: 61, offset: 9789},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 402, col: 65, offset: 9793},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 402, col: 67, offset: 9795},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 402, col: 81, offset: 9809},
	name: "_",
},
&litMatcher{
	pos: position{line: 402, col: 83, offset: 9811},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarAgg",
	pos: position{line: 424, col: 1, offset: 10272},
	expr: &actionExpr{
	pos: position{line: 424, col: 11, offset: 10282},
	run: (*parser).callonVarAgg1,
	expr: &seqExpr{
	pos: position{line: 424, col: 11, offset: 10282},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 424, col: 11, offset: 10282},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 424, col: 17, offset: 10288},
	expr: &ruleRefExpr{
	pos: position{line: 424, col: 17, offset: 10288},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 424, col: 27, offset: 10298},
	name: "_",
},
&litMatcher{
	pos: position{line: 424, col: 29, offset: 10300},
	val: "var(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 424, col: 36, offset: 10307},
	name: "_",
},
&labeledExpr{
	pos: position{line: 424, col: 38, offset: 10309},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 424, col: 44, offset: 10315},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 424, col: 53, offset: 10324},
	name: "_",
},
&labeledExpr{
	pos: position{line: 424, col: 55, offset: 10326},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 424, col: 60, offset: 10331},
	expr: &seqExpr{
	pos: position{line: 424, col: 61, offset: 10332},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 424, col: 61, offset: 10332},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 424, col: 65, offset: 10336},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 424, col: 67, offset: 10338},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 424, col: 81, offset: 10352},
	name: "_",
},
&litMatcher{
	pos: position{line: 424, col: 83, offset: 10354},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DistinctCountAgg",
	pos: position{line: 446, col: 1, offset: 10820},
	expr: &actionExpr{
	pos: position{line: 446, col: 21, offset: 10840},
	run: (*parser).callonDistinctCountAgg1,
	expr: &seqExpr{
	pos: position{line: 446, col: 21, offset: 10840},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 446, col: 21, offset: 10840},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 446, col: 27, offset: 10846},
	expr: &ruleRefExpr{
	pos: position{line: 446, col: 27, offset: 10846},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 446, col: 37, offset: 10856},
	name: "_",
},
&litMatcher{
	pos: position{line: 446, col: 39, offset: 10858},
	val: "dcount(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 446, col: 49, offset: 10868},
	name: "_",
},
&labeledExpr{
	pos: position{line: 446, col: 51, offset: 10870},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 446, col: 57, offset: 10876},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 446, col: 66, offset: 10885},
	name: "_",
},
&labeledExpr{
	pos: position{line: 446, col: 68, offset: 10887},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 446, col: 73, offset: 10892},
	expr: &seqExpr{
	pos: position{line: 446, col: 74, offset: 10893},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 446, col: 74, offset: 10893},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 446, col: 78, offset: 10897},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 446, col: 80, offset: 10899},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 446, col: 94, offset: 10913},
	name: "_",
},
&litMatcher{
	pos: position{line: 446, col: 96, offset: 10915},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "ModeAgg",
	pos: position{line: 468, col: 1, offset: 11386},
	expr: &actionExpr{
	pos: position{line: 468, col: 12, offset: 11397},
	run: (*parser).callonModeAgg1,
	expr: &seqExpr{
	pos: position{line: 468, col: 12, offset: 11397},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 468, col: 12, offset: 11397},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 468, col: 18, offset: 11403},
	expr: &ruleRefExpr{
	pos: position{line: 468, col: 18, offset: 11403},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 468, col: 28, offset: 11413},
	name: "_",
},
&litMatcher{
	pos: position{line: 468, col: 30, offset: 11415},
	val: "mode(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 468, col: 38, offset: 11423},
	name: "_",
},
&labeledExpr{
	pos: position{line: 468, col: 40, offset: 11425},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 468, col: 46, offset: 11431},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 468, col: 55, offset: 11440},
	name: "_",
},
&labeledExpr{
	pos: position{line: 468, col: 57, offset: 11442},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 468, col: 62, offset: 11447},
	expr: &seqExpr{
	pos: position{line: 468, col: 63, offset: 11448},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 468, col: 63, offset: 11448},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 468, col: 67, offset: 11452},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 468, col: 69, offset: 11454},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 468, col: 83, offset: 11468},
	name: "_",
},
&litMatcher{
	pos: position{line: 468, col: 85, offset: 11470},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuantileAgg",
	pos: position{line: 490, col: 1, offset: 11932},
	expr: &actionExpr{
	pos: position{line: 490, col: 16, offset: 11947},
	run: (*parser).callonQuantileAgg1,
	expr: &seqExpr{
	pos: position{line: 490, col: 16, offset: 11947},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 490, col: 16, offset: 11947},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 490, col: 22, offset: 11953},
	expr: &ruleRefExpr{
	pos: position{line: 490, col: 22, offset: 11953},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 490, col: 32, offset: 11963},
	name: "_",
},
&litMatcher{
	pos: position{line: 490, col: 34, offset: 11965},
	val: "quantile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 490, col: 46, offset: 11977},
	name: "_",
},
&labeledExpr{
	pos: position{line: 490, col: 48, offset: 11979},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 490, col: 54, offset: 11985},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 490, col: 63, offset: 11994},
	name: "_",
},
&labeledExpr{
	pos: position{line: 490, col: 66, offset: 11997},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 490, col: 73, offset: 12004},
	expr: &seqExpr{
	pos: position{line: 490, col: 74, offset: 12005},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 490, col: 74, offset: 12005},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 490, col: 78, offset: 12009},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 490, col: 80, offset: 12011},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 490, col: 91, offset: 12022},
	name: "_",
},
&litMatcher{
	pos: position{line: 490, col: 93, offset: 12024},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 490, col: 97, offset: 12028},
	name: "_",
},
&labeledExpr{
	pos: position{line: 490, col: 99, offset: 12030},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 490, col: 103, offset: 12034},
	name: "Float",
},
},
&ruleRefExpr{
	pos: position{line: 490, col: 109, offset: 12040},
	name: "_",
},
&labeledExpr{
	pos: position{line: 490, col: 111, offset: 12042},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 490, col: 116, offset: 12047},
	expr: &seqExpr{
	pos: position{line: 490, col: 117, offset: 12048},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 490, col: 117, offset: 12048},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 490, col: 121, offset: 12052},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 490, col: 123, offset: 12054},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 490, col: 137, offset: 12068},
	name: "_",
},
&litMatcher{
	pos: position{line: 490, col: 139, offset: 12070},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MedianAgg",
	pos: position{line: 520, col: 1, offset: 12715},
	expr: &actionExpr{
	pos: position{line: 520, col: 14, offset: 12728},
	run: (*parser).callonMedianAgg1,
	expr: &seqExpr{
	pos: position{line: 520, col: 14, offset: 12728},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 520, col: 14, offset: 12728},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 520, col: 20, offset: 12734},
	expr: &ruleRefExpr{
	pos: position{line: 520, col: 20, offset: 12734},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 520, col: 30, offset: 12744},
	name: "_",
},
&litMatcher{
	pos: position{line: 520, col: 32, offset: 12746},
	val: "median(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 520, col: 42, offset: 12756},
	name: "_",
},
&labeledExpr{
	pos: position{line: 520, col: 44, offset: 12758},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 520, col: 50, offset: 12764},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 520, col: 59, offset: 12773},
	name: "_",
},
&labeledExpr{
	pos: position{line: 520, col: 62, offset: 12776},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 520, col: 69, offset: 12783},
	expr: &seqExpr{
	pos: position{line: 520, col: 70, offset: 12784},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 520, col: 70, offset: 12784},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 520, col: 74, offset: 12788},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 520, col: 76, offset: 12790},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 520, col: 87, offset: 12801},
	name: "_",
},
&labeledExpr{
	pos: position{line: 520, col: 89, offset: 12803},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 520, col: 94, offset: 12808},
	expr: &seqExpr{
	pos: position{line: 520, col: 95, offset: 12809},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 520, col: 95, offset: 12809},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 520, col: 99, offset: 12813},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 520, col: 101, offset: 12815},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 520, col: 115, offset: 12829},
	name: "_",
},
&litMatcher{
	pos: position{line: 520, col: 117, offset: 12831},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuartileAgg",
	pos: position{line: 552, col: 1, offset: 13508},
	expr: &actionExpr{
	pos: position{line: 552, col: 16, offset: 13523},
	run: (*parser).callonQuartileAgg1,
	expr: &seqExpr{
	pos: position{line: 552, col: 16, offset: 13523},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 552, col: 16, offset: 13523},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 552, col: 22, offset: 13529},
	expr: &ruleRefExpr{
	pos: position{line: 552, col: 22, offset: 13529},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 552, col: 32, offset: 13539},
	name: "_",
},
&litMatcher{
	pos: position{line: 552, col: 34, offset: 13541},
	val: "quartile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 552, col: 46, offset: 13553},
	name: "_",
},
&labeledExpr{
	pos: position{line: 552, col: 48, offset: 13555},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 552, col: 54, offset: 13561},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 552, col: 63, offset: 13570},
	name: "_",
},
&labeledExpr{
	pos: position{line: 552, col: 66, offset: 13573},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 552, col: 73, offset: 13580},
	expr: &seqExpr{
	pos: position{line: 552, col: 74, offset: 13581},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 552, col: 74, offset: 13581},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 552, col: 78, offset: 13585},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 552, col: 80, offset: 13587},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 552, col: 91, offset: 13598},
	name: "_",
},
&litMatcher{
	pos: position{line: 552, col: 93, offset: 13600},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 552, col: 97, offset: 13604},
	name: "_",
},
&labeledExpr{
	pos: position{line: 552, col: 99, offset: 13606},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 552, col: 103, offset: 13610},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 552, col: 110, offset: 13617},
	name: "_",
},
&labeledExpr{
	pos: position{line: 552, col: 112, offset: 13619},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 552, col: 117, offset: 13624},
	expr: &seqExpr{
	pos: position{line: 552, col: 118, offset: 13625},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 552, col: 118, offset: 13625},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 552, col: 122, offset: 13629},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 552, col: 124, offset: 13631},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 552, col: 138, offset: 13645},
	name: "_",
},
&litMatcher{
	pos: position{line: 552, col: 140, offset: 13647},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CovAgg",
	pos: position{line: 598, col: 1, offset: 14674},
	expr: &actionExpr{
	pos: position{line: 598, col: 11, offset: 14684},
	run: (*parser).callonCovAgg1,
	expr: &seqExpr{
	pos: position{line: 598, col: 11, offset: 14684},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 598, col: 11, offset: 14684},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 598, col: 17, offset: 14690},
	expr: &ruleRefExpr{
	pos: position{line: 598, col: 17, offset: 14690},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 598, col: 27, offset: 14700},
	name: "_",
},
&litMatcher{
	pos: position{line: 598, col: 29, offset: 14702},
	val: "cov(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 598, col: 36, offset: 14709},
	name: "_",
},
&labeledExpr{
	pos: position{line: 598, col: 38, offset: 14711},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 598, col: 45, offset: 14718},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 598, col: 54, offset: 14727},
	name: "_",
},
&litMatcher{
	pos: position{line: 598, col: 56, offset: 14729},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 598, col: 60, offset: 14733},
	name: "_",
},
&labeledExpr{
	pos: position{line: 598, col: 62, offset: 14735},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 598, col: 69, offset: 14742},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 598, col: 78, offset: 14751},
	name: "_",
},
&labeledExpr{
	pos: position{line: 598, col: 80, offset: 14753},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 598, col: 85, offset: 14758},
	expr: &seqExpr{
	pos: position{line: 598, col: 86, offset: 14759},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 598, col: 86, offset: 14759},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 598, col: 90, offset: 14763},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 598, col: 92, offset: 14765},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 598, col: 106, offset: 14779},
	name: "_",
},
&litMatcher{
	pos: position{line: 598, col: 108, offset: 14781},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CorrAgg",
	pos: position{line: 621, col: 1, offset: 15306},
	expr: &actionExpr{
	pos: position{line: 621, col: 12, offset: 15317},
	run: (*parser).callonCorrAgg1,
	expr: &seqExpr{
	pos: position{line: 621, col: 12, offset: 15317},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 621, col: 12, offset: 15317},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 621, col: 18, offset: 15323},
	expr: &ruleRefExpr{
	pos: position{line: 621, col: 18, offset: 15323},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 621, col: 28, offset: 15333},
	name: "_",
},
&litMatcher{
	pos: position{line: 621, col: 30, offset: 15335},
	val: "correlation(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 621, col: 45, offset: 15350},
	name: "_",
},
&labeledExpr{
	pos: position{line: 621, col: 47, offset: 15352},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 621, col: 54, offset: 15359},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 621, col: 63, offset: 15368},
	name: "_",
},
&litMatcher{
	pos: position{line: 621, col: 65, offset: 15370},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 621, col: 69, offset: 15374},
	name: "_",
},
&labeledExpr{
	pos: position{line: 621, col: 71, offset: 15376},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 621, col: 78, offset: 15383},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 621, col: 87, offset: 15392},
	name: "_",
},
&labeledExpr{
	pos: position{line: 621, col: 89, offset: 15394},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 621, col: 94, offset: 15399},
	expr: &seqExpr{
	pos: position{line: 621, col: 95, offset: 15400},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 621, col: 95, offset: 15400},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 621, col: 99, offset: 15404},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 621, col: 101, offset: 15406},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 621, col: 115, offset: 15420},
	name: "_",
},
&litMatcher{
	pos: position{line: 621, col: 117, offset: 15422},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PValueAgg",
	pos: position{line: 644, col: 1, offset: 15948},
	expr: &actionExpr{
	pos: position{line: 644, col: 14, offset: 15961},
	run: (*parser).callonPValueAgg1,
	expr: &seqExpr{
	pos: position{line: 644, col: 14, offset: 15961},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 644, col: 14, offset: 15961},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 644, col: 20, offset: 15967},
	expr: &ruleRefExpr{
	pos: position{line: 644, col: 20, offset: 15967},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 644, col: 30, offset: 15977},
	name: "_",
},
&litMatcher{
	pos: position{line: 644, col: 32, offset: 15979},
	val: "pvalue(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 644, col: 42, offset: 15989},
	name: "_",
},
&labeledExpr{
	pos: position{line: 644, col: 44, offset: 15991},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 644, col: 51, offset: 15998},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 644, col: 60, offset: 16007},
	name: "_",
},
&litMatcher{
	pos: position{line: 644, col: 62, offset: 16009},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 644, col: 66, offset: 16013},
	name: "_",
},
&labeledExpr{
	pos: position{line: 644, col: 68, offset: 16015},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 644, col: 75, offset: 16022},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 644, col: 84, offset: 16031},
	name: "_",
},
&labeledExpr{
	pos: position{line: 644, col: 86, offset: 16033},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 644, col: 91, offset: 16038},
	expr: &seqExpr{
	pos: position{line: 644, col: 92, offset: 16039},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 644, col: 92, offset: 16039},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 644, col: 96, offset: 16043},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 644, col: 98, offset: 16045},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 644, col: 112, offset: 16059},
	name: "_",
},
&litMatcher{
	pos: position{line: 644, col: 114, offset: 16061},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AliasG",
	pos: position{line: 669, col: 1, offset: 16602},
	expr: &actionExpr{
	pos: position{line: 669, col: 11, offset: 16612},
	run: (*parser).callonAliasG1,
	expr: &seqExpr{
	pos: position{line: 669, col: 11, offset: 16612},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 669, col: 11, offset: 16612},
	name: "_",
},
&litMatcher{
	pos: position{line: 669, col: 13, offset: 16614},
	val: "as",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 669, col: 18, offset: 16619},
	name: "_",
},
&labeledExpr{
	pos: position{line: 669, col: 20, offset: 16621},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 669, col: 26, offset: 16627},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "EqAliasG",
	pos: position{line: 676, col: 1, offset: 16779},
	expr: &actionExpr{
	pos: position{line: 676, col: 13, offset: 16791},
	run: (*parser).callonEqAliasG1,
	expr: &seqExpr{
	pos: position{line: 676, col: 13, offset: 16791},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 676, col: 13, offset: 16791},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 676, col: 19, offset: 16797},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 676, col: 28, offset: 16806},
	name: "_",
},
&litMatcher{
	pos: position{line: 676, col: 30, offset: 16808},
	val: "=",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarParamListG",
	pos: position{line: 685, col: 1, offset: 17009},
	expr: &actionExpr{
	pos: position{line: 685, col: 18, offset: 17026},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 685, col: 18, offset: 17026},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 685, col: 18, offset: 17026},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 685, col: 24, offset: 17032},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 685, col: 33, offset: 17041},
	name: "_",
},
&labeledExpr{
	pos: position{line: 685, col: 35, offset: 17043},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 685, col: 40, offset: 17048},
	expr: &seqExpr{
	pos: position{line: 685, col: 41, offset: 17049},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 685, col: 41, offset: 17049},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 685, col: 45, offset: 17053},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 685, col: 47, offset: 17055},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 685, col: 56, offset: 17064},
	name: "_",
},
	},
},
},
},
	},
},
},
},
	},
}
func (c *current) onMain1(agg, rest interface{}) (interface{}, error) {

    jobs := []interface{}{agg}
    rst := cast.ToIfaceSlice(rest)
    for _, v := range rst {
        fmt.Println("Value",v)
    }

    return jobs, nil
}

func (p *parser) callonMain1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMain1(stack["agg"], stack["rest"])
}

func (c *current) onMExpr1(first, rest interface{}) (interface{}, error) {

        return parseMathExpr(first, rest), nil
}

func (p *parser) callonMExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMExpr1(stack["first"], stack["rest"])
}

func (c *current) onMValue1(val interface{}) (interface{}, error) {

        return val, nil
}

func (p *parser) callonMValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMValue1(stack["val"])
}

func (c *current) onMOps6() (interface{}, error) {

        return string(c.text), nil
}

func (p *parser) callonMOps6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMOps6()
}

func (c *current) onMFactor2(ex interface{}) (interface{}, error) {


    e, _ := cast.TryString(ex)
    return "(" + e + ")", nil
}

func (p *parser) callonMFactor2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMFactor2(stack["ex"])
}

func (c *current) onMFactor8(ex interface{}) (interface{}, error) {

    return ex, nil
}

func (p *parser) callonMFactor8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMFactor8(stack["ex"])
}

func (c *current) onExpr1(first, rest interface{}) (interface{}, error) {

        return parseMathExpr(first, rest), nil
}

func (p *parser) callonExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpr1(stack["first"], stack["rest"])
}

func (c *current) onValue1(val interface{}) (interface{}, error) {

        return val, nil
}

func (p *parser) callonValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue1(stack["val"])
}

func (c *current) onVariable1() (interface{}, error) {

        return string(c.text), nil
}

func (p *parser) callonVariable1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariable1()
}

func (c *current) onString_Type11() (interface{}, error) {

        c.text = bytes.Replace(c.text, []byte(`\/`), []byte(`/`), -1)
        str,_ := strconv.Unquote(string(c.text))

        return cast.GetStringValue(str), nil
}

func (p *parser) callonString_Type11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onString_Type11()
}

func (c *current) onWord1() (interface{}, error) {

        return string(c.text), nil
}

func (p *parser) callonWord1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWord1()
}

func (c *current) onNumber1() (interface{}, error) {

        return strconv.ParseInt(string(c.text), 10, 64)
}

func (p *parser) callonNumber1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNumber1()
}

func (c *current) onPositiveNumber1() (interface{}, error) {

        return strconv.ParseInt(string(c.text), 10, 64)
}

func (p *parser) callonPositiveNumber1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPositiveNumber1()
}

func (c *current) onFloat1() (interface{}, error) {

    return strconv.ParseFloat(string(c.text), 64)
}

func (p *parser) callonFloat1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFloat1()
}

func (c *current) onFileJobG1(filename interface{}) (interface{}, error) {

    if fname, ok := cast.TryString(filename); ok {
        return FileJob{Filename: fname}, nil
    }
    return nil, errors.New("could not decode source file name")
}

func (p *parser) callonFileJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFileJobG1(stack["filename"])
}

func (c *current) onFakeJobG1(numData interface{}) (interface{}, error) {

    if n, ok := cast.TryInt(numData); ok {
        return FakeJob{NumData: n}, nil
    }
    return nil, errors.New("could not decode source file name")
}

func (p *parser) callonFakeJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFakeJobG1(stack["numData"])
}

func (c *current) onFileName1() (interface{}, error) {

    nm, _ := cast.TryString(string(c.text))

    return nm, nil
}

func (p *parser) callonFileName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFileName1()
}

func (c *current) onDoJobG1(job interface{}) (interface{}, error) {

        return DoNodeJob{Function: job}, nil
}

func (p *parser) callonDoJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDoJobG1(stack["job"])
}

func (c *current) onMapJobG1(job interface{}) (interface{}, error) {

        return DoNodeJob{Function: job}, nil
}

func (p *parser) callonMapJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMapJobG1(stack["job"])
}

func (c *current) onAggJobG1(params, job, rest interface{}) (interface{}, error) {

        jobSlices := cast.ToIfaceSlice(rest)
        jobs := []interface{}{job}

        for _, v := range jobSlices {
                j := cast.ToIfaceSlice(v)
                jobs = append(jobs, j[2])
        }

        ps, _ := params.([]string)

        return AggNodeJob{Functions: jobs, GroupBy: ps}, nil
}

func (p *parser) callonAggJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAggJobG1(stack["params"], stack["job"], stack["rest"])
}

func (c *current) onAggGroupBy1(param, rest interface{}) (interface{}, error) {

    p1, _ := cast.TryString(param)
    params := []string{p1}

    restSlice := cast.ToIfaceSlice(rest)
    for _, v := range restSlice {
            ps := cast.ToIfaceSlice(v)
            p2, _ := cast.TryString(ps[2])
            params = append(params, p2)
    }

    return params, nil
}

func (p *parser) callonAggGroupBy1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAggGroupBy1(stack["param"], stack["rest"])
}

func (c *current) onStdOutG1(params interface{}) (interface{}, error) {


    if header, ok := params.([]string); ok {
        return SinkJob{
            Type: "stdout",
            Header: header,
        }, nil
    }
    return SinkJob{
        Type: "stdout",
        Header: nil,
    }, nil
}

func (p *parser) callonStdOutG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStdOutG1(stack["params"])
}

func (c *current) onBlackHoleG1(params interface{}) (interface{}, error) {

        return SinkJob{Type: "blackhole"}, nil
}

func (p *parser) callonBlackHoleG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlackHoleG1(stack["params"])
}

func (c *current) onPlotG1(widget interface{}) (interface{}, error) {

    return SinkJob{Type: "plot", Args: widget}, nil
}

func (p *parser) callonPlotG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPlotG1(stack["widget"])
}

func (c *current) onBarWidget1(alias, xField, yField, args interface{}) (interface{}, error) {


    title := "Bar Chart"
	if alias != nil {
		if nm, ok := cast.TryString(alias); ok {
			title = nm
		}
	}

	xfld, _ := cast.TryString(xField)
	yfld, _ := cast.TryString(yField)

	width := int64(3)
	gap := int64(2)
	if args != nil {
	    argSlice := cast.ToIfaceSlice(args)
	    width, _ = cast.TryInt(argSlice[2])

	    gapArg := argSlice[4]
	    if gapArg != nil {
	        gapSlice := cast.ToIfaceSlice(gapArg)
	        gap, _ = cast.TryInt(gapSlice[2])
	    }
	}

	title += "(q to quit)"

	return Plot{Type: "bar", Widget: BarPlot{Title: title, XField: xfld, YField: yfld, BarWidth: int(width), BarGap: int(gap)}}, nil
}

func (p *parser) callonBarWidget1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBarWidget1(stack["alias"], stack["xField"], stack["yField"], stack["args"])
}

func (c *current) onFilterDo1(filt interface{}) (interface{}, error) {

        if f, ok := filt.(Filter); ok {
                return f, nil
        }
        return nil, errors.New("Could not parse filter")
}

func (p *parser) callonFilterDo1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFilterDo1(stack["filt"])
}

func (c *current) onFilterMulti1(first, rest interface{}) (interface{}, error) {

        return newSecondaryFilter(first, rest), nil
}

func (p *parser) callonFilterMulti1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFilterMulti1(stack["first"], stack["rest"])
}

func (c *current) onFilterFactor2(sf interface{}) (interface{}, error) {

    return sf, nil
}

func (p *parser) callonFilterFactor2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFilterFactor2(stack["sf"])
}

func (c *current) onFilterFactor8(pf interface{}) (interface{}, error) {

    return pf, nil
}

func (p *parser) callonFilterFactor8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFilterFactor8(stack["pf"])
}

func (c *current) onFilterSecOp3() (interface{}, error) {

        return string(c.text), nil
}

func (p *parser) callonFilterSecOp3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFilterSecOp3()
}

func (c *current) onFilterPrimary1(name, op, value interface{}) (interface{}, error) {

        return newPrimaryFilter(name.(string), op, value), nil
}

func (p *parser) callonFilterPrimary1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFilterPrimary1(stack["name"], stack["op"], stack["value"])
}

func (c *current) onFilterPriOp7() (interface{}, error) {

        return string(c.text), nil
}

func (p *parser) callonFilterPriOp7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFilterPriOp7()
}

func (c *current) onBranchDo1(br interface{}) (interface{}, error) {

        if branch, ok := cast.TryString(br); ok {
                return Branch{Condition: branch}, nil
        }
        return nil, errors.New("Could not decode branch syntax")
}

func (p *parser) callonBranchDo1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBranchDo1(stack["br"])
}

func (c *current) onBranchPrimary1(name, op, value interface{}) (interface{}, error) {

        return string(c.text), nil
}

func (p *parser) callonBranchPrimary1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBranchPrimary1(stack["name"], stack["op"], stack["value"])
}

func (c *current) onSelectDo1(field, rest interface{}) (interface{}, error) {

    fld, _ := cast.TryString(field)
    fields := []string{ fld }

    rst := cast.ToIfaceSlice(rest)
    for _, r := range rst {
        rstVal := cast.ToIfaceSlice(r)
        fld, _ := cast.TryString(rstVal[2])
        fields = append(fields, fld)
    }

    return Select{Fields: fields}, nil
}

func (p *parser) callonSelectDo1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSelectDo1(stack["field"], stack["rest"])
}

func (c *current) onPickDo1(desc, num interface{}) (interface{}, error) {

    n, _ := cast.TryInt(num)
    if n < 0 {
        log.Panic("Can't take less than 0 items")
    }

    dsc, _ := cast.TryString(desc)

    return Pick{Desc: dsc, Num: uint64(n)}, nil
}

func (p *parser) callonPickDo1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPickDo1(stack["desc"], stack["num"])
}

func (c *current) onSortDo1(alias interface{}) (interface{}, error) {

    fld, _ := cast.TryString(alias)
    return Sort{Field: fld}, nil
}

func (p *parser) callonSortDo1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSortDo1(stack["alias"])
}

func (c *current) onBatchDo1(num interface{}) (interface{}, error) {


    n := int64(-1)
    if num != nil {
        n, _ = cast.TryInt(num)
    }
    return Batch{Num: n}, nil
}

func (p *parser) callonBatchDo1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBatchDo1(stack["num"])
}

func (c *current) onEnrichDo1(fld, expr interface{}) (interface{}, error) {

    field, _ := cast.TryString(fld)
    ex, _ := cast.TryString(expr)

    expression, err := govaluate.NewEvaluableExpression(ex)
    if err != nil {
        panic(err)
    }

    return Enrich{Field: field, Expr: expression}, nil
}

func (p *parser) callonEnrichDo1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEnrichDo1(stack["fld"], stack["expr"])
}

func (c *current) onCountAgg1(args interface{}) (interface{}, error) {

	alias := ""
    name, filter, err := parseAggArgs(alias, args)
    if err != nil {
        return nil, err
    }

    return Count{Alias: name, Filter: filter}, nil
}

func (p *parser) callonCountAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCountAgg1(stack["args"])
}

func (c *current) onMaxAgg1(alias, field, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    fld, _ := cast.TryString(field)
    return Max{Alias: name, Field: fld, Filter: filter}, nil
}

func (p *parser) callonMaxAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMaxAgg1(stack["alias"], stack["field"], stack["args"])
}

func (c *current) onMinAgg1(alias, field, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    fld, _ := cast.TryString(field)
    return Min{Alias: name, Field: fld, Filter: filter}, nil
}

func (p *parser) callonMinAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMinAgg1(stack["alias"], stack["field"], stack["args"])
}

func (c *current) onAvgAgg1(alias, field, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    fld, _ := cast.TryString(field)
    return Avg{Alias: name, Field: fld, Filter: filter}, nil
}

func (p *parser) callonAvgAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAvgAgg1(stack["alias"], stack["field"], stack["args"])
}

func (c *current) onSumAgg1(alias, field, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    fld, _ := cast.TryString(field)
    return Sum{Alias: name, Field: fld, Filter: filter}, nil
}

func (p *parser) callonSumAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSumAgg1(stack["alias"], stack["field"], stack["args"])
}

func (c *current) onVarAgg1(alias, field, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    fld, _ := cast.TryString(field)
    return Variance{Alias: name, Field: fld, Filter: filter}, nil
}

func (p *parser) callonVarAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVarAgg1(stack["alias"], stack["field"], stack["args"])
}

func (c *current) onDistinctCountAgg1(alias, field, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    fld, _ := cast.TryString(field)
    return DistinctCount{Alias: name, Field: fld, Filter: filter}, nil
}

func (p *parser) callonDistinctCountAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDistinctCountAgg1(stack["alias"], stack["field"], stack["args"])
}

func (c *current) onModeAgg1(alias, field, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    fld, _ := cast.TryString(field)
    return Mode{Alias: name, Field: fld, Filter: filter}, nil
}

func (p *parser) callonModeAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onModeAgg1(stack["alias"], stack["field"], stack["args"])
}

func (c *current) onQuantileAgg1(alias, field, weight, qth, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    wt := ""
    if weight != nil {
        w := cast.ToIfaceSlice(weight)
        wt, _ = cast.TryString(w[2])
    }

    q, _ := cast.TryFloat(qth)

    fld, _ := cast.TryString(field)
    return Quantile{Alias: name, Field: fld, Filter: filter, Weight: wt, Qth: q}, nil
}

func (p *parser) callonQuantileAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQuantileAgg1(stack["alias"], stack["field"], stack["weight"], stack["qth"], stack["args"])
}

func (c *current) onMedianAgg1(alias, field, weight, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    wt := ""
    if weight != nil {
        w := cast.ToIfaceSlice(weight)
        wt, _ = cast.TryString(w[2])
    }

    if name == "" {
        name = "median"
    }

    fld, _ := cast.TryString(field)
    return Quantile{Alias: name, Field: fld, Filter: filter, Weight: wt, Qth: float64(0.5)}, nil
}

func (p *parser) callonMedianAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMedianAgg1(stack["alias"], stack["field"], stack["weight"], stack["args"])
}

func (c *current) onQuartileAgg1(alias, field, weight, qth, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    wt := ""
    if weight != nil {
        w := cast.ToIfaceSlice(weight)
        wt, _ = cast.TryString(w[2])
    }

    q, _ := cast.TryInt(qth)
    if q < 1 || q > 4 {
        log.Panic("Quartile must be in range (1, 4)")
    }

    if name == "" {
        if trail, ok := posMap[int(q)]; ok {
            name = fmt.Sprintf("%v%s quartile", q, trail)
        } else {
            name = fmt.Sprintf("%vth quartile", q)
        }
    }

    qn := float64(q) * 0.25
    if qn > 1 {
        qn = float64(1)
    }

    fld, _ := cast.TryString(field)
    return Quantile{Alias: name, Field: fld, Filter: filter, Weight: wt, Qth: qn}, nil
}

func (p *parser) callonQuartileAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQuartileAgg1(stack["alias"], stack["field"], stack["weight"], stack["qth"], stack["args"])
}

func (c *current) onCovAgg1(alias, field1, field2, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    fld1, _ := cast.TryString(field1)
    fld2, _ := cast.TryString(field2)
    return Covariance{Alias: name, Field1: fld1, Field2: fld2, Filter: filter}, nil
}

func (p *parser) callonCovAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCovAgg1(stack["alias"], stack["field1"], stack["field2"], stack["args"])
}

func (c *current) onCorrAgg1(alias, field1, field2, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    fld1, _ := cast.TryString(field1)
    fld2, _ := cast.TryString(field2)
    return Correlation{Alias: name, Field1: fld1, Field2: fld2, Filter: filter}, nil
}

func (p *parser) callonCorrAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCorrAgg1(stack["alias"], stack["field1"], stack["field2"], stack["args"])
}

func (c *current) onPValueAgg1(alias, field1, field2, args interface{}) (interface{}, error) {


    var name string
    var filter Filter
    var err error

    if args != nil {
        if ags := cast.ToIfaceSlice(args); ags != nil {
            name, filter, err = parseAggArgs(alias, ags[2])
        }
    } else {
        name, filter, err = parseAggArgs(alias, nil)
    }

    if err != nil {
        return nil, err
    }

    fld1, _ := cast.TryString(field1)
    fld2, _ := cast.TryString(field2)
    return PValue{Alias: name, Field1: fld1, Field2: fld2, Filter: filter}, nil
}

func (p *parser) callonPValueAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPValueAgg1(stack["alias"], stack["field1"], stack["field2"], stack["args"])
}

func (c *current) onAliasG1(alias interface{}) (interface{}, error) {

    if name, ok := cast.TryString(alias); ok {
        return name, nil
    }
    return nil, errors.New("Could not decode alias")
}

func (p *parser) callonAliasG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAliasG1(stack["alias"])
}

func (c *current) onEqAliasG1(alias interface{}) (interface{}, error) {

    if name, ok := cast.TryString(alias); ok {
        return name, nil
    }
    return nil, errors.New("Could not decode alias")
}

func (p *parser) callonEqAliasG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEqAliasG1(stack["alias"])
}

func (c *current) onVarParamListG1(first, rest interface{}) (interface{}, error) {

    return getParamList(first, rest), nil
}

func (p *parser) callonVarParamListG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVarParamListG1(stack["first"], stack["rest"])
}


var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule          = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch         = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos    position
	expr   interface{}
	run    func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs: new(errList),
		data: b,
		pt: savepoint{position: position{line: 1}},
		recover: true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v interface{}
	b bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug bool
	depth  int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules  map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth) + ">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth) + "<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth) + "MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth) + "MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}

