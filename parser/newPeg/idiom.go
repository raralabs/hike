
package newPeg

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
	name: "Command",
	pos: position{line: 5, col: 1, offset: 24},
	expr: &seqExpr{
	pos: position{line: 5, col: 12, offset: 35},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 5, col: 12, offset: 35},
	label: "command",
	expr: &ruleRefExpr{
	pos: position{line: 5, col: 20, offset: 43},
	name: "Statement",
},
},
&labeledExpr{
	pos: position{line: 5, col: 30, offset: 53},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 5, col: 36, offset: 59},
	exprs: []interface{}{
&zeroOrMoreExpr{
	pos: position{line: 5, col: 36, offset: 59},
	expr: &seqExpr{
	pos: position{line: 5, col: 37, offset: 60},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 5, col: 37, offset: 60},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 5, col: 39, offset: 62},
	name: "Statement",
},
	},
},
},
&ruleRefExpr{
	pos: position{line: 5, col: 51, offset: 74},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 5, col: 53, offset: 76},
	name: "TOK_SEMICOLON",
},
	},
},
},
	},
},
},
{
	name: "Statement",
	pos: position{line: 7, col: 1, offset: 94},
	expr: &seqExpr{
	pos: position{line: 7, col: 14, offset: 107},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 7, col: 14, offset: 107},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 7, col: 24, offset: 117},
	name: "Source",
},
},
&labeledExpr{
	pos: position{line: 7, col: 31, offset: 124},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 7, col: 37, offset: 130},
	exprs: []interface{}{
&zeroOrMoreExpr{
	pos: position{line: 7, col: 37, offset: 130},
	expr: &seqExpr{
	pos: position{line: 7, col: 38, offset: 131},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 7, col: 38, offset: 131},
	name: "_",
},
&litMatcher{
	pos: position{line: 7, col: 40, offset: 133},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 7, col: 44, offset: 137},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 7, col: 46, offset: 139},
	name: "Transform",
},
	},
},
},
&ruleRefExpr{
	pos: position{line: 7, col: 58, offset: 151},
	name: "_",
},
&litMatcher{
	pos: position{line: 7, col: 59, offset: 152},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 7, col: 62, offset: 155},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 7, col: 64, offset: 157},
	name: "Sink",
},
&ruleRefExpr{
	pos: position{line: 7, col: 69, offset: 162},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 7, col: 71, offset: 164},
	name: "TOK_SEMICOLON",
},
	},
},
},
	},
},
},
{
	name: "Source",
	pos: position{line: 9, col: 1, offset: 182},
	expr: &choiceExpr{
	pos: position{line: 9, col: 11, offset: 192},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 9, col: 11, offset: 192},
	name: "FileJobG",
},
&ruleRefExpr{
	pos: position{line: 9, col: 20, offset: 201},
	name: "FakeJobG",
},
	},
},
},
{
	name: "Transform",
	pos: position{line: 11, col: 1, offset: 213},
	expr: &choiceExpr{
	pos: position{line: 11, col: 14, offset: 226},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 11, col: 14, offset: 226},
	name: "DoJobG",
},
&ruleRefExpr{
	pos: position{line: 11, col: 21, offset: 233},
	name: "AggJobG",
},
	},
},
},
{
	name: "Sink",
	pos: position{line: 13, col: 1, offset: 244},
	expr: &choiceExpr{
	pos: position{line: 13, col: 9, offset: 252},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 13, col: 9, offset: 252},
	name: "StdOutG",
},
&ruleRefExpr{
	pos: position{line: 13, col: 19, offset: 262},
	name: "BlackHoleG",
},
&ruleRefExpr{
	pos: position{line: 13, col: 32, offset: 275},
	name: "PlotG",
},
	},
},
},
{
	name: "FileJobG",
	pos: position{line: 16, col: 1, offset: 299},
	expr: &actionExpr{
	pos: position{line: 16, col: 13, offset: 311},
	run: (*parser).callonFileJobG1,
	expr: &seqExpr{
	pos: position{line: 16, col: 13, offset: 311},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 16, col: 13, offset: 311},
	name: "_",
},
&litMatcher{
	pos: position{line: 16, col: 15, offset: 313},
	val: "file(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 16, col: 23, offset: 321},
	name: "_",
},
&labeledExpr{
	pos: position{line: 16, col: 25, offset: 323},
	label: "filename",
	expr: &ruleRefExpr{
	pos: position{line: 16, col: 34, offset: 332},
	name: "FileName",
},
},
&ruleRefExpr{
	pos: position{line: 16, col: 43, offset: 341},
	name: "_",
},
&litMatcher{
	pos: position{line: 16, col: 45, offset: 343},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FakeJobG",
	pos: position{line: 23, col: 1, offset: 525},
	expr: &actionExpr{
	pos: position{line: 23, col: 13, offset: 537},
	run: (*parser).callonFakeJobG1,
	expr: &seqExpr{
	pos: position{line: 23, col: 13, offset: 537},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 23, col: 13, offset: 537},
	name: "_",
},
&litMatcher{
	pos: position{line: 23, col: 15, offset: 539},
	val: "fake(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 23, col: 23, offset: 547},
	name: "_",
},
&labeledExpr{
	pos: position{line: 23, col: 25, offset: 549},
	label: "numData",
	expr: &ruleRefExpr{
	pos: position{line: 23, col: 33, offset: 557},
	name: "PositiveNumber",
},
},
&ruleRefExpr{
	pos: position{line: 23, col: 48, offset: 572},
	name: "_",
},
&litMatcher{
	pos: position{line: 23, col: 50, offset: 574},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggJobG",
	pos: position{line: 33, col: 1, offset: 823},
	expr: &actionExpr{
	pos: position{line: 33, col: 12, offset: 834},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 33, col: 12, offset: 834},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 33, col: 12, offset: 834},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 33, col: 18, offset: 840},
	name: "_",
},
&labeledExpr{
	pos: position{line: 33, col: 20, offset: 842},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 33, col: 27, offset: 849},
	expr: &ruleRefExpr{
	pos: position{line: 33, col: 27, offset: 849},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 33, col: 39, offset: 861},
	name: "_",
},
&labeledExpr{
	pos: position{line: 33, col: 41, offset: 863},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 33, col: 45, offset: 867},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 33, col: 53, offset: 875},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 33, col: 58, offset: 880},
	expr: &seqExpr{
	pos: position{line: 33, col: 59, offset: 881},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 33, col: 59, offset: 881},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 33, col: 63, offset: 885},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 33, col: 65, offset: 887},
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
	pos: position{line: 47, col: 1, offset: 1227},
	expr: &choiceExpr{
	pos: position{line: 47, col: 12, offset: 1238},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 47, col: 12, offset: 1238},
	name: "CountAgg",
},
&ruleRefExpr{
	pos: position{line: 47, col: 23, offset: 1249},
	name: "MaxAgg",
},
&ruleRefExpr{
	pos: position{line: 47, col: 32, offset: 1258},
	name: "MinAgg",
},
&ruleRefExpr{
	pos: position{line: 47, col: 41, offset: 1267},
	name: "AvgAgg",
},
&ruleRefExpr{
	pos: position{line: 47, col: 50, offset: 1276},
	name: "SumAgg",
},
&ruleRefExpr{
	pos: position{line: 47, col: 59, offset: 1285},
	name: "ModeAgg",
},
&ruleRefExpr{
	pos: position{line: 48, col: 14, offset: 1307},
	name: "VarAgg",
},
&ruleRefExpr{
	pos: position{line: 48, col: 23, offset: 1316},
	name: "DistinctCountAgg",
},
&ruleRefExpr{
	pos: position{line: 48, col: 42, offset: 1335},
	name: "QuantileAgg",
},
&ruleRefExpr{
	pos: position{line: 48, col: 56, offset: 1349},
	name: "MedianAgg",
},
&ruleRefExpr{
	pos: position{line: 48, col: 68, offset: 1361},
	name: "QuartileAgg",
},
&ruleRefExpr{
	pos: position{line: 49, col: 14, offset: 1387},
	name: "CovAgg",
},
&ruleRefExpr{
	pos: position{line: 49, col: 23, offset: 1396},
	name: "CorrAgg",
},
&ruleRefExpr{
	pos: position{line: 49, col: 33, offset: 1406},
	name: "PValueAgg",
},
	},
},
},
{
	name: "CountAgg",
	pos: position{line: 51, col: 1, offset: 1419},
	expr: &actionExpr{
	pos: position{line: 51, col: 13, offset: 1431},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 51, col: 13, offset: 1431},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 51, col: 13, offset: 1431},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 51, col: 19, offset: 1437},
	expr: &ruleRefExpr{
	pos: position{line: 51, col: 19, offset: 1437},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 51, col: 29, offset: 1447},
	name: "_",
},
&litMatcher{
	pos: position{line: 51, col: 31, offset: 1449},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 51, col: 40, offset: 1458},
	name: "_",
},
&labeledExpr{
	pos: position{line: 51, col: 42, offset: 1460},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 51, col: 47, offset: 1465},
	expr: &ruleRefExpr{
	pos: position{line: 51, col: 47, offset: 1465},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 51, col: 60, offset: 1478},
	name: "_",
},
&litMatcher{
	pos: position{line: 51, col: 62, offset: 1480},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MaxAgg",
	pos: position{line: 61, col: 1, offset: 1654},
	expr: &actionExpr{
	pos: position{line: 61, col: 11, offset: 1664},
	run: (*parser).callonMaxAgg1,
	expr: &seqExpr{
	pos: position{line: 61, col: 11, offset: 1664},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 61, col: 11, offset: 1664},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 61, col: 17, offset: 1670},
	expr: &ruleRefExpr{
	pos: position{line: 61, col: 17, offset: 1670},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 61, col: 27, offset: 1680},
	name: "_",
},
&litMatcher{
	pos: position{line: 61, col: 29, offset: 1682},
	val: "max(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 61, col: 36, offset: 1689},
	name: "_",
},
&labeledExpr{
	pos: position{line: 61, col: 38, offset: 1691},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 61, col: 44, offset: 1697},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 61, col: 53, offset: 1706},
	name: "_",
},
&labeledExpr{
	pos: position{line: 61, col: 55, offset: 1708},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 61, col: 60, offset: 1713},
	expr: &seqExpr{
	pos: position{line: 61, col: 61, offset: 1714},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 61, col: 61, offset: 1714},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 61, col: 65, offset: 1718},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 61, col: 67, offset: 1720},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 61, col: 81, offset: 1734},
	name: "_",
},
&litMatcher{
	pos: position{line: 61, col: 83, offset: 1736},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MinAgg",
	pos: position{line: 83, col: 1, offset: 2197},
	expr: &actionExpr{
	pos: position{line: 83, col: 11, offset: 2207},
	run: (*parser).callonMinAgg1,
	expr: &seqExpr{
	pos: position{line: 83, col: 11, offset: 2207},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 83, col: 11, offset: 2207},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 83, col: 17, offset: 2213},
	expr: &ruleRefExpr{
	pos: position{line: 83, col: 17, offset: 2213},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 83, col: 27, offset: 2223},
	name: "_",
},
&litMatcher{
	pos: position{line: 83, col: 29, offset: 2225},
	val: "min(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 83, col: 36, offset: 2232},
	name: "_",
},
&labeledExpr{
	pos: position{line: 83, col: 38, offset: 2234},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 83, col: 44, offset: 2240},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 83, col: 53, offset: 2249},
	name: "_",
},
&labeledExpr{
	pos: position{line: 83, col: 55, offset: 2251},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 83, col: 60, offset: 2256},
	expr: &seqExpr{
	pos: position{line: 83, col: 61, offset: 2257},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 83, col: 61, offset: 2257},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 83, col: 65, offset: 2261},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 83, col: 67, offset: 2263},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 83, col: 81, offset: 2277},
	name: "_",
},
&litMatcher{
	pos: position{line: 83, col: 83, offset: 2279},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AvgAgg",
	pos: position{line: 105, col: 1, offset: 2740},
	expr: &actionExpr{
	pos: position{line: 105, col: 11, offset: 2750},
	run: (*parser).callonAvgAgg1,
	expr: &seqExpr{
	pos: position{line: 105, col: 11, offset: 2750},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 105, col: 11, offset: 2750},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 105, col: 17, offset: 2756},
	expr: &ruleRefExpr{
	pos: position{line: 105, col: 17, offset: 2756},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 105, col: 27, offset: 2766},
	name: "_",
},
&litMatcher{
	pos: position{line: 105, col: 29, offset: 2768},
	val: "avg(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 105, col: 36, offset: 2775},
	name: "_",
},
&labeledExpr{
	pos: position{line: 105, col: 38, offset: 2777},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 105, col: 44, offset: 2783},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 105, col: 53, offset: 2792},
	name: "_",
},
&labeledExpr{
	pos: position{line: 105, col: 55, offset: 2794},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 105, col: 60, offset: 2799},
	expr: &seqExpr{
	pos: position{line: 105, col: 61, offset: 2800},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 105, col: 61, offset: 2800},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 105, col: 65, offset: 2804},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 105, col: 67, offset: 2806},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 105, col: 81, offset: 2820},
	name: "_",
},
&litMatcher{
	pos: position{line: 105, col: 83, offset: 2822},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SumAgg",
	pos: position{line: 127, col: 1, offset: 3283},
	expr: &actionExpr{
	pos: position{line: 127, col: 11, offset: 3293},
	run: (*parser).callonSumAgg1,
	expr: &seqExpr{
	pos: position{line: 127, col: 11, offset: 3293},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 127, col: 11, offset: 3293},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 127, col: 17, offset: 3299},
	expr: &ruleRefExpr{
	pos: position{line: 127, col: 17, offset: 3299},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 127, col: 27, offset: 3309},
	name: "_",
},
&litMatcher{
	pos: position{line: 127, col: 29, offset: 3311},
	val: "sum(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 127, col: 36, offset: 3318},
	name: "_",
},
&labeledExpr{
	pos: position{line: 127, col: 38, offset: 3320},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 127, col: 44, offset: 3326},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 127, col: 53, offset: 3335},
	name: "_",
},
&labeledExpr{
	pos: position{line: 127, col: 55, offset: 3337},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 127, col: 60, offset: 3342},
	expr: &seqExpr{
	pos: position{line: 127, col: 61, offset: 3343},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 127, col: 61, offset: 3343},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 127, col: 65, offset: 3347},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 127, col: 67, offset: 3349},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 127, col: 81, offset: 3363},
	name: "_",
},
&litMatcher{
	pos: position{line: 127, col: 83, offset: 3365},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarAgg",
	pos: position{line: 149, col: 1, offset: 3826},
	expr: &actionExpr{
	pos: position{line: 149, col: 11, offset: 3836},
	run: (*parser).callonVarAgg1,
	expr: &seqExpr{
	pos: position{line: 149, col: 11, offset: 3836},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 149, col: 11, offset: 3836},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 149, col: 17, offset: 3842},
	expr: &ruleRefExpr{
	pos: position{line: 149, col: 17, offset: 3842},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 149, col: 27, offset: 3852},
	name: "_",
},
&litMatcher{
	pos: position{line: 149, col: 29, offset: 3854},
	val: "var(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 149, col: 36, offset: 3861},
	name: "_",
},
&labeledExpr{
	pos: position{line: 149, col: 38, offset: 3863},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 149, col: 44, offset: 3869},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 149, col: 53, offset: 3878},
	name: "_",
},
&labeledExpr{
	pos: position{line: 149, col: 55, offset: 3880},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 149, col: 60, offset: 3885},
	expr: &seqExpr{
	pos: position{line: 149, col: 61, offset: 3886},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 149, col: 61, offset: 3886},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 149, col: 65, offset: 3890},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 149, col: 67, offset: 3892},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 149, col: 81, offset: 3906},
	name: "_",
},
&litMatcher{
	pos: position{line: 149, col: 83, offset: 3908},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DistinctCountAgg",
	pos: position{line: 171, col: 1, offset: 4374},
	expr: &actionExpr{
	pos: position{line: 171, col: 21, offset: 4394},
	run: (*parser).callonDistinctCountAgg1,
	expr: &seqExpr{
	pos: position{line: 171, col: 21, offset: 4394},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 171, col: 21, offset: 4394},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 171, col: 27, offset: 4400},
	expr: &ruleRefExpr{
	pos: position{line: 171, col: 27, offset: 4400},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 171, col: 37, offset: 4410},
	name: "_",
},
&litMatcher{
	pos: position{line: 171, col: 39, offset: 4412},
	val: "dcount(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 171, col: 49, offset: 4422},
	name: "_",
},
&labeledExpr{
	pos: position{line: 171, col: 51, offset: 4424},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 171, col: 57, offset: 4430},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 171, col: 66, offset: 4439},
	name: "_",
},
&labeledExpr{
	pos: position{line: 171, col: 68, offset: 4441},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 171, col: 73, offset: 4446},
	expr: &seqExpr{
	pos: position{line: 171, col: 74, offset: 4447},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 171, col: 74, offset: 4447},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 171, col: 78, offset: 4451},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 171, col: 80, offset: 4453},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 171, col: 94, offset: 4467},
	name: "_",
},
&litMatcher{
	pos: position{line: 171, col: 96, offset: 4469},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "ModeAgg",
	pos: position{line: 193, col: 1, offset: 4940},
	expr: &actionExpr{
	pos: position{line: 193, col: 12, offset: 4951},
	run: (*parser).callonModeAgg1,
	expr: &seqExpr{
	pos: position{line: 193, col: 12, offset: 4951},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 193, col: 12, offset: 4951},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 193, col: 18, offset: 4957},
	expr: &ruleRefExpr{
	pos: position{line: 193, col: 18, offset: 4957},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 193, col: 28, offset: 4967},
	name: "_",
},
&litMatcher{
	pos: position{line: 193, col: 30, offset: 4969},
	val: "mode(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 193, col: 38, offset: 4977},
	name: "_",
},
&labeledExpr{
	pos: position{line: 193, col: 40, offset: 4979},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 193, col: 46, offset: 4985},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 193, col: 55, offset: 4994},
	name: "_",
},
&labeledExpr{
	pos: position{line: 193, col: 57, offset: 4996},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 193, col: 62, offset: 5001},
	expr: &seqExpr{
	pos: position{line: 193, col: 63, offset: 5002},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 193, col: 63, offset: 5002},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 193, col: 67, offset: 5006},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 193, col: 69, offset: 5008},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 193, col: 83, offset: 5022},
	name: "_",
},
&litMatcher{
	pos: position{line: 193, col: 85, offset: 5024},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuantileAgg",
	pos: position{line: 215, col: 1, offset: 5486},
	expr: &actionExpr{
	pos: position{line: 215, col: 16, offset: 5501},
	run: (*parser).callonQuantileAgg1,
	expr: &seqExpr{
	pos: position{line: 215, col: 16, offset: 5501},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 215, col: 16, offset: 5501},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 215, col: 22, offset: 5507},
	expr: &ruleRefExpr{
	pos: position{line: 215, col: 22, offset: 5507},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 215, col: 32, offset: 5517},
	name: "_",
},
&litMatcher{
	pos: position{line: 215, col: 34, offset: 5519},
	val: "quantile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 215, col: 46, offset: 5531},
	name: "_",
},
&labeledExpr{
	pos: position{line: 215, col: 48, offset: 5533},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 215, col: 54, offset: 5539},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 215, col: 63, offset: 5548},
	name: "_",
},
&labeledExpr{
	pos: position{line: 215, col: 66, offset: 5551},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 215, col: 73, offset: 5558},
	expr: &seqExpr{
	pos: position{line: 215, col: 74, offset: 5559},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 215, col: 74, offset: 5559},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 215, col: 78, offset: 5563},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 215, col: 80, offset: 5565},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 215, col: 91, offset: 5576},
	name: "_",
},
&litMatcher{
	pos: position{line: 215, col: 93, offset: 5578},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 215, col: 97, offset: 5582},
	name: "_",
},
&labeledExpr{
	pos: position{line: 215, col: 99, offset: 5584},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 215, col: 103, offset: 5588},
	name: "Float",
},
},
&ruleRefExpr{
	pos: position{line: 215, col: 109, offset: 5594},
	name: "_",
},
&labeledExpr{
	pos: position{line: 215, col: 111, offset: 5596},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 215, col: 116, offset: 5601},
	expr: &seqExpr{
	pos: position{line: 215, col: 117, offset: 5602},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 215, col: 117, offset: 5602},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 215, col: 121, offset: 5606},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 215, col: 123, offset: 5608},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 215, col: 137, offset: 5622},
	name: "_",
},
&litMatcher{
	pos: position{line: 215, col: 139, offset: 5624},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MedianAgg",
	pos: position{line: 245, col: 1, offset: 6269},
	expr: &actionExpr{
	pos: position{line: 245, col: 14, offset: 6282},
	run: (*parser).callonMedianAgg1,
	expr: &seqExpr{
	pos: position{line: 245, col: 14, offset: 6282},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 245, col: 14, offset: 6282},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 245, col: 20, offset: 6288},
	expr: &ruleRefExpr{
	pos: position{line: 245, col: 20, offset: 6288},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 245, col: 30, offset: 6298},
	name: "_",
},
&litMatcher{
	pos: position{line: 245, col: 32, offset: 6300},
	val: "median(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 245, col: 42, offset: 6310},
	name: "_",
},
&labeledExpr{
	pos: position{line: 245, col: 44, offset: 6312},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 245, col: 50, offset: 6318},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 245, col: 59, offset: 6327},
	name: "_",
},
&labeledExpr{
	pos: position{line: 245, col: 62, offset: 6330},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 245, col: 69, offset: 6337},
	expr: &seqExpr{
	pos: position{line: 245, col: 70, offset: 6338},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 245, col: 70, offset: 6338},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 245, col: 74, offset: 6342},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 245, col: 76, offset: 6344},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 245, col: 87, offset: 6355},
	name: "_",
},
&labeledExpr{
	pos: position{line: 245, col: 89, offset: 6357},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 245, col: 94, offset: 6362},
	expr: &seqExpr{
	pos: position{line: 245, col: 95, offset: 6363},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 245, col: 95, offset: 6363},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 245, col: 99, offset: 6367},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 245, col: 101, offset: 6369},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 245, col: 115, offset: 6383},
	name: "_",
},
&litMatcher{
	pos: position{line: 245, col: 117, offset: 6385},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuartileAgg",
	pos: position{line: 277, col: 1, offset: 7062},
	expr: &actionExpr{
	pos: position{line: 277, col: 16, offset: 7077},
	run: (*parser).callonQuartileAgg1,
	expr: &seqExpr{
	pos: position{line: 277, col: 16, offset: 7077},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 277, col: 16, offset: 7077},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 277, col: 22, offset: 7083},
	expr: &ruleRefExpr{
	pos: position{line: 277, col: 22, offset: 7083},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 277, col: 32, offset: 7093},
	name: "_",
},
&litMatcher{
	pos: position{line: 277, col: 34, offset: 7095},
	val: "quartile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 277, col: 46, offset: 7107},
	name: "_",
},
&labeledExpr{
	pos: position{line: 277, col: 48, offset: 7109},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 277, col: 54, offset: 7115},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 277, col: 63, offset: 7124},
	name: "_",
},
&labeledExpr{
	pos: position{line: 277, col: 66, offset: 7127},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 277, col: 73, offset: 7134},
	expr: &seqExpr{
	pos: position{line: 277, col: 74, offset: 7135},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 277, col: 74, offset: 7135},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 277, col: 78, offset: 7139},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 277, col: 80, offset: 7141},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 277, col: 91, offset: 7152},
	name: "_",
},
&litMatcher{
	pos: position{line: 277, col: 93, offset: 7154},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 277, col: 97, offset: 7158},
	name: "_",
},
&labeledExpr{
	pos: position{line: 277, col: 99, offset: 7160},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 277, col: 103, offset: 7164},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 277, col: 110, offset: 7171},
	name: "_",
},
&labeledExpr{
	pos: position{line: 277, col: 112, offset: 7173},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 277, col: 117, offset: 7178},
	expr: &seqExpr{
	pos: position{line: 277, col: 118, offset: 7179},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 277, col: 118, offset: 7179},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 277, col: 122, offset: 7183},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 277, col: 124, offset: 7185},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 277, col: 138, offset: 7199},
	name: "_",
},
&litMatcher{
	pos: position{line: 277, col: 140, offset: 7201},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CovAgg",
	pos: position{line: 323, col: 1, offset: 8228},
	expr: &actionExpr{
	pos: position{line: 323, col: 11, offset: 8238},
	run: (*parser).callonCovAgg1,
	expr: &seqExpr{
	pos: position{line: 323, col: 11, offset: 8238},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 323, col: 11, offset: 8238},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 323, col: 17, offset: 8244},
	expr: &ruleRefExpr{
	pos: position{line: 323, col: 17, offset: 8244},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 323, col: 27, offset: 8254},
	name: "_",
},
&litMatcher{
	pos: position{line: 323, col: 29, offset: 8256},
	val: "cov(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 323, col: 36, offset: 8263},
	name: "_",
},
&labeledExpr{
	pos: position{line: 323, col: 38, offset: 8265},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 323, col: 45, offset: 8272},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 323, col: 54, offset: 8281},
	name: "_",
},
&litMatcher{
	pos: position{line: 323, col: 56, offset: 8283},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 323, col: 60, offset: 8287},
	name: "_",
},
&labeledExpr{
	pos: position{line: 323, col: 62, offset: 8289},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 323, col: 69, offset: 8296},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 323, col: 78, offset: 8305},
	name: "_",
},
&labeledExpr{
	pos: position{line: 323, col: 80, offset: 8307},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 323, col: 85, offset: 8312},
	expr: &seqExpr{
	pos: position{line: 323, col: 86, offset: 8313},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 323, col: 86, offset: 8313},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 323, col: 90, offset: 8317},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 323, col: 92, offset: 8319},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 323, col: 106, offset: 8333},
	name: "_",
},
&litMatcher{
	pos: position{line: 323, col: 108, offset: 8335},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CorrAgg",
	pos: position{line: 346, col: 1, offset: 8860},
	expr: &actionExpr{
	pos: position{line: 346, col: 12, offset: 8871},
	run: (*parser).callonCorrAgg1,
	expr: &seqExpr{
	pos: position{line: 346, col: 12, offset: 8871},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 346, col: 12, offset: 8871},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 346, col: 18, offset: 8877},
	expr: &ruleRefExpr{
	pos: position{line: 346, col: 18, offset: 8877},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 346, col: 28, offset: 8887},
	name: "_",
},
&litMatcher{
	pos: position{line: 346, col: 30, offset: 8889},
	val: "correlation(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 346, col: 45, offset: 8904},
	name: "_",
},
&labeledExpr{
	pos: position{line: 346, col: 47, offset: 8906},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 346, col: 54, offset: 8913},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 346, col: 63, offset: 8922},
	name: "_",
},
&litMatcher{
	pos: position{line: 346, col: 65, offset: 8924},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 346, col: 69, offset: 8928},
	name: "_",
},
&labeledExpr{
	pos: position{line: 346, col: 71, offset: 8930},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 346, col: 78, offset: 8937},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 346, col: 87, offset: 8946},
	name: "_",
},
&labeledExpr{
	pos: position{line: 346, col: 89, offset: 8948},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 346, col: 94, offset: 8953},
	expr: &seqExpr{
	pos: position{line: 346, col: 95, offset: 8954},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 346, col: 95, offset: 8954},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 346, col: 99, offset: 8958},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 346, col: 101, offset: 8960},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 346, col: 115, offset: 8974},
	name: "_",
},
&litMatcher{
	pos: position{line: 346, col: 117, offset: 8976},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PValueAgg",
	pos: position{line: 369, col: 1, offset: 9502},
	expr: &actionExpr{
	pos: position{line: 369, col: 14, offset: 9515},
	run: (*parser).callonPValueAgg1,
	expr: &seqExpr{
	pos: position{line: 369, col: 14, offset: 9515},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 369, col: 14, offset: 9515},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 369, col: 20, offset: 9521},
	expr: &ruleRefExpr{
	pos: position{line: 369, col: 20, offset: 9521},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 369, col: 30, offset: 9531},
	name: "_",
},
&litMatcher{
	pos: position{line: 369, col: 32, offset: 9533},
	val: "pvalue(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 369, col: 42, offset: 9543},
	name: "_",
},
&labeledExpr{
	pos: position{line: 369, col: 44, offset: 9545},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 369, col: 51, offset: 9552},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 369, col: 60, offset: 9561},
	name: "_",
},
&litMatcher{
	pos: position{line: 369, col: 62, offset: 9563},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 369, col: 66, offset: 9567},
	name: "_",
},
&labeledExpr{
	pos: position{line: 369, col: 68, offset: 9569},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 369, col: 75, offset: 9576},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 369, col: 84, offset: 9585},
	name: "_",
},
&labeledExpr{
	pos: position{line: 369, col: 86, offset: 9587},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 369, col: 91, offset: 9592},
	expr: &seqExpr{
	pos: position{line: 369, col: 92, offset: 9593},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 369, col: 92, offset: 9593},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 369, col: 96, offset: 9597},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 369, col: 98, offset: 9599},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 369, col: 112, offset: 9613},
	name: "_",
},
&litMatcher{
	pos: position{line: 369, col: 114, offset: 9615},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DoJobG",
	pos: position{line: 393, col: 1, offset: 10167},
	expr: &actionExpr{
	pos: position{line: 393, col: 11, offset: 10177},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 393, col: 11, offset: 10177},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 393, col: 16, offset: 10182},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 393, col: 16, offset: 10182},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 393, col: 27, offset: 10193},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 393, col: 38, offset: 10204},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 393, col: 49, offset: 10215},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 393, col: 58, offset: 10224},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 393, col: 67, offset: 10233},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 399, col: 1, offset: 10317},
	expr: &actionExpr{
	pos: position{line: 399, col: 13, offset: 10329},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 399, col: 13, offset: 10329},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 399, col: 13, offset: 10329},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 399, col: 23, offset: 10339},
	name: "_",
},
&labeledExpr{
	pos: position{line: 399, col: 25, offset: 10341},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 399, col: 28, offset: 10344},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 399, col: 42, offset: 10358},
	name: "_",
},
&litMatcher{
	pos: position{line: 399, col: 44, offset: 10360},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 406, col: 1, offset: 10555},
	expr: &actionExpr{
	pos: position{line: 406, col: 18, offset: 10572},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 406, col: 18, offset: 10572},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 406, col: 18, offset: 10572},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 406, col: 23, offset: 10577},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 406, col: 28, offset: 10582},
	name: "_",
},
&labeledExpr{
	pos: position{line: 406, col: 30, offset: 10584},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 406, col: 33, offset: 10587},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 406, col: 45, offset: 10599},
	name: "_",
},
&labeledExpr{
	pos: position{line: 406, col: 47, offset: 10601},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 406, col: 54, offset: 10608},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 410, col: 1, offset: 10659},
	expr: &ruleRefExpr{
	pos: position{line: 410, col: 16, offset: 10674},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 413, col: 1, offset: 10708},
	expr: &actionExpr{
	pos: position{line: 413, col: 13, offset: 10720},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 413, col: 13, offset: 10720},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 413, col: 13, offset: 10720},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 413, col: 23, offset: 10730},
	name: "_",
},
&labeledExpr{
	pos: position{line: 413, col: 25, offset: 10732},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 413, col: 31, offset: 10738},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 413, col: 40, offset: 10747},
	name: "_",
},
&labeledExpr{
	pos: position{line: 413, col: 42, offset: 10749},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 413, col: 47, offset: 10754},
	expr: &seqExpr{
	pos: position{line: 413, col: 48, offset: 10755},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 413, col: 48, offset: 10755},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 413, col: 52, offset: 10759},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 413, col: 54, offset: 10761},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 413, col: 65, offset: 10772},
	name: "_",
},
&litMatcher{
	pos: position{line: 413, col: 67, offset: 10774},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 428, col: 1, offset: 11111},
	expr: &actionExpr{
	pos: position{line: 428, col: 11, offset: 11121},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 428, col: 11, offset: 11121},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 428, col: 11, offset: 11121},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 428, col: 18, offset: 11128},
	name: "_",
},
&labeledExpr{
	pos: position{line: 428, col: 20, offset: 11130},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 428, col: 25, offset: 11135},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 428, col: 34, offset: 11144},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 428, col: 38, offset: 11148},
	name: "_",
},
&labeledExpr{
	pos: position{line: 428, col: 40, offset: 11150},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 428, col: 44, offset: 11154},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 428, col: 51, offset: 11161},
	name: "_",
},
&litMatcher{
	pos: position{line: 428, col: 53, offset: 11163},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 440, col: 1, offset: 11385},
	expr: &actionExpr{
	pos: position{line: 440, col: 11, offset: 11395},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 440, col: 11, offset: 11395},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 440, col: 11, offset: 11395},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 440, col: 20, offset: 11404},
	name: "_",
},
&litMatcher{
	pos: position{line: 440, col: 22, offset: 11406},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 440, col: 27, offset: 11411},
	name: "_",
},
&labeledExpr{
	pos: position{line: 440, col: 29, offset: 11413},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 440, col: 35, offset: 11419},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 446, col: 1, offset: 11525},
	expr: &actionExpr{
	pos: position{line: 446, col: 12, offset: 11536},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 446, col: 12, offset: 11536},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 446, col: 12, offset: 11536},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 446, col: 21, offset: 11545},
	name: "_",
},
&labeledExpr{
	pos: position{line: 446, col: 23, offset: 11547},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 446, col: 27, offset: 11551},
	expr: &ruleRefExpr{
	pos: position{line: 446, col: 27, offset: 11551},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 446, col: 35, offset: 11559},
	name: "_",
},
&litMatcher{
	pos: position{line: 446, col: 37, offset: 11561},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 457, col: 1, offset: 11708},
	expr: &actionExpr{
	pos: position{line: 457, col: 13, offset: 11720},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 457, col: 13, offset: 11720},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 457, col: 13, offset: 11720},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 457, col: 23, offset: 11730},
	name: "_",
},
&labeledExpr{
	pos: position{line: 457, col: 25, offset: 11732},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 457, col: 30, offset: 11737},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 457, col: 42, offset: 11749},
	name: "_",
},
&litMatcher{
	pos: position{line: 457, col: 44, offset: 11751},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 466, col: 1, offset: 12037},
	expr: &actionExpr{
	pos: position{line: 466, col: 16, offset: 12052},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 466, col: 16, offset: 12052},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 466, col: 16, offset: 12052},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 466, col: 22, offset: 12058},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 466, col: 35, offset: 12071},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 466, col: 40, offset: 12076},
	expr: &seqExpr{
	pos: position{line: 466, col: 41, offset: 12077},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 466, col: 41, offset: 12077},
	name: "_",
},
&labeledExpr{
	pos: position{line: 466, col: 43, offset: 12079},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 466, col: 46, offset: 12082},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 466, col: 58, offset: 12094},
	name: "_",
},
&labeledExpr{
	pos: position{line: 466, col: 60, offset: 12096},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 466, col: 63, offset: 12099},
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
	pos: position{line: 471, col: 1, offset: 12257},
	expr: &choiceExpr{
	pos: position{line: 471, col: 17, offset: 12273},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 471, col: 17, offset: 12273},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 471, col: 17, offset: 12273},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 471, col: 17, offset: 12273},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 471, col: 21, offset: 12277},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 471, col: 24, offset: 12280},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 471, col: 36, offset: 12292},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 473, col: 5, offset: 12323},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 473, col: 5, offset: 12323},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 473, col: 8, offset: 12326},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 477, col: 1, offset: 12368},
	expr: &choiceExpr{
	pos: position{line: 477, col: 16, offset: 12383},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 477, col: 16, offset: 12383},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 477, col: 24, offset: 12391},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 477, col: 24, offset: 12391},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 484, col: 1, offset: 12570},
	expr: &actionExpr{
	pos: position{line: 484, col: 18, offset: 12587},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 484, col: 18, offset: 12587},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 484, col: 18, offset: 12587},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 484, col: 23, offset: 12592},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 484, col: 28, offset: 12597},
	name: "_",
},
&labeledExpr{
	pos: position{line: 484, col: 30, offset: 12599},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 484, col: 33, offset: 12602},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 484, col: 45, offset: 12614},
	name: "_",
},
&labeledExpr{
	pos: position{line: 484, col: 47, offset: 12616},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 484, col: 53, offset: 12622},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 484, col: 59, offset: 12628},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 488, col: 1, offset: 12702},
	expr: &choiceExpr{
	pos: position{line: 488, col: 16, offset: 12717},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 488, col: 16, offset: 12717},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 488, col: 23, offset: 12724},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 488, col: 30, offset: 12731},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 488, col: 36, offset: 12737},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 488, col: 42, offset: 12743},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 488, col: 49, offset: 12750},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 488, col: 49, offset: 12750},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MapJobG",
	pos: position{line: 493, col: 1, offset: 12821},
	expr: &actionExpr{
	pos: position{line: 493, col: 12, offset: 12832},
	run: (*parser).callonMapJobG1,
	expr: &seqExpr{
	pos: position{line: 493, col: 12, offset: 12832},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 493, col: 12, offset: 12832},
	val: "map",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 493, col: 18, offset: 12838},
	name: "_",
},
&labeledExpr{
	pos: position{line: 493, col: 20, offset: 12840},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 493, col: 25, offset: 12845},
	name: "EnrichDo",
},
},
	},
},
},
},
{
	name: "EnrichDo",
	pos: position{line: 497, col: 1, offset: 12926},
	expr: &actionExpr{
	pos: position{line: 497, col: 13, offset: 12938},
	run: (*parser).callonEnrichDo1,
	expr: &seqExpr{
	pos: position{line: 497, col: 13, offset: 12938},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 497, col: 13, offset: 12938},
	label: "fld",
	expr: &ruleRefExpr{
	pos: position{line: 497, col: 17, offset: 12942},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 497, col: 26, offset: 12951},
	name: "_",
},
&litMatcher{
	pos: position{line: 497, col: 28, offset: 12953},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 497, col: 32, offset: 12957},
	name: "_",
},
&labeledExpr{
	pos: position{line: 497, col: 34, offset: 12959},
	label: "expr",
	expr: &ruleRefExpr{
	pos: position{line: 497, col: 39, offset: 12964},
	name: "MExpr",
},
},
	},
},
},
},
{
	name: "StdOutG",
	pos: position{line: 513, col: 1, offset: 13278},
	expr: &actionExpr{
	pos: position{line: 513, col: 12, offset: 13289},
	run: (*parser).callonStdOutG1,
	expr: &seqExpr{
	pos: position{line: 513, col: 12, offset: 13289},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 513, col: 12, offset: 13289},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 513, col: 22, offset: 13299},
	name: "_",
},
&labeledExpr{
	pos: position{line: 513, col: 24, offset: 13301},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 513, col: 31, offset: 13308},
	expr: &ruleRefExpr{
	pos: position{line: 513, col: 32, offset: 13309},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 513, col: 48, offset: 13325},
	name: "_",
},
&litMatcher{
	pos: position{line: 513, col: 50, offset: 13327},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BlackHoleG",
	pos: position{line: 527, col: 1, offset: 13573},
	expr: &actionExpr{
	pos: position{line: 527, col: 15, offset: 13587},
	run: (*parser).callonBlackHoleG1,
	expr: &seqExpr{
	pos: position{line: 527, col: 15, offset: 13587},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 527, col: 15, offset: 13587},
	val: "blackhole(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 527, col: 28, offset: 13600},
	name: "_",
},
&labeledExpr{
	pos: position{line: 527, col: 30, offset: 13602},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 527, col: 37, offset: 13609},
	expr: &ruleRefExpr{
	pos: position{line: 527, col: 38, offset: 13610},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 527, col: 54, offset: 13626},
	name: "_",
},
&litMatcher{
	pos: position{line: 527, col: 56, offset: 13628},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PlotG",
	pos: position{line: 531, col: 1, offset: 13688},
	expr: &actionExpr{
	pos: position{line: 531, col: 10, offset: 13697},
	run: (*parser).callonPlotG1,
	expr: &seqExpr{
	pos: position{line: 531, col: 10, offset: 13697},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 531, col: 10, offset: 13697},
	val: "plot",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 531, col: 17, offset: 13704},
	name: "_",
},
&labeledExpr{
	pos: position{line: 531, col: 19, offset: 13706},
	label: "widget",
	expr: &ruleRefExpr{
	pos: position{line: 531, col: 27, offset: 13714},
	name: "BarWidget",
},
},
	},
},
},
},
{
	name: "FileName",
	pos: position{line: 537, col: 1, offset: 13805},
	expr: &actionExpr{
	pos: position{line: 537, col: 13, offset: 13817},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 537, col: 13, offset: 13817},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 537, col: 13, offset: 13817},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 537, col: 22, offset: 13826},
	expr: &seqExpr{
	pos: position{line: 537, col: 23, offset: 13827},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 537, col: 23, offset: 13827},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 537, col: 27, offset: 13831},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 537, col: 31, offset: 13835},
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
	name: "MExpr",
	pos: position{line: 543, col: 1, offset: 13936},
	expr: &actionExpr{
	pos: position{line: 543, col: 10, offset: 13945},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 543, col: 10, offset: 13945},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 543, col: 10, offset: 13945},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 543, col: 16, offset: 13951},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 543, col: 23, offset: 13958},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 543, col: 28, offset: 13963},
	expr: &seqExpr{
	pos: position{line: 543, col: 29, offset: 13964},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 543, col: 29, offset: 13964},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 543, col: 31, offset: 13966},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 543, col: 36, offset: 13971},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 543, col: 38, offset: 13973},
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
	pos: position{line: 547, col: 1, offset: 14038},
	expr: &actionExpr{
	pos: position{line: 547, col: 11, offset: 14048},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 547, col: 11, offset: 14048},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 547, col: 16, offset: 14053},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 547, col: 16, offset: 14053},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 547, col: 24, offset: 14061},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 547, col: 33, offset: 14070},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 547, col: 48, offset: 14085},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 547, col: 59, offset: 14096},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 551, col: 1, offset: 14138},
	expr: &choiceExpr{
	pos: position{line: 551, col: 9, offset: 14146},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 551, col: 9, offset: 14146},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 551, col: 15, offset: 14152},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 551, col: 21, offset: 14158},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 551, col: 27, offset: 14164},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 551, col: 33, offset: 14170},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 551, col: 33, offset: 14170},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 555, col: 1, offset: 14218},
	expr: &choiceExpr{
	pos: position{line: 555, col: 12, offset: 14229},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 555, col: 12, offset: 14229},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 555, col: 12, offset: 14229},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 555, col: 12, offset: 14229},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 555, col: 16, offset: 14233},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 555, col: 19, offset: 14236},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 555, col: 25, offset: 14242},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 559, col: 5, offset: 14318},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 559, col: 5, offset: 14318},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 559, col: 8, offset: 14321},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 566, col: 1, offset: 14398},
	expr: &actionExpr{
	pos: position{line: 566, col: 9, offset: 14406},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 566, col: 9, offset: 14406},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 566, col: 9, offset: 14406},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 566, col: 15, offset: 14412},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 566, col: 21, offset: 14418},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 566, col: 26, offset: 14423},
	expr: &seqExpr{
	pos: position{line: 566, col: 27, offset: 14424},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 566, col: 27, offset: 14424},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 566, col: 29, offset: 14426},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 566, col: 34, offset: 14431},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 566, col: 36, offset: 14433},
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
	pos: position{line: 570, col: 1, offset: 14497},
	expr: &actionExpr{
	pos: position{line: 570, col: 10, offset: 14506},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 570, col: 10, offset: 14506},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 570, col: 15, offset: 14511},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 570, col: 15, offset: 14511},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 570, col: 23, offset: 14519},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 570, col: 32, offset: 14528},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 570, col: 47, offset: 14543},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 574, col: 1, offset: 14586},
	expr: &actionExpr{
	pos: position{line: 574, col: 13, offset: 14598},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 574, col: 13, offset: 14598},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 574, col: 13, offset: 14598},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 574, col: 20, offset: 14605},
	expr: &choiceExpr{
	pos: position{line: 574, col: 21, offset: 14606},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 574, col: 21, offset: 14606},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 574, col: 31, offset: 14616},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 574, col: 39, offset: 14624},
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
	pos: position{line: 578, col: 1, offset: 14674},
	expr: &actionExpr{
	pos: position{line: 578, col: 17, offset: 14690},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 578, col: 17, offset: 14690},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 578, col: 17, offset: 14690},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 578, col: 21, offset: 14694},
	expr: &choiceExpr{
	pos: position{line: 578, col: 23, offset: 14696},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 578, col: 23, offset: 14696},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 578, col: 23, offset: 14696},
	expr: &ruleRefExpr{
	pos: position{line: 578, col: 24, offset: 14697},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 578, col: 36, offset: 14709,
},
	},
},
&seqExpr{
	pos: position{line: 578, col: 40, offset: 14713},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 578, col: 40, offset: 14713},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 578, col: 45, offset: 14718},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 578, col: 63, offset: 14736},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 585, col: 1, offset: 14917},
	expr: &charClassMatcher{
	pos: position{line: 585, col: 16, offset: 14932},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 587, col: 1, offset: 14950},
	expr: &ruleRefExpr{
	pos: position{line: 587, col: 19, offset: 14968},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 589, col: 1, offset: 14988},
	expr: &charClassMatcher{
	pos: position{line: 589, col: 21, offset: 15008},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 591, col: 1, offset: 15023},
	expr: &actionExpr{
	pos: position{line: 591, col: 9, offset: 15031},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 591, col: 9, offset: 15031},
	expr: &choiceExpr{
	pos: position{line: 591, col: 10, offset: 15032},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 591, col: 10, offset: 15032},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 591, col: 19, offset: 15041},
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
	pos: position{line: 595, col: 1, offset: 15091},
	expr: &choiceExpr{
	pos: position{line: 595, col: 11, offset: 15101},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 595, col: 11, offset: 15101},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 595, col: 19, offset: 15109},
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
	pos: position{line: 597, col: 1, offset: 15118},
	expr: &actionExpr{
	pos: position{line: 597, col: 11, offset: 15128},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 597, col: 11, offset: 15128},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 597, col: 11, offset: 15128},
	expr: &litMatcher{
	pos: position{line: 597, col: 11, offset: 15128},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 597, col: 16, offset: 15133},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 601, col: 1, offset: 15206},
	expr: &actionExpr{
	pos: position{line: 601, col: 19, offset: 15224},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 601, col: 19, offset: 15224},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 605, col: 1, offset: 15297},
	expr: &actionExpr{
	pos: position{line: 605, col: 10, offset: 15306},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 605, col: 10, offset: 15306},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 605, col: 10, offset: 15306},
	expr: &litMatcher{
	pos: position{line: 605, col: 10, offset: 15306},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 605, col: 15, offset: 15311},
	expr: &ruleRefExpr{
	pos: position{line: 605, col: 15, offset: 15311},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 605, col: 24, offset: 15320},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 605, col: 28, offset: 15324},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 609, col: 1, offset: 15391},
	expr: &choiceExpr{
	pos: position{line: 609, col: 12, offset: 15402},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 609, col: 12, offset: 15402},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 609, col: 18, offset: 15408},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 609, col: 18, offset: 15408},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 609, col: 38, offset: 15428},
	expr: &ruleRefExpr{
	pos: position{line: 609, col: 38, offset: 15428},
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
	pos: position{line: 611, col: 1, offset: 15445},
	expr: &charClassMatcher{
	pos: position{line: 611, col: 17, offset: 15461},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 613, col: 1, offset: 15470},
	expr: &charClassMatcher{
	pos: position{line: 613, col: 24, offset: 15493},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 615, col: 1, offset: 15502},
	expr: &zeroOrMoreExpr{
	pos: position{line: 615, col: 19, offset: 15520},
	expr: &charClassMatcher{
	pos: position{line: 615, col: 19, offset: 15520},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 617, col: 1, offset: 15534},
	expr: &charClassMatcher{
	pos: position{line: 617, col: 18, offset: 15551},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EOF",
	pos: position{line: 619, col: 1, offset: 15564},
	expr: &notExpr{
	pos: position{line: 619, col: 8, offset: 15571},
	expr: &anyMatcher{
	line: 619, col: 9, offset: 15572,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 621, col: 1, offset: 15577},
	expr: &litMatcher{
	pos: position{line: 621, col: 17, offset: 15593},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 627, col: 1, offset: 15680},
	expr: &actionExpr{
	pos: position{line: 627, col: 18, offset: 15697},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 627, col: 18, offset: 15697},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 627, col: 18, offset: 15697},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 627, col: 24, offset: 15703},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 627, col: 33, offset: 15712},
	name: "_",
},
&labeledExpr{
	pos: position{line: 627, col: 35, offset: 15714},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 627, col: 40, offset: 15719},
	expr: &seqExpr{
	pos: position{line: 627, col: 41, offset: 15720},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 627, col: 41, offset: 15720},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 627, col: 45, offset: 15724},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 627, col: 47, offset: 15726},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 627, col: 56, offset: 15735},
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

func (c *current) onCountAgg1(alias, args interface{}) (interface{}, error) {


    name, filter, err := parseAggArgs(alias, args)
    if err != nil {
        return nil, err
    }

    return Count{Alias: name, Filter: filter}, nil
}

func (p *parser) callonCountAgg1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCountAgg1(stack["alias"], stack["args"])
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

func (c *current) onDoJobG1(job interface{}) (interface{}, error) {

        return DoNodeJob{Function: job}, nil
}

func (p *parser) callonDoJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDoJobG1(stack["job"])
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

func (c *current) onMapJobG1(job interface{}) (interface{}, error) {

        return DoNodeJob{Function: job}, nil
}

func (p *parser) callonMapJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMapJobG1(stack["job"])
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

func (c *current) onFileName1() (interface{}, error) {

    nm, _ := cast.TryString(string(c.text))

    return nm, nil
}

func (p *parser) callonFileName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFileName1()
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

