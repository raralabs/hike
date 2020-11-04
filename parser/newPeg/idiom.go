
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
	pos: position{line: 25, col: 1, offset: 620},
	expr: &actionExpr{
	pos: position{line: 25, col: 14, offset: 633},
	run: (*parser).callonStatement1,
	expr: &seqExpr{
	pos: position{line: 25, col: 14, offset: 633},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 25, col: 14, offset: 633},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 25, col: 24, offset: 643},
	name: "Source",
},
},
&labeledExpr{
	pos: position{line: 25, col: 31, offset: 650},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 25, col: 37, offset: 656},
	exprs: []interface{}{
&zeroOrMoreExpr{
	pos: position{line: 25, col: 37, offset: 656},
	expr: &seqExpr{
	pos: position{line: 25, col: 38, offset: 657},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 25, col: 38, offset: 657},
	name: "_",
},
&litMatcher{
	pos: position{line: 25, col: 40, offset: 659},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 25, col: 44, offset: 663},
	name: "_",
},
&labeledExpr{
	pos: position{line: 25, col: 46, offset: 665},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 25, col: 56, offset: 675},
	name: "Transform",
},
},
	},
},
},
&ruleRefExpr{
	pos: position{line: 25, col: 68, offset: 687},
	name: "_",
},
&litMatcher{
	pos: position{line: 25, col: 69, offset: 688},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 25, col: 72, offset: 691},
	name: "_",
},
&labeledExpr{
	pos: position{line: 25, col: 74, offset: 693},
	label: "statement",
	expr: &ruleRefExpr{
	pos: position{line: 25, col: 84, offset: 703},
	name: "Sink",
},
},
&ruleRefExpr{
	pos: position{line: 25, col: 89, offset: 708},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 25, col: 91, offset: 710},
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
	pos: position{line: 63, col: 1, offset: 1387},
	expr: &choiceExpr{
	pos: position{line: 63, col: 11, offset: 1397},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 63, col: 11, offset: 1397},
	name: "FileJobG",
},
&ruleRefExpr{
	pos: position{line: 63, col: 20, offset: 1406},
	name: "BranchJobG",
},
	},
},
},
{
	name: "Transform",
	pos: position{line: 65, col: 1, offset: 1420},
	expr: &choiceExpr{
	pos: position{line: 65, col: 14, offset: 1433},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 65, col: 14, offset: 1433},
	name: "DoJobG",
},
&ruleRefExpr{
	pos: position{line: 65, col: 21, offset: 1440},
	name: "AggJobG",
},
	},
},
},
{
	name: "Sink",
	pos: position{line: 67, col: 1, offset: 1451},
	expr: &choiceExpr{
	pos: position{line: 67, col: 9, offset: 1459},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 67, col: 9, offset: 1459},
	name: "StdOutG",
},
&ruleRefExpr{
	pos: position{line: 67, col: 17, offset: 1467},
	name: "SinkInto",
},
	},
},
},
{
	name: "FileJobG",
	pos: position{line: 71, col: 1, offset: 1496},
	expr: &actionExpr{
	pos: position{line: 71, col: 13, offset: 1508},
	run: (*parser).callonFileJobG1,
	expr: &seqExpr{
	pos: position{line: 71, col: 13, offset: 1508},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 71, col: 13, offset: 1508},
	name: "_",
},
&litMatcher{
	pos: position{line: 71, col: 15, offset: 1510},
	val: "file(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 71, col: 23, offset: 1518},
	name: "_",
},
&labeledExpr{
	pos: position{line: 71, col: 25, offset: 1520},
	label: "filename",
	expr: &ruleRefExpr{
	pos: position{line: 71, col: 34, offset: 1529},
	name: "FileName",
},
},
&ruleRefExpr{
	pos: position{line: 71, col: 43, offset: 1538},
	name: "_",
},
&litMatcher{
	pos: position{line: 71, col: 45, offset: 1540},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchJobG",
	pos: position{line: 78, col: 1, offset: 1722},
	expr: &litMatcher{
	pos: position{line: 78, col: 15, offset: 1736},
	val: "branch",
	ignoreCase: false,
},
},
{
	name: "AggJobG",
	pos: position{line: 86, col: 1, offset: 1834},
	expr: &actionExpr{
	pos: position{line: 86, col: 12, offset: 1845},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 86, col: 12, offset: 1845},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 86, col: 12, offset: 1845},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 86, col: 18, offset: 1851},
	name: "_",
},
&labeledExpr{
	pos: position{line: 86, col: 20, offset: 1853},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 86, col: 27, offset: 1860},
	expr: &ruleRefExpr{
	pos: position{line: 86, col: 27, offset: 1860},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 86, col: 39, offset: 1872},
	name: "_",
},
&labeledExpr{
	pos: position{line: 86, col: 41, offset: 1874},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 86, col: 45, offset: 1878},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 86, col: 53, offset: 1886},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 86, col: 58, offset: 1891},
	expr: &seqExpr{
	pos: position{line: 86, col: 59, offset: 1892},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 86, col: 59, offset: 1892},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 86, col: 63, offset: 1896},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 86, col: 65, offset: 1898},
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
	pos: position{line: 100, col: 1, offset: 2238},
	expr: &ruleRefExpr{
	pos: position{line: 100, col: 12, offset: 2249},
	name: "CountAgg",
},
},
{
	name: "AggGroupBy",
	pos: position{line: 103, col: 1, offset: 2263},
	expr: &actionExpr{
	pos: position{line: 103, col: 15, offset: 2277},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 103, col: 15, offset: 2277},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 103, col: 15, offset: 2277},
	name: "_",
},
&litMatcher{
	pos: position{line: 103, col: 17, offset: 2279},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 103, col: 22, offset: 2284},
	name: "_",
},
&labeledExpr{
	pos: position{line: 103, col: 24, offset: 2286},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 103, col: 30, offset: 2292},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 103, col: 39, offset: 2301},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 103, col: 44, offset: 2306},
	expr: &seqExpr{
	pos: position{line: 103, col: 45, offset: 2307},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 103, col: 45, offset: 2307},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 103, col: 49, offset: 2311},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 103, col: 51, offset: 2313},
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
	name: "CountAgg",
	pos: position{line: 116, col: 1, offset: 2631},
	expr: &choiceExpr{
	pos: position{line: 116, col: 13, offset: 2643},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 116, col: 13, offset: 2643},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 116, col: 13, offset: 2643},
	label: "condition",
	expr: &litMatcher{
	pos: position{line: 116, col: 23, offset: 2653},
	val: "count_of_",
	ignoreCase: false,
},
},
&labeledExpr{
	pos: position{line: 116, col: 34, offset: 2664},
	label: "entity",
	expr: &ruleRefExpr{
	pos: position{line: 116, col: 41, offset: 2671},
	name: "Identifier",
},
},
&litMatcher{
	pos: position{line: 116, col: 51, offset: 2681},
	val: "_",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 116, col: 54, offset: 2684},
	label: "comparator",
	expr: &ruleRefExpr{
	pos: position{line: 116, col: 65, offset: 2695},
	name: "Comparator",
},
},
&litMatcher{
	pos: position{line: 116, col: 75, offset: 2705},
	val: "_",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 116, col: 78, offset: 2708},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 116, col: 84, offset: 2714},
	name: "Integer",
},
},
	},
},
&actionExpr{
	pos: position{line: 116, col: 92, offset: 2722},
	run: (*parser).callonCountAgg13,
	expr: &seqExpr{
	pos: position{line: 116, col: 92, offset: 2722},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 116, col: 92, offset: 2722},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 116, col: 98, offset: 2728},
	name: "_",
},
&litMatcher{
	pos: position{line: 116, col: 99, offset: 2729},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 116, col: 102, offset: 2732},
	name: "_",
},
&litMatcher{
	pos: position{line: 116, col: 103, offset: 2733},
	val: "count()",
	ignoreCase: false,
},
	},
},
},
	},
},
},
{
	name: "DoJobG",
	pos: position{line: 122, col: 1, offset: 2807},
	expr: &actionExpr{
	pos: position{line: 122, col: 11, offset: 2817},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 122, col: 11, offset: 2817},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 122, col: 16, offset: 2822},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 122, col: 16, offset: 2822},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 122, col: 27, offset: 2833},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 122, col: 38, offset: 2844},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 122, col: 49, offset: 2855},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 122, col: 58, offset: 2864},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 122, col: 67, offset: 2873},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchJobG",
	pos: position{line: 127, col: 1, offset: 2955},
	expr: &actionExpr{
	pos: position{line: 127, col: 15, offset: 2969},
	run: (*parser).callonBranchJobG1,
	expr: &seqExpr{
	pos: position{line: 127, col: 15, offset: 2969},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 127, col: 15, offset: 2969},
	val: "branch",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 127, col: 23, offset: 2977},
	name: "_",
},
&litMatcher{
	pos: position{line: 127, col: 24, offset: 2978},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 127, col: 27, offset: 2981},
	label: "_var",
	expr: &ruleRefExpr{
	pos: position{line: 127, col: 32, offset: 2986},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 127, col: 40, offset: 2994},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 132, col: 1, offset: 3053},
	expr: &actionExpr{
	pos: position{line: 132, col: 13, offset: 3065},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 132, col: 13, offset: 3065},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 132, col: 13, offset: 3065},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 132, col: 23, offset: 3075},
	name: "_",
},
&labeledExpr{
	pos: position{line: 132, col: 25, offset: 3077},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 132, col: 28, offset: 3080},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 132, col: 42, offset: 3094},
	name: "_",
},
&litMatcher{
	pos: position{line: 132, col: 44, offset: 3096},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 139, col: 1, offset: 3291},
	expr: &actionExpr{
	pos: position{line: 139, col: 18, offset: 3308},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 139, col: 18, offset: 3308},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 139, col: 18, offset: 3308},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 139, col: 23, offset: 3313},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 139, col: 28, offset: 3318},
	name: "_",
},
&labeledExpr{
	pos: position{line: 139, col: 30, offset: 3320},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 139, col: 33, offset: 3323},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 139, col: 45, offset: 3335},
	name: "_",
},
&labeledExpr{
	pos: position{line: 139, col: 47, offset: 3337},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 139, col: 54, offset: 3344},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 143, col: 1, offset: 3395},
	expr: &ruleRefExpr{
	pos: position{line: 143, col: 16, offset: 3410},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 146, col: 1, offset: 3444},
	expr: &actionExpr{
	pos: position{line: 146, col: 13, offset: 3456},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 146, col: 13, offset: 3456},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 146, col: 13, offset: 3456},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 146, col: 23, offset: 3466},
	name: "_",
},
&labeledExpr{
	pos: position{line: 146, col: 25, offset: 3468},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 146, col: 31, offset: 3474},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 146, col: 40, offset: 3483},
	name: "_",
},
&labeledExpr{
	pos: position{line: 146, col: 42, offset: 3485},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 146, col: 47, offset: 3490},
	expr: &seqExpr{
	pos: position{line: 146, col: 48, offset: 3491},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 146, col: 48, offset: 3491},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 146, col: 52, offset: 3495},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 146, col: 54, offset: 3497},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 146, col: 65, offset: 3508},
	name: "_",
},
&litMatcher{
	pos: position{line: 146, col: 67, offset: 3510},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 161, col: 1, offset: 3847},
	expr: &actionExpr{
	pos: position{line: 161, col: 11, offset: 3857},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 161, col: 11, offset: 3857},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 161, col: 11, offset: 3857},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 161, col: 18, offset: 3864},
	name: "_",
},
&labeledExpr{
	pos: position{line: 161, col: 20, offset: 3866},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 161, col: 25, offset: 3871},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 161, col: 34, offset: 3880},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 161, col: 38, offset: 3884},
	name: "_",
},
&labeledExpr{
	pos: position{line: 161, col: 40, offset: 3886},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 161, col: 44, offset: 3890},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 161, col: 51, offset: 3897},
	name: "_",
},
&litMatcher{
	pos: position{line: 161, col: 53, offset: 3899},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 173, col: 1, offset: 4121},
	expr: &actionExpr{
	pos: position{line: 173, col: 11, offset: 4131},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 173, col: 11, offset: 4131},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 173, col: 11, offset: 4131},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 173, col: 20, offset: 4140},
	name: "_",
},
&litMatcher{
	pos: position{line: 173, col: 22, offset: 4142},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 173, col: 27, offset: 4147},
	name: "_",
},
&labeledExpr{
	pos: position{line: 173, col: 29, offset: 4149},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 173, col: 35, offset: 4155},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 179, col: 1, offset: 4261},
	expr: &actionExpr{
	pos: position{line: 179, col: 12, offset: 4272},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 179, col: 12, offset: 4272},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 179, col: 12, offset: 4272},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 179, col: 21, offset: 4281},
	name: "_",
},
&labeledExpr{
	pos: position{line: 179, col: 23, offset: 4283},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 179, col: 27, offset: 4287},
	expr: &ruleRefExpr{
	pos: position{line: 179, col: 27, offset: 4287},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 179, col: 35, offset: 4295},
	name: "_",
},
&litMatcher{
	pos: position{line: 179, col: 37, offset: 4297},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 189, col: 1, offset: 4442},
	expr: &actionExpr{
	pos: position{line: 189, col: 13, offset: 4454},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 189, col: 13, offset: 4454},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 189, col: 13, offset: 4454},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 189, col: 23, offset: 4464},
	name: "_",
},
&labeledExpr{
	pos: position{line: 189, col: 25, offset: 4466},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 189, col: 30, offset: 4471},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 189, col: 42, offset: 4483},
	name: "_",
},
&litMatcher{
	pos: position{line: 189, col: 44, offset: 4485},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 199, col: 1, offset: 4773},
	expr: &actionExpr{
	pos: position{line: 199, col: 16, offset: 4788},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 199, col: 16, offset: 4788},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 199, col: 16, offset: 4788},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 199, col: 22, offset: 4794},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 199, col: 35, offset: 4807},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 199, col: 40, offset: 4812},
	expr: &seqExpr{
	pos: position{line: 199, col: 41, offset: 4813},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 199, col: 41, offset: 4813},
	name: "_",
},
&labeledExpr{
	pos: position{line: 199, col: 43, offset: 4815},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 199, col: 46, offset: 4818},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 199, col: 58, offset: 4830},
	name: "_",
},
&labeledExpr{
	pos: position{line: 199, col: 60, offset: 4832},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 199, col: 63, offset: 4835},
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
	pos: position{line: 204, col: 1, offset: 4993},
	expr: &choiceExpr{
	pos: position{line: 204, col: 17, offset: 5009},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 204, col: 17, offset: 5009},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 204, col: 17, offset: 5009},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 204, col: 17, offset: 5009},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 204, col: 21, offset: 5013},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 204, col: 24, offset: 5016},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 204, col: 36, offset: 5028},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 206, col: 5, offset: 5059},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 206, col: 5, offset: 5059},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 206, col: 8, offset: 5062},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 210, col: 1, offset: 5104},
	expr: &choiceExpr{
	pos: position{line: 210, col: 16, offset: 5119},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 210, col: 16, offset: 5119},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 210, col: 24, offset: 5127},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 210, col: 24, offset: 5127},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 217, col: 1, offset: 5306},
	expr: &actionExpr{
	pos: position{line: 217, col: 18, offset: 5323},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 217, col: 18, offset: 5323},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 217, col: 18, offset: 5323},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 217, col: 23, offset: 5328},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 217, col: 28, offset: 5333},
	name: "_",
},
&labeledExpr{
	pos: position{line: 217, col: 30, offset: 5335},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 217, col: 33, offset: 5338},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 217, col: 45, offset: 5350},
	name: "_",
},
&labeledExpr{
	pos: position{line: 217, col: 47, offset: 5352},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 217, col: 53, offset: 5358},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 217, col: 59, offset: 5364},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 221, col: 1, offset: 5438},
	expr: &choiceExpr{
	pos: position{line: 221, col: 16, offset: 5453},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 221, col: 16, offset: 5453},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 221, col: 23, offset: 5460},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 221, col: 30, offset: 5467},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 221, col: 36, offset: 5473},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 221, col: 42, offset: 5479},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 221, col: 49, offset: 5486},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 221, col: 49, offset: 5486},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 230, col: 1, offset: 5596},
	expr: &actionExpr{
	pos: position{line: 230, col: 12, offset: 5607},
	run: (*parser).callonStdOutG1,
	expr: &seqExpr{
	pos: position{line: 230, col: 12, offset: 5607},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 230, col: 12, offset: 5607},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 230, col: 22, offset: 5617},
	name: "_",
},
&labeledExpr{
	pos: position{line: 230, col: 24, offset: 5619},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 230, col: 31, offset: 5626},
	expr: &ruleRefExpr{
	pos: position{line: 230, col: 32, offset: 5627},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 230, col: 48, offset: 5643},
	name: "_",
},
&litMatcher{
	pos: position{line: 230, col: 50, offset: 5645},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SinkInto",
	pos: position{line: 244, col: 1, offset: 5891},
	expr: &actionExpr{
	pos: position{line: 244, col: 12, offset: 5902},
	run: (*parser).callonSinkInto1,
	expr: &seqExpr{
	pos: position{line: 244, col: 12, offset: 5902},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 244, col: 12, offset: 5902},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 244, col: 18, offset: 5908},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 244, col: 27, offset: 5917},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 244, col: 33, offset: 5923},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 244, col: 33, offset: 5923},
	name: "_",
},
&litMatcher{
	pos: position{line: 244, col: 35, offset: 5925},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 244, col: 39, offset: 5929},
	name: "_",
},
&litMatcher{
	pos: position{line: 244, col: 41, offset: 5931},
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
	pos: position{line: 255, col: 1, offset: 6092},
	expr: &actionExpr{
	pos: position{line: 255, col: 13, offset: 6104},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 255, col: 13, offset: 6104},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 255, col: 13, offset: 6104},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 255, col: 22, offset: 6113},
	expr: &seqExpr{
	pos: position{line: 255, col: 23, offset: 6114},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 255, col: 23, offset: 6114},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 255, col: 27, offset: 6118},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 255, col: 31, offset: 6122},
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
	pos: position{line: 261, col: 1, offset: 6223},
	expr: &actionExpr{
	pos: position{line: 261, col: 10, offset: 6232},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 261, col: 10, offset: 6232},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 261, col: 10, offset: 6232},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 261, col: 16, offset: 6238},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 261, col: 23, offset: 6245},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 261, col: 28, offset: 6250},
	expr: &seqExpr{
	pos: position{line: 261, col: 29, offset: 6251},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 261, col: 29, offset: 6251},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 261, col: 31, offset: 6253},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 261, col: 36, offset: 6258},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 261, col: 38, offset: 6260},
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
	pos: position{line: 265, col: 1, offset: 6325},
	expr: &actionExpr{
	pos: position{line: 265, col: 11, offset: 6335},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 265, col: 11, offset: 6335},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 265, col: 16, offset: 6340},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 265, col: 16, offset: 6340},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 265, col: 24, offset: 6348},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 265, col: 33, offset: 6357},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 265, col: 48, offset: 6372},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 265, col: 59, offset: 6383},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 269, col: 1, offset: 6425},
	expr: &choiceExpr{
	pos: position{line: 269, col: 9, offset: 6433},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 269, col: 9, offset: 6433},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 269, col: 15, offset: 6439},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 269, col: 21, offset: 6445},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 269, col: 27, offset: 6451},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 269, col: 33, offset: 6457},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 269, col: 33, offset: 6457},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 273, col: 1, offset: 6505},
	expr: &choiceExpr{
	pos: position{line: 273, col: 12, offset: 6516},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 273, col: 12, offset: 6516},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 273, col: 12, offset: 6516},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 273, col: 12, offset: 6516},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 273, col: 16, offset: 6520},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 273, col: 19, offset: 6523},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 273, col: 25, offset: 6529},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 277, col: 5, offset: 6605},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 277, col: 5, offset: 6605},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 277, col: 8, offset: 6608},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 284, col: 1, offset: 6685},
	expr: &actionExpr{
	pos: position{line: 284, col: 9, offset: 6693},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 284, col: 9, offset: 6693},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 284, col: 9, offset: 6693},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 284, col: 15, offset: 6699},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 284, col: 21, offset: 6705},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 284, col: 26, offset: 6710},
	expr: &seqExpr{
	pos: position{line: 284, col: 27, offset: 6711},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 284, col: 27, offset: 6711},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 284, col: 29, offset: 6713},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 284, col: 34, offset: 6718},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 284, col: 36, offset: 6720},
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
	pos: position{line: 288, col: 1, offset: 6784},
	expr: &actionExpr{
	pos: position{line: 288, col: 10, offset: 6793},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 288, col: 10, offset: 6793},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 288, col: 15, offset: 6798},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 288, col: 15, offset: 6798},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 288, col: 23, offset: 6806},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 288, col: 32, offset: 6815},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 288, col: 47, offset: 6830},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 292, col: 1, offset: 6873},
	expr: &actionExpr{
	pos: position{line: 292, col: 13, offset: 6885},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 292, col: 13, offset: 6885},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 292, col: 13, offset: 6885},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 292, col: 20, offset: 6892},
	expr: &choiceExpr{
	pos: position{line: 292, col: 21, offset: 6893},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 292, col: 21, offset: 6893},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 292, col: 31, offset: 6903},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 292, col: 39, offset: 6911},
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
	pos: position{line: 296, col: 1, offset: 6961},
	expr: &actionExpr{
	pos: position{line: 296, col: 17, offset: 6977},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 296, col: 17, offset: 6977},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 296, col: 17, offset: 6977},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 296, col: 21, offset: 6981},
	expr: &choiceExpr{
	pos: position{line: 296, col: 23, offset: 6983},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 296, col: 23, offset: 6983},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 296, col: 23, offset: 6983},
	expr: &ruleRefExpr{
	pos: position{line: 296, col: 24, offset: 6984},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 296, col: 36, offset: 6996,
},
	},
},
&seqExpr{
	pos: position{line: 296, col: 40, offset: 7000},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 296, col: 40, offset: 7000},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 296, col: 45, offset: 7005},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 296, col: 63, offset: 7023},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "EscapedChar",
	pos: position{line: 303, col: 1, offset: 7204},
	expr: &charClassMatcher{
	pos: position{line: 303, col: 16, offset: 7219},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 305, col: 1, offset: 7237},
	expr: &ruleRefExpr{
	pos: position{line: 305, col: 19, offset: 7255},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 307, col: 1, offset: 7275},
	expr: &charClassMatcher{
	pos: position{line: 307, col: 21, offset: 7295},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 309, col: 1, offset: 7310},
	expr: &actionExpr{
	pos: position{line: 309, col: 9, offset: 7318},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 309, col: 9, offset: 7318},
	expr: &choiceExpr{
	pos: position{line: 309, col: 10, offset: 7319},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 309, col: 10, offset: 7319},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 309, col: 19, offset: 7328},
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
	pos: position{line: 313, col: 1, offset: 7378},
	expr: &choiceExpr{
	pos: position{line: 313, col: 11, offset: 7388},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 313, col: 11, offset: 7388},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 313, col: 19, offset: 7396},
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
	pos: position{line: 315, col: 1, offset: 7405},
	expr: &actionExpr{
	pos: position{line: 315, col: 11, offset: 7415},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 315, col: 11, offset: 7415},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 315, col: 11, offset: 7415},
	expr: &litMatcher{
	pos: position{line: 315, col: 11, offset: 7415},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 315, col: 16, offset: 7420},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 319, col: 1, offset: 7493},
	expr: &actionExpr{
	pos: position{line: 319, col: 19, offset: 7511},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 319, col: 19, offset: 7511},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 323, col: 1, offset: 7584},
	expr: &actionExpr{
	pos: position{line: 323, col: 10, offset: 7593},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 323, col: 10, offset: 7593},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 323, col: 10, offset: 7593},
	expr: &litMatcher{
	pos: position{line: 323, col: 10, offset: 7593},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 323, col: 15, offset: 7598},
	expr: &ruleRefExpr{
	pos: position{line: 323, col: 15, offset: 7598},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 323, col: 24, offset: 7607},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 323, col: 28, offset: 7611},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 327, col: 1, offset: 7678},
	expr: &choiceExpr{
	pos: position{line: 327, col: 12, offset: 7689},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 327, col: 12, offset: 7689},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 327, col: 18, offset: 7695},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 327, col: 18, offset: 7695},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 327, col: 38, offset: 7715},
	expr: &ruleRefExpr{
	pos: position{line: 327, col: 38, offset: 7715},
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
	pos: position{line: 329, col: 1, offset: 7732},
	expr: &charClassMatcher{
	pos: position{line: 329, col: 17, offset: 7748},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 331, col: 1, offset: 7757},
	expr: &charClassMatcher{
	pos: position{line: 331, col: 24, offset: 7780},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 333, col: 1, offset: 7789},
	expr: &zeroOrMoreExpr{
	pos: position{line: 333, col: 19, offset: 7807},
	expr: &charClassMatcher{
	pos: position{line: 333, col: 19, offset: 7807},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 335, col: 1, offset: 7821},
	expr: &charClassMatcher{
	pos: position{line: 335, col: 18, offset: 7838},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EOF",
	pos: position{line: 337, col: 1, offset: 7851},
	expr: &notExpr{
	pos: position{line: 337, col: 8, offset: 7858},
	expr: &anyMatcher{
	line: 337, col: 9, offset: 7859,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 339, col: 1, offset: 7864},
	expr: &litMatcher{
	pos: position{line: 339, col: 17, offset: 7880},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 345, col: 1, offset: 7967},
	expr: &actionExpr{
	pos: position{line: 345, col: 18, offset: 7984},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 345, col: 18, offset: 7984},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 345, col: 18, offset: 7984},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 345, col: 24, offset: 7990},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 345, col: 33, offset: 7999},
	name: "_",
},
&labeledExpr{
	pos: position{line: 345, col: 35, offset: 8001},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 345, col: 40, offset: 8006},
	expr: &seqExpr{
	pos: position{line: 345, col: 41, offset: 8007},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 345, col: 41, offset: 8007},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 345, col: 45, offset: 8011},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 345, col: 47, offset: 8013},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 345, col: 56, offset: 8022},
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
}
func (c *current) onCommand1(command, rest interface{}) (interface{}, error) {

    commands := []interface{}{(command.([]interface{}))}
    	for _,val := range(cast.ToIfaceSlice(rest)){
    		_,ok := val.([]interface{})
    		if ok{
    			newVal := val.([]interface{})
    			if len(newVal)>0{
    				newCmds := newVal[0]
    				nextCmd := cast.ToIfaceSlice(newCmds)
    				for _,cmd := range(nextCmd){
    					if len(cast.ToIfaceSlice(cmd))>0{
    						commands = append(commands,cast.ToIfaceSlice(cmd))
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
    				for _,item := range(newVal){
    					_,ok:= item.([]interface{})
    					if ok {
							castedItem := cast.ToIfaceSlice(item)
							if len(castedItem) == 4 {
								stages = append(stages, castedItem[3])
							}
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

func (c *current) onCountAgg13() (interface{}, error) {

    return "count",nil
}

func (p *parser) callonCountAgg13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCountAgg13()
}

func (c *current) onDoJobG1(job interface{}) (interface{}, error) {

        return DoNodeJob{Function: job}, nil
}

func (p *parser) callonDoJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDoJobG1(stack["job"])
}

func (c *current) onBranchJobG1(_var interface{}) (interface{}, error) {
		newVar,_:= cast.TryString(_var)
		fmt.Println("herer",newVar)
        return Branch{Condition:newVar},nil
}

func (p *parser) callonBranchJobG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBranchJobG1(stack["_var"])
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

func (c *current) onVarParamListG1(first, rest interface{}) (interface{}, error) {

    return getParamList(first, rest), nil
}

func (p *parser) callonVarParamListG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVarParamListG1(stack["first"], stack["rest"])
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

