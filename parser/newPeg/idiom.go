
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
},
{
	name: "Statement",
	pos: position{line: 36, col: 1, offset: 901},
	expr: &actionExpr{
	pos: position{line: 36, col: 14, offset: 914},
	run: (*parser).callonStatement1,
	expr: &seqExpr{
	pos: position{line: 36, col: 14, offset: 914},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 36, col: 14, offset: 914},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 36, col: 24, offset: 924},
	name: "Source",
},
},
&labeledExpr{
	pos: position{line: 36, col: 31, offset: 931},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 36, col: 37, offset: 937},
	exprs: []interface{}{
&zeroOrMoreExpr{
	pos: position{line: 36, col: 37, offset: 937},
	expr: &seqExpr{
	pos: position{line: 36, col: 38, offset: 938},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 36, col: 38, offset: 938},
	name: "_",
},
&litMatcher{
	pos: position{line: 36, col: 40, offset: 940},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 36, col: 44, offset: 944},
	name: "_",
},
&labeledExpr{
	pos: position{line: 36, col: 46, offset: 946},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 36, col: 56, offset: 956},
	name: "Transform",
},
},
	},
},
},
&ruleRefExpr{
	pos: position{line: 36, col: 68, offset: 968},
	name: "_",
},
&litMatcher{
	pos: position{line: 36, col: 69, offset: 969},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 36, col: 72, offset: 972},
	name: "_",
},
&labeledExpr{
	pos: position{line: 36, col: 74, offset: 974},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 36, col: 84, offset: 984},
	name: "Sink",
},
},
&ruleRefExpr{
	pos: position{line: 36, col: 89, offset: 989},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 36, col: 91, offset: 991},
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
	name: "Source",
	pos: position{line: 71, col: 1, offset: 1662},
	expr: &choiceExpr{
	pos: position{line: 71, col: 11, offset: 1672},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 71, col: 11, offset: 1672},
	name: "BranchJobG",
},
&ruleRefExpr{
	pos: position{line: 71, col: 22, offset: 1683},
	name: "FileJobG",
},
	},
},
},
{
	name: "Transform",
	pos: position{line: 73, col: 1, offset: 1695},
	expr: &choiceExpr{
	pos: position{line: 73, col: 14, offset: 1708},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 73, col: 14, offset: 1708},
	name: "AggJobG",
},
&ruleRefExpr{
	pos: position{line: 73, col: 22, offset: 1716},
	name: "DoJobG",
},
	},
},
},
{
	name: "Sink",
	pos: position{line: 75, col: 1, offset: 1726},
	expr: &choiceExpr{
	pos: position{line: 75, col: 9, offset: 1734},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 75, col: 9, offset: 1734},
	name: "StdOutG",
},
&ruleRefExpr{
	pos: position{line: 75, col: 17, offset: 1742},
	name: "SinkInto",
},
	},
},
},
{
	name: "FileJobG",
	pos: position{line: 79, col: 1, offset: 1771},
	expr: &actionExpr{
	pos: position{line: 79, col: 13, offset: 1783},
	run: (*parser).callonFileJobG1,
	expr: &seqExpr{
	pos: position{line: 79, col: 13, offset: 1783},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 79, col: 13, offset: 1783},
	name: "_",
},
&litMatcher{
	pos: position{line: 79, col: 15, offset: 1785},
	val: "file(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 79, col: 23, offset: 1793},
	name: "_",
},
&labeledExpr{
	pos: position{line: 79, col: 25, offset: 1795},
	label: "filename",
	expr: &ruleRefExpr{
	pos: position{line: 79, col: 34, offset: 1804},
	name: "FileName",
},
},
&ruleRefExpr{
	pos: position{line: 79, col: 43, offset: 1813},
	name: "_",
},
&litMatcher{
	pos: position{line: 79, col: 45, offset: 1815},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchJobG",
	pos: position{line: 86, col: 1, offset: 1997},
	expr: &actionExpr{
	pos: position{line: 86, col: 15, offset: 2011},
	run: (*parser).callonBranchJobG1,
	expr: &seqExpr{
	pos: position{line: 86, col: 15, offset: 2011},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 86, col: 15, offset: 2011},
	val: "branch",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 86, col: 23, offset: 2019},
	name: "_",
},
&litMatcher{
	pos: position{line: 86, col: 24, offset: 2020},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 86, col: 28, offset: 2024},
	name: "_",
},
&labeledExpr{
	pos: position{line: 86, col: 30, offset: 2026},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 86, col: 33, offset: 2029},
	name: "IDENTIFIER",
},
},
&litMatcher{
	pos: position{line: 86, col: 44, offset: 2040},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggJobG",
	pos: position{line: 99, col: 1, offset: 2306},
	expr: &actionExpr{
	pos: position{line: 99, col: 12, offset: 2317},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 99, col: 12, offset: 2317},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 99, col: 12, offset: 2317},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 99, col: 18, offset: 2323},
	name: "_",
},
&labeledExpr{
	pos: position{line: 99, col: 20, offset: 2325},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 99, col: 27, offset: 2332},
	expr: &ruleRefExpr{
	pos: position{line: 99, col: 27, offset: 2332},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 99, col: 39, offset: 2344},
	name: "_",
},
&labeledExpr{
	pos: position{line: 99, col: 41, offset: 2346},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 99, col: 45, offset: 2350},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 99, col: 53, offset: 2358},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 99, col: 58, offset: 2363},
	expr: &seqExpr{
	pos: position{line: 99, col: 59, offset: 2364},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 99, col: 59, offset: 2364},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 99, col: 63, offset: 2368},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 99, col: 65, offset: 2370},
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
	pos: position{line: 114, col: 1, offset: 2712},
	expr: &ruleRefExpr{
	pos: position{line: 114, col: 12, offset: 2723},
	name: "CountAgg",
},
},
{
	name: "CountAgg",
	pos: position{line: 116, col: 1, offset: 2735},
	expr: &actionExpr{
	pos: position{line: 116, col: 13, offset: 2747},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 116, col: 13, offset: 2747},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 116, col: 13, offset: 2747},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 116, col: 19, offset: 2753},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 116, col: 28, offset: 2762},
	name: "_",
},
&litMatcher{
	pos: position{line: 116, col: 30, offset: 2764},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 116, col: 39, offset: 2773},
	name: "_",
},
&labeledExpr{
	pos: position{line: 116, col: 41, offset: 2775},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 116, col: 46, offset: 2780},
	expr: &ruleRefExpr{
	pos: position{line: 116, col: 46, offset: 2780},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 116, col: 59, offset: 2793},
	name: "_",
},
&litMatcher{
	pos: position{line: 116, col: 61, offset: 2795},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggGroupBy",
	pos: position{line: 125, col: 1, offset: 2951},
	expr: &actionExpr{
	pos: position{line: 125, col: 15, offset: 2965},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 125, col: 15, offset: 2965},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 125, col: 15, offset: 2965},
	name: "_",
},
&litMatcher{
	pos: position{line: 125, col: 17, offset: 2967},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 125, col: 22, offset: 2972},
	name: "_",
},
&labeledExpr{
	pos: position{line: 125, col: 24, offset: 2974},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 125, col: 30, offset: 2980},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 125, col: 39, offset: 2989},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 125, col: 44, offset: 2994},
	expr: &seqExpr{
	pos: position{line: 125, col: 45, offset: 2995},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 125, col: 45, offset: 2995},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 125, col: 49, offset: 2999},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 125, col: 51, offset: 3001},
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
	pos: position{line: 142, col: 1, offset: 3356},
	expr: &actionExpr{
	pos: position{line: 142, col: 11, offset: 3366},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 142, col: 11, offset: 3366},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 142, col: 16, offset: 3371},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 142, col: 16, offset: 3371},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 142, col: 27, offset: 3382},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 142, col: 38, offset: 3393},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 142, col: 49, offset: 3404},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 142, col: 58, offset: 3413},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 142, col: 67, offset: 3422},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 148, col: 1, offset: 3506},
	expr: &actionExpr{
	pos: position{line: 148, col: 13, offset: 3518},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 148, col: 13, offset: 3518},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 148, col: 13, offset: 3518},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 148, col: 23, offset: 3528},
	name: "_",
},
&labeledExpr{
	pos: position{line: 148, col: 25, offset: 3530},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 148, col: 28, offset: 3533},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 148, col: 42, offset: 3547},
	name: "_",
},
&litMatcher{
	pos: position{line: 148, col: 44, offset: 3549},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 155, col: 1, offset: 3744},
	expr: &actionExpr{
	pos: position{line: 155, col: 18, offset: 3761},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 155, col: 18, offset: 3761},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 155, col: 18, offset: 3761},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 155, col: 23, offset: 3766},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 155, col: 28, offset: 3771},
	name: "_",
},
&labeledExpr{
	pos: position{line: 155, col: 30, offset: 3773},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 155, col: 33, offset: 3776},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 155, col: 45, offset: 3788},
	name: "_",
},
&labeledExpr{
	pos: position{line: 155, col: 47, offset: 3790},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 155, col: 54, offset: 3797},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 159, col: 1, offset: 3848},
	expr: &ruleRefExpr{
	pos: position{line: 159, col: 16, offset: 3863},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 162, col: 1, offset: 3897},
	expr: &actionExpr{
	pos: position{line: 162, col: 13, offset: 3909},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 162, col: 13, offset: 3909},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 162, col: 13, offset: 3909},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 162, col: 23, offset: 3919},
	name: "_",
},
&labeledExpr{
	pos: position{line: 162, col: 25, offset: 3921},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 162, col: 31, offset: 3927},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 162, col: 40, offset: 3936},
	name: "_",
},
&labeledExpr{
	pos: position{line: 162, col: 42, offset: 3938},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 162, col: 47, offset: 3943},
	expr: &seqExpr{
	pos: position{line: 162, col: 48, offset: 3944},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 162, col: 48, offset: 3944},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 162, col: 52, offset: 3948},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 162, col: 54, offset: 3950},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 162, col: 65, offset: 3961},
	name: "_",
},
&litMatcher{
	pos: position{line: 162, col: 67, offset: 3963},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 177, col: 1, offset: 4300},
	expr: &actionExpr{
	pos: position{line: 177, col: 11, offset: 4310},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 177, col: 11, offset: 4310},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 177, col: 11, offset: 4310},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 177, col: 18, offset: 4317},
	name: "_",
},
&labeledExpr{
	pos: position{line: 177, col: 20, offset: 4319},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 177, col: 25, offset: 4324},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 177, col: 34, offset: 4333},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 177, col: 38, offset: 4337},
	name: "_",
},
&labeledExpr{
	pos: position{line: 177, col: 40, offset: 4339},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 177, col: 44, offset: 4343},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 177, col: 51, offset: 4350},
	name: "_",
},
&litMatcher{
	pos: position{line: 177, col: 53, offset: 4352},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 189, col: 1, offset: 4574},
	expr: &actionExpr{
	pos: position{line: 189, col: 11, offset: 4584},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 189, col: 11, offset: 4584},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 189, col: 11, offset: 4584},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 189, col: 20, offset: 4593},
	name: "_",
},
&litMatcher{
	pos: position{line: 189, col: 22, offset: 4595},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 189, col: 27, offset: 4600},
	name: "_",
},
&labeledExpr{
	pos: position{line: 189, col: 29, offset: 4602},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 189, col: 35, offset: 4608},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 195, col: 1, offset: 4714},
	expr: &actionExpr{
	pos: position{line: 195, col: 12, offset: 4725},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 195, col: 12, offset: 4725},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 195, col: 12, offset: 4725},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 195, col: 21, offset: 4734},
	name: "_",
},
&labeledExpr{
	pos: position{line: 195, col: 23, offset: 4736},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 195, col: 27, offset: 4740},
	expr: &ruleRefExpr{
	pos: position{line: 195, col: 27, offset: 4740},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 195, col: 35, offset: 4748},
	name: "_",
},
&litMatcher{
	pos: position{line: 195, col: 37, offset: 4750},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 205, col: 1, offset: 4895},
	expr: &actionExpr{
	pos: position{line: 205, col: 13, offset: 4907},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 205, col: 13, offset: 4907},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 205, col: 13, offset: 4907},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 205, col: 23, offset: 4917},
	name: "_",
},
&labeledExpr{
	pos: position{line: 205, col: 25, offset: 4919},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 205, col: 30, offset: 4924},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 205, col: 42, offset: 4936},
	name: "_",
},
&litMatcher{
	pos: position{line: 205, col: 44, offset: 4938},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 215, col: 1, offset: 5226},
	expr: &actionExpr{
	pos: position{line: 215, col: 16, offset: 5241},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 215, col: 16, offset: 5241},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 215, col: 16, offset: 5241},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 215, col: 22, offset: 5247},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 215, col: 35, offset: 5260},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 215, col: 40, offset: 5265},
	expr: &seqExpr{
	pos: position{line: 215, col: 41, offset: 5266},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 215, col: 41, offset: 5266},
	name: "_",
},
&labeledExpr{
	pos: position{line: 215, col: 43, offset: 5268},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 215, col: 46, offset: 5271},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 215, col: 58, offset: 5283},
	name: "_",
},
&labeledExpr{
	pos: position{line: 215, col: 60, offset: 5285},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 215, col: 63, offset: 5288},
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
	pos: position{line: 220, col: 1, offset: 5446},
	expr: &choiceExpr{
	pos: position{line: 220, col: 17, offset: 5462},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 220, col: 17, offset: 5462},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 220, col: 17, offset: 5462},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 220, col: 17, offset: 5462},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 220, col: 21, offset: 5466},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 220, col: 24, offset: 5469},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 220, col: 36, offset: 5481},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 222, col: 5, offset: 5512},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 222, col: 5, offset: 5512},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 222, col: 8, offset: 5515},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 226, col: 1, offset: 5557},
	expr: &choiceExpr{
	pos: position{line: 226, col: 16, offset: 5572},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 226, col: 16, offset: 5572},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 226, col: 24, offset: 5580},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 226, col: 24, offset: 5580},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 233, col: 1, offset: 5759},
	expr: &actionExpr{
	pos: position{line: 233, col: 18, offset: 5776},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 233, col: 18, offset: 5776},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 233, col: 18, offset: 5776},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 233, col: 23, offset: 5781},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 233, col: 28, offset: 5786},
	name: "_",
},
&labeledExpr{
	pos: position{line: 233, col: 30, offset: 5788},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 233, col: 33, offset: 5791},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 233, col: 45, offset: 5803},
	name: "_",
},
&labeledExpr{
	pos: position{line: 233, col: 47, offset: 5805},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 233, col: 53, offset: 5811},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 233, col: 59, offset: 5817},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 237, col: 1, offset: 5891},
	expr: &choiceExpr{
	pos: position{line: 237, col: 16, offset: 5906},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 237, col: 16, offset: 5906},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 237, col: 23, offset: 5913},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 237, col: 30, offset: 5920},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 237, col: 36, offset: 5926},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 237, col: 42, offset: 5932},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 237, col: 49, offset: 5939},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 237, col: 49, offset: 5939},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 246, col: 1, offset: 6049},
	expr: &actionExpr{
	pos: position{line: 246, col: 12, offset: 6060},
	run: (*parser).callonStdOutG1,
	expr: &seqExpr{
	pos: position{line: 246, col: 12, offset: 6060},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 246, col: 12, offset: 6060},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 246, col: 22, offset: 6070},
	name: "_",
},
&labeledExpr{
	pos: position{line: 246, col: 24, offset: 6072},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 246, col: 31, offset: 6079},
	expr: &ruleRefExpr{
	pos: position{line: 246, col: 32, offset: 6080},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 246, col: 48, offset: 6096},
	name: "_",
},
&litMatcher{
	pos: position{line: 246, col: 50, offset: 6098},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SinkInto",
	pos: position{line: 260, col: 1, offset: 6344},
	expr: &actionExpr{
	pos: position{line: 260, col: 12, offset: 6355},
	run: (*parser).callonSinkInto1,
	expr: &seqExpr{
	pos: position{line: 260, col: 12, offset: 6355},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 260, col: 12, offset: 6355},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 260, col: 18, offset: 6361},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 260, col: 27, offset: 6370},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 260, col: 33, offset: 6376},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 260, col: 33, offset: 6376},
	name: "_",
},
&litMatcher{
	pos: position{line: 260, col: 35, offset: 6378},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 260, col: 39, offset: 6382},
	name: "_",
},
&litMatcher{
	pos: position{line: 260, col: 41, offset: 6384},
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
	name: "FileName",
	pos: position{line: 271, col: 1, offset: 6545},
	expr: &actionExpr{
	pos: position{line: 271, col: 13, offset: 6557},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 271, col: 13, offset: 6557},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 271, col: 13, offset: 6557},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 271, col: 22, offset: 6566},
	expr: &seqExpr{
	pos: position{line: 271, col: 23, offset: 6567},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 271, col: 23, offset: 6567},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 271, col: 27, offset: 6571},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 271, col: 31, offset: 6575},
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
	pos: position{line: 277, col: 1, offset: 6676},
	expr: &actionExpr{
	pos: position{line: 277, col: 10, offset: 6685},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 277, col: 10, offset: 6685},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 277, col: 10, offset: 6685},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 277, col: 16, offset: 6691},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 277, col: 23, offset: 6698},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 277, col: 28, offset: 6703},
	expr: &seqExpr{
	pos: position{line: 277, col: 29, offset: 6704},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 277, col: 29, offset: 6704},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 277, col: 31, offset: 6706},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 277, col: 36, offset: 6711},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 277, col: 38, offset: 6713},
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
	pos: position{line: 281, col: 1, offset: 6778},
	expr: &actionExpr{
	pos: position{line: 281, col: 11, offset: 6788},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 281, col: 11, offset: 6788},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 281, col: 16, offset: 6793},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 281, col: 16, offset: 6793},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 281, col: 24, offset: 6801},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 281, col: 33, offset: 6810},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 281, col: 48, offset: 6825},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 281, col: 59, offset: 6836},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 285, col: 1, offset: 6878},
	expr: &choiceExpr{
	pos: position{line: 285, col: 9, offset: 6886},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 285, col: 9, offset: 6886},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 285, col: 15, offset: 6892},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 285, col: 21, offset: 6898},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 285, col: 27, offset: 6904},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 285, col: 33, offset: 6910},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 285, col: 33, offset: 6910},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 289, col: 1, offset: 6958},
	expr: &choiceExpr{
	pos: position{line: 289, col: 12, offset: 6969},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 289, col: 12, offset: 6969},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 289, col: 12, offset: 6969},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 289, col: 12, offset: 6969},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 289, col: 16, offset: 6973},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 289, col: 19, offset: 6976},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 289, col: 25, offset: 6982},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 293, col: 5, offset: 7058},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 293, col: 5, offset: 7058},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 293, col: 8, offset: 7061},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 300, col: 1, offset: 7138},
	expr: &actionExpr{
	pos: position{line: 300, col: 9, offset: 7146},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 300, col: 9, offset: 7146},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 300, col: 9, offset: 7146},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 300, col: 15, offset: 7152},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 300, col: 21, offset: 7158},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 300, col: 26, offset: 7163},
	expr: &seqExpr{
	pos: position{line: 300, col: 27, offset: 7164},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 300, col: 27, offset: 7164},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 300, col: 29, offset: 7166},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 300, col: 34, offset: 7171},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 300, col: 36, offset: 7173},
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
	pos: position{line: 304, col: 1, offset: 7237},
	expr: &actionExpr{
	pos: position{line: 304, col: 10, offset: 7246},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 304, col: 10, offset: 7246},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 304, col: 15, offset: 7251},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 304, col: 15, offset: 7251},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 304, col: 23, offset: 7259},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 304, col: 32, offset: 7268},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 304, col: 47, offset: 7283},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 308, col: 1, offset: 7326},
	expr: &actionExpr{
	pos: position{line: 308, col: 13, offset: 7338},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 308, col: 13, offset: 7338},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 308, col: 13, offset: 7338},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 308, col: 20, offset: 7345},
	expr: &choiceExpr{
	pos: position{line: 308, col: 21, offset: 7346},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 308, col: 21, offset: 7346},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 308, col: 31, offset: 7356},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 308, col: 39, offset: 7364},
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
	pos: position{line: 312, col: 1, offset: 7414},
	expr: &actionExpr{
	pos: position{line: 312, col: 17, offset: 7430},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 312, col: 17, offset: 7430},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 312, col: 17, offset: 7430},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 312, col: 21, offset: 7434},
	expr: &choiceExpr{
	pos: position{line: 312, col: 23, offset: 7436},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 312, col: 23, offset: 7436},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 312, col: 23, offset: 7436},
	expr: &ruleRefExpr{
	pos: position{line: 312, col: 24, offset: 7437},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 312, col: 36, offset: 7449,
},
	},
},
&seqExpr{
	pos: position{line: 312, col: 40, offset: 7453},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 312, col: 40, offset: 7453},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 312, col: 45, offset: 7458},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 312, col: 63, offset: 7476},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 319, col: 1, offset: 7657},
	expr: &charClassMatcher{
	pos: position{line: 319, col: 16, offset: 7672},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 321, col: 1, offset: 7690},
	expr: &ruleRefExpr{
	pos: position{line: 321, col: 19, offset: 7708},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 323, col: 1, offset: 7728},
	expr: &charClassMatcher{
	pos: position{line: 323, col: 21, offset: 7748},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 325, col: 1, offset: 7763},
	expr: &actionExpr{
	pos: position{line: 325, col: 9, offset: 7771},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 325, col: 9, offset: 7771},
	expr: &choiceExpr{
	pos: position{line: 325, col: 10, offset: 7772},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 325, col: 10, offset: 7772},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 325, col: 19, offset: 7781},
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
	pos: position{line: 329, col: 1, offset: 7831},
	expr: &choiceExpr{
	pos: position{line: 329, col: 11, offset: 7841},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 329, col: 11, offset: 7841},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 329, col: 19, offset: 7849},
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
	pos: position{line: 331, col: 1, offset: 7858},
	expr: &actionExpr{
	pos: position{line: 331, col: 11, offset: 7868},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 331, col: 11, offset: 7868},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 331, col: 11, offset: 7868},
	expr: &litMatcher{
	pos: position{line: 331, col: 11, offset: 7868},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 331, col: 16, offset: 7873},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 335, col: 1, offset: 7946},
	expr: &actionExpr{
	pos: position{line: 335, col: 19, offset: 7964},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 335, col: 19, offset: 7964},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 339, col: 1, offset: 8037},
	expr: &actionExpr{
	pos: position{line: 339, col: 10, offset: 8046},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 339, col: 10, offset: 8046},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 339, col: 10, offset: 8046},
	expr: &litMatcher{
	pos: position{line: 339, col: 10, offset: 8046},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 339, col: 15, offset: 8051},
	expr: &ruleRefExpr{
	pos: position{line: 339, col: 15, offset: 8051},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 339, col: 24, offset: 8060},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 339, col: 28, offset: 8064},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 343, col: 1, offset: 8131},
	expr: &choiceExpr{
	pos: position{line: 343, col: 12, offset: 8142},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 343, col: 12, offset: 8142},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 343, col: 18, offset: 8148},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 343, col: 18, offset: 8148},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 343, col: 38, offset: 8168},
	expr: &ruleRefExpr{
	pos: position{line: 343, col: 38, offset: 8168},
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
	pos: position{line: 345, col: 1, offset: 8185},
	expr: &charClassMatcher{
	pos: position{line: 345, col: 17, offset: 8201},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 347, col: 1, offset: 8210},
	expr: &charClassMatcher{
	pos: position{line: 347, col: 24, offset: 8233},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 349, col: 1, offset: 8242},
	expr: &zeroOrMoreExpr{
	pos: position{line: 349, col: 19, offset: 8260},
	expr: &charClassMatcher{
	pos: position{line: 349, col: 19, offset: 8260},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 351, col: 1, offset: 8274},
	expr: &charClassMatcher{
	pos: position{line: 351, col: 18, offset: 8291},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "IDENTIFIER",
	pos: position{line: 352, col: 1, offset: 8302},
	expr: &actionExpr{
	pos: position{line: 352, col: 15, offset: 8316},
	run: (*parser).callonIDENTIFIER1,
	expr: &seqExpr{
	pos: position{line: 352, col: 15, offset: 8316},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 352, col: 15, offset: 8316},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 352, col: 22, offset: 8323},
	expr: &choiceExpr{
	pos: position{line: 352, col: 23, offset: 8324},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 352, col: 23, offset: 8324},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 352, col: 33, offset: 8334},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 352, col: 41, offset: 8342},
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
	pos: position{line: 355, col: 1, offset: 8415},
	expr: &notExpr{
	pos: position{line: 355, col: 8, offset: 8422},
	expr: &anyMatcher{
	line: 355, col: 9, offset: 8423,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 357, col: 1, offset: 8428},
	expr: &litMatcher{
	pos: position{line: 357, col: 17, offset: 8444},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 363, col: 1, offset: 8531},
	expr: &actionExpr{
	pos: position{line: 363, col: 18, offset: 8548},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 363, col: 18, offset: 8548},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 363, col: 18, offset: 8548},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 363, col: 24, offset: 8554},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 363, col: 33, offset: 8563},
	name: "_",
},
&labeledExpr{
	pos: position{line: 363, col: 35, offset: 8565},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 363, col: 40, offset: 8570},
	expr: &seqExpr{
	pos: position{line: 363, col: 41, offset: 8571},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 363, col: 41, offset: 8571},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 363, col: 45, offset: 8575},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 363, col: 47, offset: 8577},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 363, col: 56, offset: 8586},
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
	pos: position{line: 366, col: 1, offset: 8639},
	expr: &actionExpr{
	pos: position{line: 366, col: 13, offset: 8651},
	run: (*parser).callonEqAliasG1,
	expr: &seqExpr{
	pos: position{line: 366, col: 13, offset: 8651},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 366, col: 13, offset: 8651},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 366, col: 19, offset: 8657},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 366, col: 28, offset: 8666},
	name: "_",
},
&litMatcher{
	pos: position{line: 366, col: 30, offset: 8668},
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

    			if len(newVal) > 0 {
    				firstVal := newVal[0]
    				_, ok1 := firstVal.([]interface{})
    				if ok1 {
    					finalVal := cast.ToIfaceSlice(firstVal)
    					if len(finalVal) == 4 {
    						stages = append(stages, finalVal[3])
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
        return FileJob{Filename: fname}, nil
    }
    return nil, errors.New("could not decode source file name")
}

func (p *parser) callonFileJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFileJobG1(stack["filename"])
}

func (c *current) onBranchJobG1(br interface{}) (interface{}, error) {

    if branchName,ok := cast.TryString(br);ok{
    return BranchJob{BranchName:branchName},nil
    }
    return nil,errors.New("could not identify the branch name")
}

func (p *parser) callonBranchJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBranchJobG1(stack["br"])
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

func (c *current) onSinkInto1(first, rest interface{}) (interface{}, error) {

    streamTo,_ := first.(string)
    return SinkJob{
        Type: "into",
        Args : streamTo,
        },nil
    
}

func (p *parser) callonSinkInto1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSinkInto1(stack["first"], stack["rest"])
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

