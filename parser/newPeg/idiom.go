
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
&ruleRefExpr{
	pos: position{line: 71, col: 31, offset: 1692},
	name: "FakeJobG",
},
	},
},
},
{
	name: "Transform",
	pos: position{line: 73, col: 1, offset: 1704},
	expr: &choiceExpr{
	pos: position{line: 73, col: 14, offset: 1717},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 73, col: 14, offset: 1717},
	name: "AggJobG",
},
&ruleRefExpr{
	pos: position{line: 73, col: 22, offset: 1725},
	name: "DoJobG",
},
	},
},
},
{
	name: "Sink",
	pos: position{line: 75, col: 1, offset: 1735},
	expr: &choiceExpr{
	pos: position{line: 75, col: 9, offset: 1743},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 75, col: 9, offset: 1743},
	name: "StdOutG",
},
&ruleRefExpr{
	pos: position{line: 75, col: 17, offset: 1751},
	name: "SinkInto",
},
	},
},
},
{
	name: "FileJobG",
	pos: position{line: 79, col: 1, offset: 1780},
	expr: &actionExpr{
	pos: position{line: 79, col: 13, offset: 1792},
	run: (*parser).callonFileJobG1,
	expr: &seqExpr{
	pos: position{line: 79, col: 13, offset: 1792},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 79, col: 13, offset: 1792},
	name: "_",
},
&litMatcher{
	pos: position{line: 79, col: 15, offset: 1794},
	val: "file(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 79, col: 23, offset: 1802},
	name: "_",
},
&labeledExpr{
	pos: position{line: 79, col: 25, offset: 1804},
	label: "filename",
	expr: &ruleRefExpr{
	pos: position{line: 79, col: 34, offset: 1813},
	name: "FileName",
},
},
&ruleRefExpr{
	pos: position{line: 79, col: 43, offset: 1822},
	name: "_",
},
&litMatcher{
	pos: position{line: 79, col: 45, offset: 1824},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FakeJobG",
	pos: position{line: 86, col: 1, offset: 2039},
	expr: &actionExpr{
	pos: position{line: 86, col: 13, offset: 2051},
	run: (*parser).callonFakeJobG1,
	expr: &seqExpr{
	pos: position{line: 86, col: 13, offset: 2051},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 86, col: 13, offset: 2051},
	name: "_",
},
&litMatcher{
	pos: position{line: 86, col: 15, offset: 2053},
	val: "fake(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 86, col: 23, offset: 2061},
	name: "_",
},
&labeledExpr{
	pos: position{line: 86, col: 25, offset: 2063},
	label: "numData",
	expr: &ruleRefExpr{
	pos: position{line: 86, col: 33, offset: 2071},
	name: "PositiveNumber",
},
},
&ruleRefExpr{
	pos: position{line: 86, col: 48, offset: 2086},
	name: "_",
},
&litMatcher{
	pos: position{line: 86, col: 50, offset: 2088},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchJobG",
	pos: position{line: 93, col: 1, offset: 2287},
	expr: &actionExpr{
	pos: position{line: 93, col: 15, offset: 2301},
	run: (*parser).callonBranchJobG1,
	expr: &seqExpr{
	pos: position{line: 93, col: 15, offset: 2301},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 93, col: 15, offset: 2301},
	val: "branch",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 93, col: 23, offset: 2309},
	name: "_",
},
&litMatcher{
	pos: position{line: 93, col: 24, offset: 2310},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 93, col: 28, offset: 2314},
	name: "_",
},
&labeledExpr{
	pos: position{line: 93, col: 30, offset: 2316},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 93, col: 33, offset: 2319},
	name: "IDENTIFIER",
},
},
&litMatcher{
	pos: position{line: 93, col: 44, offset: 2330},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggJobG",
	pos: position{line: 106, col: 1, offset: 2632},
	expr: &actionExpr{
	pos: position{line: 106, col: 12, offset: 2643},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 106, col: 12, offset: 2643},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 106, col: 12, offset: 2643},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 106, col: 18, offset: 2649},
	name: "_",
},
&labeledExpr{
	pos: position{line: 106, col: 20, offset: 2651},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 106, col: 27, offset: 2658},
	expr: &ruleRefExpr{
	pos: position{line: 106, col: 27, offset: 2658},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 106, col: 39, offset: 2670},
	name: "_",
},
&labeledExpr{
	pos: position{line: 106, col: 41, offset: 2672},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 106, col: 45, offset: 2676},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 106, col: 53, offset: 2684},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 106, col: 58, offset: 2689},
	expr: &seqExpr{
	pos: position{line: 106, col: 59, offset: 2690},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 106, col: 59, offset: 2690},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 106, col: 63, offset: 2694},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 106, col: 65, offset: 2696},
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
	pos: position{line: 121, col: 1, offset: 3071},
	expr: &choiceExpr{
	pos: position{line: 121, col: 12, offset: 3082},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 121, col: 12, offset: 3082},
	name: "CountAgg",
},
&ruleRefExpr{
	pos: position{line: 121, col: 23, offset: 3093},
	name: "MaxAgg",
},
&ruleRefExpr{
	pos: position{line: 121, col: 32, offset: 3102},
	name: "MinAgg",
},
&ruleRefExpr{
	pos: position{line: 121, col: 41, offset: 3111},
	name: "AvgAgg",
},
&ruleRefExpr{
	pos: position{line: 121, col: 50, offset: 3120},
	name: "SumAgg",
},
&ruleRefExpr{
	pos: position{line: 121, col: 59, offset: 3129},
	name: "ModeAgg",
},
&ruleRefExpr{
	pos: position{line: 122, col: 13, offset: 3150},
	name: "VarAgg",
},
&ruleRefExpr{
	pos: position{line: 122, col: 22, offset: 3159},
	name: "DistinctCountAgg",
},
&ruleRefExpr{
	pos: position{line: 122, col: 41, offset: 3178},
	name: "QuantileAgg",
},
&ruleRefExpr{
	pos: position{line: 122, col: 55, offset: 3192},
	name: "MedianAgg",
},
&ruleRefExpr{
	pos: position{line: 122, col: 67, offset: 3204},
	name: "QuartileAgg",
},
&ruleRefExpr{
	pos: position{line: 123, col: 13, offset: 3229},
	name: "CovAgg",
},
&ruleRefExpr{
	pos: position{line: 123, col: 22, offset: 3238},
	name: "CorrAgg",
},
&ruleRefExpr{
	pos: position{line: 123, col: 32, offset: 3248},
	name: "PValueAgg",
},
	},
},
},
{
	name: "CountAgg",
	pos: position{line: 126, col: 1, offset: 3292},
	expr: &actionExpr{
	pos: position{line: 126, col: 13, offset: 3304},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 126, col: 13, offset: 3304},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 126, col: 13, offset: 3304},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 126, col: 19, offset: 3310},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 126, col: 28, offset: 3319},
	name: "_",
},
&litMatcher{
	pos: position{line: 126, col: 30, offset: 3321},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 126, col: 39, offset: 3330},
	name: "_",
},
&labeledExpr{
	pos: position{line: 126, col: 41, offset: 3332},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 126, col: 46, offset: 3337},
	expr: &ruleRefExpr{
	pos: position{line: 126, col: 46, offset: 3337},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 126, col: 59, offset: 3350},
	name: "_",
},
&litMatcher{
	pos: position{line: 126, col: 61, offset: 3352},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MaxAgg",
	pos: position{line: 133, col: 1, offset: 3504},
	expr: &actionExpr{
	pos: position{line: 133, col: 11, offset: 3514},
	run: (*parser).callonMaxAgg1,
	expr: &seqExpr{
	pos: position{line: 133, col: 11, offset: 3514},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 133, col: 11, offset: 3514},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 133, col: 17, offset: 3520},
	expr: &ruleRefExpr{
	pos: position{line: 133, col: 17, offset: 3520},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 133, col: 27, offset: 3530},
	name: "_",
},
&litMatcher{
	pos: position{line: 133, col: 29, offset: 3532},
	val: "max(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 133, col: 36, offset: 3539},
	name: "_",
},
&labeledExpr{
	pos: position{line: 133, col: 38, offset: 3541},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 133, col: 44, offset: 3547},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 133, col: 53, offset: 3556},
	name: "_",
},
&labeledExpr{
	pos: position{line: 133, col: 55, offset: 3558},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 133, col: 60, offset: 3563},
	expr: &seqExpr{
	pos: position{line: 133, col: 61, offset: 3564},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 133, col: 61, offset: 3564},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 133, col: 65, offset: 3568},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 133, col: 67, offset: 3570},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 133, col: 81, offset: 3584},
	name: "_",
},
&litMatcher{
	pos: position{line: 133, col: 83, offset: 3586},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MinAgg",
	pos: position{line: 155, col: 1, offset: 4047},
	expr: &actionExpr{
	pos: position{line: 155, col: 11, offset: 4057},
	run: (*parser).callonMinAgg1,
	expr: &seqExpr{
	pos: position{line: 155, col: 11, offset: 4057},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 155, col: 11, offset: 4057},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 155, col: 17, offset: 4063},
	expr: &ruleRefExpr{
	pos: position{line: 155, col: 17, offset: 4063},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 155, col: 27, offset: 4073},
	name: "_",
},
&litMatcher{
	pos: position{line: 155, col: 29, offset: 4075},
	val: "min(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 155, col: 36, offset: 4082},
	name: "_",
},
&labeledExpr{
	pos: position{line: 155, col: 38, offset: 4084},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 155, col: 44, offset: 4090},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 155, col: 53, offset: 4099},
	name: "_",
},
&labeledExpr{
	pos: position{line: 155, col: 55, offset: 4101},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 155, col: 60, offset: 4106},
	expr: &seqExpr{
	pos: position{line: 155, col: 61, offset: 4107},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 155, col: 61, offset: 4107},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 155, col: 65, offset: 4111},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 155, col: 67, offset: 4113},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 155, col: 81, offset: 4127},
	name: "_",
},
&litMatcher{
	pos: position{line: 155, col: 83, offset: 4129},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AvgAgg",
	pos: position{line: 177, col: 1, offset: 4590},
	expr: &actionExpr{
	pos: position{line: 177, col: 11, offset: 4600},
	run: (*parser).callonAvgAgg1,
	expr: &seqExpr{
	pos: position{line: 177, col: 11, offset: 4600},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 177, col: 11, offset: 4600},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 177, col: 17, offset: 4606},
	expr: &ruleRefExpr{
	pos: position{line: 177, col: 17, offset: 4606},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 177, col: 27, offset: 4616},
	name: "_",
},
&litMatcher{
	pos: position{line: 177, col: 29, offset: 4618},
	val: "avg(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 177, col: 36, offset: 4625},
	name: "_",
},
&labeledExpr{
	pos: position{line: 177, col: 38, offset: 4627},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 177, col: 44, offset: 4633},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 177, col: 53, offset: 4642},
	name: "_",
},
&labeledExpr{
	pos: position{line: 177, col: 55, offset: 4644},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 177, col: 60, offset: 4649},
	expr: &seqExpr{
	pos: position{line: 177, col: 61, offset: 4650},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 177, col: 61, offset: 4650},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 177, col: 65, offset: 4654},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 177, col: 67, offset: 4656},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 177, col: 81, offset: 4670},
	name: "_",
},
&litMatcher{
	pos: position{line: 177, col: 83, offset: 4672},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SumAgg",
	pos: position{line: 199, col: 1, offset: 5133},
	expr: &actionExpr{
	pos: position{line: 199, col: 11, offset: 5143},
	run: (*parser).callonSumAgg1,
	expr: &seqExpr{
	pos: position{line: 199, col: 11, offset: 5143},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 199, col: 11, offset: 5143},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 199, col: 17, offset: 5149},
	expr: &ruleRefExpr{
	pos: position{line: 199, col: 17, offset: 5149},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 199, col: 27, offset: 5159},
	name: "_",
},
&litMatcher{
	pos: position{line: 199, col: 29, offset: 5161},
	val: "sum(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 199, col: 36, offset: 5168},
	name: "_",
},
&labeledExpr{
	pos: position{line: 199, col: 38, offset: 5170},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 199, col: 44, offset: 5176},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 199, col: 53, offset: 5185},
	name: "_",
},
&labeledExpr{
	pos: position{line: 199, col: 55, offset: 5187},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 199, col: 60, offset: 5192},
	expr: &seqExpr{
	pos: position{line: 199, col: 61, offset: 5193},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 199, col: 61, offset: 5193},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 199, col: 65, offset: 5197},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 199, col: 67, offset: 5199},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 199, col: 81, offset: 5213},
	name: "_",
},
&litMatcher{
	pos: position{line: 199, col: 83, offset: 5215},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarAgg",
	pos: position{line: 221, col: 1, offset: 5676},
	expr: &actionExpr{
	pos: position{line: 221, col: 11, offset: 5686},
	run: (*parser).callonVarAgg1,
	expr: &seqExpr{
	pos: position{line: 221, col: 11, offset: 5686},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 221, col: 11, offset: 5686},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 221, col: 17, offset: 5692},
	expr: &ruleRefExpr{
	pos: position{line: 221, col: 17, offset: 5692},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 221, col: 27, offset: 5702},
	name: "_",
},
&litMatcher{
	pos: position{line: 221, col: 29, offset: 5704},
	val: "var(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 221, col: 36, offset: 5711},
	name: "_",
},
&labeledExpr{
	pos: position{line: 221, col: 38, offset: 5713},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 221, col: 44, offset: 5719},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 221, col: 53, offset: 5728},
	name: "_",
},
&labeledExpr{
	pos: position{line: 221, col: 55, offset: 5730},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 221, col: 60, offset: 5735},
	expr: &seqExpr{
	pos: position{line: 221, col: 61, offset: 5736},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 221, col: 61, offset: 5736},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 221, col: 65, offset: 5740},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 221, col: 67, offset: 5742},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 221, col: 81, offset: 5756},
	name: "_",
},
&litMatcher{
	pos: position{line: 221, col: 83, offset: 5758},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DistinctCountAgg",
	pos: position{line: 243, col: 1, offset: 6224},
	expr: &actionExpr{
	pos: position{line: 243, col: 21, offset: 6244},
	run: (*parser).callonDistinctCountAgg1,
	expr: &seqExpr{
	pos: position{line: 243, col: 21, offset: 6244},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 243, col: 21, offset: 6244},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 243, col: 27, offset: 6250},
	expr: &ruleRefExpr{
	pos: position{line: 243, col: 27, offset: 6250},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 243, col: 37, offset: 6260},
	name: "_",
},
&litMatcher{
	pos: position{line: 243, col: 39, offset: 6262},
	val: "dcount(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 243, col: 49, offset: 6272},
	name: "_",
},
&labeledExpr{
	pos: position{line: 243, col: 51, offset: 6274},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 243, col: 57, offset: 6280},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 243, col: 66, offset: 6289},
	name: "_",
},
&labeledExpr{
	pos: position{line: 243, col: 68, offset: 6291},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 243, col: 73, offset: 6296},
	expr: &seqExpr{
	pos: position{line: 243, col: 74, offset: 6297},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 243, col: 74, offset: 6297},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 243, col: 78, offset: 6301},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 243, col: 80, offset: 6303},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 243, col: 94, offset: 6317},
	name: "_",
},
&litMatcher{
	pos: position{line: 243, col: 96, offset: 6319},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "ModeAgg",
	pos: position{line: 265, col: 1, offset: 6790},
	expr: &actionExpr{
	pos: position{line: 265, col: 12, offset: 6801},
	run: (*parser).callonModeAgg1,
	expr: &seqExpr{
	pos: position{line: 265, col: 12, offset: 6801},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 265, col: 12, offset: 6801},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 265, col: 18, offset: 6807},
	expr: &ruleRefExpr{
	pos: position{line: 265, col: 18, offset: 6807},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 265, col: 28, offset: 6817},
	name: "_",
},
&litMatcher{
	pos: position{line: 265, col: 30, offset: 6819},
	val: "mode(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 265, col: 38, offset: 6827},
	name: "_",
},
&labeledExpr{
	pos: position{line: 265, col: 40, offset: 6829},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 265, col: 46, offset: 6835},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 265, col: 55, offset: 6844},
	name: "_",
},
&labeledExpr{
	pos: position{line: 265, col: 57, offset: 6846},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 265, col: 62, offset: 6851},
	expr: &seqExpr{
	pos: position{line: 265, col: 63, offset: 6852},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 265, col: 63, offset: 6852},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 265, col: 67, offset: 6856},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 265, col: 69, offset: 6858},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 265, col: 83, offset: 6872},
	name: "_",
},
&litMatcher{
	pos: position{line: 265, col: 85, offset: 6874},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuantileAgg",
	pos: position{line: 287, col: 1, offset: 7336},
	expr: &actionExpr{
	pos: position{line: 287, col: 16, offset: 7351},
	run: (*parser).callonQuantileAgg1,
	expr: &seqExpr{
	pos: position{line: 287, col: 16, offset: 7351},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 287, col: 16, offset: 7351},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 287, col: 22, offset: 7357},
	expr: &ruleRefExpr{
	pos: position{line: 287, col: 22, offset: 7357},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 287, col: 32, offset: 7367},
	name: "_",
},
&litMatcher{
	pos: position{line: 287, col: 34, offset: 7369},
	val: "quantile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 287, col: 46, offset: 7381},
	name: "_",
},
&labeledExpr{
	pos: position{line: 287, col: 48, offset: 7383},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 287, col: 54, offset: 7389},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 287, col: 63, offset: 7398},
	name: "_",
},
&labeledExpr{
	pos: position{line: 287, col: 66, offset: 7401},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 287, col: 73, offset: 7408},
	expr: &seqExpr{
	pos: position{line: 287, col: 74, offset: 7409},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 287, col: 74, offset: 7409},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 287, col: 78, offset: 7413},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 287, col: 80, offset: 7415},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 287, col: 91, offset: 7426},
	name: "_",
},
&litMatcher{
	pos: position{line: 287, col: 93, offset: 7428},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 287, col: 97, offset: 7432},
	name: "_",
},
&labeledExpr{
	pos: position{line: 287, col: 99, offset: 7434},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 287, col: 103, offset: 7438},
	name: "Float",
},
},
&ruleRefExpr{
	pos: position{line: 287, col: 109, offset: 7444},
	name: "_",
},
&labeledExpr{
	pos: position{line: 287, col: 111, offset: 7446},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 287, col: 116, offset: 7451},
	expr: &seqExpr{
	pos: position{line: 287, col: 117, offset: 7452},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 287, col: 117, offset: 7452},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 287, col: 121, offset: 7456},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 287, col: 123, offset: 7458},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 287, col: 137, offset: 7472},
	name: "_",
},
&litMatcher{
	pos: position{line: 287, col: 139, offset: 7474},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MedianAgg",
	pos: position{line: 317, col: 1, offset: 8119},
	expr: &actionExpr{
	pos: position{line: 317, col: 14, offset: 8132},
	run: (*parser).callonMedianAgg1,
	expr: &seqExpr{
	pos: position{line: 317, col: 14, offset: 8132},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 317, col: 14, offset: 8132},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 317, col: 20, offset: 8138},
	expr: &ruleRefExpr{
	pos: position{line: 317, col: 20, offset: 8138},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 317, col: 30, offset: 8148},
	name: "_",
},
&litMatcher{
	pos: position{line: 317, col: 32, offset: 8150},
	val: "median(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 317, col: 42, offset: 8160},
	name: "_",
},
&labeledExpr{
	pos: position{line: 317, col: 44, offset: 8162},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 317, col: 50, offset: 8168},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 317, col: 59, offset: 8177},
	name: "_",
},
&labeledExpr{
	pos: position{line: 317, col: 62, offset: 8180},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 317, col: 69, offset: 8187},
	expr: &seqExpr{
	pos: position{line: 317, col: 70, offset: 8188},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 317, col: 70, offset: 8188},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 317, col: 74, offset: 8192},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 317, col: 76, offset: 8194},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 317, col: 87, offset: 8205},
	name: "_",
},
&labeledExpr{
	pos: position{line: 317, col: 89, offset: 8207},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 317, col: 94, offset: 8212},
	expr: &seqExpr{
	pos: position{line: 317, col: 95, offset: 8213},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 317, col: 95, offset: 8213},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 317, col: 99, offset: 8217},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 317, col: 101, offset: 8219},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 317, col: 115, offset: 8233},
	name: "_",
},
&litMatcher{
	pos: position{line: 317, col: 117, offset: 8235},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuartileAgg",
	pos: position{line: 349, col: 1, offset: 8912},
	expr: &actionExpr{
	pos: position{line: 349, col: 16, offset: 8927},
	run: (*parser).callonQuartileAgg1,
	expr: &seqExpr{
	pos: position{line: 349, col: 16, offset: 8927},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 349, col: 16, offset: 8927},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 349, col: 22, offset: 8933},
	expr: &ruleRefExpr{
	pos: position{line: 349, col: 22, offset: 8933},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 349, col: 32, offset: 8943},
	name: "_",
},
&litMatcher{
	pos: position{line: 349, col: 34, offset: 8945},
	val: "quartile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 349, col: 46, offset: 8957},
	name: "_",
},
&labeledExpr{
	pos: position{line: 349, col: 48, offset: 8959},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 349, col: 54, offset: 8965},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 349, col: 63, offset: 8974},
	name: "_",
},
&labeledExpr{
	pos: position{line: 349, col: 66, offset: 8977},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 349, col: 73, offset: 8984},
	expr: &seqExpr{
	pos: position{line: 349, col: 74, offset: 8985},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 349, col: 74, offset: 8985},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 349, col: 78, offset: 8989},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 349, col: 80, offset: 8991},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 349, col: 91, offset: 9002},
	name: "_",
},
&litMatcher{
	pos: position{line: 349, col: 93, offset: 9004},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 349, col: 97, offset: 9008},
	name: "_",
},
&labeledExpr{
	pos: position{line: 349, col: 99, offset: 9010},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 349, col: 103, offset: 9014},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 349, col: 110, offset: 9021},
	name: "_",
},
&labeledExpr{
	pos: position{line: 349, col: 112, offset: 9023},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 349, col: 117, offset: 9028},
	expr: &seqExpr{
	pos: position{line: 349, col: 118, offset: 9029},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 349, col: 118, offset: 9029},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 349, col: 122, offset: 9033},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 349, col: 124, offset: 9035},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 349, col: 138, offset: 9049},
	name: "_",
},
&litMatcher{
	pos: position{line: 349, col: 140, offset: 9051},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CovAgg",
	pos: position{line: 395, col: 1, offset: 10078},
	expr: &actionExpr{
	pos: position{line: 395, col: 11, offset: 10088},
	run: (*parser).callonCovAgg1,
	expr: &seqExpr{
	pos: position{line: 395, col: 11, offset: 10088},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 395, col: 11, offset: 10088},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 395, col: 17, offset: 10094},
	expr: &ruleRefExpr{
	pos: position{line: 395, col: 17, offset: 10094},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 395, col: 27, offset: 10104},
	name: "_",
},
&litMatcher{
	pos: position{line: 395, col: 29, offset: 10106},
	val: "cov(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 395, col: 36, offset: 10113},
	name: "_",
},
&labeledExpr{
	pos: position{line: 395, col: 38, offset: 10115},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 395, col: 45, offset: 10122},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 395, col: 54, offset: 10131},
	name: "_",
},
&litMatcher{
	pos: position{line: 395, col: 56, offset: 10133},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 395, col: 60, offset: 10137},
	name: "_",
},
&labeledExpr{
	pos: position{line: 395, col: 62, offset: 10139},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 395, col: 69, offset: 10146},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 395, col: 78, offset: 10155},
	name: "_",
},
&labeledExpr{
	pos: position{line: 395, col: 80, offset: 10157},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 395, col: 85, offset: 10162},
	expr: &seqExpr{
	pos: position{line: 395, col: 86, offset: 10163},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 395, col: 86, offset: 10163},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 395, col: 90, offset: 10167},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 395, col: 92, offset: 10169},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 395, col: 106, offset: 10183},
	name: "_",
},
&litMatcher{
	pos: position{line: 395, col: 108, offset: 10185},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CorrAgg",
	pos: position{line: 418, col: 1, offset: 10710},
	expr: &actionExpr{
	pos: position{line: 418, col: 12, offset: 10721},
	run: (*parser).callonCorrAgg1,
	expr: &seqExpr{
	pos: position{line: 418, col: 12, offset: 10721},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 418, col: 12, offset: 10721},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 418, col: 18, offset: 10727},
	expr: &ruleRefExpr{
	pos: position{line: 418, col: 18, offset: 10727},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 418, col: 28, offset: 10737},
	name: "_",
},
&litMatcher{
	pos: position{line: 418, col: 30, offset: 10739},
	val: "correlation(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 418, col: 45, offset: 10754},
	name: "_",
},
&labeledExpr{
	pos: position{line: 418, col: 47, offset: 10756},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 418, col: 54, offset: 10763},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 418, col: 63, offset: 10772},
	name: "_",
},
&litMatcher{
	pos: position{line: 418, col: 65, offset: 10774},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 418, col: 69, offset: 10778},
	name: "_",
},
&labeledExpr{
	pos: position{line: 418, col: 71, offset: 10780},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 418, col: 78, offset: 10787},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 418, col: 87, offset: 10796},
	name: "_",
},
&labeledExpr{
	pos: position{line: 418, col: 89, offset: 10798},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 418, col: 94, offset: 10803},
	expr: &seqExpr{
	pos: position{line: 418, col: 95, offset: 10804},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 418, col: 95, offset: 10804},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 418, col: 99, offset: 10808},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 418, col: 101, offset: 10810},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 418, col: 115, offset: 10824},
	name: "_",
},
&litMatcher{
	pos: position{line: 418, col: 117, offset: 10826},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PValueAgg",
	pos: position{line: 441, col: 1, offset: 11352},
	expr: &actionExpr{
	pos: position{line: 441, col: 14, offset: 11365},
	run: (*parser).callonPValueAgg1,
	expr: &seqExpr{
	pos: position{line: 441, col: 14, offset: 11365},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 441, col: 14, offset: 11365},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 441, col: 20, offset: 11371},
	expr: &ruleRefExpr{
	pos: position{line: 441, col: 20, offset: 11371},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 441, col: 30, offset: 11381},
	name: "_",
},
&litMatcher{
	pos: position{line: 441, col: 32, offset: 11383},
	val: "pvalue(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 441, col: 42, offset: 11393},
	name: "_",
},
&labeledExpr{
	pos: position{line: 441, col: 44, offset: 11395},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 441, col: 51, offset: 11402},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 441, col: 60, offset: 11411},
	name: "_",
},
&litMatcher{
	pos: position{line: 441, col: 62, offset: 11413},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 441, col: 66, offset: 11417},
	name: "_",
},
&labeledExpr{
	pos: position{line: 441, col: 68, offset: 11419},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 441, col: 75, offset: 11426},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 441, col: 84, offset: 11435},
	name: "_",
},
&labeledExpr{
	pos: position{line: 441, col: 86, offset: 11437},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 441, col: 91, offset: 11442},
	expr: &seqExpr{
	pos: position{line: 441, col: 92, offset: 11443},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 441, col: 92, offset: 11443},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 441, col: 96, offset: 11447},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 441, col: 98, offset: 11449},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 441, col: 112, offset: 11463},
	name: "_",
},
&litMatcher{
	pos: position{line: 441, col: 114, offset: 11465},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggGroupBy",
	pos: position{line: 466, col: 1, offset: 11990},
	expr: &actionExpr{
	pos: position{line: 466, col: 15, offset: 12004},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 466, col: 15, offset: 12004},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 466, col: 15, offset: 12004},
	name: "_",
},
&litMatcher{
	pos: position{line: 466, col: 17, offset: 12006},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 466, col: 22, offset: 12011},
	name: "_",
},
&labeledExpr{
	pos: position{line: 466, col: 24, offset: 12013},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 466, col: 30, offset: 12019},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 466, col: 39, offset: 12028},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 466, col: 44, offset: 12033},
	expr: &seqExpr{
	pos: position{line: 466, col: 45, offset: 12034},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 466, col: 45, offset: 12034},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 466, col: 49, offset: 12038},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 466, col: 51, offset: 12040},
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
	pos: position{line: 481, col: 1, offset: 12391},
	expr: &actionExpr{
	pos: position{line: 481, col: 11, offset: 12401},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 481, col: 11, offset: 12401},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 481, col: 16, offset: 12406},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 481, col: 16, offset: 12406},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 481, col: 27, offset: 12417},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 481, col: 38, offset: 12428},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 481, col: 49, offset: 12439},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 481, col: 58, offset: 12448},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 481, col: 67, offset: 12457},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 487, col: 1, offset: 12576},
	expr: &actionExpr{
	pos: position{line: 487, col: 13, offset: 12588},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 487, col: 13, offset: 12588},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 487, col: 13, offset: 12588},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 487, col: 23, offset: 12598},
	name: "_",
},
&labeledExpr{
	pos: position{line: 487, col: 25, offset: 12600},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 487, col: 28, offset: 12603},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 487, col: 42, offset: 12617},
	name: "_",
},
&litMatcher{
	pos: position{line: 487, col: 44, offset: 12619},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 494, col: 1, offset: 12814},
	expr: &actionExpr{
	pos: position{line: 494, col: 18, offset: 12831},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 494, col: 18, offset: 12831},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 494, col: 18, offset: 12831},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 494, col: 23, offset: 12836},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 494, col: 28, offset: 12841},
	name: "_",
},
&labeledExpr{
	pos: position{line: 494, col: 30, offset: 12843},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 494, col: 33, offset: 12846},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 494, col: 45, offset: 12858},
	name: "_",
},
&labeledExpr{
	pos: position{line: 494, col: 47, offset: 12860},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 494, col: 54, offset: 12867},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 498, col: 1, offset: 12918},
	expr: &ruleRefExpr{
	pos: position{line: 498, col: 16, offset: 12933},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 501, col: 1, offset: 12967},
	expr: &actionExpr{
	pos: position{line: 501, col: 13, offset: 12979},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 501, col: 13, offset: 12979},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 501, col: 13, offset: 12979},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 501, col: 23, offset: 12989},
	name: "_",
},
&labeledExpr{
	pos: position{line: 501, col: 25, offset: 12991},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 501, col: 31, offset: 12997},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 501, col: 40, offset: 13006},
	name: "_",
},
&labeledExpr{
	pos: position{line: 501, col: 42, offset: 13008},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 501, col: 47, offset: 13013},
	expr: &seqExpr{
	pos: position{line: 501, col: 48, offset: 13014},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 501, col: 48, offset: 13014},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 501, col: 52, offset: 13018},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 501, col: 54, offset: 13020},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 501, col: 65, offset: 13031},
	name: "_",
},
&litMatcher{
	pos: position{line: 501, col: 67, offset: 13033},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 516, col: 1, offset: 13370},
	expr: &actionExpr{
	pos: position{line: 516, col: 11, offset: 13380},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 516, col: 11, offset: 13380},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 516, col: 11, offset: 13380},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 516, col: 18, offset: 13387},
	name: "_",
},
&labeledExpr{
	pos: position{line: 516, col: 20, offset: 13389},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 516, col: 25, offset: 13394},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 516, col: 34, offset: 13403},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 516, col: 38, offset: 13407},
	name: "_",
},
&labeledExpr{
	pos: position{line: 516, col: 40, offset: 13409},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 516, col: 44, offset: 13413},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 516, col: 51, offset: 13420},
	name: "_",
},
&litMatcher{
	pos: position{line: 516, col: 53, offset: 13422},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 528, col: 1, offset: 13644},
	expr: &actionExpr{
	pos: position{line: 528, col: 11, offset: 13654},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 528, col: 11, offset: 13654},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 528, col: 11, offset: 13654},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 528, col: 20, offset: 13663},
	name: "_",
},
&litMatcher{
	pos: position{line: 528, col: 22, offset: 13665},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 528, col: 27, offset: 13670},
	name: "_",
},
&labeledExpr{
	pos: position{line: 528, col: 29, offset: 13672},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 528, col: 35, offset: 13678},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 534, col: 1, offset: 13784},
	expr: &actionExpr{
	pos: position{line: 534, col: 12, offset: 13795},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 534, col: 12, offset: 13795},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 534, col: 12, offset: 13795},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 534, col: 21, offset: 13804},
	name: "_",
},
&labeledExpr{
	pos: position{line: 534, col: 23, offset: 13806},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 534, col: 27, offset: 13810},
	expr: &ruleRefExpr{
	pos: position{line: 534, col: 27, offset: 13810},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 534, col: 35, offset: 13818},
	name: "_",
},
&litMatcher{
	pos: position{line: 534, col: 37, offset: 13820},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 544, col: 1, offset: 13965},
	expr: &actionExpr{
	pos: position{line: 544, col: 13, offset: 13977},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 544, col: 13, offset: 13977},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 544, col: 13, offset: 13977},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 544, col: 23, offset: 13987},
	name: "_",
},
&labeledExpr{
	pos: position{line: 544, col: 25, offset: 13989},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 544, col: 30, offset: 13994},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 544, col: 42, offset: 14006},
	name: "_",
},
&litMatcher{
	pos: position{line: 544, col: 44, offset: 14008},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 554, col: 1, offset: 14296},
	expr: &actionExpr{
	pos: position{line: 554, col: 16, offset: 14311},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 554, col: 16, offset: 14311},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 554, col: 16, offset: 14311},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 554, col: 22, offset: 14317},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 554, col: 35, offset: 14330},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 554, col: 40, offset: 14335},
	expr: &seqExpr{
	pos: position{line: 554, col: 41, offset: 14336},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 554, col: 41, offset: 14336},
	name: "_",
},
&labeledExpr{
	pos: position{line: 554, col: 43, offset: 14338},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 554, col: 46, offset: 14341},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 554, col: 58, offset: 14353},
	name: "_",
},
&labeledExpr{
	pos: position{line: 554, col: 60, offset: 14355},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 554, col: 63, offset: 14358},
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
	pos: position{line: 559, col: 1, offset: 14516},
	expr: &choiceExpr{
	pos: position{line: 559, col: 17, offset: 14532},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 559, col: 17, offset: 14532},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 559, col: 17, offset: 14532},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 559, col: 17, offset: 14532},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 559, col: 21, offset: 14536},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 559, col: 24, offset: 14539},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 559, col: 36, offset: 14551},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 561, col: 5, offset: 14582},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 561, col: 5, offset: 14582},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 561, col: 8, offset: 14585},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 565, col: 1, offset: 14627},
	expr: &choiceExpr{
	pos: position{line: 565, col: 16, offset: 14642},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 565, col: 16, offset: 14642},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 565, col: 24, offset: 14650},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 565, col: 24, offset: 14650},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 572, col: 1, offset: 14829},
	expr: &actionExpr{
	pos: position{line: 572, col: 18, offset: 14846},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 572, col: 18, offset: 14846},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 572, col: 18, offset: 14846},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 572, col: 23, offset: 14851},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 572, col: 28, offset: 14856},
	name: "_",
},
&labeledExpr{
	pos: position{line: 572, col: 30, offset: 14858},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 572, col: 33, offset: 14861},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 572, col: 45, offset: 14873},
	name: "_",
},
&labeledExpr{
	pos: position{line: 572, col: 47, offset: 14875},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 572, col: 53, offset: 14881},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 572, col: 59, offset: 14887},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 576, col: 1, offset: 14961},
	expr: &choiceExpr{
	pos: position{line: 576, col: 16, offset: 14976},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 576, col: 16, offset: 14976},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 576, col: 23, offset: 14983},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 576, col: 30, offset: 14990},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 576, col: 36, offset: 14996},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 576, col: 42, offset: 15002},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 576, col: 49, offset: 15009},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 576, col: 49, offset: 15009},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 585, col: 1, offset: 15119},
	expr: &actionExpr{
	pos: position{line: 585, col: 12, offset: 15130},
	run: (*parser).callonStdOutG1,
	expr: &seqExpr{
	pos: position{line: 585, col: 12, offset: 15130},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 585, col: 12, offset: 15130},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 585, col: 22, offset: 15140},
	name: "_",
},
&labeledExpr{
	pos: position{line: 585, col: 24, offset: 15142},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 585, col: 31, offset: 15149},
	expr: &ruleRefExpr{
	pos: position{line: 585, col: 32, offset: 15150},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 585, col: 48, offset: 15166},
	name: "_",
},
&litMatcher{
	pos: position{line: 585, col: 50, offset: 15168},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SinkInto",
	pos: position{line: 599, col: 1, offset: 15447},
	expr: &actionExpr{
	pos: position{line: 599, col: 12, offset: 15458},
	run: (*parser).callonSinkInto1,
	expr: &seqExpr{
	pos: position{line: 599, col: 12, offset: 15458},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 599, col: 12, offset: 15458},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 599, col: 18, offset: 15464},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 599, col: 27, offset: 15473},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 599, col: 33, offset: 15479},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 599, col: 33, offset: 15479},
	name: "_",
},
&litMatcher{
	pos: position{line: 599, col: 35, offset: 15481},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 599, col: 39, offset: 15485},
	name: "_",
},
&litMatcher{
	pos: position{line: 599, col: 41, offset: 15487},
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
	pos: position{line: 610, col: 1, offset: 15664},
	expr: &actionExpr{
	pos: position{line: 610, col: 13, offset: 15676},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 610, col: 13, offset: 15676},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 610, col: 13, offset: 15676},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 610, col: 22, offset: 15685},
	expr: &seqExpr{
	pos: position{line: 610, col: 23, offset: 15686},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 610, col: 23, offset: 15686},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 610, col: 27, offset: 15690},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 610, col: 31, offset: 15694},
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
	pos: position{line: 616, col: 1, offset: 15795},
	expr: &actionExpr{
	pos: position{line: 616, col: 10, offset: 15804},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 616, col: 10, offset: 15804},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 616, col: 10, offset: 15804},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 616, col: 16, offset: 15810},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 616, col: 23, offset: 15817},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 616, col: 28, offset: 15822},
	expr: &seqExpr{
	pos: position{line: 616, col: 29, offset: 15823},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 616, col: 29, offset: 15823},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 616, col: 31, offset: 15825},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 616, col: 36, offset: 15830},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 616, col: 38, offset: 15832},
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
	pos: position{line: 620, col: 1, offset: 15897},
	expr: &actionExpr{
	pos: position{line: 620, col: 11, offset: 15907},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 620, col: 11, offset: 15907},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 620, col: 16, offset: 15912},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 620, col: 16, offset: 15912},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 620, col: 24, offset: 15920},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 620, col: 33, offset: 15929},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 620, col: 48, offset: 15944},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 620, col: 59, offset: 15955},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 624, col: 1, offset: 15997},
	expr: &choiceExpr{
	pos: position{line: 624, col: 9, offset: 16005},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 624, col: 9, offset: 16005},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 624, col: 15, offset: 16011},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 624, col: 21, offset: 16017},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 624, col: 27, offset: 16023},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 624, col: 33, offset: 16029},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 624, col: 33, offset: 16029},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 628, col: 1, offset: 16077},
	expr: &choiceExpr{
	pos: position{line: 628, col: 12, offset: 16088},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 628, col: 12, offset: 16088},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 628, col: 12, offset: 16088},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 628, col: 12, offset: 16088},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 628, col: 16, offset: 16092},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 628, col: 19, offset: 16095},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 628, col: 25, offset: 16101},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 632, col: 5, offset: 16177},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 632, col: 5, offset: 16177},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 632, col: 8, offset: 16180},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 639, col: 1, offset: 16257},
	expr: &actionExpr{
	pos: position{line: 639, col: 9, offset: 16265},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 639, col: 9, offset: 16265},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 639, col: 9, offset: 16265},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 639, col: 15, offset: 16271},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 639, col: 21, offset: 16277},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 639, col: 26, offset: 16282},
	expr: &seqExpr{
	pos: position{line: 639, col: 27, offset: 16283},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 639, col: 27, offset: 16283},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 639, col: 29, offset: 16285},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 639, col: 34, offset: 16290},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 639, col: 36, offset: 16292},
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
	pos: position{line: 643, col: 1, offset: 16356},
	expr: &actionExpr{
	pos: position{line: 643, col: 10, offset: 16365},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 643, col: 10, offset: 16365},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 643, col: 15, offset: 16370},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 643, col: 15, offset: 16370},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 643, col: 23, offset: 16378},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 643, col: 32, offset: 16387},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 643, col: 47, offset: 16402},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 647, col: 1, offset: 16445},
	expr: &actionExpr{
	pos: position{line: 647, col: 13, offset: 16457},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 647, col: 13, offset: 16457},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 647, col: 13, offset: 16457},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 647, col: 20, offset: 16464},
	expr: &choiceExpr{
	pos: position{line: 647, col: 21, offset: 16465},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 647, col: 21, offset: 16465},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 647, col: 31, offset: 16475},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 647, col: 39, offset: 16483},
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
	pos: position{line: 651, col: 1, offset: 16533},
	expr: &actionExpr{
	pos: position{line: 651, col: 17, offset: 16549},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 651, col: 17, offset: 16549},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 651, col: 17, offset: 16549},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 651, col: 21, offset: 16553},
	expr: &choiceExpr{
	pos: position{line: 651, col: 23, offset: 16555},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 651, col: 23, offset: 16555},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 651, col: 23, offset: 16555},
	expr: &ruleRefExpr{
	pos: position{line: 651, col: 24, offset: 16556},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 651, col: 36, offset: 16568,
},
	},
},
&seqExpr{
	pos: position{line: 651, col: 40, offset: 16572},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 651, col: 40, offset: 16572},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 651, col: 45, offset: 16577},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 651, col: 63, offset: 16595},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 658, col: 1, offset: 16776},
	expr: &charClassMatcher{
	pos: position{line: 658, col: 16, offset: 16791},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 660, col: 1, offset: 16809},
	expr: &ruleRefExpr{
	pos: position{line: 660, col: 19, offset: 16827},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 662, col: 1, offset: 16847},
	expr: &charClassMatcher{
	pos: position{line: 662, col: 21, offset: 16867},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 664, col: 1, offset: 16882},
	expr: &actionExpr{
	pos: position{line: 664, col: 9, offset: 16890},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 664, col: 9, offset: 16890},
	expr: &choiceExpr{
	pos: position{line: 664, col: 10, offset: 16891},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 664, col: 10, offset: 16891},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 664, col: 19, offset: 16900},
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
	pos: position{line: 668, col: 1, offset: 16950},
	expr: &choiceExpr{
	pos: position{line: 668, col: 11, offset: 16960},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 668, col: 11, offset: 16960},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 668, col: 19, offset: 16968},
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
	pos: position{line: 670, col: 1, offset: 16977},
	expr: &actionExpr{
	pos: position{line: 670, col: 11, offset: 16987},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 670, col: 11, offset: 16987},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 670, col: 11, offset: 16987},
	expr: &litMatcher{
	pos: position{line: 670, col: 11, offset: 16987},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 670, col: 16, offset: 16992},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 674, col: 1, offset: 17065},
	expr: &actionExpr{
	pos: position{line: 674, col: 19, offset: 17083},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 674, col: 19, offset: 17083},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 678, col: 1, offset: 17156},
	expr: &actionExpr{
	pos: position{line: 678, col: 10, offset: 17165},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 678, col: 10, offset: 17165},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 678, col: 10, offset: 17165},
	expr: &litMatcher{
	pos: position{line: 678, col: 10, offset: 17165},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 678, col: 15, offset: 17170},
	expr: &ruleRefExpr{
	pos: position{line: 678, col: 15, offset: 17170},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 678, col: 24, offset: 17179},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 678, col: 28, offset: 17183},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 682, col: 1, offset: 17250},
	expr: &choiceExpr{
	pos: position{line: 682, col: 12, offset: 17261},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 682, col: 12, offset: 17261},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 682, col: 18, offset: 17267},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 682, col: 18, offset: 17267},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 682, col: 38, offset: 17287},
	expr: &ruleRefExpr{
	pos: position{line: 682, col: 38, offset: 17287},
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
	pos: position{line: 684, col: 1, offset: 17304},
	expr: &charClassMatcher{
	pos: position{line: 684, col: 17, offset: 17320},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 686, col: 1, offset: 17329},
	expr: &charClassMatcher{
	pos: position{line: 686, col: 24, offset: 17352},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 688, col: 1, offset: 17361},
	expr: &zeroOrMoreExpr{
	pos: position{line: 688, col: 19, offset: 17379},
	expr: &charClassMatcher{
	pos: position{line: 688, col: 19, offset: 17379},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 690, col: 1, offset: 17393},
	expr: &charClassMatcher{
	pos: position{line: 690, col: 18, offset: 17410},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "IDENTIFIER",
	pos: position{line: 691, col: 1, offset: 17421},
	expr: &actionExpr{
	pos: position{line: 691, col: 15, offset: 17435},
	run: (*parser).callonIDENTIFIER1,
	expr: &seqExpr{
	pos: position{line: 691, col: 15, offset: 17435},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 691, col: 15, offset: 17435},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 691, col: 22, offset: 17442},
	expr: &choiceExpr{
	pos: position{line: 691, col: 23, offset: 17443},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 691, col: 23, offset: 17443},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 691, col: 33, offset: 17453},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 691, col: 41, offset: 17461},
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
	pos: position{line: 694, col: 1, offset: 17534},
	expr: &notExpr{
	pos: position{line: 694, col: 8, offset: 17541},
	expr: &anyMatcher{
	line: 694, col: 9, offset: 17542,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 696, col: 1, offset: 17547},
	expr: &litMatcher{
	pos: position{line: 696, col: 17, offset: 17563},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 702, col: 1, offset: 17650},
	expr: &actionExpr{
	pos: position{line: 702, col: 18, offset: 17667},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 702, col: 18, offset: 17667},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 702, col: 18, offset: 17667},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 702, col: 24, offset: 17673},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 702, col: 33, offset: 17682},
	name: "_",
},
&labeledExpr{
	pos: position{line: 702, col: 35, offset: 17684},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 702, col: 40, offset: 17689},
	expr: &seqExpr{
	pos: position{line: 702, col: 41, offset: 17690},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 702, col: 41, offset: 17690},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 702, col: 45, offset: 17694},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 702, col: 47, offset: 17696},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 702, col: 56, offset: 17705},
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
	pos: position{line: 705, col: 1, offset: 17758},
	expr: &actionExpr{
	pos: position{line: 705, col: 13, offset: 17770},
	run: (*parser).callonEqAliasG1,
	expr: &seqExpr{
	pos: position{line: 705, col: 13, offset: 17770},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 705, col: 13, offset: 17770},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 705, col: 19, offset: 17776},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 705, col: 28, offset: 17785},
	name: "_",
},
&litMatcher{
	pos: position{line: 705, col: 30, offset: 17787},
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

