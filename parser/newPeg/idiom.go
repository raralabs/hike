
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
	"github.com/raralabs/canal/utils/cast"
)


var g = &grammar {
	rules: []*rule{
{
	name: "Command",
	pos: position{line: 5, col: 1, offset: 24},
	expr: &actionExpr{
	pos: position{line: 5, col: 12, offset: 35},
	run: (*parser).callonCommand1,
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
	pos: position{line: 5, col: 52, offset: 75},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 5, col: 54, offset: 77},
	name: "TOK_SEMICOLON",
},
	},
},
},
	},
},
},
},
{
	name: "Statement",
	pos: position{line: 36, col: 1, offset: 903},
	expr: &actionExpr{
	pos: position{line: 36, col: 14, offset: 916},
	run: (*parser).callonStatement1,
	expr: &seqExpr{
	pos: position{line: 36, col: 14, offset: 916},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 36, col: 14, offset: 916},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 36, col: 24, offset: 926},
	name: "Source",
},
},
&labeledExpr{
	pos: position{line: 36, col: 31, offset: 933},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 36, col: 37, offset: 939},
	exprs: []interface{}{
&zeroOrMoreExpr{
	pos: position{line: 36, col: 37, offset: 939},
	expr: &seqExpr{
	pos: position{line: 36, col: 38, offset: 940},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 36, col: 38, offset: 940},
	name: "_",
},
&litMatcher{
	pos: position{line: 36, col: 40, offset: 942},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 36, col: 44, offset: 946},
	name: "_",
},
&labeledExpr{
	pos: position{line: 36, col: 46, offset: 948},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 36, col: 56, offset: 958},
	name: "Transform",
},
},
	},
},
},
&ruleRefExpr{
	pos: position{line: 36, col: 68, offset: 970},
	name: "_",
},
&litMatcher{
	pos: position{line: 36, col: 69, offset: 971},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 36, col: 72, offset: 974},
	name: "_",
},
&labeledExpr{
	pos: position{line: 36, col: 74, offset: 976},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 36, col: 84, offset: 986},
	name: "Sink",
},
},
&ruleRefExpr{
	pos: position{line: 36, col: 89, offset: 991},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 36, col: 91, offset: 993},
	name: "TOK_SEMICOLON",
},
&ruleRefExpr{
	pos: position{line: 36, col: 105, offset: 1007},
	name: "_",
},
	},
},
},
	},
},
},
},
{
	name: "Source",
	pos: position{line: 71, col: 1, offset: 1708},
	expr: &choiceExpr{
	pos: position{line: 71, col: 11, offset: 1718},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 71, col: 11, offset: 1718},
	name: "BranchJobG",
},
&ruleRefExpr{
	pos: position{line: 71, col: 22, offset: 1729},
	name: "FileJobG",
},
&ruleRefExpr{
	pos: position{line: 71, col: 31, offset: 1738},
	name: "FakeJobG",
},
&ruleRefExpr{
	pos: position{line: 71, col: 40, offset: 1747},
	name: "SecSource",
},
	},
},
},
{
	name: "Transform",
	pos: position{line: 73, col: 1, offset: 1760},
	expr: &choiceExpr{
	pos: position{line: 73, col: 14, offset: 1773},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 73, col: 14, offset: 1773},
	name: "AggJobG",
},
&ruleRefExpr{
	pos: position{line: 73, col: 22, offset: 1781},
	name: "DoJobG",
},
&ruleRefExpr{
	pos: position{line: 73, col: 29, offset: 1788},
	name: "JoinJobG",
},
	},
},
},
{
	name: "Sink",
	pos: position{line: 75, col: 1, offset: 1800},
	expr: &choiceExpr{
	pos: position{line: 75, col: 9, offset: 1808},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 75, col: 9, offset: 1808},
	name: "StdOutG",
},
&ruleRefExpr{
	pos: position{line: 75, col: 17, offset: 1816},
	name: "SinkInto",
},
&ruleRefExpr{
	pos: position{line: 75, col: 26, offset: 1825},
	name: "PlotG",
},
	},
},
},
{
	name: "FileJobG",
	pos: position{line: 79, col: 1, offset: 1851},
	expr: &actionExpr{
	pos: position{line: 79, col: 13, offset: 1863},
	run: (*parser).callonFileJobG1,
	expr: &seqExpr{
	pos: position{line: 79, col: 13, offset: 1863},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 79, col: 13, offset: 1863},
	name: "_",
},
&litMatcher{
	pos: position{line: 79, col: 15, offset: 1865},
	val: "file(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 79, col: 23, offset: 1873},
	name: "_",
},
&labeledExpr{
	pos: position{line: 79, col: 25, offset: 1875},
	label: "filename",
	expr: &ruleRefExpr{
	pos: position{line: 79, col: 34, offset: 1884},
	name: "FileName",
},
},
&ruleRefExpr{
	pos: position{line: 79, col: 43, offset: 1893},
	name: "_",
},
&litMatcher{
	pos: position{line: 79, col: 45, offset: 1895},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FakeJobG",
	pos: position{line: 86, col: 1, offset: 2110},
	expr: &actionExpr{
	pos: position{line: 86, col: 13, offset: 2122},
	run: (*parser).callonFakeJobG1,
	expr: &seqExpr{
	pos: position{line: 86, col: 13, offset: 2122},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 86, col: 13, offset: 2122},
	name: "_",
},
&litMatcher{
	pos: position{line: 86, col: 15, offset: 2124},
	val: "fake(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 86, col: 23, offset: 2132},
	name: "_",
},
&labeledExpr{
	pos: position{line: 86, col: 25, offset: 2134},
	label: "numData",
	expr: &ruleRefExpr{
	pos: position{line: 86, col: 33, offset: 2142},
	name: "PositiveNumber",
},
},
&ruleRefExpr{
	pos: position{line: 86, col: 48, offset: 2157},
	name: "_",
},
&litMatcher{
	pos: position{line: 86, col: 50, offset: 2159},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchJobG",
	pos: position{line: 93, col: 1, offset: 2358},
	expr: &actionExpr{
	pos: position{line: 93, col: 15, offset: 2372},
	run: (*parser).callonBranchJobG1,
	expr: &seqExpr{
	pos: position{line: 93, col: 15, offset: 2372},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 93, col: 15, offset: 2372},
	val: "branch",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 93, col: 23, offset: 2380},
	name: "_",
},
&litMatcher{
	pos: position{line: 93, col: 24, offset: 2381},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 93, col: 28, offset: 2385},
	name: "_",
},
&labeledExpr{
	pos: position{line: 93, col: 30, offset: 2387},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 93, col: 33, offset: 2390},
	name: "IDENTIFIER",
},
},
&litMatcher{
	pos: position{line: 93, col: 44, offset: 2401},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SecSource",
	pos: position{line: 100, col: 1, offset: 2617},
	expr: &actionExpr{
	pos: position{line: 100, col: 14, offset: 2630},
	run: (*parser).callonSecSource1,
	expr: &seqExpr{
	pos: position{line: 100, col: 14, offset: 2630},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 100, col: 14, offset: 2630},
	label: "src1",
	expr: &ruleRefExpr{
	pos: position{line: 100, col: 19, offset: 2635},
	name: "IDENTIFIER",
},
},
&labeledExpr{
	pos: position{line: 100, col: 30, offset: 2646},
	label: "rest",
	expr: &zeroOrOneExpr{
	pos: position{line: 100, col: 35, offset: 2651},
	expr: &seqExpr{
	pos: position{line: 100, col: 36, offset: 2652},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 100, col: 36, offset: 2652},
	name: "_",
},
&litMatcher{
	pos: position{line: 100, col: 38, offset: 2654},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 100, col: 41, offset: 2657},
	name: "_",
},
&labeledExpr{
	pos: position{line: 100, col: 43, offset: 2659},
	label: "src2",
	expr: &ruleRefExpr{
	pos: position{line: 100, col: 48, offset: 2664},
	name: "IDENTIFIER",
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
	name: "JoinJobG",
	pos: position{line: 119, col: 1, offset: 3119},
	expr: &actionExpr{
	pos: position{line: 119, col: 13, offset: 3131},
	run: (*parser).callonJoinJobG1,
	expr: &seqExpr{
	pos: position{line: 119, col: 13, offset: 3131},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 119, col: 13, offset: 3131},
	val: "join",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 119, col: 20, offset: 3138},
	name: "_",
},
&labeledExpr{
	pos: position{line: 119, col: 22, offset: 3140},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 119, col: 26, offset: 3144},
	name: "JoinJobs",
},
},
	},
},
},
},
{
	name: "JoinJobs",
	pos: position{line: 123, col: 1, offset: 3217},
	expr: &ruleRefExpr{
	pos: position{line: 123, col: 12, offset: 3228},
	name: "InnerJoinG",
},
},
{
	name: "InnerJoinG",
	pos: position{line: 125, col: 1, offset: 3242},
	expr: &actionExpr{
	pos: position{line: 125, col: 14, offset: 3255},
	run: (*parser).callonInnerJoinG1,
	expr: &seqExpr{
	pos: position{line: 125, col: 14, offset: 3255},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 125, col: 14, offset: 3255},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 125, col: 20, offset: 3261},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 125, col: 29, offset: 3270},
	name: "_",
},
&litMatcher{
	pos: position{line: 125, col: 31, offset: 3272},
	val: "inner",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 125, col: 39, offset: 3280},
	name: "_",
},
&litMatcher{
	pos: position{line: 125, col: 41, offset: 3282},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 125, col: 45, offset: 3286},
	name: "_",
},
&labeledExpr{
	pos: position{line: 125, col: 47, offset: 3288},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 125, col: 53, offset: 3294},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 125, col: 59, offset: 3300},
	name: "_",
},
&litMatcher{
	pos: position{line: 125, col: 61, offset: 3302},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "Query",
	pos: position{line: 132, col: 1, offset: 3495},
	expr: &actionExpr{
	pos: position{line: 132, col: 10, offset: 3504},
	run: (*parser).callonQuery1,
	expr: &seqExpr{
	pos: position{line: 132, col: 10, offset: 3504},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 132, col: 10, offset: 3504},
	val: "select",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 132, col: 19, offset: 3513},
	name: "_",
},
&labeledExpr{
	pos: position{line: 132, col: 21, offset: 3515},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 132, col: 27, offset: 3521},
	name: "Fields",
},
},
&ruleRefExpr{
	pos: position{line: 132, col: 34, offset: 3528},
	name: "_",
},
&litMatcher{
	pos: position{line: 132, col: 36, offset: 3530},
	val: "on",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 132, col: 41, offset: 3535},
	name: "_",
},
&labeledExpr{
	pos: position{line: 132, col: 43, offset: 3537},
	label: "condition",
	expr: &ruleRefExpr{
	pos: position{line: 132, col: 53, offset: 3547},
	name: "JoinCondition",
},
},
&ruleRefExpr{
	pos: position{line: 132, col: 67, offset: 3561},
	name: "_",
},
	},
},
},
},
{
	name: "Fields",
	pos: position{line: 136, col: 1, offset: 3673},
	expr: &choiceExpr{
	pos: position{line: 136, col: 11, offset: 3683},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 136, col: 11, offset: 3683},
	name: "AllFields",
},
&ruleRefExpr{
	pos: position{line: 136, col: 21, offset: 3693},
	name: "SelFields",
},
	},
},
},
{
	name: "AllFields",
	pos: position{line: 138, col: 1, offset: 3706},
	expr: &actionExpr{
	pos: position{line: 138, col: 14, offset: 3719},
	run: (*parser).callonAllFields1,
	expr: &litMatcher{
	pos: position{line: 138, col: 14, offset: 3719},
	val: "*",
	ignoreCase: false,
},
},
},
{
	name: "SelFields",
	pos: position{line: 144, col: 1, offset: 3811},
	expr: &actionExpr{
	pos: position{line: 144, col: 14, offset: 3824},
	run: (*parser).callonSelFields1,
	expr: &seqExpr{
	pos: position{line: 144, col: 14, offset: 3824},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 144, col: 14, offset: 3824},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 144, col: 20, offset: 3830},
	name: "Word",
},
},
&ruleRefExpr{
	pos: position{line: 144, col: 25, offset: 3835},
	name: "_",
},
&labeledExpr{
	pos: position{line: 144, col: 28, offset: 3838},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 144, col: 33, offset: 3843},
	expr: &seqExpr{
	pos: position{line: 144, col: 34, offset: 3844},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 144, col: 34, offset: 3844},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 144, col: 38, offset: 3848},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 144, col: 40, offset: 3850},
	name: "Word",
},
&ruleRefExpr{
	pos: position{line: 144, col: 45, offset: 3855},
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
{
	name: "JoinCondition",
	pos: position{line: 153, col: 1, offset: 4106},
	expr: &actionExpr{
	pos: position{line: 153, col: 18, offset: 4123},
	run: (*parser).callonJoinCondition1,
	expr: &seqExpr{
	pos: position{line: 153, col: 18, offset: 4123},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 153, col: 18, offset: 4123},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 153, col: 23, offset: 4128},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 153, col: 35, offset: 4140},
	name: "_",
},
&labeledExpr{
	pos: position{line: 153, col: 37, offset: 4142},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 153, col: 40, offset: 4145},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 153, col: 52, offset: 4157},
	name: "_",
},
&labeledExpr{
	pos: position{line: 153, col: 54, offset: 4159},
	label: "right",
	expr: &ruleRefExpr{
	pos: position{line: 153, col: 60, offset: 4165},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 153, col: 72, offset: 4177},
	name: "_",
},
	},
},
},
},
{
	name: "JoinFactors",
	pos: position{line: 158, col: 1, offset: 4299},
	expr: &actionExpr{
	pos: position{line: 158, col: 16, offset: 4314},
	run: (*parser).callonJoinFactors1,
	expr: &seqExpr{
	pos: position{line: 158, col: 16, offset: 4314},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 158, col: 16, offset: 4314},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 158, col: 22, offset: 4320},
	name: "Word",
},
},
&ruleRefExpr{
	pos: position{line: 158, col: 27, offset: 4325},
	name: "_",
},
&labeledExpr{
	pos: position{line: 158, col: 29, offset: 4327},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 158, col: 34, offset: 4332},
	expr: &seqExpr{
	pos: position{line: 158, col: 35, offset: 4333},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 158, col: 35, offset: 4333},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 158, col: 39, offset: 4337},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 158, col: 41, offset: 4339},
	name: "Word",
},
&ruleRefExpr{
	pos: position{line: 158, col: 46, offset: 4344},
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
{
	name: "AggJobG",
	pos: position{line: 179, col: 1, offset: 4711},
	expr: &actionExpr{
	pos: position{line: 179, col: 12, offset: 4722},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 179, col: 12, offset: 4722},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 179, col: 12, offset: 4722},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 179, col: 18, offset: 4728},
	name: "_",
},
&labeledExpr{
	pos: position{line: 179, col: 20, offset: 4730},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 179, col: 27, offset: 4737},
	expr: &ruleRefExpr{
	pos: position{line: 179, col: 27, offset: 4737},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 179, col: 39, offset: 4749},
	name: "_",
},
&labeledExpr{
	pos: position{line: 179, col: 41, offset: 4751},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 179, col: 45, offset: 4755},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 179, col: 53, offset: 4763},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 179, col: 58, offset: 4768},
	expr: &seqExpr{
	pos: position{line: 179, col: 59, offset: 4769},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 179, col: 59, offset: 4769},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 179, col: 63, offset: 4773},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 179, col: 65, offset: 4775},
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
	pos: position{line: 194, col: 1, offset: 5150},
	expr: &choiceExpr{
	pos: position{line: 194, col: 12, offset: 5161},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 194, col: 12, offset: 5161},
	name: "CountAgg",
},
&ruleRefExpr{
	pos: position{line: 194, col: 23, offset: 5172},
	name: "MaxAgg",
},
&ruleRefExpr{
	pos: position{line: 194, col: 32, offset: 5181},
	name: "MinAgg",
},
&ruleRefExpr{
	pos: position{line: 194, col: 41, offset: 5190},
	name: "AvgAgg",
},
&ruleRefExpr{
	pos: position{line: 194, col: 50, offset: 5199},
	name: "SumAgg",
},
&ruleRefExpr{
	pos: position{line: 194, col: 59, offset: 5208},
	name: "ModeAgg",
},
&ruleRefExpr{
	pos: position{line: 195, col: 13, offset: 5229},
	name: "VarAgg",
},
&ruleRefExpr{
	pos: position{line: 195, col: 22, offset: 5238},
	name: "DistinctCountAgg",
},
&ruleRefExpr{
	pos: position{line: 195, col: 41, offset: 5257},
	name: "QuantileAgg",
},
&ruleRefExpr{
	pos: position{line: 195, col: 55, offset: 5271},
	name: "MedianAgg",
},
&ruleRefExpr{
	pos: position{line: 195, col: 67, offset: 5283},
	name: "QuartileAgg",
},
&ruleRefExpr{
	pos: position{line: 196, col: 13, offset: 5308},
	name: "CovAgg",
},
&ruleRefExpr{
	pos: position{line: 196, col: 22, offset: 5317},
	name: "CorrAgg",
},
&ruleRefExpr{
	pos: position{line: 196, col: 32, offset: 5327},
	name: "PValueAgg",
},
	},
},
},
{
	name: "CountAgg",
	pos: position{line: 199, col: 1, offset: 5371},
	expr: &actionExpr{
	pos: position{line: 199, col: 13, offset: 5383},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 199, col: 13, offset: 5383},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 199, col: 13, offset: 5383},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 199, col: 19, offset: 5389},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 199, col: 28, offset: 5398},
	name: "_",
},
&litMatcher{
	pos: position{line: 199, col: 30, offset: 5400},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 199, col: 39, offset: 5409},
	name: "_",
},
&labeledExpr{
	pos: position{line: 199, col: 41, offset: 5411},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 199, col: 46, offset: 5416},
	expr: &ruleRefExpr{
	pos: position{line: 199, col: 46, offset: 5416},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 199, col: 59, offset: 5429},
	name: "_",
},
&litMatcher{
	pos: position{line: 199, col: 61, offset: 5431},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MaxAgg",
	pos: position{line: 206, col: 1, offset: 5583},
	expr: &actionExpr{
	pos: position{line: 206, col: 11, offset: 5593},
	run: (*parser).callonMaxAgg1,
	expr: &seqExpr{
	pos: position{line: 206, col: 11, offset: 5593},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 206, col: 11, offset: 5593},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 206, col: 17, offset: 5599},
	expr: &ruleRefExpr{
	pos: position{line: 206, col: 17, offset: 5599},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 206, col: 27, offset: 5609},
	name: "_",
},
&litMatcher{
	pos: position{line: 206, col: 29, offset: 5611},
	val: "max(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 206, col: 36, offset: 5618},
	name: "_",
},
&labeledExpr{
	pos: position{line: 206, col: 38, offset: 5620},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 206, col: 44, offset: 5626},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 206, col: 53, offset: 5635},
	name: "_",
},
&labeledExpr{
	pos: position{line: 206, col: 55, offset: 5637},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 206, col: 60, offset: 5642},
	expr: &seqExpr{
	pos: position{line: 206, col: 61, offset: 5643},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 206, col: 61, offset: 5643},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 206, col: 65, offset: 5647},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 206, col: 67, offset: 5649},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 206, col: 81, offset: 5663},
	name: "_",
},
&litMatcher{
	pos: position{line: 206, col: 83, offset: 5665},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MinAgg",
	pos: position{line: 228, col: 1, offset: 6126},
	expr: &actionExpr{
	pos: position{line: 228, col: 11, offset: 6136},
	run: (*parser).callonMinAgg1,
	expr: &seqExpr{
	pos: position{line: 228, col: 11, offset: 6136},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 228, col: 11, offset: 6136},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 228, col: 17, offset: 6142},
	expr: &ruleRefExpr{
	pos: position{line: 228, col: 17, offset: 6142},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 228, col: 27, offset: 6152},
	name: "_",
},
&litMatcher{
	pos: position{line: 228, col: 29, offset: 6154},
	val: "min(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 228, col: 36, offset: 6161},
	name: "_",
},
&labeledExpr{
	pos: position{line: 228, col: 38, offset: 6163},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 228, col: 44, offset: 6169},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 228, col: 53, offset: 6178},
	name: "_",
},
&labeledExpr{
	pos: position{line: 228, col: 55, offset: 6180},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 228, col: 60, offset: 6185},
	expr: &seqExpr{
	pos: position{line: 228, col: 61, offset: 6186},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 228, col: 61, offset: 6186},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 228, col: 65, offset: 6190},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 228, col: 67, offset: 6192},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 228, col: 81, offset: 6206},
	name: "_",
},
&litMatcher{
	pos: position{line: 228, col: 83, offset: 6208},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AvgAgg",
	pos: position{line: 250, col: 1, offset: 6669},
	expr: &actionExpr{
	pos: position{line: 250, col: 11, offset: 6679},
	run: (*parser).callonAvgAgg1,
	expr: &seqExpr{
	pos: position{line: 250, col: 11, offset: 6679},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 250, col: 11, offset: 6679},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 250, col: 17, offset: 6685},
	expr: &ruleRefExpr{
	pos: position{line: 250, col: 17, offset: 6685},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 250, col: 27, offset: 6695},
	name: "_",
},
&litMatcher{
	pos: position{line: 250, col: 29, offset: 6697},
	val: "avg(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 250, col: 36, offset: 6704},
	name: "_",
},
&labeledExpr{
	pos: position{line: 250, col: 38, offset: 6706},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 250, col: 44, offset: 6712},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 250, col: 53, offset: 6721},
	name: "_",
},
&labeledExpr{
	pos: position{line: 250, col: 55, offset: 6723},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 250, col: 60, offset: 6728},
	expr: &seqExpr{
	pos: position{line: 250, col: 61, offset: 6729},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 250, col: 61, offset: 6729},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 250, col: 65, offset: 6733},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 250, col: 67, offset: 6735},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 250, col: 81, offset: 6749},
	name: "_",
},
&litMatcher{
	pos: position{line: 250, col: 83, offset: 6751},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SumAgg",
	pos: position{line: 272, col: 1, offset: 7212},
	expr: &actionExpr{
	pos: position{line: 272, col: 11, offset: 7222},
	run: (*parser).callonSumAgg1,
	expr: &seqExpr{
	pos: position{line: 272, col: 11, offset: 7222},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 272, col: 11, offset: 7222},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 272, col: 17, offset: 7228},
	expr: &ruleRefExpr{
	pos: position{line: 272, col: 17, offset: 7228},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 272, col: 27, offset: 7238},
	name: "_",
},
&litMatcher{
	pos: position{line: 272, col: 29, offset: 7240},
	val: "sum(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 272, col: 36, offset: 7247},
	name: "_",
},
&labeledExpr{
	pos: position{line: 272, col: 38, offset: 7249},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 272, col: 44, offset: 7255},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 272, col: 53, offset: 7264},
	name: "_",
},
&labeledExpr{
	pos: position{line: 272, col: 55, offset: 7266},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 272, col: 60, offset: 7271},
	expr: &seqExpr{
	pos: position{line: 272, col: 61, offset: 7272},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 272, col: 61, offset: 7272},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 272, col: 65, offset: 7276},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 272, col: 67, offset: 7278},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 272, col: 81, offset: 7292},
	name: "_",
},
&litMatcher{
	pos: position{line: 272, col: 83, offset: 7294},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarAgg",
	pos: position{line: 294, col: 1, offset: 7755},
	expr: &actionExpr{
	pos: position{line: 294, col: 11, offset: 7765},
	run: (*parser).callonVarAgg1,
	expr: &seqExpr{
	pos: position{line: 294, col: 11, offset: 7765},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 294, col: 11, offset: 7765},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 294, col: 17, offset: 7771},
	expr: &ruleRefExpr{
	pos: position{line: 294, col: 17, offset: 7771},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 294, col: 27, offset: 7781},
	name: "_",
},
&litMatcher{
	pos: position{line: 294, col: 29, offset: 7783},
	val: "var(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 294, col: 36, offset: 7790},
	name: "_",
},
&labeledExpr{
	pos: position{line: 294, col: 38, offset: 7792},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 294, col: 44, offset: 7798},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 294, col: 53, offset: 7807},
	name: "_",
},
&labeledExpr{
	pos: position{line: 294, col: 55, offset: 7809},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 294, col: 60, offset: 7814},
	expr: &seqExpr{
	pos: position{line: 294, col: 61, offset: 7815},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 294, col: 61, offset: 7815},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 294, col: 65, offset: 7819},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 294, col: 67, offset: 7821},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 294, col: 81, offset: 7835},
	name: "_",
},
&litMatcher{
	pos: position{line: 294, col: 83, offset: 7837},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DistinctCountAgg",
	pos: position{line: 316, col: 1, offset: 8303},
	expr: &actionExpr{
	pos: position{line: 316, col: 21, offset: 8323},
	run: (*parser).callonDistinctCountAgg1,
	expr: &seqExpr{
	pos: position{line: 316, col: 21, offset: 8323},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 316, col: 21, offset: 8323},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 316, col: 27, offset: 8329},
	expr: &ruleRefExpr{
	pos: position{line: 316, col: 27, offset: 8329},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 316, col: 37, offset: 8339},
	name: "_",
},
&litMatcher{
	pos: position{line: 316, col: 39, offset: 8341},
	val: "dcount(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 316, col: 49, offset: 8351},
	name: "_",
},
&labeledExpr{
	pos: position{line: 316, col: 51, offset: 8353},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 316, col: 57, offset: 8359},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 316, col: 66, offset: 8368},
	name: "_",
},
&labeledExpr{
	pos: position{line: 316, col: 68, offset: 8370},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 316, col: 73, offset: 8375},
	expr: &seqExpr{
	pos: position{line: 316, col: 74, offset: 8376},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 316, col: 74, offset: 8376},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 316, col: 78, offset: 8380},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 316, col: 80, offset: 8382},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 316, col: 94, offset: 8396},
	name: "_",
},
&litMatcher{
	pos: position{line: 316, col: 96, offset: 8398},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "ModeAgg",
	pos: position{line: 338, col: 1, offset: 8869},
	expr: &actionExpr{
	pos: position{line: 338, col: 12, offset: 8880},
	run: (*parser).callonModeAgg1,
	expr: &seqExpr{
	pos: position{line: 338, col: 12, offset: 8880},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 338, col: 12, offset: 8880},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 338, col: 18, offset: 8886},
	expr: &ruleRefExpr{
	pos: position{line: 338, col: 18, offset: 8886},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 338, col: 28, offset: 8896},
	name: "_",
},
&litMatcher{
	pos: position{line: 338, col: 30, offset: 8898},
	val: "mode(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 338, col: 38, offset: 8906},
	name: "_",
},
&labeledExpr{
	pos: position{line: 338, col: 40, offset: 8908},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 338, col: 46, offset: 8914},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 338, col: 55, offset: 8923},
	name: "_",
},
&labeledExpr{
	pos: position{line: 338, col: 57, offset: 8925},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 338, col: 62, offset: 8930},
	expr: &seqExpr{
	pos: position{line: 338, col: 63, offset: 8931},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 338, col: 63, offset: 8931},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 338, col: 67, offset: 8935},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 338, col: 69, offset: 8937},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 338, col: 83, offset: 8951},
	name: "_",
},
&litMatcher{
	pos: position{line: 338, col: 85, offset: 8953},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuantileAgg",
	pos: position{line: 360, col: 1, offset: 9415},
	expr: &actionExpr{
	pos: position{line: 360, col: 16, offset: 9430},
	run: (*parser).callonQuantileAgg1,
	expr: &seqExpr{
	pos: position{line: 360, col: 16, offset: 9430},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 360, col: 16, offset: 9430},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 360, col: 22, offset: 9436},
	expr: &ruleRefExpr{
	pos: position{line: 360, col: 22, offset: 9436},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 360, col: 32, offset: 9446},
	name: "_",
},
&litMatcher{
	pos: position{line: 360, col: 34, offset: 9448},
	val: "quantile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 360, col: 46, offset: 9460},
	name: "_",
},
&labeledExpr{
	pos: position{line: 360, col: 48, offset: 9462},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 360, col: 54, offset: 9468},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 360, col: 63, offset: 9477},
	name: "_",
},
&labeledExpr{
	pos: position{line: 360, col: 66, offset: 9480},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 360, col: 73, offset: 9487},
	expr: &seqExpr{
	pos: position{line: 360, col: 74, offset: 9488},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 360, col: 74, offset: 9488},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 360, col: 78, offset: 9492},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 360, col: 80, offset: 9494},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 360, col: 91, offset: 9505},
	name: "_",
},
&litMatcher{
	pos: position{line: 360, col: 93, offset: 9507},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 360, col: 97, offset: 9511},
	name: "_",
},
&labeledExpr{
	pos: position{line: 360, col: 99, offset: 9513},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 360, col: 103, offset: 9517},
	name: "Float",
},
},
&ruleRefExpr{
	pos: position{line: 360, col: 109, offset: 9523},
	name: "_",
},
&labeledExpr{
	pos: position{line: 360, col: 111, offset: 9525},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 360, col: 116, offset: 9530},
	expr: &seqExpr{
	pos: position{line: 360, col: 117, offset: 9531},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 360, col: 117, offset: 9531},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 360, col: 121, offset: 9535},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 360, col: 123, offset: 9537},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 360, col: 137, offset: 9551},
	name: "_",
},
&litMatcher{
	pos: position{line: 360, col: 139, offset: 9553},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MedianAgg",
	pos: position{line: 390, col: 1, offset: 10198},
	expr: &actionExpr{
	pos: position{line: 390, col: 14, offset: 10211},
	run: (*parser).callonMedianAgg1,
	expr: &seqExpr{
	pos: position{line: 390, col: 14, offset: 10211},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 390, col: 14, offset: 10211},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 390, col: 20, offset: 10217},
	expr: &ruleRefExpr{
	pos: position{line: 390, col: 20, offset: 10217},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 390, col: 30, offset: 10227},
	name: "_",
},
&litMatcher{
	pos: position{line: 390, col: 32, offset: 10229},
	val: "median(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 390, col: 42, offset: 10239},
	name: "_",
},
&labeledExpr{
	pos: position{line: 390, col: 44, offset: 10241},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 390, col: 50, offset: 10247},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 390, col: 59, offset: 10256},
	name: "_",
},
&labeledExpr{
	pos: position{line: 390, col: 62, offset: 10259},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 390, col: 69, offset: 10266},
	expr: &seqExpr{
	pos: position{line: 390, col: 70, offset: 10267},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 390, col: 70, offset: 10267},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 390, col: 74, offset: 10271},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 390, col: 76, offset: 10273},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 390, col: 87, offset: 10284},
	name: "_",
},
&labeledExpr{
	pos: position{line: 390, col: 89, offset: 10286},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 390, col: 94, offset: 10291},
	expr: &seqExpr{
	pos: position{line: 390, col: 95, offset: 10292},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 390, col: 95, offset: 10292},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 390, col: 99, offset: 10296},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 390, col: 101, offset: 10298},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 390, col: 115, offset: 10312},
	name: "_",
},
&litMatcher{
	pos: position{line: 390, col: 117, offset: 10314},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuartileAgg",
	pos: position{line: 422, col: 1, offset: 10991},
	expr: &actionExpr{
	pos: position{line: 422, col: 16, offset: 11006},
	run: (*parser).callonQuartileAgg1,
	expr: &seqExpr{
	pos: position{line: 422, col: 16, offset: 11006},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 422, col: 16, offset: 11006},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 422, col: 22, offset: 11012},
	expr: &ruleRefExpr{
	pos: position{line: 422, col: 22, offset: 11012},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 422, col: 32, offset: 11022},
	name: "_",
},
&litMatcher{
	pos: position{line: 422, col: 34, offset: 11024},
	val: "quartile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 422, col: 46, offset: 11036},
	name: "_",
},
&labeledExpr{
	pos: position{line: 422, col: 48, offset: 11038},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 422, col: 54, offset: 11044},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 422, col: 63, offset: 11053},
	name: "_",
},
&labeledExpr{
	pos: position{line: 422, col: 66, offset: 11056},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 422, col: 73, offset: 11063},
	expr: &seqExpr{
	pos: position{line: 422, col: 74, offset: 11064},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 422, col: 74, offset: 11064},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 422, col: 78, offset: 11068},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 422, col: 80, offset: 11070},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 422, col: 91, offset: 11081},
	name: "_",
},
&litMatcher{
	pos: position{line: 422, col: 93, offset: 11083},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 422, col: 97, offset: 11087},
	name: "_",
},
&labeledExpr{
	pos: position{line: 422, col: 99, offset: 11089},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 422, col: 103, offset: 11093},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 422, col: 110, offset: 11100},
	name: "_",
},
&labeledExpr{
	pos: position{line: 422, col: 112, offset: 11102},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 422, col: 117, offset: 11107},
	expr: &seqExpr{
	pos: position{line: 422, col: 118, offset: 11108},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 422, col: 118, offset: 11108},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 422, col: 122, offset: 11112},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 422, col: 124, offset: 11114},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 422, col: 138, offset: 11128},
	name: "_",
},
&litMatcher{
	pos: position{line: 422, col: 140, offset: 11130},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CovAgg",
	pos: position{line: 468, col: 1, offset: 12157},
	expr: &actionExpr{
	pos: position{line: 468, col: 11, offset: 12167},
	run: (*parser).callonCovAgg1,
	expr: &seqExpr{
	pos: position{line: 468, col: 11, offset: 12167},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 468, col: 11, offset: 12167},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 468, col: 17, offset: 12173},
	expr: &ruleRefExpr{
	pos: position{line: 468, col: 17, offset: 12173},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 468, col: 27, offset: 12183},
	name: "_",
},
&litMatcher{
	pos: position{line: 468, col: 29, offset: 12185},
	val: "cov(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 468, col: 36, offset: 12192},
	name: "_",
},
&labeledExpr{
	pos: position{line: 468, col: 38, offset: 12194},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 468, col: 45, offset: 12201},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 468, col: 54, offset: 12210},
	name: "_",
},
&litMatcher{
	pos: position{line: 468, col: 56, offset: 12212},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 468, col: 60, offset: 12216},
	name: "_",
},
&labeledExpr{
	pos: position{line: 468, col: 62, offset: 12218},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 468, col: 69, offset: 12225},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 468, col: 78, offset: 12234},
	name: "_",
},
&labeledExpr{
	pos: position{line: 468, col: 80, offset: 12236},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 468, col: 85, offset: 12241},
	expr: &seqExpr{
	pos: position{line: 468, col: 86, offset: 12242},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 468, col: 86, offset: 12242},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 468, col: 90, offset: 12246},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 468, col: 92, offset: 12248},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 468, col: 106, offset: 12262},
	name: "_",
},
&litMatcher{
	pos: position{line: 468, col: 108, offset: 12264},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CorrAgg",
	pos: position{line: 491, col: 1, offset: 12789},
	expr: &actionExpr{
	pos: position{line: 491, col: 12, offset: 12800},
	run: (*parser).callonCorrAgg1,
	expr: &seqExpr{
	pos: position{line: 491, col: 12, offset: 12800},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 491, col: 12, offset: 12800},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 491, col: 18, offset: 12806},
	expr: &ruleRefExpr{
	pos: position{line: 491, col: 18, offset: 12806},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 491, col: 28, offset: 12816},
	name: "_",
},
&litMatcher{
	pos: position{line: 491, col: 30, offset: 12818},
	val: "correlation(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 491, col: 45, offset: 12833},
	name: "_",
},
&labeledExpr{
	pos: position{line: 491, col: 47, offset: 12835},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 491, col: 54, offset: 12842},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 491, col: 63, offset: 12851},
	name: "_",
},
&litMatcher{
	pos: position{line: 491, col: 65, offset: 12853},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 491, col: 69, offset: 12857},
	name: "_",
},
&labeledExpr{
	pos: position{line: 491, col: 71, offset: 12859},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 491, col: 78, offset: 12866},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 491, col: 87, offset: 12875},
	name: "_",
},
&labeledExpr{
	pos: position{line: 491, col: 89, offset: 12877},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 491, col: 94, offset: 12882},
	expr: &seqExpr{
	pos: position{line: 491, col: 95, offset: 12883},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 491, col: 95, offset: 12883},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 491, col: 99, offset: 12887},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 491, col: 101, offset: 12889},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 491, col: 115, offset: 12903},
	name: "_",
},
&litMatcher{
	pos: position{line: 491, col: 117, offset: 12905},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PValueAgg",
	pos: position{line: 514, col: 1, offset: 13431},
	expr: &actionExpr{
	pos: position{line: 514, col: 14, offset: 13444},
	run: (*parser).callonPValueAgg1,
	expr: &seqExpr{
	pos: position{line: 514, col: 14, offset: 13444},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 514, col: 14, offset: 13444},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 514, col: 20, offset: 13450},
	expr: &ruleRefExpr{
	pos: position{line: 514, col: 20, offset: 13450},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 514, col: 30, offset: 13460},
	name: "_",
},
&litMatcher{
	pos: position{line: 514, col: 32, offset: 13462},
	val: "pvalue(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 514, col: 42, offset: 13472},
	name: "_",
},
&labeledExpr{
	pos: position{line: 514, col: 44, offset: 13474},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 514, col: 51, offset: 13481},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 514, col: 60, offset: 13490},
	name: "_",
},
&litMatcher{
	pos: position{line: 514, col: 62, offset: 13492},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 514, col: 66, offset: 13496},
	name: "_",
},
&labeledExpr{
	pos: position{line: 514, col: 68, offset: 13498},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 514, col: 75, offset: 13505},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 514, col: 84, offset: 13514},
	name: "_",
},
&labeledExpr{
	pos: position{line: 514, col: 86, offset: 13516},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 514, col: 91, offset: 13521},
	expr: &seqExpr{
	pos: position{line: 514, col: 92, offset: 13522},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 514, col: 92, offset: 13522},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 514, col: 96, offset: 13526},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 514, col: 98, offset: 13528},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 514, col: 112, offset: 13542},
	name: "_",
},
&litMatcher{
	pos: position{line: 514, col: 114, offset: 13544},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggGroupBy",
	pos: position{line: 539, col: 1, offset: 14069},
	expr: &actionExpr{
	pos: position{line: 539, col: 15, offset: 14083},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 539, col: 15, offset: 14083},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 539, col: 15, offset: 14083},
	name: "_",
},
&litMatcher{
	pos: position{line: 539, col: 17, offset: 14085},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 539, col: 22, offset: 14090},
	name: "_",
},
&labeledExpr{
	pos: position{line: 539, col: 24, offset: 14092},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 539, col: 30, offset: 14098},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 539, col: 39, offset: 14107},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 539, col: 44, offset: 14112},
	expr: &seqExpr{
	pos: position{line: 539, col: 45, offset: 14113},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 539, col: 45, offset: 14113},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 539, col: 49, offset: 14117},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 539, col: 51, offset: 14119},
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
	pos: position{line: 554, col: 1, offset: 14470},
	expr: &actionExpr{
	pos: position{line: 554, col: 11, offset: 14480},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 554, col: 11, offset: 14480},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 554, col: 16, offset: 14485},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 554, col: 16, offset: 14485},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 554, col: 27, offset: 14496},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 554, col: 38, offset: 14507},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 554, col: 49, offset: 14518},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 554, col: 58, offset: 14527},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 554, col: 67, offset: 14536},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 560, col: 1, offset: 14655},
	expr: &actionExpr{
	pos: position{line: 560, col: 13, offset: 14667},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 560, col: 13, offset: 14667},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 560, col: 13, offset: 14667},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 560, col: 23, offset: 14677},
	name: "_",
},
&labeledExpr{
	pos: position{line: 560, col: 25, offset: 14679},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 560, col: 28, offset: 14682},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 560, col: 42, offset: 14696},
	name: "_",
},
&litMatcher{
	pos: position{line: 560, col: 44, offset: 14698},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 567, col: 1, offset: 14893},
	expr: &actionExpr{
	pos: position{line: 567, col: 18, offset: 14910},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 567, col: 18, offset: 14910},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 567, col: 18, offset: 14910},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 567, col: 23, offset: 14915},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 567, col: 28, offset: 14920},
	name: "_",
},
&labeledExpr{
	pos: position{line: 567, col: 30, offset: 14922},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 567, col: 33, offset: 14925},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 567, col: 45, offset: 14937},
	name: "_",
},
&labeledExpr{
	pos: position{line: 567, col: 47, offset: 14939},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 567, col: 54, offset: 14946},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 571, col: 1, offset: 14997},
	expr: &ruleRefExpr{
	pos: position{line: 571, col: 16, offset: 15012},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 574, col: 1, offset: 15046},
	expr: &actionExpr{
	pos: position{line: 574, col: 13, offset: 15058},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 574, col: 13, offset: 15058},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 574, col: 13, offset: 15058},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 574, col: 23, offset: 15068},
	name: "_",
},
&labeledExpr{
	pos: position{line: 574, col: 25, offset: 15070},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 574, col: 31, offset: 15076},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 574, col: 40, offset: 15085},
	name: "_",
},
&labeledExpr{
	pos: position{line: 574, col: 42, offset: 15087},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 574, col: 47, offset: 15092},
	expr: &seqExpr{
	pos: position{line: 574, col: 48, offset: 15093},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 574, col: 48, offset: 15093},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 574, col: 52, offset: 15097},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 574, col: 54, offset: 15099},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 574, col: 65, offset: 15110},
	name: "_",
},
&litMatcher{
	pos: position{line: 574, col: 67, offset: 15112},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 589, col: 1, offset: 15449},
	expr: &actionExpr{
	pos: position{line: 589, col: 11, offset: 15459},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 589, col: 11, offset: 15459},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 589, col: 11, offset: 15459},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 589, col: 18, offset: 15466},
	name: "_",
},
&labeledExpr{
	pos: position{line: 589, col: 20, offset: 15468},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 589, col: 25, offset: 15473},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 589, col: 34, offset: 15482},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 589, col: 38, offset: 15486},
	name: "_",
},
&labeledExpr{
	pos: position{line: 589, col: 40, offset: 15488},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 589, col: 44, offset: 15492},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 589, col: 51, offset: 15499},
	name: "_",
},
&litMatcher{
	pos: position{line: 589, col: 53, offset: 15501},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 601, col: 1, offset: 15723},
	expr: &actionExpr{
	pos: position{line: 601, col: 11, offset: 15733},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 601, col: 11, offset: 15733},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 601, col: 11, offset: 15733},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 601, col: 20, offset: 15742},
	name: "_",
},
&litMatcher{
	pos: position{line: 601, col: 22, offset: 15744},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 601, col: 27, offset: 15749},
	name: "_",
},
&labeledExpr{
	pos: position{line: 601, col: 29, offset: 15751},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 601, col: 35, offset: 15757},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 607, col: 1, offset: 15863},
	expr: &actionExpr{
	pos: position{line: 607, col: 12, offset: 15874},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 607, col: 12, offset: 15874},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 607, col: 12, offset: 15874},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 607, col: 21, offset: 15883},
	name: "_",
},
&labeledExpr{
	pos: position{line: 607, col: 23, offset: 15885},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 607, col: 27, offset: 15889},
	expr: &ruleRefExpr{
	pos: position{line: 607, col: 27, offset: 15889},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 607, col: 35, offset: 15897},
	name: "_",
},
&litMatcher{
	pos: position{line: 607, col: 37, offset: 15899},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 617, col: 1, offset: 16044},
	expr: &actionExpr{
	pos: position{line: 617, col: 13, offset: 16056},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 617, col: 13, offset: 16056},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 617, col: 13, offset: 16056},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 617, col: 23, offset: 16066},
	name: "_",
},
&labeledExpr{
	pos: position{line: 617, col: 25, offset: 16068},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 617, col: 30, offset: 16073},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 617, col: 42, offset: 16085},
	name: "_",
},
&litMatcher{
	pos: position{line: 617, col: 44, offset: 16087},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 627, col: 1, offset: 16375},
	expr: &actionExpr{
	pos: position{line: 627, col: 16, offset: 16390},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 627, col: 16, offset: 16390},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 627, col: 16, offset: 16390},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 627, col: 22, offset: 16396},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 627, col: 35, offset: 16409},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 627, col: 40, offset: 16414},
	expr: &seqExpr{
	pos: position{line: 627, col: 41, offset: 16415},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 627, col: 41, offset: 16415},
	name: "_",
},
&labeledExpr{
	pos: position{line: 627, col: 43, offset: 16417},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 627, col: 46, offset: 16420},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 627, col: 58, offset: 16432},
	name: "_",
},
&labeledExpr{
	pos: position{line: 627, col: 60, offset: 16434},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 627, col: 63, offset: 16437},
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
	pos: position{line: 632, col: 1, offset: 16595},
	expr: &choiceExpr{
	pos: position{line: 632, col: 17, offset: 16611},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 632, col: 17, offset: 16611},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 632, col: 17, offset: 16611},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 632, col: 17, offset: 16611},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 632, col: 21, offset: 16615},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 632, col: 24, offset: 16618},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 632, col: 36, offset: 16630},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 634, col: 5, offset: 16661},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 634, col: 5, offset: 16661},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 634, col: 8, offset: 16664},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 638, col: 1, offset: 16706},
	expr: &choiceExpr{
	pos: position{line: 638, col: 16, offset: 16721},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 638, col: 16, offset: 16721},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 638, col: 24, offset: 16729},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 638, col: 24, offset: 16729},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 645, col: 1, offset: 16908},
	expr: &actionExpr{
	pos: position{line: 645, col: 18, offset: 16925},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 645, col: 18, offset: 16925},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 645, col: 18, offset: 16925},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 645, col: 23, offset: 16930},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 645, col: 28, offset: 16935},
	name: "_",
},
&labeledExpr{
	pos: position{line: 645, col: 30, offset: 16937},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 645, col: 33, offset: 16940},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 645, col: 45, offset: 16952},
	name: "_",
},
&labeledExpr{
	pos: position{line: 645, col: 47, offset: 16954},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 645, col: 53, offset: 16960},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 645, col: 59, offset: 16966},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 649, col: 1, offset: 17040},
	expr: &choiceExpr{
	pos: position{line: 649, col: 16, offset: 17055},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 649, col: 16, offset: 17055},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 649, col: 23, offset: 17062},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 649, col: 30, offset: 17069},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 649, col: 36, offset: 17075},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 649, col: 42, offset: 17081},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 649, col: 49, offset: 17088},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 649, col: 49, offset: 17088},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 658, col: 1, offset: 17198},
	expr: &actionExpr{
	pos: position{line: 658, col: 12, offset: 17209},
	run: (*parser).callonStdOutG1,
	expr: &seqExpr{
	pos: position{line: 658, col: 12, offset: 17209},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 658, col: 12, offset: 17209},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 658, col: 22, offset: 17219},
	name: "_",
},
&labeledExpr{
	pos: position{line: 658, col: 24, offset: 17221},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 658, col: 31, offset: 17228},
	expr: &ruleRefExpr{
	pos: position{line: 658, col: 32, offset: 17229},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 658, col: 48, offset: 17245},
	name: "_",
},
&litMatcher{
	pos: position{line: 658, col: 50, offset: 17247},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SinkInto",
	pos: position{line: 672, col: 1, offset: 17526},
	expr: &actionExpr{
	pos: position{line: 672, col: 12, offset: 17537},
	run: (*parser).callonSinkInto1,
	expr: &seqExpr{
	pos: position{line: 672, col: 12, offset: 17537},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 672, col: 12, offset: 17537},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 672, col: 18, offset: 17543},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 672, col: 27, offset: 17552},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 672, col: 33, offset: 17558},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 672, col: 33, offset: 17558},
	name: "_",
},
&litMatcher{
	pos: position{line: 672, col: 35, offset: 17560},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 672, col: 39, offset: 17564},
	name: "_",
},
&litMatcher{
	pos: position{line: 672, col: 41, offset: 17566},
	val: "into()",
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
	name: "BlackHoleG",
	pos: position{line: 680, col: 1, offset: 17722},
	expr: &actionExpr{
	pos: position{line: 680, col: 15, offset: 17736},
	run: (*parser).callonBlackHoleG1,
	expr: &seqExpr{
	pos: position{line: 680, col: 15, offset: 17736},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 680, col: 15, offset: 17736},
	val: "blackhole(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 680, col: 28, offset: 17749},
	name: "_",
},
&labeledExpr{
	pos: position{line: 680, col: 30, offset: 17751},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 680, col: 37, offset: 17758},
	expr: &ruleRefExpr{
	pos: position{line: 680, col: 38, offset: 17759},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 680, col: 54, offset: 17775},
	name: "_",
},
&litMatcher{
	pos: position{line: 680, col: 56, offset: 17777},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PlotG",
	pos: position{line: 684, col: 1, offset: 17835},
	expr: &actionExpr{
	pos: position{line: 684, col: 10, offset: 17844},
	run: (*parser).callonPlotG1,
	expr: &seqExpr{
	pos: position{line: 684, col: 10, offset: 17844},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 684, col: 10, offset: 17844},
	val: "plot",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 684, col: 17, offset: 17851},
	name: "_",
},
&labeledExpr{
	pos: position{line: 684, col: 19, offset: 17853},
	label: "widget",
	expr: &ruleRefExpr{
	pos: position{line: 684, col: 27, offset: 17861},
	name: "BarWidget",
},
},
	},
},
},
},
{
	name: "BarWidget",
	pos: position{line: 688, col: 1, offset: 17961},
	expr: &actionExpr{
	pos: position{line: 688, col: 14, offset: 17974},
	run: (*parser).callonBarWidget1,
	expr: &seqExpr{
	pos: position{line: 688, col: 14, offset: 17974},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 688, col: 14, offset: 17974},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 688, col: 20, offset: 17980},
	expr: &ruleRefExpr{
	pos: position{line: 688, col: 20, offset: 17980},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 688, col: 30, offset: 17990},
	name: "_",
},
&litMatcher{
	pos: position{line: 688, col: 32, offset: 17992},
	val: "bar(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 688, col: 39, offset: 17999},
	name: "_",
},
&labeledExpr{
	pos: position{line: 688, col: 41, offset: 18001},
	label: "xField",
	expr: &ruleRefExpr{
	pos: position{line: 688, col: 48, offset: 18008},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 688, col: 57, offset: 18017},
	name: "_",
},
&litMatcher{
	pos: position{line: 688, col: 59, offset: 18019},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 688, col: 63, offset: 18023},
	name: "_",
},
&labeledExpr{
	pos: position{line: 688, col: 65, offset: 18025},
	label: "yField",
	expr: &ruleRefExpr{
	pos: position{line: 688, col: 72, offset: 18032},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 688, col: 81, offset: 18041},
	name: "_",
},
&labeledExpr{
	pos: position{line: 688, col: 83, offset: 18043},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 688, col: 88, offset: 18048},
	expr: &seqExpr{
	pos: position{line: 688, col: 89, offset: 18049},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 688, col: 89, offset: 18049},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 688, col: 93, offset: 18053},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 688, col: 95, offset: 18055},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 688, col: 102, offset: 18062},
	name: "_",
},
&zeroOrOneExpr{
	pos: position{line: 688, col: 104, offset: 18064},
	expr: &seqExpr{
	pos: position{line: 688, col: 105, offset: 18065},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 688, col: 105, offset: 18065},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 688, col: 109, offset: 18069},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 688, col: 111, offset: 18071},
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
	pos: position{line: 688, col: 122, offset: 18082},
	name: "_",
},
&litMatcher{
	pos: position{line: 688, col: 124, offset: 18084},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FileName",
	pos: position{line: 721, col: 1, offset: 18769},
	expr: &actionExpr{
	pos: position{line: 721, col: 13, offset: 18781},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 721, col: 13, offset: 18781},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 721, col: 13, offset: 18781},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 721, col: 22, offset: 18790},
	expr: &seqExpr{
	pos: position{line: 721, col: 23, offset: 18791},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 721, col: 23, offset: 18791},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 721, col: 27, offset: 18795},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 721, col: 31, offset: 18799},
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
	pos: position{line: 727, col: 1, offset: 18900},
	expr: &actionExpr{
	pos: position{line: 727, col: 10, offset: 18909},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 727, col: 10, offset: 18909},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 727, col: 10, offset: 18909},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 727, col: 16, offset: 18915},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 727, col: 23, offset: 18922},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 727, col: 28, offset: 18927},
	expr: &seqExpr{
	pos: position{line: 727, col: 29, offset: 18928},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 727, col: 29, offset: 18928},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 727, col: 31, offset: 18930},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 727, col: 36, offset: 18935},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 727, col: 38, offset: 18937},
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
	pos: position{line: 731, col: 1, offset: 19002},
	expr: &actionExpr{
	pos: position{line: 731, col: 11, offset: 19012},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 731, col: 11, offset: 19012},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 731, col: 16, offset: 19017},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 731, col: 16, offset: 19017},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 731, col: 24, offset: 19025},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 731, col: 33, offset: 19034},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 731, col: 48, offset: 19049},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 731, col: 59, offset: 19060},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 735, col: 1, offset: 19102},
	expr: &choiceExpr{
	pos: position{line: 735, col: 9, offset: 19110},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 735, col: 9, offset: 19110},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 735, col: 15, offset: 19116},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 735, col: 21, offset: 19122},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 735, col: 27, offset: 19128},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 735, col: 33, offset: 19134},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 735, col: 33, offset: 19134},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 739, col: 1, offset: 19182},
	expr: &choiceExpr{
	pos: position{line: 739, col: 12, offset: 19193},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 739, col: 12, offset: 19193},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 739, col: 12, offset: 19193},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 739, col: 12, offset: 19193},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 739, col: 16, offset: 19197},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 739, col: 19, offset: 19200},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 739, col: 25, offset: 19206},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 743, col: 5, offset: 19282},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 743, col: 5, offset: 19282},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 743, col: 8, offset: 19285},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 750, col: 1, offset: 19362},
	expr: &actionExpr{
	pos: position{line: 750, col: 9, offset: 19370},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 750, col: 9, offset: 19370},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 750, col: 9, offset: 19370},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 750, col: 15, offset: 19376},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 750, col: 21, offset: 19382},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 750, col: 26, offset: 19387},
	expr: &seqExpr{
	pos: position{line: 750, col: 27, offset: 19388},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 750, col: 27, offset: 19388},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 750, col: 29, offset: 19390},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 750, col: 34, offset: 19395},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 750, col: 36, offset: 19397},
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
	pos: position{line: 754, col: 1, offset: 19461},
	expr: &actionExpr{
	pos: position{line: 754, col: 10, offset: 19470},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 754, col: 10, offset: 19470},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 754, col: 15, offset: 19475},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 754, col: 15, offset: 19475},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 754, col: 23, offset: 19483},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 754, col: 32, offset: 19492},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 754, col: 47, offset: 19507},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 758, col: 1, offset: 19550},
	expr: &actionExpr{
	pos: position{line: 758, col: 13, offset: 19562},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 758, col: 13, offset: 19562},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 758, col: 13, offset: 19562},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 758, col: 20, offset: 19569},
	expr: &choiceExpr{
	pos: position{line: 758, col: 21, offset: 19570},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 758, col: 21, offset: 19570},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 758, col: 31, offset: 19580},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 758, col: 39, offset: 19588},
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
	pos: position{line: 762, col: 1, offset: 19638},
	expr: &actionExpr{
	pos: position{line: 762, col: 17, offset: 19654},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 762, col: 17, offset: 19654},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 762, col: 17, offset: 19654},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 762, col: 21, offset: 19658},
	expr: &choiceExpr{
	pos: position{line: 762, col: 23, offset: 19660},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 762, col: 23, offset: 19660},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 762, col: 23, offset: 19660},
	expr: &ruleRefExpr{
	pos: position{line: 762, col: 24, offset: 19661},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 762, col: 36, offset: 19673,
},
	},
},
&seqExpr{
	pos: position{line: 762, col: 40, offset: 19677},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 762, col: 40, offset: 19677},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 762, col: 45, offset: 19682},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 762, col: 63, offset: 19700},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 769, col: 1, offset: 19881},
	expr: &charClassMatcher{
	pos: position{line: 769, col: 16, offset: 19896},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 771, col: 1, offset: 19914},
	expr: &ruleRefExpr{
	pos: position{line: 771, col: 19, offset: 19932},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 773, col: 1, offset: 19952},
	expr: &charClassMatcher{
	pos: position{line: 773, col: 21, offset: 19972},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 775, col: 1, offset: 19987},
	expr: &actionExpr{
	pos: position{line: 775, col: 9, offset: 19995},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 775, col: 9, offset: 19995},
	expr: &choiceExpr{
	pos: position{line: 775, col: 10, offset: 19996},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 775, col: 10, offset: 19996},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 775, col: 19, offset: 20005},
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
	pos: position{line: 779, col: 1, offset: 20055},
	expr: &choiceExpr{
	pos: position{line: 779, col: 11, offset: 20065},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 779, col: 11, offset: 20065},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 779, col: 19, offset: 20073},
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
	pos: position{line: 781, col: 1, offset: 20082},
	expr: &actionExpr{
	pos: position{line: 781, col: 11, offset: 20092},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 781, col: 11, offset: 20092},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 781, col: 11, offset: 20092},
	expr: &litMatcher{
	pos: position{line: 781, col: 11, offset: 20092},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 781, col: 16, offset: 20097},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 785, col: 1, offset: 20170},
	expr: &actionExpr{
	pos: position{line: 785, col: 19, offset: 20188},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 785, col: 19, offset: 20188},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 789, col: 1, offset: 20261},
	expr: &actionExpr{
	pos: position{line: 789, col: 10, offset: 20270},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 789, col: 10, offset: 20270},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 789, col: 10, offset: 20270},
	expr: &litMatcher{
	pos: position{line: 789, col: 10, offset: 20270},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 789, col: 15, offset: 20275},
	expr: &ruleRefExpr{
	pos: position{line: 789, col: 15, offset: 20275},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 789, col: 24, offset: 20284},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 789, col: 28, offset: 20288},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 793, col: 1, offset: 20355},
	expr: &choiceExpr{
	pos: position{line: 793, col: 12, offset: 20366},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 793, col: 12, offset: 20366},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 793, col: 18, offset: 20372},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 793, col: 18, offset: 20372},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 793, col: 38, offset: 20392},
	expr: &ruleRefExpr{
	pos: position{line: 793, col: 38, offset: 20392},
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
	pos: position{line: 795, col: 1, offset: 20409},
	expr: &charClassMatcher{
	pos: position{line: 795, col: 17, offset: 20425},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 797, col: 1, offset: 20434},
	expr: &charClassMatcher{
	pos: position{line: 797, col: 24, offset: 20457},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 799, col: 1, offset: 20466},
	expr: &zeroOrMoreExpr{
	pos: position{line: 799, col: 19, offset: 20484},
	expr: &charClassMatcher{
	pos: position{line: 799, col: 19, offset: 20484},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 801, col: 1, offset: 20498},
	expr: &charClassMatcher{
	pos: position{line: 801, col: 18, offset: 20515},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "IDENTIFIER",
	pos: position{line: 802, col: 1, offset: 20526},
	expr: &actionExpr{
	pos: position{line: 802, col: 15, offset: 20540},
	run: (*parser).callonIDENTIFIER1,
	expr: &seqExpr{
	pos: position{line: 802, col: 15, offset: 20540},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 802, col: 15, offset: 20540},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 802, col: 22, offset: 20547},
	expr: &choiceExpr{
	pos: position{line: 802, col: 23, offset: 20548},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 802, col: 23, offset: 20548},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 802, col: 33, offset: 20558},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 802, col: 41, offset: 20566},
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
	name: "EOF",
	pos: position{line: 805, col: 1, offset: 20639},
	expr: &notExpr{
	pos: position{line: 805, col: 8, offset: 20646},
	expr: &anyMatcher{
	line: 805, col: 9, offset: 20647,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 807, col: 1, offset: 20652},
	expr: &litMatcher{
	pos: position{line: 807, col: 17, offset: 20668},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 813, col: 1, offset: 20755},
	expr: &actionExpr{
	pos: position{line: 813, col: 18, offset: 20772},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 813, col: 18, offset: 20772},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 813, col: 18, offset: 20772},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 813, col: 24, offset: 20778},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 813, col: 33, offset: 20787},
	name: "_",
},
&labeledExpr{
	pos: position{line: 813, col: 35, offset: 20789},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 813, col: 40, offset: 20794},
	expr: &seqExpr{
	pos: position{line: 813, col: 41, offset: 20795},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 813, col: 41, offset: 20795},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 813, col: 45, offset: 20799},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 813, col: 47, offset: 20801},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 813, col: 56, offset: 20810},
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
{
	name: "EqAliasG",
	pos: position{line: 816, col: 1, offset: 20863},
	expr: &actionExpr{
	pos: position{line: 816, col: 13, offset: 20875},
	run: (*parser).callonEqAliasG1,
	expr: &seqExpr{
	pos: position{line: 816, col: 13, offset: 20875},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 816, col: 13, offset: 20875},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 816, col: 19, offset: 20881},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 816, col: 28, offset: 20890},
	name: "_",
},
&litMatcher{
	pos: position{line: 816, col: 30, offset: 20892},
	val: "=",
	ignoreCase: false,
},
	},
},
},
},
	},
}
func (c *current) onCommand1(command, rest interface{}) (interface{}, error) {

        commands := []interface{}{(command.([]interface{}))}
        for _,val := range(cast.ToIfaceSlice(rest)){
        		_,ok := val.([]interface{})
        		if ok{
        			newVal := val.([]interface{})
        			if len(newVal)>0{
        				for _,item := range(newVal){
        					_,ok := item.([]interface{})
        					if ok{
        						for _,subItem := range cast.ToIfaceSlice(item){
        							_,ok := subItem.([]interface{})
        							if ok{
        								castedSubItem := cast.ToIfaceSlice(subItem)
        								if len(castedSubItem)>0{
        										commands = append(commands,castedSubItem)

    									}
    								}
    							}
    						}
    					}

        			}
        		}
        	}
            return commands,nil
    
}

func (p *parser) callonCommand1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCommand1(stack["command"], stack["rest"])
}

func (c *current) onStatement1(statement, rest interface{}) (interface{}, error) {

    stages := []interface{}{statement}
        	rst := cast.ToIfaceSlice(rest)

        	for _,val := range(rst){

        		_,ok := val.([]interface{})
        		if ok {
        			newVal := cast.ToIfaceSlice(val)
        			for _,item := range newVal{
    					_, success := item.([]interface{})
    					if success{
    						newItem := cast.ToIfaceSlice(item)

    						if len(newItem) == 4 {
    							stages = append(stages, newItem[3])
    						}
    					}
    				}
        		}else{
        			_,ok:= val.([]uint8)
        			if !ok{
        				stages = append(stages,val)

        			}

        		}

        	}

        	return stages,nil
    
}

func (p *parser) callonStatement1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement1(stack["statement"], stack["rest"])
}

func (c *current) onFileJobG1(filename interface{}) (interface{}, error) {

    if fname, ok := cast.TryString(filename); ok {
        return SourceJob{Type: FILEJOB,OperateOn:SrcFile{FileName:fname}},nil
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
       return SourceJob{Type:FAKEJOB,OperateOn:SrcFake{Number:n}},nil
    }
    return nil, errors.New("could not decode source file name")
}

func (p *parser) callonFakeJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFakeJobG1(stack["numData"])
}

func (c *current) onBranchJobG1(br interface{}) (interface{}, error) {

    if branchName,ok := cast.TryString(br);ok{
    return SourceJob{Type:BRANCHJOB,OperateOn:SrcBranch{Identifier:branchName}},nil
    }
    return nil,errors.New("could not identify the branch name")
}

func (p *parser) callonBranchJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBranchJobG1(stack["br"])
}

func (c *current) onSecSource1(src1, rest interface{}) (interface{}, error) {

    var secSources []string
     if  source1,ok := cast.TryString(src1);ok{
            secSources = append(secSources,source1)
        }
     if rest!= nil{
    		restVal := cast.ToIfaceSlice(rest)
    		source2 := restVal[len(restVal)-1]
    		secSources = append(secSources,source2.(string))
    	}

    return SourceJob{Type:SECSOURCE,OperateOn:SrcSec{Sources:secSources}},nil
}

func (p *parser) callonSecSource1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSecSource1(stack["src1"], stack["rest"])
}

func (c *current) onJoinJobG1(job interface{}) (interface{}, error) {

    return TransformJob{Type:JOINJOB,OperateOn:job},nil
}

func (p *parser) callonJoinJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onJoinJobG1(stack["job"])
}

func (c *current) onInnerJoinG1(alias, query interface{}) (interface{}, error) {

     name,err := parseJoinArgs(alias)
     if err !=nil{
         return nil,err
         }
     return JoinNodeJob{Alias:name,Type:"INNER",Attributes:query.(JoinAttributes)},nil
}

func (p *parser) callonInnerJoinG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInnerJoinG1(stack["alias"], stack["query"])
}

func (c *current) onQuery1(field, condition interface{}) (interface{}, error) {

    return JoinAttributes{SelectFields:field.([]string),JoinCondition:condition.(JoinConditions)},nil
}

func (p *parser) callonQuery1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQuery1(stack["field"], stack["condition"])
}

func (c *current) onAllFields1() (interface{}, error) {

    var fields []string
    fields = append(fields,"*")
    return fields,nil
}

func (p *parser) callonAllFields1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAllFields1()
}

func (c *current) onSelFields1(first, rest interface{}) (interface{}, error) {

    var fields []string
    fields = append(fields,first.(string))
    for _,val := range rest.([]interface{}){
        castedVal := val.([]interface{})
        fields = append(fields,castedVal[2].(string))
    }
    return fields,nil
}

func (p *parser) callonSelFields1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSelFields1(stack["first"], stack["rest"])
}

func (c *current) onJoinCondition1(left, op, right interface{}) (interface{}, error) {

        return JoinConditions{LeftFields:left.([]string),RightFields:right.([]string),JoinOperator:"=="},nil
}

func (p *parser) callonJoinCondition1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onJoinCondition1(stack["left"], stack["op"], stack["right"])
}

func (c *current) onJoinFactors1(first, rest interface{}) (interface{}, error) {

    var factors []string
    factors = append(factors,first.(string))
    for _,factor := range rest.([]interface{}){
        castedVal := factor.([]interface{})
        if len(castedVal) >=4{
            factors = append(factors,castedVal[2].(string))
        }
    }
return factors,nil
}

func (p *parser) callonJoinFactors1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onJoinFactors1(stack["first"], stack["rest"])
}

func (c *current) onAggJobG1(params, job, rest interface{}) (interface{}, error) {

        jobSlices := cast.ToIfaceSlice(rest)
        jobs := []interface{}{job}

        for _, v := range jobSlices {
                j := cast.ToIfaceSlice(v)
                jobs = append(jobs, j[2])
        }

        ps, _ := params.([]string)
        return TransformJob{Type:AGGJOB,OperateOn:AggNodeJob{Functions:jobs,GroupBy: ps}},nil

}

func (p *parser) callonAggJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAggJobG1(stack["params"], stack["job"], stack["rest"])
}

func (c *current) onCountAgg1(alias, args interface{}) (interface{}, error) {

    name,filter,err := parseAggArgs(alias,args)
    if err !=nil{
    return nil,err
    }
    return Count{Alias:name,Filter:filter},nil
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

func (c *current) onDoJobG1(job interface{}) (interface{}, error) {

        return TransformJob{Type:DOJOB,OperateOn:DoNodeJob{Function: job}}, nil
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

func (c *current) onStdOutG1(params interface{}) (interface{}, error) {


    if header, ok := params.([]string); ok {
        return SinkJob{
            Type: STDOUT,
            OperateOn:StdOut{Header:header},
            },nil
    }
    return SinkJob{
        Type: STDOUT,
        OperateOn:StdOut{Header:nil},
    }, nil
}

func (p *parser) callonStdOutG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStdOutG1(stack["params"])
}

func (c *current) onSinkInto1(first, rest interface{}) (interface{}, error) {

    streamTo,_ := first.(string)
    return SinkJob{
        Type: INTO,
        OperateOn:InTo{StreamTo:streamTo},
        },nil
    
}

func (p *parser) callonSinkInto1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSinkInto1(stack["first"], stack["rest"])
}

func (c *current) onBlackHoleG1(params interface{}) (interface{}, error) {

        return SinkJob{Type: BLACKHOLE}, nil
}

func (p *parser) callonBlackHoleG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlackHoleG1(stack["params"])
}

func (c *current) onPlotG1(widget interface{}) (interface{}, error) {

    return SinkJob{Type: PLOT, OperateOn:Plot{Type:"plot",Widget: widget}}, nil
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

func (c *current) onIDENTIFIER1() (interface{}, error) {

                   return string(c.text), nil
              
}

func (p *parser) callonIDENTIFIER1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIDENTIFIER1()
}

func (c *current) onVarParamListG1(first, rest interface{}) (interface{}, error) {

    return getParamList(first, rest), nil
}

func (p *parser) callonVarParamListG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVarParamListG1(stack["first"], stack["rest"])
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

