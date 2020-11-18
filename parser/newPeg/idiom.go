
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
	},
},
},
{
	name: "Sink",
	pos: position{line: 75, col: 1, offset: 1791},
	expr: &choiceExpr{
	pos: position{line: 75, col: 9, offset: 1799},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 75, col: 9, offset: 1799},
	name: "StdOutG",
},
&ruleRefExpr{
	pos: position{line: 75, col: 17, offset: 1807},
	name: "SinkInto",
},
&ruleRefExpr{
	pos: position{line: 75, col: 26, offset: 1816},
	name: "PlotG",
},
	},
},
},
{
	name: "FileJobG",
	pos: position{line: 79, col: 1, offset: 1842},
	expr: &actionExpr{
	pos: position{line: 79, col: 13, offset: 1854},
	run: (*parser).callonFileJobG1,
	expr: &seqExpr{
	pos: position{line: 79, col: 13, offset: 1854},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 79, col: 13, offset: 1854},
	name: "_",
},
&litMatcher{
	pos: position{line: 79, col: 15, offset: 1856},
	val: "file(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 79, col: 23, offset: 1864},
	name: "_",
},
&labeledExpr{
	pos: position{line: 79, col: 25, offset: 1866},
	label: "filename",
	expr: &ruleRefExpr{
	pos: position{line: 79, col: 34, offset: 1875},
	name: "FileName",
},
},
&ruleRefExpr{
	pos: position{line: 79, col: 43, offset: 1884},
	name: "_",
},
&litMatcher{
	pos: position{line: 79, col: 45, offset: 1886},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FakeJobG",
	pos: position{line: 86, col: 1, offset: 2101},
	expr: &actionExpr{
	pos: position{line: 86, col: 13, offset: 2113},
	run: (*parser).callonFakeJobG1,
	expr: &seqExpr{
	pos: position{line: 86, col: 13, offset: 2113},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 86, col: 13, offset: 2113},
	name: "_",
},
&litMatcher{
	pos: position{line: 86, col: 15, offset: 2115},
	val: "fake(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 86, col: 23, offset: 2123},
	name: "_",
},
&labeledExpr{
	pos: position{line: 86, col: 25, offset: 2125},
	label: "numData",
	expr: &ruleRefExpr{
	pos: position{line: 86, col: 33, offset: 2133},
	name: "PositiveNumber",
},
},
&ruleRefExpr{
	pos: position{line: 86, col: 48, offset: 2148},
	name: "_",
},
&litMatcher{
	pos: position{line: 86, col: 50, offset: 2150},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchJobG",
	pos: position{line: 93, col: 1, offset: 2349},
	expr: &actionExpr{
	pos: position{line: 93, col: 15, offset: 2363},
	run: (*parser).callonBranchJobG1,
	expr: &seqExpr{
	pos: position{line: 93, col: 15, offset: 2363},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 93, col: 15, offset: 2363},
	val: "branch",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 93, col: 23, offset: 2371},
	name: "_",
},
&litMatcher{
	pos: position{line: 93, col: 24, offset: 2372},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 93, col: 28, offset: 2376},
	name: "_",
},
&labeledExpr{
	pos: position{line: 93, col: 30, offset: 2378},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 93, col: 33, offset: 2381},
	name: "IDENTIFIER",
},
},
&litMatcher{
	pos: position{line: 93, col: 44, offset: 2392},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SecSource",
	pos: position{line: 100, col: 1, offset: 2608},
	expr: &actionExpr{
	pos: position{line: 100, col: 14, offset: 2621},
	run: (*parser).callonSecSource1,
	expr: &seqExpr{
	pos: position{line: 100, col: 14, offset: 2621},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 100, col: 14, offset: 2621},
	label: "src1",
	expr: &ruleRefExpr{
	pos: position{line: 100, col: 19, offset: 2626},
	name: "IDENTIFIER",
},
},
&ruleRefExpr{
	pos: position{line: 100, col: 30, offset: 2637},
	name: "_",
},
&litMatcher{
	pos: position{line: 100, col: 32, offset: 2639},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 100, col: 35, offset: 2642},
	name: "_",
},
&labeledExpr{
	pos: position{line: 100, col: 37, offset: 2644},
	label: "src2",
	expr: &ruleRefExpr{
	pos: position{line: 100, col: 42, offset: 2649},
	name: "IDENTIFIER",
},
},
	},
},
},
},
{
	name: "AggJobG",
	pos: position{line: 117, col: 1, offset: 3068},
	expr: &actionExpr{
	pos: position{line: 117, col: 12, offset: 3079},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 117, col: 12, offset: 3079},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 117, col: 12, offset: 3079},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 117, col: 18, offset: 3085},
	name: "_",
},
&labeledExpr{
	pos: position{line: 117, col: 20, offset: 3087},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 117, col: 27, offset: 3094},
	expr: &ruleRefExpr{
	pos: position{line: 117, col: 27, offset: 3094},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 117, col: 39, offset: 3106},
	name: "_",
},
&labeledExpr{
	pos: position{line: 117, col: 41, offset: 3108},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 117, col: 45, offset: 3112},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 117, col: 53, offset: 3120},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 117, col: 58, offset: 3125},
	expr: &seqExpr{
	pos: position{line: 117, col: 59, offset: 3126},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 117, col: 59, offset: 3126},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 117, col: 63, offset: 3130},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 117, col: 65, offset: 3132},
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
	pos: position{line: 132, col: 1, offset: 3507},
	expr: &choiceExpr{
	pos: position{line: 132, col: 12, offset: 3518},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 132, col: 12, offset: 3518},
	name: "CountAgg",
},
&ruleRefExpr{
	pos: position{line: 132, col: 23, offset: 3529},
	name: "MaxAgg",
},
&ruleRefExpr{
	pos: position{line: 132, col: 32, offset: 3538},
	name: "MinAgg",
},
&ruleRefExpr{
	pos: position{line: 132, col: 41, offset: 3547},
	name: "AvgAgg",
},
&ruleRefExpr{
	pos: position{line: 132, col: 50, offset: 3556},
	name: "SumAgg",
},
&ruleRefExpr{
	pos: position{line: 132, col: 59, offset: 3565},
	name: "ModeAgg",
},
&ruleRefExpr{
	pos: position{line: 133, col: 13, offset: 3586},
	name: "VarAgg",
},
&ruleRefExpr{
	pos: position{line: 133, col: 22, offset: 3595},
	name: "DistinctCountAgg",
},
&ruleRefExpr{
	pos: position{line: 133, col: 41, offset: 3614},
	name: "QuantileAgg",
},
&ruleRefExpr{
	pos: position{line: 133, col: 55, offset: 3628},
	name: "MedianAgg",
},
&ruleRefExpr{
	pos: position{line: 133, col: 67, offset: 3640},
	name: "QuartileAgg",
},
&ruleRefExpr{
	pos: position{line: 134, col: 13, offset: 3665},
	name: "CovAgg",
},
&ruleRefExpr{
	pos: position{line: 134, col: 22, offset: 3674},
	name: "CorrAgg",
},
&ruleRefExpr{
	pos: position{line: 134, col: 32, offset: 3684},
	name: "PValueAgg",
},
	},
},
},
{
	name: "CountAgg",
	pos: position{line: 137, col: 1, offset: 3728},
	expr: &actionExpr{
	pos: position{line: 137, col: 13, offset: 3740},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 137, col: 13, offset: 3740},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 137, col: 13, offset: 3740},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 137, col: 19, offset: 3746},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 137, col: 28, offset: 3755},
	name: "_",
},
&litMatcher{
	pos: position{line: 137, col: 30, offset: 3757},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 137, col: 39, offset: 3766},
	name: "_",
},
&labeledExpr{
	pos: position{line: 137, col: 41, offset: 3768},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 137, col: 46, offset: 3773},
	expr: &ruleRefExpr{
	pos: position{line: 137, col: 46, offset: 3773},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 137, col: 59, offset: 3786},
	name: "_",
},
&litMatcher{
	pos: position{line: 137, col: 61, offset: 3788},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MaxAgg",
	pos: position{line: 144, col: 1, offset: 3940},
	expr: &actionExpr{
	pos: position{line: 144, col: 11, offset: 3950},
	run: (*parser).callonMaxAgg1,
	expr: &seqExpr{
	pos: position{line: 144, col: 11, offset: 3950},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 144, col: 11, offset: 3950},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 144, col: 17, offset: 3956},
	expr: &ruleRefExpr{
	pos: position{line: 144, col: 17, offset: 3956},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 144, col: 27, offset: 3966},
	name: "_",
},
&litMatcher{
	pos: position{line: 144, col: 29, offset: 3968},
	val: "max(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 144, col: 36, offset: 3975},
	name: "_",
},
&labeledExpr{
	pos: position{line: 144, col: 38, offset: 3977},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 144, col: 44, offset: 3983},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 144, col: 53, offset: 3992},
	name: "_",
},
&labeledExpr{
	pos: position{line: 144, col: 55, offset: 3994},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 144, col: 60, offset: 3999},
	expr: &seqExpr{
	pos: position{line: 144, col: 61, offset: 4000},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 144, col: 61, offset: 4000},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 144, col: 65, offset: 4004},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 144, col: 67, offset: 4006},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 144, col: 81, offset: 4020},
	name: "_",
},
&litMatcher{
	pos: position{line: 144, col: 83, offset: 4022},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MinAgg",
	pos: position{line: 166, col: 1, offset: 4483},
	expr: &actionExpr{
	pos: position{line: 166, col: 11, offset: 4493},
	run: (*parser).callonMinAgg1,
	expr: &seqExpr{
	pos: position{line: 166, col: 11, offset: 4493},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 166, col: 11, offset: 4493},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 166, col: 17, offset: 4499},
	expr: &ruleRefExpr{
	pos: position{line: 166, col: 17, offset: 4499},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 166, col: 27, offset: 4509},
	name: "_",
},
&litMatcher{
	pos: position{line: 166, col: 29, offset: 4511},
	val: "min(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 166, col: 36, offset: 4518},
	name: "_",
},
&labeledExpr{
	pos: position{line: 166, col: 38, offset: 4520},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 166, col: 44, offset: 4526},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 166, col: 53, offset: 4535},
	name: "_",
},
&labeledExpr{
	pos: position{line: 166, col: 55, offset: 4537},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 166, col: 60, offset: 4542},
	expr: &seqExpr{
	pos: position{line: 166, col: 61, offset: 4543},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 166, col: 61, offset: 4543},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 166, col: 65, offset: 4547},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 166, col: 67, offset: 4549},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 166, col: 81, offset: 4563},
	name: "_",
},
&litMatcher{
	pos: position{line: 166, col: 83, offset: 4565},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AvgAgg",
	pos: position{line: 188, col: 1, offset: 5026},
	expr: &actionExpr{
	pos: position{line: 188, col: 11, offset: 5036},
	run: (*parser).callonAvgAgg1,
	expr: &seqExpr{
	pos: position{line: 188, col: 11, offset: 5036},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 188, col: 11, offset: 5036},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 188, col: 17, offset: 5042},
	expr: &ruleRefExpr{
	pos: position{line: 188, col: 17, offset: 5042},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 188, col: 27, offset: 5052},
	name: "_",
},
&litMatcher{
	pos: position{line: 188, col: 29, offset: 5054},
	val: "avg(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 188, col: 36, offset: 5061},
	name: "_",
},
&labeledExpr{
	pos: position{line: 188, col: 38, offset: 5063},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 188, col: 44, offset: 5069},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 188, col: 53, offset: 5078},
	name: "_",
},
&labeledExpr{
	pos: position{line: 188, col: 55, offset: 5080},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 188, col: 60, offset: 5085},
	expr: &seqExpr{
	pos: position{line: 188, col: 61, offset: 5086},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 188, col: 61, offset: 5086},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 188, col: 65, offset: 5090},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 188, col: 67, offset: 5092},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 188, col: 81, offset: 5106},
	name: "_",
},
&litMatcher{
	pos: position{line: 188, col: 83, offset: 5108},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SumAgg",
	pos: position{line: 210, col: 1, offset: 5569},
	expr: &actionExpr{
	pos: position{line: 210, col: 11, offset: 5579},
	run: (*parser).callonSumAgg1,
	expr: &seqExpr{
	pos: position{line: 210, col: 11, offset: 5579},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 210, col: 11, offset: 5579},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 210, col: 17, offset: 5585},
	expr: &ruleRefExpr{
	pos: position{line: 210, col: 17, offset: 5585},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 210, col: 27, offset: 5595},
	name: "_",
},
&litMatcher{
	pos: position{line: 210, col: 29, offset: 5597},
	val: "sum(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 210, col: 36, offset: 5604},
	name: "_",
},
&labeledExpr{
	pos: position{line: 210, col: 38, offset: 5606},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 210, col: 44, offset: 5612},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 210, col: 53, offset: 5621},
	name: "_",
},
&labeledExpr{
	pos: position{line: 210, col: 55, offset: 5623},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 210, col: 60, offset: 5628},
	expr: &seqExpr{
	pos: position{line: 210, col: 61, offset: 5629},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 210, col: 61, offset: 5629},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 210, col: 65, offset: 5633},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 210, col: 67, offset: 5635},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 210, col: 81, offset: 5649},
	name: "_",
},
&litMatcher{
	pos: position{line: 210, col: 83, offset: 5651},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarAgg",
	pos: position{line: 232, col: 1, offset: 6112},
	expr: &actionExpr{
	pos: position{line: 232, col: 11, offset: 6122},
	run: (*parser).callonVarAgg1,
	expr: &seqExpr{
	pos: position{line: 232, col: 11, offset: 6122},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 232, col: 11, offset: 6122},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 232, col: 17, offset: 6128},
	expr: &ruleRefExpr{
	pos: position{line: 232, col: 17, offset: 6128},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 232, col: 27, offset: 6138},
	name: "_",
},
&litMatcher{
	pos: position{line: 232, col: 29, offset: 6140},
	val: "var(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 232, col: 36, offset: 6147},
	name: "_",
},
&labeledExpr{
	pos: position{line: 232, col: 38, offset: 6149},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 232, col: 44, offset: 6155},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 232, col: 53, offset: 6164},
	name: "_",
},
&labeledExpr{
	pos: position{line: 232, col: 55, offset: 6166},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 232, col: 60, offset: 6171},
	expr: &seqExpr{
	pos: position{line: 232, col: 61, offset: 6172},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 232, col: 61, offset: 6172},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 232, col: 65, offset: 6176},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 232, col: 67, offset: 6178},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 232, col: 81, offset: 6192},
	name: "_",
},
&litMatcher{
	pos: position{line: 232, col: 83, offset: 6194},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DistinctCountAgg",
	pos: position{line: 254, col: 1, offset: 6660},
	expr: &actionExpr{
	pos: position{line: 254, col: 21, offset: 6680},
	run: (*parser).callonDistinctCountAgg1,
	expr: &seqExpr{
	pos: position{line: 254, col: 21, offset: 6680},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 254, col: 21, offset: 6680},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 254, col: 27, offset: 6686},
	expr: &ruleRefExpr{
	pos: position{line: 254, col: 27, offset: 6686},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 254, col: 37, offset: 6696},
	name: "_",
},
&litMatcher{
	pos: position{line: 254, col: 39, offset: 6698},
	val: "dcount(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 254, col: 49, offset: 6708},
	name: "_",
},
&labeledExpr{
	pos: position{line: 254, col: 51, offset: 6710},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 254, col: 57, offset: 6716},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 254, col: 66, offset: 6725},
	name: "_",
},
&labeledExpr{
	pos: position{line: 254, col: 68, offset: 6727},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 254, col: 73, offset: 6732},
	expr: &seqExpr{
	pos: position{line: 254, col: 74, offset: 6733},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 254, col: 74, offset: 6733},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 254, col: 78, offset: 6737},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 254, col: 80, offset: 6739},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 254, col: 94, offset: 6753},
	name: "_",
},
&litMatcher{
	pos: position{line: 254, col: 96, offset: 6755},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "ModeAgg",
	pos: position{line: 276, col: 1, offset: 7226},
	expr: &actionExpr{
	pos: position{line: 276, col: 12, offset: 7237},
	run: (*parser).callonModeAgg1,
	expr: &seqExpr{
	pos: position{line: 276, col: 12, offset: 7237},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 276, col: 12, offset: 7237},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 276, col: 18, offset: 7243},
	expr: &ruleRefExpr{
	pos: position{line: 276, col: 18, offset: 7243},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 276, col: 28, offset: 7253},
	name: "_",
},
&litMatcher{
	pos: position{line: 276, col: 30, offset: 7255},
	val: "mode(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 276, col: 38, offset: 7263},
	name: "_",
},
&labeledExpr{
	pos: position{line: 276, col: 40, offset: 7265},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 276, col: 46, offset: 7271},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 276, col: 55, offset: 7280},
	name: "_",
},
&labeledExpr{
	pos: position{line: 276, col: 57, offset: 7282},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 276, col: 62, offset: 7287},
	expr: &seqExpr{
	pos: position{line: 276, col: 63, offset: 7288},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 276, col: 63, offset: 7288},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 276, col: 67, offset: 7292},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 276, col: 69, offset: 7294},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 276, col: 83, offset: 7308},
	name: "_",
},
&litMatcher{
	pos: position{line: 276, col: 85, offset: 7310},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuantileAgg",
	pos: position{line: 298, col: 1, offset: 7772},
	expr: &actionExpr{
	pos: position{line: 298, col: 16, offset: 7787},
	run: (*parser).callonQuantileAgg1,
	expr: &seqExpr{
	pos: position{line: 298, col: 16, offset: 7787},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 298, col: 16, offset: 7787},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 298, col: 22, offset: 7793},
	expr: &ruleRefExpr{
	pos: position{line: 298, col: 22, offset: 7793},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 298, col: 32, offset: 7803},
	name: "_",
},
&litMatcher{
	pos: position{line: 298, col: 34, offset: 7805},
	val: "quantile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 298, col: 46, offset: 7817},
	name: "_",
},
&labeledExpr{
	pos: position{line: 298, col: 48, offset: 7819},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 298, col: 54, offset: 7825},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 298, col: 63, offset: 7834},
	name: "_",
},
&labeledExpr{
	pos: position{line: 298, col: 66, offset: 7837},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 298, col: 73, offset: 7844},
	expr: &seqExpr{
	pos: position{line: 298, col: 74, offset: 7845},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 298, col: 74, offset: 7845},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 298, col: 78, offset: 7849},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 298, col: 80, offset: 7851},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 298, col: 91, offset: 7862},
	name: "_",
},
&litMatcher{
	pos: position{line: 298, col: 93, offset: 7864},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 298, col: 97, offset: 7868},
	name: "_",
},
&labeledExpr{
	pos: position{line: 298, col: 99, offset: 7870},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 298, col: 103, offset: 7874},
	name: "Float",
},
},
&ruleRefExpr{
	pos: position{line: 298, col: 109, offset: 7880},
	name: "_",
},
&labeledExpr{
	pos: position{line: 298, col: 111, offset: 7882},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 298, col: 116, offset: 7887},
	expr: &seqExpr{
	pos: position{line: 298, col: 117, offset: 7888},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 298, col: 117, offset: 7888},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 298, col: 121, offset: 7892},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 298, col: 123, offset: 7894},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 298, col: 137, offset: 7908},
	name: "_",
},
&litMatcher{
	pos: position{line: 298, col: 139, offset: 7910},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MedianAgg",
	pos: position{line: 328, col: 1, offset: 8555},
	expr: &actionExpr{
	pos: position{line: 328, col: 14, offset: 8568},
	run: (*parser).callonMedianAgg1,
	expr: &seqExpr{
	pos: position{line: 328, col: 14, offset: 8568},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 328, col: 14, offset: 8568},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 328, col: 20, offset: 8574},
	expr: &ruleRefExpr{
	pos: position{line: 328, col: 20, offset: 8574},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 328, col: 30, offset: 8584},
	name: "_",
},
&litMatcher{
	pos: position{line: 328, col: 32, offset: 8586},
	val: "median(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 328, col: 42, offset: 8596},
	name: "_",
},
&labeledExpr{
	pos: position{line: 328, col: 44, offset: 8598},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 328, col: 50, offset: 8604},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 328, col: 59, offset: 8613},
	name: "_",
},
&labeledExpr{
	pos: position{line: 328, col: 62, offset: 8616},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 328, col: 69, offset: 8623},
	expr: &seqExpr{
	pos: position{line: 328, col: 70, offset: 8624},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 328, col: 70, offset: 8624},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 328, col: 74, offset: 8628},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 328, col: 76, offset: 8630},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 328, col: 87, offset: 8641},
	name: "_",
},
&labeledExpr{
	pos: position{line: 328, col: 89, offset: 8643},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 328, col: 94, offset: 8648},
	expr: &seqExpr{
	pos: position{line: 328, col: 95, offset: 8649},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 328, col: 95, offset: 8649},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 328, col: 99, offset: 8653},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 328, col: 101, offset: 8655},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 328, col: 115, offset: 8669},
	name: "_",
},
&litMatcher{
	pos: position{line: 328, col: 117, offset: 8671},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuartileAgg",
	pos: position{line: 360, col: 1, offset: 9348},
	expr: &actionExpr{
	pos: position{line: 360, col: 16, offset: 9363},
	run: (*parser).callonQuartileAgg1,
	expr: &seqExpr{
	pos: position{line: 360, col: 16, offset: 9363},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 360, col: 16, offset: 9363},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 360, col: 22, offset: 9369},
	expr: &ruleRefExpr{
	pos: position{line: 360, col: 22, offset: 9369},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 360, col: 32, offset: 9379},
	name: "_",
},
&litMatcher{
	pos: position{line: 360, col: 34, offset: 9381},
	val: "quartile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 360, col: 46, offset: 9393},
	name: "_",
},
&labeledExpr{
	pos: position{line: 360, col: 48, offset: 9395},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 360, col: 54, offset: 9401},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 360, col: 63, offset: 9410},
	name: "_",
},
&labeledExpr{
	pos: position{line: 360, col: 66, offset: 9413},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 360, col: 73, offset: 9420},
	expr: &seqExpr{
	pos: position{line: 360, col: 74, offset: 9421},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 360, col: 74, offset: 9421},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 360, col: 78, offset: 9425},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 360, col: 80, offset: 9427},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 360, col: 91, offset: 9438},
	name: "_",
},
&litMatcher{
	pos: position{line: 360, col: 93, offset: 9440},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 360, col: 97, offset: 9444},
	name: "_",
},
&labeledExpr{
	pos: position{line: 360, col: 99, offset: 9446},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 360, col: 103, offset: 9450},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 360, col: 110, offset: 9457},
	name: "_",
},
&labeledExpr{
	pos: position{line: 360, col: 112, offset: 9459},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 360, col: 117, offset: 9464},
	expr: &seqExpr{
	pos: position{line: 360, col: 118, offset: 9465},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 360, col: 118, offset: 9465},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 360, col: 122, offset: 9469},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 360, col: 124, offset: 9471},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 360, col: 138, offset: 9485},
	name: "_",
},
&litMatcher{
	pos: position{line: 360, col: 140, offset: 9487},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CovAgg",
	pos: position{line: 406, col: 1, offset: 10514},
	expr: &actionExpr{
	pos: position{line: 406, col: 11, offset: 10524},
	run: (*parser).callonCovAgg1,
	expr: &seqExpr{
	pos: position{line: 406, col: 11, offset: 10524},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 406, col: 11, offset: 10524},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 406, col: 17, offset: 10530},
	expr: &ruleRefExpr{
	pos: position{line: 406, col: 17, offset: 10530},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 406, col: 27, offset: 10540},
	name: "_",
},
&litMatcher{
	pos: position{line: 406, col: 29, offset: 10542},
	val: "cov(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 406, col: 36, offset: 10549},
	name: "_",
},
&labeledExpr{
	pos: position{line: 406, col: 38, offset: 10551},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 406, col: 45, offset: 10558},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 406, col: 54, offset: 10567},
	name: "_",
},
&litMatcher{
	pos: position{line: 406, col: 56, offset: 10569},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 406, col: 60, offset: 10573},
	name: "_",
},
&labeledExpr{
	pos: position{line: 406, col: 62, offset: 10575},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 406, col: 69, offset: 10582},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 406, col: 78, offset: 10591},
	name: "_",
},
&labeledExpr{
	pos: position{line: 406, col: 80, offset: 10593},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 406, col: 85, offset: 10598},
	expr: &seqExpr{
	pos: position{line: 406, col: 86, offset: 10599},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 406, col: 86, offset: 10599},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 406, col: 90, offset: 10603},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 406, col: 92, offset: 10605},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 406, col: 106, offset: 10619},
	name: "_",
},
&litMatcher{
	pos: position{line: 406, col: 108, offset: 10621},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CorrAgg",
	pos: position{line: 429, col: 1, offset: 11146},
	expr: &actionExpr{
	pos: position{line: 429, col: 12, offset: 11157},
	run: (*parser).callonCorrAgg1,
	expr: &seqExpr{
	pos: position{line: 429, col: 12, offset: 11157},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 429, col: 12, offset: 11157},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 429, col: 18, offset: 11163},
	expr: &ruleRefExpr{
	pos: position{line: 429, col: 18, offset: 11163},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 429, col: 28, offset: 11173},
	name: "_",
},
&litMatcher{
	pos: position{line: 429, col: 30, offset: 11175},
	val: "correlation(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 429, col: 45, offset: 11190},
	name: "_",
},
&labeledExpr{
	pos: position{line: 429, col: 47, offset: 11192},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 429, col: 54, offset: 11199},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 429, col: 63, offset: 11208},
	name: "_",
},
&litMatcher{
	pos: position{line: 429, col: 65, offset: 11210},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 429, col: 69, offset: 11214},
	name: "_",
},
&labeledExpr{
	pos: position{line: 429, col: 71, offset: 11216},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 429, col: 78, offset: 11223},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 429, col: 87, offset: 11232},
	name: "_",
},
&labeledExpr{
	pos: position{line: 429, col: 89, offset: 11234},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 429, col: 94, offset: 11239},
	expr: &seqExpr{
	pos: position{line: 429, col: 95, offset: 11240},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 429, col: 95, offset: 11240},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 429, col: 99, offset: 11244},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 429, col: 101, offset: 11246},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 429, col: 115, offset: 11260},
	name: "_",
},
&litMatcher{
	pos: position{line: 429, col: 117, offset: 11262},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PValueAgg",
	pos: position{line: 452, col: 1, offset: 11788},
	expr: &actionExpr{
	pos: position{line: 452, col: 14, offset: 11801},
	run: (*parser).callonPValueAgg1,
	expr: &seqExpr{
	pos: position{line: 452, col: 14, offset: 11801},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 452, col: 14, offset: 11801},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 452, col: 20, offset: 11807},
	expr: &ruleRefExpr{
	pos: position{line: 452, col: 20, offset: 11807},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 452, col: 30, offset: 11817},
	name: "_",
},
&litMatcher{
	pos: position{line: 452, col: 32, offset: 11819},
	val: "pvalue(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 452, col: 42, offset: 11829},
	name: "_",
},
&labeledExpr{
	pos: position{line: 452, col: 44, offset: 11831},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 452, col: 51, offset: 11838},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 452, col: 60, offset: 11847},
	name: "_",
},
&litMatcher{
	pos: position{line: 452, col: 62, offset: 11849},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 452, col: 66, offset: 11853},
	name: "_",
},
&labeledExpr{
	pos: position{line: 452, col: 68, offset: 11855},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 452, col: 75, offset: 11862},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 452, col: 84, offset: 11871},
	name: "_",
},
&labeledExpr{
	pos: position{line: 452, col: 86, offset: 11873},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 452, col: 91, offset: 11878},
	expr: &seqExpr{
	pos: position{line: 452, col: 92, offset: 11879},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 452, col: 92, offset: 11879},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 452, col: 96, offset: 11883},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 452, col: 98, offset: 11885},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 452, col: 112, offset: 11899},
	name: "_",
},
&litMatcher{
	pos: position{line: 452, col: 114, offset: 11901},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggGroupBy",
	pos: position{line: 477, col: 1, offset: 12426},
	expr: &actionExpr{
	pos: position{line: 477, col: 15, offset: 12440},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 477, col: 15, offset: 12440},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 477, col: 15, offset: 12440},
	name: "_",
},
&litMatcher{
	pos: position{line: 477, col: 17, offset: 12442},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 477, col: 22, offset: 12447},
	name: "_",
},
&labeledExpr{
	pos: position{line: 477, col: 24, offset: 12449},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 477, col: 30, offset: 12455},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 477, col: 39, offset: 12464},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 477, col: 44, offset: 12469},
	expr: &seqExpr{
	pos: position{line: 477, col: 45, offset: 12470},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 477, col: 45, offset: 12470},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 477, col: 49, offset: 12474},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 477, col: 51, offset: 12476},
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
	pos: position{line: 492, col: 1, offset: 12827},
	expr: &actionExpr{
	pos: position{line: 492, col: 11, offset: 12837},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 492, col: 11, offset: 12837},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 492, col: 16, offset: 12842},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 492, col: 16, offset: 12842},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 492, col: 27, offset: 12853},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 492, col: 38, offset: 12864},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 492, col: 49, offset: 12875},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 492, col: 58, offset: 12884},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 492, col: 67, offset: 12893},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 498, col: 1, offset: 13012},
	expr: &actionExpr{
	pos: position{line: 498, col: 13, offset: 13024},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 498, col: 13, offset: 13024},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 498, col: 13, offset: 13024},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 498, col: 23, offset: 13034},
	name: "_",
},
&labeledExpr{
	pos: position{line: 498, col: 25, offset: 13036},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 498, col: 28, offset: 13039},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 498, col: 42, offset: 13053},
	name: "_",
},
&litMatcher{
	pos: position{line: 498, col: 44, offset: 13055},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 505, col: 1, offset: 13250},
	expr: &actionExpr{
	pos: position{line: 505, col: 18, offset: 13267},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 505, col: 18, offset: 13267},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 505, col: 18, offset: 13267},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 505, col: 23, offset: 13272},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 505, col: 28, offset: 13277},
	name: "_",
},
&labeledExpr{
	pos: position{line: 505, col: 30, offset: 13279},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 505, col: 33, offset: 13282},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 505, col: 45, offset: 13294},
	name: "_",
},
&labeledExpr{
	pos: position{line: 505, col: 47, offset: 13296},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 505, col: 54, offset: 13303},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 509, col: 1, offset: 13354},
	expr: &ruleRefExpr{
	pos: position{line: 509, col: 16, offset: 13369},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 512, col: 1, offset: 13403},
	expr: &actionExpr{
	pos: position{line: 512, col: 13, offset: 13415},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 512, col: 13, offset: 13415},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 512, col: 13, offset: 13415},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 512, col: 23, offset: 13425},
	name: "_",
},
&labeledExpr{
	pos: position{line: 512, col: 25, offset: 13427},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 512, col: 31, offset: 13433},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 512, col: 40, offset: 13442},
	name: "_",
},
&labeledExpr{
	pos: position{line: 512, col: 42, offset: 13444},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 512, col: 47, offset: 13449},
	expr: &seqExpr{
	pos: position{line: 512, col: 48, offset: 13450},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 512, col: 48, offset: 13450},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 512, col: 52, offset: 13454},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 512, col: 54, offset: 13456},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 512, col: 65, offset: 13467},
	name: "_",
},
&litMatcher{
	pos: position{line: 512, col: 67, offset: 13469},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 527, col: 1, offset: 13806},
	expr: &actionExpr{
	pos: position{line: 527, col: 11, offset: 13816},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 527, col: 11, offset: 13816},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 527, col: 11, offset: 13816},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 527, col: 18, offset: 13823},
	name: "_",
},
&labeledExpr{
	pos: position{line: 527, col: 20, offset: 13825},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 527, col: 25, offset: 13830},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 527, col: 34, offset: 13839},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 527, col: 38, offset: 13843},
	name: "_",
},
&labeledExpr{
	pos: position{line: 527, col: 40, offset: 13845},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 527, col: 44, offset: 13849},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 527, col: 51, offset: 13856},
	name: "_",
},
&litMatcher{
	pos: position{line: 527, col: 53, offset: 13858},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 539, col: 1, offset: 14080},
	expr: &actionExpr{
	pos: position{line: 539, col: 11, offset: 14090},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 539, col: 11, offset: 14090},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 539, col: 11, offset: 14090},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 539, col: 20, offset: 14099},
	name: "_",
},
&litMatcher{
	pos: position{line: 539, col: 22, offset: 14101},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 539, col: 27, offset: 14106},
	name: "_",
},
&labeledExpr{
	pos: position{line: 539, col: 29, offset: 14108},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 539, col: 35, offset: 14114},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 545, col: 1, offset: 14220},
	expr: &actionExpr{
	pos: position{line: 545, col: 12, offset: 14231},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 545, col: 12, offset: 14231},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 545, col: 12, offset: 14231},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 545, col: 21, offset: 14240},
	name: "_",
},
&labeledExpr{
	pos: position{line: 545, col: 23, offset: 14242},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 545, col: 27, offset: 14246},
	expr: &ruleRefExpr{
	pos: position{line: 545, col: 27, offset: 14246},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 545, col: 35, offset: 14254},
	name: "_",
},
&litMatcher{
	pos: position{line: 545, col: 37, offset: 14256},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 555, col: 1, offset: 14401},
	expr: &actionExpr{
	pos: position{line: 555, col: 13, offset: 14413},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 555, col: 13, offset: 14413},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 555, col: 13, offset: 14413},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 555, col: 23, offset: 14423},
	name: "_",
},
&labeledExpr{
	pos: position{line: 555, col: 25, offset: 14425},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 555, col: 30, offset: 14430},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 555, col: 42, offset: 14442},
	name: "_",
},
&litMatcher{
	pos: position{line: 555, col: 44, offset: 14444},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 565, col: 1, offset: 14732},
	expr: &actionExpr{
	pos: position{line: 565, col: 16, offset: 14747},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 565, col: 16, offset: 14747},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 565, col: 16, offset: 14747},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 565, col: 22, offset: 14753},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 565, col: 35, offset: 14766},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 565, col: 40, offset: 14771},
	expr: &seqExpr{
	pos: position{line: 565, col: 41, offset: 14772},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 565, col: 41, offset: 14772},
	name: "_",
},
&labeledExpr{
	pos: position{line: 565, col: 43, offset: 14774},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 565, col: 46, offset: 14777},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 565, col: 58, offset: 14789},
	name: "_",
},
&labeledExpr{
	pos: position{line: 565, col: 60, offset: 14791},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 565, col: 63, offset: 14794},
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
	pos: position{line: 570, col: 1, offset: 14952},
	expr: &choiceExpr{
	pos: position{line: 570, col: 17, offset: 14968},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 570, col: 17, offset: 14968},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 570, col: 17, offset: 14968},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 570, col: 17, offset: 14968},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 570, col: 21, offset: 14972},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 570, col: 24, offset: 14975},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 570, col: 36, offset: 14987},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 572, col: 5, offset: 15018},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 572, col: 5, offset: 15018},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 572, col: 8, offset: 15021},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 576, col: 1, offset: 15063},
	expr: &choiceExpr{
	pos: position{line: 576, col: 16, offset: 15078},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 576, col: 16, offset: 15078},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 576, col: 24, offset: 15086},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 576, col: 24, offset: 15086},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 583, col: 1, offset: 15265},
	expr: &actionExpr{
	pos: position{line: 583, col: 18, offset: 15282},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 583, col: 18, offset: 15282},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 583, col: 18, offset: 15282},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 583, col: 23, offset: 15287},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 583, col: 28, offset: 15292},
	name: "_",
},
&labeledExpr{
	pos: position{line: 583, col: 30, offset: 15294},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 583, col: 33, offset: 15297},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 583, col: 45, offset: 15309},
	name: "_",
},
&labeledExpr{
	pos: position{line: 583, col: 47, offset: 15311},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 583, col: 53, offset: 15317},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 583, col: 59, offset: 15323},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 587, col: 1, offset: 15397},
	expr: &choiceExpr{
	pos: position{line: 587, col: 16, offset: 15412},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 587, col: 16, offset: 15412},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 587, col: 23, offset: 15419},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 587, col: 30, offset: 15426},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 587, col: 36, offset: 15432},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 587, col: 42, offset: 15438},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 587, col: 49, offset: 15445},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 587, col: 49, offset: 15445},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 596, col: 1, offset: 15555},
	expr: &actionExpr{
	pos: position{line: 596, col: 12, offset: 15566},
	run: (*parser).callonStdOutG1,
	expr: &seqExpr{
	pos: position{line: 596, col: 12, offset: 15566},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 596, col: 12, offset: 15566},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 596, col: 22, offset: 15576},
	name: "_",
},
&labeledExpr{
	pos: position{line: 596, col: 24, offset: 15578},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 596, col: 31, offset: 15585},
	expr: &ruleRefExpr{
	pos: position{line: 596, col: 32, offset: 15586},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 596, col: 48, offset: 15602},
	name: "_",
},
&litMatcher{
	pos: position{line: 596, col: 50, offset: 15604},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SinkInto",
	pos: position{line: 610, col: 1, offset: 15883},
	expr: &actionExpr{
	pos: position{line: 610, col: 12, offset: 15894},
	run: (*parser).callonSinkInto1,
	expr: &seqExpr{
	pos: position{line: 610, col: 12, offset: 15894},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 610, col: 12, offset: 15894},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 610, col: 18, offset: 15900},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 610, col: 27, offset: 15909},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 610, col: 33, offset: 15915},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 610, col: 33, offset: 15915},
	name: "_",
},
&litMatcher{
	pos: position{line: 610, col: 35, offset: 15917},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 610, col: 39, offset: 15921},
	name: "_",
},
&litMatcher{
	pos: position{line: 610, col: 41, offset: 15923},
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
	pos: position{line: 618, col: 1, offset: 16079},
	expr: &actionExpr{
	pos: position{line: 618, col: 15, offset: 16093},
	run: (*parser).callonBlackHoleG1,
	expr: &seqExpr{
	pos: position{line: 618, col: 15, offset: 16093},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 618, col: 15, offset: 16093},
	val: "blackhole(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 618, col: 28, offset: 16106},
	name: "_",
},
&labeledExpr{
	pos: position{line: 618, col: 30, offset: 16108},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 618, col: 37, offset: 16115},
	expr: &ruleRefExpr{
	pos: position{line: 618, col: 38, offset: 16116},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 618, col: 54, offset: 16132},
	name: "_",
},
&litMatcher{
	pos: position{line: 618, col: 56, offset: 16134},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PlotG",
	pos: position{line: 622, col: 1, offset: 16192},
	expr: &actionExpr{
	pos: position{line: 622, col: 10, offset: 16201},
	run: (*parser).callonPlotG1,
	expr: &seqExpr{
	pos: position{line: 622, col: 10, offset: 16201},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 622, col: 10, offset: 16201},
	val: "plot",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 622, col: 17, offset: 16208},
	name: "_",
},
&labeledExpr{
	pos: position{line: 622, col: 19, offset: 16210},
	label: "widget",
	expr: &ruleRefExpr{
	pos: position{line: 622, col: 27, offset: 16218},
	name: "BarWidget",
},
},
	},
},
},
},
{
	name: "BarWidget",
	pos: position{line: 626, col: 1, offset: 16318},
	expr: &actionExpr{
	pos: position{line: 626, col: 14, offset: 16331},
	run: (*parser).callonBarWidget1,
	expr: &seqExpr{
	pos: position{line: 626, col: 14, offset: 16331},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 626, col: 14, offset: 16331},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 626, col: 20, offset: 16337},
	expr: &ruleRefExpr{
	pos: position{line: 626, col: 20, offset: 16337},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 626, col: 30, offset: 16347},
	name: "_",
},
&litMatcher{
	pos: position{line: 626, col: 32, offset: 16349},
	val: "bar(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 626, col: 39, offset: 16356},
	name: "_",
},
&labeledExpr{
	pos: position{line: 626, col: 41, offset: 16358},
	label: "xField",
	expr: &ruleRefExpr{
	pos: position{line: 626, col: 48, offset: 16365},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 626, col: 57, offset: 16374},
	name: "_",
},
&litMatcher{
	pos: position{line: 626, col: 59, offset: 16376},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 626, col: 63, offset: 16380},
	name: "_",
},
&labeledExpr{
	pos: position{line: 626, col: 65, offset: 16382},
	label: "yField",
	expr: &ruleRefExpr{
	pos: position{line: 626, col: 72, offset: 16389},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 626, col: 81, offset: 16398},
	name: "_",
},
&labeledExpr{
	pos: position{line: 626, col: 83, offset: 16400},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 626, col: 88, offset: 16405},
	expr: &seqExpr{
	pos: position{line: 626, col: 89, offset: 16406},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 626, col: 89, offset: 16406},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 626, col: 93, offset: 16410},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 626, col: 95, offset: 16412},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 626, col: 102, offset: 16419},
	name: "_",
},
&zeroOrOneExpr{
	pos: position{line: 626, col: 104, offset: 16421},
	expr: &seqExpr{
	pos: position{line: 626, col: 105, offset: 16422},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 626, col: 105, offset: 16422},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 626, col: 109, offset: 16426},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 626, col: 111, offset: 16428},
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
	pos: position{line: 626, col: 122, offset: 16439},
	name: "_",
},
&litMatcher{
	pos: position{line: 626, col: 124, offset: 16441},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FileName",
	pos: position{line: 659, col: 1, offset: 17126},
	expr: &actionExpr{
	pos: position{line: 659, col: 13, offset: 17138},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 659, col: 13, offset: 17138},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 659, col: 13, offset: 17138},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 659, col: 22, offset: 17147},
	expr: &seqExpr{
	pos: position{line: 659, col: 23, offset: 17148},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 659, col: 23, offset: 17148},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 659, col: 27, offset: 17152},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 659, col: 31, offset: 17156},
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
	pos: position{line: 665, col: 1, offset: 17257},
	expr: &actionExpr{
	pos: position{line: 665, col: 10, offset: 17266},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 665, col: 10, offset: 17266},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 665, col: 10, offset: 17266},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 665, col: 16, offset: 17272},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 665, col: 23, offset: 17279},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 665, col: 28, offset: 17284},
	expr: &seqExpr{
	pos: position{line: 665, col: 29, offset: 17285},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 665, col: 29, offset: 17285},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 665, col: 31, offset: 17287},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 665, col: 36, offset: 17292},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 665, col: 38, offset: 17294},
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
	pos: position{line: 669, col: 1, offset: 17359},
	expr: &actionExpr{
	pos: position{line: 669, col: 11, offset: 17369},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 669, col: 11, offset: 17369},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 669, col: 16, offset: 17374},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 669, col: 16, offset: 17374},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 669, col: 24, offset: 17382},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 669, col: 33, offset: 17391},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 669, col: 48, offset: 17406},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 669, col: 59, offset: 17417},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 673, col: 1, offset: 17459},
	expr: &choiceExpr{
	pos: position{line: 673, col: 9, offset: 17467},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 673, col: 9, offset: 17467},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 673, col: 15, offset: 17473},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 673, col: 21, offset: 17479},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 673, col: 27, offset: 17485},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 673, col: 33, offset: 17491},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 673, col: 33, offset: 17491},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 677, col: 1, offset: 17539},
	expr: &choiceExpr{
	pos: position{line: 677, col: 12, offset: 17550},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 677, col: 12, offset: 17550},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 677, col: 12, offset: 17550},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 677, col: 12, offset: 17550},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 677, col: 16, offset: 17554},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 677, col: 19, offset: 17557},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 677, col: 25, offset: 17563},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 681, col: 5, offset: 17639},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 681, col: 5, offset: 17639},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 681, col: 8, offset: 17642},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 688, col: 1, offset: 17719},
	expr: &actionExpr{
	pos: position{line: 688, col: 9, offset: 17727},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 688, col: 9, offset: 17727},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 688, col: 9, offset: 17727},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 688, col: 15, offset: 17733},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 688, col: 21, offset: 17739},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 688, col: 26, offset: 17744},
	expr: &seqExpr{
	pos: position{line: 688, col: 27, offset: 17745},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 688, col: 27, offset: 17745},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 688, col: 29, offset: 17747},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 688, col: 34, offset: 17752},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 688, col: 36, offset: 17754},
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
	pos: position{line: 692, col: 1, offset: 17818},
	expr: &actionExpr{
	pos: position{line: 692, col: 10, offset: 17827},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 692, col: 10, offset: 17827},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 692, col: 15, offset: 17832},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 692, col: 15, offset: 17832},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 692, col: 23, offset: 17840},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 692, col: 32, offset: 17849},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 692, col: 47, offset: 17864},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 696, col: 1, offset: 17907},
	expr: &actionExpr{
	pos: position{line: 696, col: 13, offset: 17919},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 696, col: 13, offset: 17919},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 696, col: 13, offset: 17919},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 696, col: 20, offset: 17926},
	expr: &choiceExpr{
	pos: position{line: 696, col: 21, offset: 17927},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 696, col: 21, offset: 17927},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 696, col: 31, offset: 17937},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 696, col: 39, offset: 17945},
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
	pos: position{line: 700, col: 1, offset: 17995},
	expr: &actionExpr{
	pos: position{line: 700, col: 17, offset: 18011},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 700, col: 17, offset: 18011},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 700, col: 17, offset: 18011},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 700, col: 21, offset: 18015},
	expr: &choiceExpr{
	pos: position{line: 700, col: 23, offset: 18017},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 700, col: 23, offset: 18017},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 700, col: 23, offset: 18017},
	expr: &ruleRefExpr{
	pos: position{line: 700, col: 24, offset: 18018},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 700, col: 36, offset: 18030,
},
	},
},
&seqExpr{
	pos: position{line: 700, col: 40, offset: 18034},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 700, col: 40, offset: 18034},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 700, col: 45, offset: 18039},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 700, col: 63, offset: 18057},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 707, col: 1, offset: 18238},
	expr: &charClassMatcher{
	pos: position{line: 707, col: 16, offset: 18253},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 709, col: 1, offset: 18271},
	expr: &ruleRefExpr{
	pos: position{line: 709, col: 19, offset: 18289},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 711, col: 1, offset: 18309},
	expr: &charClassMatcher{
	pos: position{line: 711, col: 21, offset: 18329},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 713, col: 1, offset: 18344},
	expr: &actionExpr{
	pos: position{line: 713, col: 9, offset: 18352},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 713, col: 9, offset: 18352},
	expr: &choiceExpr{
	pos: position{line: 713, col: 10, offset: 18353},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 713, col: 10, offset: 18353},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 713, col: 19, offset: 18362},
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
	pos: position{line: 717, col: 1, offset: 18412},
	expr: &choiceExpr{
	pos: position{line: 717, col: 11, offset: 18422},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 717, col: 11, offset: 18422},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 717, col: 19, offset: 18430},
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
	pos: position{line: 719, col: 1, offset: 18439},
	expr: &actionExpr{
	pos: position{line: 719, col: 11, offset: 18449},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 719, col: 11, offset: 18449},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 719, col: 11, offset: 18449},
	expr: &litMatcher{
	pos: position{line: 719, col: 11, offset: 18449},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 719, col: 16, offset: 18454},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 723, col: 1, offset: 18527},
	expr: &actionExpr{
	pos: position{line: 723, col: 19, offset: 18545},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 723, col: 19, offset: 18545},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 727, col: 1, offset: 18618},
	expr: &actionExpr{
	pos: position{line: 727, col: 10, offset: 18627},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 727, col: 10, offset: 18627},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 727, col: 10, offset: 18627},
	expr: &litMatcher{
	pos: position{line: 727, col: 10, offset: 18627},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 727, col: 15, offset: 18632},
	expr: &ruleRefExpr{
	pos: position{line: 727, col: 15, offset: 18632},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 727, col: 24, offset: 18641},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 727, col: 28, offset: 18645},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 731, col: 1, offset: 18712},
	expr: &choiceExpr{
	pos: position{line: 731, col: 12, offset: 18723},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 731, col: 12, offset: 18723},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 731, col: 18, offset: 18729},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 731, col: 18, offset: 18729},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 731, col: 38, offset: 18749},
	expr: &ruleRefExpr{
	pos: position{line: 731, col: 38, offset: 18749},
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
	pos: position{line: 733, col: 1, offset: 18766},
	expr: &charClassMatcher{
	pos: position{line: 733, col: 17, offset: 18782},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 735, col: 1, offset: 18791},
	expr: &charClassMatcher{
	pos: position{line: 735, col: 24, offset: 18814},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 737, col: 1, offset: 18823},
	expr: &zeroOrMoreExpr{
	pos: position{line: 737, col: 19, offset: 18841},
	expr: &charClassMatcher{
	pos: position{line: 737, col: 19, offset: 18841},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 739, col: 1, offset: 18855},
	expr: &charClassMatcher{
	pos: position{line: 739, col: 18, offset: 18872},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "IDENTIFIER",
	pos: position{line: 740, col: 1, offset: 18883},
	expr: &actionExpr{
	pos: position{line: 740, col: 15, offset: 18897},
	run: (*parser).callonIDENTIFIER1,
	expr: &seqExpr{
	pos: position{line: 740, col: 15, offset: 18897},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 740, col: 15, offset: 18897},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 740, col: 22, offset: 18904},
	expr: &choiceExpr{
	pos: position{line: 740, col: 23, offset: 18905},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 740, col: 23, offset: 18905},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 740, col: 33, offset: 18915},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 740, col: 41, offset: 18923},
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
	pos: position{line: 743, col: 1, offset: 18996},
	expr: &notExpr{
	pos: position{line: 743, col: 8, offset: 19003},
	expr: &anyMatcher{
	line: 743, col: 9, offset: 19004,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 745, col: 1, offset: 19009},
	expr: &litMatcher{
	pos: position{line: 745, col: 17, offset: 19025},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 751, col: 1, offset: 19112},
	expr: &actionExpr{
	pos: position{line: 751, col: 18, offset: 19129},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 751, col: 18, offset: 19129},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 751, col: 18, offset: 19129},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 751, col: 24, offset: 19135},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 751, col: 33, offset: 19144},
	name: "_",
},
&labeledExpr{
	pos: position{line: 751, col: 35, offset: 19146},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 751, col: 40, offset: 19151},
	expr: &seqExpr{
	pos: position{line: 751, col: 41, offset: 19152},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 751, col: 41, offset: 19152},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 751, col: 45, offset: 19156},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 751, col: 47, offset: 19158},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 751, col: 56, offset: 19167},
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
	pos: position{line: 754, col: 1, offset: 19220},
	expr: &actionExpr{
	pos: position{line: 754, col: 13, offset: 19232},
	run: (*parser).callonEqAliasG1,
	expr: &seqExpr{
	pos: position{line: 754, col: 13, offset: 19232},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 754, col: 13, offset: 19232},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 754, col: 19, offset: 19238},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 754, col: 28, offset: 19247},
	name: "_",
},
&litMatcher{
	pos: position{line: 754, col: 30, offset: 19249},
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

func (c *current) onSecSource1(src1, src2 interface{}) (interface{}, error) {

    var secSources []string
    if  source1,ok := cast.TryString(src1);ok{
        secSources = append(secSources,source1)
    }
    if source2,ok := cast.TryString(src2);ok{
        secSources = append(secSources,source2)
    }
    return SourceJob{Type:SECSOURCE,OperateOn:SrcSec{Sources:secSources}},nil
}

func (p *parser) callonSecSource1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSecSource1(stack["src1"], stack["src2"])
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

