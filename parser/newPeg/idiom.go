
package newPeg
import(
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
&ruleRefExpr{
	pos: position{line: 79, col: 38, offset: 2232},
	name: "MapJobG",
},
	},
},
},
{
	name: "Sink",
	pos: position{line: 81, col: 1, offset: 2243},
	expr: &choiceExpr{
	pos: position{line: 81, col: 9, offset: 2251},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 81, col: 9, offset: 2251},
	name: "StdOutG",
},
&ruleRefExpr{
	pos: position{line: 81, col: 17, offset: 2259},
	name: "SinkInto",
},
&ruleRefExpr{
	pos: position{line: 81, col: 26, offset: 2268},
	name: "PlotG",
},
	},
},
},
{
	name: "FileJobG",
	pos: position{line: 85, col: 1, offset: 2294},
	expr: &choiceExpr{
	pos: position{line: 85, col: 13, offset: 2306},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 85, col: 13, offset: 2306},
	run: (*parser).callonFileJobG2,
	expr: &seqExpr{
	pos: position{line: 85, col: 13, offset: 2306},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 85, col: 13, offset: 2306},
	name: "_",
},
&litMatcher{
	pos: position{line: 85, col: 15, offset: 2308},
	val: "file(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 85, col: 23, offset: 2316},
	name: "_",
},
&labeledExpr{
	pos: position{line: 85, col: 25, offset: 2318},
	label: "filename",
	expr: &ruleRefExpr{
	pos: position{line: 85, col: 34, offset: 2327},
	name: "FileName",
},
},
&ruleRefExpr{
	pos: position{line: 85, col: 43, offset: 2336},
	name: "_",
},
&litMatcher{
	pos: position{line: 85, col: 45, offset: 2338},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 90, col: 4, offset: 2551},
	run: (*parser).callonFileJobG11,
	expr: &seqExpr{
	pos: position{line: 90, col: 4, offset: 2551},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 90, col: 4, offset: 2551},
	name: "_",
},
&litMatcher{
	pos: position{line: 90, col: 6, offset: 2553},
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
	pos: position{line: 94, col: 1, offset: 2617},
	expr: &choiceExpr{
	pos: position{line: 94, col: 13, offset: 2629},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 94, col: 13, offset: 2629},
	run: (*parser).callonFakeJobG2,
	expr: &seqExpr{
	pos: position{line: 94, col: 13, offset: 2629},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 94, col: 13, offset: 2629},
	name: "_",
},
&litMatcher{
	pos: position{line: 94, col: 15, offset: 2631},
	val: "fake",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 94, col: 21, offset: 2637},
	name: "_",
},
&litMatcher{
	pos: position{line: 94, col: 23, offset: 2639},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 94, col: 27, offset: 2643},
	name: "_",
},
&labeledExpr{
	pos: position{line: 94, col: 29, offset: 2645},
	label: "numData",
	expr: &ruleRefExpr{
	pos: position{line: 94, col: 37, offset: 2653},
	name: "PositiveNumber",
},
},
&ruleRefExpr{
	pos: position{line: 94, col: 52, offset: 2668},
	name: "_",
},
&litMatcher{
	pos: position{line: 94, col: 54, offset: 2670},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 99, col: 3, offset: 2856},
	run: (*parser).callonFakeJobG13,
	expr: &seqExpr{
	pos: position{line: 99, col: 3, offset: 2856},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 99, col: 3, offset: 2856},
	name: "_",
},
&litMatcher{
	pos: position{line: 99, col: 5, offset: 2858},
	val: "fake",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 99, col: 11, offset: 2864},
	name: "_",
},
&litMatcher{
	pos: position{line: 99, col: 13, offset: 2866},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 99, col: 17, offset: 2870},
	name: "_",
},
&litMatcher{
	pos: position{line: 99, col: 19, offset: 2872},
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
	pos: position{line: 103, col: 1, offset: 2956},
	expr: &actionExpr{
	pos: position{line: 103, col: 15, offset: 2970},
	run: (*parser).callonBranchJobG1,
	expr: &seqExpr{
	pos: position{line: 103, col: 15, offset: 2970},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 103, col: 15, offset: 2970},
	val: "branch",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 103, col: 23, offset: 2978},
	name: "_",
},
&litMatcher{
	pos: position{line: 103, col: 24, offset: 2979},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 103, col: 28, offset: 2983},
	name: "_",
},
&labeledExpr{
	pos: position{line: 103, col: 30, offset: 2985},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 103, col: 33, offset: 2988},
	name: "IDENTIFIER",
},
},
&litMatcher{
	pos: position{line: 103, col: 44, offset: 2999},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SecSource",
	pos: position{line: 110, col: 1, offset: 3215},
	expr: &choiceExpr{
	pos: position{line: 110, col: 14, offset: 3228},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 110, col: 14, offset: 3228},
	run: (*parser).callonSecSource2,
	expr: &seqExpr{
	pos: position{line: 110, col: 14, offset: 3228},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 110, col: 14, offset: 3228},
	label: "src1",
	expr: &ruleRefExpr{
	pos: position{line: 110, col: 19, offset: 3233},
	name: "IDENTIFIER",
},
},
&labeledExpr{
	pos: position{line: 110, col: 30, offset: 3244},
	label: "rest",
	expr: &zeroOrOneExpr{
	pos: position{line: 110, col: 35, offset: 3249},
	expr: &seqExpr{
	pos: position{line: 110, col: 36, offset: 3250},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 110, col: 36, offset: 3250},
	name: "_",
},
&litMatcher{
	pos: position{line: 110, col: 38, offset: 3252},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 110, col: 41, offset: 3255},
	name: "_",
},
&labeledExpr{
	pos: position{line: 110, col: 43, offset: 3257},
	label: "src2",
	expr: &ruleRefExpr{
	pos: position{line: 110, col: 48, offset: 3262},
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
	pos: position{line: 122, col: 3, offset: 3671},
	run: (*parser).callonSecSource14,
	expr: &seqExpr{
	pos: position{line: 122, col: 3, offset: 3671},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 122, col: 3, offset: 3671},
	label: "src1",
	expr: &ruleRefExpr{
	pos: position{line: 122, col: 8, offset: 3676},
	name: "IDENTIFIER",
},
},
&labeledExpr{
	pos: position{line: 122, col: 19, offset: 3687},
	label: "rest",
	expr: &ruleRefExpr{
	pos: position{line: 122, col: 25, offset: 3693},
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
	pos: position{line: 132, col: 1, offset: 3830},
	expr: &actionExpr{
	pos: position{line: 132, col: 13, offset: 3842},
	run: (*parser).callonJoinJobG1,
	expr: &seqExpr{
	pos: position{line: 132, col: 13, offset: 3842},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 132, col: 13, offset: 3842},
	val: "join",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 132, col: 20, offset: 3849},
	name: "_",
},
&labeledExpr{
	pos: position{line: 132, col: 22, offset: 3851},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 132, col: 26, offset: 3855},
	name: "JoinJobs",
},
},
	},
},
},
},
{
	name: "JoinJobs",
	pos: position{line: 136, col: 1, offset: 3928},
	expr: &choiceExpr{
	pos: position{line: 136, col: 12, offset: 3939},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 136, col: 12, offset: 3939},
	name: "InnerJoinG",
},
&ruleRefExpr{
	pos: position{line: 136, col: 23, offset: 3950},
	name: "OuterJoinG",
},
	},
},
},
{
	name: "InnerJoinG",
	pos: position{line: 138, col: 1, offset: 3964},
	expr: &choiceExpr{
	pos: position{line: 138, col: 14, offset: 3977},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 138, col: 14, offset: 3977},
	run: (*parser).callonInnerJoinG2,
	expr: &seqExpr{
	pos: position{line: 138, col: 14, offset: 3977},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 138, col: 14, offset: 3977},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 138, col: 20, offset: 3983},
	expr: &ruleRefExpr{
	pos: position{line: 138, col: 20, offset: 3983},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 138, col: 30, offset: 3993},
	name: "_",
},
&litMatcher{
	pos: position{line: 138, col: 32, offset: 3995},
	val: "inner",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 138, col: 40, offset: 4003},
	name: "_",
},
&litMatcher{
	pos: position{line: 138, col: 42, offset: 4005},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 138, col: 46, offset: 4009},
	name: "_",
},
&labeledExpr{
	pos: position{line: 138, col: 48, offset: 4011},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 138, col: 54, offset: 4017},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 138, col: 60, offset: 4023},
	name: "_",
},
&litMatcher{
	pos: position{line: 138, col: 62, offset: 4025},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 144, col: 4, offset: 4218},
	run: (*parser).callonInnerJoinG16,
	expr: &seqExpr{
	pos: position{line: 144, col: 4, offset: 4218},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 144, col: 4, offset: 4218},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 144, col: 10, offset: 4224},
	expr: &ruleRefExpr{
	pos: position{line: 144, col: 10, offset: 4224},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 144, col: 20, offset: 4234},
	name: "_",
},
&litMatcher{
	pos: position{line: 144, col: 22, offset: 4236},
	val: "inner",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 144, col: 30, offset: 4244},
	name: "_",
},
&litMatcher{
	pos: position{line: 144, col: 32, offset: 4246},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 144, col: 36, offset: 4250},
	name: "_",
},
&litMatcher{
	pos: position{line: 144, col: 38, offset: 4252},
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
	pos: position{line: 148, col: 1, offset: 4312},
	expr: &choiceExpr{
	pos: position{line: 148, col: 15, offset: 4326},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 148, col: 15, offset: 4326},
	name: "LeftOuterG",
},
&ruleRefExpr{
	pos: position{line: 148, col: 26, offset: 4337},
	name: "RightOuterG",
},
&ruleRefExpr{
	pos: position{line: 148, col: 38, offset: 4349},
	name: "FullOuterG",
},
	},
},
},
{
	name: "LeftOuterG",
	pos: position{line: 150, col: 1, offset: 4363},
	expr: &choiceExpr{
	pos: position{line: 150, col: 15, offset: 4377},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 150, col: 15, offset: 4377},
	run: (*parser).callonLeftOuterG2,
	expr: &seqExpr{
	pos: position{line: 150, col: 15, offset: 4377},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 150, col: 15, offset: 4377},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 150, col: 21, offset: 4383},
	expr: &ruleRefExpr{
	pos: position{line: 150, col: 21, offset: 4383},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 150, col: 31, offset: 4393},
	name: "_",
},
&litMatcher{
	pos: position{line: 150, col: 33, offset: 4395},
	val: "leftouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 150, col: 45, offset: 4407},
	name: "_",
},
&litMatcher{
	pos: position{line: 150, col: 47, offset: 4409},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 150, col: 51, offset: 4413},
	name: "_",
},
&labeledExpr{
	pos: position{line: 150, col: 53, offset: 4415},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 150, col: 59, offset: 4421},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 150, col: 65, offset: 4427},
	name: "_",
},
&litMatcher{
	pos: position{line: 150, col: 67, offset: 4429},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 156, col: 12, offset: 4700},
	run: (*parser).callonLeftOuterG16,
	expr: &seqExpr{
	pos: position{line: 156, col: 12, offset: 4700},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 156, col: 12, offset: 4700},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 156, col: 18, offset: 4706},
	expr: &ruleRefExpr{
	pos: position{line: 156, col: 18, offset: 4706},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 156, col: 28, offset: 4716},
	name: "_",
},
&litMatcher{
	pos: position{line: 156, col: 30, offset: 4718},
	val: "leftouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 156, col: 42, offset: 4730},
	name: "_",
},
&litMatcher{
	pos: position{line: 156, col: 44, offset: 4732},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 156, col: 48, offset: 4736},
	name: "_",
},
&litMatcher{
	pos: position{line: 156, col: 50, offset: 4738},
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
	pos: position{line: 160, col: 1, offset: 4817},
	expr: &choiceExpr{
	pos: position{line: 160, col: 15, offset: 4831},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 160, col: 15, offset: 4831},
	run: (*parser).callonRightOuterG2,
	expr: &seqExpr{
	pos: position{line: 160, col: 15, offset: 4831},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 160, col: 15, offset: 4831},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 160, col: 21, offset: 4837},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 160, col: 30, offset: 4846},
	name: "_",
},
&litMatcher{
	pos: position{line: 160, col: 32, offset: 4848},
	val: "rightouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 160, col: 45, offset: 4861},
	name: "_",
},
&litMatcher{
	pos: position{line: 160, col: 47, offset: 4863},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 160, col: 51, offset: 4867},
	name: "_",
},
&labeledExpr{
	pos: position{line: 160, col: 53, offset: 4869},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 160, col: 59, offset: 4875},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 160, col: 65, offset: 4881},
	name: "_",
},
&litMatcher{
	pos: position{line: 160, col: 67, offset: 4883},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 166, col: 17, offset: 5159},
	run: (*parser).callonRightOuterG15,
	expr: &seqExpr{
	pos: position{line: 166, col: 17, offset: 5159},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 166, col: 17, offset: 5159},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 166, col: 23, offset: 5165},
	expr: &ruleRefExpr{
	pos: position{line: 166, col: 23, offset: 5165},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 166, col: 33, offset: 5175},
	name: "_",
},
&litMatcher{
	pos: position{line: 166, col: 35, offset: 5177},
	val: "rightouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 166, col: 48, offset: 5190},
	name: "_",
},
&litMatcher{
	pos: position{line: 166, col: 50, offset: 5192},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 166, col: 54, offset: 5196},
	name: "_",
},
&litMatcher{
	pos: position{line: 166, col: 56, offset: 5198},
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
	pos: position{line: 169, col: 1, offset: 5305},
	expr: &choiceExpr{
	pos: position{line: 169, col: 15, offset: 5319},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 169, col: 15, offset: 5319},
	run: (*parser).callonFullOuterG2,
	expr: &seqExpr{
	pos: position{line: 169, col: 15, offset: 5319},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 169, col: 15, offset: 5319},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 169, col: 21, offset: 5325},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 169, col: 30, offset: 5334},
	name: "_",
},
&litMatcher{
	pos: position{line: 169, col: 32, offset: 5336},
	val: "rightouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 169, col: 45, offset: 5349},
	name: "_",
},
&litMatcher{
	pos: position{line: 169, col: 47, offset: 5351},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 169, col: 51, offset: 5355},
	name: "_",
},
&labeledExpr{
	pos: position{line: 169, col: 53, offset: 5357},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 169, col: 59, offset: 5363},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 169, col: 65, offset: 5369},
	name: "_",
},
&litMatcher{
	pos: position{line: 169, col: 67, offset: 5371},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 175, col: 17, offset: 5648},
	run: (*parser).callonFullOuterG15,
	expr: &seqExpr{
	pos: position{line: 175, col: 17, offset: 5648},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 175, col: 17, offset: 5648},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 175, col: 23, offset: 5654},
	expr: &ruleRefExpr{
	pos: position{line: 175, col: 23, offset: 5654},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 175, col: 33, offset: 5664},
	name: "_",
},
&litMatcher{
	pos: position{line: 175, col: 35, offset: 5666},
	val: "fullouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 175, col: 47, offset: 5678},
	name: "_",
},
&litMatcher{
	pos: position{line: 175, col: 49, offset: 5680},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 175, col: 53, offset: 5684},
	name: "_",
},
&litMatcher{
	pos: position{line: 175, col: 55, offset: 5686},
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
	pos: position{line: 180, col: 1, offset: 5780},
	expr: &choiceExpr{
	pos: position{line: 180, col: 10, offset: 5789},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 180, col: 10, offset: 5789},
	run: (*parser).callonQuery2,
	expr: &seqExpr{
	pos: position{line: 180, col: 10, offset: 5789},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 180, col: 10, offset: 5789},
	val: "select",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 180, col: 19, offset: 5798},
	name: "_",
},
&labeledExpr{
	pos: position{line: 180, col: 21, offset: 5800},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 180, col: 27, offset: 5806},
	name: "Fields",
},
},
&ruleRefExpr{
	pos: position{line: 180, col: 34, offset: 5813},
	name: "_",
},
&litMatcher{
	pos: position{line: 180, col: 36, offset: 5815},
	val: "on",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 180, col: 41, offset: 5820},
	name: "_",
},
&labeledExpr{
	pos: position{line: 180, col: 43, offset: 5822},
	label: "condition",
	expr: &ruleRefExpr{
	pos: position{line: 180, col: 53, offset: 5832},
	name: "JoinCondition",
},
},
&ruleRefExpr{
	pos: position{line: 180, col: 67, offset: 5846},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 182, col: 4, offset: 5956},
	run: (*parser).callonQuery14,
	expr: &seqExpr{
	pos: position{line: 182, col: 4, offset: 5956},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 182, col: 4, offset: 5956},
	val: "select",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 182, col: 13, offset: 5965},
	name: "_",
},
&litMatcher{
	pos: position{line: 182, col: 15, offset: 5967},
	val: "on",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 182, col: 20, offset: 5972},
	name: "_",
},
&labeledExpr{
	pos: position{line: 182, col: 22, offset: 5974},
	label: "condition",
	expr: &ruleRefExpr{
	pos: position{line: 182, col: 32, offset: 5984},
	name: "JoinCondition",
},
},
&ruleRefExpr{
	pos: position{line: 182, col: 46, offset: 5998},
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
	pos: position{line: 186, col: 1, offset: 6083},
	expr: &choiceExpr{
	pos: position{line: 186, col: 11, offset: 6093},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 186, col: 11, offset: 6093},
	name: "AllFields",
},
&ruleRefExpr{
	pos: position{line: 186, col: 21, offset: 6103},
	name: "SelFields",
},
	},
},
},
{
	name: "AllFields",
	pos: position{line: 188, col: 1, offset: 6116},
	expr: &actionExpr{
	pos: position{line: 188, col: 14, offset: 6129},
	run: (*parser).callonAllFields1,
	expr: &litMatcher{
	pos: position{line: 188, col: 14, offset: 6129},
	val: "*",
	ignoreCase: false,
},
},
},
{
	name: "SelFields",
	pos: position{line: 194, col: 1, offset: 6221},
	expr: &actionExpr{
	pos: position{line: 194, col: 14, offset: 6234},
	run: (*parser).callonSelFields1,
	expr: &seqExpr{
	pos: position{line: 194, col: 14, offset: 6234},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 194, col: 14, offset: 6234},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 194, col: 20, offset: 6240},
	name: "Word",
},
},
&ruleRefExpr{
	pos: position{line: 194, col: 25, offset: 6245},
	name: "_",
},
&labeledExpr{
	pos: position{line: 194, col: 28, offset: 6248},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 194, col: 33, offset: 6253},
	expr: &seqExpr{
	pos: position{line: 194, col: 34, offset: 6254},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 194, col: 34, offset: 6254},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 194, col: 38, offset: 6258},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 194, col: 40, offset: 6260},
	name: "Word",
},
&ruleRefExpr{
	pos: position{line: 194, col: 45, offset: 6265},
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
	pos: position{line: 204, col: 1, offset: 6518},
	expr: &choiceExpr{
	pos: position{line: 204, col: 18, offset: 6535},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 204, col: 18, offset: 6535},
	run: (*parser).callonJoinCondition2,
	expr: &seqExpr{
	pos: position{line: 204, col: 18, offset: 6535},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 204, col: 18, offset: 6535},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 204, col: 23, offset: 6540},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 204, col: 35, offset: 6552},
	name: "_",
},
&labeledExpr{
	pos: position{line: 204, col: 37, offset: 6554},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 204, col: 40, offset: 6557},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 204, col: 52, offset: 6569},
	name: "_",
},
&labeledExpr{
	pos: position{line: 204, col: 54, offset: 6571},
	label: "right",
	expr: &ruleRefExpr{
	pos: position{line: 204, col: 60, offset: 6577},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 204, col: 72, offset: 6589},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 206, col: 4, offset: 6707},
	run: (*parser).callonJoinCondition13,
	expr: &seqExpr{
	pos: position{line: 206, col: 4, offset: 6707},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 206, col: 4, offset: 6707},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 206, col: 9, offset: 6712},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 206, col: 21, offset: 6724},
	name: "_",
},
&labeledExpr{
	pos: position{line: 206, col: 23, offset: 6726},
	label: "right",
	expr: &ruleRefExpr{
	pos: position{line: 206, col: 29, offset: 6732},
	name: "JoinFacotors",
},
},
&ruleRefExpr{
	pos: position{line: 206, col: 42, offset: 6745},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 208, col: 3, offset: 6828},
	run: (*parser).callonJoinCondition21,
	expr: &seqExpr{
	pos: position{line: 208, col: 3, offset: 6828},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 208, col: 3, offset: 6828},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 208, col: 8, offset: 6833},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 208, col: 20, offset: 6845},
	name: "_",
},
&labeledExpr{
	pos: position{line: 208, col: 22, offset: 6847},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 208, col: 25, offset: 6850},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 208, col: 37, offset: 6862},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 210, col: 3, offset: 6944},
	run: (*parser).callonJoinCondition29,
	expr: &seqExpr{
	pos: position{line: 210, col: 3, offset: 6944},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 210, col: 3, offset: 6944},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 210, col: 8, offset: 6949},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 210, col: 20, offset: 6961},
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
	pos: position{line: 215, col: 1, offset: 7056},
	expr: &actionExpr{
	pos: position{line: 215, col: 16, offset: 7071},
	run: (*parser).callonJoinFactors1,
	expr: &seqExpr{
	pos: position{line: 215, col: 16, offset: 7071},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 215, col: 16, offset: 7071},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 215, col: 22, offset: 7077},
	name: "Word",
},
},
&ruleRefExpr{
	pos: position{line: 215, col: 27, offset: 7082},
	name: "_",
},
&labeledExpr{
	pos: position{line: 215, col: 29, offset: 7084},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 215, col: 34, offset: 7089},
	expr: &seqExpr{
	pos: position{line: 215, col: 35, offset: 7090},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 215, col: 35, offset: 7090},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 215, col: 39, offset: 7094},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 215, col: 41, offset: 7096},
	name: "Word",
},
&ruleRefExpr{
	pos: position{line: 215, col: 46, offset: 7101},
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
	name: "MapJobG",
	pos: position{line: 228, col: 1, offset: 7439},
	expr: &actionExpr{
	pos: position{line: 228, col: 12, offset: 7450},
	run: (*parser).callonMapJobG1,
	expr: &seqExpr{
	pos: position{line: 228, col: 12, offset: 7450},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 228, col: 12, offset: 7450},
	val: "map",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 228, col: 18, offset: 7456},
	name: "_",
},
&labeledExpr{
	pos: position{line: 228, col: 20, offset: 7458},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 228, col: 25, offset: 7463},
	name: "EnrichDo",
},
},
	},
},
},
},
{
	name: "EnrichDo",
	pos: position{line: 233, col: 1, offset: 7577},
	expr: &actionExpr{
	pos: position{line: 233, col: 13, offset: 7589},
	run: (*parser).callonEnrichDo1,
	expr: &seqExpr{
	pos: position{line: 233, col: 13, offset: 7589},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 233, col: 13, offset: 7589},
	label: "fld",
	expr: &ruleRefExpr{
	pos: position{line: 233, col: 17, offset: 7593},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 233, col: 26, offset: 7602},
	name: "_",
},
&litMatcher{
	pos: position{line: 233, col: 28, offset: 7604},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 233, col: 32, offset: 7608},
	name: "_",
},
&labeledExpr{
	pos: position{line: 233, col: 34, offset: 7610},
	label: "expr",
	expr: &ruleRefExpr{
	pos: position{line: 233, col: 39, offset: 7615},
	name: "MExpr",
},
},
	},
},
},
},
{
	name: "AggJobG",
	pos: position{line: 247, col: 1, offset: 7915},
	expr: &actionExpr{
	pos: position{line: 247, col: 12, offset: 7926},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 247, col: 12, offset: 7926},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 247, col: 12, offset: 7926},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 247, col: 18, offset: 7932},
	name: "_",
},
&labeledExpr{
	pos: position{line: 247, col: 20, offset: 7934},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 247, col: 27, offset: 7941},
	expr: &ruleRefExpr{
	pos: position{line: 247, col: 27, offset: 7941},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 247, col: 39, offset: 7953},
	name: "_",
},
&labeledExpr{
	pos: position{line: 247, col: 41, offset: 7955},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 247, col: 45, offset: 7959},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 247, col: 53, offset: 7967},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 247, col: 58, offset: 7972},
	expr: &seqExpr{
	pos: position{line: 247, col: 59, offset: 7973},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 247, col: 59, offset: 7973},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 247, col: 63, offset: 7977},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 247, col: 65, offset: 7979},
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
	pos: position{line: 262, col: 1, offset: 8354},
	expr: &choiceExpr{
	pos: position{line: 262, col: 12, offset: 8365},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 262, col: 12, offset: 8365},
	name: "CountAgg",
},
&ruleRefExpr{
	pos: position{line: 262, col: 23, offset: 8376},
	name: "MaxAgg",
},
&ruleRefExpr{
	pos: position{line: 262, col: 32, offset: 8385},
	name: "MinAgg",
},
&ruleRefExpr{
	pos: position{line: 262, col: 41, offset: 8394},
	name: "AvgAgg",
},
&ruleRefExpr{
	pos: position{line: 262, col: 50, offset: 8403},
	name: "SumAgg",
},
&ruleRefExpr{
	pos: position{line: 262, col: 59, offset: 8412},
	name: "ModeAgg",
},
&ruleRefExpr{
	pos: position{line: 263, col: 13, offset: 8433},
	name: "VarAgg",
},
&ruleRefExpr{
	pos: position{line: 263, col: 22, offset: 8442},
	name: "DistinctCountAgg",
},
&ruleRefExpr{
	pos: position{line: 263, col: 41, offset: 8461},
	name: "QuantileAgg",
},
&ruleRefExpr{
	pos: position{line: 263, col: 55, offset: 8475},
	name: "MedianAgg",
},
&ruleRefExpr{
	pos: position{line: 263, col: 67, offset: 8487},
	name: "QuartileAgg",
},
&ruleRefExpr{
	pos: position{line: 264, col: 13, offset: 8512},
	name: "CovAgg",
},
&ruleRefExpr{
	pos: position{line: 264, col: 22, offset: 8521},
	name: "CorrAgg",
},
&ruleRefExpr{
	pos: position{line: 264, col: 32, offset: 8531},
	name: "PValueAgg",
},
	},
},
},
{
	name: "CountAgg",
	pos: position{line: 267, col: 1, offset: 8575},
	expr: &actionExpr{
	pos: position{line: 267, col: 13, offset: 8587},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 267, col: 13, offset: 8587},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 267, col: 13, offset: 8587},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 267, col: 19, offset: 8593},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 267, col: 28, offset: 8602},
	name: "_",
},
&litMatcher{
	pos: position{line: 267, col: 30, offset: 8604},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 267, col: 39, offset: 8613},
	name: "_",
},
&labeledExpr{
	pos: position{line: 267, col: 41, offset: 8615},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 267, col: 46, offset: 8620},
	expr: &ruleRefExpr{
	pos: position{line: 267, col: 46, offset: 8620},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 267, col: 59, offset: 8633},
	name: "_",
},
&litMatcher{
	pos: position{line: 267, col: 61, offset: 8635},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MaxAgg",
	pos: position{line: 274, col: 1, offset: 8787},
	expr: &actionExpr{
	pos: position{line: 274, col: 11, offset: 8797},
	run: (*parser).callonMaxAgg1,
	expr: &seqExpr{
	pos: position{line: 274, col: 11, offset: 8797},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 274, col: 11, offset: 8797},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 274, col: 17, offset: 8803},
	expr: &ruleRefExpr{
	pos: position{line: 274, col: 17, offset: 8803},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 274, col: 27, offset: 8813},
	name: "_",
},
&litMatcher{
	pos: position{line: 274, col: 29, offset: 8815},
	val: "max(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 274, col: 36, offset: 8822},
	name: "_",
},
&labeledExpr{
	pos: position{line: 274, col: 38, offset: 8824},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 274, col: 44, offset: 8830},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 274, col: 53, offset: 8839},
	name: "_",
},
&labeledExpr{
	pos: position{line: 274, col: 55, offset: 8841},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 274, col: 60, offset: 8846},
	expr: &seqExpr{
	pos: position{line: 274, col: 61, offset: 8847},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 274, col: 61, offset: 8847},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 274, col: 65, offset: 8851},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 274, col: 67, offset: 8853},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 274, col: 81, offset: 8867},
	name: "_",
},
&litMatcher{
	pos: position{line: 274, col: 83, offset: 8869},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MinAgg",
	pos: position{line: 296, col: 1, offset: 9330},
	expr: &actionExpr{
	pos: position{line: 296, col: 11, offset: 9340},
	run: (*parser).callonMinAgg1,
	expr: &seqExpr{
	pos: position{line: 296, col: 11, offset: 9340},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 296, col: 11, offset: 9340},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 296, col: 17, offset: 9346},
	expr: &ruleRefExpr{
	pos: position{line: 296, col: 17, offset: 9346},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 296, col: 27, offset: 9356},
	name: "_",
},
&litMatcher{
	pos: position{line: 296, col: 29, offset: 9358},
	val: "min(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 296, col: 36, offset: 9365},
	name: "_",
},
&labeledExpr{
	pos: position{line: 296, col: 38, offset: 9367},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 296, col: 44, offset: 9373},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 296, col: 53, offset: 9382},
	name: "_",
},
&labeledExpr{
	pos: position{line: 296, col: 55, offset: 9384},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 296, col: 60, offset: 9389},
	expr: &seqExpr{
	pos: position{line: 296, col: 61, offset: 9390},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 296, col: 61, offset: 9390},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 296, col: 65, offset: 9394},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 296, col: 67, offset: 9396},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 296, col: 81, offset: 9410},
	name: "_",
},
&litMatcher{
	pos: position{line: 296, col: 83, offset: 9412},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AvgAgg",
	pos: position{line: 318, col: 1, offset: 9873},
	expr: &actionExpr{
	pos: position{line: 318, col: 11, offset: 9883},
	run: (*parser).callonAvgAgg1,
	expr: &seqExpr{
	pos: position{line: 318, col: 11, offset: 9883},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 318, col: 11, offset: 9883},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 318, col: 17, offset: 9889},
	expr: &ruleRefExpr{
	pos: position{line: 318, col: 17, offset: 9889},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 318, col: 27, offset: 9899},
	name: "_",
},
&litMatcher{
	pos: position{line: 318, col: 29, offset: 9901},
	val: "avg(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 318, col: 36, offset: 9908},
	name: "_",
},
&labeledExpr{
	pos: position{line: 318, col: 38, offset: 9910},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 318, col: 44, offset: 9916},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 318, col: 53, offset: 9925},
	name: "_",
},
&labeledExpr{
	pos: position{line: 318, col: 55, offset: 9927},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 318, col: 60, offset: 9932},
	expr: &seqExpr{
	pos: position{line: 318, col: 61, offset: 9933},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 318, col: 61, offset: 9933},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 318, col: 65, offset: 9937},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 318, col: 67, offset: 9939},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 318, col: 81, offset: 9953},
	name: "_",
},
&litMatcher{
	pos: position{line: 318, col: 83, offset: 9955},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SumAgg",
	pos: position{line: 340, col: 1, offset: 10416},
	expr: &actionExpr{
	pos: position{line: 340, col: 11, offset: 10426},
	run: (*parser).callonSumAgg1,
	expr: &seqExpr{
	pos: position{line: 340, col: 11, offset: 10426},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 340, col: 11, offset: 10426},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 340, col: 17, offset: 10432},
	expr: &ruleRefExpr{
	pos: position{line: 340, col: 17, offset: 10432},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 340, col: 27, offset: 10442},
	name: "_",
},
&litMatcher{
	pos: position{line: 340, col: 29, offset: 10444},
	val: "sum(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 340, col: 36, offset: 10451},
	name: "_",
},
&labeledExpr{
	pos: position{line: 340, col: 38, offset: 10453},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 340, col: 44, offset: 10459},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 340, col: 53, offset: 10468},
	name: "_",
},
&labeledExpr{
	pos: position{line: 340, col: 55, offset: 10470},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 340, col: 60, offset: 10475},
	expr: &seqExpr{
	pos: position{line: 340, col: 61, offset: 10476},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 340, col: 61, offset: 10476},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 340, col: 65, offset: 10480},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 340, col: 67, offset: 10482},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 340, col: 81, offset: 10496},
	name: "_",
},
&litMatcher{
	pos: position{line: 340, col: 83, offset: 10498},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarAgg",
	pos: position{line: 362, col: 1, offset: 10959},
	expr: &actionExpr{
	pos: position{line: 362, col: 11, offset: 10969},
	run: (*parser).callonVarAgg1,
	expr: &seqExpr{
	pos: position{line: 362, col: 11, offset: 10969},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 362, col: 11, offset: 10969},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 362, col: 17, offset: 10975},
	expr: &ruleRefExpr{
	pos: position{line: 362, col: 17, offset: 10975},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 362, col: 27, offset: 10985},
	name: "_",
},
&litMatcher{
	pos: position{line: 362, col: 29, offset: 10987},
	val: "var(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 362, col: 36, offset: 10994},
	name: "_",
},
&labeledExpr{
	pos: position{line: 362, col: 38, offset: 10996},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 362, col: 44, offset: 11002},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 362, col: 53, offset: 11011},
	name: "_",
},
&labeledExpr{
	pos: position{line: 362, col: 55, offset: 11013},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 362, col: 60, offset: 11018},
	expr: &seqExpr{
	pos: position{line: 362, col: 61, offset: 11019},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 362, col: 61, offset: 11019},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 362, col: 65, offset: 11023},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 362, col: 67, offset: 11025},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 362, col: 81, offset: 11039},
	name: "_",
},
&litMatcher{
	pos: position{line: 362, col: 83, offset: 11041},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DistinctCountAgg",
	pos: position{line: 384, col: 1, offset: 11507},
	expr: &actionExpr{
	pos: position{line: 384, col: 21, offset: 11527},
	run: (*parser).callonDistinctCountAgg1,
	expr: &seqExpr{
	pos: position{line: 384, col: 21, offset: 11527},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 384, col: 21, offset: 11527},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 384, col: 27, offset: 11533},
	expr: &ruleRefExpr{
	pos: position{line: 384, col: 27, offset: 11533},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 384, col: 37, offset: 11543},
	name: "_",
},
&litMatcher{
	pos: position{line: 384, col: 39, offset: 11545},
	val: "dcount(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 384, col: 49, offset: 11555},
	name: "_",
},
&labeledExpr{
	pos: position{line: 384, col: 51, offset: 11557},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 384, col: 57, offset: 11563},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 384, col: 66, offset: 11572},
	name: "_",
},
&labeledExpr{
	pos: position{line: 384, col: 68, offset: 11574},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 384, col: 73, offset: 11579},
	expr: &seqExpr{
	pos: position{line: 384, col: 74, offset: 11580},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 384, col: 74, offset: 11580},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 384, col: 78, offset: 11584},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 384, col: 80, offset: 11586},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 384, col: 94, offset: 11600},
	name: "_",
},
&litMatcher{
	pos: position{line: 384, col: 96, offset: 11602},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "ModeAgg",
	pos: position{line: 406, col: 1, offset: 12073},
	expr: &actionExpr{
	pos: position{line: 406, col: 12, offset: 12084},
	run: (*parser).callonModeAgg1,
	expr: &seqExpr{
	pos: position{line: 406, col: 12, offset: 12084},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 406, col: 12, offset: 12084},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 406, col: 18, offset: 12090},
	expr: &ruleRefExpr{
	pos: position{line: 406, col: 18, offset: 12090},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 406, col: 28, offset: 12100},
	name: "_",
},
&litMatcher{
	pos: position{line: 406, col: 30, offset: 12102},
	val: "mode(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 406, col: 38, offset: 12110},
	name: "_",
},
&labeledExpr{
	pos: position{line: 406, col: 40, offset: 12112},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 406, col: 46, offset: 12118},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 406, col: 55, offset: 12127},
	name: "_",
},
&labeledExpr{
	pos: position{line: 406, col: 57, offset: 12129},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 406, col: 62, offset: 12134},
	expr: &seqExpr{
	pos: position{line: 406, col: 63, offset: 12135},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 406, col: 63, offset: 12135},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 406, col: 67, offset: 12139},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 406, col: 69, offset: 12141},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 406, col: 83, offset: 12155},
	name: "_",
},
&litMatcher{
	pos: position{line: 406, col: 85, offset: 12157},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuantileAgg",
	pos: position{line: 428, col: 1, offset: 12619},
	expr: &actionExpr{
	pos: position{line: 428, col: 16, offset: 12634},
	run: (*parser).callonQuantileAgg1,
	expr: &seqExpr{
	pos: position{line: 428, col: 16, offset: 12634},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 428, col: 16, offset: 12634},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 428, col: 22, offset: 12640},
	expr: &ruleRefExpr{
	pos: position{line: 428, col: 22, offset: 12640},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 428, col: 32, offset: 12650},
	name: "_",
},
&litMatcher{
	pos: position{line: 428, col: 34, offset: 12652},
	val: "quantile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 428, col: 46, offset: 12664},
	name: "_",
},
&labeledExpr{
	pos: position{line: 428, col: 48, offset: 12666},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 428, col: 54, offset: 12672},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 428, col: 63, offset: 12681},
	name: "_",
},
&labeledExpr{
	pos: position{line: 428, col: 66, offset: 12684},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 428, col: 73, offset: 12691},
	expr: &seqExpr{
	pos: position{line: 428, col: 74, offset: 12692},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 428, col: 74, offset: 12692},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 428, col: 78, offset: 12696},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 428, col: 80, offset: 12698},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 428, col: 91, offset: 12709},
	name: "_",
},
&litMatcher{
	pos: position{line: 428, col: 93, offset: 12711},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 428, col: 97, offset: 12715},
	name: "_",
},
&labeledExpr{
	pos: position{line: 428, col: 99, offset: 12717},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 428, col: 103, offset: 12721},
	name: "Float",
},
},
&ruleRefExpr{
	pos: position{line: 428, col: 109, offset: 12727},
	name: "_",
},
&labeledExpr{
	pos: position{line: 428, col: 111, offset: 12729},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 428, col: 116, offset: 12734},
	expr: &seqExpr{
	pos: position{line: 428, col: 117, offset: 12735},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 428, col: 117, offset: 12735},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 428, col: 121, offset: 12739},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 428, col: 123, offset: 12741},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 428, col: 137, offset: 12755},
	name: "_",
},
&litMatcher{
	pos: position{line: 428, col: 139, offset: 12757},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MedianAgg",
	pos: position{line: 458, col: 1, offset: 13402},
	expr: &actionExpr{
	pos: position{line: 458, col: 14, offset: 13415},
	run: (*parser).callonMedianAgg1,
	expr: &seqExpr{
	pos: position{line: 458, col: 14, offset: 13415},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 458, col: 14, offset: 13415},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 458, col: 20, offset: 13421},
	expr: &ruleRefExpr{
	pos: position{line: 458, col: 20, offset: 13421},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 458, col: 30, offset: 13431},
	name: "_",
},
&litMatcher{
	pos: position{line: 458, col: 32, offset: 13433},
	val: "median(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 458, col: 42, offset: 13443},
	name: "_",
},
&labeledExpr{
	pos: position{line: 458, col: 44, offset: 13445},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 458, col: 50, offset: 13451},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 458, col: 59, offset: 13460},
	name: "_",
},
&labeledExpr{
	pos: position{line: 458, col: 62, offset: 13463},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 458, col: 69, offset: 13470},
	expr: &seqExpr{
	pos: position{line: 458, col: 70, offset: 13471},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 458, col: 70, offset: 13471},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 458, col: 74, offset: 13475},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 458, col: 76, offset: 13477},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 458, col: 87, offset: 13488},
	name: "_",
},
&labeledExpr{
	pos: position{line: 458, col: 89, offset: 13490},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 458, col: 94, offset: 13495},
	expr: &seqExpr{
	pos: position{line: 458, col: 95, offset: 13496},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 458, col: 95, offset: 13496},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 458, col: 99, offset: 13500},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 458, col: 101, offset: 13502},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 458, col: 115, offset: 13516},
	name: "_",
},
&litMatcher{
	pos: position{line: 458, col: 117, offset: 13518},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuartileAgg",
	pos: position{line: 490, col: 1, offset: 14195},
	expr: &actionExpr{
	pos: position{line: 490, col: 16, offset: 14210},
	run: (*parser).callonQuartileAgg1,
	expr: &seqExpr{
	pos: position{line: 490, col: 16, offset: 14210},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 490, col: 16, offset: 14210},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 490, col: 22, offset: 14216},
	expr: &ruleRefExpr{
	pos: position{line: 490, col: 22, offset: 14216},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 490, col: 32, offset: 14226},
	name: "_",
},
&litMatcher{
	pos: position{line: 490, col: 34, offset: 14228},
	val: "quartile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 490, col: 46, offset: 14240},
	name: "_",
},
&labeledExpr{
	pos: position{line: 490, col: 48, offset: 14242},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 490, col: 54, offset: 14248},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 490, col: 63, offset: 14257},
	name: "_",
},
&labeledExpr{
	pos: position{line: 490, col: 66, offset: 14260},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 490, col: 73, offset: 14267},
	expr: &seqExpr{
	pos: position{line: 490, col: 74, offset: 14268},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 490, col: 74, offset: 14268},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 490, col: 78, offset: 14272},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 490, col: 80, offset: 14274},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 490, col: 91, offset: 14285},
	name: "_",
},
&litMatcher{
	pos: position{line: 490, col: 93, offset: 14287},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 490, col: 97, offset: 14291},
	name: "_",
},
&labeledExpr{
	pos: position{line: 490, col: 99, offset: 14293},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 490, col: 103, offset: 14297},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 490, col: 110, offset: 14304},
	name: "_",
},
&labeledExpr{
	pos: position{line: 490, col: 112, offset: 14306},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 490, col: 117, offset: 14311},
	expr: &seqExpr{
	pos: position{line: 490, col: 118, offset: 14312},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 490, col: 118, offset: 14312},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 490, col: 122, offset: 14316},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 490, col: 124, offset: 14318},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 490, col: 138, offset: 14332},
	name: "_",
},
&litMatcher{
	pos: position{line: 490, col: 140, offset: 14334},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CovAgg",
	pos: position{line: 536, col: 1, offset: 15361},
	expr: &actionExpr{
	pos: position{line: 536, col: 11, offset: 15371},
	run: (*parser).callonCovAgg1,
	expr: &seqExpr{
	pos: position{line: 536, col: 11, offset: 15371},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 536, col: 11, offset: 15371},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 536, col: 17, offset: 15377},
	expr: &ruleRefExpr{
	pos: position{line: 536, col: 17, offset: 15377},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 536, col: 27, offset: 15387},
	name: "_",
},
&litMatcher{
	pos: position{line: 536, col: 29, offset: 15389},
	val: "cov(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 536, col: 36, offset: 15396},
	name: "_",
},
&labeledExpr{
	pos: position{line: 536, col: 38, offset: 15398},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 536, col: 45, offset: 15405},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 536, col: 54, offset: 15414},
	name: "_",
},
&litMatcher{
	pos: position{line: 536, col: 56, offset: 15416},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 536, col: 60, offset: 15420},
	name: "_",
},
&labeledExpr{
	pos: position{line: 536, col: 62, offset: 15422},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 536, col: 69, offset: 15429},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 536, col: 78, offset: 15438},
	name: "_",
},
&labeledExpr{
	pos: position{line: 536, col: 80, offset: 15440},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 536, col: 85, offset: 15445},
	expr: &seqExpr{
	pos: position{line: 536, col: 86, offset: 15446},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 536, col: 86, offset: 15446},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 536, col: 90, offset: 15450},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 536, col: 92, offset: 15452},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 536, col: 106, offset: 15466},
	name: "_",
},
&litMatcher{
	pos: position{line: 536, col: 108, offset: 15468},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CorrAgg",
	pos: position{line: 559, col: 1, offset: 15993},
	expr: &actionExpr{
	pos: position{line: 559, col: 12, offset: 16004},
	run: (*parser).callonCorrAgg1,
	expr: &seqExpr{
	pos: position{line: 559, col: 12, offset: 16004},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 559, col: 12, offset: 16004},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 559, col: 18, offset: 16010},
	expr: &ruleRefExpr{
	pos: position{line: 559, col: 18, offset: 16010},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 559, col: 28, offset: 16020},
	name: "_",
},
&litMatcher{
	pos: position{line: 559, col: 30, offset: 16022},
	val: "correlation(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 559, col: 45, offset: 16037},
	name: "_",
},
&labeledExpr{
	pos: position{line: 559, col: 47, offset: 16039},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 559, col: 54, offset: 16046},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 559, col: 63, offset: 16055},
	name: "_",
},
&litMatcher{
	pos: position{line: 559, col: 65, offset: 16057},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 559, col: 69, offset: 16061},
	name: "_",
},
&labeledExpr{
	pos: position{line: 559, col: 71, offset: 16063},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 559, col: 78, offset: 16070},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 559, col: 87, offset: 16079},
	name: "_",
},
&labeledExpr{
	pos: position{line: 559, col: 89, offset: 16081},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 559, col: 94, offset: 16086},
	expr: &seqExpr{
	pos: position{line: 559, col: 95, offset: 16087},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 559, col: 95, offset: 16087},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 559, col: 99, offset: 16091},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 559, col: 101, offset: 16093},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 559, col: 115, offset: 16107},
	name: "_",
},
&litMatcher{
	pos: position{line: 559, col: 117, offset: 16109},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PValueAgg",
	pos: position{line: 582, col: 1, offset: 16635},
	expr: &actionExpr{
	pos: position{line: 582, col: 14, offset: 16648},
	run: (*parser).callonPValueAgg1,
	expr: &seqExpr{
	pos: position{line: 582, col: 14, offset: 16648},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 582, col: 14, offset: 16648},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 582, col: 20, offset: 16654},
	expr: &ruleRefExpr{
	pos: position{line: 582, col: 20, offset: 16654},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 582, col: 30, offset: 16664},
	name: "_",
},
&litMatcher{
	pos: position{line: 582, col: 32, offset: 16666},
	val: "pvalue(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 582, col: 42, offset: 16676},
	name: "_",
},
&labeledExpr{
	pos: position{line: 582, col: 44, offset: 16678},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 582, col: 51, offset: 16685},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 582, col: 60, offset: 16694},
	name: "_",
},
&litMatcher{
	pos: position{line: 582, col: 62, offset: 16696},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 582, col: 66, offset: 16700},
	name: "_",
},
&labeledExpr{
	pos: position{line: 582, col: 68, offset: 16702},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 582, col: 75, offset: 16709},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 582, col: 84, offset: 16718},
	name: "_",
},
&labeledExpr{
	pos: position{line: 582, col: 86, offset: 16720},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 582, col: 91, offset: 16725},
	expr: &seqExpr{
	pos: position{line: 582, col: 92, offset: 16726},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 582, col: 92, offset: 16726},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 582, col: 96, offset: 16730},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 582, col: 98, offset: 16732},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 582, col: 112, offset: 16746},
	name: "_",
},
&litMatcher{
	pos: position{line: 582, col: 114, offset: 16748},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggGroupBy",
	pos: position{line: 607, col: 1, offset: 17273},
	expr: &actionExpr{
	pos: position{line: 607, col: 15, offset: 17287},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 607, col: 15, offset: 17287},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 607, col: 15, offset: 17287},
	name: "_",
},
&litMatcher{
	pos: position{line: 607, col: 17, offset: 17289},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 607, col: 22, offset: 17294},
	name: "_",
},
&labeledExpr{
	pos: position{line: 607, col: 24, offset: 17296},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 607, col: 30, offset: 17302},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 607, col: 39, offset: 17311},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 607, col: 44, offset: 17316},
	expr: &seqExpr{
	pos: position{line: 607, col: 45, offset: 17317},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 607, col: 45, offset: 17317},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 607, col: 49, offset: 17321},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 607, col: 51, offset: 17323},
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
	pos: position{line: 622, col: 1, offset: 17674},
	expr: &actionExpr{
	pos: position{line: 622, col: 11, offset: 17684},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 622, col: 11, offset: 17684},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 622, col: 16, offset: 17689},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 622, col: 16, offset: 17689},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 622, col: 27, offset: 17700},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 622, col: 38, offset: 17711},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 622, col: 49, offset: 17722},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 622, col: 58, offset: 17731},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 622, col: 67, offset: 17740},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 628, col: 1, offset: 17859},
	expr: &actionExpr{
	pos: position{line: 628, col: 13, offset: 17871},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 628, col: 13, offset: 17871},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 628, col: 13, offset: 17871},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 628, col: 23, offset: 17881},
	name: "_",
},
&labeledExpr{
	pos: position{line: 628, col: 25, offset: 17883},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 628, col: 28, offset: 17886},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 628, col: 42, offset: 17900},
	name: "_",
},
&litMatcher{
	pos: position{line: 628, col: 44, offset: 17902},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 635, col: 1, offset: 18097},
	expr: &actionExpr{
	pos: position{line: 635, col: 18, offset: 18114},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 635, col: 18, offset: 18114},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 635, col: 18, offset: 18114},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 635, col: 23, offset: 18119},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 635, col: 28, offset: 18124},
	name: "_",
},
&labeledExpr{
	pos: position{line: 635, col: 30, offset: 18126},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 635, col: 33, offset: 18129},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 635, col: 45, offset: 18141},
	name: "_",
},
&labeledExpr{
	pos: position{line: 635, col: 47, offset: 18143},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 635, col: 54, offset: 18150},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 639, col: 1, offset: 18201},
	expr: &ruleRefExpr{
	pos: position{line: 639, col: 16, offset: 18216},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 642, col: 1, offset: 18250},
	expr: &actionExpr{
	pos: position{line: 642, col: 13, offset: 18262},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 642, col: 13, offset: 18262},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 642, col: 13, offset: 18262},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 642, col: 23, offset: 18272},
	name: "_",
},
&labeledExpr{
	pos: position{line: 642, col: 25, offset: 18274},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 642, col: 31, offset: 18280},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 642, col: 40, offset: 18289},
	name: "_",
},
&labeledExpr{
	pos: position{line: 642, col: 42, offset: 18291},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 642, col: 47, offset: 18296},
	expr: &seqExpr{
	pos: position{line: 642, col: 48, offset: 18297},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 642, col: 48, offset: 18297},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 642, col: 52, offset: 18301},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 642, col: 54, offset: 18303},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 642, col: 65, offset: 18314},
	name: "_",
},
&litMatcher{
	pos: position{line: 642, col: 67, offset: 18316},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 657, col: 1, offset: 18653},
	expr: &actionExpr{
	pos: position{line: 657, col: 11, offset: 18663},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 657, col: 11, offset: 18663},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 657, col: 11, offset: 18663},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 657, col: 18, offset: 18670},
	name: "_",
},
&labeledExpr{
	pos: position{line: 657, col: 20, offset: 18672},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 657, col: 25, offset: 18677},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 657, col: 34, offset: 18686},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 657, col: 38, offset: 18690},
	name: "_",
},
&labeledExpr{
	pos: position{line: 657, col: 40, offset: 18692},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 657, col: 44, offset: 18696},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 657, col: 51, offset: 18703},
	name: "_",
},
&litMatcher{
	pos: position{line: 657, col: 53, offset: 18705},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 669, col: 1, offset: 18927},
	expr: &actionExpr{
	pos: position{line: 669, col: 11, offset: 18937},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 669, col: 11, offset: 18937},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 669, col: 11, offset: 18937},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 669, col: 20, offset: 18946},
	name: "_",
},
&litMatcher{
	pos: position{line: 669, col: 22, offset: 18948},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 669, col: 27, offset: 18953},
	name: "_",
},
&labeledExpr{
	pos: position{line: 669, col: 29, offset: 18955},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 669, col: 35, offset: 18961},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 675, col: 1, offset: 19067},
	expr: &actionExpr{
	pos: position{line: 675, col: 12, offset: 19078},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 675, col: 12, offset: 19078},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 675, col: 12, offset: 19078},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 675, col: 21, offset: 19087},
	name: "_",
},
&labeledExpr{
	pos: position{line: 675, col: 23, offset: 19089},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 675, col: 27, offset: 19093},
	expr: &ruleRefExpr{
	pos: position{line: 675, col: 27, offset: 19093},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 675, col: 35, offset: 19101},
	name: "_",
},
&litMatcher{
	pos: position{line: 675, col: 37, offset: 19103},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 685, col: 1, offset: 19248},
	expr: &actionExpr{
	pos: position{line: 685, col: 13, offset: 19260},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 685, col: 13, offset: 19260},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 685, col: 13, offset: 19260},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 685, col: 23, offset: 19270},
	name: "_",
},
&labeledExpr{
	pos: position{line: 685, col: 25, offset: 19272},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 685, col: 30, offset: 19277},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 685, col: 42, offset: 19289},
	name: "_",
},
&litMatcher{
	pos: position{line: 685, col: 44, offset: 19291},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 695, col: 1, offset: 19579},
	expr: &actionExpr{
	pos: position{line: 695, col: 16, offset: 19594},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 695, col: 16, offset: 19594},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 695, col: 16, offset: 19594},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 695, col: 22, offset: 19600},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 695, col: 35, offset: 19613},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 695, col: 40, offset: 19618},
	expr: &seqExpr{
	pos: position{line: 695, col: 41, offset: 19619},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 695, col: 41, offset: 19619},
	name: "_",
},
&labeledExpr{
	pos: position{line: 695, col: 43, offset: 19621},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 695, col: 46, offset: 19624},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 695, col: 58, offset: 19636},
	name: "_",
},
&labeledExpr{
	pos: position{line: 695, col: 60, offset: 19638},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 695, col: 63, offset: 19641},
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
	pos: position{line: 700, col: 1, offset: 19799},
	expr: &choiceExpr{
	pos: position{line: 700, col: 17, offset: 19815},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 700, col: 17, offset: 19815},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 700, col: 17, offset: 19815},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 700, col: 17, offset: 19815},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 700, col: 21, offset: 19819},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 700, col: 24, offset: 19822},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 700, col: 36, offset: 19834},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 702, col: 5, offset: 19865},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 702, col: 5, offset: 19865},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 702, col: 8, offset: 19868},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 706, col: 1, offset: 19910},
	expr: &choiceExpr{
	pos: position{line: 706, col: 16, offset: 19925},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 706, col: 16, offset: 19925},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 706, col: 24, offset: 19933},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 706, col: 24, offset: 19933},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 713, col: 1, offset: 20112},
	expr: &actionExpr{
	pos: position{line: 713, col: 18, offset: 20129},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 713, col: 18, offset: 20129},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 713, col: 18, offset: 20129},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 713, col: 23, offset: 20134},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 713, col: 28, offset: 20139},
	name: "_",
},
&labeledExpr{
	pos: position{line: 713, col: 30, offset: 20141},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 713, col: 33, offset: 20144},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 713, col: 45, offset: 20156},
	name: "_",
},
&labeledExpr{
	pos: position{line: 713, col: 47, offset: 20158},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 713, col: 53, offset: 20164},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 713, col: 59, offset: 20170},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 717, col: 1, offset: 20244},
	expr: &choiceExpr{
	pos: position{line: 717, col: 16, offset: 20259},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 717, col: 16, offset: 20259},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 717, col: 23, offset: 20266},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 717, col: 30, offset: 20273},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 717, col: 36, offset: 20279},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 717, col: 42, offset: 20285},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 717, col: 49, offset: 20292},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 717, col: 49, offset: 20292},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 726, col: 1, offset: 20402},
	expr: &choiceExpr{
	pos: position{line: 726, col: 12, offset: 20413},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 726, col: 12, offset: 20413},
	run: (*parser).callonStdOutG2,
	expr: &seqExpr{
	pos: position{line: 726, col: 12, offset: 20413},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 726, col: 12, offset: 20413},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 726, col: 22, offset: 20423},
	name: "_",
},
&labeledExpr{
	pos: position{line: 726, col: 24, offset: 20425},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 726, col: 31, offset: 20432},
	expr: &ruleRefExpr{
	pos: position{line: 726, col: 32, offset: 20433},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 726, col: 48, offset: 20449},
	name: "_",
},
&litMatcher{
	pos: position{line: 726, col: 50, offset: 20451},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 738, col: 3, offset: 20727},
	run: (*parser).callonStdOutG11,
	expr: &litMatcher{
	pos: position{line: 738, col: 3, offset: 20727},
	val: "stdout",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "SinkInto",
	pos: position{line: 742, col: 1, offset: 20797},
	expr: &choiceExpr{
	pos: position{line: 742, col: 12, offset: 20808},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 742, col: 12, offset: 20808},
	run: (*parser).callonSinkInto2,
	expr: &seqExpr{
	pos: position{line: 742, col: 12, offset: 20808},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 742, col: 12, offset: 20808},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 742, col: 18, offset: 20814},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 742, col: 27, offset: 20823},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 742, col: 33, offset: 20829},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 742, col: 33, offset: 20829},
	name: "_",
},
&litMatcher{
	pos: position{line: 742, col: 35, offset: 20831},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 742, col: 39, offset: 20835},
	name: "_",
},
&litMatcher{
	pos: position{line: 742, col: 41, offset: 20837},
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
	pos: position{line: 748, col: 7, offset: 20990},
	run: (*parser).callonSinkInto12,
	expr: &seqExpr{
	pos: position{line: 748, col: 7, offset: 20990},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 748, col: 7, offset: 20990},
	val: "into",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 748, col: 14, offset: 20997},
	name: "_",
},
&litMatcher{
	pos: position{line: 748, col: 16, offset: 20999},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 748, col: 20, offset: 21003},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 748, col: 22, offset: 21005},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 748, col: 31, offset: 21014},
	name: "_",
},
&litMatcher{
	pos: position{line: 748, col: 33, offset: 21016},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 750, col: 7, offset: 21112},
	run: (*parser).callonSinkInto21,
	expr: &seqExpr{
	pos: position{line: 750, col: 7, offset: 21112},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 750, col: 7, offset: 21112},
	val: "into",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 750, col: 14, offset: 21119},
	name: "_",
},
&litMatcher{
	pos: position{line: 750, col: 16, offset: 21121},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 750, col: 20, offset: 21125},
	name: "_",
},
&litMatcher{
	pos: position{line: 750, col: 22, offset: 21127},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 752, col: 7, offset: 21266},
	run: (*parser).callonSinkInto28,
	expr: &seqExpr{
	pos: position{line: 752, col: 7, offset: 21266},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 752, col: 7, offset: 21266},
	val: "into",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 752, col: 14, offset: 21273},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 754, col: 7, offset: 21366},
	run: (*parser).callonSinkInto32,
	expr: &seqExpr{
	pos: position{line: 754, col: 7, offset: 21366},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 754, col: 7, offset: 21366},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 754, col: 13, offset: 21372},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 754, col: 22, offset: 21381},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 754, col: 29, offset: 21388},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 754, col: 29, offset: 21388},
	name: "_",
},
&litMatcher{
	pos: position{line: 754, col: 31, offset: 21390},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 754, col: 35, offset: 21394},
	name: "_",
},
&litMatcher{
	pos: position{line: 754, col: 37, offset: 21396},
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
	pos: position{line: 758, col: 1, offset: 21467},
	expr: &choiceExpr{
	pos: position{line: 758, col: 15, offset: 21481},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 758, col: 15, offset: 21481},
	run: (*parser).callonBlackHoleG2,
	expr: &seqExpr{
	pos: position{line: 758, col: 15, offset: 21481},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 758, col: 15, offset: 21481},
	val: "blackhole(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 758, col: 28, offset: 21494},
	name: "_",
},
&labeledExpr{
	pos: position{line: 758, col: 30, offset: 21496},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 758, col: 37, offset: 21503},
	expr: &ruleRefExpr{
	pos: position{line: 758, col: 38, offset: 21504},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 758, col: 54, offset: 21520},
	name: "_",
},
&litMatcher{
	pos: position{line: 758, col: 56, offset: 21522},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 760, col: 3, offset: 21577},
	run: (*parser).callonBlackHoleG11,
	expr: &seqExpr{
	pos: position{line: 760, col: 3, offset: 21577},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 760, col: 3, offset: 21577},
	val: "blackhole",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 760, col: 15, offset: 21589},
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
	pos: position{line: 764, col: 1, offset: 21674},
	expr: &choiceExpr{
	pos: position{line: 764, col: 10, offset: 21683},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 764, col: 10, offset: 21683},
	run: (*parser).callonPlotG2,
	expr: &seqExpr{
	pos: position{line: 764, col: 10, offset: 21683},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 764, col: 10, offset: 21683},
	val: "plot",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 764, col: 17, offset: 21690},
	name: "_",
},
&labeledExpr{
	pos: position{line: 764, col: 19, offset: 21692},
	label: "widget",
	expr: &ruleRefExpr{
	pos: position{line: 764, col: 27, offset: 21700},
	name: "BarWidget",
},
},
	},
},
},
&actionExpr{
	pos: position{line: 766, col: 3, offset: 21797},
	run: (*parser).callonPlotG8,
	expr: &seqExpr{
	pos: position{line: 766, col: 3, offset: 21797},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 766, col: 3, offset: 21797},
	val: "plot",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 766, col: 10, offset: 21804},
	name: "_",
},
	},
},
},
	},
},
},
{
	name: "BarWidget",
	pos: position{line: 770, col: 1, offset: 21864},
	expr: &actionExpr{
	pos: position{line: 770, col: 14, offset: 21877},
	run: (*parser).callonBarWidget1,
	expr: &seqExpr{
	pos: position{line: 770, col: 14, offset: 21877},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 770, col: 14, offset: 21877},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 770, col: 20, offset: 21883},
	expr: &ruleRefExpr{
	pos: position{line: 770, col: 20, offset: 21883},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 770, col: 30, offset: 21893},
	name: "_",
},
&litMatcher{
	pos: position{line: 770, col: 32, offset: 21895},
	val: "bar(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 770, col: 39, offset: 21902},
	name: "_",
},
&labeledExpr{
	pos: position{line: 770, col: 41, offset: 21904},
	label: "xField",
	expr: &ruleRefExpr{
	pos: position{line: 770, col: 48, offset: 21911},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 770, col: 57, offset: 21920},
	name: "_",
},
&litMatcher{
	pos: position{line: 770, col: 59, offset: 21922},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 770, col: 63, offset: 21926},
	name: "_",
},
&labeledExpr{
	pos: position{line: 770, col: 65, offset: 21928},
	label: "yField",
	expr: &ruleRefExpr{
	pos: position{line: 770, col: 72, offset: 21935},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 770, col: 81, offset: 21944},
	name: "_",
},
&labeledExpr{
	pos: position{line: 770, col: 83, offset: 21946},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 770, col: 88, offset: 21951},
	expr: &seqExpr{
	pos: position{line: 770, col: 89, offset: 21952},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 770, col: 89, offset: 21952},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 770, col: 93, offset: 21956},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 770, col: 95, offset: 21958},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 770, col: 102, offset: 21965},
	name: "_",
},
&zeroOrOneExpr{
	pos: position{line: 770, col: 104, offset: 21967},
	expr: &seqExpr{
	pos: position{line: 770, col: 105, offset: 21968},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 770, col: 105, offset: 21968},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 770, col: 109, offset: 21972},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 770, col: 111, offset: 21974},
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
	pos: position{line: 770, col: 122, offset: 21985},
	name: "_",
},
&litMatcher{
	pos: position{line: 770, col: 124, offset: 21987},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FileName",
	pos: position{line: 803, col: 1, offset: 22672},
	expr: &actionExpr{
	pos: position{line: 803, col: 13, offset: 22684},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 803, col: 13, offset: 22684},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 803, col: 13, offset: 22684},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 803, col: 22, offset: 22693},
	expr: &seqExpr{
	pos: position{line: 803, col: 23, offset: 22694},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 803, col: 23, offset: 22694},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 803, col: 27, offset: 22698},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 803, col: 31, offset: 22702},
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
	pos: position{line: 808, col: 1, offset: 22801},
	expr: &actionExpr{
	pos: position{line: 808, col: 10, offset: 22810},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 808, col: 10, offset: 22810},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 808, col: 10, offset: 22810},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 808, col: 16, offset: 22816},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 808, col: 23, offset: 22823},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 808, col: 28, offset: 22828},
	expr: &seqExpr{
	pos: position{line: 808, col: 29, offset: 22829},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 808, col: 29, offset: 22829},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 808, col: 31, offset: 22831},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 808, col: 36, offset: 22836},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 808, col: 38, offset: 22838},
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
	pos: position{line: 812, col: 1, offset: 22903},
	expr: &actionExpr{
	pos: position{line: 812, col: 11, offset: 22913},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 812, col: 11, offset: 22913},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 812, col: 16, offset: 22918},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 812, col: 16, offset: 22918},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 812, col: 24, offset: 22926},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 812, col: 33, offset: 22935},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 812, col: 48, offset: 22950},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 812, col: 59, offset: 22961},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 816, col: 1, offset: 23003},
	expr: &choiceExpr{
	pos: position{line: 816, col: 9, offset: 23011},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 816, col: 9, offset: 23011},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 816, col: 15, offset: 23017},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 816, col: 21, offset: 23023},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 816, col: 27, offset: 23029},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 816, col: 33, offset: 23035},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 816, col: 33, offset: 23035},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 820, col: 1, offset: 23083},
	expr: &choiceExpr{
	pos: position{line: 820, col: 12, offset: 23094},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 820, col: 12, offset: 23094},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 820, col: 12, offset: 23094},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 820, col: 12, offset: 23094},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 820, col: 16, offset: 23098},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 820, col: 19, offset: 23101},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 820, col: 25, offset: 23107},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 824, col: 5, offset: 23183},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 824, col: 5, offset: 23183},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 824, col: 8, offset: 23186},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 831, col: 1, offset: 23263},
	expr: &actionExpr{
	pos: position{line: 831, col: 9, offset: 23271},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 831, col: 9, offset: 23271},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 831, col: 9, offset: 23271},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 831, col: 15, offset: 23277},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 831, col: 21, offset: 23283},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 831, col: 26, offset: 23288},
	expr: &seqExpr{
	pos: position{line: 831, col: 27, offset: 23289},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 831, col: 27, offset: 23289},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 831, col: 29, offset: 23291},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 831, col: 34, offset: 23296},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 831, col: 36, offset: 23298},
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
	pos: position{line: 835, col: 1, offset: 23362},
	expr: &actionExpr{
	pos: position{line: 835, col: 10, offset: 23371},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 835, col: 10, offset: 23371},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 835, col: 15, offset: 23376},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 835, col: 15, offset: 23376},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 835, col: 23, offset: 23384},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 835, col: 32, offset: 23393},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 835, col: 47, offset: 23408},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 839, col: 1, offset: 23451},
	expr: &actionExpr{
	pos: position{line: 839, col: 13, offset: 23463},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 839, col: 13, offset: 23463},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 839, col: 13, offset: 23463},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 839, col: 20, offset: 23470},
	expr: &choiceExpr{
	pos: position{line: 839, col: 21, offset: 23471},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 839, col: 21, offset: 23471},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 839, col: 31, offset: 23481},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 839, col: 39, offset: 23489},
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
	pos: position{line: 843, col: 1, offset: 23539},
	expr: &actionExpr{
	pos: position{line: 843, col: 17, offset: 23555},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 843, col: 17, offset: 23555},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 843, col: 17, offset: 23555},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 843, col: 21, offset: 23559},
	expr: &choiceExpr{
	pos: position{line: 843, col: 23, offset: 23561},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 843, col: 23, offset: 23561},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 843, col: 23, offset: 23561},
	expr: &ruleRefExpr{
	pos: position{line: 843, col: 24, offset: 23562},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 843, col: 36, offset: 23574,
},
	},
},
&seqExpr{
	pos: position{line: 843, col: 40, offset: 23578},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 843, col: 40, offset: 23578},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 843, col: 45, offset: 23583},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 843, col: 63, offset: 23601},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 850, col: 1, offset: 23782},
	expr: &charClassMatcher{
	pos: position{line: 850, col: 16, offset: 23797},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 852, col: 1, offset: 23815},
	expr: &ruleRefExpr{
	pos: position{line: 852, col: 19, offset: 23833},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 854, col: 1, offset: 23853},
	expr: &charClassMatcher{
	pos: position{line: 854, col: 21, offset: 23873},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 856, col: 1, offset: 23888},
	expr: &actionExpr{
	pos: position{line: 856, col: 9, offset: 23896},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 856, col: 9, offset: 23896},
	expr: &choiceExpr{
	pos: position{line: 856, col: 10, offset: 23897},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 856, col: 10, offset: 23897},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 856, col: 19, offset: 23906},
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
	pos: position{line: 860, col: 1, offset: 23956},
	expr: &choiceExpr{
	pos: position{line: 860, col: 11, offset: 23966},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 860, col: 11, offset: 23966},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 860, col: 19, offset: 23974},
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
	pos: position{line: 862, col: 1, offset: 23983},
	expr: &actionExpr{
	pos: position{line: 862, col: 11, offset: 23993},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 862, col: 11, offset: 23993},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 862, col: 11, offset: 23993},
	expr: &litMatcher{
	pos: position{line: 862, col: 11, offset: 23993},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 862, col: 16, offset: 23998},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 866, col: 1, offset: 24071},
	expr: &actionExpr{
	pos: position{line: 866, col: 19, offset: 24089},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 866, col: 19, offset: 24089},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 870, col: 1, offset: 24162},
	expr: &actionExpr{
	pos: position{line: 870, col: 10, offset: 24171},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 870, col: 10, offset: 24171},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 870, col: 10, offset: 24171},
	expr: &litMatcher{
	pos: position{line: 870, col: 10, offset: 24171},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 870, col: 15, offset: 24176},
	expr: &ruleRefExpr{
	pos: position{line: 870, col: 15, offset: 24176},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 870, col: 24, offset: 24185},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 870, col: 28, offset: 24189},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 874, col: 1, offset: 24256},
	expr: &choiceExpr{
	pos: position{line: 874, col: 12, offset: 24267},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 874, col: 12, offset: 24267},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 874, col: 18, offset: 24273},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 874, col: 18, offset: 24273},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 874, col: 38, offset: 24293},
	expr: &ruleRefExpr{
	pos: position{line: 874, col: 38, offset: 24293},
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
	pos: position{line: 876, col: 1, offset: 24310},
	expr: &charClassMatcher{
	pos: position{line: 876, col: 17, offset: 24326},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 878, col: 1, offset: 24335},
	expr: &charClassMatcher{
	pos: position{line: 878, col: 24, offset: 24358},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 880, col: 1, offset: 24367},
	expr: &zeroOrMoreExpr{
	pos: position{line: 880, col: 19, offset: 24385},
	expr: &charClassMatcher{
	pos: position{line: 880, col: 19, offset: 24385},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 882, col: 1, offset: 24399},
	expr: &charClassMatcher{
	pos: position{line: 882, col: 18, offset: 24416},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "IDENTIFIER",
	pos: position{line: 883, col: 1, offset: 24427},
	expr: &actionExpr{
	pos: position{line: 883, col: 15, offset: 24441},
	run: (*parser).callonIDENTIFIER1,
	expr: &seqExpr{
	pos: position{line: 883, col: 15, offset: 24441},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 883, col: 15, offset: 24441},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 883, col: 22, offset: 24448},
	expr: &choiceExpr{
	pos: position{line: 883, col: 23, offset: 24449},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 883, col: 23, offset: 24449},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 883, col: 33, offset: 24459},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 883, col: 41, offset: 24467},
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
	pos: position{line: 886, col: 1, offset: 24540},
	expr: &notExpr{
	pos: position{line: 886, col: 8, offset: 24547},
	expr: &anyMatcher{
	line: 886, col: 9, offset: 24548,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 888, col: 1, offset: 24553},
	expr: &litMatcher{
	pos: position{line: 888, col: 17, offset: 24569},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 894, col: 1, offset: 24656},
	expr: &actionExpr{
	pos: position{line: 894, col: 18, offset: 24673},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 894, col: 18, offset: 24673},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 894, col: 18, offset: 24673},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 894, col: 24, offset: 24679},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 894, col: 33, offset: 24688},
	name: "_",
},
&labeledExpr{
	pos: position{line: 894, col: 35, offset: 24690},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 894, col: 40, offset: 24695},
	expr: &seqExpr{
	pos: position{line: 894, col: 41, offset: 24696},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 894, col: 41, offset: 24696},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 894, col: 45, offset: 24700},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 894, col: 47, offset: 24702},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 894, col: 56, offset: 24711},
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
	pos: position{line: 897, col: 1, offset: 24764},
	expr: &actionExpr{
	pos: position{line: 897, col: 13, offset: 24776},
	run: (*parser).callonEqAliasG1,
	expr: &seqExpr{
	pos: position{line: 897, col: 13, offset: 24776},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 897, col: 13, offset: 24776},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 897, col: 19, offset: 24782},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 897, col: 28, offset: 24791},
	name: "_",
},
&litMatcher{
	pos: position{line: 897, col: 30, offset: 24793},
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

func (c *current) onMapJobG1(job interface{}) (interface{}, error) {

    return TransformJob{Type:DOJOB, OperateOn:DoNodeJob{Function: job}}, nil
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

func (c *current) onPlotG2(widget interface{}) (interface{}, error) {

    return SinkJob{Type: PLOT, OperateOn:Plot{Type:"plot",Widget: widget}}, nil
}

func (p *parser) callonPlotG2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPlotG2(stack["widget"])
}

func (c *current) onPlotG8() (interface{}, error) {

 return nil,errors.New("widget missing in plot")
}

func (p *parser) callonPlotG8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPlotG8()
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

