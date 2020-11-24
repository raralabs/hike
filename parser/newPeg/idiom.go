
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
	expr: &choiceExpr{
	pos: position{line: 123, col: 12, offset: 3228},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 123, col: 12, offset: 3228},
	name: "InnerJoinG",
},
&ruleRefExpr{
	pos: position{line: 123, col: 23, offset: 3239},
	name: "OuterJoinG",
},
	},
},
},
{
	name: "InnerJoinG",
	pos: position{line: 125, col: 1, offset: 3253},
	expr: &actionExpr{
	pos: position{line: 125, col: 14, offset: 3266},
	run: (*parser).callonInnerJoinG1,
	expr: &seqExpr{
	pos: position{line: 125, col: 14, offset: 3266},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 125, col: 14, offset: 3266},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 125, col: 20, offset: 3272},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 125, col: 29, offset: 3281},
	name: "_",
},
&litMatcher{
	pos: position{line: 125, col: 31, offset: 3283},
	val: "inner",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 125, col: 39, offset: 3291},
	name: "_",
},
&litMatcher{
	pos: position{line: 125, col: 41, offset: 3293},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 125, col: 45, offset: 3297},
	name: "_",
},
&labeledExpr{
	pos: position{line: 125, col: 47, offset: 3299},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 125, col: 53, offset: 3305},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 125, col: 59, offset: 3311},
	name: "_",
},
&litMatcher{
	pos: position{line: 125, col: 61, offset: 3313},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "OuterJoinG",
	pos: position{line: 133, col: 1, offset: 3508},
	expr: &choiceExpr{
	pos: position{line: 133, col: 15, offset: 3522},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 133, col: 15, offset: 3522},
	name: "LeftOuterG",
},
&ruleRefExpr{
	pos: position{line: 133, col: 26, offset: 3533},
	name: "RightOuterG",
},
&ruleRefExpr{
	pos: position{line: 133, col: 38, offset: 3545},
	name: "FullOuterG",
},
	},
},
},
{
	name: "LeftOuterG",
	pos: position{line: 135, col: 1, offset: 3559},
	expr: &actionExpr{
	pos: position{line: 135, col: 15, offset: 3573},
	run: (*parser).callonLeftOuterG1,
	expr: &seqExpr{
	pos: position{line: 135, col: 15, offset: 3573},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 135, col: 15, offset: 3573},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 135, col: 21, offset: 3579},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 135, col: 30, offset: 3588},
	name: "_",
},
&litMatcher{
	pos: position{line: 135, col: 32, offset: 3590},
	val: "leftouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 135, col: 44, offset: 3602},
	name: "_",
},
&litMatcher{
	pos: position{line: 135, col: 46, offset: 3604},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 135, col: 50, offset: 3608},
	name: "_",
},
&labeledExpr{
	pos: position{line: 135, col: 52, offset: 3610},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 135, col: 58, offset: 3616},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 135, col: 64, offset: 3622},
	name: "_",
},
&litMatcher{
	pos: position{line: 135, col: 66, offset: 3624},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "RightOuterG",
	pos: position{line: 143, col: 1, offset: 3925},
	expr: &actionExpr{
	pos: position{line: 143, col: 15, offset: 3939},
	run: (*parser).callonRightOuterG1,
	expr: &seqExpr{
	pos: position{line: 143, col: 15, offset: 3939},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 143, col: 15, offset: 3939},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 143, col: 21, offset: 3945},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 143, col: 30, offset: 3954},
	name: "_",
},
&litMatcher{
	pos: position{line: 143, col: 32, offset: 3956},
	val: "rightouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 143, col: 45, offset: 3969},
	name: "_",
},
&litMatcher{
	pos: position{line: 143, col: 47, offset: 3971},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 143, col: 51, offset: 3975},
	name: "_",
},
&labeledExpr{
	pos: position{line: 143, col: 53, offset: 3977},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 143, col: 59, offset: 3983},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 143, col: 65, offset: 3989},
	name: "_",
},
&litMatcher{
	pos: position{line: 143, col: 67, offset: 3991},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FullOuterG",
	pos: position{line: 150, col: 1, offset: 4296},
	expr: &litMatcher{
	pos: position{line: 150, col: 15, offset: 4310},
	val: "shakya",
	ignoreCase: false,
},
},
{
	name: "Query",
	pos: position{line: 152, col: 1, offset: 4322},
	expr: &actionExpr{
	pos: position{line: 152, col: 10, offset: 4331},
	run: (*parser).callonQuery1,
	expr: &seqExpr{
	pos: position{line: 152, col: 10, offset: 4331},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 152, col: 10, offset: 4331},
	val: "select",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 152, col: 19, offset: 4340},
	name: "_",
},
&labeledExpr{
	pos: position{line: 152, col: 21, offset: 4342},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 152, col: 27, offset: 4348},
	name: "Fields",
},
},
&ruleRefExpr{
	pos: position{line: 152, col: 34, offset: 4355},
	name: "_",
},
&litMatcher{
	pos: position{line: 152, col: 36, offset: 4357},
	val: "on",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 152, col: 41, offset: 4362},
	name: "_",
},
&labeledExpr{
	pos: position{line: 152, col: 43, offset: 4364},
	label: "condition",
	expr: &ruleRefExpr{
	pos: position{line: 152, col: 53, offset: 4374},
	name: "JoinCondition",
},
},
&ruleRefExpr{
	pos: position{line: 152, col: 67, offset: 4388},
	name: "_",
},
	},
},
},
},
{
	name: "Fields",
	pos: position{line: 156, col: 1, offset: 4500},
	expr: &choiceExpr{
	pos: position{line: 156, col: 11, offset: 4510},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 156, col: 11, offset: 4510},
	name: "AllFields",
},
&ruleRefExpr{
	pos: position{line: 156, col: 21, offset: 4520},
	name: "SelFields",
},
	},
},
},
{
	name: "AllFields",
	pos: position{line: 158, col: 1, offset: 4533},
	expr: &actionExpr{
	pos: position{line: 158, col: 14, offset: 4546},
	run: (*parser).callonAllFields1,
	expr: &litMatcher{
	pos: position{line: 158, col: 14, offset: 4546},
	val: "*",
	ignoreCase: false,
},
},
},
{
	name: "SelFields",
	pos: position{line: 164, col: 1, offset: 4638},
	expr: &actionExpr{
	pos: position{line: 164, col: 14, offset: 4651},
	run: (*parser).callonSelFields1,
	expr: &seqExpr{
	pos: position{line: 164, col: 14, offset: 4651},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 164, col: 14, offset: 4651},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 164, col: 20, offset: 4657},
	name: "Word",
},
},
&ruleRefExpr{
	pos: position{line: 164, col: 25, offset: 4662},
	name: "_",
},
&labeledExpr{
	pos: position{line: 164, col: 28, offset: 4665},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 164, col: 33, offset: 4670},
	expr: &seqExpr{
	pos: position{line: 164, col: 34, offset: 4671},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 164, col: 34, offset: 4671},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 164, col: 38, offset: 4675},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 164, col: 40, offset: 4677},
	name: "Word",
},
&ruleRefExpr{
	pos: position{line: 164, col: 45, offset: 4682},
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
	pos: position{line: 173, col: 1, offset: 4933},
	expr: &actionExpr{
	pos: position{line: 173, col: 18, offset: 4950},
	run: (*parser).callonJoinCondition1,
	expr: &seqExpr{
	pos: position{line: 173, col: 18, offset: 4950},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 173, col: 18, offset: 4950},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 173, col: 23, offset: 4955},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 173, col: 35, offset: 4967},
	name: "_",
},
&labeledExpr{
	pos: position{line: 173, col: 37, offset: 4969},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 173, col: 40, offset: 4972},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 173, col: 52, offset: 4984},
	name: "_",
},
&labeledExpr{
	pos: position{line: 173, col: 54, offset: 4986},
	label: "right",
	expr: &ruleRefExpr{
	pos: position{line: 173, col: 60, offset: 4992},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 173, col: 72, offset: 5004},
	name: "_",
},
	},
},
},
},
{
	name: "JoinFactors",
	pos: position{line: 178, col: 1, offset: 5126},
	expr: &actionExpr{
	pos: position{line: 178, col: 16, offset: 5141},
	run: (*parser).callonJoinFactors1,
	expr: &seqExpr{
	pos: position{line: 178, col: 16, offset: 5141},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 178, col: 16, offset: 5141},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 178, col: 22, offset: 5147},
	name: "Word",
},
},
&ruleRefExpr{
	pos: position{line: 178, col: 27, offset: 5152},
	name: "_",
},
&labeledExpr{
	pos: position{line: 178, col: 29, offset: 5154},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 178, col: 34, offset: 5159},
	expr: &seqExpr{
	pos: position{line: 178, col: 35, offset: 5160},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 178, col: 35, offset: 5160},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 178, col: 39, offset: 5164},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 178, col: 41, offset: 5166},
	name: "Word",
},
&ruleRefExpr{
	pos: position{line: 178, col: 46, offset: 5171},
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
	pos: position{line: 199, col: 1, offset: 5538},
	expr: &actionExpr{
	pos: position{line: 199, col: 12, offset: 5549},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 199, col: 12, offset: 5549},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 199, col: 12, offset: 5549},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 199, col: 18, offset: 5555},
	name: "_",
},
&labeledExpr{
	pos: position{line: 199, col: 20, offset: 5557},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 199, col: 27, offset: 5564},
	expr: &ruleRefExpr{
	pos: position{line: 199, col: 27, offset: 5564},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 199, col: 39, offset: 5576},
	name: "_",
},
&labeledExpr{
	pos: position{line: 199, col: 41, offset: 5578},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 199, col: 45, offset: 5582},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 199, col: 53, offset: 5590},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 199, col: 58, offset: 5595},
	expr: &seqExpr{
	pos: position{line: 199, col: 59, offset: 5596},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 199, col: 59, offset: 5596},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 199, col: 63, offset: 5600},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 199, col: 65, offset: 5602},
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
	pos: position{line: 214, col: 1, offset: 5977},
	expr: &choiceExpr{
	pos: position{line: 214, col: 12, offset: 5988},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 214, col: 12, offset: 5988},
	name: "CountAgg",
},
&ruleRefExpr{
	pos: position{line: 214, col: 23, offset: 5999},
	name: "MaxAgg",
},
&ruleRefExpr{
	pos: position{line: 214, col: 32, offset: 6008},
	name: "MinAgg",
},
&ruleRefExpr{
	pos: position{line: 214, col: 41, offset: 6017},
	name: "AvgAgg",
},
&ruleRefExpr{
	pos: position{line: 214, col: 50, offset: 6026},
	name: "SumAgg",
},
&ruleRefExpr{
	pos: position{line: 214, col: 59, offset: 6035},
	name: "ModeAgg",
},
&ruleRefExpr{
	pos: position{line: 215, col: 13, offset: 6056},
	name: "VarAgg",
},
&ruleRefExpr{
	pos: position{line: 215, col: 22, offset: 6065},
	name: "DistinctCountAgg",
},
&ruleRefExpr{
	pos: position{line: 215, col: 41, offset: 6084},
	name: "QuantileAgg",
},
&ruleRefExpr{
	pos: position{line: 215, col: 55, offset: 6098},
	name: "MedianAgg",
},
&ruleRefExpr{
	pos: position{line: 215, col: 67, offset: 6110},
	name: "QuartileAgg",
},
&ruleRefExpr{
	pos: position{line: 216, col: 13, offset: 6135},
	name: "CovAgg",
},
&ruleRefExpr{
	pos: position{line: 216, col: 22, offset: 6144},
	name: "CorrAgg",
},
&ruleRefExpr{
	pos: position{line: 216, col: 32, offset: 6154},
	name: "PValueAgg",
},
	},
},
},
{
	name: "CountAgg",
	pos: position{line: 219, col: 1, offset: 6198},
	expr: &actionExpr{
	pos: position{line: 219, col: 13, offset: 6210},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 219, col: 13, offset: 6210},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 219, col: 13, offset: 6210},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 219, col: 19, offset: 6216},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 219, col: 28, offset: 6225},
	name: "_",
},
&litMatcher{
	pos: position{line: 219, col: 30, offset: 6227},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 219, col: 39, offset: 6236},
	name: "_",
},
&labeledExpr{
	pos: position{line: 219, col: 41, offset: 6238},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 219, col: 46, offset: 6243},
	expr: &ruleRefExpr{
	pos: position{line: 219, col: 46, offset: 6243},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 219, col: 59, offset: 6256},
	name: "_",
},
&litMatcher{
	pos: position{line: 219, col: 61, offset: 6258},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MaxAgg",
	pos: position{line: 226, col: 1, offset: 6410},
	expr: &actionExpr{
	pos: position{line: 226, col: 11, offset: 6420},
	run: (*parser).callonMaxAgg1,
	expr: &seqExpr{
	pos: position{line: 226, col: 11, offset: 6420},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 226, col: 11, offset: 6420},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 226, col: 17, offset: 6426},
	expr: &ruleRefExpr{
	pos: position{line: 226, col: 17, offset: 6426},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 226, col: 27, offset: 6436},
	name: "_",
},
&litMatcher{
	pos: position{line: 226, col: 29, offset: 6438},
	val: "max(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 226, col: 36, offset: 6445},
	name: "_",
},
&labeledExpr{
	pos: position{line: 226, col: 38, offset: 6447},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 226, col: 44, offset: 6453},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 226, col: 53, offset: 6462},
	name: "_",
},
&labeledExpr{
	pos: position{line: 226, col: 55, offset: 6464},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 226, col: 60, offset: 6469},
	expr: &seqExpr{
	pos: position{line: 226, col: 61, offset: 6470},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 226, col: 61, offset: 6470},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 226, col: 65, offset: 6474},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 226, col: 67, offset: 6476},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 226, col: 81, offset: 6490},
	name: "_",
},
&litMatcher{
	pos: position{line: 226, col: 83, offset: 6492},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MinAgg",
	pos: position{line: 248, col: 1, offset: 6953},
	expr: &actionExpr{
	pos: position{line: 248, col: 11, offset: 6963},
	run: (*parser).callonMinAgg1,
	expr: &seqExpr{
	pos: position{line: 248, col: 11, offset: 6963},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 248, col: 11, offset: 6963},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 248, col: 17, offset: 6969},
	expr: &ruleRefExpr{
	pos: position{line: 248, col: 17, offset: 6969},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 248, col: 27, offset: 6979},
	name: "_",
},
&litMatcher{
	pos: position{line: 248, col: 29, offset: 6981},
	val: "min(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 248, col: 36, offset: 6988},
	name: "_",
},
&labeledExpr{
	pos: position{line: 248, col: 38, offset: 6990},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 248, col: 44, offset: 6996},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 248, col: 53, offset: 7005},
	name: "_",
},
&labeledExpr{
	pos: position{line: 248, col: 55, offset: 7007},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 248, col: 60, offset: 7012},
	expr: &seqExpr{
	pos: position{line: 248, col: 61, offset: 7013},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 248, col: 61, offset: 7013},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 248, col: 65, offset: 7017},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 248, col: 67, offset: 7019},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 248, col: 81, offset: 7033},
	name: "_",
},
&litMatcher{
	pos: position{line: 248, col: 83, offset: 7035},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AvgAgg",
	pos: position{line: 270, col: 1, offset: 7496},
	expr: &actionExpr{
	pos: position{line: 270, col: 11, offset: 7506},
	run: (*parser).callonAvgAgg1,
	expr: &seqExpr{
	pos: position{line: 270, col: 11, offset: 7506},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 270, col: 11, offset: 7506},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 270, col: 17, offset: 7512},
	expr: &ruleRefExpr{
	pos: position{line: 270, col: 17, offset: 7512},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 270, col: 27, offset: 7522},
	name: "_",
},
&litMatcher{
	pos: position{line: 270, col: 29, offset: 7524},
	val: "avg(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 270, col: 36, offset: 7531},
	name: "_",
},
&labeledExpr{
	pos: position{line: 270, col: 38, offset: 7533},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 270, col: 44, offset: 7539},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 270, col: 53, offset: 7548},
	name: "_",
},
&labeledExpr{
	pos: position{line: 270, col: 55, offset: 7550},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 270, col: 60, offset: 7555},
	expr: &seqExpr{
	pos: position{line: 270, col: 61, offset: 7556},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 270, col: 61, offset: 7556},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 270, col: 65, offset: 7560},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 270, col: 67, offset: 7562},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 270, col: 81, offset: 7576},
	name: "_",
},
&litMatcher{
	pos: position{line: 270, col: 83, offset: 7578},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SumAgg",
	pos: position{line: 292, col: 1, offset: 8039},
	expr: &actionExpr{
	pos: position{line: 292, col: 11, offset: 8049},
	run: (*parser).callonSumAgg1,
	expr: &seqExpr{
	pos: position{line: 292, col: 11, offset: 8049},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 292, col: 11, offset: 8049},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 292, col: 17, offset: 8055},
	expr: &ruleRefExpr{
	pos: position{line: 292, col: 17, offset: 8055},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 292, col: 27, offset: 8065},
	name: "_",
},
&litMatcher{
	pos: position{line: 292, col: 29, offset: 8067},
	val: "sum(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 292, col: 36, offset: 8074},
	name: "_",
},
&labeledExpr{
	pos: position{line: 292, col: 38, offset: 8076},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 292, col: 44, offset: 8082},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 292, col: 53, offset: 8091},
	name: "_",
},
&labeledExpr{
	pos: position{line: 292, col: 55, offset: 8093},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 292, col: 60, offset: 8098},
	expr: &seqExpr{
	pos: position{line: 292, col: 61, offset: 8099},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 292, col: 61, offset: 8099},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 292, col: 65, offset: 8103},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 292, col: 67, offset: 8105},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 292, col: 81, offset: 8119},
	name: "_",
},
&litMatcher{
	pos: position{line: 292, col: 83, offset: 8121},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarAgg",
	pos: position{line: 314, col: 1, offset: 8582},
	expr: &actionExpr{
	pos: position{line: 314, col: 11, offset: 8592},
	run: (*parser).callonVarAgg1,
	expr: &seqExpr{
	pos: position{line: 314, col: 11, offset: 8592},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 314, col: 11, offset: 8592},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 314, col: 17, offset: 8598},
	expr: &ruleRefExpr{
	pos: position{line: 314, col: 17, offset: 8598},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 314, col: 27, offset: 8608},
	name: "_",
},
&litMatcher{
	pos: position{line: 314, col: 29, offset: 8610},
	val: "var(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 314, col: 36, offset: 8617},
	name: "_",
},
&labeledExpr{
	pos: position{line: 314, col: 38, offset: 8619},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 314, col: 44, offset: 8625},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 314, col: 53, offset: 8634},
	name: "_",
},
&labeledExpr{
	pos: position{line: 314, col: 55, offset: 8636},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 314, col: 60, offset: 8641},
	expr: &seqExpr{
	pos: position{line: 314, col: 61, offset: 8642},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 314, col: 61, offset: 8642},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 314, col: 65, offset: 8646},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 314, col: 67, offset: 8648},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 314, col: 81, offset: 8662},
	name: "_",
},
&litMatcher{
	pos: position{line: 314, col: 83, offset: 8664},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DistinctCountAgg",
	pos: position{line: 336, col: 1, offset: 9130},
	expr: &actionExpr{
	pos: position{line: 336, col: 21, offset: 9150},
	run: (*parser).callonDistinctCountAgg1,
	expr: &seqExpr{
	pos: position{line: 336, col: 21, offset: 9150},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 336, col: 21, offset: 9150},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 336, col: 27, offset: 9156},
	expr: &ruleRefExpr{
	pos: position{line: 336, col: 27, offset: 9156},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 336, col: 37, offset: 9166},
	name: "_",
},
&litMatcher{
	pos: position{line: 336, col: 39, offset: 9168},
	val: "dcount(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 336, col: 49, offset: 9178},
	name: "_",
},
&labeledExpr{
	pos: position{line: 336, col: 51, offset: 9180},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 336, col: 57, offset: 9186},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 336, col: 66, offset: 9195},
	name: "_",
},
&labeledExpr{
	pos: position{line: 336, col: 68, offset: 9197},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 336, col: 73, offset: 9202},
	expr: &seqExpr{
	pos: position{line: 336, col: 74, offset: 9203},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 336, col: 74, offset: 9203},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 336, col: 78, offset: 9207},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 336, col: 80, offset: 9209},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 336, col: 94, offset: 9223},
	name: "_",
},
&litMatcher{
	pos: position{line: 336, col: 96, offset: 9225},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "ModeAgg",
	pos: position{line: 358, col: 1, offset: 9696},
	expr: &actionExpr{
	pos: position{line: 358, col: 12, offset: 9707},
	run: (*parser).callonModeAgg1,
	expr: &seqExpr{
	pos: position{line: 358, col: 12, offset: 9707},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 358, col: 12, offset: 9707},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 358, col: 18, offset: 9713},
	expr: &ruleRefExpr{
	pos: position{line: 358, col: 18, offset: 9713},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 358, col: 28, offset: 9723},
	name: "_",
},
&litMatcher{
	pos: position{line: 358, col: 30, offset: 9725},
	val: "mode(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 358, col: 38, offset: 9733},
	name: "_",
},
&labeledExpr{
	pos: position{line: 358, col: 40, offset: 9735},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 358, col: 46, offset: 9741},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 358, col: 55, offset: 9750},
	name: "_",
},
&labeledExpr{
	pos: position{line: 358, col: 57, offset: 9752},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 358, col: 62, offset: 9757},
	expr: &seqExpr{
	pos: position{line: 358, col: 63, offset: 9758},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 358, col: 63, offset: 9758},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 358, col: 67, offset: 9762},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 358, col: 69, offset: 9764},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 358, col: 83, offset: 9778},
	name: "_",
},
&litMatcher{
	pos: position{line: 358, col: 85, offset: 9780},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuantileAgg",
	pos: position{line: 380, col: 1, offset: 10242},
	expr: &actionExpr{
	pos: position{line: 380, col: 16, offset: 10257},
	run: (*parser).callonQuantileAgg1,
	expr: &seqExpr{
	pos: position{line: 380, col: 16, offset: 10257},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 380, col: 16, offset: 10257},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 380, col: 22, offset: 10263},
	expr: &ruleRefExpr{
	pos: position{line: 380, col: 22, offset: 10263},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 380, col: 32, offset: 10273},
	name: "_",
},
&litMatcher{
	pos: position{line: 380, col: 34, offset: 10275},
	val: "quantile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 380, col: 46, offset: 10287},
	name: "_",
},
&labeledExpr{
	pos: position{line: 380, col: 48, offset: 10289},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 380, col: 54, offset: 10295},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 380, col: 63, offset: 10304},
	name: "_",
},
&labeledExpr{
	pos: position{line: 380, col: 66, offset: 10307},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 380, col: 73, offset: 10314},
	expr: &seqExpr{
	pos: position{line: 380, col: 74, offset: 10315},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 380, col: 74, offset: 10315},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 380, col: 78, offset: 10319},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 380, col: 80, offset: 10321},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 380, col: 91, offset: 10332},
	name: "_",
},
&litMatcher{
	pos: position{line: 380, col: 93, offset: 10334},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 380, col: 97, offset: 10338},
	name: "_",
},
&labeledExpr{
	pos: position{line: 380, col: 99, offset: 10340},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 380, col: 103, offset: 10344},
	name: "Float",
},
},
&ruleRefExpr{
	pos: position{line: 380, col: 109, offset: 10350},
	name: "_",
},
&labeledExpr{
	pos: position{line: 380, col: 111, offset: 10352},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 380, col: 116, offset: 10357},
	expr: &seqExpr{
	pos: position{line: 380, col: 117, offset: 10358},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 380, col: 117, offset: 10358},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 380, col: 121, offset: 10362},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 380, col: 123, offset: 10364},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 380, col: 137, offset: 10378},
	name: "_",
},
&litMatcher{
	pos: position{line: 380, col: 139, offset: 10380},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MedianAgg",
	pos: position{line: 410, col: 1, offset: 11025},
	expr: &actionExpr{
	pos: position{line: 410, col: 14, offset: 11038},
	run: (*parser).callonMedianAgg1,
	expr: &seqExpr{
	pos: position{line: 410, col: 14, offset: 11038},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 410, col: 14, offset: 11038},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 410, col: 20, offset: 11044},
	expr: &ruleRefExpr{
	pos: position{line: 410, col: 20, offset: 11044},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 410, col: 30, offset: 11054},
	name: "_",
},
&litMatcher{
	pos: position{line: 410, col: 32, offset: 11056},
	val: "median(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 410, col: 42, offset: 11066},
	name: "_",
},
&labeledExpr{
	pos: position{line: 410, col: 44, offset: 11068},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 410, col: 50, offset: 11074},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 410, col: 59, offset: 11083},
	name: "_",
},
&labeledExpr{
	pos: position{line: 410, col: 62, offset: 11086},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 410, col: 69, offset: 11093},
	expr: &seqExpr{
	pos: position{line: 410, col: 70, offset: 11094},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 410, col: 70, offset: 11094},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 410, col: 74, offset: 11098},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 410, col: 76, offset: 11100},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 410, col: 87, offset: 11111},
	name: "_",
},
&labeledExpr{
	pos: position{line: 410, col: 89, offset: 11113},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 410, col: 94, offset: 11118},
	expr: &seqExpr{
	pos: position{line: 410, col: 95, offset: 11119},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 410, col: 95, offset: 11119},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 410, col: 99, offset: 11123},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 410, col: 101, offset: 11125},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 410, col: 115, offset: 11139},
	name: "_",
},
&litMatcher{
	pos: position{line: 410, col: 117, offset: 11141},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuartileAgg",
	pos: position{line: 442, col: 1, offset: 11818},
	expr: &actionExpr{
	pos: position{line: 442, col: 16, offset: 11833},
	run: (*parser).callonQuartileAgg1,
	expr: &seqExpr{
	pos: position{line: 442, col: 16, offset: 11833},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 442, col: 16, offset: 11833},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 442, col: 22, offset: 11839},
	expr: &ruleRefExpr{
	pos: position{line: 442, col: 22, offset: 11839},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 442, col: 32, offset: 11849},
	name: "_",
},
&litMatcher{
	pos: position{line: 442, col: 34, offset: 11851},
	val: "quartile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 442, col: 46, offset: 11863},
	name: "_",
},
&labeledExpr{
	pos: position{line: 442, col: 48, offset: 11865},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 442, col: 54, offset: 11871},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 442, col: 63, offset: 11880},
	name: "_",
},
&labeledExpr{
	pos: position{line: 442, col: 66, offset: 11883},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 442, col: 73, offset: 11890},
	expr: &seqExpr{
	pos: position{line: 442, col: 74, offset: 11891},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 442, col: 74, offset: 11891},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 442, col: 78, offset: 11895},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 442, col: 80, offset: 11897},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 442, col: 91, offset: 11908},
	name: "_",
},
&litMatcher{
	pos: position{line: 442, col: 93, offset: 11910},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 442, col: 97, offset: 11914},
	name: "_",
},
&labeledExpr{
	pos: position{line: 442, col: 99, offset: 11916},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 442, col: 103, offset: 11920},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 442, col: 110, offset: 11927},
	name: "_",
},
&labeledExpr{
	pos: position{line: 442, col: 112, offset: 11929},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 442, col: 117, offset: 11934},
	expr: &seqExpr{
	pos: position{line: 442, col: 118, offset: 11935},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 442, col: 118, offset: 11935},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 442, col: 122, offset: 11939},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 442, col: 124, offset: 11941},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 442, col: 138, offset: 11955},
	name: "_",
},
&litMatcher{
	pos: position{line: 442, col: 140, offset: 11957},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CovAgg",
	pos: position{line: 488, col: 1, offset: 12984},
	expr: &actionExpr{
	pos: position{line: 488, col: 11, offset: 12994},
	run: (*parser).callonCovAgg1,
	expr: &seqExpr{
	pos: position{line: 488, col: 11, offset: 12994},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 488, col: 11, offset: 12994},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 488, col: 17, offset: 13000},
	expr: &ruleRefExpr{
	pos: position{line: 488, col: 17, offset: 13000},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 488, col: 27, offset: 13010},
	name: "_",
},
&litMatcher{
	pos: position{line: 488, col: 29, offset: 13012},
	val: "cov(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 488, col: 36, offset: 13019},
	name: "_",
},
&labeledExpr{
	pos: position{line: 488, col: 38, offset: 13021},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 488, col: 45, offset: 13028},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 488, col: 54, offset: 13037},
	name: "_",
},
&litMatcher{
	pos: position{line: 488, col: 56, offset: 13039},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 488, col: 60, offset: 13043},
	name: "_",
},
&labeledExpr{
	pos: position{line: 488, col: 62, offset: 13045},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 488, col: 69, offset: 13052},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 488, col: 78, offset: 13061},
	name: "_",
},
&labeledExpr{
	pos: position{line: 488, col: 80, offset: 13063},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 488, col: 85, offset: 13068},
	expr: &seqExpr{
	pos: position{line: 488, col: 86, offset: 13069},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 488, col: 86, offset: 13069},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 488, col: 90, offset: 13073},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 488, col: 92, offset: 13075},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 488, col: 106, offset: 13089},
	name: "_",
},
&litMatcher{
	pos: position{line: 488, col: 108, offset: 13091},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CorrAgg",
	pos: position{line: 511, col: 1, offset: 13616},
	expr: &actionExpr{
	pos: position{line: 511, col: 12, offset: 13627},
	run: (*parser).callonCorrAgg1,
	expr: &seqExpr{
	pos: position{line: 511, col: 12, offset: 13627},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 511, col: 12, offset: 13627},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 511, col: 18, offset: 13633},
	expr: &ruleRefExpr{
	pos: position{line: 511, col: 18, offset: 13633},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 511, col: 28, offset: 13643},
	name: "_",
},
&litMatcher{
	pos: position{line: 511, col: 30, offset: 13645},
	val: "correlation(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 511, col: 45, offset: 13660},
	name: "_",
},
&labeledExpr{
	pos: position{line: 511, col: 47, offset: 13662},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 511, col: 54, offset: 13669},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 511, col: 63, offset: 13678},
	name: "_",
},
&litMatcher{
	pos: position{line: 511, col: 65, offset: 13680},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 511, col: 69, offset: 13684},
	name: "_",
},
&labeledExpr{
	pos: position{line: 511, col: 71, offset: 13686},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 511, col: 78, offset: 13693},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 511, col: 87, offset: 13702},
	name: "_",
},
&labeledExpr{
	pos: position{line: 511, col: 89, offset: 13704},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 511, col: 94, offset: 13709},
	expr: &seqExpr{
	pos: position{line: 511, col: 95, offset: 13710},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 511, col: 95, offset: 13710},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 511, col: 99, offset: 13714},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 511, col: 101, offset: 13716},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 511, col: 115, offset: 13730},
	name: "_",
},
&litMatcher{
	pos: position{line: 511, col: 117, offset: 13732},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PValueAgg",
	pos: position{line: 534, col: 1, offset: 14258},
	expr: &actionExpr{
	pos: position{line: 534, col: 14, offset: 14271},
	run: (*parser).callonPValueAgg1,
	expr: &seqExpr{
	pos: position{line: 534, col: 14, offset: 14271},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 534, col: 14, offset: 14271},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 534, col: 20, offset: 14277},
	expr: &ruleRefExpr{
	pos: position{line: 534, col: 20, offset: 14277},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 534, col: 30, offset: 14287},
	name: "_",
},
&litMatcher{
	pos: position{line: 534, col: 32, offset: 14289},
	val: "pvalue(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 534, col: 42, offset: 14299},
	name: "_",
},
&labeledExpr{
	pos: position{line: 534, col: 44, offset: 14301},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 534, col: 51, offset: 14308},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 534, col: 60, offset: 14317},
	name: "_",
},
&litMatcher{
	pos: position{line: 534, col: 62, offset: 14319},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 534, col: 66, offset: 14323},
	name: "_",
},
&labeledExpr{
	pos: position{line: 534, col: 68, offset: 14325},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 534, col: 75, offset: 14332},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 534, col: 84, offset: 14341},
	name: "_",
},
&labeledExpr{
	pos: position{line: 534, col: 86, offset: 14343},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 534, col: 91, offset: 14348},
	expr: &seqExpr{
	pos: position{line: 534, col: 92, offset: 14349},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 534, col: 92, offset: 14349},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 534, col: 96, offset: 14353},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 534, col: 98, offset: 14355},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 534, col: 112, offset: 14369},
	name: "_",
},
&litMatcher{
	pos: position{line: 534, col: 114, offset: 14371},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggGroupBy",
	pos: position{line: 559, col: 1, offset: 14896},
	expr: &actionExpr{
	pos: position{line: 559, col: 15, offset: 14910},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 559, col: 15, offset: 14910},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 559, col: 15, offset: 14910},
	name: "_",
},
&litMatcher{
	pos: position{line: 559, col: 17, offset: 14912},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 559, col: 22, offset: 14917},
	name: "_",
},
&labeledExpr{
	pos: position{line: 559, col: 24, offset: 14919},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 559, col: 30, offset: 14925},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 559, col: 39, offset: 14934},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 559, col: 44, offset: 14939},
	expr: &seqExpr{
	pos: position{line: 559, col: 45, offset: 14940},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 559, col: 45, offset: 14940},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 559, col: 49, offset: 14944},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 559, col: 51, offset: 14946},
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
	pos: position{line: 574, col: 1, offset: 15297},
	expr: &actionExpr{
	pos: position{line: 574, col: 11, offset: 15307},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 574, col: 11, offset: 15307},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 574, col: 16, offset: 15312},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 574, col: 16, offset: 15312},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 574, col: 27, offset: 15323},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 574, col: 38, offset: 15334},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 574, col: 49, offset: 15345},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 574, col: 58, offset: 15354},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 574, col: 67, offset: 15363},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 580, col: 1, offset: 15482},
	expr: &actionExpr{
	pos: position{line: 580, col: 13, offset: 15494},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 580, col: 13, offset: 15494},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 580, col: 13, offset: 15494},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 580, col: 23, offset: 15504},
	name: "_",
},
&labeledExpr{
	pos: position{line: 580, col: 25, offset: 15506},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 580, col: 28, offset: 15509},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 580, col: 42, offset: 15523},
	name: "_",
},
&litMatcher{
	pos: position{line: 580, col: 44, offset: 15525},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 587, col: 1, offset: 15720},
	expr: &actionExpr{
	pos: position{line: 587, col: 18, offset: 15737},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 587, col: 18, offset: 15737},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 587, col: 18, offset: 15737},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 587, col: 23, offset: 15742},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 587, col: 28, offset: 15747},
	name: "_",
},
&labeledExpr{
	pos: position{line: 587, col: 30, offset: 15749},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 587, col: 33, offset: 15752},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 587, col: 45, offset: 15764},
	name: "_",
},
&labeledExpr{
	pos: position{line: 587, col: 47, offset: 15766},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 587, col: 54, offset: 15773},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 591, col: 1, offset: 15824},
	expr: &ruleRefExpr{
	pos: position{line: 591, col: 16, offset: 15839},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 594, col: 1, offset: 15873},
	expr: &actionExpr{
	pos: position{line: 594, col: 13, offset: 15885},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 594, col: 13, offset: 15885},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 594, col: 13, offset: 15885},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 594, col: 23, offset: 15895},
	name: "_",
},
&labeledExpr{
	pos: position{line: 594, col: 25, offset: 15897},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 594, col: 31, offset: 15903},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 594, col: 40, offset: 15912},
	name: "_",
},
&labeledExpr{
	pos: position{line: 594, col: 42, offset: 15914},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 594, col: 47, offset: 15919},
	expr: &seqExpr{
	pos: position{line: 594, col: 48, offset: 15920},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 594, col: 48, offset: 15920},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 594, col: 52, offset: 15924},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 594, col: 54, offset: 15926},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 594, col: 65, offset: 15937},
	name: "_",
},
&litMatcher{
	pos: position{line: 594, col: 67, offset: 15939},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 609, col: 1, offset: 16276},
	expr: &actionExpr{
	pos: position{line: 609, col: 11, offset: 16286},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 609, col: 11, offset: 16286},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 609, col: 11, offset: 16286},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 609, col: 18, offset: 16293},
	name: "_",
},
&labeledExpr{
	pos: position{line: 609, col: 20, offset: 16295},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 609, col: 25, offset: 16300},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 609, col: 34, offset: 16309},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 609, col: 38, offset: 16313},
	name: "_",
},
&labeledExpr{
	pos: position{line: 609, col: 40, offset: 16315},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 609, col: 44, offset: 16319},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 609, col: 51, offset: 16326},
	name: "_",
},
&litMatcher{
	pos: position{line: 609, col: 53, offset: 16328},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 621, col: 1, offset: 16550},
	expr: &actionExpr{
	pos: position{line: 621, col: 11, offset: 16560},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 621, col: 11, offset: 16560},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 621, col: 11, offset: 16560},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 621, col: 20, offset: 16569},
	name: "_",
},
&litMatcher{
	pos: position{line: 621, col: 22, offset: 16571},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 621, col: 27, offset: 16576},
	name: "_",
},
&labeledExpr{
	pos: position{line: 621, col: 29, offset: 16578},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 621, col: 35, offset: 16584},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 627, col: 1, offset: 16690},
	expr: &actionExpr{
	pos: position{line: 627, col: 12, offset: 16701},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 627, col: 12, offset: 16701},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 627, col: 12, offset: 16701},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 627, col: 21, offset: 16710},
	name: "_",
},
&labeledExpr{
	pos: position{line: 627, col: 23, offset: 16712},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 627, col: 27, offset: 16716},
	expr: &ruleRefExpr{
	pos: position{line: 627, col: 27, offset: 16716},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 627, col: 35, offset: 16724},
	name: "_",
},
&litMatcher{
	pos: position{line: 627, col: 37, offset: 16726},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 637, col: 1, offset: 16871},
	expr: &actionExpr{
	pos: position{line: 637, col: 13, offset: 16883},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 637, col: 13, offset: 16883},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 637, col: 13, offset: 16883},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 637, col: 23, offset: 16893},
	name: "_",
},
&labeledExpr{
	pos: position{line: 637, col: 25, offset: 16895},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 637, col: 30, offset: 16900},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 637, col: 42, offset: 16912},
	name: "_",
},
&litMatcher{
	pos: position{line: 637, col: 44, offset: 16914},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 647, col: 1, offset: 17202},
	expr: &actionExpr{
	pos: position{line: 647, col: 16, offset: 17217},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 647, col: 16, offset: 17217},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 647, col: 16, offset: 17217},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 647, col: 22, offset: 17223},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 647, col: 35, offset: 17236},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 647, col: 40, offset: 17241},
	expr: &seqExpr{
	pos: position{line: 647, col: 41, offset: 17242},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 647, col: 41, offset: 17242},
	name: "_",
},
&labeledExpr{
	pos: position{line: 647, col: 43, offset: 17244},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 647, col: 46, offset: 17247},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 647, col: 58, offset: 17259},
	name: "_",
},
&labeledExpr{
	pos: position{line: 647, col: 60, offset: 17261},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 647, col: 63, offset: 17264},
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
	pos: position{line: 652, col: 1, offset: 17422},
	expr: &choiceExpr{
	pos: position{line: 652, col: 17, offset: 17438},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 652, col: 17, offset: 17438},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 652, col: 17, offset: 17438},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 652, col: 17, offset: 17438},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 652, col: 21, offset: 17442},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 652, col: 24, offset: 17445},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 652, col: 36, offset: 17457},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 654, col: 5, offset: 17488},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 654, col: 5, offset: 17488},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 654, col: 8, offset: 17491},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 658, col: 1, offset: 17533},
	expr: &choiceExpr{
	pos: position{line: 658, col: 16, offset: 17548},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 658, col: 16, offset: 17548},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 658, col: 24, offset: 17556},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 658, col: 24, offset: 17556},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 665, col: 1, offset: 17735},
	expr: &actionExpr{
	pos: position{line: 665, col: 18, offset: 17752},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 665, col: 18, offset: 17752},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 665, col: 18, offset: 17752},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 665, col: 23, offset: 17757},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 665, col: 28, offset: 17762},
	name: "_",
},
&labeledExpr{
	pos: position{line: 665, col: 30, offset: 17764},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 665, col: 33, offset: 17767},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 665, col: 45, offset: 17779},
	name: "_",
},
&labeledExpr{
	pos: position{line: 665, col: 47, offset: 17781},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 665, col: 53, offset: 17787},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 665, col: 59, offset: 17793},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 669, col: 1, offset: 17867},
	expr: &choiceExpr{
	pos: position{line: 669, col: 16, offset: 17882},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 669, col: 16, offset: 17882},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 669, col: 23, offset: 17889},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 669, col: 30, offset: 17896},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 669, col: 36, offset: 17902},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 669, col: 42, offset: 17908},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 669, col: 49, offset: 17915},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 669, col: 49, offset: 17915},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 678, col: 1, offset: 18025},
	expr: &actionExpr{
	pos: position{line: 678, col: 12, offset: 18036},
	run: (*parser).callonStdOutG1,
	expr: &seqExpr{
	pos: position{line: 678, col: 12, offset: 18036},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 678, col: 12, offset: 18036},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 678, col: 22, offset: 18046},
	name: "_",
},
&labeledExpr{
	pos: position{line: 678, col: 24, offset: 18048},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 678, col: 31, offset: 18055},
	expr: &ruleRefExpr{
	pos: position{line: 678, col: 32, offset: 18056},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 678, col: 48, offset: 18072},
	name: "_",
},
&litMatcher{
	pos: position{line: 678, col: 50, offset: 18074},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SinkInto",
	pos: position{line: 692, col: 1, offset: 18353},
	expr: &actionExpr{
	pos: position{line: 692, col: 12, offset: 18364},
	run: (*parser).callonSinkInto1,
	expr: &seqExpr{
	pos: position{line: 692, col: 12, offset: 18364},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 692, col: 12, offset: 18364},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 692, col: 18, offset: 18370},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 692, col: 27, offset: 18379},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 692, col: 33, offset: 18385},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 692, col: 33, offset: 18385},
	name: "_",
},
&litMatcher{
	pos: position{line: 692, col: 35, offset: 18387},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 692, col: 39, offset: 18391},
	name: "_",
},
&litMatcher{
	pos: position{line: 692, col: 41, offset: 18393},
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
	pos: position{line: 700, col: 1, offset: 18549},
	expr: &actionExpr{
	pos: position{line: 700, col: 15, offset: 18563},
	run: (*parser).callonBlackHoleG1,
	expr: &seqExpr{
	pos: position{line: 700, col: 15, offset: 18563},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 700, col: 15, offset: 18563},
	val: "blackhole(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 700, col: 28, offset: 18576},
	name: "_",
},
&labeledExpr{
	pos: position{line: 700, col: 30, offset: 18578},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 700, col: 37, offset: 18585},
	expr: &ruleRefExpr{
	pos: position{line: 700, col: 38, offset: 18586},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 700, col: 54, offset: 18602},
	name: "_",
},
&litMatcher{
	pos: position{line: 700, col: 56, offset: 18604},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PlotG",
	pos: position{line: 704, col: 1, offset: 18662},
	expr: &actionExpr{
	pos: position{line: 704, col: 10, offset: 18671},
	run: (*parser).callonPlotG1,
	expr: &seqExpr{
	pos: position{line: 704, col: 10, offset: 18671},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 704, col: 10, offset: 18671},
	val: "plot",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 704, col: 17, offset: 18678},
	name: "_",
},
&labeledExpr{
	pos: position{line: 704, col: 19, offset: 18680},
	label: "widget",
	expr: &ruleRefExpr{
	pos: position{line: 704, col: 27, offset: 18688},
	name: "BarWidget",
},
},
	},
},
},
},
{
	name: "BarWidget",
	pos: position{line: 708, col: 1, offset: 18788},
	expr: &actionExpr{
	pos: position{line: 708, col: 14, offset: 18801},
	run: (*parser).callonBarWidget1,
	expr: &seqExpr{
	pos: position{line: 708, col: 14, offset: 18801},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 708, col: 14, offset: 18801},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 708, col: 20, offset: 18807},
	expr: &ruleRefExpr{
	pos: position{line: 708, col: 20, offset: 18807},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 708, col: 30, offset: 18817},
	name: "_",
},
&litMatcher{
	pos: position{line: 708, col: 32, offset: 18819},
	val: "bar(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 708, col: 39, offset: 18826},
	name: "_",
},
&labeledExpr{
	pos: position{line: 708, col: 41, offset: 18828},
	label: "xField",
	expr: &ruleRefExpr{
	pos: position{line: 708, col: 48, offset: 18835},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 708, col: 57, offset: 18844},
	name: "_",
},
&litMatcher{
	pos: position{line: 708, col: 59, offset: 18846},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 708, col: 63, offset: 18850},
	name: "_",
},
&labeledExpr{
	pos: position{line: 708, col: 65, offset: 18852},
	label: "yField",
	expr: &ruleRefExpr{
	pos: position{line: 708, col: 72, offset: 18859},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 708, col: 81, offset: 18868},
	name: "_",
},
&labeledExpr{
	pos: position{line: 708, col: 83, offset: 18870},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 708, col: 88, offset: 18875},
	expr: &seqExpr{
	pos: position{line: 708, col: 89, offset: 18876},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 708, col: 89, offset: 18876},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 708, col: 93, offset: 18880},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 708, col: 95, offset: 18882},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 708, col: 102, offset: 18889},
	name: "_",
},
&zeroOrOneExpr{
	pos: position{line: 708, col: 104, offset: 18891},
	expr: &seqExpr{
	pos: position{line: 708, col: 105, offset: 18892},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 708, col: 105, offset: 18892},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 708, col: 109, offset: 18896},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 708, col: 111, offset: 18898},
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
	pos: position{line: 708, col: 122, offset: 18909},
	name: "_",
},
&litMatcher{
	pos: position{line: 708, col: 124, offset: 18911},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FileName",
	pos: position{line: 741, col: 1, offset: 19596},
	expr: &actionExpr{
	pos: position{line: 741, col: 13, offset: 19608},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 741, col: 13, offset: 19608},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 741, col: 13, offset: 19608},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 741, col: 22, offset: 19617},
	expr: &seqExpr{
	pos: position{line: 741, col: 23, offset: 19618},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 741, col: 23, offset: 19618},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 741, col: 27, offset: 19622},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 741, col: 31, offset: 19626},
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
	pos: position{line: 747, col: 1, offset: 19727},
	expr: &actionExpr{
	pos: position{line: 747, col: 10, offset: 19736},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 747, col: 10, offset: 19736},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 747, col: 10, offset: 19736},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 747, col: 16, offset: 19742},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 747, col: 23, offset: 19749},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 747, col: 28, offset: 19754},
	expr: &seqExpr{
	pos: position{line: 747, col: 29, offset: 19755},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 747, col: 29, offset: 19755},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 747, col: 31, offset: 19757},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 747, col: 36, offset: 19762},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 747, col: 38, offset: 19764},
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
	pos: position{line: 751, col: 1, offset: 19829},
	expr: &actionExpr{
	pos: position{line: 751, col: 11, offset: 19839},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 751, col: 11, offset: 19839},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 751, col: 16, offset: 19844},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 751, col: 16, offset: 19844},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 751, col: 24, offset: 19852},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 751, col: 33, offset: 19861},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 751, col: 48, offset: 19876},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 751, col: 59, offset: 19887},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 755, col: 1, offset: 19929},
	expr: &choiceExpr{
	pos: position{line: 755, col: 9, offset: 19937},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 755, col: 9, offset: 19937},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 755, col: 15, offset: 19943},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 755, col: 21, offset: 19949},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 755, col: 27, offset: 19955},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 755, col: 33, offset: 19961},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 755, col: 33, offset: 19961},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 759, col: 1, offset: 20009},
	expr: &choiceExpr{
	pos: position{line: 759, col: 12, offset: 20020},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 759, col: 12, offset: 20020},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 759, col: 12, offset: 20020},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 759, col: 12, offset: 20020},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 759, col: 16, offset: 20024},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 759, col: 19, offset: 20027},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 759, col: 25, offset: 20033},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 763, col: 5, offset: 20109},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 763, col: 5, offset: 20109},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 763, col: 8, offset: 20112},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 770, col: 1, offset: 20189},
	expr: &actionExpr{
	pos: position{line: 770, col: 9, offset: 20197},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 770, col: 9, offset: 20197},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 770, col: 9, offset: 20197},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 770, col: 15, offset: 20203},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 770, col: 21, offset: 20209},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 770, col: 26, offset: 20214},
	expr: &seqExpr{
	pos: position{line: 770, col: 27, offset: 20215},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 770, col: 27, offset: 20215},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 770, col: 29, offset: 20217},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 770, col: 34, offset: 20222},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 770, col: 36, offset: 20224},
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
	pos: position{line: 774, col: 1, offset: 20288},
	expr: &actionExpr{
	pos: position{line: 774, col: 10, offset: 20297},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 774, col: 10, offset: 20297},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 774, col: 15, offset: 20302},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 774, col: 15, offset: 20302},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 774, col: 23, offset: 20310},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 774, col: 32, offset: 20319},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 774, col: 47, offset: 20334},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 778, col: 1, offset: 20377},
	expr: &actionExpr{
	pos: position{line: 778, col: 13, offset: 20389},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 778, col: 13, offset: 20389},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 778, col: 13, offset: 20389},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 778, col: 20, offset: 20396},
	expr: &choiceExpr{
	pos: position{line: 778, col: 21, offset: 20397},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 778, col: 21, offset: 20397},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 778, col: 31, offset: 20407},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 778, col: 39, offset: 20415},
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
	pos: position{line: 782, col: 1, offset: 20465},
	expr: &actionExpr{
	pos: position{line: 782, col: 17, offset: 20481},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 782, col: 17, offset: 20481},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 782, col: 17, offset: 20481},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 782, col: 21, offset: 20485},
	expr: &choiceExpr{
	pos: position{line: 782, col: 23, offset: 20487},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 782, col: 23, offset: 20487},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 782, col: 23, offset: 20487},
	expr: &ruleRefExpr{
	pos: position{line: 782, col: 24, offset: 20488},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 782, col: 36, offset: 20500,
},
	},
},
&seqExpr{
	pos: position{line: 782, col: 40, offset: 20504},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 782, col: 40, offset: 20504},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 782, col: 45, offset: 20509},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 782, col: 63, offset: 20527},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 789, col: 1, offset: 20708},
	expr: &charClassMatcher{
	pos: position{line: 789, col: 16, offset: 20723},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 791, col: 1, offset: 20741},
	expr: &ruleRefExpr{
	pos: position{line: 791, col: 19, offset: 20759},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 793, col: 1, offset: 20779},
	expr: &charClassMatcher{
	pos: position{line: 793, col: 21, offset: 20799},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 795, col: 1, offset: 20814},
	expr: &actionExpr{
	pos: position{line: 795, col: 9, offset: 20822},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 795, col: 9, offset: 20822},
	expr: &choiceExpr{
	pos: position{line: 795, col: 10, offset: 20823},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 795, col: 10, offset: 20823},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 795, col: 19, offset: 20832},
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
	pos: position{line: 799, col: 1, offset: 20882},
	expr: &choiceExpr{
	pos: position{line: 799, col: 11, offset: 20892},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 799, col: 11, offset: 20892},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 799, col: 19, offset: 20900},
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
	pos: position{line: 801, col: 1, offset: 20909},
	expr: &actionExpr{
	pos: position{line: 801, col: 11, offset: 20919},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 801, col: 11, offset: 20919},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 801, col: 11, offset: 20919},
	expr: &litMatcher{
	pos: position{line: 801, col: 11, offset: 20919},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 801, col: 16, offset: 20924},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 805, col: 1, offset: 20997},
	expr: &actionExpr{
	pos: position{line: 805, col: 19, offset: 21015},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 805, col: 19, offset: 21015},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 809, col: 1, offset: 21088},
	expr: &actionExpr{
	pos: position{line: 809, col: 10, offset: 21097},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 809, col: 10, offset: 21097},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 809, col: 10, offset: 21097},
	expr: &litMatcher{
	pos: position{line: 809, col: 10, offset: 21097},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 809, col: 15, offset: 21102},
	expr: &ruleRefExpr{
	pos: position{line: 809, col: 15, offset: 21102},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 809, col: 24, offset: 21111},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 809, col: 28, offset: 21115},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 813, col: 1, offset: 21182},
	expr: &choiceExpr{
	pos: position{line: 813, col: 12, offset: 21193},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 813, col: 12, offset: 21193},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 813, col: 18, offset: 21199},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 813, col: 18, offset: 21199},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 813, col: 38, offset: 21219},
	expr: &ruleRefExpr{
	pos: position{line: 813, col: 38, offset: 21219},
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
	pos: position{line: 815, col: 1, offset: 21236},
	expr: &charClassMatcher{
	pos: position{line: 815, col: 17, offset: 21252},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 817, col: 1, offset: 21261},
	expr: &charClassMatcher{
	pos: position{line: 817, col: 24, offset: 21284},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 819, col: 1, offset: 21293},
	expr: &zeroOrMoreExpr{
	pos: position{line: 819, col: 19, offset: 21311},
	expr: &charClassMatcher{
	pos: position{line: 819, col: 19, offset: 21311},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 821, col: 1, offset: 21325},
	expr: &charClassMatcher{
	pos: position{line: 821, col: 18, offset: 21342},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "IDENTIFIER",
	pos: position{line: 822, col: 1, offset: 21353},
	expr: &actionExpr{
	pos: position{line: 822, col: 15, offset: 21367},
	run: (*parser).callonIDENTIFIER1,
	expr: &seqExpr{
	pos: position{line: 822, col: 15, offset: 21367},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 822, col: 15, offset: 21367},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 822, col: 22, offset: 21374},
	expr: &choiceExpr{
	pos: position{line: 822, col: 23, offset: 21375},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 822, col: 23, offset: 21375},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 822, col: 33, offset: 21385},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 822, col: 41, offset: 21393},
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
	pos: position{line: 825, col: 1, offset: 21466},
	expr: &notExpr{
	pos: position{line: 825, col: 8, offset: 21473},
	expr: &anyMatcher{
	line: 825, col: 9, offset: 21474,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 827, col: 1, offset: 21479},
	expr: &litMatcher{
	pos: position{line: 827, col: 17, offset: 21495},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 833, col: 1, offset: 21582},
	expr: &actionExpr{
	pos: position{line: 833, col: 18, offset: 21599},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 833, col: 18, offset: 21599},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 833, col: 18, offset: 21599},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 833, col: 24, offset: 21605},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 833, col: 33, offset: 21614},
	name: "_",
},
&labeledExpr{
	pos: position{line: 833, col: 35, offset: 21616},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 833, col: 40, offset: 21621},
	expr: &seqExpr{
	pos: position{line: 833, col: 41, offset: 21622},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 833, col: 41, offset: 21622},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 833, col: 45, offset: 21626},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 833, col: 47, offset: 21628},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 833, col: 56, offset: 21637},
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
	pos: position{line: 836, col: 1, offset: 21690},
	expr: &actionExpr{
	pos: position{line: 836, col: 13, offset: 21702},
	run: (*parser).callonEqAliasG1,
	expr: &seqExpr{
	pos: position{line: 836, col: 13, offset: 21702},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 836, col: 13, offset: 21702},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 836, col: 19, offset: 21708},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 836, col: 28, offset: 21717},
	name: "_",
},
&litMatcher{
	pos: position{line: 836, col: 30, offset: 21719},
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

func (c *current) onLeftOuterG1(alias, query interface{}) (interface{}, error) {

                   name,err := parseJoinArgs(alias)
                   if err !=nil{
                      return nil,err
                      }
                   return JoinNodeJob{Alias:name,Type:"OUTER",Attributes:query.(JoinAttributes),JoinSubType:"LEFTOUTER"},nil
              
}

func (p *parser) callonLeftOuterG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLeftOuterG1(stack["alias"], stack["query"])
}

func (c *current) onRightOuterG1(alias, query interface{}) (interface{}, error) {

                   name,err := parseJoinArgs(alias)
                   if err !=nil{
                      return nil,err
                      }
                   return JoinNodeJob{Alias:name,Type:"OUTER",Attributes:query.(JoinAttributes),JoinSubType:"RIGHTOUTER"},nil
                   
}

func (p *parser) callonRightOuterG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRightOuterG1(stack["alias"], stack["query"])
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

