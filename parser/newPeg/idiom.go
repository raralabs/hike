
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
	expr: &choiceExpr{
	pos: position{line: 114, col: 12, offset: 2723},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 114, col: 12, offset: 2723},
	name: "CountAgg",
},
&ruleRefExpr{
	pos: position{line: 114, col: 23, offset: 2734},
	name: "MaxAgg",
},
&ruleRefExpr{
	pos: position{line: 114, col: 32, offset: 2743},
	name: "MinAgg",
},
&ruleRefExpr{
	pos: position{line: 114, col: 41, offset: 2752},
	name: "AvgAgg",
},
&ruleRefExpr{
	pos: position{line: 114, col: 50, offset: 2761},
	name: "SumAgg",
},
&ruleRefExpr{
	pos: position{line: 114, col: 59, offset: 2770},
	name: "ModeAgg",
},
&ruleRefExpr{
	pos: position{line: 115, col: 13, offset: 2791},
	name: "VarAgg",
},
&ruleRefExpr{
	pos: position{line: 115, col: 22, offset: 2800},
	name: "DistinctCountAgg",
},
&ruleRefExpr{
	pos: position{line: 115, col: 41, offset: 2819},
	name: "QuantileAgg",
},
&ruleRefExpr{
	pos: position{line: 115, col: 55, offset: 2833},
	name: "MedianAgg",
},
&ruleRefExpr{
	pos: position{line: 115, col: 67, offset: 2845},
	name: "QuartileAgg",
},
&ruleRefExpr{
	pos: position{line: 116, col: 13, offset: 2870},
	name: "CovAgg",
},
&ruleRefExpr{
	pos: position{line: 116, col: 22, offset: 2879},
	name: "CorrAgg",
},
&ruleRefExpr{
	pos: position{line: 116, col: 32, offset: 2889},
	name: "PValueAgg",
},
	},
},
},
{
	name: "CountAgg",
	pos: position{line: 119, col: 1, offset: 2933},
	expr: &actionExpr{
	pos: position{line: 119, col: 13, offset: 2945},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 119, col: 13, offset: 2945},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 119, col: 13, offset: 2945},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 119, col: 19, offset: 2951},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 119, col: 28, offset: 2960},
	name: "_",
},
&litMatcher{
	pos: position{line: 119, col: 30, offset: 2962},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 119, col: 39, offset: 2971},
	name: "_",
},
&labeledExpr{
	pos: position{line: 119, col: 41, offset: 2973},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 119, col: 46, offset: 2978},
	expr: &ruleRefExpr{
	pos: position{line: 119, col: 46, offset: 2978},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 119, col: 59, offset: 2991},
	name: "_",
},
&litMatcher{
	pos: position{line: 119, col: 61, offset: 2993},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MaxAgg",
	pos: position{line: 126, col: 1, offset: 3145},
	expr: &actionExpr{
	pos: position{line: 126, col: 11, offset: 3155},
	run: (*parser).callonMaxAgg1,
	expr: &seqExpr{
	pos: position{line: 126, col: 11, offset: 3155},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 126, col: 11, offset: 3155},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 126, col: 17, offset: 3161},
	expr: &ruleRefExpr{
	pos: position{line: 126, col: 17, offset: 3161},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 126, col: 27, offset: 3171},
	name: "_",
},
&litMatcher{
	pos: position{line: 126, col: 29, offset: 3173},
	val: "max(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 126, col: 36, offset: 3180},
	name: "_",
},
&labeledExpr{
	pos: position{line: 126, col: 38, offset: 3182},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 126, col: 44, offset: 3188},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 126, col: 53, offset: 3197},
	name: "_",
},
&labeledExpr{
	pos: position{line: 126, col: 55, offset: 3199},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 126, col: 60, offset: 3204},
	expr: &seqExpr{
	pos: position{line: 126, col: 61, offset: 3205},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 126, col: 61, offset: 3205},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 126, col: 65, offset: 3209},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 126, col: 67, offset: 3211},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 126, col: 81, offset: 3225},
	name: "_",
},
&litMatcher{
	pos: position{line: 126, col: 83, offset: 3227},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MinAgg",
	pos: position{line: 148, col: 1, offset: 3688},
	expr: &actionExpr{
	pos: position{line: 148, col: 11, offset: 3698},
	run: (*parser).callonMinAgg1,
	expr: &seqExpr{
	pos: position{line: 148, col: 11, offset: 3698},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 148, col: 11, offset: 3698},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 148, col: 17, offset: 3704},
	expr: &ruleRefExpr{
	pos: position{line: 148, col: 17, offset: 3704},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 148, col: 27, offset: 3714},
	name: "_",
},
&litMatcher{
	pos: position{line: 148, col: 29, offset: 3716},
	val: "min(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 148, col: 36, offset: 3723},
	name: "_",
},
&labeledExpr{
	pos: position{line: 148, col: 38, offset: 3725},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 148, col: 44, offset: 3731},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 148, col: 53, offset: 3740},
	name: "_",
},
&labeledExpr{
	pos: position{line: 148, col: 55, offset: 3742},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 148, col: 60, offset: 3747},
	expr: &seqExpr{
	pos: position{line: 148, col: 61, offset: 3748},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 148, col: 61, offset: 3748},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 148, col: 65, offset: 3752},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 148, col: 67, offset: 3754},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 148, col: 81, offset: 3768},
	name: "_",
},
&litMatcher{
	pos: position{line: 148, col: 83, offset: 3770},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AvgAgg",
	pos: position{line: 170, col: 1, offset: 4231},
	expr: &actionExpr{
	pos: position{line: 170, col: 11, offset: 4241},
	run: (*parser).callonAvgAgg1,
	expr: &seqExpr{
	pos: position{line: 170, col: 11, offset: 4241},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 170, col: 11, offset: 4241},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 170, col: 17, offset: 4247},
	expr: &ruleRefExpr{
	pos: position{line: 170, col: 17, offset: 4247},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 170, col: 27, offset: 4257},
	name: "_",
},
&litMatcher{
	pos: position{line: 170, col: 29, offset: 4259},
	val: "avg(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 170, col: 36, offset: 4266},
	name: "_",
},
&labeledExpr{
	pos: position{line: 170, col: 38, offset: 4268},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 170, col: 44, offset: 4274},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 170, col: 53, offset: 4283},
	name: "_",
},
&labeledExpr{
	pos: position{line: 170, col: 55, offset: 4285},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 170, col: 60, offset: 4290},
	expr: &seqExpr{
	pos: position{line: 170, col: 61, offset: 4291},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 170, col: 61, offset: 4291},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 170, col: 65, offset: 4295},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 170, col: 67, offset: 4297},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 170, col: 81, offset: 4311},
	name: "_",
},
&litMatcher{
	pos: position{line: 170, col: 83, offset: 4313},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SumAgg",
	pos: position{line: 192, col: 1, offset: 4774},
	expr: &actionExpr{
	pos: position{line: 192, col: 11, offset: 4784},
	run: (*parser).callonSumAgg1,
	expr: &seqExpr{
	pos: position{line: 192, col: 11, offset: 4784},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 192, col: 11, offset: 4784},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 192, col: 17, offset: 4790},
	expr: &ruleRefExpr{
	pos: position{line: 192, col: 17, offset: 4790},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 192, col: 27, offset: 4800},
	name: "_",
},
&litMatcher{
	pos: position{line: 192, col: 29, offset: 4802},
	val: "sum(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 192, col: 36, offset: 4809},
	name: "_",
},
&labeledExpr{
	pos: position{line: 192, col: 38, offset: 4811},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 192, col: 44, offset: 4817},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 192, col: 53, offset: 4826},
	name: "_",
},
&labeledExpr{
	pos: position{line: 192, col: 55, offset: 4828},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 192, col: 60, offset: 4833},
	expr: &seqExpr{
	pos: position{line: 192, col: 61, offset: 4834},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 192, col: 61, offset: 4834},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 192, col: 65, offset: 4838},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 192, col: 67, offset: 4840},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 192, col: 81, offset: 4854},
	name: "_",
},
&litMatcher{
	pos: position{line: 192, col: 83, offset: 4856},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarAgg",
	pos: position{line: 214, col: 1, offset: 5317},
	expr: &actionExpr{
	pos: position{line: 214, col: 11, offset: 5327},
	run: (*parser).callonVarAgg1,
	expr: &seqExpr{
	pos: position{line: 214, col: 11, offset: 5327},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 214, col: 11, offset: 5327},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 214, col: 17, offset: 5333},
	expr: &ruleRefExpr{
	pos: position{line: 214, col: 17, offset: 5333},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 214, col: 27, offset: 5343},
	name: "_",
},
&litMatcher{
	pos: position{line: 214, col: 29, offset: 5345},
	val: "var(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 214, col: 36, offset: 5352},
	name: "_",
},
&labeledExpr{
	pos: position{line: 214, col: 38, offset: 5354},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 214, col: 44, offset: 5360},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 214, col: 53, offset: 5369},
	name: "_",
},
&labeledExpr{
	pos: position{line: 214, col: 55, offset: 5371},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 214, col: 60, offset: 5376},
	expr: &seqExpr{
	pos: position{line: 214, col: 61, offset: 5377},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 214, col: 61, offset: 5377},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 214, col: 65, offset: 5381},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 214, col: 67, offset: 5383},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 214, col: 81, offset: 5397},
	name: "_",
},
&litMatcher{
	pos: position{line: 214, col: 83, offset: 5399},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DistinctCountAgg",
	pos: position{line: 236, col: 1, offset: 5865},
	expr: &actionExpr{
	pos: position{line: 236, col: 21, offset: 5885},
	run: (*parser).callonDistinctCountAgg1,
	expr: &seqExpr{
	pos: position{line: 236, col: 21, offset: 5885},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 236, col: 21, offset: 5885},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 236, col: 27, offset: 5891},
	expr: &ruleRefExpr{
	pos: position{line: 236, col: 27, offset: 5891},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 236, col: 37, offset: 5901},
	name: "_",
},
&litMatcher{
	pos: position{line: 236, col: 39, offset: 5903},
	val: "dcount(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 236, col: 49, offset: 5913},
	name: "_",
},
&labeledExpr{
	pos: position{line: 236, col: 51, offset: 5915},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 236, col: 57, offset: 5921},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 236, col: 66, offset: 5930},
	name: "_",
},
&labeledExpr{
	pos: position{line: 236, col: 68, offset: 5932},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 236, col: 73, offset: 5937},
	expr: &seqExpr{
	pos: position{line: 236, col: 74, offset: 5938},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 236, col: 74, offset: 5938},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 236, col: 78, offset: 5942},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 236, col: 80, offset: 5944},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 236, col: 94, offset: 5958},
	name: "_",
},
&litMatcher{
	pos: position{line: 236, col: 96, offset: 5960},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "ModeAgg",
	pos: position{line: 258, col: 1, offset: 6431},
	expr: &actionExpr{
	pos: position{line: 258, col: 12, offset: 6442},
	run: (*parser).callonModeAgg1,
	expr: &seqExpr{
	pos: position{line: 258, col: 12, offset: 6442},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 258, col: 12, offset: 6442},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 258, col: 18, offset: 6448},
	expr: &ruleRefExpr{
	pos: position{line: 258, col: 18, offset: 6448},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 258, col: 28, offset: 6458},
	name: "_",
},
&litMatcher{
	pos: position{line: 258, col: 30, offset: 6460},
	val: "mode(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 258, col: 38, offset: 6468},
	name: "_",
},
&labeledExpr{
	pos: position{line: 258, col: 40, offset: 6470},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 258, col: 46, offset: 6476},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 258, col: 55, offset: 6485},
	name: "_",
},
&labeledExpr{
	pos: position{line: 258, col: 57, offset: 6487},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 258, col: 62, offset: 6492},
	expr: &seqExpr{
	pos: position{line: 258, col: 63, offset: 6493},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 258, col: 63, offset: 6493},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 258, col: 67, offset: 6497},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 258, col: 69, offset: 6499},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 258, col: 83, offset: 6513},
	name: "_",
},
&litMatcher{
	pos: position{line: 258, col: 85, offset: 6515},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuantileAgg",
	pos: position{line: 280, col: 1, offset: 6977},
	expr: &actionExpr{
	pos: position{line: 280, col: 16, offset: 6992},
	run: (*parser).callonQuantileAgg1,
	expr: &seqExpr{
	pos: position{line: 280, col: 16, offset: 6992},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 280, col: 16, offset: 6992},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 280, col: 22, offset: 6998},
	expr: &ruleRefExpr{
	pos: position{line: 280, col: 22, offset: 6998},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 280, col: 32, offset: 7008},
	name: "_",
},
&litMatcher{
	pos: position{line: 280, col: 34, offset: 7010},
	val: "quantile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 280, col: 46, offset: 7022},
	name: "_",
},
&labeledExpr{
	pos: position{line: 280, col: 48, offset: 7024},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 280, col: 54, offset: 7030},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 280, col: 63, offset: 7039},
	name: "_",
},
&labeledExpr{
	pos: position{line: 280, col: 66, offset: 7042},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 280, col: 73, offset: 7049},
	expr: &seqExpr{
	pos: position{line: 280, col: 74, offset: 7050},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 280, col: 74, offset: 7050},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 280, col: 78, offset: 7054},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 280, col: 80, offset: 7056},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 280, col: 91, offset: 7067},
	name: "_",
},
&litMatcher{
	pos: position{line: 280, col: 93, offset: 7069},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 280, col: 97, offset: 7073},
	name: "_",
},
&labeledExpr{
	pos: position{line: 280, col: 99, offset: 7075},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 280, col: 103, offset: 7079},
	name: "Float",
},
},
&ruleRefExpr{
	pos: position{line: 280, col: 109, offset: 7085},
	name: "_",
},
&labeledExpr{
	pos: position{line: 280, col: 111, offset: 7087},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 280, col: 116, offset: 7092},
	expr: &seqExpr{
	pos: position{line: 280, col: 117, offset: 7093},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 280, col: 117, offset: 7093},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 280, col: 121, offset: 7097},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 280, col: 123, offset: 7099},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 280, col: 137, offset: 7113},
	name: "_",
},
&litMatcher{
	pos: position{line: 280, col: 139, offset: 7115},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MedianAgg",
	pos: position{line: 310, col: 1, offset: 7760},
	expr: &actionExpr{
	pos: position{line: 310, col: 14, offset: 7773},
	run: (*parser).callonMedianAgg1,
	expr: &seqExpr{
	pos: position{line: 310, col: 14, offset: 7773},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 310, col: 14, offset: 7773},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 310, col: 20, offset: 7779},
	expr: &ruleRefExpr{
	pos: position{line: 310, col: 20, offset: 7779},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 310, col: 30, offset: 7789},
	name: "_",
},
&litMatcher{
	pos: position{line: 310, col: 32, offset: 7791},
	val: "median(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 310, col: 42, offset: 7801},
	name: "_",
},
&labeledExpr{
	pos: position{line: 310, col: 44, offset: 7803},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 310, col: 50, offset: 7809},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 310, col: 59, offset: 7818},
	name: "_",
},
&labeledExpr{
	pos: position{line: 310, col: 62, offset: 7821},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 310, col: 69, offset: 7828},
	expr: &seqExpr{
	pos: position{line: 310, col: 70, offset: 7829},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 310, col: 70, offset: 7829},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 310, col: 74, offset: 7833},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 310, col: 76, offset: 7835},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 310, col: 87, offset: 7846},
	name: "_",
},
&labeledExpr{
	pos: position{line: 310, col: 89, offset: 7848},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 310, col: 94, offset: 7853},
	expr: &seqExpr{
	pos: position{line: 310, col: 95, offset: 7854},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 310, col: 95, offset: 7854},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 310, col: 99, offset: 7858},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 310, col: 101, offset: 7860},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 310, col: 115, offset: 7874},
	name: "_",
},
&litMatcher{
	pos: position{line: 310, col: 117, offset: 7876},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuartileAgg",
	pos: position{line: 342, col: 1, offset: 8553},
	expr: &actionExpr{
	pos: position{line: 342, col: 16, offset: 8568},
	run: (*parser).callonQuartileAgg1,
	expr: &seqExpr{
	pos: position{line: 342, col: 16, offset: 8568},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 342, col: 16, offset: 8568},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 342, col: 22, offset: 8574},
	expr: &ruleRefExpr{
	pos: position{line: 342, col: 22, offset: 8574},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 342, col: 32, offset: 8584},
	name: "_",
},
&litMatcher{
	pos: position{line: 342, col: 34, offset: 8586},
	val: "quartile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 342, col: 46, offset: 8598},
	name: "_",
},
&labeledExpr{
	pos: position{line: 342, col: 48, offset: 8600},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 342, col: 54, offset: 8606},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 342, col: 63, offset: 8615},
	name: "_",
},
&labeledExpr{
	pos: position{line: 342, col: 66, offset: 8618},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 342, col: 73, offset: 8625},
	expr: &seqExpr{
	pos: position{line: 342, col: 74, offset: 8626},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 342, col: 74, offset: 8626},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 342, col: 78, offset: 8630},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 342, col: 80, offset: 8632},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 342, col: 91, offset: 8643},
	name: "_",
},
&litMatcher{
	pos: position{line: 342, col: 93, offset: 8645},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 342, col: 97, offset: 8649},
	name: "_",
},
&labeledExpr{
	pos: position{line: 342, col: 99, offset: 8651},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 342, col: 103, offset: 8655},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 342, col: 110, offset: 8662},
	name: "_",
},
&labeledExpr{
	pos: position{line: 342, col: 112, offset: 8664},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 342, col: 117, offset: 8669},
	expr: &seqExpr{
	pos: position{line: 342, col: 118, offset: 8670},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 342, col: 118, offset: 8670},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 342, col: 122, offset: 8674},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 342, col: 124, offset: 8676},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 342, col: 138, offset: 8690},
	name: "_",
},
&litMatcher{
	pos: position{line: 342, col: 140, offset: 8692},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CovAgg",
	pos: position{line: 388, col: 1, offset: 9719},
	expr: &actionExpr{
	pos: position{line: 388, col: 11, offset: 9729},
	run: (*parser).callonCovAgg1,
	expr: &seqExpr{
	pos: position{line: 388, col: 11, offset: 9729},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 388, col: 11, offset: 9729},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 388, col: 17, offset: 9735},
	expr: &ruleRefExpr{
	pos: position{line: 388, col: 17, offset: 9735},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 388, col: 27, offset: 9745},
	name: "_",
},
&litMatcher{
	pos: position{line: 388, col: 29, offset: 9747},
	val: "cov(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 388, col: 36, offset: 9754},
	name: "_",
},
&labeledExpr{
	pos: position{line: 388, col: 38, offset: 9756},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 388, col: 45, offset: 9763},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 388, col: 54, offset: 9772},
	name: "_",
},
&litMatcher{
	pos: position{line: 388, col: 56, offset: 9774},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 388, col: 60, offset: 9778},
	name: "_",
},
&labeledExpr{
	pos: position{line: 388, col: 62, offset: 9780},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 388, col: 69, offset: 9787},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 388, col: 78, offset: 9796},
	name: "_",
},
&labeledExpr{
	pos: position{line: 388, col: 80, offset: 9798},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 388, col: 85, offset: 9803},
	expr: &seqExpr{
	pos: position{line: 388, col: 86, offset: 9804},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 388, col: 86, offset: 9804},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 388, col: 90, offset: 9808},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 388, col: 92, offset: 9810},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 388, col: 106, offset: 9824},
	name: "_",
},
&litMatcher{
	pos: position{line: 388, col: 108, offset: 9826},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CorrAgg",
	pos: position{line: 411, col: 1, offset: 10351},
	expr: &actionExpr{
	pos: position{line: 411, col: 12, offset: 10362},
	run: (*parser).callonCorrAgg1,
	expr: &seqExpr{
	pos: position{line: 411, col: 12, offset: 10362},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 411, col: 12, offset: 10362},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 411, col: 18, offset: 10368},
	expr: &ruleRefExpr{
	pos: position{line: 411, col: 18, offset: 10368},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 411, col: 28, offset: 10378},
	name: "_",
},
&litMatcher{
	pos: position{line: 411, col: 30, offset: 10380},
	val: "correlation(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 411, col: 45, offset: 10395},
	name: "_",
},
&labeledExpr{
	pos: position{line: 411, col: 47, offset: 10397},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 411, col: 54, offset: 10404},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 411, col: 63, offset: 10413},
	name: "_",
},
&litMatcher{
	pos: position{line: 411, col: 65, offset: 10415},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 411, col: 69, offset: 10419},
	name: "_",
},
&labeledExpr{
	pos: position{line: 411, col: 71, offset: 10421},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 411, col: 78, offset: 10428},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 411, col: 87, offset: 10437},
	name: "_",
},
&labeledExpr{
	pos: position{line: 411, col: 89, offset: 10439},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 411, col: 94, offset: 10444},
	expr: &seqExpr{
	pos: position{line: 411, col: 95, offset: 10445},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 411, col: 95, offset: 10445},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 411, col: 99, offset: 10449},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 411, col: 101, offset: 10451},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 411, col: 115, offset: 10465},
	name: "_",
},
&litMatcher{
	pos: position{line: 411, col: 117, offset: 10467},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PValueAgg",
	pos: position{line: 434, col: 1, offset: 10993},
	expr: &actionExpr{
	pos: position{line: 434, col: 14, offset: 11006},
	run: (*parser).callonPValueAgg1,
	expr: &seqExpr{
	pos: position{line: 434, col: 14, offset: 11006},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 434, col: 14, offset: 11006},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 434, col: 20, offset: 11012},
	expr: &ruleRefExpr{
	pos: position{line: 434, col: 20, offset: 11012},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 434, col: 30, offset: 11022},
	name: "_",
},
&litMatcher{
	pos: position{line: 434, col: 32, offset: 11024},
	val: "pvalue(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 434, col: 42, offset: 11034},
	name: "_",
},
&labeledExpr{
	pos: position{line: 434, col: 44, offset: 11036},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 434, col: 51, offset: 11043},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 434, col: 60, offset: 11052},
	name: "_",
},
&litMatcher{
	pos: position{line: 434, col: 62, offset: 11054},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 434, col: 66, offset: 11058},
	name: "_",
},
&labeledExpr{
	pos: position{line: 434, col: 68, offset: 11060},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 434, col: 75, offset: 11067},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 434, col: 84, offset: 11076},
	name: "_",
},
&labeledExpr{
	pos: position{line: 434, col: 86, offset: 11078},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 434, col: 91, offset: 11083},
	expr: &seqExpr{
	pos: position{line: 434, col: 92, offset: 11084},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 434, col: 92, offset: 11084},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 434, col: 96, offset: 11088},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 434, col: 98, offset: 11090},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 434, col: 112, offset: 11104},
	name: "_",
},
&litMatcher{
	pos: position{line: 434, col: 114, offset: 11106},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggGroupBy",
	pos: position{line: 459, col: 1, offset: 11631},
	expr: &actionExpr{
	pos: position{line: 459, col: 15, offset: 11645},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 459, col: 15, offset: 11645},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 459, col: 15, offset: 11645},
	name: "_",
},
&litMatcher{
	pos: position{line: 459, col: 17, offset: 11647},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 459, col: 22, offset: 11652},
	name: "_",
},
&labeledExpr{
	pos: position{line: 459, col: 24, offset: 11654},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 459, col: 30, offset: 11660},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 459, col: 39, offset: 11669},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 459, col: 44, offset: 11674},
	expr: &seqExpr{
	pos: position{line: 459, col: 45, offset: 11675},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 459, col: 45, offset: 11675},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 459, col: 49, offset: 11679},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 459, col: 51, offset: 11681},
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
	pos: position{line: 474, col: 1, offset: 12032},
	expr: &actionExpr{
	pos: position{line: 474, col: 11, offset: 12042},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 474, col: 11, offset: 12042},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 474, col: 16, offset: 12047},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 474, col: 16, offset: 12047},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 474, col: 27, offset: 12058},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 474, col: 38, offset: 12069},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 474, col: 49, offset: 12080},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 474, col: 58, offset: 12089},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 474, col: 67, offset: 12098},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 480, col: 1, offset: 12182},
	expr: &actionExpr{
	pos: position{line: 480, col: 13, offset: 12194},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 480, col: 13, offset: 12194},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 480, col: 13, offset: 12194},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 480, col: 23, offset: 12204},
	name: "_",
},
&labeledExpr{
	pos: position{line: 480, col: 25, offset: 12206},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 480, col: 28, offset: 12209},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 480, col: 42, offset: 12223},
	name: "_",
},
&litMatcher{
	pos: position{line: 480, col: 44, offset: 12225},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 487, col: 1, offset: 12420},
	expr: &actionExpr{
	pos: position{line: 487, col: 18, offset: 12437},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 487, col: 18, offset: 12437},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 487, col: 18, offset: 12437},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 487, col: 23, offset: 12442},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 487, col: 28, offset: 12447},
	name: "_",
},
&labeledExpr{
	pos: position{line: 487, col: 30, offset: 12449},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 487, col: 33, offset: 12452},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 487, col: 45, offset: 12464},
	name: "_",
},
&labeledExpr{
	pos: position{line: 487, col: 47, offset: 12466},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 487, col: 54, offset: 12473},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 491, col: 1, offset: 12524},
	expr: &ruleRefExpr{
	pos: position{line: 491, col: 16, offset: 12539},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 494, col: 1, offset: 12573},
	expr: &actionExpr{
	pos: position{line: 494, col: 13, offset: 12585},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 494, col: 13, offset: 12585},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 494, col: 13, offset: 12585},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 494, col: 23, offset: 12595},
	name: "_",
},
&labeledExpr{
	pos: position{line: 494, col: 25, offset: 12597},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 494, col: 31, offset: 12603},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 494, col: 40, offset: 12612},
	name: "_",
},
&labeledExpr{
	pos: position{line: 494, col: 42, offset: 12614},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 494, col: 47, offset: 12619},
	expr: &seqExpr{
	pos: position{line: 494, col: 48, offset: 12620},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 494, col: 48, offset: 12620},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 494, col: 52, offset: 12624},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 494, col: 54, offset: 12626},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 494, col: 65, offset: 12637},
	name: "_",
},
&litMatcher{
	pos: position{line: 494, col: 67, offset: 12639},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 509, col: 1, offset: 12976},
	expr: &actionExpr{
	pos: position{line: 509, col: 11, offset: 12986},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 509, col: 11, offset: 12986},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 509, col: 11, offset: 12986},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 509, col: 18, offset: 12993},
	name: "_",
},
&labeledExpr{
	pos: position{line: 509, col: 20, offset: 12995},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 509, col: 25, offset: 13000},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 509, col: 34, offset: 13009},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 509, col: 38, offset: 13013},
	name: "_",
},
&labeledExpr{
	pos: position{line: 509, col: 40, offset: 13015},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 509, col: 44, offset: 13019},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 509, col: 51, offset: 13026},
	name: "_",
},
&litMatcher{
	pos: position{line: 509, col: 53, offset: 13028},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 521, col: 1, offset: 13250},
	expr: &actionExpr{
	pos: position{line: 521, col: 11, offset: 13260},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 521, col: 11, offset: 13260},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 521, col: 11, offset: 13260},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 521, col: 20, offset: 13269},
	name: "_",
},
&litMatcher{
	pos: position{line: 521, col: 22, offset: 13271},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 521, col: 27, offset: 13276},
	name: "_",
},
&labeledExpr{
	pos: position{line: 521, col: 29, offset: 13278},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 521, col: 35, offset: 13284},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 527, col: 1, offset: 13390},
	expr: &actionExpr{
	pos: position{line: 527, col: 12, offset: 13401},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 527, col: 12, offset: 13401},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 527, col: 12, offset: 13401},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 527, col: 21, offset: 13410},
	name: "_",
},
&labeledExpr{
	pos: position{line: 527, col: 23, offset: 13412},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 527, col: 27, offset: 13416},
	expr: &ruleRefExpr{
	pos: position{line: 527, col: 27, offset: 13416},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 527, col: 35, offset: 13424},
	name: "_",
},
&litMatcher{
	pos: position{line: 527, col: 37, offset: 13426},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 537, col: 1, offset: 13571},
	expr: &actionExpr{
	pos: position{line: 537, col: 13, offset: 13583},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 537, col: 13, offset: 13583},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 537, col: 13, offset: 13583},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 537, col: 23, offset: 13593},
	name: "_",
},
&labeledExpr{
	pos: position{line: 537, col: 25, offset: 13595},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 537, col: 30, offset: 13600},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 537, col: 42, offset: 13612},
	name: "_",
},
&litMatcher{
	pos: position{line: 537, col: 44, offset: 13614},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 547, col: 1, offset: 13902},
	expr: &actionExpr{
	pos: position{line: 547, col: 16, offset: 13917},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 547, col: 16, offset: 13917},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 547, col: 16, offset: 13917},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 547, col: 22, offset: 13923},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 547, col: 35, offset: 13936},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 547, col: 40, offset: 13941},
	expr: &seqExpr{
	pos: position{line: 547, col: 41, offset: 13942},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 547, col: 41, offset: 13942},
	name: "_",
},
&labeledExpr{
	pos: position{line: 547, col: 43, offset: 13944},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 547, col: 46, offset: 13947},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 547, col: 58, offset: 13959},
	name: "_",
},
&labeledExpr{
	pos: position{line: 547, col: 60, offset: 13961},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 547, col: 63, offset: 13964},
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
	pos: position{line: 552, col: 1, offset: 14122},
	expr: &choiceExpr{
	pos: position{line: 552, col: 17, offset: 14138},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 552, col: 17, offset: 14138},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 552, col: 17, offset: 14138},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 552, col: 17, offset: 14138},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 552, col: 21, offset: 14142},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 552, col: 24, offset: 14145},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 552, col: 36, offset: 14157},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 554, col: 5, offset: 14188},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 554, col: 5, offset: 14188},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 554, col: 8, offset: 14191},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 558, col: 1, offset: 14233},
	expr: &choiceExpr{
	pos: position{line: 558, col: 16, offset: 14248},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 558, col: 16, offset: 14248},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 558, col: 24, offset: 14256},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 558, col: 24, offset: 14256},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 565, col: 1, offset: 14435},
	expr: &actionExpr{
	pos: position{line: 565, col: 18, offset: 14452},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 565, col: 18, offset: 14452},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 565, col: 18, offset: 14452},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 565, col: 23, offset: 14457},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 565, col: 28, offset: 14462},
	name: "_",
},
&labeledExpr{
	pos: position{line: 565, col: 30, offset: 14464},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 565, col: 33, offset: 14467},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 565, col: 45, offset: 14479},
	name: "_",
},
&labeledExpr{
	pos: position{line: 565, col: 47, offset: 14481},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 565, col: 53, offset: 14487},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 565, col: 59, offset: 14493},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 569, col: 1, offset: 14567},
	expr: &choiceExpr{
	pos: position{line: 569, col: 16, offset: 14582},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 569, col: 16, offset: 14582},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 569, col: 23, offset: 14589},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 569, col: 30, offset: 14596},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 569, col: 36, offset: 14602},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 569, col: 42, offset: 14608},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 569, col: 49, offset: 14615},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 569, col: 49, offset: 14615},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 578, col: 1, offset: 14725},
	expr: &actionExpr{
	pos: position{line: 578, col: 12, offset: 14736},
	run: (*parser).callonStdOutG1,
	expr: &seqExpr{
	pos: position{line: 578, col: 12, offset: 14736},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 578, col: 12, offset: 14736},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 578, col: 22, offset: 14746},
	name: "_",
},
&labeledExpr{
	pos: position{line: 578, col: 24, offset: 14748},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 578, col: 31, offset: 14755},
	expr: &ruleRefExpr{
	pos: position{line: 578, col: 32, offset: 14756},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 578, col: 48, offset: 14772},
	name: "_",
},
&litMatcher{
	pos: position{line: 578, col: 50, offset: 14774},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SinkInto",
	pos: position{line: 592, col: 1, offset: 15020},
	expr: &actionExpr{
	pos: position{line: 592, col: 12, offset: 15031},
	run: (*parser).callonSinkInto1,
	expr: &seqExpr{
	pos: position{line: 592, col: 12, offset: 15031},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 592, col: 12, offset: 15031},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 592, col: 18, offset: 15037},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 592, col: 27, offset: 15046},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 592, col: 33, offset: 15052},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 592, col: 33, offset: 15052},
	name: "_",
},
&litMatcher{
	pos: position{line: 592, col: 35, offset: 15054},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 592, col: 39, offset: 15058},
	name: "_",
},
&litMatcher{
	pos: position{line: 592, col: 41, offset: 15060},
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
	pos: position{line: 602, col: 1, offset: 15196},
	expr: &actionExpr{
	pos: position{line: 602, col: 13, offset: 15208},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 602, col: 13, offset: 15208},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 602, col: 13, offset: 15208},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 602, col: 22, offset: 15217},
	expr: &seqExpr{
	pos: position{line: 602, col: 23, offset: 15218},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 602, col: 23, offset: 15218},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 602, col: 27, offset: 15222},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 602, col: 31, offset: 15226},
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
	pos: position{line: 608, col: 1, offset: 15327},
	expr: &actionExpr{
	pos: position{line: 608, col: 10, offset: 15336},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 608, col: 10, offset: 15336},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 608, col: 10, offset: 15336},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 608, col: 16, offset: 15342},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 608, col: 23, offset: 15349},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 608, col: 28, offset: 15354},
	expr: &seqExpr{
	pos: position{line: 608, col: 29, offset: 15355},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 608, col: 29, offset: 15355},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 608, col: 31, offset: 15357},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 608, col: 36, offset: 15362},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 608, col: 38, offset: 15364},
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
	pos: position{line: 612, col: 1, offset: 15429},
	expr: &actionExpr{
	pos: position{line: 612, col: 11, offset: 15439},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 612, col: 11, offset: 15439},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 612, col: 16, offset: 15444},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 612, col: 16, offset: 15444},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 612, col: 24, offset: 15452},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 612, col: 33, offset: 15461},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 612, col: 48, offset: 15476},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 612, col: 59, offset: 15487},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 616, col: 1, offset: 15529},
	expr: &choiceExpr{
	pos: position{line: 616, col: 9, offset: 15537},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 616, col: 9, offset: 15537},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 616, col: 15, offset: 15543},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 616, col: 21, offset: 15549},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 616, col: 27, offset: 15555},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 616, col: 33, offset: 15561},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 616, col: 33, offset: 15561},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 620, col: 1, offset: 15609},
	expr: &choiceExpr{
	pos: position{line: 620, col: 12, offset: 15620},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 620, col: 12, offset: 15620},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 620, col: 12, offset: 15620},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 620, col: 12, offset: 15620},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 620, col: 16, offset: 15624},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 620, col: 19, offset: 15627},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 620, col: 25, offset: 15633},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 624, col: 5, offset: 15709},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 624, col: 5, offset: 15709},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 624, col: 8, offset: 15712},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 631, col: 1, offset: 15789},
	expr: &actionExpr{
	pos: position{line: 631, col: 9, offset: 15797},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 631, col: 9, offset: 15797},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 631, col: 9, offset: 15797},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 631, col: 15, offset: 15803},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 631, col: 21, offset: 15809},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 631, col: 26, offset: 15814},
	expr: &seqExpr{
	pos: position{line: 631, col: 27, offset: 15815},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 631, col: 27, offset: 15815},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 631, col: 29, offset: 15817},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 631, col: 34, offset: 15822},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 631, col: 36, offset: 15824},
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
	pos: position{line: 635, col: 1, offset: 15888},
	expr: &actionExpr{
	pos: position{line: 635, col: 10, offset: 15897},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 635, col: 10, offset: 15897},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 635, col: 15, offset: 15902},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 635, col: 15, offset: 15902},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 635, col: 23, offset: 15910},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 635, col: 32, offset: 15919},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 635, col: 47, offset: 15934},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 639, col: 1, offset: 15977},
	expr: &actionExpr{
	pos: position{line: 639, col: 13, offset: 15989},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 639, col: 13, offset: 15989},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 639, col: 13, offset: 15989},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 639, col: 20, offset: 15996},
	expr: &choiceExpr{
	pos: position{line: 639, col: 21, offset: 15997},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 639, col: 21, offset: 15997},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 639, col: 31, offset: 16007},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 639, col: 39, offset: 16015},
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
	pos: position{line: 643, col: 1, offset: 16065},
	expr: &actionExpr{
	pos: position{line: 643, col: 17, offset: 16081},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 643, col: 17, offset: 16081},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 643, col: 17, offset: 16081},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 643, col: 21, offset: 16085},
	expr: &choiceExpr{
	pos: position{line: 643, col: 23, offset: 16087},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 643, col: 23, offset: 16087},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 643, col: 23, offset: 16087},
	expr: &ruleRefExpr{
	pos: position{line: 643, col: 24, offset: 16088},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 643, col: 36, offset: 16100,
},
	},
},
&seqExpr{
	pos: position{line: 643, col: 40, offset: 16104},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 643, col: 40, offset: 16104},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 643, col: 45, offset: 16109},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 643, col: 63, offset: 16127},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 650, col: 1, offset: 16308},
	expr: &charClassMatcher{
	pos: position{line: 650, col: 16, offset: 16323},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 652, col: 1, offset: 16341},
	expr: &ruleRefExpr{
	pos: position{line: 652, col: 19, offset: 16359},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 654, col: 1, offset: 16379},
	expr: &charClassMatcher{
	pos: position{line: 654, col: 21, offset: 16399},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 656, col: 1, offset: 16414},
	expr: &actionExpr{
	pos: position{line: 656, col: 9, offset: 16422},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 656, col: 9, offset: 16422},
	expr: &choiceExpr{
	pos: position{line: 656, col: 10, offset: 16423},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 656, col: 10, offset: 16423},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 656, col: 19, offset: 16432},
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
	pos: position{line: 660, col: 1, offset: 16482},
	expr: &choiceExpr{
	pos: position{line: 660, col: 11, offset: 16492},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 660, col: 11, offset: 16492},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 660, col: 19, offset: 16500},
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
	pos: position{line: 662, col: 1, offset: 16509},
	expr: &actionExpr{
	pos: position{line: 662, col: 11, offset: 16519},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 662, col: 11, offset: 16519},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 662, col: 11, offset: 16519},
	expr: &litMatcher{
	pos: position{line: 662, col: 11, offset: 16519},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 662, col: 16, offset: 16524},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 666, col: 1, offset: 16597},
	expr: &actionExpr{
	pos: position{line: 666, col: 19, offset: 16615},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 666, col: 19, offset: 16615},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 670, col: 1, offset: 16688},
	expr: &actionExpr{
	pos: position{line: 670, col: 10, offset: 16697},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 670, col: 10, offset: 16697},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 670, col: 10, offset: 16697},
	expr: &litMatcher{
	pos: position{line: 670, col: 10, offset: 16697},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 670, col: 15, offset: 16702},
	expr: &ruleRefExpr{
	pos: position{line: 670, col: 15, offset: 16702},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 670, col: 24, offset: 16711},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 670, col: 28, offset: 16715},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 674, col: 1, offset: 16782},
	expr: &choiceExpr{
	pos: position{line: 674, col: 12, offset: 16793},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 674, col: 12, offset: 16793},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 674, col: 18, offset: 16799},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 674, col: 18, offset: 16799},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 674, col: 38, offset: 16819},
	expr: &ruleRefExpr{
	pos: position{line: 674, col: 38, offset: 16819},
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
	pos: position{line: 676, col: 1, offset: 16836},
	expr: &charClassMatcher{
	pos: position{line: 676, col: 17, offset: 16852},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 678, col: 1, offset: 16861},
	expr: &charClassMatcher{
	pos: position{line: 678, col: 24, offset: 16884},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 680, col: 1, offset: 16893},
	expr: &zeroOrMoreExpr{
	pos: position{line: 680, col: 19, offset: 16911},
	expr: &charClassMatcher{
	pos: position{line: 680, col: 19, offset: 16911},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 682, col: 1, offset: 16925},
	expr: &charClassMatcher{
	pos: position{line: 682, col: 18, offset: 16942},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "IDENTIFIER",
	pos: position{line: 683, col: 1, offset: 16953},
	expr: &actionExpr{
	pos: position{line: 683, col: 15, offset: 16967},
	run: (*parser).callonIDENTIFIER1,
	expr: &seqExpr{
	pos: position{line: 683, col: 15, offset: 16967},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 683, col: 15, offset: 16967},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 683, col: 22, offset: 16974},
	expr: &choiceExpr{
	pos: position{line: 683, col: 23, offset: 16975},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 683, col: 23, offset: 16975},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 683, col: 33, offset: 16985},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 683, col: 41, offset: 16993},
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
	pos: position{line: 686, col: 1, offset: 17066},
	expr: &notExpr{
	pos: position{line: 686, col: 8, offset: 17073},
	expr: &anyMatcher{
	line: 686, col: 9, offset: 17074,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 688, col: 1, offset: 17079},
	expr: &litMatcher{
	pos: position{line: 688, col: 17, offset: 17095},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 694, col: 1, offset: 17182},
	expr: &actionExpr{
	pos: position{line: 694, col: 18, offset: 17199},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 694, col: 18, offset: 17199},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 694, col: 18, offset: 17199},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 694, col: 24, offset: 17205},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 694, col: 33, offset: 17214},
	name: "_",
},
&labeledExpr{
	pos: position{line: 694, col: 35, offset: 17216},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 694, col: 40, offset: 17221},
	expr: &seqExpr{
	pos: position{line: 694, col: 41, offset: 17222},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 694, col: 41, offset: 17222},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 694, col: 45, offset: 17226},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 694, col: 47, offset: 17228},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 694, col: 56, offset: 17237},
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
	pos: position{line: 697, col: 1, offset: 17290},
	expr: &actionExpr{
	pos: position{line: 697, col: 13, offset: 17302},
	run: (*parser).callonEqAliasG1,
	expr: &seqExpr{
	pos: position{line: 697, col: 13, offset: 17302},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 697, col: 13, offset: 17302},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 697, col: 19, offset: 17308},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 697, col: 28, offset: 17317},
	name: "_",
},
&litMatcher{
	pos: position{line: 697, col: 30, offset: 17319},
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
    return Into{
        StreamTo:streamTo,
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

