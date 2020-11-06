
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
	pos: position{line: 86, col: 1, offset: 2030},
	expr: &actionExpr{
	pos: position{line: 86, col: 15, offset: 2044},
	run: (*parser).callonBranchJobG1,
	expr: &seqExpr{
	pos: position{line: 86, col: 15, offset: 2044},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 86, col: 15, offset: 2044},
	val: "branch",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 86, col: 23, offset: 2052},
	name: "_",
},
&litMatcher{
	pos: position{line: 86, col: 24, offset: 2053},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 86, col: 28, offset: 2057},
	name: "_",
},
&labeledExpr{
	pos: position{line: 86, col: 30, offset: 2059},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 86, col: 33, offset: 2062},
	name: "IDENTIFIER",
},
},
&litMatcher{
	pos: position{line: 86, col: 44, offset: 2073},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggJobG",
	pos: position{line: 99, col: 1, offset: 2375},
	expr: &actionExpr{
	pos: position{line: 99, col: 12, offset: 2386},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 99, col: 12, offset: 2386},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 99, col: 12, offset: 2386},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 99, col: 18, offset: 2392},
	name: "_",
},
&labeledExpr{
	pos: position{line: 99, col: 20, offset: 2394},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 99, col: 27, offset: 2401},
	expr: &ruleRefExpr{
	pos: position{line: 99, col: 27, offset: 2401},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 99, col: 39, offset: 2413},
	name: "_",
},
&labeledExpr{
	pos: position{line: 99, col: 41, offset: 2415},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 99, col: 45, offset: 2419},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 99, col: 53, offset: 2427},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 99, col: 58, offset: 2432},
	expr: &seqExpr{
	pos: position{line: 99, col: 59, offset: 2433},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 99, col: 59, offset: 2433},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 99, col: 63, offset: 2437},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 99, col: 65, offset: 2439},
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
	pos: position{line: 114, col: 1, offset: 2814},
	expr: &choiceExpr{
	pos: position{line: 114, col: 12, offset: 2825},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 114, col: 12, offset: 2825},
	name: "CountAgg",
},
&ruleRefExpr{
	pos: position{line: 114, col: 23, offset: 2836},
	name: "MaxAgg",
},
&ruleRefExpr{
	pos: position{line: 114, col: 32, offset: 2845},
	name: "MinAgg",
},
&ruleRefExpr{
	pos: position{line: 114, col: 41, offset: 2854},
	name: "AvgAgg",
},
&ruleRefExpr{
	pos: position{line: 114, col: 50, offset: 2863},
	name: "SumAgg",
},
&ruleRefExpr{
	pos: position{line: 114, col: 59, offset: 2872},
	name: "ModeAgg",
},
&ruleRefExpr{
	pos: position{line: 115, col: 13, offset: 2893},
	name: "VarAgg",
},
&ruleRefExpr{
	pos: position{line: 115, col: 22, offset: 2902},
	name: "DistinctCountAgg",
},
&ruleRefExpr{
	pos: position{line: 115, col: 41, offset: 2921},
	name: "QuantileAgg",
},
&ruleRefExpr{
	pos: position{line: 115, col: 55, offset: 2935},
	name: "MedianAgg",
},
&ruleRefExpr{
	pos: position{line: 115, col: 67, offset: 2947},
	name: "QuartileAgg",
},
&ruleRefExpr{
	pos: position{line: 116, col: 13, offset: 2972},
	name: "CovAgg",
},
&ruleRefExpr{
	pos: position{line: 116, col: 22, offset: 2981},
	name: "CorrAgg",
},
&ruleRefExpr{
	pos: position{line: 116, col: 32, offset: 2991},
	name: "PValueAgg",
},
	},
},
},
{
	name: "CountAgg",
	pos: position{line: 119, col: 1, offset: 3035},
	expr: &actionExpr{
	pos: position{line: 119, col: 13, offset: 3047},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 119, col: 13, offset: 3047},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 119, col: 13, offset: 3047},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 119, col: 19, offset: 3053},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 119, col: 28, offset: 3062},
	name: "_",
},
&litMatcher{
	pos: position{line: 119, col: 30, offset: 3064},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 119, col: 39, offset: 3073},
	name: "_",
},
&labeledExpr{
	pos: position{line: 119, col: 41, offset: 3075},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 119, col: 46, offset: 3080},
	expr: &ruleRefExpr{
	pos: position{line: 119, col: 46, offset: 3080},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 119, col: 59, offset: 3093},
	name: "_",
},
&litMatcher{
	pos: position{line: 119, col: 61, offset: 3095},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MaxAgg",
	pos: position{line: 126, col: 1, offset: 3247},
	expr: &actionExpr{
	pos: position{line: 126, col: 11, offset: 3257},
	run: (*parser).callonMaxAgg1,
	expr: &seqExpr{
	pos: position{line: 126, col: 11, offset: 3257},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 126, col: 11, offset: 3257},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 126, col: 17, offset: 3263},
	expr: &ruleRefExpr{
	pos: position{line: 126, col: 17, offset: 3263},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 126, col: 27, offset: 3273},
	name: "_",
},
&litMatcher{
	pos: position{line: 126, col: 29, offset: 3275},
	val: "max(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 126, col: 36, offset: 3282},
	name: "_",
},
&labeledExpr{
	pos: position{line: 126, col: 38, offset: 3284},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 126, col: 44, offset: 3290},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 126, col: 53, offset: 3299},
	name: "_",
},
&labeledExpr{
	pos: position{line: 126, col: 55, offset: 3301},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 126, col: 60, offset: 3306},
	expr: &seqExpr{
	pos: position{line: 126, col: 61, offset: 3307},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 126, col: 61, offset: 3307},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 126, col: 65, offset: 3311},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 126, col: 67, offset: 3313},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 126, col: 81, offset: 3327},
	name: "_",
},
&litMatcher{
	pos: position{line: 126, col: 83, offset: 3329},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MinAgg",
	pos: position{line: 148, col: 1, offset: 3790},
	expr: &actionExpr{
	pos: position{line: 148, col: 11, offset: 3800},
	run: (*parser).callonMinAgg1,
	expr: &seqExpr{
	pos: position{line: 148, col: 11, offset: 3800},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 148, col: 11, offset: 3800},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 148, col: 17, offset: 3806},
	expr: &ruleRefExpr{
	pos: position{line: 148, col: 17, offset: 3806},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 148, col: 27, offset: 3816},
	name: "_",
},
&litMatcher{
	pos: position{line: 148, col: 29, offset: 3818},
	val: "min(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 148, col: 36, offset: 3825},
	name: "_",
},
&labeledExpr{
	pos: position{line: 148, col: 38, offset: 3827},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 148, col: 44, offset: 3833},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 148, col: 53, offset: 3842},
	name: "_",
},
&labeledExpr{
	pos: position{line: 148, col: 55, offset: 3844},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 148, col: 60, offset: 3849},
	expr: &seqExpr{
	pos: position{line: 148, col: 61, offset: 3850},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 148, col: 61, offset: 3850},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 148, col: 65, offset: 3854},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 148, col: 67, offset: 3856},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 148, col: 81, offset: 3870},
	name: "_",
},
&litMatcher{
	pos: position{line: 148, col: 83, offset: 3872},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AvgAgg",
	pos: position{line: 170, col: 1, offset: 4333},
	expr: &actionExpr{
	pos: position{line: 170, col: 11, offset: 4343},
	run: (*parser).callonAvgAgg1,
	expr: &seqExpr{
	pos: position{line: 170, col: 11, offset: 4343},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 170, col: 11, offset: 4343},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 170, col: 17, offset: 4349},
	expr: &ruleRefExpr{
	pos: position{line: 170, col: 17, offset: 4349},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 170, col: 27, offset: 4359},
	name: "_",
},
&litMatcher{
	pos: position{line: 170, col: 29, offset: 4361},
	val: "avg(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 170, col: 36, offset: 4368},
	name: "_",
},
&labeledExpr{
	pos: position{line: 170, col: 38, offset: 4370},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 170, col: 44, offset: 4376},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 170, col: 53, offset: 4385},
	name: "_",
},
&labeledExpr{
	pos: position{line: 170, col: 55, offset: 4387},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 170, col: 60, offset: 4392},
	expr: &seqExpr{
	pos: position{line: 170, col: 61, offset: 4393},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 170, col: 61, offset: 4393},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 170, col: 65, offset: 4397},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 170, col: 67, offset: 4399},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 170, col: 81, offset: 4413},
	name: "_",
},
&litMatcher{
	pos: position{line: 170, col: 83, offset: 4415},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SumAgg",
	pos: position{line: 192, col: 1, offset: 4876},
	expr: &actionExpr{
	pos: position{line: 192, col: 11, offset: 4886},
	run: (*parser).callonSumAgg1,
	expr: &seqExpr{
	pos: position{line: 192, col: 11, offset: 4886},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 192, col: 11, offset: 4886},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 192, col: 17, offset: 4892},
	expr: &ruleRefExpr{
	pos: position{line: 192, col: 17, offset: 4892},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 192, col: 27, offset: 4902},
	name: "_",
},
&litMatcher{
	pos: position{line: 192, col: 29, offset: 4904},
	val: "sum(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 192, col: 36, offset: 4911},
	name: "_",
},
&labeledExpr{
	pos: position{line: 192, col: 38, offset: 4913},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 192, col: 44, offset: 4919},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 192, col: 53, offset: 4928},
	name: "_",
},
&labeledExpr{
	pos: position{line: 192, col: 55, offset: 4930},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 192, col: 60, offset: 4935},
	expr: &seqExpr{
	pos: position{line: 192, col: 61, offset: 4936},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 192, col: 61, offset: 4936},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 192, col: 65, offset: 4940},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 192, col: 67, offset: 4942},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 192, col: 81, offset: 4956},
	name: "_",
},
&litMatcher{
	pos: position{line: 192, col: 83, offset: 4958},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarAgg",
	pos: position{line: 214, col: 1, offset: 5419},
	expr: &actionExpr{
	pos: position{line: 214, col: 11, offset: 5429},
	run: (*parser).callonVarAgg1,
	expr: &seqExpr{
	pos: position{line: 214, col: 11, offset: 5429},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 214, col: 11, offset: 5429},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 214, col: 17, offset: 5435},
	expr: &ruleRefExpr{
	pos: position{line: 214, col: 17, offset: 5435},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 214, col: 27, offset: 5445},
	name: "_",
},
&litMatcher{
	pos: position{line: 214, col: 29, offset: 5447},
	val: "var(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 214, col: 36, offset: 5454},
	name: "_",
},
&labeledExpr{
	pos: position{line: 214, col: 38, offset: 5456},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 214, col: 44, offset: 5462},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 214, col: 53, offset: 5471},
	name: "_",
},
&labeledExpr{
	pos: position{line: 214, col: 55, offset: 5473},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 214, col: 60, offset: 5478},
	expr: &seqExpr{
	pos: position{line: 214, col: 61, offset: 5479},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 214, col: 61, offset: 5479},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 214, col: 65, offset: 5483},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 214, col: 67, offset: 5485},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 214, col: 81, offset: 5499},
	name: "_",
},
&litMatcher{
	pos: position{line: 214, col: 83, offset: 5501},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DistinctCountAgg",
	pos: position{line: 236, col: 1, offset: 5967},
	expr: &actionExpr{
	pos: position{line: 236, col: 21, offset: 5987},
	run: (*parser).callonDistinctCountAgg1,
	expr: &seqExpr{
	pos: position{line: 236, col: 21, offset: 5987},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 236, col: 21, offset: 5987},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 236, col: 27, offset: 5993},
	expr: &ruleRefExpr{
	pos: position{line: 236, col: 27, offset: 5993},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 236, col: 37, offset: 6003},
	name: "_",
},
&litMatcher{
	pos: position{line: 236, col: 39, offset: 6005},
	val: "dcount(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 236, col: 49, offset: 6015},
	name: "_",
},
&labeledExpr{
	pos: position{line: 236, col: 51, offset: 6017},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 236, col: 57, offset: 6023},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 236, col: 66, offset: 6032},
	name: "_",
},
&labeledExpr{
	pos: position{line: 236, col: 68, offset: 6034},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 236, col: 73, offset: 6039},
	expr: &seqExpr{
	pos: position{line: 236, col: 74, offset: 6040},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 236, col: 74, offset: 6040},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 236, col: 78, offset: 6044},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 236, col: 80, offset: 6046},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 236, col: 94, offset: 6060},
	name: "_",
},
&litMatcher{
	pos: position{line: 236, col: 96, offset: 6062},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "ModeAgg",
	pos: position{line: 258, col: 1, offset: 6533},
	expr: &actionExpr{
	pos: position{line: 258, col: 12, offset: 6544},
	run: (*parser).callonModeAgg1,
	expr: &seqExpr{
	pos: position{line: 258, col: 12, offset: 6544},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 258, col: 12, offset: 6544},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 258, col: 18, offset: 6550},
	expr: &ruleRefExpr{
	pos: position{line: 258, col: 18, offset: 6550},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 258, col: 28, offset: 6560},
	name: "_",
},
&litMatcher{
	pos: position{line: 258, col: 30, offset: 6562},
	val: "mode(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 258, col: 38, offset: 6570},
	name: "_",
},
&labeledExpr{
	pos: position{line: 258, col: 40, offset: 6572},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 258, col: 46, offset: 6578},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 258, col: 55, offset: 6587},
	name: "_",
},
&labeledExpr{
	pos: position{line: 258, col: 57, offset: 6589},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 258, col: 62, offset: 6594},
	expr: &seqExpr{
	pos: position{line: 258, col: 63, offset: 6595},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 258, col: 63, offset: 6595},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 258, col: 67, offset: 6599},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 258, col: 69, offset: 6601},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 258, col: 83, offset: 6615},
	name: "_",
},
&litMatcher{
	pos: position{line: 258, col: 85, offset: 6617},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuantileAgg",
	pos: position{line: 280, col: 1, offset: 7079},
	expr: &actionExpr{
	pos: position{line: 280, col: 16, offset: 7094},
	run: (*parser).callonQuantileAgg1,
	expr: &seqExpr{
	pos: position{line: 280, col: 16, offset: 7094},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 280, col: 16, offset: 7094},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 280, col: 22, offset: 7100},
	expr: &ruleRefExpr{
	pos: position{line: 280, col: 22, offset: 7100},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 280, col: 32, offset: 7110},
	name: "_",
},
&litMatcher{
	pos: position{line: 280, col: 34, offset: 7112},
	val: "quantile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 280, col: 46, offset: 7124},
	name: "_",
},
&labeledExpr{
	pos: position{line: 280, col: 48, offset: 7126},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 280, col: 54, offset: 7132},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 280, col: 63, offset: 7141},
	name: "_",
},
&labeledExpr{
	pos: position{line: 280, col: 66, offset: 7144},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 280, col: 73, offset: 7151},
	expr: &seqExpr{
	pos: position{line: 280, col: 74, offset: 7152},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 280, col: 74, offset: 7152},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 280, col: 78, offset: 7156},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 280, col: 80, offset: 7158},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 280, col: 91, offset: 7169},
	name: "_",
},
&litMatcher{
	pos: position{line: 280, col: 93, offset: 7171},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 280, col: 97, offset: 7175},
	name: "_",
},
&labeledExpr{
	pos: position{line: 280, col: 99, offset: 7177},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 280, col: 103, offset: 7181},
	name: "Float",
},
},
&ruleRefExpr{
	pos: position{line: 280, col: 109, offset: 7187},
	name: "_",
},
&labeledExpr{
	pos: position{line: 280, col: 111, offset: 7189},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 280, col: 116, offset: 7194},
	expr: &seqExpr{
	pos: position{line: 280, col: 117, offset: 7195},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 280, col: 117, offset: 7195},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 280, col: 121, offset: 7199},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 280, col: 123, offset: 7201},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 280, col: 137, offset: 7215},
	name: "_",
},
&litMatcher{
	pos: position{line: 280, col: 139, offset: 7217},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MedianAgg",
	pos: position{line: 310, col: 1, offset: 7862},
	expr: &actionExpr{
	pos: position{line: 310, col: 14, offset: 7875},
	run: (*parser).callonMedianAgg1,
	expr: &seqExpr{
	pos: position{line: 310, col: 14, offset: 7875},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 310, col: 14, offset: 7875},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 310, col: 20, offset: 7881},
	expr: &ruleRefExpr{
	pos: position{line: 310, col: 20, offset: 7881},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 310, col: 30, offset: 7891},
	name: "_",
},
&litMatcher{
	pos: position{line: 310, col: 32, offset: 7893},
	val: "median(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 310, col: 42, offset: 7903},
	name: "_",
},
&labeledExpr{
	pos: position{line: 310, col: 44, offset: 7905},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 310, col: 50, offset: 7911},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 310, col: 59, offset: 7920},
	name: "_",
},
&labeledExpr{
	pos: position{line: 310, col: 62, offset: 7923},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 310, col: 69, offset: 7930},
	expr: &seqExpr{
	pos: position{line: 310, col: 70, offset: 7931},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 310, col: 70, offset: 7931},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 310, col: 74, offset: 7935},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 310, col: 76, offset: 7937},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 310, col: 87, offset: 7948},
	name: "_",
},
&labeledExpr{
	pos: position{line: 310, col: 89, offset: 7950},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 310, col: 94, offset: 7955},
	expr: &seqExpr{
	pos: position{line: 310, col: 95, offset: 7956},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 310, col: 95, offset: 7956},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 310, col: 99, offset: 7960},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 310, col: 101, offset: 7962},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 310, col: 115, offset: 7976},
	name: "_",
},
&litMatcher{
	pos: position{line: 310, col: 117, offset: 7978},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuartileAgg",
	pos: position{line: 342, col: 1, offset: 8655},
	expr: &actionExpr{
	pos: position{line: 342, col: 16, offset: 8670},
	run: (*parser).callonQuartileAgg1,
	expr: &seqExpr{
	pos: position{line: 342, col: 16, offset: 8670},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 342, col: 16, offset: 8670},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 342, col: 22, offset: 8676},
	expr: &ruleRefExpr{
	pos: position{line: 342, col: 22, offset: 8676},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 342, col: 32, offset: 8686},
	name: "_",
},
&litMatcher{
	pos: position{line: 342, col: 34, offset: 8688},
	val: "quartile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 342, col: 46, offset: 8700},
	name: "_",
},
&labeledExpr{
	pos: position{line: 342, col: 48, offset: 8702},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 342, col: 54, offset: 8708},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 342, col: 63, offset: 8717},
	name: "_",
},
&labeledExpr{
	pos: position{line: 342, col: 66, offset: 8720},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 342, col: 73, offset: 8727},
	expr: &seqExpr{
	pos: position{line: 342, col: 74, offset: 8728},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 342, col: 74, offset: 8728},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 342, col: 78, offset: 8732},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 342, col: 80, offset: 8734},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 342, col: 91, offset: 8745},
	name: "_",
},
&litMatcher{
	pos: position{line: 342, col: 93, offset: 8747},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 342, col: 97, offset: 8751},
	name: "_",
},
&labeledExpr{
	pos: position{line: 342, col: 99, offset: 8753},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 342, col: 103, offset: 8757},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 342, col: 110, offset: 8764},
	name: "_",
},
&labeledExpr{
	pos: position{line: 342, col: 112, offset: 8766},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 342, col: 117, offset: 8771},
	expr: &seqExpr{
	pos: position{line: 342, col: 118, offset: 8772},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 342, col: 118, offset: 8772},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 342, col: 122, offset: 8776},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 342, col: 124, offset: 8778},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 342, col: 138, offset: 8792},
	name: "_",
},
&litMatcher{
	pos: position{line: 342, col: 140, offset: 8794},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CovAgg",
	pos: position{line: 388, col: 1, offset: 9821},
	expr: &actionExpr{
	pos: position{line: 388, col: 11, offset: 9831},
	run: (*parser).callonCovAgg1,
	expr: &seqExpr{
	pos: position{line: 388, col: 11, offset: 9831},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 388, col: 11, offset: 9831},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 388, col: 17, offset: 9837},
	expr: &ruleRefExpr{
	pos: position{line: 388, col: 17, offset: 9837},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 388, col: 27, offset: 9847},
	name: "_",
},
&litMatcher{
	pos: position{line: 388, col: 29, offset: 9849},
	val: "cov(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 388, col: 36, offset: 9856},
	name: "_",
},
&labeledExpr{
	pos: position{line: 388, col: 38, offset: 9858},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 388, col: 45, offset: 9865},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 388, col: 54, offset: 9874},
	name: "_",
},
&litMatcher{
	pos: position{line: 388, col: 56, offset: 9876},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 388, col: 60, offset: 9880},
	name: "_",
},
&labeledExpr{
	pos: position{line: 388, col: 62, offset: 9882},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 388, col: 69, offset: 9889},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 388, col: 78, offset: 9898},
	name: "_",
},
&labeledExpr{
	pos: position{line: 388, col: 80, offset: 9900},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 388, col: 85, offset: 9905},
	expr: &seqExpr{
	pos: position{line: 388, col: 86, offset: 9906},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 388, col: 86, offset: 9906},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 388, col: 90, offset: 9910},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 388, col: 92, offset: 9912},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 388, col: 106, offset: 9926},
	name: "_",
},
&litMatcher{
	pos: position{line: 388, col: 108, offset: 9928},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CorrAgg",
	pos: position{line: 411, col: 1, offset: 10453},
	expr: &actionExpr{
	pos: position{line: 411, col: 12, offset: 10464},
	run: (*parser).callonCorrAgg1,
	expr: &seqExpr{
	pos: position{line: 411, col: 12, offset: 10464},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 411, col: 12, offset: 10464},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 411, col: 18, offset: 10470},
	expr: &ruleRefExpr{
	pos: position{line: 411, col: 18, offset: 10470},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 411, col: 28, offset: 10480},
	name: "_",
},
&litMatcher{
	pos: position{line: 411, col: 30, offset: 10482},
	val: "correlation(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 411, col: 45, offset: 10497},
	name: "_",
},
&labeledExpr{
	pos: position{line: 411, col: 47, offset: 10499},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 411, col: 54, offset: 10506},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 411, col: 63, offset: 10515},
	name: "_",
},
&litMatcher{
	pos: position{line: 411, col: 65, offset: 10517},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 411, col: 69, offset: 10521},
	name: "_",
},
&labeledExpr{
	pos: position{line: 411, col: 71, offset: 10523},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 411, col: 78, offset: 10530},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 411, col: 87, offset: 10539},
	name: "_",
},
&labeledExpr{
	pos: position{line: 411, col: 89, offset: 10541},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 411, col: 94, offset: 10546},
	expr: &seqExpr{
	pos: position{line: 411, col: 95, offset: 10547},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 411, col: 95, offset: 10547},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 411, col: 99, offset: 10551},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 411, col: 101, offset: 10553},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 411, col: 115, offset: 10567},
	name: "_",
},
&litMatcher{
	pos: position{line: 411, col: 117, offset: 10569},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PValueAgg",
	pos: position{line: 434, col: 1, offset: 11095},
	expr: &actionExpr{
	pos: position{line: 434, col: 14, offset: 11108},
	run: (*parser).callonPValueAgg1,
	expr: &seqExpr{
	pos: position{line: 434, col: 14, offset: 11108},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 434, col: 14, offset: 11108},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 434, col: 20, offset: 11114},
	expr: &ruleRefExpr{
	pos: position{line: 434, col: 20, offset: 11114},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 434, col: 30, offset: 11124},
	name: "_",
},
&litMatcher{
	pos: position{line: 434, col: 32, offset: 11126},
	val: "pvalue(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 434, col: 42, offset: 11136},
	name: "_",
},
&labeledExpr{
	pos: position{line: 434, col: 44, offset: 11138},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 434, col: 51, offset: 11145},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 434, col: 60, offset: 11154},
	name: "_",
},
&litMatcher{
	pos: position{line: 434, col: 62, offset: 11156},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 434, col: 66, offset: 11160},
	name: "_",
},
&labeledExpr{
	pos: position{line: 434, col: 68, offset: 11162},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 434, col: 75, offset: 11169},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 434, col: 84, offset: 11178},
	name: "_",
},
&labeledExpr{
	pos: position{line: 434, col: 86, offset: 11180},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 434, col: 91, offset: 11185},
	expr: &seqExpr{
	pos: position{line: 434, col: 92, offset: 11186},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 434, col: 92, offset: 11186},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 434, col: 96, offset: 11190},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 434, col: 98, offset: 11192},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 434, col: 112, offset: 11206},
	name: "_",
},
&litMatcher{
	pos: position{line: 434, col: 114, offset: 11208},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggGroupBy",
	pos: position{line: 459, col: 1, offset: 11733},
	expr: &actionExpr{
	pos: position{line: 459, col: 15, offset: 11747},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 459, col: 15, offset: 11747},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 459, col: 15, offset: 11747},
	name: "_",
},
&litMatcher{
	pos: position{line: 459, col: 17, offset: 11749},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 459, col: 22, offset: 11754},
	name: "_",
},
&labeledExpr{
	pos: position{line: 459, col: 24, offset: 11756},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 459, col: 30, offset: 11762},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 459, col: 39, offset: 11771},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 459, col: 44, offset: 11776},
	expr: &seqExpr{
	pos: position{line: 459, col: 45, offset: 11777},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 459, col: 45, offset: 11777},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 459, col: 49, offset: 11781},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 459, col: 51, offset: 11783},
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
	pos: position{line: 474, col: 1, offset: 12134},
	expr: &actionExpr{
	pos: position{line: 474, col: 11, offset: 12144},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 474, col: 11, offset: 12144},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 474, col: 16, offset: 12149},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 474, col: 16, offset: 12149},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 474, col: 27, offset: 12160},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 474, col: 38, offset: 12171},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 474, col: 49, offset: 12182},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 474, col: 58, offset: 12191},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 474, col: 67, offset: 12200},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 480, col: 1, offset: 12284},
	expr: &actionExpr{
	pos: position{line: 480, col: 13, offset: 12296},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 480, col: 13, offset: 12296},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 480, col: 13, offset: 12296},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 480, col: 23, offset: 12306},
	name: "_",
},
&labeledExpr{
	pos: position{line: 480, col: 25, offset: 12308},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 480, col: 28, offset: 12311},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 480, col: 42, offset: 12325},
	name: "_",
},
&litMatcher{
	pos: position{line: 480, col: 44, offset: 12327},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 487, col: 1, offset: 12522},
	expr: &actionExpr{
	pos: position{line: 487, col: 18, offset: 12539},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 487, col: 18, offset: 12539},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 487, col: 18, offset: 12539},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 487, col: 23, offset: 12544},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 487, col: 28, offset: 12549},
	name: "_",
},
&labeledExpr{
	pos: position{line: 487, col: 30, offset: 12551},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 487, col: 33, offset: 12554},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 487, col: 45, offset: 12566},
	name: "_",
},
&labeledExpr{
	pos: position{line: 487, col: 47, offset: 12568},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 487, col: 54, offset: 12575},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 491, col: 1, offset: 12626},
	expr: &ruleRefExpr{
	pos: position{line: 491, col: 16, offset: 12641},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 494, col: 1, offset: 12675},
	expr: &actionExpr{
	pos: position{line: 494, col: 13, offset: 12687},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 494, col: 13, offset: 12687},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 494, col: 13, offset: 12687},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 494, col: 23, offset: 12697},
	name: "_",
},
&labeledExpr{
	pos: position{line: 494, col: 25, offset: 12699},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 494, col: 31, offset: 12705},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 494, col: 40, offset: 12714},
	name: "_",
},
&labeledExpr{
	pos: position{line: 494, col: 42, offset: 12716},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 494, col: 47, offset: 12721},
	expr: &seqExpr{
	pos: position{line: 494, col: 48, offset: 12722},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 494, col: 48, offset: 12722},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 494, col: 52, offset: 12726},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 494, col: 54, offset: 12728},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 494, col: 65, offset: 12739},
	name: "_",
},
&litMatcher{
	pos: position{line: 494, col: 67, offset: 12741},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 509, col: 1, offset: 13078},
	expr: &actionExpr{
	pos: position{line: 509, col: 11, offset: 13088},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 509, col: 11, offset: 13088},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 509, col: 11, offset: 13088},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 509, col: 18, offset: 13095},
	name: "_",
},
&labeledExpr{
	pos: position{line: 509, col: 20, offset: 13097},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 509, col: 25, offset: 13102},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 509, col: 34, offset: 13111},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 509, col: 38, offset: 13115},
	name: "_",
},
&labeledExpr{
	pos: position{line: 509, col: 40, offset: 13117},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 509, col: 44, offset: 13121},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 509, col: 51, offset: 13128},
	name: "_",
},
&litMatcher{
	pos: position{line: 509, col: 53, offset: 13130},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 521, col: 1, offset: 13352},
	expr: &actionExpr{
	pos: position{line: 521, col: 11, offset: 13362},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 521, col: 11, offset: 13362},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 521, col: 11, offset: 13362},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 521, col: 20, offset: 13371},
	name: "_",
},
&litMatcher{
	pos: position{line: 521, col: 22, offset: 13373},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 521, col: 27, offset: 13378},
	name: "_",
},
&labeledExpr{
	pos: position{line: 521, col: 29, offset: 13380},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 521, col: 35, offset: 13386},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 527, col: 1, offset: 13492},
	expr: &actionExpr{
	pos: position{line: 527, col: 12, offset: 13503},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 527, col: 12, offset: 13503},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 527, col: 12, offset: 13503},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 527, col: 21, offset: 13512},
	name: "_",
},
&labeledExpr{
	pos: position{line: 527, col: 23, offset: 13514},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 527, col: 27, offset: 13518},
	expr: &ruleRefExpr{
	pos: position{line: 527, col: 27, offset: 13518},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 527, col: 35, offset: 13526},
	name: "_",
},
&litMatcher{
	pos: position{line: 527, col: 37, offset: 13528},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 537, col: 1, offset: 13673},
	expr: &actionExpr{
	pos: position{line: 537, col: 13, offset: 13685},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 537, col: 13, offset: 13685},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 537, col: 13, offset: 13685},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 537, col: 23, offset: 13695},
	name: "_",
},
&labeledExpr{
	pos: position{line: 537, col: 25, offset: 13697},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 537, col: 30, offset: 13702},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 537, col: 42, offset: 13714},
	name: "_",
},
&litMatcher{
	pos: position{line: 537, col: 44, offset: 13716},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 547, col: 1, offset: 14042},
	expr: &actionExpr{
	pos: position{line: 547, col: 16, offset: 14057},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 547, col: 16, offset: 14057},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 547, col: 16, offset: 14057},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 547, col: 22, offset: 14063},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 547, col: 35, offset: 14076},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 547, col: 40, offset: 14081},
	expr: &seqExpr{
	pos: position{line: 547, col: 41, offset: 14082},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 547, col: 41, offset: 14082},
	name: "_",
},
&labeledExpr{
	pos: position{line: 547, col: 43, offset: 14084},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 547, col: 46, offset: 14087},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 547, col: 58, offset: 14099},
	name: "_",
},
&labeledExpr{
	pos: position{line: 547, col: 60, offset: 14101},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 547, col: 63, offset: 14104},
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
	pos: position{line: 552, col: 1, offset: 14262},
	expr: &choiceExpr{
	pos: position{line: 552, col: 17, offset: 14278},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 552, col: 17, offset: 14278},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 552, col: 17, offset: 14278},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 552, col: 17, offset: 14278},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 552, col: 21, offset: 14282},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 552, col: 24, offset: 14285},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 552, col: 36, offset: 14297},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 554, col: 5, offset: 14328},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 554, col: 5, offset: 14328},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 554, col: 8, offset: 14331},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 558, col: 1, offset: 14373},
	expr: &choiceExpr{
	pos: position{line: 558, col: 16, offset: 14388},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 558, col: 16, offset: 14388},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 558, col: 24, offset: 14396},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 558, col: 24, offset: 14396},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 565, col: 1, offset: 14575},
	expr: &actionExpr{
	pos: position{line: 565, col: 18, offset: 14592},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 565, col: 18, offset: 14592},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 565, col: 18, offset: 14592},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 565, col: 23, offset: 14597},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 565, col: 28, offset: 14602},
	name: "_",
},
&labeledExpr{
	pos: position{line: 565, col: 30, offset: 14604},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 565, col: 33, offset: 14607},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 565, col: 45, offset: 14619},
	name: "_",
},
&labeledExpr{
	pos: position{line: 565, col: 47, offset: 14621},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 565, col: 53, offset: 14627},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 565, col: 59, offset: 14633},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 569, col: 1, offset: 14707},
	expr: &choiceExpr{
	pos: position{line: 569, col: 16, offset: 14722},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 569, col: 16, offset: 14722},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 569, col: 23, offset: 14729},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 569, col: 30, offset: 14736},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 569, col: 36, offset: 14742},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 569, col: 42, offset: 14748},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 569, col: 49, offset: 14755},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 569, col: 49, offset: 14755},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 578, col: 1, offset: 14865},
	expr: &actionExpr{
	pos: position{line: 578, col: 12, offset: 14876},
	run: (*parser).callonStdOutG1,
	expr: &seqExpr{
	pos: position{line: 578, col: 12, offset: 14876},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 578, col: 12, offset: 14876},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 578, col: 22, offset: 14886},
	name: "_",
},
&labeledExpr{
	pos: position{line: 578, col: 24, offset: 14888},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 578, col: 31, offset: 14895},
	expr: &ruleRefExpr{
	pos: position{line: 578, col: 32, offset: 14896},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 578, col: 48, offset: 14912},
	name: "_",
},
&litMatcher{
	pos: position{line: 578, col: 50, offset: 14914},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SinkInto",
	pos: position{line: 592, col: 1, offset: 15193},
	expr: &actionExpr{
	pos: position{line: 592, col: 12, offset: 15204},
	run: (*parser).callonSinkInto1,
	expr: &seqExpr{
	pos: position{line: 592, col: 12, offset: 15204},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 592, col: 12, offset: 15204},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 592, col: 18, offset: 15210},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 592, col: 27, offset: 15219},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 592, col: 33, offset: 15225},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 592, col: 33, offset: 15225},
	name: "_",
},
&litMatcher{
	pos: position{line: 592, col: 35, offset: 15227},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 592, col: 39, offset: 15231},
	name: "_",
},
&litMatcher{
	pos: position{line: 592, col: 41, offset: 15233},
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
	pos: position{line: 603, col: 1, offset: 15410},
	expr: &actionExpr{
	pos: position{line: 603, col: 13, offset: 15422},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 603, col: 13, offset: 15422},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 603, col: 13, offset: 15422},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 603, col: 22, offset: 15431},
	expr: &seqExpr{
	pos: position{line: 603, col: 23, offset: 15432},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 603, col: 23, offset: 15432},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 603, col: 27, offset: 15436},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 603, col: 31, offset: 15440},
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
	pos: position{line: 609, col: 1, offset: 15541},
	expr: &actionExpr{
	pos: position{line: 609, col: 10, offset: 15550},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 609, col: 10, offset: 15550},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 609, col: 10, offset: 15550},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 609, col: 16, offset: 15556},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 609, col: 23, offset: 15563},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 609, col: 28, offset: 15568},
	expr: &seqExpr{
	pos: position{line: 609, col: 29, offset: 15569},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 609, col: 29, offset: 15569},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 609, col: 31, offset: 15571},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 609, col: 36, offset: 15576},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 609, col: 38, offset: 15578},
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
	pos: position{line: 613, col: 1, offset: 15643},
	expr: &actionExpr{
	pos: position{line: 613, col: 11, offset: 15653},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 613, col: 11, offset: 15653},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 613, col: 16, offset: 15658},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 613, col: 16, offset: 15658},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 613, col: 24, offset: 15666},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 613, col: 33, offset: 15675},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 613, col: 48, offset: 15690},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 613, col: 59, offset: 15701},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 617, col: 1, offset: 15743},
	expr: &choiceExpr{
	pos: position{line: 617, col: 9, offset: 15751},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 617, col: 9, offset: 15751},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 617, col: 15, offset: 15757},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 617, col: 21, offset: 15763},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 617, col: 27, offset: 15769},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 617, col: 33, offset: 15775},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 617, col: 33, offset: 15775},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 621, col: 1, offset: 15823},
	expr: &choiceExpr{
	pos: position{line: 621, col: 12, offset: 15834},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 621, col: 12, offset: 15834},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 621, col: 12, offset: 15834},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 621, col: 12, offset: 15834},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 621, col: 16, offset: 15838},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 621, col: 19, offset: 15841},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 621, col: 25, offset: 15847},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 625, col: 5, offset: 15923},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 625, col: 5, offset: 15923},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 625, col: 8, offset: 15926},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 632, col: 1, offset: 16003},
	expr: &actionExpr{
	pos: position{line: 632, col: 9, offset: 16011},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 632, col: 9, offset: 16011},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 632, col: 9, offset: 16011},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 632, col: 15, offset: 16017},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 632, col: 21, offset: 16023},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 632, col: 26, offset: 16028},
	expr: &seqExpr{
	pos: position{line: 632, col: 27, offset: 16029},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 632, col: 27, offset: 16029},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 632, col: 29, offset: 16031},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 632, col: 34, offset: 16036},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 632, col: 36, offset: 16038},
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
	pos: position{line: 636, col: 1, offset: 16102},
	expr: &actionExpr{
	pos: position{line: 636, col: 10, offset: 16111},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 636, col: 10, offset: 16111},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 636, col: 15, offset: 16116},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 636, col: 15, offset: 16116},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 636, col: 23, offset: 16124},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 636, col: 32, offset: 16133},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 636, col: 47, offset: 16148},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 640, col: 1, offset: 16191},
	expr: &actionExpr{
	pos: position{line: 640, col: 13, offset: 16203},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 640, col: 13, offset: 16203},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 640, col: 13, offset: 16203},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 640, col: 20, offset: 16210},
	expr: &choiceExpr{
	pos: position{line: 640, col: 21, offset: 16211},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 640, col: 21, offset: 16211},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 640, col: 31, offset: 16221},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 640, col: 39, offset: 16229},
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
	pos: position{line: 644, col: 1, offset: 16279},
	expr: &actionExpr{
	pos: position{line: 644, col: 17, offset: 16295},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 644, col: 17, offset: 16295},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 644, col: 17, offset: 16295},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 644, col: 21, offset: 16299},
	expr: &choiceExpr{
	pos: position{line: 644, col: 23, offset: 16301},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 644, col: 23, offset: 16301},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 644, col: 23, offset: 16301},
	expr: &ruleRefExpr{
	pos: position{line: 644, col: 24, offset: 16302},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 644, col: 36, offset: 16314,
},
	},
},
&seqExpr{
	pos: position{line: 644, col: 40, offset: 16318},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 644, col: 40, offset: 16318},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 644, col: 45, offset: 16323},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 644, col: 63, offset: 16341},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 651, col: 1, offset: 16522},
	expr: &charClassMatcher{
	pos: position{line: 651, col: 16, offset: 16537},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 653, col: 1, offset: 16555},
	expr: &ruleRefExpr{
	pos: position{line: 653, col: 19, offset: 16573},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 655, col: 1, offset: 16593},
	expr: &charClassMatcher{
	pos: position{line: 655, col: 21, offset: 16613},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 657, col: 1, offset: 16628},
	expr: &actionExpr{
	pos: position{line: 657, col: 9, offset: 16636},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 657, col: 9, offset: 16636},
	expr: &choiceExpr{
	pos: position{line: 657, col: 10, offset: 16637},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 657, col: 10, offset: 16637},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 657, col: 19, offset: 16646},
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
	pos: position{line: 661, col: 1, offset: 16696},
	expr: &choiceExpr{
	pos: position{line: 661, col: 11, offset: 16706},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 661, col: 11, offset: 16706},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 661, col: 19, offset: 16714},
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
	pos: position{line: 663, col: 1, offset: 16723},
	expr: &actionExpr{
	pos: position{line: 663, col: 11, offset: 16733},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 663, col: 11, offset: 16733},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 663, col: 11, offset: 16733},
	expr: &litMatcher{
	pos: position{line: 663, col: 11, offset: 16733},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 663, col: 16, offset: 16738},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 667, col: 1, offset: 16811},
	expr: &actionExpr{
	pos: position{line: 667, col: 19, offset: 16829},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 667, col: 19, offset: 16829},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 671, col: 1, offset: 16902},
	expr: &actionExpr{
	pos: position{line: 671, col: 10, offset: 16911},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 671, col: 10, offset: 16911},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 671, col: 10, offset: 16911},
	expr: &litMatcher{
	pos: position{line: 671, col: 10, offset: 16911},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 671, col: 15, offset: 16916},
	expr: &ruleRefExpr{
	pos: position{line: 671, col: 15, offset: 16916},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 671, col: 24, offset: 16925},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 671, col: 28, offset: 16929},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 675, col: 1, offset: 16996},
	expr: &choiceExpr{
	pos: position{line: 675, col: 12, offset: 17007},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 675, col: 12, offset: 17007},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 675, col: 18, offset: 17013},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 675, col: 18, offset: 17013},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 675, col: 38, offset: 17033},
	expr: &ruleRefExpr{
	pos: position{line: 675, col: 38, offset: 17033},
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
	pos: position{line: 677, col: 1, offset: 17050},
	expr: &charClassMatcher{
	pos: position{line: 677, col: 17, offset: 17066},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 679, col: 1, offset: 17075},
	expr: &charClassMatcher{
	pos: position{line: 679, col: 24, offset: 17098},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 681, col: 1, offset: 17107},
	expr: &zeroOrMoreExpr{
	pos: position{line: 681, col: 19, offset: 17125},
	expr: &charClassMatcher{
	pos: position{line: 681, col: 19, offset: 17125},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 683, col: 1, offset: 17139},
	expr: &charClassMatcher{
	pos: position{line: 683, col: 18, offset: 17156},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "IDENTIFIER",
	pos: position{line: 684, col: 1, offset: 17167},
	expr: &actionExpr{
	pos: position{line: 684, col: 15, offset: 17181},
	run: (*parser).callonIDENTIFIER1,
	expr: &seqExpr{
	pos: position{line: 684, col: 15, offset: 17181},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 684, col: 15, offset: 17181},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 684, col: 22, offset: 17188},
	expr: &choiceExpr{
	pos: position{line: 684, col: 23, offset: 17189},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 684, col: 23, offset: 17189},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 684, col: 33, offset: 17199},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 684, col: 41, offset: 17207},
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
	pos: position{line: 687, col: 1, offset: 17280},
	expr: &notExpr{
	pos: position{line: 687, col: 8, offset: 17287},
	expr: &anyMatcher{
	line: 687, col: 9, offset: 17288,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 689, col: 1, offset: 17293},
	expr: &litMatcher{
	pos: position{line: 689, col: 17, offset: 17309},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 695, col: 1, offset: 17396},
	expr: &actionExpr{
	pos: position{line: 695, col: 18, offset: 17413},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 695, col: 18, offset: 17413},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 695, col: 18, offset: 17413},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 695, col: 24, offset: 17419},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 695, col: 33, offset: 17428},
	name: "_",
},
&labeledExpr{
	pos: position{line: 695, col: 35, offset: 17430},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 695, col: 40, offset: 17435},
	expr: &seqExpr{
	pos: position{line: 695, col: 41, offset: 17436},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 695, col: 41, offset: 17436},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 695, col: 45, offset: 17440},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 695, col: 47, offset: 17442},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 695, col: 56, offset: 17451},
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
	pos: position{line: 698, col: 1, offset: 17504},
	expr: &actionExpr{
	pos: position{line: 698, col: 13, offset: 17516},
	run: (*parser).callonEqAliasG1,
	expr: &seqExpr{
	pos: position{line: 698, col: 13, offset: 17516},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 698, col: 13, offset: 17516},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 698, col: 19, offset: 17522},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 698, col: 28, offset: 17531},
	name: "_",
},
&litMatcher{
	pos: position{line: 698, col: 30, offset: 17533},
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
        return SourceJob{Type: FILEJOB,OperateOn:SrcFile{fileName:fname}},nil
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
    return SourceJob{Type:BRANCHJOB,OperateOn:SrcBranch{identifier:branchName}},nil
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
                return TransformJob{Type:DOJOB,OperateOn:f}, nil
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
        OperateOn:InTo{streamTo:streamTo},
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

