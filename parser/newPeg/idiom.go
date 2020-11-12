
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
&ruleRefExpr{
	pos: position{line: 75, col: 26, offset: 1760},
	name: "PlotG",
},
	},
},
},
{
	name: "FileJobG",
	pos: position{line: 79, col: 1, offset: 1786},
	expr: &actionExpr{
	pos: position{line: 79, col: 13, offset: 1798},
	run: (*parser).callonFileJobG1,
	expr: &seqExpr{
	pos: position{line: 79, col: 13, offset: 1798},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 79, col: 13, offset: 1798},
	name: "_",
},
&litMatcher{
	pos: position{line: 79, col: 15, offset: 1800},
	val: "file(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 79, col: 23, offset: 1808},
	name: "_",
},
&labeledExpr{
	pos: position{line: 79, col: 25, offset: 1810},
	label: "filename",
	expr: &ruleRefExpr{
	pos: position{line: 79, col: 34, offset: 1819},
	name: "FileName",
},
},
&ruleRefExpr{
	pos: position{line: 79, col: 43, offset: 1828},
	name: "_",
},
&litMatcher{
	pos: position{line: 79, col: 45, offset: 1830},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FakeJobG",
	pos: position{line: 86, col: 1, offset: 2045},
	expr: &actionExpr{
	pos: position{line: 86, col: 13, offset: 2057},
	run: (*parser).callonFakeJobG1,
	expr: &seqExpr{
	pos: position{line: 86, col: 13, offset: 2057},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 86, col: 13, offset: 2057},
	name: "_",
},
&litMatcher{
	pos: position{line: 86, col: 15, offset: 2059},
	val: "fake(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 86, col: 23, offset: 2067},
	name: "_",
},
&labeledExpr{
	pos: position{line: 86, col: 25, offset: 2069},
	label: "numData",
	expr: &ruleRefExpr{
	pos: position{line: 86, col: 33, offset: 2077},
	name: "PositiveNumber",
},
},
&ruleRefExpr{
	pos: position{line: 86, col: 48, offset: 2092},
	name: "_",
},
&litMatcher{
	pos: position{line: 86, col: 50, offset: 2094},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchJobG",
	pos: position{line: 93, col: 1, offset: 2293},
	expr: &actionExpr{
	pos: position{line: 93, col: 15, offset: 2307},
	run: (*parser).callonBranchJobG1,
	expr: &seqExpr{
	pos: position{line: 93, col: 15, offset: 2307},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 93, col: 15, offset: 2307},
	val: "branch",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 93, col: 23, offset: 2315},
	name: "_",
},
&litMatcher{
	pos: position{line: 93, col: 24, offset: 2316},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 93, col: 28, offset: 2320},
	name: "_",
},
&labeledExpr{
	pos: position{line: 93, col: 30, offset: 2322},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 93, col: 33, offset: 2325},
	name: "IDENTIFIER",
},
},
&litMatcher{
	pos: position{line: 93, col: 44, offset: 2336},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggJobG",
	pos: position{line: 106, col: 1, offset: 2638},
	expr: &actionExpr{
	pos: position{line: 106, col: 12, offset: 2649},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 106, col: 12, offset: 2649},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 106, col: 12, offset: 2649},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 106, col: 18, offset: 2655},
	name: "_",
},
&labeledExpr{
	pos: position{line: 106, col: 20, offset: 2657},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 106, col: 27, offset: 2664},
	expr: &ruleRefExpr{
	pos: position{line: 106, col: 27, offset: 2664},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 106, col: 39, offset: 2676},
	name: "_",
},
&labeledExpr{
	pos: position{line: 106, col: 41, offset: 2678},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 106, col: 45, offset: 2682},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 106, col: 53, offset: 2690},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 106, col: 58, offset: 2695},
	expr: &seqExpr{
	pos: position{line: 106, col: 59, offset: 2696},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 106, col: 59, offset: 2696},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 106, col: 63, offset: 2700},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 106, col: 65, offset: 2702},
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
	pos: position{line: 121, col: 1, offset: 3077},
	expr: &choiceExpr{
	pos: position{line: 121, col: 12, offset: 3088},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 121, col: 12, offset: 3088},
	name: "CountAgg",
},
&ruleRefExpr{
	pos: position{line: 121, col: 23, offset: 3099},
	name: "MaxAgg",
},
&ruleRefExpr{
	pos: position{line: 121, col: 32, offset: 3108},
	name: "MinAgg",
},
&ruleRefExpr{
	pos: position{line: 121, col: 41, offset: 3117},
	name: "AvgAgg",
},
&ruleRefExpr{
	pos: position{line: 121, col: 50, offset: 3126},
	name: "SumAgg",
},
&ruleRefExpr{
	pos: position{line: 121, col: 59, offset: 3135},
	name: "ModeAgg",
},
&ruleRefExpr{
	pos: position{line: 122, col: 13, offset: 3156},
	name: "VarAgg",
},
&ruleRefExpr{
	pos: position{line: 122, col: 22, offset: 3165},
	name: "DistinctCountAgg",
},
&ruleRefExpr{
	pos: position{line: 122, col: 41, offset: 3184},
	name: "QuantileAgg",
},
&ruleRefExpr{
	pos: position{line: 122, col: 55, offset: 3198},
	name: "MedianAgg",
},
&ruleRefExpr{
	pos: position{line: 122, col: 67, offset: 3210},
	name: "QuartileAgg",
},
&ruleRefExpr{
	pos: position{line: 123, col: 13, offset: 3235},
	name: "CovAgg",
},
&ruleRefExpr{
	pos: position{line: 123, col: 22, offset: 3244},
	name: "CorrAgg",
},
&ruleRefExpr{
	pos: position{line: 123, col: 32, offset: 3254},
	name: "PValueAgg",
},
	},
},
},
{
	name: "CountAgg",
	pos: position{line: 126, col: 1, offset: 3298},
	expr: &actionExpr{
	pos: position{line: 126, col: 13, offset: 3310},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 126, col: 13, offset: 3310},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 126, col: 13, offset: 3310},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 126, col: 19, offset: 3316},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 126, col: 28, offset: 3325},
	name: "_",
},
&litMatcher{
	pos: position{line: 126, col: 30, offset: 3327},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 126, col: 39, offset: 3336},
	name: "_",
},
&labeledExpr{
	pos: position{line: 126, col: 41, offset: 3338},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 126, col: 46, offset: 3343},
	expr: &ruleRefExpr{
	pos: position{line: 126, col: 46, offset: 3343},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 126, col: 59, offset: 3356},
	name: "_",
},
&litMatcher{
	pos: position{line: 126, col: 61, offset: 3358},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MaxAgg",
	pos: position{line: 133, col: 1, offset: 3510},
	expr: &actionExpr{
	pos: position{line: 133, col: 11, offset: 3520},
	run: (*parser).callonMaxAgg1,
	expr: &seqExpr{
	pos: position{line: 133, col: 11, offset: 3520},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 133, col: 11, offset: 3520},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 133, col: 17, offset: 3526},
	expr: &ruleRefExpr{
	pos: position{line: 133, col: 17, offset: 3526},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 133, col: 27, offset: 3536},
	name: "_",
},
&litMatcher{
	pos: position{line: 133, col: 29, offset: 3538},
	val: "max(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 133, col: 36, offset: 3545},
	name: "_",
},
&labeledExpr{
	pos: position{line: 133, col: 38, offset: 3547},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 133, col: 44, offset: 3553},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 133, col: 53, offset: 3562},
	name: "_",
},
&labeledExpr{
	pos: position{line: 133, col: 55, offset: 3564},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 133, col: 60, offset: 3569},
	expr: &seqExpr{
	pos: position{line: 133, col: 61, offset: 3570},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 133, col: 61, offset: 3570},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 133, col: 65, offset: 3574},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 133, col: 67, offset: 3576},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 133, col: 81, offset: 3590},
	name: "_",
},
&litMatcher{
	pos: position{line: 133, col: 83, offset: 3592},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MinAgg",
	pos: position{line: 155, col: 1, offset: 4053},
	expr: &actionExpr{
	pos: position{line: 155, col: 11, offset: 4063},
	run: (*parser).callonMinAgg1,
	expr: &seqExpr{
	pos: position{line: 155, col: 11, offset: 4063},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 155, col: 11, offset: 4063},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 155, col: 17, offset: 4069},
	expr: &ruleRefExpr{
	pos: position{line: 155, col: 17, offset: 4069},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 155, col: 27, offset: 4079},
	name: "_",
},
&litMatcher{
	pos: position{line: 155, col: 29, offset: 4081},
	val: "min(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 155, col: 36, offset: 4088},
	name: "_",
},
&labeledExpr{
	pos: position{line: 155, col: 38, offset: 4090},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 155, col: 44, offset: 4096},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 155, col: 53, offset: 4105},
	name: "_",
},
&labeledExpr{
	pos: position{line: 155, col: 55, offset: 4107},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 155, col: 60, offset: 4112},
	expr: &seqExpr{
	pos: position{line: 155, col: 61, offset: 4113},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 155, col: 61, offset: 4113},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 155, col: 65, offset: 4117},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 155, col: 67, offset: 4119},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 155, col: 81, offset: 4133},
	name: "_",
},
&litMatcher{
	pos: position{line: 155, col: 83, offset: 4135},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AvgAgg",
	pos: position{line: 177, col: 1, offset: 4596},
	expr: &actionExpr{
	pos: position{line: 177, col: 11, offset: 4606},
	run: (*parser).callonAvgAgg1,
	expr: &seqExpr{
	pos: position{line: 177, col: 11, offset: 4606},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 177, col: 11, offset: 4606},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 177, col: 17, offset: 4612},
	expr: &ruleRefExpr{
	pos: position{line: 177, col: 17, offset: 4612},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 177, col: 27, offset: 4622},
	name: "_",
},
&litMatcher{
	pos: position{line: 177, col: 29, offset: 4624},
	val: "avg(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 177, col: 36, offset: 4631},
	name: "_",
},
&labeledExpr{
	pos: position{line: 177, col: 38, offset: 4633},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 177, col: 44, offset: 4639},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 177, col: 53, offset: 4648},
	name: "_",
},
&labeledExpr{
	pos: position{line: 177, col: 55, offset: 4650},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 177, col: 60, offset: 4655},
	expr: &seqExpr{
	pos: position{line: 177, col: 61, offset: 4656},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 177, col: 61, offset: 4656},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 177, col: 65, offset: 4660},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 177, col: 67, offset: 4662},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 177, col: 81, offset: 4676},
	name: "_",
},
&litMatcher{
	pos: position{line: 177, col: 83, offset: 4678},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SumAgg",
	pos: position{line: 199, col: 1, offset: 5139},
	expr: &actionExpr{
	pos: position{line: 199, col: 11, offset: 5149},
	run: (*parser).callonSumAgg1,
	expr: &seqExpr{
	pos: position{line: 199, col: 11, offset: 5149},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 199, col: 11, offset: 5149},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 199, col: 17, offset: 5155},
	expr: &ruleRefExpr{
	pos: position{line: 199, col: 17, offset: 5155},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 199, col: 27, offset: 5165},
	name: "_",
},
&litMatcher{
	pos: position{line: 199, col: 29, offset: 5167},
	val: "sum(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 199, col: 36, offset: 5174},
	name: "_",
},
&labeledExpr{
	pos: position{line: 199, col: 38, offset: 5176},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 199, col: 44, offset: 5182},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 199, col: 53, offset: 5191},
	name: "_",
},
&labeledExpr{
	pos: position{line: 199, col: 55, offset: 5193},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 199, col: 60, offset: 5198},
	expr: &seqExpr{
	pos: position{line: 199, col: 61, offset: 5199},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 199, col: 61, offset: 5199},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 199, col: 65, offset: 5203},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 199, col: 67, offset: 5205},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 199, col: 81, offset: 5219},
	name: "_",
},
&litMatcher{
	pos: position{line: 199, col: 83, offset: 5221},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarAgg",
	pos: position{line: 221, col: 1, offset: 5682},
	expr: &actionExpr{
	pos: position{line: 221, col: 11, offset: 5692},
	run: (*parser).callonVarAgg1,
	expr: &seqExpr{
	pos: position{line: 221, col: 11, offset: 5692},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 221, col: 11, offset: 5692},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 221, col: 17, offset: 5698},
	expr: &ruleRefExpr{
	pos: position{line: 221, col: 17, offset: 5698},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 221, col: 27, offset: 5708},
	name: "_",
},
&litMatcher{
	pos: position{line: 221, col: 29, offset: 5710},
	val: "var(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 221, col: 36, offset: 5717},
	name: "_",
},
&labeledExpr{
	pos: position{line: 221, col: 38, offset: 5719},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 221, col: 44, offset: 5725},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 221, col: 53, offset: 5734},
	name: "_",
},
&labeledExpr{
	pos: position{line: 221, col: 55, offset: 5736},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 221, col: 60, offset: 5741},
	expr: &seqExpr{
	pos: position{line: 221, col: 61, offset: 5742},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 221, col: 61, offset: 5742},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 221, col: 65, offset: 5746},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 221, col: 67, offset: 5748},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 221, col: 81, offset: 5762},
	name: "_",
},
&litMatcher{
	pos: position{line: 221, col: 83, offset: 5764},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DistinctCountAgg",
	pos: position{line: 243, col: 1, offset: 6230},
	expr: &actionExpr{
	pos: position{line: 243, col: 21, offset: 6250},
	run: (*parser).callonDistinctCountAgg1,
	expr: &seqExpr{
	pos: position{line: 243, col: 21, offset: 6250},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 243, col: 21, offset: 6250},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 243, col: 27, offset: 6256},
	expr: &ruleRefExpr{
	pos: position{line: 243, col: 27, offset: 6256},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 243, col: 37, offset: 6266},
	name: "_",
},
&litMatcher{
	pos: position{line: 243, col: 39, offset: 6268},
	val: "dcount(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 243, col: 49, offset: 6278},
	name: "_",
},
&labeledExpr{
	pos: position{line: 243, col: 51, offset: 6280},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 243, col: 57, offset: 6286},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 243, col: 66, offset: 6295},
	name: "_",
},
&labeledExpr{
	pos: position{line: 243, col: 68, offset: 6297},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 243, col: 73, offset: 6302},
	expr: &seqExpr{
	pos: position{line: 243, col: 74, offset: 6303},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 243, col: 74, offset: 6303},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 243, col: 78, offset: 6307},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 243, col: 80, offset: 6309},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 243, col: 94, offset: 6323},
	name: "_",
},
&litMatcher{
	pos: position{line: 243, col: 96, offset: 6325},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "ModeAgg",
	pos: position{line: 265, col: 1, offset: 6796},
	expr: &actionExpr{
	pos: position{line: 265, col: 12, offset: 6807},
	run: (*parser).callonModeAgg1,
	expr: &seqExpr{
	pos: position{line: 265, col: 12, offset: 6807},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 265, col: 12, offset: 6807},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 265, col: 18, offset: 6813},
	expr: &ruleRefExpr{
	pos: position{line: 265, col: 18, offset: 6813},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 265, col: 28, offset: 6823},
	name: "_",
},
&litMatcher{
	pos: position{line: 265, col: 30, offset: 6825},
	val: "mode(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 265, col: 38, offset: 6833},
	name: "_",
},
&labeledExpr{
	pos: position{line: 265, col: 40, offset: 6835},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 265, col: 46, offset: 6841},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 265, col: 55, offset: 6850},
	name: "_",
},
&labeledExpr{
	pos: position{line: 265, col: 57, offset: 6852},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 265, col: 62, offset: 6857},
	expr: &seqExpr{
	pos: position{line: 265, col: 63, offset: 6858},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 265, col: 63, offset: 6858},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 265, col: 67, offset: 6862},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 265, col: 69, offset: 6864},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 265, col: 83, offset: 6878},
	name: "_",
},
&litMatcher{
	pos: position{line: 265, col: 85, offset: 6880},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuantileAgg",
	pos: position{line: 287, col: 1, offset: 7342},
	expr: &actionExpr{
	pos: position{line: 287, col: 16, offset: 7357},
	run: (*parser).callonQuantileAgg1,
	expr: &seqExpr{
	pos: position{line: 287, col: 16, offset: 7357},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 287, col: 16, offset: 7357},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 287, col: 22, offset: 7363},
	expr: &ruleRefExpr{
	pos: position{line: 287, col: 22, offset: 7363},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 287, col: 32, offset: 7373},
	name: "_",
},
&litMatcher{
	pos: position{line: 287, col: 34, offset: 7375},
	val: "quantile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 287, col: 46, offset: 7387},
	name: "_",
},
&labeledExpr{
	pos: position{line: 287, col: 48, offset: 7389},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 287, col: 54, offset: 7395},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 287, col: 63, offset: 7404},
	name: "_",
},
&labeledExpr{
	pos: position{line: 287, col: 66, offset: 7407},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 287, col: 73, offset: 7414},
	expr: &seqExpr{
	pos: position{line: 287, col: 74, offset: 7415},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 287, col: 74, offset: 7415},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 287, col: 78, offset: 7419},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 287, col: 80, offset: 7421},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 287, col: 91, offset: 7432},
	name: "_",
},
&litMatcher{
	pos: position{line: 287, col: 93, offset: 7434},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 287, col: 97, offset: 7438},
	name: "_",
},
&labeledExpr{
	pos: position{line: 287, col: 99, offset: 7440},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 287, col: 103, offset: 7444},
	name: "Float",
},
},
&ruleRefExpr{
	pos: position{line: 287, col: 109, offset: 7450},
	name: "_",
},
&labeledExpr{
	pos: position{line: 287, col: 111, offset: 7452},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 287, col: 116, offset: 7457},
	expr: &seqExpr{
	pos: position{line: 287, col: 117, offset: 7458},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 287, col: 117, offset: 7458},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 287, col: 121, offset: 7462},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 287, col: 123, offset: 7464},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 287, col: 137, offset: 7478},
	name: "_",
},
&litMatcher{
	pos: position{line: 287, col: 139, offset: 7480},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MedianAgg",
	pos: position{line: 317, col: 1, offset: 8125},
	expr: &actionExpr{
	pos: position{line: 317, col: 14, offset: 8138},
	run: (*parser).callonMedianAgg1,
	expr: &seqExpr{
	pos: position{line: 317, col: 14, offset: 8138},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 317, col: 14, offset: 8138},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 317, col: 20, offset: 8144},
	expr: &ruleRefExpr{
	pos: position{line: 317, col: 20, offset: 8144},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 317, col: 30, offset: 8154},
	name: "_",
},
&litMatcher{
	pos: position{line: 317, col: 32, offset: 8156},
	val: "median(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 317, col: 42, offset: 8166},
	name: "_",
},
&labeledExpr{
	pos: position{line: 317, col: 44, offset: 8168},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 317, col: 50, offset: 8174},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 317, col: 59, offset: 8183},
	name: "_",
},
&labeledExpr{
	pos: position{line: 317, col: 62, offset: 8186},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 317, col: 69, offset: 8193},
	expr: &seqExpr{
	pos: position{line: 317, col: 70, offset: 8194},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 317, col: 70, offset: 8194},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 317, col: 74, offset: 8198},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 317, col: 76, offset: 8200},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 317, col: 87, offset: 8211},
	name: "_",
},
&labeledExpr{
	pos: position{line: 317, col: 89, offset: 8213},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 317, col: 94, offset: 8218},
	expr: &seqExpr{
	pos: position{line: 317, col: 95, offset: 8219},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 317, col: 95, offset: 8219},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 317, col: 99, offset: 8223},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 317, col: 101, offset: 8225},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 317, col: 115, offset: 8239},
	name: "_",
},
&litMatcher{
	pos: position{line: 317, col: 117, offset: 8241},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuartileAgg",
	pos: position{line: 349, col: 1, offset: 8918},
	expr: &actionExpr{
	pos: position{line: 349, col: 16, offset: 8933},
	run: (*parser).callonQuartileAgg1,
	expr: &seqExpr{
	pos: position{line: 349, col: 16, offset: 8933},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 349, col: 16, offset: 8933},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 349, col: 22, offset: 8939},
	expr: &ruleRefExpr{
	pos: position{line: 349, col: 22, offset: 8939},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 349, col: 32, offset: 8949},
	name: "_",
},
&litMatcher{
	pos: position{line: 349, col: 34, offset: 8951},
	val: "quartile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 349, col: 46, offset: 8963},
	name: "_",
},
&labeledExpr{
	pos: position{line: 349, col: 48, offset: 8965},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 349, col: 54, offset: 8971},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 349, col: 63, offset: 8980},
	name: "_",
},
&labeledExpr{
	pos: position{line: 349, col: 66, offset: 8983},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 349, col: 73, offset: 8990},
	expr: &seqExpr{
	pos: position{line: 349, col: 74, offset: 8991},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 349, col: 74, offset: 8991},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 349, col: 78, offset: 8995},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 349, col: 80, offset: 8997},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 349, col: 91, offset: 9008},
	name: "_",
},
&litMatcher{
	pos: position{line: 349, col: 93, offset: 9010},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 349, col: 97, offset: 9014},
	name: "_",
},
&labeledExpr{
	pos: position{line: 349, col: 99, offset: 9016},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 349, col: 103, offset: 9020},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 349, col: 110, offset: 9027},
	name: "_",
},
&labeledExpr{
	pos: position{line: 349, col: 112, offset: 9029},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 349, col: 117, offset: 9034},
	expr: &seqExpr{
	pos: position{line: 349, col: 118, offset: 9035},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 349, col: 118, offset: 9035},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 349, col: 122, offset: 9039},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 349, col: 124, offset: 9041},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 349, col: 138, offset: 9055},
	name: "_",
},
&litMatcher{
	pos: position{line: 349, col: 140, offset: 9057},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CovAgg",
	pos: position{line: 395, col: 1, offset: 10084},
	expr: &actionExpr{
	pos: position{line: 395, col: 11, offset: 10094},
	run: (*parser).callonCovAgg1,
	expr: &seqExpr{
	pos: position{line: 395, col: 11, offset: 10094},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 395, col: 11, offset: 10094},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 395, col: 17, offset: 10100},
	expr: &ruleRefExpr{
	pos: position{line: 395, col: 17, offset: 10100},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 395, col: 27, offset: 10110},
	name: "_",
},
&litMatcher{
	pos: position{line: 395, col: 29, offset: 10112},
	val: "cov(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 395, col: 36, offset: 10119},
	name: "_",
},
&labeledExpr{
	pos: position{line: 395, col: 38, offset: 10121},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 395, col: 45, offset: 10128},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 395, col: 54, offset: 10137},
	name: "_",
},
&litMatcher{
	pos: position{line: 395, col: 56, offset: 10139},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 395, col: 60, offset: 10143},
	name: "_",
},
&labeledExpr{
	pos: position{line: 395, col: 62, offset: 10145},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 395, col: 69, offset: 10152},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 395, col: 78, offset: 10161},
	name: "_",
},
&labeledExpr{
	pos: position{line: 395, col: 80, offset: 10163},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 395, col: 85, offset: 10168},
	expr: &seqExpr{
	pos: position{line: 395, col: 86, offset: 10169},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 395, col: 86, offset: 10169},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 395, col: 90, offset: 10173},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 395, col: 92, offset: 10175},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 395, col: 106, offset: 10189},
	name: "_",
},
&litMatcher{
	pos: position{line: 395, col: 108, offset: 10191},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CorrAgg",
	pos: position{line: 418, col: 1, offset: 10716},
	expr: &actionExpr{
	pos: position{line: 418, col: 12, offset: 10727},
	run: (*parser).callonCorrAgg1,
	expr: &seqExpr{
	pos: position{line: 418, col: 12, offset: 10727},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 418, col: 12, offset: 10727},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 418, col: 18, offset: 10733},
	expr: &ruleRefExpr{
	pos: position{line: 418, col: 18, offset: 10733},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 418, col: 28, offset: 10743},
	name: "_",
},
&litMatcher{
	pos: position{line: 418, col: 30, offset: 10745},
	val: "correlation(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 418, col: 45, offset: 10760},
	name: "_",
},
&labeledExpr{
	pos: position{line: 418, col: 47, offset: 10762},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 418, col: 54, offset: 10769},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 418, col: 63, offset: 10778},
	name: "_",
},
&litMatcher{
	pos: position{line: 418, col: 65, offset: 10780},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 418, col: 69, offset: 10784},
	name: "_",
},
&labeledExpr{
	pos: position{line: 418, col: 71, offset: 10786},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 418, col: 78, offset: 10793},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 418, col: 87, offset: 10802},
	name: "_",
},
&labeledExpr{
	pos: position{line: 418, col: 89, offset: 10804},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 418, col: 94, offset: 10809},
	expr: &seqExpr{
	pos: position{line: 418, col: 95, offset: 10810},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 418, col: 95, offset: 10810},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 418, col: 99, offset: 10814},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 418, col: 101, offset: 10816},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 418, col: 115, offset: 10830},
	name: "_",
},
&litMatcher{
	pos: position{line: 418, col: 117, offset: 10832},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PValueAgg",
	pos: position{line: 441, col: 1, offset: 11358},
	expr: &actionExpr{
	pos: position{line: 441, col: 14, offset: 11371},
	run: (*parser).callonPValueAgg1,
	expr: &seqExpr{
	pos: position{line: 441, col: 14, offset: 11371},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 441, col: 14, offset: 11371},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 441, col: 20, offset: 11377},
	expr: &ruleRefExpr{
	pos: position{line: 441, col: 20, offset: 11377},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 441, col: 30, offset: 11387},
	name: "_",
},
&litMatcher{
	pos: position{line: 441, col: 32, offset: 11389},
	val: "pvalue(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 441, col: 42, offset: 11399},
	name: "_",
},
&labeledExpr{
	pos: position{line: 441, col: 44, offset: 11401},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 441, col: 51, offset: 11408},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 441, col: 60, offset: 11417},
	name: "_",
},
&litMatcher{
	pos: position{line: 441, col: 62, offset: 11419},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 441, col: 66, offset: 11423},
	name: "_",
},
&labeledExpr{
	pos: position{line: 441, col: 68, offset: 11425},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 441, col: 75, offset: 11432},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 441, col: 84, offset: 11441},
	name: "_",
},
&labeledExpr{
	pos: position{line: 441, col: 86, offset: 11443},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 441, col: 91, offset: 11448},
	expr: &seqExpr{
	pos: position{line: 441, col: 92, offset: 11449},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 441, col: 92, offset: 11449},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 441, col: 96, offset: 11453},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 441, col: 98, offset: 11455},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 441, col: 112, offset: 11469},
	name: "_",
},
&litMatcher{
	pos: position{line: 441, col: 114, offset: 11471},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggGroupBy",
	pos: position{line: 466, col: 1, offset: 11996},
	expr: &actionExpr{
	pos: position{line: 466, col: 15, offset: 12010},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 466, col: 15, offset: 12010},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 466, col: 15, offset: 12010},
	name: "_",
},
&litMatcher{
	pos: position{line: 466, col: 17, offset: 12012},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 466, col: 22, offset: 12017},
	name: "_",
},
&labeledExpr{
	pos: position{line: 466, col: 24, offset: 12019},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 466, col: 30, offset: 12025},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 466, col: 39, offset: 12034},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 466, col: 44, offset: 12039},
	expr: &seqExpr{
	pos: position{line: 466, col: 45, offset: 12040},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 466, col: 45, offset: 12040},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 466, col: 49, offset: 12044},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 466, col: 51, offset: 12046},
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
	pos: position{line: 481, col: 1, offset: 12397},
	expr: &actionExpr{
	pos: position{line: 481, col: 11, offset: 12407},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 481, col: 11, offset: 12407},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 481, col: 16, offset: 12412},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 481, col: 16, offset: 12412},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 481, col: 27, offset: 12423},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 481, col: 38, offset: 12434},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 481, col: 49, offset: 12445},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 481, col: 58, offset: 12454},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 481, col: 67, offset: 12463},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 487, col: 1, offset: 12582},
	expr: &actionExpr{
	pos: position{line: 487, col: 13, offset: 12594},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 487, col: 13, offset: 12594},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 487, col: 13, offset: 12594},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 487, col: 23, offset: 12604},
	name: "_",
},
&labeledExpr{
	pos: position{line: 487, col: 25, offset: 12606},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 487, col: 28, offset: 12609},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 487, col: 42, offset: 12623},
	name: "_",
},
&litMatcher{
	pos: position{line: 487, col: 44, offset: 12625},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 494, col: 1, offset: 12820},
	expr: &actionExpr{
	pos: position{line: 494, col: 18, offset: 12837},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 494, col: 18, offset: 12837},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 494, col: 18, offset: 12837},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 494, col: 23, offset: 12842},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 494, col: 28, offset: 12847},
	name: "_",
},
&labeledExpr{
	pos: position{line: 494, col: 30, offset: 12849},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 494, col: 33, offset: 12852},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 494, col: 45, offset: 12864},
	name: "_",
},
&labeledExpr{
	pos: position{line: 494, col: 47, offset: 12866},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 494, col: 54, offset: 12873},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 498, col: 1, offset: 12924},
	expr: &ruleRefExpr{
	pos: position{line: 498, col: 16, offset: 12939},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 501, col: 1, offset: 12973},
	expr: &actionExpr{
	pos: position{line: 501, col: 13, offset: 12985},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 501, col: 13, offset: 12985},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 501, col: 13, offset: 12985},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 501, col: 23, offset: 12995},
	name: "_",
},
&labeledExpr{
	pos: position{line: 501, col: 25, offset: 12997},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 501, col: 31, offset: 13003},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 501, col: 40, offset: 13012},
	name: "_",
},
&labeledExpr{
	pos: position{line: 501, col: 42, offset: 13014},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 501, col: 47, offset: 13019},
	expr: &seqExpr{
	pos: position{line: 501, col: 48, offset: 13020},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 501, col: 48, offset: 13020},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 501, col: 52, offset: 13024},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 501, col: 54, offset: 13026},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 501, col: 65, offset: 13037},
	name: "_",
},
&litMatcher{
	pos: position{line: 501, col: 67, offset: 13039},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 516, col: 1, offset: 13376},
	expr: &actionExpr{
	pos: position{line: 516, col: 11, offset: 13386},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 516, col: 11, offset: 13386},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 516, col: 11, offset: 13386},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 516, col: 18, offset: 13393},
	name: "_",
},
&labeledExpr{
	pos: position{line: 516, col: 20, offset: 13395},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 516, col: 25, offset: 13400},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 516, col: 34, offset: 13409},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 516, col: 38, offset: 13413},
	name: "_",
},
&labeledExpr{
	pos: position{line: 516, col: 40, offset: 13415},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 516, col: 44, offset: 13419},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 516, col: 51, offset: 13426},
	name: "_",
},
&litMatcher{
	pos: position{line: 516, col: 53, offset: 13428},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 528, col: 1, offset: 13650},
	expr: &actionExpr{
	pos: position{line: 528, col: 11, offset: 13660},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 528, col: 11, offset: 13660},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 528, col: 11, offset: 13660},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 528, col: 20, offset: 13669},
	name: "_",
},
&litMatcher{
	pos: position{line: 528, col: 22, offset: 13671},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 528, col: 27, offset: 13676},
	name: "_",
},
&labeledExpr{
	pos: position{line: 528, col: 29, offset: 13678},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 528, col: 35, offset: 13684},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 534, col: 1, offset: 13790},
	expr: &actionExpr{
	pos: position{line: 534, col: 12, offset: 13801},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 534, col: 12, offset: 13801},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 534, col: 12, offset: 13801},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 534, col: 21, offset: 13810},
	name: "_",
},
&labeledExpr{
	pos: position{line: 534, col: 23, offset: 13812},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 534, col: 27, offset: 13816},
	expr: &ruleRefExpr{
	pos: position{line: 534, col: 27, offset: 13816},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 534, col: 35, offset: 13824},
	name: "_",
},
&litMatcher{
	pos: position{line: 534, col: 37, offset: 13826},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 544, col: 1, offset: 13971},
	expr: &actionExpr{
	pos: position{line: 544, col: 13, offset: 13983},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 544, col: 13, offset: 13983},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 544, col: 13, offset: 13983},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 544, col: 23, offset: 13993},
	name: "_",
},
&labeledExpr{
	pos: position{line: 544, col: 25, offset: 13995},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 544, col: 30, offset: 14000},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 544, col: 42, offset: 14012},
	name: "_",
},
&litMatcher{
	pos: position{line: 544, col: 44, offset: 14014},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 554, col: 1, offset: 14302},
	expr: &actionExpr{
	pos: position{line: 554, col: 16, offset: 14317},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 554, col: 16, offset: 14317},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 554, col: 16, offset: 14317},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 554, col: 22, offset: 14323},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 554, col: 35, offset: 14336},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 554, col: 40, offset: 14341},
	expr: &seqExpr{
	pos: position{line: 554, col: 41, offset: 14342},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 554, col: 41, offset: 14342},
	name: "_",
},
&labeledExpr{
	pos: position{line: 554, col: 43, offset: 14344},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 554, col: 46, offset: 14347},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 554, col: 58, offset: 14359},
	name: "_",
},
&labeledExpr{
	pos: position{line: 554, col: 60, offset: 14361},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 554, col: 63, offset: 14364},
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
	pos: position{line: 559, col: 1, offset: 14522},
	expr: &choiceExpr{
	pos: position{line: 559, col: 17, offset: 14538},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 559, col: 17, offset: 14538},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 559, col: 17, offset: 14538},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 559, col: 17, offset: 14538},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 559, col: 21, offset: 14542},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 559, col: 24, offset: 14545},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 559, col: 36, offset: 14557},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 561, col: 5, offset: 14588},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 561, col: 5, offset: 14588},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 561, col: 8, offset: 14591},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 565, col: 1, offset: 14633},
	expr: &choiceExpr{
	pos: position{line: 565, col: 16, offset: 14648},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 565, col: 16, offset: 14648},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 565, col: 24, offset: 14656},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 565, col: 24, offset: 14656},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 572, col: 1, offset: 14835},
	expr: &actionExpr{
	pos: position{line: 572, col: 18, offset: 14852},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 572, col: 18, offset: 14852},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 572, col: 18, offset: 14852},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 572, col: 23, offset: 14857},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 572, col: 28, offset: 14862},
	name: "_",
},
&labeledExpr{
	pos: position{line: 572, col: 30, offset: 14864},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 572, col: 33, offset: 14867},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 572, col: 45, offset: 14879},
	name: "_",
},
&labeledExpr{
	pos: position{line: 572, col: 47, offset: 14881},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 572, col: 53, offset: 14887},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 572, col: 59, offset: 14893},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 576, col: 1, offset: 14967},
	expr: &choiceExpr{
	pos: position{line: 576, col: 16, offset: 14982},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 576, col: 16, offset: 14982},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 576, col: 23, offset: 14989},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 576, col: 30, offset: 14996},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 576, col: 36, offset: 15002},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 576, col: 42, offset: 15008},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 576, col: 49, offset: 15015},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 576, col: 49, offset: 15015},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 585, col: 1, offset: 15125},
	expr: &actionExpr{
	pos: position{line: 585, col: 12, offset: 15136},
	run: (*parser).callonStdOutG1,
	expr: &seqExpr{
	pos: position{line: 585, col: 12, offset: 15136},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 585, col: 12, offset: 15136},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 585, col: 22, offset: 15146},
	name: "_",
},
&labeledExpr{
	pos: position{line: 585, col: 24, offset: 15148},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 585, col: 31, offset: 15155},
	expr: &ruleRefExpr{
	pos: position{line: 585, col: 32, offset: 15156},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 585, col: 48, offset: 15172},
	name: "_",
},
&litMatcher{
	pos: position{line: 585, col: 50, offset: 15174},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SinkInto",
	pos: position{line: 599, col: 1, offset: 15453},
	expr: &actionExpr{
	pos: position{line: 599, col: 12, offset: 15464},
	run: (*parser).callonSinkInto1,
	expr: &seqExpr{
	pos: position{line: 599, col: 12, offset: 15464},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 599, col: 12, offset: 15464},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 599, col: 18, offset: 15470},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 599, col: 27, offset: 15479},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 599, col: 33, offset: 15485},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 599, col: 33, offset: 15485},
	name: "_",
},
&litMatcher{
	pos: position{line: 599, col: 35, offset: 15487},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 599, col: 39, offset: 15491},
	name: "_",
},
&litMatcher{
	pos: position{line: 599, col: 41, offset: 15493},
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
	pos: position{line: 607, col: 1, offset: 15649},
	expr: &actionExpr{
	pos: position{line: 607, col: 15, offset: 15663},
	run: (*parser).callonBlackHoleG1,
	expr: &seqExpr{
	pos: position{line: 607, col: 15, offset: 15663},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 607, col: 15, offset: 15663},
	val: "blackhole(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 607, col: 28, offset: 15676},
	name: "_",
},
&labeledExpr{
	pos: position{line: 607, col: 30, offset: 15678},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 607, col: 37, offset: 15685},
	expr: &ruleRefExpr{
	pos: position{line: 607, col: 38, offset: 15686},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 607, col: 54, offset: 15702},
	name: "_",
},
&litMatcher{
	pos: position{line: 607, col: 56, offset: 15704},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PlotG",
	pos: position{line: 611, col: 1, offset: 15762},
	expr: &actionExpr{
	pos: position{line: 611, col: 10, offset: 15771},
	run: (*parser).callonPlotG1,
	expr: &seqExpr{
	pos: position{line: 611, col: 10, offset: 15771},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 611, col: 10, offset: 15771},
	val: "plot",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 611, col: 17, offset: 15778},
	name: "_",
},
&labeledExpr{
	pos: position{line: 611, col: 19, offset: 15780},
	label: "widget",
	expr: &ruleRefExpr{
	pos: position{line: 611, col: 27, offset: 15788},
	name: "BarWidget",
},
},
	},
},
},
},
{
	name: "BarWidget",
	pos: position{line: 615, col: 1, offset: 15858},
	expr: &actionExpr{
	pos: position{line: 615, col: 14, offset: 15871},
	run: (*parser).callonBarWidget1,
	expr: &seqExpr{
	pos: position{line: 615, col: 14, offset: 15871},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 615, col: 14, offset: 15871},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 615, col: 20, offset: 15877},
	expr: &ruleRefExpr{
	pos: position{line: 615, col: 20, offset: 15877},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 615, col: 30, offset: 15887},
	name: "_",
},
&litMatcher{
	pos: position{line: 615, col: 32, offset: 15889},
	val: "bar(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 615, col: 39, offset: 15896},
	name: "_",
},
&labeledExpr{
	pos: position{line: 615, col: 41, offset: 15898},
	label: "xField",
	expr: &ruleRefExpr{
	pos: position{line: 615, col: 48, offset: 15905},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 615, col: 57, offset: 15914},
	name: "_",
},
&litMatcher{
	pos: position{line: 615, col: 59, offset: 15916},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 615, col: 63, offset: 15920},
	name: "_",
},
&labeledExpr{
	pos: position{line: 615, col: 65, offset: 15922},
	label: "yField",
	expr: &ruleRefExpr{
	pos: position{line: 615, col: 72, offset: 15929},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 615, col: 81, offset: 15938},
	name: "_",
},
&labeledExpr{
	pos: position{line: 615, col: 83, offset: 15940},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 615, col: 88, offset: 15945},
	expr: &seqExpr{
	pos: position{line: 615, col: 89, offset: 15946},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 615, col: 89, offset: 15946},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 615, col: 93, offset: 15950},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 615, col: 95, offset: 15952},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 615, col: 102, offset: 15959},
	name: "_",
},
&zeroOrOneExpr{
	pos: position{line: 615, col: 104, offset: 15961},
	expr: &seqExpr{
	pos: position{line: 615, col: 105, offset: 15962},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 615, col: 105, offset: 15962},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 615, col: 109, offset: 15966},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 615, col: 111, offset: 15968},
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
	pos: position{line: 615, col: 122, offset: 15979},
	name: "_",
},
&litMatcher{
	pos: position{line: 615, col: 124, offset: 15981},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FileName",
	pos: position{line: 648, col: 1, offset: 16666},
	expr: &actionExpr{
	pos: position{line: 648, col: 13, offset: 16678},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 648, col: 13, offset: 16678},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 648, col: 13, offset: 16678},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 648, col: 22, offset: 16687},
	expr: &seqExpr{
	pos: position{line: 648, col: 23, offset: 16688},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 648, col: 23, offset: 16688},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 648, col: 27, offset: 16692},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 648, col: 31, offset: 16696},
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
	pos: position{line: 654, col: 1, offset: 16797},
	expr: &actionExpr{
	pos: position{line: 654, col: 10, offset: 16806},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 654, col: 10, offset: 16806},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 654, col: 10, offset: 16806},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 654, col: 16, offset: 16812},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 654, col: 23, offset: 16819},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 654, col: 28, offset: 16824},
	expr: &seqExpr{
	pos: position{line: 654, col: 29, offset: 16825},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 654, col: 29, offset: 16825},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 654, col: 31, offset: 16827},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 654, col: 36, offset: 16832},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 654, col: 38, offset: 16834},
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
	pos: position{line: 658, col: 1, offset: 16899},
	expr: &actionExpr{
	pos: position{line: 658, col: 11, offset: 16909},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 658, col: 11, offset: 16909},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 658, col: 16, offset: 16914},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 658, col: 16, offset: 16914},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 658, col: 24, offset: 16922},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 658, col: 33, offset: 16931},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 658, col: 48, offset: 16946},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 658, col: 59, offset: 16957},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 662, col: 1, offset: 16999},
	expr: &choiceExpr{
	pos: position{line: 662, col: 9, offset: 17007},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 662, col: 9, offset: 17007},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 662, col: 15, offset: 17013},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 662, col: 21, offset: 17019},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 662, col: 27, offset: 17025},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 662, col: 33, offset: 17031},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 662, col: 33, offset: 17031},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 666, col: 1, offset: 17079},
	expr: &choiceExpr{
	pos: position{line: 666, col: 12, offset: 17090},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 666, col: 12, offset: 17090},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 666, col: 12, offset: 17090},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 666, col: 12, offset: 17090},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 666, col: 16, offset: 17094},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 666, col: 19, offset: 17097},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 666, col: 25, offset: 17103},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 670, col: 5, offset: 17179},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 670, col: 5, offset: 17179},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 670, col: 8, offset: 17182},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 677, col: 1, offset: 17259},
	expr: &actionExpr{
	pos: position{line: 677, col: 9, offset: 17267},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 677, col: 9, offset: 17267},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 677, col: 9, offset: 17267},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 677, col: 15, offset: 17273},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 677, col: 21, offset: 17279},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 677, col: 26, offset: 17284},
	expr: &seqExpr{
	pos: position{line: 677, col: 27, offset: 17285},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 677, col: 27, offset: 17285},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 677, col: 29, offset: 17287},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 677, col: 34, offset: 17292},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 677, col: 36, offset: 17294},
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
	pos: position{line: 681, col: 1, offset: 17358},
	expr: &actionExpr{
	pos: position{line: 681, col: 10, offset: 17367},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 681, col: 10, offset: 17367},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 681, col: 15, offset: 17372},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 681, col: 15, offset: 17372},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 681, col: 23, offset: 17380},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 681, col: 32, offset: 17389},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 681, col: 47, offset: 17404},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 685, col: 1, offset: 17447},
	expr: &actionExpr{
	pos: position{line: 685, col: 13, offset: 17459},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 685, col: 13, offset: 17459},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 685, col: 13, offset: 17459},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 685, col: 20, offset: 17466},
	expr: &choiceExpr{
	pos: position{line: 685, col: 21, offset: 17467},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 685, col: 21, offset: 17467},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 685, col: 31, offset: 17477},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 685, col: 39, offset: 17485},
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
	pos: position{line: 689, col: 1, offset: 17535},
	expr: &actionExpr{
	pos: position{line: 689, col: 17, offset: 17551},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 689, col: 17, offset: 17551},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 689, col: 17, offset: 17551},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 689, col: 21, offset: 17555},
	expr: &choiceExpr{
	pos: position{line: 689, col: 23, offset: 17557},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 689, col: 23, offset: 17557},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 689, col: 23, offset: 17557},
	expr: &ruleRefExpr{
	pos: position{line: 689, col: 24, offset: 17558},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 689, col: 36, offset: 17570,
},
	},
},
&seqExpr{
	pos: position{line: 689, col: 40, offset: 17574},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 689, col: 40, offset: 17574},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 689, col: 45, offset: 17579},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 689, col: 63, offset: 17597},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 696, col: 1, offset: 17778},
	expr: &charClassMatcher{
	pos: position{line: 696, col: 16, offset: 17793},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 698, col: 1, offset: 17811},
	expr: &ruleRefExpr{
	pos: position{line: 698, col: 19, offset: 17829},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 700, col: 1, offset: 17849},
	expr: &charClassMatcher{
	pos: position{line: 700, col: 21, offset: 17869},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 702, col: 1, offset: 17884},
	expr: &actionExpr{
	pos: position{line: 702, col: 9, offset: 17892},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 702, col: 9, offset: 17892},
	expr: &choiceExpr{
	pos: position{line: 702, col: 10, offset: 17893},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 702, col: 10, offset: 17893},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 702, col: 19, offset: 17902},
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
	pos: position{line: 706, col: 1, offset: 17952},
	expr: &choiceExpr{
	pos: position{line: 706, col: 11, offset: 17962},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 706, col: 11, offset: 17962},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 706, col: 19, offset: 17970},
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
	pos: position{line: 708, col: 1, offset: 17979},
	expr: &actionExpr{
	pos: position{line: 708, col: 11, offset: 17989},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 708, col: 11, offset: 17989},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 708, col: 11, offset: 17989},
	expr: &litMatcher{
	pos: position{line: 708, col: 11, offset: 17989},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 708, col: 16, offset: 17994},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 712, col: 1, offset: 18067},
	expr: &actionExpr{
	pos: position{line: 712, col: 19, offset: 18085},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 712, col: 19, offset: 18085},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 716, col: 1, offset: 18158},
	expr: &actionExpr{
	pos: position{line: 716, col: 10, offset: 18167},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 716, col: 10, offset: 18167},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 716, col: 10, offset: 18167},
	expr: &litMatcher{
	pos: position{line: 716, col: 10, offset: 18167},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 716, col: 15, offset: 18172},
	expr: &ruleRefExpr{
	pos: position{line: 716, col: 15, offset: 18172},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 716, col: 24, offset: 18181},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 716, col: 28, offset: 18185},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 720, col: 1, offset: 18252},
	expr: &choiceExpr{
	pos: position{line: 720, col: 12, offset: 18263},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 720, col: 12, offset: 18263},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 720, col: 18, offset: 18269},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 720, col: 18, offset: 18269},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 720, col: 38, offset: 18289},
	expr: &ruleRefExpr{
	pos: position{line: 720, col: 38, offset: 18289},
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
	pos: position{line: 722, col: 1, offset: 18306},
	expr: &charClassMatcher{
	pos: position{line: 722, col: 17, offset: 18322},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 724, col: 1, offset: 18331},
	expr: &charClassMatcher{
	pos: position{line: 724, col: 24, offset: 18354},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 726, col: 1, offset: 18363},
	expr: &zeroOrMoreExpr{
	pos: position{line: 726, col: 19, offset: 18381},
	expr: &charClassMatcher{
	pos: position{line: 726, col: 19, offset: 18381},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 728, col: 1, offset: 18395},
	expr: &charClassMatcher{
	pos: position{line: 728, col: 18, offset: 18412},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "IDENTIFIER",
	pos: position{line: 729, col: 1, offset: 18423},
	expr: &actionExpr{
	pos: position{line: 729, col: 15, offset: 18437},
	run: (*parser).callonIDENTIFIER1,
	expr: &seqExpr{
	pos: position{line: 729, col: 15, offset: 18437},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 729, col: 15, offset: 18437},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 729, col: 22, offset: 18444},
	expr: &choiceExpr{
	pos: position{line: 729, col: 23, offset: 18445},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 729, col: 23, offset: 18445},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 729, col: 33, offset: 18455},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 729, col: 41, offset: 18463},
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
	pos: position{line: 732, col: 1, offset: 18536},
	expr: &notExpr{
	pos: position{line: 732, col: 8, offset: 18543},
	expr: &anyMatcher{
	line: 732, col: 9, offset: 18544,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 734, col: 1, offset: 18549},
	expr: &litMatcher{
	pos: position{line: 734, col: 17, offset: 18565},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 740, col: 1, offset: 18652},
	expr: &actionExpr{
	pos: position{line: 740, col: 18, offset: 18669},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 740, col: 18, offset: 18669},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 740, col: 18, offset: 18669},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 740, col: 24, offset: 18675},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 740, col: 33, offset: 18684},
	name: "_",
},
&labeledExpr{
	pos: position{line: 740, col: 35, offset: 18686},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 740, col: 40, offset: 18691},
	expr: &seqExpr{
	pos: position{line: 740, col: 41, offset: 18692},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 740, col: 41, offset: 18692},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 740, col: 45, offset: 18696},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 740, col: 47, offset: 18698},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 740, col: 56, offset: 18707},
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
	pos: position{line: 743, col: 1, offset: 18760},
	expr: &actionExpr{
	pos: position{line: 743, col: 13, offset: 18772},
	run: (*parser).callonEqAliasG1,
	expr: &seqExpr{
	pos: position{line: 743, col: 13, offset: 18772},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 743, col: 13, offset: 18772},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 743, col: 19, offset: 18778},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 743, col: 28, offset: 18787},
	name: "_",
},
&litMatcher{
	pos: position{line: 743, col: 30, offset: 18789},
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

func (c *current) onBlackHoleG1(params interface{}) (interface{}, error) {

        return SinkJob{Type: BLACKHOLE}, nil
}

func (p *parser) callonBlackHoleG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlackHoleG1(stack["params"])
}

func (c *current) onPlotG1(widget interface{}) (interface{}, error) {

    return SinkJob{Type: PLOT}, nil
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

