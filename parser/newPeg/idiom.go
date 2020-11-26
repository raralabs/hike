
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
	expr: &choiceExpr{
	pos: position{line: 5, col: 12, offset: 35},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 5, col: 12, offset: 35},
	run: (*parser).callonCommand2,
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
&actionExpr{
	pos: position{line: 32, col: 7, offset: 896},
	run: (*parser).callonCommand14,
	expr: &seqExpr{
	pos: position{line: 32, col: 7, offset: 896},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 32, col: 7, offset: 896},
	label: "command",
	expr: &ruleRefExpr{
	pos: position{line: 32, col: 15, offset: 904},
	name: "Statement",
},
},
&labeledExpr{
	pos: position{line: 32, col: 25, offset: 914},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 32, col: 31, offset: 920},
	exprs: []interface{}{
&zeroOrMoreExpr{
	pos: position{line: 32, col: 31, offset: 920},
	expr: &seqExpr{
	pos: position{line: 32, col: 32, offset: 921},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 32, col: 32, offset: 921},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 32, col: 34, offset: 923},
	name: "Statement",
},
	},
},
},
&ruleRefExpr{
	pos: position{line: 32, col: 47, offset: 936},
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
{
	name: "Statement",
	pos: position{line: 38, col: 1, offset: 1023},
	expr: &choiceExpr{
	pos: position{line: 38, col: 14, offset: 1036},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 38, col: 14, offset: 1036},
	run: (*parser).callonStatement2,
	expr: &seqExpr{
	pos: position{line: 38, col: 14, offset: 1036},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 38, col: 14, offset: 1036},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 38, col: 24, offset: 1046},
	name: "Source",
},
},
&labeledExpr{
	pos: position{line: 38, col: 31, offset: 1053},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 38, col: 37, offset: 1059},
	exprs: []interface{}{
&zeroOrMoreExpr{
	pos: position{line: 38, col: 37, offset: 1059},
	expr: &seqExpr{
	pos: position{line: 38, col: 38, offset: 1060},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 38, col: 38, offset: 1060},
	name: "_",
},
&litMatcher{
	pos: position{line: 38, col: 40, offset: 1062},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 38, col: 44, offset: 1066},
	name: "_",
},
&labeledExpr{
	pos: position{line: 38, col: 46, offset: 1068},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 38, col: 56, offset: 1078},
	name: "Transform",
},
},
	},
},
},
&ruleRefExpr{
	pos: position{line: 38, col: 68, offset: 1090},
	name: "_",
},
&litMatcher{
	pos: position{line: 38, col: 69, offset: 1091},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 38, col: 72, offset: 1094},
	name: "_",
},
&labeledExpr{
	pos: position{line: 38, col: 74, offset: 1096},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 38, col: 84, offset: 1106},
	name: "Sink",
},
},
&ruleRefExpr{
	pos: position{line: 38, col: 89, offset: 1111},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 38, col: 91, offset: 1113},
	name: "TOK_SEMICOLON",
},
&ruleRefExpr{
	pos: position{line: 38, col: 105, offset: 1127},
	name: "_",
},
	},
},
},
	},
},
},
&actionExpr{
	pos: position{line: 69, col: 7, offset: 1821},
	run: (*parser).callonStatement23,
	expr: &seqExpr{
	pos: position{line: 69, col: 7, offset: 1821},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 69, col: 7, offset: 1821},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 69, col: 17, offset: 1831},
	name: "Source",
},
},
&labeledExpr{
	pos: position{line: 69, col: 24, offset: 1838},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 69, col: 30, offset: 1844},
	exprs: []interface{}{
&zeroOrMoreExpr{
	pos: position{line: 69, col: 30, offset: 1844},
	expr: &seqExpr{
	pos: position{line: 69, col: 31, offset: 1845},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 69, col: 31, offset: 1845},
	name: "_",
},
&litMatcher{
	pos: position{line: 69, col: 33, offset: 1847},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 69, col: 37, offset: 1851},
	name: "_",
},
&labeledExpr{
	pos: position{line: 69, col: 39, offset: 1853},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 69, col: 49, offset: 1863},
	name: "Transform",
},
},
	},
},
},
&ruleRefExpr{
	pos: position{line: 69, col: 61, offset: 1875},
	name: "_",
},
&litMatcher{
	pos: position{line: 69, col: 62, offset: 1876},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 69, col: 65, offset: 1879},
	name: "_",
},
&labeledExpr{
	pos: position{line: 69, col: 67, offset: 1881},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 69, col: 77, offset: 1891},
	name: "Sink",
},
},
&ruleRefExpr{
	pos: position{line: 69, col: 82, offset: 1896},
	name: "_",
},
	},
},
},
	},
},
},
&actionExpr{
	pos: position{line: 71, col: 7, offset: 1990},
	run: (*parser).callonStatement42,
	expr: &seqExpr{
	pos: position{line: 71, col: 7, offset: 1990},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 71, col: 7, offset: 1990},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 71, col: 17, offset: 2000},
	name: "Source",
},
},
&labeledExpr{
	pos: position{line: 71, col: 24, offset: 2007},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 71, col: 30, offset: 2013},
	exprs: []interface{}{
&zeroOrMoreExpr{
	pos: position{line: 71, col: 30, offset: 2013},
	expr: &seqExpr{
	pos: position{line: 71, col: 31, offset: 2014},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 71, col: 31, offset: 2014},
	name: "_",
},
&litMatcher{
	pos: position{line: 71, col: 33, offset: 2016},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 71, col: 37, offset: 2020},
	name: "_",
},
&labeledExpr{
	pos: position{line: 71, col: 39, offset: 2022},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 71, col: 49, offset: 2032},
	name: "Transform",
},
},
	},
},
},
&ruleRefExpr{
	pos: position{line: 71, col: 61, offset: 2044},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 71, col: 63, offset: 2046},
	name: "TOK_SEMICOLON",
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
	name: "Source",
	pos: position{line: 77, col: 1, offset: 2143},
	expr: &choiceExpr{
	pos: position{line: 77, col: 11, offset: 2153},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 77, col: 11, offset: 2153},
	name: "BranchJobG",
},
&ruleRefExpr{
	pos: position{line: 77, col: 22, offset: 2164},
	name: "FileJobG",
},
&ruleRefExpr{
	pos: position{line: 77, col: 31, offset: 2173},
	name: "FakeJobG",
},
&ruleRefExpr{
	pos: position{line: 77, col: 40, offset: 2182},
	name: "SecSource",
},
	},
},
},
{
	name: "Transform",
	pos: position{line: 79, col: 1, offset: 2195},
	expr: &choiceExpr{
	pos: position{line: 79, col: 14, offset: 2208},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 79, col: 14, offset: 2208},
	name: "AggJobG",
},
&ruleRefExpr{
	pos: position{line: 79, col: 22, offset: 2216},
	name: "DoJobG",
},
&ruleRefExpr{
	pos: position{line: 79, col: 29, offset: 2223},
	name: "JoinJobG",
},
	},
},
},
{
	name: "Sink",
	pos: position{line: 81, col: 1, offset: 2235},
	expr: &choiceExpr{
	pos: position{line: 81, col: 9, offset: 2243},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 81, col: 9, offset: 2243},
	name: "StdOutG",
},
&ruleRefExpr{
	pos: position{line: 81, col: 17, offset: 2251},
	name: "SinkInto",
},
&ruleRefExpr{
	pos: position{line: 81, col: 26, offset: 2260},
	name: "PlotG",
},
	},
},
},
{
	name: "FileJobG",
	pos: position{line: 85, col: 1, offset: 2286},
	expr: &choiceExpr{
	pos: position{line: 85, col: 13, offset: 2298},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 85, col: 13, offset: 2298},
	run: (*parser).callonFileJobG2,
	expr: &seqExpr{
	pos: position{line: 85, col: 13, offset: 2298},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 85, col: 13, offset: 2298},
	name: "_",
},
&litMatcher{
	pos: position{line: 85, col: 15, offset: 2300},
	val: "file(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 85, col: 23, offset: 2308},
	name: "_",
},
&labeledExpr{
	pos: position{line: 85, col: 25, offset: 2310},
	label: "filename",
	expr: &ruleRefExpr{
	pos: position{line: 85, col: 34, offset: 2319},
	name: "FileName",
},
},
&ruleRefExpr{
	pos: position{line: 85, col: 43, offset: 2328},
	name: "_",
},
&litMatcher{
	pos: position{line: 85, col: 45, offset: 2330},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 90, col: 4, offset: 2543},
	run: (*parser).callonFileJobG11,
	expr: &seqExpr{
	pos: position{line: 90, col: 4, offset: 2543},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 90, col: 4, offset: 2543},
	name: "_",
},
&litMatcher{
	pos: position{line: 90, col: 6, offset: 2545},
	val: "file()",
	ignoreCase: false,
},
	},
},
},
	},
},
},
{
	name: "FakeJobG",
	pos: position{line: 94, col: 1, offset: 2609},
	expr: &choiceExpr{
	pos: position{line: 94, col: 13, offset: 2621},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 94, col: 13, offset: 2621},
	run: (*parser).callonFakeJobG2,
	expr: &seqExpr{
	pos: position{line: 94, col: 13, offset: 2621},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 94, col: 13, offset: 2621},
	name: "_",
},
&litMatcher{
	pos: position{line: 94, col: 15, offset: 2623},
	val: "fake",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 94, col: 21, offset: 2629},
	name: "_",
},
&litMatcher{
	pos: position{line: 94, col: 23, offset: 2631},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 94, col: 27, offset: 2635},
	name: "_",
},
&labeledExpr{
	pos: position{line: 94, col: 29, offset: 2637},
	label: "numData",
	expr: &ruleRefExpr{
	pos: position{line: 94, col: 37, offset: 2645},
	name: "PositiveNumber",
},
},
&ruleRefExpr{
	pos: position{line: 94, col: 52, offset: 2660},
	name: "_",
},
&litMatcher{
	pos: position{line: 94, col: 54, offset: 2662},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 99, col: 3, offset: 2848},
	run: (*parser).callonFakeJobG13,
	expr: &seqExpr{
	pos: position{line: 99, col: 3, offset: 2848},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 99, col: 3, offset: 2848},
	name: "_",
},
&litMatcher{
	pos: position{line: 99, col: 5, offset: 2850},
	val: "fake",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 99, col: 11, offset: 2856},
	name: "_",
},
&litMatcher{
	pos: position{line: 99, col: 13, offset: 2858},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 99, col: 17, offset: 2862},
	name: "_",
},
&litMatcher{
	pos: position{line: 99, col: 19, offset: 2864},
	val: ")",
	ignoreCase: false,
},
	},
},
},
	},
},
},
{
	name: "BranchJobG",
	pos: position{line: 103, col: 1, offset: 2948},
	expr: &actionExpr{
	pos: position{line: 103, col: 15, offset: 2962},
	run: (*parser).callonBranchJobG1,
	expr: &seqExpr{
	pos: position{line: 103, col: 15, offset: 2962},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 103, col: 15, offset: 2962},
	val: "branch",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 103, col: 23, offset: 2970},
	name: "_",
},
&litMatcher{
	pos: position{line: 103, col: 24, offset: 2971},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 103, col: 28, offset: 2975},
	name: "_",
},
&labeledExpr{
	pos: position{line: 103, col: 30, offset: 2977},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 103, col: 33, offset: 2980},
	name: "IDENTIFIER",
},
},
&litMatcher{
	pos: position{line: 103, col: 44, offset: 2991},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SecSource",
	pos: position{line: 110, col: 1, offset: 3207},
	expr: &choiceExpr{
	pos: position{line: 110, col: 14, offset: 3220},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 110, col: 14, offset: 3220},
	run: (*parser).callonSecSource2,
	expr: &seqExpr{
	pos: position{line: 110, col: 14, offset: 3220},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 110, col: 14, offset: 3220},
	label: "src1",
	expr: &ruleRefExpr{
	pos: position{line: 110, col: 19, offset: 3225},
	name: "IDENTIFIER",
},
},
&labeledExpr{
	pos: position{line: 110, col: 30, offset: 3236},
	label: "rest",
	expr: &zeroOrOneExpr{
	pos: position{line: 110, col: 35, offset: 3241},
	expr: &seqExpr{
	pos: position{line: 110, col: 36, offset: 3242},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 110, col: 36, offset: 3242},
	name: "_",
},
&litMatcher{
	pos: position{line: 110, col: 38, offset: 3244},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 110, col: 41, offset: 3247},
	name: "_",
},
&labeledExpr{
	pos: position{line: 110, col: 43, offset: 3249},
	label: "src2",
	expr: &ruleRefExpr{
	pos: position{line: 110, col: 48, offset: 3254},
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
&actionExpr{
	pos: position{line: 122, col: 3, offset: 3663},
	run: (*parser).callonSecSource14,
	expr: &seqExpr{
	pos: position{line: 122, col: 3, offset: 3663},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 122, col: 3, offset: 3663},
	label: "src1",
	expr: &ruleRefExpr{
	pos: position{line: 122, col: 8, offset: 3668},
	name: "IDENTIFIER",
},
},
&labeledExpr{
	pos: position{line: 122, col: 19, offset: 3679},
	label: "rest",
	expr: &ruleRefExpr{
	pos: position{line: 122, col: 25, offset: 3685},
	name: "SecSource",
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
	pos: position{line: 131, col: 1, offset: 3820},
	expr: &actionExpr{
	pos: position{line: 131, col: 13, offset: 3832},
	run: (*parser).callonJoinJobG1,
	expr: &seqExpr{
	pos: position{line: 131, col: 13, offset: 3832},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 131, col: 13, offset: 3832},
	val: "join",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 131, col: 20, offset: 3839},
	name: "_",
},
&labeledExpr{
	pos: position{line: 131, col: 22, offset: 3841},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 131, col: 26, offset: 3845},
	name: "JoinJobs",
},
},
	},
},
},
},
{
	name: "JoinJobs",
	pos: position{line: 135, col: 1, offset: 3918},
	expr: &choiceExpr{
	pos: position{line: 135, col: 12, offset: 3929},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 135, col: 12, offset: 3929},
	name: "InnerJoinG",
},
&ruleRefExpr{
	pos: position{line: 135, col: 23, offset: 3940},
	name: "OuterJoinG",
},
	},
},
},
{
	name: "InnerJoinG",
	pos: position{line: 137, col: 1, offset: 3954},
	expr: &choiceExpr{
	pos: position{line: 137, col: 14, offset: 3967},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 137, col: 14, offset: 3967},
	run: (*parser).callonInnerJoinG2,
	expr: &seqExpr{
	pos: position{line: 137, col: 14, offset: 3967},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 137, col: 14, offset: 3967},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 137, col: 20, offset: 3973},
	expr: &ruleRefExpr{
	pos: position{line: 137, col: 20, offset: 3973},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 137, col: 30, offset: 3983},
	name: "_",
},
&litMatcher{
	pos: position{line: 137, col: 32, offset: 3985},
	val: "inner",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 137, col: 40, offset: 3993},
	name: "_",
},
&litMatcher{
	pos: position{line: 137, col: 42, offset: 3995},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 137, col: 46, offset: 3999},
	name: "_",
},
&labeledExpr{
	pos: position{line: 137, col: 48, offset: 4001},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 137, col: 54, offset: 4007},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 137, col: 60, offset: 4013},
	name: "_",
},
&litMatcher{
	pos: position{line: 137, col: 62, offset: 4015},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 143, col: 4, offset: 4208},
	run: (*parser).callonInnerJoinG16,
	expr: &seqExpr{
	pos: position{line: 143, col: 4, offset: 4208},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 143, col: 4, offset: 4208},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 143, col: 10, offset: 4214},
	expr: &ruleRefExpr{
	pos: position{line: 143, col: 10, offset: 4214},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 143, col: 20, offset: 4224},
	name: "_",
},
&litMatcher{
	pos: position{line: 143, col: 22, offset: 4226},
	val: "inner",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 143, col: 30, offset: 4234},
	name: "_",
},
&litMatcher{
	pos: position{line: 143, col: 32, offset: 4236},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 143, col: 36, offset: 4240},
	name: "_",
},
&litMatcher{
	pos: position{line: 143, col: 38, offset: 4242},
	val: ")",
	ignoreCase: false,
},
	},
},
},
	},
},
},
{
	name: "OuterJoinG",
	pos: position{line: 147, col: 1, offset: 4302},
	expr: &choiceExpr{
	pos: position{line: 147, col: 15, offset: 4316},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 147, col: 15, offset: 4316},
	name: "LeftOuterG",
},
&ruleRefExpr{
	pos: position{line: 147, col: 26, offset: 4327},
	name: "RightOuterG",
},
&ruleRefExpr{
	pos: position{line: 147, col: 38, offset: 4339},
	name: "FullOuterG",
},
	},
},
},
{
	name: "LeftOuterG",
	pos: position{line: 149, col: 1, offset: 4353},
	expr: &choiceExpr{
	pos: position{line: 149, col: 15, offset: 4367},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 149, col: 15, offset: 4367},
	run: (*parser).callonLeftOuterG2,
	expr: &seqExpr{
	pos: position{line: 149, col: 15, offset: 4367},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 149, col: 15, offset: 4367},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 149, col: 21, offset: 4373},
	expr: &ruleRefExpr{
	pos: position{line: 149, col: 21, offset: 4373},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 149, col: 31, offset: 4383},
	name: "_",
},
&litMatcher{
	pos: position{line: 149, col: 33, offset: 4385},
	val: "leftouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 149, col: 45, offset: 4397},
	name: "_",
},
&litMatcher{
	pos: position{line: 149, col: 47, offset: 4399},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 149, col: 51, offset: 4403},
	name: "_",
},
&labeledExpr{
	pos: position{line: 149, col: 53, offset: 4405},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 149, col: 59, offset: 4411},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 149, col: 65, offset: 4417},
	name: "_",
},
&litMatcher{
	pos: position{line: 149, col: 67, offset: 4419},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 155, col: 12, offset: 4690},
	run: (*parser).callonLeftOuterG16,
	expr: &seqExpr{
	pos: position{line: 155, col: 12, offset: 4690},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 155, col: 12, offset: 4690},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 155, col: 18, offset: 4696},
	expr: &ruleRefExpr{
	pos: position{line: 155, col: 18, offset: 4696},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 155, col: 28, offset: 4706},
	name: "_",
},
&litMatcher{
	pos: position{line: 155, col: 30, offset: 4708},
	val: "leftouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 155, col: 42, offset: 4720},
	name: "_",
},
&litMatcher{
	pos: position{line: 155, col: 44, offset: 4722},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 155, col: 48, offset: 4726},
	name: "_",
},
&litMatcher{
	pos: position{line: 155, col: 50, offset: 4728},
	val: ")",
	ignoreCase: false,
},
	},
},
},
	},
},
},
{
	name: "RightOuterG",
	pos: position{line: 159, col: 1, offset: 4807},
	expr: &choiceExpr{
	pos: position{line: 159, col: 15, offset: 4821},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 159, col: 15, offset: 4821},
	run: (*parser).callonRightOuterG2,
	expr: &seqExpr{
	pos: position{line: 159, col: 15, offset: 4821},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 159, col: 15, offset: 4821},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 159, col: 21, offset: 4827},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 159, col: 30, offset: 4836},
	name: "_",
},
&litMatcher{
	pos: position{line: 159, col: 32, offset: 4838},
	val: "rightouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 159, col: 45, offset: 4851},
	name: "_",
},
&litMatcher{
	pos: position{line: 159, col: 47, offset: 4853},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 159, col: 51, offset: 4857},
	name: "_",
},
&labeledExpr{
	pos: position{line: 159, col: 53, offset: 4859},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 159, col: 59, offset: 4865},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 159, col: 65, offset: 4871},
	name: "_",
},
&litMatcher{
	pos: position{line: 159, col: 67, offset: 4873},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 165, col: 17, offset: 5149},
	run: (*parser).callonRightOuterG15,
	expr: &seqExpr{
	pos: position{line: 165, col: 17, offset: 5149},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 165, col: 17, offset: 5149},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 165, col: 23, offset: 5155},
	expr: &ruleRefExpr{
	pos: position{line: 165, col: 23, offset: 5155},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 165, col: 33, offset: 5165},
	name: "_",
},
&litMatcher{
	pos: position{line: 165, col: 35, offset: 5167},
	val: "rightouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 165, col: 48, offset: 5180},
	name: "_",
},
&litMatcher{
	pos: position{line: 165, col: 50, offset: 5182},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 165, col: 54, offset: 5186},
	name: "_",
},
&litMatcher{
	pos: position{line: 165, col: 56, offset: 5188},
	val: ")",
	ignoreCase: false,
},
	},
},
},
	},
},
},
{
	name: "FullOuterG",
	pos: position{line: 168, col: 1, offset: 5295},
	expr: &choiceExpr{
	pos: position{line: 168, col: 15, offset: 5309},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 168, col: 15, offset: 5309},
	run: (*parser).callonFullOuterG2,
	expr: &seqExpr{
	pos: position{line: 168, col: 15, offset: 5309},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 168, col: 15, offset: 5309},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 168, col: 21, offset: 5315},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 168, col: 30, offset: 5324},
	name: "_",
},
&litMatcher{
	pos: position{line: 168, col: 32, offset: 5326},
	val: "rightouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 168, col: 45, offset: 5339},
	name: "_",
},
&litMatcher{
	pos: position{line: 168, col: 47, offset: 5341},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 168, col: 51, offset: 5345},
	name: "_",
},
&labeledExpr{
	pos: position{line: 168, col: 53, offset: 5347},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 168, col: 59, offset: 5353},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 168, col: 65, offset: 5359},
	name: "_",
},
&litMatcher{
	pos: position{line: 168, col: 67, offset: 5361},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 174, col: 17, offset: 5638},
	run: (*parser).callonFullOuterG15,
	expr: &seqExpr{
	pos: position{line: 174, col: 17, offset: 5638},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 174, col: 17, offset: 5638},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 174, col: 23, offset: 5644},
	expr: &ruleRefExpr{
	pos: position{line: 174, col: 23, offset: 5644},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 174, col: 33, offset: 5654},
	name: "_",
},
&litMatcher{
	pos: position{line: 174, col: 35, offset: 5656},
	val: "fullouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 174, col: 47, offset: 5668},
	name: "_",
},
&litMatcher{
	pos: position{line: 174, col: 49, offset: 5670},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 174, col: 53, offset: 5674},
	name: "_",
},
&litMatcher{
	pos: position{line: 174, col: 55, offset: 5676},
	val: ")",
	ignoreCase: false,
},
	},
},
},
	},
},
},
{
	name: "Query",
	pos: position{line: 178, col: 1, offset: 5768},
	expr: &choiceExpr{
	pos: position{line: 178, col: 10, offset: 5777},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 178, col: 10, offset: 5777},
	run: (*parser).callonQuery2,
	expr: &seqExpr{
	pos: position{line: 178, col: 10, offset: 5777},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 178, col: 10, offset: 5777},
	val: "select",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 178, col: 19, offset: 5786},
	name: "_",
},
&labeledExpr{
	pos: position{line: 178, col: 21, offset: 5788},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 178, col: 27, offset: 5794},
	name: "Fields",
},
},
&ruleRefExpr{
	pos: position{line: 178, col: 34, offset: 5801},
	name: "_",
},
&litMatcher{
	pos: position{line: 178, col: 36, offset: 5803},
	val: "on",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 178, col: 41, offset: 5808},
	name: "_",
},
&labeledExpr{
	pos: position{line: 178, col: 43, offset: 5810},
	label: "condition",
	expr: &ruleRefExpr{
	pos: position{line: 178, col: 53, offset: 5820},
	name: "JoinCondition",
},
},
&ruleRefExpr{
	pos: position{line: 178, col: 67, offset: 5834},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 180, col: 4, offset: 5944},
	run: (*parser).callonQuery14,
	expr: &seqExpr{
	pos: position{line: 180, col: 4, offset: 5944},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 180, col: 4, offset: 5944},
	val: "select",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 180, col: 13, offset: 5953},
	name: "_",
},
&litMatcher{
	pos: position{line: 180, col: 15, offset: 5955},
	val: "on",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 180, col: 20, offset: 5960},
	name: "_",
},
&labeledExpr{
	pos: position{line: 180, col: 22, offset: 5962},
	label: "condition",
	expr: &ruleRefExpr{
	pos: position{line: 180, col: 32, offset: 5972},
	name: "JoinCondition",
},
},
&ruleRefExpr{
	pos: position{line: 180, col: 46, offset: 5986},
	name: "_",
},
	},
},
},
	},
},
},
{
	name: "Fields",
	pos: position{line: 184, col: 1, offset: 6071},
	expr: &choiceExpr{
	pos: position{line: 184, col: 11, offset: 6081},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 184, col: 11, offset: 6081},
	name: "AllFields",
},
&ruleRefExpr{
	pos: position{line: 184, col: 21, offset: 6091},
	name: "SelFields",
},
	},
},
},
{
	name: "AllFields",
	pos: position{line: 186, col: 1, offset: 6104},
	expr: &actionExpr{
	pos: position{line: 186, col: 14, offset: 6117},
	run: (*parser).callonAllFields1,
	expr: &litMatcher{
	pos: position{line: 186, col: 14, offset: 6117},
	val: "*",
	ignoreCase: false,
},
},
},
{
	name: "SelFields",
	pos: position{line: 192, col: 1, offset: 6209},
	expr: &actionExpr{
	pos: position{line: 192, col: 14, offset: 6222},
	run: (*parser).callonSelFields1,
	expr: &seqExpr{
	pos: position{line: 192, col: 14, offset: 6222},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 192, col: 14, offset: 6222},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 192, col: 20, offset: 6228},
	name: "Word",
},
},
&ruleRefExpr{
	pos: position{line: 192, col: 25, offset: 6233},
	name: "_",
},
&labeledExpr{
	pos: position{line: 192, col: 28, offset: 6236},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 192, col: 33, offset: 6241},
	expr: &seqExpr{
	pos: position{line: 192, col: 34, offset: 6242},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 192, col: 34, offset: 6242},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 192, col: 38, offset: 6246},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 192, col: 40, offset: 6248},
	name: "Word",
},
&ruleRefExpr{
	pos: position{line: 192, col: 45, offset: 6253},
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
	pos: position{line: 202, col: 1, offset: 6506},
	expr: &choiceExpr{
	pos: position{line: 202, col: 18, offset: 6523},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 202, col: 18, offset: 6523},
	run: (*parser).callonJoinCondition2,
	expr: &seqExpr{
	pos: position{line: 202, col: 18, offset: 6523},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 202, col: 18, offset: 6523},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 202, col: 23, offset: 6528},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 202, col: 35, offset: 6540},
	name: "_",
},
&labeledExpr{
	pos: position{line: 202, col: 37, offset: 6542},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 202, col: 40, offset: 6545},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 202, col: 52, offset: 6557},
	name: "_",
},
&labeledExpr{
	pos: position{line: 202, col: 54, offset: 6559},
	label: "right",
	expr: &ruleRefExpr{
	pos: position{line: 202, col: 60, offset: 6565},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 202, col: 72, offset: 6577},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 204, col: 4, offset: 6695},
	run: (*parser).callonJoinCondition13,
	expr: &seqExpr{
	pos: position{line: 204, col: 4, offset: 6695},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 204, col: 4, offset: 6695},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 204, col: 9, offset: 6700},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 204, col: 21, offset: 6712},
	name: "_",
},
&labeledExpr{
	pos: position{line: 204, col: 23, offset: 6714},
	label: "right",
	expr: &ruleRefExpr{
	pos: position{line: 204, col: 29, offset: 6720},
	name: "JoinFacotors",
},
},
&ruleRefExpr{
	pos: position{line: 204, col: 42, offset: 6733},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 206, col: 3, offset: 6816},
	run: (*parser).callonJoinCondition21,
	expr: &seqExpr{
	pos: position{line: 206, col: 3, offset: 6816},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 206, col: 3, offset: 6816},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 206, col: 8, offset: 6821},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 206, col: 20, offset: 6833},
	name: "_",
},
&labeledExpr{
	pos: position{line: 206, col: 22, offset: 6835},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 206, col: 25, offset: 6838},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 206, col: 37, offset: 6850},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 208, col: 3, offset: 6932},
	run: (*parser).callonJoinCondition29,
	expr: &seqExpr{
	pos: position{line: 208, col: 3, offset: 6932},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 208, col: 3, offset: 6932},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 208, col: 8, offset: 6937},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 208, col: 20, offset: 6949},
	name: "_",
},
	},
},
},
	},
},
},
{
	name: "JoinFactors",
	pos: position{line: 213, col: 1, offset: 7044},
	expr: &actionExpr{
	pos: position{line: 213, col: 16, offset: 7059},
	run: (*parser).callonJoinFactors1,
	expr: &seqExpr{
	pos: position{line: 213, col: 16, offset: 7059},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 213, col: 16, offset: 7059},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 213, col: 22, offset: 7065},
	name: "Word",
},
},
&ruleRefExpr{
	pos: position{line: 213, col: 27, offset: 7070},
	name: "_",
},
&labeledExpr{
	pos: position{line: 213, col: 29, offset: 7072},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 213, col: 34, offset: 7077},
	expr: &seqExpr{
	pos: position{line: 213, col: 35, offset: 7078},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 213, col: 35, offset: 7078},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 213, col: 39, offset: 7082},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 213, col: 41, offset: 7084},
	name: "Word",
},
&ruleRefExpr{
	pos: position{line: 213, col: 46, offset: 7089},
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
	pos: position{line: 229, col: 1, offset: 7446},
	expr: &actionExpr{
	pos: position{line: 229, col: 12, offset: 7457},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 229, col: 12, offset: 7457},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 229, col: 12, offset: 7457},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 229, col: 18, offset: 7463},
	name: "_",
},
&labeledExpr{
	pos: position{line: 229, col: 20, offset: 7465},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 229, col: 27, offset: 7472},
	expr: &ruleRefExpr{
	pos: position{line: 229, col: 27, offset: 7472},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 229, col: 39, offset: 7484},
	name: "_",
},
&labeledExpr{
	pos: position{line: 229, col: 41, offset: 7486},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 229, col: 45, offset: 7490},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 229, col: 53, offset: 7498},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 229, col: 58, offset: 7503},
	expr: &seqExpr{
	pos: position{line: 229, col: 59, offset: 7504},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 229, col: 59, offset: 7504},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 229, col: 63, offset: 7508},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 229, col: 65, offset: 7510},
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
	pos: position{line: 244, col: 1, offset: 7885},
	expr: &choiceExpr{
	pos: position{line: 244, col: 12, offset: 7896},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 244, col: 12, offset: 7896},
	name: "CountAgg",
},
&ruleRefExpr{
	pos: position{line: 244, col: 23, offset: 7907},
	name: "MaxAgg",
},
&ruleRefExpr{
	pos: position{line: 244, col: 32, offset: 7916},
	name: "MinAgg",
},
&ruleRefExpr{
	pos: position{line: 244, col: 41, offset: 7925},
	name: "AvgAgg",
},
&ruleRefExpr{
	pos: position{line: 244, col: 50, offset: 7934},
	name: "SumAgg",
},
&ruleRefExpr{
	pos: position{line: 244, col: 59, offset: 7943},
	name: "ModeAgg",
},
&ruleRefExpr{
	pos: position{line: 245, col: 13, offset: 7964},
	name: "VarAgg",
},
&ruleRefExpr{
	pos: position{line: 245, col: 22, offset: 7973},
	name: "DistinctCountAgg",
},
&ruleRefExpr{
	pos: position{line: 245, col: 41, offset: 7992},
	name: "QuantileAgg",
},
&ruleRefExpr{
	pos: position{line: 245, col: 55, offset: 8006},
	name: "MedianAgg",
},
&ruleRefExpr{
	pos: position{line: 245, col: 67, offset: 8018},
	name: "QuartileAgg",
},
&ruleRefExpr{
	pos: position{line: 246, col: 13, offset: 8043},
	name: "CovAgg",
},
&ruleRefExpr{
	pos: position{line: 246, col: 22, offset: 8052},
	name: "CorrAgg",
},
&ruleRefExpr{
	pos: position{line: 246, col: 32, offset: 8062},
	name: "PValueAgg",
},
	},
},
},
{
	name: "CountAgg",
	pos: position{line: 249, col: 1, offset: 8106},
	expr: &actionExpr{
	pos: position{line: 249, col: 13, offset: 8118},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 249, col: 13, offset: 8118},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 249, col: 13, offset: 8118},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 249, col: 19, offset: 8124},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 249, col: 28, offset: 8133},
	name: "_",
},
&litMatcher{
	pos: position{line: 249, col: 30, offset: 8135},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 249, col: 39, offset: 8144},
	name: "_",
},
&labeledExpr{
	pos: position{line: 249, col: 41, offset: 8146},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 249, col: 46, offset: 8151},
	expr: &ruleRefExpr{
	pos: position{line: 249, col: 46, offset: 8151},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 249, col: 59, offset: 8164},
	name: "_",
},
&litMatcher{
	pos: position{line: 249, col: 61, offset: 8166},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MaxAgg",
	pos: position{line: 256, col: 1, offset: 8318},
	expr: &actionExpr{
	pos: position{line: 256, col: 11, offset: 8328},
	run: (*parser).callonMaxAgg1,
	expr: &seqExpr{
	pos: position{line: 256, col: 11, offset: 8328},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 256, col: 11, offset: 8328},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 256, col: 17, offset: 8334},
	expr: &ruleRefExpr{
	pos: position{line: 256, col: 17, offset: 8334},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 256, col: 27, offset: 8344},
	name: "_",
},
&litMatcher{
	pos: position{line: 256, col: 29, offset: 8346},
	val: "max(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 256, col: 36, offset: 8353},
	name: "_",
},
&labeledExpr{
	pos: position{line: 256, col: 38, offset: 8355},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 256, col: 44, offset: 8361},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 256, col: 53, offset: 8370},
	name: "_",
},
&labeledExpr{
	pos: position{line: 256, col: 55, offset: 8372},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 256, col: 60, offset: 8377},
	expr: &seqExpr{
	pos: position{line: 256, col: 61, offset: 8378},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 256, col: 61, offset: 8378},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 256, col: 65, offset: 8382},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 256, col: 67, offset: 8384},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 256, col: 81, offset: 8398},
	name: "_",
},
&litMatcher{
	pos: position{line: 256, col: 83, offset: 8400},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MinAgg",
	pos: position{line: 278, col: 1, offset: 8861},
	expr: &actionExpr{
	pos: position{line: 278, col: 11, offset: 8871},
	run: (*parser).callonMinAgg1,
	expr: &seqExpr{
	pos: position{line: 278, col: 11, offset: 8871},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 278, col: 11, offset: 8871},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 278, col: 17, offset: 8877},
	expr: &ruleRefExpr{
	pos: position{line: 278, col: 17, offset: 8877},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 278, col: 27, offset: 8887},
	name: "_",
},
&litMatcher{
	pos: position{line: 278, col: 29, offset: 8889},
	val: "min(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 278, col: 36, offset: 8896},
	name: "_",
},
&labeledExpr{
	pos: position{line: 278, col: 38, offset: 8898},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 278, col: 44, offset: 8904},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 278, col: 53, offset: 8913},
	name: "_",
},
&labeledExpr{
	pos: position{line: 278, col: 55, offset: 8915},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 278, col: 60, offset: 8920},
	expr: &seqExpr{
	pos: position{line: 278, col: 61, offset: 8921},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 278, col: 61, offset: 8921},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 278, col: 65, offset: 8925},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 278, col: 67, offset: 8927},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 278, col: 81, offset: 8941},
	name: "_",
},
&litMatcher{
	pos: position{line: 278, col: 83, offset: 8943},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AvgAgg",
	pos: position{line: 300, col: 1, offset: 9404},
	expr: &actionExpr{
	pos: position{line: 300, col: 11, offset: 9414},
	run: (*parser).callonAvgAgg1,
	expr: &seqExpr{
	pos: position{line: 300, col: 11, offset: 9414},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 300, col: 11, offset: 9414},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 300, col: 17, offset: 9420},
	expr: &ruleRefExpr{
	pos: position{line: 300, col: 17, offset: 9420},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 300, col: 27, offset: 9430},
	name: "_",
},
&litMatcher{
	pos: position{line: 300, col: 29, offset: 9432},
	val: "avg(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 300, col: 36, offset: 9439},
	name: "_",
},
&labeledExpr{
	pos: position{line: 300, col: 38, offset: 9441},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 300, col: 44, offset: 9447},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 300, col: 53, offset: 9456},
	name: "_",
},
&labeledExpr{
	pos: position{line: 300, col: 55, offset: 9458},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 300, col: 60, offset: 9463},
	expr: &seqExpr{
	pos: position{line: 300, col: 61, offset: 9464},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 300, col: 61, offset: 9464},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 300, col: 65, offset: 9468},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 300, col: 67, offset: 9470},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 300, col: 81, offset: 9484},
	name: "_",
},
&litMatcher{
	pos: position{line: 300, col: 83, offset: 9486},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SumAgg",
	pos: position{line: 322, col: 1, offset: 9947},
	expr: &actionExpr{
	pos: position{line: 322, col: 11, offset: 9957},
	run: (*parser).callonSumAgg1,
	expr: &seqExpr{
	pos: position{line: 322, col: 11, offset: 9957},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 322, col: 11, offset: 9957},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 322, col: 17, offset: 9963},
	expr: &ruleRefExpr{
	pos: position{line: 322, col: 17, offset: 9963},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 322, col: 27, offset: 9973},
	name: "_",
},
&litMatcher{
	pos: position{line: 322, col: 29, offset: 9975},
	val: "sum(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 322, col: 36, offset: 9982},
	name: "_",
},
&labeledExpr{
	pos: position{line: 322, col: 38, offset: 9984},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 322, col: 44, offset: 9990},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 322, col: 53, offset: 9999},
	name: "_",
},
&labeledExpr{
	pos: position{line: 322, col: 55, offset: 10001},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 322, col: 60, offset: 10006},
	expr: &seqExpr{
	pos: position{line: 322, col: 61, offset: 10007},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 322, col: 61, offset: 10007},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 322, col: 65, offset: 10011},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 322, col: 67, offset: 10013},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 322, col: 81, offset: 10027},
	name: "_",
},
&litMatcher{
	pos: position{line: 322, col: 83, offset: 10029},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarAgg",
	pos: position{line: 344, col: 1, offset: 10490},
	expr: &actionExpr{
	pos: position{line: 344, col: 11, offset: 10500},
	run: (*parser).callonVarAgg1,
	expr: &seqExpr{
	pos: position{line: 344, col: 11, offset: 10500},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 344, col: 11, offset: 10500},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 344, col: 17, offset: 10506},
	expr: &ruleRefExpr{
	pos: position{line: 344, col: 17, offset: 10506},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 344, col: 27, offset: 10516},
	name: "_",
},
&litMatcher{
	pos: position{line: 344, col: 29, offset: 10518},
	val: "var(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 344, col: 36, offset: 10525},
	name: "_",
},
&labeledExpr{
	pos: position{line: 344, col: 38, offset: 10527},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 344, col: 44, offset: 10533},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 344, col: 53, offset: 10542},
	name: "_",
},
&labeledExpr{
	pos: position{line: 344, col: 55, offset: 10544},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 344, col: 60, offset: 10549},
	expr: &seqExpr{
	pos: position{line: 344, col: 61, offset: 10550},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 344, col: 61, offset: 10550},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 344, col: 65, offset: 10554},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 344, col: 67, offset: 10556},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 344, col: 81, offset: 10570},
	name: "_",
},
&litMatcher{
	pos: position{line: 344, col: 83, offset: 10572},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DistinctCountAgg",
	pos: position{line: 366, col: 1, offset: 11038},
	expr: &actionExpr{
	pos: position{line: 366, col: 21, offset: 11058},
	run: (*parser).callonDistinctCountAgg1,
	expr: &seqExpr{
	pos: position{line: 366, col: 21, offset: 11058},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 366, col: 21, offset: 11058},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 366, col: 27, offset: 11064},
	expr: &ruleRefExpr{
	pos: position{line: 366, col: 27, offset: 11064},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 366, col: 37, offset: 11074},
	name: "_",
},
&litMatcher{
	pos: position{line: 366, col: 39, offset: 11076},
	val: "dcount(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 366, col: 49, offset: 11086},
	name: "_",
},
&labeledExpr{
	pos: position{line: 366, col: 51, offset: 11088},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 366, col: 57, offset: 11094},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 366, col: 66, offset: 11103},
	name: "_",
},
&labeledExpr{
	pos: position{line: 366, col: 68, offset: 11105},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 366, col: 73, offset: 11110},
	expr: &seqExpr{
	pos: position{line: 366, col: 74, offset: 11111},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 366, col: 74, offset: 11111},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 366, col: 78, offset: 11115},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 366, col: 80, offset: 11117},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 366, col: 94, offset: 11131},
	name: "_",
},
&litMatcher{
	pos: position{line: 366, col: 96, offset: 11133},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "ModeAgg",
	pos: position{line: 388, col: 1, offset: 11604},
	expr: &actionExpr{
	pos: position{line: 388, col: 12, offset: 11615},
	run: (*parser).callonModeAgg1,
	expr: &seqExpr{
	pos: position{line: 388, col: 12, offset: 11615},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 388, col: 12, offset: 11615},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 388, col: 18, offset: 11621},
	expr: &ruleRefExpr{
	pos: position{line: 388, col: 18, offset: 11621},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 388, col: 28, offset: 11631},
	name: "_",
},
&litMatcher{
	pos: position{line: 388, col: 30, offset: 11633},
	val: "mode(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 388, col: 38, offset: 11641},
	name: "_",
},
&labeledExpr{
	pos: position{line: 388, col: 40, offset: 11643},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 388, col: 46, offset: 11649},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 388, col: 55, offset: 11658},
	name: "_",
},
&labeledExpr{
	pos: position{line: 388, col: 57, offset: 11660},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 388, col: 62, offset: 11665},
	expr: &seqExpr{
	pos: position{line: 388, col: 63, offset: 11666},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 388, col: 63, offset: 11666},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 388, col: 67, offset: 11670},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 388, col: 69, offset: 11672},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 388, col: 83, offset: 11686},
	name: "_",
},
&litMatcher{
	pos: position{line: 388, col: 85, offset: 11688},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuantileAgg",
	pos: position{line: 410, col: 1, offset: 12150},
	expr: &actionExpr{
	pos: position{line: 410, col: 16, offset: 12165},
	run: (*parser).callonQuantileAgg1,
	expr: &seqExpr{
	pos: position{line: 410, col: 16, offset: 12165},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 410, col: 16, offset: 12165},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 410, col: 22, offset: 12171},
	expr: &ruleRefExpr{
	pos: position{line: 410, col: 22, offset: 12171},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 410, col: 32, offset: 12181},
	name: "_",
},
&litMatcher{
	pos: position{line: 410, col: 34, offset: 12183},
	val: "quantile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 410, col: 46, offset: 12195},
	name: "_",
},
&labeledExpr{
	pos: position{line: 410, col: 48, offset: 12197},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 410, col: 54, offset: 12203},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 410, col: 63, offset: 12212},
	name: "_",
},
&labeledExpr{
	pos: position{line: 410, col: 66, offset: 12215},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 410, col: 73, offset: 12222},
	expr: &seqExpr{
	pos: position{line: 410, col: 74, offset: 12223},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 410, col: 74, offset: 12223},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 410, col: 78, offset: 12227},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 410, col: 80, offset: 12229},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 410, col: 91, offset: 12240},
	name: "_",
},
&litMatcher{
	pos: position{line: 410, col: 93, offset: 12242},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 410, col: 97, offset: 12246},
	name: "_",
},
&labeledExpr{
	pos: position{line: 410, col: 99, offset: 12248},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 410, col: 103, offset: 12252},
	name: "Float",
},
},
&ruleRefExpr{
	pos: position{line: 410, col: 109, offset: 12258},
	name: "_",
},
&labeledExpr{
	pos: position{line: 410, col: 111, offset: 12260},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 410, col: 116, offset: 12265},
	expr: &seqExpr{
	pos: position{line: 410, col: 117, offset: 12266},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 410, col: 117, offset: 12266},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 410, col: 121, offset: 12270},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 410, col: 123, offset: 12272},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 410, col: 137, offset: 12286},
	name: "_",
},
&litMatcher{
	pos: position{line: 410, col: 139, offset: 12288},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MedianAgg",
	pos: position{line: 440, col: 1, offset: 12933},
	expr: &actionExpr{
	pos: position{line: 440, col: 14, offset: 12946},
	run: (*parser).callonMedianAgg1,
	expr: &seqExpr{
	pos: position{line: 440, col: 14, offset: 12946},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 440, col: 14, offset: 12946},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 440, col: 20, offset: 12952},
	expr: &ruleRefExpr{
	pos: position{line: 440, col: 20, offset: 12952},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 440, col: 30, offset: 12962},
	name: "_",
},
&litMatcher{
	pos: position{line: 440, col: 32, offset: 12964},
	val: "median(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 440, col: 42, offset: 12974},
	name: "_",
},
&labeledExpr{
	pos: position{line: 440, col: 44, offset: 12976},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 440, col: 50, offset: 12982},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 440, col: 59, offset: 12991},
	name: "_",
},
&labeledExpr{
	pos: position{line: 440, col: 62, offset: 12994},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 440, col: 69, offset: 13001},
	expr: &seqExpr{
	pos: position{line: 440, col: 70, offset: 13002},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 440, col: 70, offset: 13002},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 440, col: 74, offset: 13006},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 440, col: 76, offset: 13008},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 440, col: 87, offset: 13019},
	name: "_",
},
&labeledExpr{
	pos: position{line: 440, col: 89, offset: 13021},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 440, col: 94, offset: 13026},
	expr: &seqExpr{
	pos: position{line: 440, col: 95, offset: 13027},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 440, col: 95, offset: 13027},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 440, col: 99, offset: 13031},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 440, col: 101, offset: 13033},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 440, col: 115, offset: 13047},
	name: "_",
},
&litMatcher{
	pos: position{line: 440, col: 117, offset: 13049},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuartileAgg",
	pos: position{line: 472, col: 1, offset: 13726},
	expr: &actionExpr{
	pos: position{line: 472, col: 16, offset: 13741},
	run: (*parser).callonQuartileAgg1,
	expr: &seqExpr{
	pos: position{line: 472, col: 16, offset: 13741},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 472, col: 16, offset: 13741},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 472, col: 22, offset: 13747},
	expr: &ruleRefExpr{
	pos: position{line: 472, col: 22, offset: 13747},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 472, col: 32, offset: 13757},
	name: "_",
},
&litMatcher{
	pos: position{line: 472, col: 34, offset: 13759},
	val: "quartile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 472, col: 46, offset: 13771},
	name: "_",
},
&labeledExpr{
	pos: position{line: 472, col: 48, offset: 13773},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 472, col: 54, offset: 13779},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 472, col: 63, offset: 13788},
	name: "_",
},
&labeledExpr{
	pos: position{line: 472, col: 66, offset: 13791},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 472, col: 73, offset: 13798},
	expr: &seqExpr{
	pos: position{line: 472, col: 74, offset: 13799},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 472, col: 74, offset: 13799},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 472, col: 78, offset: 13803},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 472, col: 80, offset: 13805},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 472, col: 91, offset: 13816},
	name: "_",
},
&litMatcher{
	pos: position{line: 472, col: 93, offset: 13818},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 472, col: 97, offset: 13822},
	name: "_",
},
&labeledExpr{
	pos: position{line: 472, col: 99, offset: 13824},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 472, col: 103, offset: 13828},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 472, col: 110, offset: 13835},
	name: "_",
},
&labeledExpr{
	pos: position{line: 472, col: 112, offset: 13837},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 472, col: 117, offset: 13842},
	expr: &seqExpr{
	pos: position{line: 472, col: 118, offset: 13843},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 472, col: 118, offset: 13843},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 472, col: 122, offset: 13847},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 472, col: 124, offset: 13849},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 472, col: 138, offset: 13863},
	name: "_",
},
&litMatcher{
	pos: position{line: 472, col: 140, offset: 13865},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CovAgg",
	pos: position{line: 518, col: 1, offset: 14892},
	expr: &actionExpr{
	pos: position{line: 518, col: 11, offset: 14902},
	run: (*parser).callonCovAgg1,
	expr: &seqExpr{
	pos: position{line: 518, col: 11, offset: 14902},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 518, col: 11, offset: 14902},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 518, col: 17, offset: 14908},
	expr: &ruleRefExpr{
	pos: position{line: 518, col: 17, offset: 14908},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 518, col: 27, offset: 14918},
	name: "_",
},
&litMatcher{
	pos: position{line: 518, col: 29, offset: 14920},
	val: "cov(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 518, col: 36, offset: 14927},
	name: "_",
},
&labeledExpr{
	pos: position{line: 518, col: 38, offset: 14929},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 518, col: 45, offset: 14936},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 518, col: 54, offset: 14945},
	name: "_",
},
&litMatcher{
	pos: position{line: 518, col: 56, offset: 14947},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 518, col: 60, offset: 14951},
	name: "_",
},
&labeledExpr{
	pos: position{line: 518, col: 62, offset: 14953},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 518, col: 69, offset: 14960},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 518, col: 78, offset: 14969},
	name: "_",
},
&labeledExpr{
	pos: position{line: 518, col: 80, offset: 14971},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 518, col: 85, offset: 14976},
	expr: &seqExpr{
	pos: position{line: 518, col: 86, offset: 14977},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 518, col: 86, offset: 14977},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 518, col: 90, offset: 14981},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 518, col: 92, offset: 14983},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 518, col: 106, offset: 14997},
	name: "_",
},
&litMatcher{
	pos: position{line: 518, col: 108, offset: 14999},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CorrAgg",
	pos: position{line: 541, col: 1, offset: 15524},
	expr: &actionExpr{
	pos: position{line: 541, col: 12, offset: 15535},
	run: (*parser).callonCorrAgg1,
	expr: &seqExpr{
	pos: position{line: 541, col: 12, offset: 15535},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 541, col: 12, offset: 15535},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 541, col: 18, offset: 15541},
	expr: &ruleRefExpr{
	pos: position{line: 541, col: 18, offset: 15541},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 541, col: 28, offset: 15551},
	name: "_",
},
&litMatcher{
	pos: position{line: 541, col: 30, offset: 15553},
	val: "correlation(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 541, col: 45, offset: 15568},
	name: "_",
},
&labeledExpr{
	pos: position{line: 541, col: 47, offset: 15570},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 541, col: 54, offset: 15577},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 541, col: 63, offset: 15586},
	name: "_",
},
&litMatcher{
	pos: position{line: 541, col: 65, offset: 15588},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 541, col: 69, offset: 15592},
	name: "_",
},
&labeledExpr{
	pos: position{line: 541, col: 71, offset: 15594},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 541, col: 78, offset: 15601},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 541, col: 87, offset: 15610},
	name: "_",
},
&labeledExpr{
	pos: position{line: 541, col: 89, offset: 15612},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 541, col: 94, offset: 15617},
	expr: &seqExpr{
	pos: position{line: 541, col: 95, offset: 15618},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 541, col: 95, offset: 15618},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 541, col: 99, offset: 15622},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 541, col: 101, offset: 15624},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 541, col: 115, offset: 15638},
	name: "_",
},
&litMatcher{
	pos: position{line: 541, col: 117, offset: 15640},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PValueAgg",
	pos: position{line: 564, col: 1, offset: 16166},
	expr: &actionExpr{
	pos: position{line: 564, col: 14, offset: 16179},
	run: (*parser).callonPValueAgg1,
	expr: &seqExpr{
	pos: position{line: 564, col: 14, offset: 16179},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 564, col: 14, offset: 16179},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 564, col: 20, offset: 16185},
	expr: &ruleRefExpr{
	pos: position{line: 564, col: 20, offset: 16185},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 564, col: 30, offset: 16195},
	name: "_",
},
&litMatcher{
	pos: position{line: 564, col: 32, offset: 16197},
	val: "pvalue(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 564, col: 42, offset: 16207},
	name: "_",
},
&labeledExpr{
	pos: position{line: 564, col: 44, offset: 16209},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 564, col: 51, offset: 16216},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 564, col: 60, offset: 16225},
	name: "_",
},
&litMatcher{
	pos: position{line: 564, col: 62, offset: 16227},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 564, col: 66, offset: 16231},
	name: "_",
},
&labeledExpr{
	pos: position{line: 564, col: 68, offset: 16233},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 564, col: 75, offset: 16240},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 564, col: 84, offset: 16249},
	name: "_",
},
&labeledExpr{
	pos: position{line: 564, col: 86, offset: 16251},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 564, col: 91, offset: 16256},
	expr: &seqExpr{
	pos: position{line: 564, col: 92, offset: 16257},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 564, col: 92, offset: 16257},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 564, col: 96, offset: 16261},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 564, col: 98, offset: 16263},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 564, col: 112, offset: 16277},
	name: "_",
},
&litMatcher{
	pos: position{line: 564, col: 114, offset: 16279},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggGroupBy",
	pos: position{line: 589, col: 1, offset: 16804},
	expr: &actionExpr{
	pos: position{line: 589, col: 15, offset: 16818},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 589, col: 15, offset: 16818},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 589, col: 15, offset: 16818},
	name: "_",
},
&litMatcher{
	pos: position{line: 589, col: 17, offset: 16820},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 589, col: 22, offset: 16825},
	name: "_",
},
&labeledExpr{
	pos: position{line: 589, col: 24, offset: 16827},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 589, col: 30, offset: 16833},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 589, col: 39, offset: 16842},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 589, col: 44, offset: 16847},
	expr: &seqExpr{
	pos: position{line: 589, col: 45, offset: 16848},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 589, col: 45, offset: 16848},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 589, col: 49, offset: 16852},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 589, col: 51, offset: 16854},
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
	pos: position{line: 604, col: 1, offset: 17205},
	expr: &actionExpr{
	pos: position{line: 604, col: 11, offset: 17215},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 604, col: 11, offset: 17215},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 604, col: 16, offset: 17220},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 604, col: 16, offset: 17220},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 604, col: 27, offset: 17231},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 604, col: 38, offset: 17242},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 604, col: 49, offset: 17253},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 604, col: 58, offset: 17262},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 604, col: 67, offset: 17271},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 610, col: 1, offset: 17390},
	expr: &actionExpr{
	pos: position{line: 610, col: 13, offset: 17402},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 610, col: 13, offset: 17402},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 610, col: 13, offset: 17402},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 610, col: 23, offset: 17412},
	name: "_",
},
&labeledExpr{
	pos: position{line: 610, col: 25, offset: 17414},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 610, col: 28, offset: 17417},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 610, col: 42, offset: 17431},
	name: "_",
},
&litMatcher{
	pos: position{line: 610, col: 44, offset: 17433},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 617, col: 1, offset: 17628},
	expr: &actionExpr{
	pos: position{line: 617, col: 18, offset: 17645},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 617, col: 18, offset: 17645},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 617, col: 18, offset: 17645},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 617, col: 23, offset: 17650},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 617, col: 28, offset: 17655},
	name: "_",
},
&labeledExpr{
	pos: position{line: 617, col: 30, offset: 17657},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 617, col: 33, offset: 17660},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 617, col: 45, offset: 17672},
	name: "_",
},
&labeledExpr{
	pos: position{line: 617, col: 47, offset: 17674},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 617, col: 54, offset: 17681},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 621, col: 1, offset: 17732},
	expr: &ruleRefExpr{
	pos: position{line: 621, col: 16, offset: 17747},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 624, col: 1, offset: 17781},
	expr: &actionExpr{
	pos: position{line: 624, col: 13, offset: 17793},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 624, col: 13, offset: 17793},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 624, col: 13, offset: 17793},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 624, col: 23, offset: 17803},
	name: "_",
},
&labeledExpr{
	pos: position{line: 624, col: 25, offset: 17805},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 624, col: 31, offset: 17811},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 624, col: 40, offset: 17820},
	name: "_",
},
&labeledExpr{
	pos: position{line: 624, col: 42, offset: 17822},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 624, col: 47, offset: 17827},
	expr: &seqExpr{
	pos: position{line: 624, col: 48, offset: 17828},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 624, col: 48, offset: 17828},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 624, col: 52, offset: 17832},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 624, col: 54, offset: 17834},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 624, col: 65, offset: 17845},
	name: "_",
},
&litMatcher{
	pos: position{line: 624, col: 67, offset: 17847},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 639, col: 1, offset: 18184},
	expr: &actionExpr{
	pos: position{line: 639, col: 11, offset: 18194},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 639, col: 11, offset: 18194},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 639, col: 11, offset: 18194},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 639, col: 18, offset: 18201},
	name: "_",
},
&labeledExpr{
	pos: position{line: 639, col: 20, offset: 18203},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 639, col: 25, offset: 18208},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 639, col: 34, offset: 18217},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 639, col: 38, offset: 18221},
	name: "_",
},
&labeledExpr{
	pos: position{line: 639, col: 40, offset: 18223},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 639, col: 44, offset: 18227},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 639, col: 51, offset: 18234},
	name: "_",
},
&litMatcher{
	pos: position{line: 639, col: 53, offset: 18236},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 651, col: 1, offset: 18458},
	expr: &actionExpr{
	pos: position{line: 651, col: 11, offset: 18468},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 651, col: 11, offset: 18468},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 651, col: 11, offset: 18468},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 651, col: 20, offset: 18477},
	name: "_",
},
&litMatcher{
	pos: position{line: 651, col: 22, offset: 18479},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 651, col: 27, offset: 18484},
	name: "_",
},
&labeledExpr{
	pos: position{line: 651, col: 29, offset: 18486},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 651, col: 35, offset: 18492},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 657, col: 1, offset: 18598},
	expr: &actionExpr{
	pos: position{line: 657, col: 12, offset: 18609},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 657, col: 12, offset: 18609},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 657, col: 12, offset: 18609},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 657, col: 21, offset: 18618},
	name: "_",
},
&labeledExpr{
	pos: position{line: 657, col: 23, offset: 18620},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 657, col: 27, offset: 18624},
	expr: &ruleRefExpr{
	pos: position{line: 657, col: 27, offset: 18624},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 657, col: 35, offset: 18632},
	name: "_",
},
&litMatcher{
	pos: position{line: 657, col: 37, offset: 18634},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 667, col: 1, offset: 18779},
	expr: &actionExpr{
	pos: position{line: 667, col: 13, offset: 18791},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 667, col: 13, offset: 18791},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 667, col: 13, offset: 18791},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 667, col: 23, offset: 18801},
	name: "_",
},
&labeledExpr{
	pos: position{line: 667, col: 25, offset: 18803},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 667, col: 30, offset: 18808},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 667, col: 42, offset: 18820},
	name: "_",
},
&litMatcher{
	pos: position{line: 667, col: 44, offset: 18822},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 677, col: 1, offset: 19110},
	expr: &actionExpr{
	pos: position{line: 677, col: 16, offset: 19125},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 677, col: 16, offset: 19125},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 677, col: 16, offset: 19125},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 677, col: 22, offset: 19131},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 677, col: 35, offset: 19144},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 677, col: 40, offset: 19149},
	expr: &seqExpr{
	pos: position{line: 677, col: 41, offset: 19150},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 677, col: 41, offset: 19150},
	name: "_",
},
&labeledExpr{
	pos: position{line: 677, col: 43, offset: 19152},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 677, col: 46, offset: 19155},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 677, col: 58, offset: 19167},
	name: "_",
},
&labeledExpr{
	pos: position{line: 677, col: 60, offset: 19169},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 677, col: 63, offset: 19172},
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
	pos: position{line: 682, col: 1, offset: 19330},
	expr: &choiceExpr{
	pos: position{line: 682, col: 17, offset: 19346},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 682, col: 17, offset: 19346},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 682, col: 17, offset: 19346},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 682, col: 17, offset: 19346},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 682, col: 21, offset: 19350},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 682, col: 24, offset: 19353},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 682, col: 36, offset: 19365},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 684, col: 5, offset: 19396},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 684, col: 5, offset: 19396},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 684, col: 8, offset: 19399},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 688, col: 1, offset: 19441},
	expr: &choiceExpr{
	pos: position{line: 688, col: 16, offset: 19456},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 688, col: 16, offset: 19456},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 688, col: 24, offset: 19464},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 688, col: 24, offset: 19464},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 695, col: 1, offset: 19643},
	expr: &actionExpr{
	pos: position{line: 695, col: 18, offset: 19660},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 695, col: 18, offset: 19660},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 695, col: 18, offset: 19660},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 695, col: 23, offset: 19665},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 695, col: 28, offset: 19670},
	name: "_",
},
&labeledExpr{
	pos: position{line: 695, col: 30, offset: 19672},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 695, col: 33, offset: 19675},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 695, col: 45, offset: 19687},
	name: "_",
},
&labeledExpr{
	pos: position{line: 695, col: 47, offset: 19689},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 695, col: 53, offset: 19695},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 695, col: 59, offset: 19701},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 699, col: 1, offset: 19775},
	expr: &choiceExpr{
	pos: position{line: 699, col: 16, offset: 19790},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 699, col: 16, offset: 19790},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 699, col: 23, offset: 19797},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 699, col: 30, offset: 19804},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 699, col: 36, offset: 19810},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 699, col: 42, offset: 19816},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 699, col: 49, offset: 19823},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 699, col: 49, offset: 19823},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 708, col: 1, offset: 19933},
	expr: &choiceExpr{
	pos: position{line: 708, col: 12, offset: 19944},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 708, col: 12, offset: 19944},
	run: (*parser).callonStdOutG2,
	expr: &seqExpr{
	pos: position{line: 708, col: 12, offset: 19944},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 708, col: 12, offset: 19944},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 708, col: 22, offset: 19954},
	name: "_",
},
&labeledExpr{
	pos: position{line: 708, col: 24, offset: 19956},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 708, col: 31, offset: 19963},
	expr: &ruleRefExpr{
	pos: position{line: 708, col: 32, offset: 19964},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 708, col: 48, offset: 19980},
	name: "_",
},
&litMatcher{
	pos: position{line: 708, col: 50, offset: 19982},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 720, col: 3, offset: 20258},
	run: (*parser).callonStdOutG11,
	expr: &litMatcher{
	pos: position{line: 720, col: 3, offset: 20258},
	val: "stdout",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "SinkInto",
	pos: position{line: 724, col: 1, offset: 20328},
	expr: &choiceExpr{
	pos: position{line: 724, col: 12, offset: 20339},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 724, col: 12, offset: 20339},
	run: (*parser).callonSinkInto2,
	expr: &seqExpr{
	pos: position{line: 724, col: 12, offset: 20339},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 724, col: 12, offset: 20339},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 724, col: 18, offset: 20345},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 724, col: 27, offset: 20354},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 724, col: 33, offset: 20360},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 724, col: 33, offset: 20360},
	name: "_",
},
&litMatcher{
	pos: position{line: 724, col: 35, offset: 20362},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 724, col: 39, offset: 20366},
	name: "_",
},
&litMatcher{
	pos: position{line: 724, col: 41, offset: 20368},
	val: "into()",
	ignoreCase: false,
},
	},
},
},
	},
},
},
&actionExpr{
	pos: position{line: 730, col: 7, offset: 20521},
	run: (*parser).callonSinkInto12,
	expr: &seqExpr{
	pos: position{line: 730, col: 7, offset: 20521},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 730, col: 7, offset: 20521},
	val: "into",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 730, col: 14, offset: 20528},
	name: "_",
},
&litMatcher{
	pos: position{line: 730, col: 16, offset: 20530},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 730, col: 20, offset: 20534},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 730, col: 22, offset: 20536},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 730, col: 31, offset: 20545},
	name: "_",
},
&litMatcher{
	pos: position{line: 730, col: 33, offset: 20547},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 732, col: 7, offset: 20643},
	run: (*parser).callonSinkInto21,
	expr: &seqExpr{
	pos: position{line: 732, col: 7, offset: 20643},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 732, col: 7, offset: 20643},
	val: "into",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 732, col: 14, offset: 20650},
	name: "_",
},
&litMatcher{
	pos: position{line: 732, col: 16, offset: 20652},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 732, col: 20, offset: 20656},
	name: "_",
},
&litMatcher{
	pos: position{line: 732, col: 22, offset: 20658},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 734, col: 7, offset: 20797},
	run: (*parser).callonSinkInto28,
	expr: &seqExpr{
	pos: position{line: 734, col: 7, offset: 20797},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 734, col: 7, offset: 20797},
	val: "into",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 734, col: 14, offset: 20804},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 736, col: 7, offset: 20897},
	run: (*parser).callonSinkInto32,
	expr: &seqExpr{
	pos: position{line: 736, col: 7, offset: 20897},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 736, col: 7, offset: 20897},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 736, col: 13, offset: 20903},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 736, col: 22, offset: 20912},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 736, col: 29, offset: 20919},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 736, col: 29, offset: 20919},
	name: "_",
},
&litMatcher{
	pos: position{line: 736, col: 31, offset: 20921},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 736, col: 35, offset: 20925},
	name: "_",
},
&litMatcher{
	pos: position{line: 736, col: 37, offset: 20927},
	val: "into",
	ignoreCase: false,
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
	name: "BlackHoleG",
	pos: position{line: 740, col: 1, offset: 20998},
	expr: &choiceExpr{
	pos: position{line: 740, col: 15, offset: 21012},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 740, col: 15, offset: 21012},
	run: (*parser).callonBlackHoleG2,
	expr: &seqExpr{
	pos: position{line: 740, col: 15, offset: 21012},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 740, col: 15, offset: 21012},
	val: "blackhole(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 740, col: 28, offset: 21025},
	name: "_",
},
&labeledExpr{
	pos: position{line: 740, col: 30, offset: 21027},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 740, col: 37, offset: 21034},
	expr: &ruleRefExpr{
	pos: position{line: 740, col: 38, offset: 21035},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 740, col: 54, offset: 21051},
	name: "_",
},
&litMatcher{
	pos: position{line: 740, col: 56, offset: 21053},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 742, col: 3, offset: 21108},
	run: (*parser).callonBlackHoleG11,
	expr: &seqExpr{
	pos: position{line: 742, col: 3, offset: 21108},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 742, col: 3, offset: 21108},
	val: "blackhole",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 742, col: 15, offset: 21120},
	name: "_",
},
	},
},
},
	},
},
},
{
	name: "PlotG",
	pos: position{line: 746, col: 1, offset: 21205},
	expr: &actionExpr{
	pos: position{line: 746, col: 10, offset: 21214},
	run: (*parser).callonPlotG1,
	expr: &seqExpr{
	pos: position{line: 746, col: 10, offset: 21214},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 746, col: 10, offset: 21214},
	val: "plot",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 746, col: 17, offset: 21221},
	name: "_",
},
&labeledExpr{
	pos: position{line: 746, col: 19, offset: 21223},
	label: "widget",
	expr: &ruleRefExpr{
	pos: position{line: 746, col: 27, offset: 21231},
	name: "BarWidget",
},
},
	},
},
},
},
{
	name: "BarWidget",
	pos: position{line: 750, col: 1, offset: 21331},
	expr: &actionExpr{
	pos: position{line: 750, col: 14, offset: 21344},
	run: (*parser).callonBarWidget1,
	expr: &seqExpr{
	pos: position{line: 750, col: 14, offset: 21344},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 750, col: 14, offset: 21344},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 750, col: 20, offset: 21350},
	expr: &ruleRefExpr{
	pos: position{line: 750, col: 20, offset: 21350},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 750, col: 30, offset: 21360},
	name: "_",
},
&litMatcher{
	pos: position{line: 750, col: 32, offset: 21362},
	val: "bar(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 750, col: 39, offset: 21369},
	name: "_",
},
&labeledExpr{
	pos: position{line: 750, col: 41, offset: 21371},
	label: "xField",
	expr: &ruleRefExpr{
	pos: position{line: 750, col: 48, offset: 21378},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 750, col: 57, offset: 21387},
	name: "_",
},
&litMatcher{
	pos: position{line: 750, col: 59, offset: 21389},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 750, col: 63, offset: 21393},
	name: "_",
},
&labeledExpr{
	pos: position{line: 750, col: 65, offset: 21395},
	label: "yField",
	expr: &ruleRefExpr{
	pos: position{line: 750, col: 72, offset: 21402},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 750, col: 81, offset: 21411},
	name: "_",
},
&labeledExpr{
	pos: position{line: 750, col: 83, offset: 21413},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 750, col: 88, offset: 21418},
	expr: &seqExpr{
	pos: position{line: 750, col: 89, offset: 21419},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 750, col: 89, offset: 21419},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 750, col: 93, offset: 21423},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 750, col: 95, offset: 21425},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 750, col: 102, offset: 21432},
	name: "_",
},
&zeroOrOneExpr{
	pos: position{line: 750, col: 104, offset: 21434},
	expr: &seqExpr{
	pos: position{line: 750, col: 105, offset: 21435},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 750, col: 105, offset: 21435},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 750, col: 109, offset: 21439},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 750, col: 111, offset: 21441},
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
	pos: position{line: 750, col: 122, offset: 21452},
	name: "_",
},
&litMatcher{
	pos: position{line: 750, col: 124, offset: 21454},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FileName",
	pos: position{line: 783, col: 1, offset: 22139},
	expr: &actionExpr{
	pos: position{line: 783, col: 13, offset: 22151},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 783, col: 13, offset: 22151},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 783, col: 13, offset: 22151},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 783, col: 22, offset: 22160},
	expr: &seqExpr{
	pos: position{line: 783, col: 23, offset: 22161},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 783, col: 23, offset: 22161},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 783, col: 27, offset: 22165},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 783, col: 31, offset: 22169},
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
	pos: position{line: 788, col: 1, offset: 22268},
	expr: &actionExpr{
	pos: position{line: 788, col: 10, offset: 22277},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 788, col: 10, offset: 22277},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 788, col: 10, offset: 22277},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 788, col: 16, offset: 22283},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 788, col: 23, offset: 22290},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 788, col: 28, offset: 22295},
	expr: &seqExpr{
	pos: position{line: 788, col: 29, offset: 22296},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 788, col: 29, offset: 22296},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 788, col: 31, offset: 22298},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 788, col: 36, offset: 22303},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 788, col: 38, offset: 22305},
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
	pos: position{line: 792, col: 1, offset: 22370},
	expr: &actionExpr{
	pos: position{line: 792, col: 11, offset: 22380},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 792, col: 11, offset: 22380},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 792, col: 16, offset: 22385},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 792, col: 16, offset: 22385},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 792, col: 24, offset: 22393},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 792, col: 33, offset: 22402},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 792, col: 48, offset: 22417},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 792, col: 59, offset: 22428},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 796, col: 1, offset: 22470},
	expr: &choiceExpr{
	pos: position{line: 796, col: 9, offset: 22478},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 796, col: 9, offset: 22478},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 796, col: 15, offset: 22484},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 796, col: 21, offset: 22490},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 796, col: 27, offset: 22496},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 796, col: 33, offset: 22502},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 796, col: 33, offset: 22502},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 800, col: 1, offset: 22550},
	expr: &choiceExpr{
	pos: position{line: 800, col: 12, offset: 22561},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 800, col: 12, offset: 22561},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 800, col: 12, offset: 22561},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 800, col: 12, offset: 22561},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 800, col: 16, offset: 22565},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 800, col: 19, offset: 22568},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 800, col: 25, offset: 22574},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 804, col: 5, offset: 22650},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 804, col: 5, offset: 22650},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 804, col: 8, offset: 22653},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 811, col: 1, offset: 22730},
	expr: &actionExpr{
	pos: position{line: 811, col: 9, offset: 22738},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 811, col: 9, offset: 22738},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 811, col: 9, offset: 22738},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 811, col: 15, offset: 22744},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 811, col: 21, offset: 22750},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 811, col: 26, offset: 22755},
	expr: &seqExpr{
	pos: position{line: 811, col: 27, offset: 22756},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 811, col: 27, offset: 22756},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 811, col: 29, offset: 22758},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 811, col: 34, offset: 22763},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 811, col: 36, offset: 22765},
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
	pos: position{line: 815, col: 1, offset: 22829},
	expr: &actionExpr{
	pos: position{line: 815, col: 10, offset: 22838},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 815, col: 10, offset: 22838},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 815, col: 15, offset: 22843},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 815, col: 15, offset: 22843},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 815, col: 23, offset: 22851},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 815, col: 32, offset: 22860},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 815, col: 47, offset: 22875},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 819, col: 1, offset: 22918},
	expr: &actionExpr{
	pos: position{line: 819, col: 13, offset: 22930},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 819, col: 13, offset: 22930},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 819, col: 13, offset: 22930},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 819, col: 20, offset: 22937},
	expr: &choiceExpr{
	pos: position{line: 819, col: 21, offset: 22938},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 819, col: 21, offset: 22938},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 819, col: 31, offset: 22948},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 819, col: 39, offset: 22956},
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
	pos: position{line: 823, col: 1, offset: 23006},
	expr: &actionExpr{
	pos: position{line: 823, col: 17, offset: 23022},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 823, col: 17, offset: 23022},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 823, col: 17, offset: 23022},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 823, col: 21, offset: 23026},
	expr: &choiceExpr{
	pos: position{line: 823, col: 23, offset: 23028},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 823, col: 23, offset: 23028},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 823, col: 23, offset: 23028},
	expr: &ruleRefExpr{
	pos: position{line: 823, col: 24, offset: 23029},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 823, col: 36, offset: 23041,
},
	},
},
&seqExpr{
	pos: position{line: 823, col: 40, offset: 23045},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 823, col: 40, offset: 23045},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 823, col: 45, offset: 23050},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 823, col: 63, offset: 23068},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 830, col: 1, offset: 23249},
	expr: &charClassMatcher{
	pos: position{line: 830, col: 16, offset: 23264},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 832, col: 1, offset: 23282},
	expr: &ruleRefExpr{
	pos: position{line: 832, col: 19, offset: 23300},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 834, col: 1, offset: 23320},
	expr: &charClassMatcher{
	pos: position{line: 834, col: 21, offset: 23340},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 836, col: 1, offset: 23355},
	expr: &actionExpr{
	pos: position{line: 836, col: 9, offset: 23363},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 836, col: 9, offset: 23363},
	expr: &choiceExpr{
	pos: position{line: 836, col: 10, offset: 23364},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 836, col: 10, offset: 23364},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 836, col: 19, offset: 23373},
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
	pos: position{line: 840, col: 1, offset: 23423},
	expr: &choiceExpr{
	pos: position{line: 840, col: 11, offset: 23433},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 840, col: 11, offset: 23433},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 840, col: 19, offset: 23441},
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
	pos: position{line: 842, col: 1, offset: 23450},
	expr: &actionExpr{
	pos: position{line: 842, col: 11, offset: 23460},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 842, col: 11, offset: 23460},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 842, col: 11, offset: 23460},
	expr: &litMatcher{
	pos: position{line: 842, col: 11, offset: 23460},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 842, col: 16, offset: 23465},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 846, col: 1, offset: 23538},
	expr: &actionExpr{
	pos: position{line: 846, col: 19, offset: 23556},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 846, col: 19, offset: 23556},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 850, col: 1, offset: 23629},
	expr: &actionExpr{
	pos: position{line: 850, col: 10, offset: 23638},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 850, col: 10, offset: 23638},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 850, col: 10, offset: 23638},
	expr: &litMatcher{
	pos: position{line: 850, col: 10, offset: 23638},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 850, col: 15, offset: 23643},
	expr: &ruleRefExpr{
	pos: position{line: 850, col: 15, offset: 23643},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 850, col: 24, offset: 23652},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 850, col: 28, offset: 23656},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 854, col: 1, offset: 23723},
	expr: &choiceExpr{
	pos: position{line: 854, col: 12, offset: 23734},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 854, col: 12, offset: 23734},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 854, col: 18, offset: 23740},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 854, col: 18, offset: 23740},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 854, col: 38, offset: 23760},
	expr: &ruleRefExpr{
	pos: position{line: 854, col: 38, offset: 23760},
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
	pos: position{line: 856, col: 1, offset: 23777},
	expr: &charClassMatcher{
	pos: position{line: 856, col: 17, offset: 23793},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 858, col: 1, offset: 23802},
	expr: &charClassMatcher{
	pos: position{line: 858, col: 24, offset: 23825},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 860, col: 1, offset: 23834},
	expr: &zeroOrMoreExpr{
	pos: position{line: 860, col: 19, offset: 23852},
	expr: &charClassMatcher{
	pos: position{line: 860, col: 19, offset: 23852},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 862, col: 1, offset: 23866},
	expr: &charClassMatcher{
	pos: position{line: 862, col: 18, offset: 23883},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "IDENTIFIER",
	pos: position{line: 863, col: 1, offset: 23894},
	expr: &actionExpr{
	pos: position{line: 863, col: 15, offset: 23908},
	run: (*parser).callonIDENTIFIER1,
	expr: &seqExpr{
	pos: position{line: 863, col: 15, offset: 23908},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 863, col: 15, offset: 23908},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 863, col: 22, offset: 23915},
	expr: &choiceExpr{
	pos: position{line: 863, col: 23, offset: 23916},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 863, col: 23, offset: 23916},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 863, col: 33, offset: 23926},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 863, col: 41, offset: 23934},
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
	pos: position{line: 866, col: 1, offset: 24007},
	expr: &notExpr{
	pos: position{line: 866, col: 8, offset: 24014},
	expr: &anyMatcher{
	line: 866, col: 9, offset: 24015,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 868, col: 1, offset: 24020},
	expr: &litMatcher{
	pos: position{line: 868, col: 17, offset: 24036},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 874, col: 1, offset: 24123},
	expr: &actionExpr{
	pos: position{line: 874, col: 18, offset: 24140},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 874, col: 18, offset: 24140},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 874, col: 18, offset: 24140},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 874, col: 24, offset: 24146},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 874, col: 33, offset: 24155},
	name: "_",
},
&labeledExpr{
	pos: position{line: 874, col: 35, offset: 24157},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 874, col: 40, offset: 24162},
	expr: &seqExpr{
	pos: position{line: 874, col: 41, offset: 24163},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 874, col: 41, offset: 24163},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 874, col: 45, offset: 24167},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 874, col: 47, offset: 24169},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 874, col: 56, offset: 24178},
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
	pos: position{line: 877, col: 1, offset: 24231},
	expr: &actionExpr{
	pos: position{line: 877, col: 13, offset: 24243},
	run: (*parser).callonEqAliasG1,
	expr: &seqExpr{
	pos: position{line: 877, col: 13, offset: 24243},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 877, col: 13, offset: 24243},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 877, col: 19, offset: 24249},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 877, col: 28, offset: 24258},
	name: "_",
},
&litMatcher{
	pos: position{line: 877, col: 30, offset: 24260},
	val: "=",
	ignoreCase: false,
},
	},
},
},
},
	},
}
func (c *current) onCommand2(command, rest interface{}) (interface{}, error) {

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

func (p *parser) callonCommand2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCommand2(stack["command"], stack["rest"])
}

func (c *current) onCommand14(command, rest interface{}) (interface{}, error) {

        return nil, errors.New("Missing ; at the end of Command")
    
}

func (p *parser) callonCommand14() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCommand14(stack["command"], stack["rest"])
}

func (c *current) onStatement2(statement, rest interface{}) (interface{}, error) {

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

func (p *parser) callonStatement2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement2(stack["statement"], stack["rest"])
}

func (c *current) onStatement23(statement, rest interface{}) (interface{}, error) {

        return nil,errors.New("In Statement: Missing ; at the end of statement")
    
}

func (p *parser) callonStatement23() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement23(stack["statement"], stack["rest"])
}

func (c *current) onStatement42(statement, rest interface{}) (interface{}, error) {

        return nil,errors.New("statement should end with a sink")
    
}

func (p *parser) callonStatement42() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement42(stack["statement"], stack["rest"])
}

func (c *current) onFileJobG2(filename interface{}) (interface{}, error) {

    if fname, ok := cast.TryString(filename); ok {
        return SourceJob{Type: FILEJOB,OperateOn:SrcFile{FileName:fname}},nil
    }
    return nil, errors.New("could not decode source file name")
}

func (p *parser) callonFileJobG2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFileJobG2(stack["filename"])
}

func (c *current) onFileJobG11() (interface{}, error) {

    return nil,errors.New("File Name Missing")
}

func (p *parser) callonFileJobG11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFileJobG11()
}

func (c *current) onFakeJobG2(numData interface{}) (interface{}, error) {

    if n, ok := cast.TryInt(numData); ok {
       return SourceJob{Type:FAKEJOB,OperateOn:SrcFake{Number:n}},nil
    }
    return nil, errors.New("could not decode number")
}

func (p *parser) callonFakeJobG2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFakeJobG2(stack["numData"])
}

func (c *current) onFakeJobG13() (interface{}, error) {

    return nil,errors.New("Mention the number of data to be generated")
}

func (p *parser) callonFakeJobG13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFakeJobG13()
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

func (c *current) onSecSource2(src1, rest interface{}) (interface{}, error) {

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

func (p *parser) callonSecSource2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSecSource2(stack["src1"], stack["rest"])
}

func (c *current) onSecSource14(src1, rest interface{}) (interface{}, error) {

    return nil,errors.New("comma expected to separate seconday sources")
}

func (p *parser) callonSecSource14() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSecSource14(stack["src1"], stack["rest"])
}

func (c *current) onJoinJobG1(job interface{}) (interface{}, error) {

    return TransformJob{Type:JOINJOB,OperateOn:job},nil
}

func (p *parser) callonJoinJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onJoinJobG1(stack["job"])
}

func (c *current) onInnerJoinG2(alias, query interface{}) (interface{}, error) {

     name,err := parseJoinArgs(alias)
     if err !=nil{
         return nil,err
         }
     return JoinNodeJob{Alias:name,Type:"INNER",Attributes:query.(JoinAttributes)},nil
}

func (p *parser) callonInnerJoinG2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInnerJoinG2(stack["alias"], stack["query"])
}

func (c *current) onInnerJoinG16(alias interface{}) (interface{}, error) {

return nil,errors.New("Expected query after (")
}

func (p *parser) callonInnerJoinG16() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInnerJoinG16(stack["alias"])
}

func (c *current) onLeftOuterG2(alias, query interface{}) (interface{}, error) {

                name,err := parseJoinArgs(alias)
                if err !=nil{
                   return nil,err
                   }
         return JoinNodeJob{Alias:name,Type:"OUTER",Attributes:query.(JoinAttributes),JoinSubType:"LEFTOUTER"},nil
         
}

func (p *parser) callonLeftOuterG2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLeftOuterG2(stack["alias"], stack["query"])
}

func (c *current) onLeftOuterG16(alias interface{}) (interface{}, error) {

         return nil, errors.New("Expected query after (")
         
}

func (p *parser) callonLeftOuterG16() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLeftOuterG16(stack["alias"])
}

func (c *current) onRightOuterG2(alias, query interface{}) (interface{}, error) {

              name,err := parseJoinArgs(alias)
              if err !=nil{
                  return nil,err
                  }
              return JoinNodeJob{Alias:name,Type:"OUTER",Attributes:query.(JoinAttributes),JoinSubType:"RIGHTOUTER"},nil
              
}

func (p *parser) callonRightOuterG2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRightOuterG2(stack["alias"], stack["query"])
}

func (c *current) onRightOuterG15(alias interface{}) (interface{}, error) {

                        return nil, errors.New("Expected query after (")
                        
}

func (p *parser) callonRightOuterG15() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRightOuterG15(stack["alias"])
}

func (c *current) onFullOuterG2(alias, query interface{}) (interface{}, error) {

              name,err := parseJoinArgs(alias)
              if err !=nil{
                    return nil,err
                 }
              return JoinNodeJob{Alias:name,Type:"OUTER",Attributes:query.(JoinAttributes),JoinSubType:"RIGHTOUTER"},nil
              
}

func (p *parser) callonFullOuterG2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFullOuterG2(stack["alias"], stack["query"])
}

func (c *current) onFullOuterG15(alias interface{}) (interface{}, error) {

                 return nil, errors.New("Expected query after (")
              
}

func (p *parser) callonFullOuterG15() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFullOuterG15(stack["alias"])
}

func (c *current) onQuery2(field, condition interface{}) (interface{}, error) {

    return JoinAttributes{SelectFields:field.([]string),JoinCondition:condition.(JoinConditions)},nil
}

func (p *parser) callonQuery2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQuery2(stack["field"], stack["condition"])
}

func (c *current) onQuery14(condition interface{}) (interface{}, error) {

    return nil, errors.New("select fields missing. use * for all fields")
}

func (p *parser) callonQuery14() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQuery14(stack["condition"])
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

func (c *current) onJoinCondition2(left, op, right interface{}) (interface{}, error) {

        return JoinConditions{LeftFields:left.([]string),RightFields:right.([]string),JoinOperator:"=="},nil
}

func (p *parser) callonJoinCondition2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onJoinCondition2(stack["left"], stack["op"], stack["right"])
}

func (c *current) onJoinCondition13(left, right interface{}) (interface{}, error) {

    return nil, errors.New("missing operator in between the join factors")
}

func (p *parser) callonJoinCondition13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onJoinCondition13(stack["left"], stack["right"])
}

func (c *current) onJoinCondition21(left, op interface{}) (interface{}, error) {

    return nil, errors.New("missing right factors of the join condition")
}

func (p *parser) callonJoinCondition21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onJoinCondition21(stack["left"], stack["op"])
}

func (c *current) onJoinCondition29(left interface{}) (interface{}, error) {

    return nil, errors.New("missing operator and right factor in join condition")
}

func (p *parser) callonJoinCondition29() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onJoinCondition29(stack["left"])
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

func (c *current) onStdOutG2(params interface{}) (interface{}, error) {


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

func (p *parser) callonStdOutG2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStdOutG2(stack["params"])
}

func (c *current) onStdOutG11() (interface{}, error) {

    return nil,errors.New("missing () after stdout")
}

func (p *parser) callonStdOutG11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStdOutG11()
}

func (c *current) onSinkInto2(first, rest interface{}) (interface{}, error) {

    streamTo,_ := first.(string)
    return SinkJob{
        Type: INTO,
        OperateOn:InTo{StreamTo:streamTo},
        },nil
    
}

func (p *parser) callonSinkInto2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSinkInto2(stack["first"], stack["rest"])
}

func (c *current) onSinkInto12() (interface{}, error) {

    return nil, errors.New("improper syntax, proper syntax is Variable = into();")
    
}

func (p *parser) callonSinkInto12() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSinkInto12()
}

func (c *current) onSinkInto21() (interface{}, error) {

    return nil, errors.New("improper syntax, secondary sink name not given; syntax is sink name of variable type = into() ;")
    
}

func (p *parser) callonSinkInto21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSinkInto21()
}

func (c *current) onSinkInto28() (interface{}, error) {

    return nil,errors.New("improper syntax,proper syntax is Variable = into();")
    
}

func (p *parser) callonSinkInto28() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSinkInto28()
}

func (c *current) onSinkInto32(first, rest interface{}) (interface{}, error) {

    return nil,errors.New("missing () after into")
    
}

func (p *parser) callonSinkInto32() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSinkInto32(stack["first"], stack["rest"])
}

func (c *current) onBlackHoleG2(params interface{}) (interface{}, error) {

        return SinkJob{Type: BLACKHOLE}, nil
}

func (p *parser) callonBlackHoleG2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlackHoleG2(stack["params"])
}

func (c *current) onBlackHoleG11() (interface{}, error) {

    return nil,errors.New("improper syntax,proper syntax is blackhole();")
}

func (p *parser) callonBlackHoleG11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlackHoleG11()
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

