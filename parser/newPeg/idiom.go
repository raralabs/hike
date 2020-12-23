
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
	pos: position{line: 6, col: 1, offset: 124},
	expr: &actionExpr{
	pos: position{line: 6, col: 12, offset: 135},
	run: (*parser).callonCommand1,
	expr: &seqExpr{
	pos: position{line: 6, col: 12, offset: 135},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 6, col: 12, offset: 135},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 6, col: 18, offset: 141},
	name: "Statement",
},
},
&labeledExpr{
	pos: position{line: 6, col: 28, offset: 151},
	label: "second",
	expr: &zeroOrMoreExpr{
	pos: position{line: 6, col: 35, offset: 158},
	expr: &seqExpr{
	pos: position{line: 6, col: 36, offset: 159},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 6, col: 36, offset: 159},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 6, col: 38, offset: 161},
	name: "Statement",
},
	},
},
},
},
&labeledExpr{
	pos: position{line: 6, col: 51, offset: 174},
	label: "last",
	expr: &seqExpr{
	pos: position{line: 6, col: 58, offset: 181},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 6, col: 58, offset: 181},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 6, col: 60, offset: 183},
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
	pos: position{line: 19, col: 1, offset: 592},
	expr: &actionExpr{
	pos: position{line: 19, col: 14, offset: 605},
	run: (*parser).callonStatement1,
	expr: &seqExpr{
	pos: position{line: 19, col: 14, offset: 605},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 19, col: 14, offset: 605},
	label: "source",
	expr: &ruleRefExpr{
	pos: position{line: 19, col: 21, offset: 612},
	name: "Source",
},
},
&labeledExpr{
	pos: position{line: 19, col: 28, offset: 619},
	label: "transform",
	expr: &zeroOrMoreExpr{
	pos: position{line: 19, col: 38, offset: 629},
	expr: &seqExpr{
	pos: position{line: 19, col: 39, offset: 630},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 19, col: 39, offset: 630},
	name: "_",
},
&litMatcher{
	pos: position{line: 19, col: 41, offset: 632},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 19, col: 45, offset: 636},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 19, col: 47, offset: 638},
	name: "Transform",
},
	},
},
},
},
&labeledExpr{
	pos: position{line: 19, col: 59, offset: 650},
	label: "sink",
	expr: &seqExpr{
	pos: position{line: 19, col: 65, offset: 656},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 19, col: 65, offset: 656},
	name: "_",
},
&litMatcher{
	pos: position{line: 19, col: 67, offset: 658},
	val: "|",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 19, col: 71, offset: 662},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 19, col: 73, offset: 664},
	name: "Sink",
},
	},
},
},
&labeledExpr{
	pos: position{line: 19, col: 79, offset: 670},
	label: "last",
	expr: &seqExpr{
	pos: position{line: 19, col: 86, offset: 677},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 19, col: 86, offset: 677},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 19, col: 88, offset: 679},
	name: "TOK_SEMICOLON",
},
&ruleRefExpr{
	pos: position{line: 19, col: 102, offset: 693},
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
	pos: position{line: 32, col: 1, offset: 1105},
	expr: &choiceExpr{
	pos: position{line: 32, col: 11, offset: 1115},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 32, col: 11, offset: 1115},
	name: "BranchJobG",
},
&ruleRefExpr{
	pos: position{line: 32, col: 22, offset: 1126},
	name: "FileJobG",
},
&ruleRefExpr{
	pos: position{line: 32, col: 31, offset: 1135},
	name: "FakeJobG",
},
&ruleRefExpr{
	pos: position{line: 32, col: 40, offset: 1144},
	name: "LiftBridgeReaderG",
},
&ruleRefExpr{
	pos: position{line: 32, col: 58, offset: 1162},
	name: "SecSource",
},
	},
},
},
{
	name: "Transform",
	pos: position{line: 34, col: 1, offset: 1175},
	expr: &choiceExpr{
	pos: position{line: 34, col: 14, offset: 1188},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 34, col: 14, offset: 1188},
	name: "AggJobG",
},
&ruleRefExpr{
	pos: position{line: 34, col: 22, offset: 1196},
	name: "DoJobG",
},
&ruleRefExpr{
	pos: position{line: 34, col: 29, offset: 1203},
	name: "JoinJobG",
},
&ruleRefExpr{
	pos: position{line: 34, col: 38, offset: 1212},
	name: "MapJobG",
},
	},
},
},
{
	name: "Sink",
	pos: position{line: 36, col: 1, offset: 1223},
	expr: &choiceExpr{
	pos: position{line: 36, col: 9, offset: 1231},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 36, col: 9, offset: 1231},
	name: "StdOutG",
},
&ruleRefExpr{
	pos: position{line: 36, col: 17, offset: 1239},
	name: "SinkInto",
},
&ruleRefExpr{
	pos: position{line: 36, col: 26, offset: 1248},
	name: "PlotG",
},
	},
},
},
{
	name: "FileJobG",
	pos: position{line: 40, col: 1, offset: 1274},
	expr: &choiceExpr{
	pos: position{line: 40, col: 13, offset: 1286},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 40, col: 13, offset: 1286},
	run: (*parser).callonFileJobG2,
	expr: &seqExpr{
	pos: position{line: 40, col: 13, offset: 1286},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 40, col: 13, offset: 1286},
	name: "_",
},
&litMatcher{
	pos: position{line: 40, col: 15, offset: 1288},
	val: "file(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 40, col: 23, offset: 1296},
	name: "_",
},
&labeledExpr{
	pos: position{line: 40, col: 25, offset: 1298},
	label: "filename",
	expr: &ruleRefExpr{
	pos: position{line: 40, col: 34, offset: 1307},
	name: "FileName",
},
},
&ruleRefExpr{
	pos: position{line: 40, col: 43, offset: 1316},
	name: "_",
},
&litMatcher{
	pos: position{line: 40, col: 45, offset: 1318},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 45, col: 4, offset: 1531},
	run: (*parser).callonFileJobG11,
	expr: &seqExpr{
	pos: position{line: 45, col: 4, offset: 1531},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 45, col: 4, offset: 1531},
	name: "_",
},
&litMatcher{
	pos: position{line: 45, col: 6, offset: 1533},
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
	pos: position{line: 49, col: 1, offset: 1597},
	expr: &choiceExpr{
	pos: position{line: 49, col: 13, offset: 1609},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 49, col: 13, offset: 1609},
	run: (*parser).callonFakeJobG2,
	expr: &seqExpr{
	pos: position{line: 49, col: 13, offset: 1609},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 49, col: 13, offset: 1609},
	name: "_",
},
&litMatcher{
	pos: position{line: 49, col: 15, offset: 1611},
	val: "fake",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 49, col: 21, offset: 1617},
	name: "_",
},
&litMatcher{
	pos: position{line: 49, col: 23, offset: 1619},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 49, col: 27, offset: 1623},
	name: "_",
},
&labeledExpr{
	pos: position{line: 49, col: 29, offset: 1625},
	label: "numData",
	expr: &ruleRefExpr{
	pos: position{line: 49, col: 37, offset: 1633},
	name: "PositiveNumber",
},
},
&ruleRefExpr{
	pos: position{line: 49, col: 52, offset: 1648},
	name: "_",
},
&litMatcher{
	pos: position{line: 49, col: 54, offset: 1650},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 54, col: 3, offset: 1836},
	run: (*parser).callonFakeJobG13,
	expr: &seqExpr{
	pos: position{line: 54, col: 3, offset: 1836},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 54, col: 3, offset: 1836},
	name: "_",
},
&litMatcher{
	pos: position{line: 54, col: 5, offset: 1838},
	val: "fake",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 54, col: 11, offset: 1844},
	name: "_",
},
&litMatcher{
	pos: position{line: 54, col: 13, offset: 1846},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 54, col: 17, offset: 1850},
	name: "_",
},
&litMatcher{
	pos: position{line: 54, col: 19, offset: 1852},
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
	pos: position{line: 58, col: 1, offset: 1936},
	expr: &actionExpr{
	pos: position{line: 58, col: 15, offset: 1950},
	run: (*parser).callonBranchJobG1,
	expr: &seqExpr{
	pos: position{line: 58, col: 15, offset: 1950},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 58, col: 15, offset: 1950},
	val: "branch",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 58, col: 23, offset: 1958},
	name: "_",
},
&litMatcher{
	pos: position{line: 58, col: 24, offset: 1959},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 58, col: 28, offset: 1963},
	name: "_",
},
&labeledExpr{
	pos: position{line: 58, col: 30, offset: 1965},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 58, col: 33, offset: 1968},
	name: "IDENTIFIER",
},
},
&litMatcher{
	pos: position{line: 58, col: 44, offset: 1979},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SecSource",
	pos: position{line: 65, col: 1, offset: 2195},
	expr: &choiceExpr{
	pos: position{line: 65, col: 14, offset: 2208},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 65, col: 14, offset: 2208},
	run: (*parser).callonSecSource2,
	expr: &seqExpr{
	pos: position{line: 65, col: 14, offset: 2208},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 65, col: 14, offset: 2208},
	label: "src1",
	expr: &ruleRefExpr{
	pos: position{line: 65, col: 19, offset: 2213},
	name: "IDENTIFIER",
},
},
&labeledExpr{
	pos: position{line: 65, col: 30, offset: 2224},
	label: "rest",
	expr: &zeroOrOneExpr{
	pos: position{line: 65, col: 35, offset: 2229},
	expr: &seqExpr{
	pos: position{line: 65, col: 36, offset: 2230},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 65, col: 36, offset: 2230},
	name: "_",
},
&litMatcher{
	pos: position{line: 65, col: 38, offset: 2232},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 65, col: 41, offset: 2235},
	name: "_",
},
&labeledExpr{
	pos: position{line: 65, col: 43, offset: 2237},
	label: "src2",
	expr: &ruleRefExpr{
	pos: position{line: 65, col: 48, offset: 2242},
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
	pos: position{line: 77, col: 3, offset: 2651},
	run: (*parser).callonSecSource14,
	expr: &seqExpr{
	pos: position{line: 77, col: 3, offset: 2651},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 77, col: 3, offset: 2651},
	label: "src1",
	expr: &ruleRefExpr{
	pos: position{line: 77, col: 8, offset: 2656},
	name: "IDENTIFIER",
},
},
&labeledExpr{
	pos: position{line: 77, col: 19, offset: 2667},
	label: "rest",
	expr: &ruleRefExpr{
	pos: position{line: 77, col: 25, offset: 2673},
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
	name: "LiftBridgeReaderG",
	pos: position{line: 81, col: 1, offset: 2765},
	expr: &actionExpr{
	pos: position{line: 81, col: 22, offset: 2786},
	run: (*parser).callonLiftBridgeReaderG1,
	expr: &seqExpr{
	pos: position{line: 81, col: 22, offset: 2786},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 81, col: 22, offset: 2786},
	val: "liftreader",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 81, col: 35, offset: 2799},
	name: "_",
},
&litMatcher{
	pos: position{line: 81, col: 37, offset: 2801},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 81, col: 41, offset: 2805},
	name: "_",
},
&labeledExpr{
	pos: position{line: 81, col: 43, offset: 2807},
	label: "params",
	expr: &ruleRefExpr{
	pos: position{line: 81, col: 50, offset: 2814},
	name: "LiftReaderParams",
},
},
&ruleRefExpr{
	pos: position{line: 81, col: 67, offset: 2831},
	name: "_",
},
&litMatcher{
	pos: position{line: 81, col: 69, offset: 2833},
	val: ")",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 81, col: 73, offset: 2837},
	name: "_",
},
	},
},
},
},
{
	name: "LiftReaderParams",
	pos: position{line: 86, col: 1, offset: 2942},
	expr: &actionExpr{
	pos: position{line: 86, col: 20, offset: 2961},
	run: (*parser).callonLiftReaderParams1,
	expr: &seqExpr{
	pos: position{line: 86, col: 20, offset: 2961},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 86, col: 20, offset: 2961},
	label: "ip",
	expr: &zeroOrOneExpr{
	pos: position{line: 86, col: 23, offset: 2964},
	expr: &ruleRefExpr{
	pos: position{line: 86, col: 23, offset: 2964},
	name: "Ip",
},
},
},
&ruleRefExpr{
	pos: position{line: 86, col: 27, offset: 2968},
	name: "_",
},
&litMatcher{
	pos: position{line: 86, col: 29, offset: 2970},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 86, col: 33, offset: 2974},
	name: "_",
},
&labeledExpr{
	pos: position{line: 86, col: 35, offset: 2976},
	label: "streamName",
	expr: &ruleRefExpr{
	pos: position{line: 86, col: 46, offset: 2987},
	name: "StreamName",
},
},
&ruleRefExpr{
	pos: position{line: 86, col: 57, offset: 2998},
	name: "_",
},
&litMatcher{
	pos: position{line: 86, col: 59, offset: 3000},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 86, col: 63, offset: 3004},
	name: "_",
},
&labeledExpr{
	pos: position{line: 86, col: 65, offset: 3006},
	label: "streamSubject",
	expr: &ruleRefExpr{
	pos: position{line: 86, col: 79, offset: 3020},
	name: "StreamSubject",
},
},
&ruleRefExpr{
	pos: position{line: 86, col: 93, offset: 3034},
	name: "_",
},
	},
},
},
},
{
	name: "JoinJobG",
	pos: position{line: 96, col: 1, offset: 3200},
	expr: &actionExpr{
	pos: position{line: 96, col: 13, offset: 3212},
	run: (*parser).callonJoinJobG1,
	expr: &seqExpr{
	pos: position{line: 96, col: 13, offset: 3212},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 96, col: 13, offset: 3212},
	val: "join",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 96, col: 20, offset: 3219},
	name: "_",
},
&labeledExpr{
	pos: position{line: 96, col: 22, offset: 3221},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 96, col: 26, offset: 3225},
	name: "JoinJobs",
},
},
	},
},
},
},
{
	name: "JoinJobs",
	pos: position{line: 100, col: 1, offset: 3298},
	expr: &choiceExpr{
	pos: position{line: 100, col: 12, offset: 3309},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 100, col: 12, offset: 3309},
	name: "InnerJoinG",
},
&ruleRefExpr{
	pos: position{line: 100, col: 23, offset: 3320},
	name: "OuterJoinG",
},
	},
},
},
{
	name: "InnerJoinG",
	pos: position{line: 102, col: 1, offset: 3334},
	expr: &choiceExpr{
	pos: position{line: 102, col: 14, offset: 3347},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 102, col: 14, offset: 3347},
	run: (*parser).callonInnerJoinG2,
	expr: &seqExpr{
	pos: position{line: 102, col: 14, offset: 3347},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 102, col: 14, offset: 3347},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 102, col: 20, offset: 3353},
	expr: &ruleRefExpr{
	pos: position{line: 102, col: 20, offset: 3353},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 102, col: 30, offset: 3363},
	name: "_",
},
&litMatcher{
	pos: position{line: 102, col: 32, offset: 3365},
	val: "inner",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 102, col: 40, offset: 3373},
	name: "_",
},
&litMatcher{
	pos: position{line: 102, col: 42, offset: 3375},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 102, col: 46, offset: 3379},
	name: "_",
},
&labeledExpr{
	pos: position{line: 102, col: 48, offset: 3381},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 102, col: 54, offset: 3387},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 102, col: 60, offset: 3393},
	name: "_",
},
&litMatcher{
	pos: position{line: 102, col: 62, offset: 3395},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 108, col: 4, offset: 3588},
	run: (*parser).callonInnerJoinG16,
	expr: &seqExpr{
	pos: position{line: 108, col: 4, offset: 3588},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 108, col: 4, offset: 3588},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 108, col: 10, offset: 3594},
	expr: &ruleRefExpr{
	pos: position{line: 108, col: 10, offset: 3594},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 108, col: 20, offset: 3604},
	name: "_",
},
&litMatcher{
	pos: position{line: 108, col: 22, offset: 3606},
	val: "inner",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 108, col: 30, offset: 3614},
	name: "_",
},
&litMatcher{
	pos: position{line: 108, col: 32, offset: 3616},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 108, col: 36, offset: 3620},
	name: "_",
},
&litMatcher{
	pos: position{line: 108, col: 38, offset: 3622},
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
	pos: position{line: 112, col: 1, offset: 3682},
	expr: &choiceExpr{
	pos: position{line: 112, col: 15, offset: 3696},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 112, col: 15, offset: 3696},
	name: "LeftOuterG",
},
&ruleRefExpr{
	pos: position{line: 112, col: 26, offset: 3707},
	name: "RightOuterG",
},
&ruleRefExpr{
	pos: position{line: 112, col: 38, offset: 3719},
	name: "FullOuterG",
},
	},
},
},
{
	name: "LeftOuterG",
	pos: position{line: 114, col: 1, offset: 3733},
	expr: &choiceExpr{
	pos: position{line: 114, col: 15, offset: 3747},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 114, col: 15, offset: 3747},
	run: (*parser).callonLeftOuterG2,
	expr: &seqExpr{
	pos: position{line: 114, col: 15, offset: 3747},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 114, col: 15, offset: 3747},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 114, col: 21, offset: 3753},
	expr: &ruleRefExpr{
	pos: position{line: 114, col: 21, offset: 3753},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 114, col: 31, offset: 3763},
	name: "_",
},
&litMatcher{
	pos: position{line: 114, col: 33, offset: 3765},
	val: "leftouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 114, col: 45, offset: 3777},
	name: "_",
},
&litMatcher{
	pos: position{line: 114, col: 47, offset: 3779},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 114, col: 51, offset: 3783},
	name: "_",
},
&labeledExpr{
	pos: position{line: 114, col: 53, offset: 3785},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 114, col: 59, offset: 3791},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 114, col: 65, offset: 3797},
	name: "_",
},
&litMatcher{
	pos: position{line: 114, col: 67, offset: 3799},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 120, col: 12, offset: 4070},
	run: (*parser).callonLeftOuterG16,
	expr: &seqExpr{
	pos: position{line: 120, col: 12, offset: 4070},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 120, col: 12, offset: 4070},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 120, col: 18, offset: 4076},
	expr: &ruleRefExpr{
	pos: position{line: 120, col: 18, offset: 4076},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 120, col: 28, offset: 4086},
	name: "_",
},
&litMatcher{
	pos: position{line: 120, col: 30, offset: 4088},
	val: "leftouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 120, col: 42, offset: 4100},
	name: "_",
},
&litMatcher{
	pos: position{line: 120, col: 44, offset: 4102},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 120, col: 48, offset: 4106},
	name: "_",
},
&litMatcher{
	pos: position{line: 120, col: 50, offset: 4108},
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
	pos: position{line: 124, col: 1, offset: 4187},
	expr: &choiceExpr{
	pos: position{line: 124, col: 15, offset: 4201},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 124, col: 15, offset: 4201},
	run: (*parser).callonRightOuterG2,
	expr: &seqExpr{
	pos: position{line: 124, col: 15, offset: 4201},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 124, col: 15, offset: 4201},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 124, col: 21, offset: 4207},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 124, col: 30, offset: 4216},
	name: "_",
},
&litMatcher{
	pos: position{line: 124, col: 32, offset: 4218},
	val: "rightouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 124, col: 45, offset: 4231},
	name: "_",
},
&litMatcher{
	pos: position{line: 124, col: 47, offset: 4233},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 124, col: 51, offset: 4237},
	name: "_",
},
&labeledExpr{
	pos: position{line: 124, col: 53, offset: 4239},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 124, col: 59, offset: 4245},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 124, col: 65, offset: 4251},
	name: "_",
},
&litMatcher{
	pos: position{line: 124, col: 67, offset: 4253},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 130, col: 17, offset: 4529},
	run: (*parser).callonRightOuterG15,
	expr: &seqExpr{
	pos: position{line: 130, col: 17, offset: 4529},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 130, col: 17, offset: 4529},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 130, col: 23, offset: 4535},
	expr: &ruleRefExpr{
	pos: position{line: 130, col: 23, offset: 4535},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 130, col: 33, offset: 4545},
	name: "_",
},
&litMatcher{
	pos: position{line: 130, col: 35, offset: 4547},
	val: "rightouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 130, col: 48, offset: 4560},
	name: "_",
},
&litMatcher{
	pos: position{line: 130, col: 50, offset: 4562},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 130, col: 54, offset: 4566},
	name: "_",
},
&litMatcher{
	pos: position{line: 130, col: 56, offset: 4568},
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
	pos: position{line: 133, col: 1, offset: 4675},
	expr: &choiceExpr{
	pos: position{line: 133, col: 15, offset: 4689},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 133, col: 15, offset: 4689},
	run: (*parser).callonFullOuterG2,
	expr: &seqExpr{
	pos: position{line: 133, col: 15, offset: 4689},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 133, col: 15, offset: 4689},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 133, col: 21, offset: 4695},
	name: "EqAliasG",
},
},
&ruleRefExpr{
	pos: position{line: 133, col: 30, offset: 4704},
	name: "_",
},
&litMatcher{
	pos: position{line: 133, col: 32, offset: 4706},
	val: "rightouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 133, col: 45, offset: 4719},
	name: "_",
},
&litMatcher{
	pos: position{line: 133, col: 47, offset: 4721},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 133, col: 51, offset: 4725},
	name: "_",
},
&labeledExpr{
	pos: position{line: 133, col: 53, offset: 4727},
	label: "query",
	expr: &ruleRefExpr{
	pos: position{line: 133, col: 59, offset: 4733},
	name: "Query",
},
},
&ruleRefExpr{
	pos: position{line: 133, col: 65, offset: 4739},
	name: "_",
},
&litMatcher{
	pos: position{line: 133, col: 67, offset: 4741},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 139, col: 17, offset: 5018},
	run: (*parser).callonFullOuterG15,
	expr: &seqExpr{
	pos: position{line: 139, col: 17, offset: 5018},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 139, col: 17, offset: 5018},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 139, col: 23, offset: 5024},
	expr: &ruleRefExpr{
	pos: position{line: 139, col: 23, offset: 5024},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 139, col: 33, offset: 5034},
	name: "_",
},
&litMatcher{
	pos: position{line: 139, col: 35, offset: 5036},
	val: "fullouter",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 139, col: 47, offset: 5048},
	name: "_",
},
&litMatcher{
	pos: position{line: 139, col: 49, offset: 5050},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 139, col: 53, offset: 5054},
	name: "_",
},
&litMatcher{
	pos: position{line: 139, col: 55, offset: 5056},
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
	pos: position{line: 144, col: 1, offset: 5150},
	expr: &choiceExpr{
	pos: position{line: 144, col: 10, offset: 5159},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 144, col: 10, offset: 5159},
	run: (*parser).callonQuery2,
	expr: &seqExpr{
	pos: position{line: 144, col: 10, offset: 5159},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 144, col: 10, offset: 5159},
	val: "select",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 144, col: 19, offset: 5168},
	name: "_",
},
&labeledExpr{
	pos: position{line: 144, col: 21, offset: 5170},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 144, col: 27, offset: 5176},
	name: "Fields",
},
},
&ruleRefExpr{
	pos: position{line: 144, col: 34, offset: 5183},
	name: "_",
},
&litMatcher{
	pos: position{line: 144, col: 36, offset: 5185},
	val: "on",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 144, col: 41, offset: 5190},
	name: "_",
},
&labeledExpr{
	pos: position{line: 144, col: 43, offset: 5192},
	label: "condition",
	expr: &ruleRefExpr{
	pos: position{line: 144, col: 53, offset: 5202},
	name: "JoinCondition",
},
},
&ruleRefExpr{
	pos: position{line: 144, col: 67, offset: 5216},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 146, col: 4, offset: 5326},
	run: (*parser).callonQuery14,
	expr: &seqExpr{
	pos: position{line: 146, col: 4, offset: 5326},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 146, col: 4, offset: 5326},
	val: "select",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 146, col: 13, offset: 5335},
	name: "_",
},
&litMatcher{
	pos: position{line: 146, col: 15, offset: 5337},
	val: "on",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 146, col: 20, offset: 5342},
	name: "_",
},
&labeledExpr{
	pos: position{line: 146, col: 22, offset: 5344},
	label: "condition",
	expr: &ruleRefExpr{
	pos: position{line: 146, col: 32, offset: 5354},
	name: "JoinCondition",
},
},
&ruleRefExpr{
	pos: position{line: 146, col: 46, offset: 5368},
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
	pos: position{line: 150, col: 1, offset: 5453},
	expr: &choiceExpr{
	pos: position{line: 150, col: 11, offset: 5463},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 150, col: 11, offset: 5463},
	name: "AllFields",
},
&ruleRefExpr{
	pos: position{line: 150, col: 21, offset: 5473},
	name: "SelFields",
},
	},
},
},
{
	name: "AllFields",
	pos: position{line: 152, col: 1, offset: 5486},
	expr: &actionExpr{
	pos: position{line: 152, col: 14, offset: 5499},
	run: (*parser).callonAllFields1,
	expr: &litMatcher{
	pos: position{line: 152, col: 14, offset: 5499},
	val: "*",
	ignoreCase: false,
},
},
},
{
	name: "SelFields",
	pos: position{line: 158, col: 1, offset: 5591},
	expr: &actionExpr{
	pos: position{line: 158, col: 14, offset: 5604},
	run: (*parser).callonSelFields1,
	expr: &seqExpr{
	pos: position{line: 158, col: 14, offset: 5604},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 158, col: 14, offset: 5604},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 158, col: 20, offset: 5610},
	name: "Word",
},
},
&ruleRefExpr{
	pos: position{line: 158, col: 25, offset: 5615},
	name: "_",
},
&labeledExpr{
	pos: position{line: 158, col: 28, offset: 5618},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 158, col: 33, offset: 5623},
	expr: &seqExpr{
	pos: position{line: 158, col: 34, offset: 5624},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 158, col: 34, offset: 5624},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 158, col: 38, offset: 5628},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 158, col: 40, offset: 5630},
	name: "Word",
},
&ruleRefExpr{
	pos: position{line: 158, col: 45, offset: 5635},
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
	pos: position{line: 168, col: 1, offset: 5888},
	expr: &choiceExpr{
	pos: position{line: 168, col: 18, offset: 5905},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 168, col: 18, offset: 5905},
	run: (*parser).callonJoinCondition2,
	expr: &seqExpr{
	pos: position{line: 168, col: 18, offset: 5905},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 168, col: 18, offset: 5905},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 168, col: 23, offset: 5910},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 168, col: 35, offset: 5922},
	name: "_",
},
&labeledExpr{
	pos: position{line: 168, col: 37, offset: 5924},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 168, col: 40, offset: 5927},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 168, col: 52, offset: 5939},
	name: "_",
},
&labeledExpr{
	pos: position{line: 168, col: 54, offset: 5941},
	label: "right",
	expr: &ruleRefExpr{
	pos: position{line: 168, col: 60, offset: 5947},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 168, col: 72, offset: 5959},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 170, col: 4, offset: 6077},
	run: (*parser).callonJoinCondition13,
	expr: &seqExpr{
	pos: position{line: 170, col: 4, offset: 6077},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 170, col: 4, offset: 6077},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 170, col: 9, offset: 6082},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 170, col: 21, offset: 6094},
	name: "_",
},
&labeledExpr{
	pos: position{line: 170, col: 23, offset: 6096},
	label: "right",
	expr: &ruleRefExpr{
	pos: position{line: 170, col: 29, offset: 6102},
	name: "JoinFacotors",
},
},
&ruleRefExpr{
	pos: position{line: 170, col: 42, offset: 6115},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 172, col: 3, offset: 6198},
	run: (*parser).callonJoinCondition21,
	expr: &seqExpr{
	pos: position{line: 172, col: 3, offset: 6198},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 172, col: 3, offset: 6198},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 172, col: 8, offset: 6203},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 172, col: 20, offset: 6215},
	name: "_",
},
&labeledExpr{
	pos: position{line: 172, col: 22, offset: 6217},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 172, col: 25, offset: 6220},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 172, col: 37, offset: 6232},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 174, col: 3, offset: 6314},
	run: (*parser).callonJoinCondition29,
	expr: &seqExpr{
	pos: position{line: 174, col: 3, offset: 6314},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 174, col: 3, offset: 6314},
	label: "left",
	expr: &ruleRefExpr{
	pos: position{line: 174, col: 8, offset: 6319},
	name: "JoinFactors",
},
},
&ruleRefExpr{
	pos: position{line: 174, col: 20, offset: 6331},
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
	pos: position{line: 179, col: 1, offset: 6426},
	expr: &actionExpr{
	pos: position{line: 179, col: 16, offset: 6441},
	run: (*parser).callonJoinFactors1,
	expr: &seqExpr{
	pos: position{line: 179, col: 16, offset: 6441},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 179, col: 16, offset: 6441},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 179, col: 22, offset: 6447},
	name: "Word",
},
},
&ruleRefExpr{
	pos: position{line: 179, col: 27, offset: 6452},
	name: "_",
},
&labeledExpr{
	pos: position{line: 179, col: 29, offset: 6454},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 179, col: 34, offset: 6459},
	expr: &seqExpr{
	pos: position{line: 179, col: 35, offset: 6460},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 179, col: 35, offset: 6460},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 179, col: 39, offset: 6464},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 179, col: 41, offset: 6466},
	name: "Word",
},
&ruleRefExpr{
	pos: position{line: 179, col: 46, offset: 6471},
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
	pos: position{line: 192, col: 1, offset: 6809},
	expr: &actionExpr{
	pos: position{line: 192, col: 12, offset: 6820},
	run: (*parser).callonMapJobG1,
	expr: &seqExpr{
	pos: position{line: 192, col: 12, offset: 6820},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 192, col: 12, offset: 6820},
	val: "map",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 192, col: 18, offset: 6826},
	name: "_",
},
&labeledExpr{
	pos: position{line: 192, col: 20, offset: 6828},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 192, col: 25, offset: 6833},
	name: "EnrichDo",
},
},
	},
},
},
},
{
	name: "EnrichDo",
	pos: position{line: 197, col: 1, offset: 6947},
	expr: &actionExpr{
	pos: position{line: 197, col: 13, offset: 6959},
	run: (*parser).callonEnrichDo1,
	expr: &seqExpr{
	pos: position{line: 197, col: 13, offset: 6959},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 197, col: 13, offset: 6959},
	label: "fld",
	expr: &ruleRefExpr{
	pos: position{line: 197, col: 17, offset: 6963},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 197, col: 26, offset: 6972},
	name: "_",
},
&litMatcher{
	pos: position{line: 197, col: 28, offset: 6974},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 197, col: 32, offset: 6978},
	name: "_",
},
&labeledExpr{
	pos: position{line: 197, col: 34, offset: 6980},
	label: "expr",
	expr: &ruleRefExpr{
	pos: position{line: 197, col: 39, offset: 6985},
	name: "MExpr",
},
},
	},
},
},
},
{
	name: "AggJobG",
	pos: position{line: 211, col: 1, offset: 7285},
	expr: &actionExpr{
	pos: position{line: 211, col: 12, offset: 7296},
	run: (*parser).callonAggJobG1,
	expr: &seqExpr{
	pos: position{line: 211, col: 12, offset: 7296},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 211, col: 12, offset: 7296},
	val: "agg",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 211, col: 18, offset: 7302},
	name: "_",
},
&labeledExpr{
	pos: position{line: 211, col: 20, offset: 7304},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 211, col: 27, offset: 7311},
	expr: &ruleRefExpr{
	pos: position{line: 211, col: 27, offset: 7311},
	name: "AggGroupBy",
},
},
},
&ruleRefExpr{
	pos: position{line: 211, col: 39, offset: 7323},
	name: "_",
},
&labeledExpr{
	pos: position{line: 211, col: 41, offset: 7325},
	label: "job",
	expr: &ruleRefExpr{
	pos: position{line: 211, col: 45, offset: 7329},
	name: "AggJobs",
},
},
&labeledExpr{
	pos: position{line: 211, col: 53, offset: 7337},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 211, col: 58, offset: 7342},
	expr: &seqExpr{
	pos: position{line: 211, col: 59, offset: 7343},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 211, col: 59, offset: 7343},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 211, col: 63, offset: 7347},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 211, col: 65, offset: 7349},
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
	pos: position{line: 226, col: 1, offset: 7724},
	expr: &choiceExpr{
	pos: position{line: 226, col: 12, offset: 7735},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 226, col: 12, offset: 7735},
	name: "CountAgg",
},
&ruleRefExpr{
	pos: position{line: 226, col: 23, offset: 7746},
	name: "MaxAgg",
},
&ruleRefExpr{
	pos: position{line: 226, col: 32, offset: 7755},
	name: "MinAgg",
},
&ruleRefExpr{
	pos: position{line: 226, col: 41, offset: 7764},
	name: "AvgAgg",
},
&ruleRefExpr{
	pos: position{line: 226, col: 50, offset: 7773},
	name: "SumAgg",
},
&ruleRefExpr{
	pos: position{line: 226, col: 59, offset: 7782},
	name: "ModeAgg",
},
&ruleRefExpr{
	pos: position{line: 227, col: 13, offset: 7803},
	name: "VarAgg",
},
&ruleRefExpr{
	pos: position{line: 227, col: 22, offset: 7812},
	name: "DistinctCountAgg",
},
&ruleRefExpr{
	pos: position{line: 227, col: 41, offset: 7831},
	name: "QuantileAgg",
},
&ruleRefExpr{
	pos: position{line: 227, col: 55, offset: 7845},
	name: "MedianAgg",
},
&ruleRefExpr{
	pos: position{line: 227, col: 67, offset: 7857},
	name: "QuartileAgg",
},
&ruleRefExpr{
	pos: position{line: 228, col: 13, offset: 7882},
	name: "CovAgg",
},
&ruleRefExpr{
	pos: position{line: 228, col: 22, offset: 7891},
	name: "CorrAgg",
},
&ruleRefExpr{
	pos: position{line: 228, col: 32, offset: 7901},
	name: "PValueAgg",
},
	},
},
},
{
	name: "CountAgg",
	pos: position{line: 231, col: 1, offset: 7945},
	expr: &actionExpr{
	pos: position{line: 231, col: 13, offset: 7957},
	run: (*parser).callonCountAgg1,
	expr: &seqExpr{
	pos: position{line: 231, col: 13, offset: 7957},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 231, col: 13, offset: 7957},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 231, col: 19, offset: 7963},
	expr: &ruleRefExpr{
	pos: position{line: 231, col: 19, offset: 7963},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 231, col: 29, offset: 7973},
	name: "_",
},
&litMatcher{
	pos: position{line: 231, col: 31, offset: 7975},
	val: "count(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 231, col: 40, offset: 7984},
	name: "_",
},
&labeledExpr{
	pos: position{line: 231, col: 42, offset: 7986},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 231, col: 47, offset: 7991},
	expr: &ruleRefExpr{
	pos: position{line: 231, col: 47, offset: 7991},
	name: "FilterMulti",
},
},
},
&ruleRefExpr{
	pos: position{line: 231, col: 60, offset: 8004},
	name: "_",
},
&litMatcher{
	pos: position{line: 231, col: 62, offset: 8006},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MaxAgg",
	pos: position{line: 238, col: 1, offset: 8158},
	expr: &actionExpr{
	pos: position{line: 238, col: 11, offset: 8168},
	run: (*parser).callonMaxAgg1,
	expr: &seqExpr{
	pos: position{line: 238, col: 11, offset: 8168},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 238, col: 11, offset: 8168},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 238, col: 17, offset: 8174},
	expr: &ruleRefExpr{
	pos: position{line: 238, col: 17, offset: 8174},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 238, col: 27, offset: 8184},
	name: "_",
},
&litMatcher{
	pos: position{line: 238, col: 29, offset: 8186},
	val: "max(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 238, col: 36, offset: 8193},
	name: "_",
},
&labeledExpr{
	pos: position{line: 238, col: 38, offset: 8195},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 238, col: 44, offset: 8201},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 238, col: 53, offset: 8210},
	name: "_",
},
&labeledExpr{
	pos: position{line: 238, col: 55, offset: 8212},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 238, col: 60, offset: 8217},
	expr: &seqExpr{
	pos: position{line: 238, col: 61, offset: 8218},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 238, col: 61, offset: 8218},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 238, col: 65, offset: 8222},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 238, col: 67, offset: 8224},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 238, col: 81, offset: 8238},
	name: "_",
},
&litMatcher{
	pos: position{line: 238, col: 83, offset: 8240},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MinAgg",
	pos: position{line: 260, col: 1, offset: 8701},
	expr: &actionExpr{
	pos: position{line: 260, col: 11, offset: 8711},
	run: (*parser).callonMinAgg1,
	expr: &seqExpr{
	pos: position{line: 260, col: 11, offset: 8711},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 260, col: 11, offset: 8711},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 260, col: 17, offset: 8717},
	expr: &ruleRefExpr{
	pos: position{line: 260, col: 17, offset: 8717},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 260, col: 27, offset: 8727},
	name: "_",
},
&litMatcher{
	pos: position{line: 260, col: 29, offset: 8729},
	val: "min(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 260, col: 36, offset: 8736},
	name: "_",
},
&labeledExpr{
	pos: position{line: 260, col: 38, offset: 8738},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 260, col: 44, offset: 8744},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 260, col: 53, offset: 8753},
	name: "_",
},
&labeledExpr{
	pos: position{line: 260, col: 55, offset: 8755},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 260, col: 60, offset: 8760},
	expr: &seqExpr{
	pos: position{line: 260, col: 61, offset: 8761},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 260, col: 61, offset: 8761},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 260, col: 65, offset: 8765},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 260, col: 67, offset: 8767},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 260, col: 81, offset: 8781},
	name: "_",
},
&litMatcher{
	pos: position{line: 260, col: 83, offset: 8783},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AvgAgg",
	pos: position{line: 282, col: 1, offset: 9244},
	expr: &actionExpr{
	pos: position{line: 282, col: 11, offset: 9254},
	run: (*parser).callonAvgAgg1,
	expr: &seqExpr{
	pos: position{line: 282, col: 11, offset: 9254},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 282, col: 11, offset: 9254},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 282, col: 17, offset: 9260},
	expr: &ruleRefExpr{
	pos: position{line: 282, col: 17, offset: 9260},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 282, col: 27, offset: 9270},
	name: "_",
},
&litMatcher{
	pos: position{line: 282, col: 29, offset: 9272},
	val: "avg(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 282, col: 36, offset: 9279},
	name: "_",
},
&labeledExpr{
	pos: position{line: 282, col: 38, offset: 9281},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 282, col: 44, offset: 9287},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 282, col: 53, offset: 9296},
	name: "_",
},
&labeledExpr{
	pos: position{line: 282, col: 55, offset: 9298},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 282, col: 60, offset: 9303},
	expr: &seqExpr{
	pos: position{line: 282, col: 61, offset: 9304},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 282, col: 61, offset: 9304},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 282, col: 65, offset: 9308},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 282, col: 67, offset: 9310},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 282, col: 81, offset: 9324},
	name: "_",
},
&litMatcher{
	pos: position{line: 282, col: 83, offset: 9326},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SumAgg",
	pos: position{line: 304, col: 1, offset: 9787},
	expr: &actionExpr{
	pos: position{line: 304, col: 11, offset: 9797},
	run: (*parser).callonSumAgg1,
	expr: &seqExpr{
	pos: position{line: 304, col: 11, offset: 9797},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 304, col: 11, offset: 9797},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 304, col: 17, offset: 9803},
	expr: &ruleRefExpr{
	pos: position{line: 304, col: 17, offset: 9803},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 304, col: 27, offset: 9813},
	name: "_",
},
&litMatcher{
	pos: position{line: 304, col: 29, offset: 9815},
	val: "sum(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 304, col: 36, offset: 9822},
	name: "_",
},
&labeledExpr{
	pos: position{line: 304, col: 38, offset: 9824},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 304, col: 44, offset: 9830},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 304, col: 53, offset: 9839},
	name: "_",
},
&labeledExpr{
	pos: position{line: 304, col: 55, offset: 9841},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 304, col: 60, offset: 9846},
	expr: &seqExpr{
	pos: position{line: 304, col: 61, offset: 9847},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 304, col: 61, offset: 9847},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 304, col: 65, offset: 9851},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 304, col: 67, offset: 9853},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 304, col: 81, offset: 9867},
	name: "_",
},
&litMatcher{
	pos: position{line: 304, col: 83, offset: 9869},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "VarAgg",
	pos: position{line: 326, col: 1, offset: 10330},
	expr: &actionExpr{
	pos: position{line: 326, col: 11, offset: 10340},
	run: (*parser).callonVarAgg1,
	expr: &seqExpr{
	pos: position{line: 326, col: 11, offset: 10340},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 326, col: 11, offset: 10340},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 326, col: 17, offset: 10346},
	expr: &ruleRefExpr{
	pos: position{line: 326, col: 17, offset: 10346},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 326, col: 27, offset: 10356},
	name: "_",
},
&litMatcher{
	pos: position{line: 326, col: 29, offset: 10358},
	val: "var(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 326, col: 36, offset: 10365},
	name: "_",
},
&labeledExpr{
	pos: position{line: 326, col: 38, offset: 10367},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 326, col: 44, offset: 10373},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 326, col: 53, offset: 10382},
	name: "_",
},
&labeledExpr{
	pos: position{line: 326, col: 55, offset: 10384},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 326, col: 60, offset: 10389},
	expr: &seqExpr{
	pos: position{line: 326, col: 61, offset: 10390},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 326, col: 61, offset: 10390},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 326, col: 65, offset: 10394},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 326, col: 67, offset: 10396},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 326, col: 81, offset: 10410},
	name: "_",
},
&litMatcher{
	pos: position{line: 326, col: 83, offset: 10412},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "DistinctCountAgg",
	pos: position{line: 348, col: 1, offset: 10878},
	expr: &actionExpr{
	pos: position{line: 348, col: 21, offset: 10898},
	run: (*parser).callonDistinctCountAgg1,
	expr: &seqExpr{
	pos: position{line: 348, col: 21, offset: 10898},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 348, col: 21, offset: 10898},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 348, col: 27, offset: 10904},
	expr: &ruleRefExpr{
	pos: position{line: 348, col: 27, offset: 10904},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 348, col: 37, offset: 10914},
	name: "_",
},
&litMatcher{
	pos: position{line: 348, col: 39, offset: 10916},
	val: "dcount(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 348, col: 49, offset: 10926},
	name: "_",
},
&labeledExpr{
	pos: position{line: 348, col: 51, offset: 10928},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 348, col: 57, offset: 10934},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 348, col: 66, offset: 10943},
	name: "_",
},
&labeledExpr{
	pos: position{line: 348, col: 68, offset: 10945},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 348, col: 73, offset: 10950},
	expr: &seqExpr{
	pos: position{line: 348, col: 74, offset: 10951},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 348, col: 74, offset: 10951},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 348, col: 78, offset: 10955},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 348, col: 80, offset: 10957},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 348, col: 94, offset: 10971},
	name: "_",
},
&litMatcher{
	pos: position{line: 348, col: 96, offset: 10973},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "ModeAgg",
	pos: position{line: 370, col: 1, offset: 11444},
	expr: &actionExpr{
	pos: position{line: 370, col: 12, offset: 11455},
	run: (*parser).callonModeAgg1,
	expr: &seqExpr{
	pos: position{line: 370, col: 12, offset: 11455},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 370, col: 12, offset: 11455},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 370, col: 18, offset: 11461},
	expr: &ruleRefExpr{
	pos: position{line: 370, col: 18, offset: 11461},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 370, col: 28, offset: 11471},
	name: "_",
},
&litMatcher{
	pos: position{line: 370, col: 30, offset: 11473},
	val: "mode(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 370, col: 38, offset: 11481},
	name: "_",
},
&labeledExpr{
	pos: position{line: 370, col: 40, offset: 11483},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 370, col: 46, offset: 11489},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 370, col: 55, offset: 11498},
	name: "_",
},
&labeledExpr{
	pos: position{line: 370, col: 57, offset: 11500},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 370, col: 62, offset: 11505},
	expr: &seqExpr{
	pos: position{line: 370, col: 63, offset: 11506},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 370, col: 63, offset: 11506},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 370, col: 67, offset: 11510},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 370, col: 69, offset: 11512},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 370, col: 83, offset: 11526},
	name: "_",
},
&litMatcher{
	pos: position{line: 370, col: 85, offset: 11528},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuantileAgg",
	pos: position{line: 392, col: 1, offset: 11990},
	expr: &actionExpr{
	pos: position{line: 392, col: 16, offset: 12005},
	run: (*parser).callonQuantileAgg1,
	expr: &seqExpr{
	pos: position{line: 392, col: 16, offset: 12005},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 392, col: 16, offset: 12005},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 392, col: 22, offset: 12011},
	expr: &ruleRefExpr{
	pos: position{line: 392, col: 22, offset: 12011},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 392, col: 32, offset: 12021},
	name: "_",
},
&litMatcher{
	pos: position{line: 392, col: 34, offset: 12023},
	val: "quantile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 392, col: 46, offset: 12035},
	name: "_",
},
&labeledExpr{
	pos: position{line: 392, col: 48, offset: 12037},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 392, col: 54, offset: 12043},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 392, col: 63, offset: 12052},
	name: "_",
},
&labeledExpr{
	pos: position{line: 392, col: 66, offset: 12055},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 392, col: 73, offset: 12062},
	expr: &seqExpr{
	pos: position{line: 392, col: 74, offset: 12063},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 392, col: 74, offset: 12063},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 392, col: 78, offset: 12067},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 392, col: 80, offset: 12069},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 392, col: 91, offset: 12080},
	name: "_",
},
&litMatcher{
	pos: position{line: 392, col: 93, offset: 12082},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 392, col: 97, offset: 12086},
	name: "_",
},
&labeledExpr{
	pos: position{line: 392, col: 99, offset: 12088},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 392, col: 103, offset: 12092},
	name: "Float",
},
},
&ruleRefExpr{
	pos: position{line: 392, col: 109, offset: 12098},
	name: "_",
},
&labeledExpr{
	pos: position{line: 392, col: 111, offset: 12100},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 392, col: 116, offset: 12105},
	expr: &seqExpr{
	pos: position{line: 392, col: 117, offset: 12106},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 392, col: 117, offset: 12106},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 392, col: 121, offset: 12110},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 392, col: 123, offset: 12112},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 392, col: 137, offset: 12126},
	name: "_",
},
&litMatcher{
	pos: position{line: 392, col: 139, offset: 12128},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "MedianAgg",
	pos: position{line: 422, col: 1, offset: 12773},
	expr: &actionExpr{
	pos: position{line: 422, col: 14, offset: 12786},
	run: (*parser).callonMedianAgg1,
	expr: &seqExpr{
	pos: position{line: 422, col: 14, offset: 12786},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 422, col: 14, offset: 12786},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 422, col: 20, offset: 12792},
	expr: &ruleRefExpr{
	pos: position{line: 422, col: 20, offset: 12792},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 422, col: 30, offset: 12802},
	name: "_",
},
&litMatcher{
	pos: position{line: 422, col: 32, offset: 12804},
	val: "median(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 422, col: 42, offset: 12814},
	name: "_",
},
&labeledExpr{
	pos: position{line: 422, col: 44, offset: 12816},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 422, col: 50, offset: 12822},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 422, col: 59, offset: 12831},
	name: "_",
},
&labeledExpr{
	pos: position{line: 422, col: 62, offset: 12834},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 422, col: 69, offset: 12841},
	expr: &seqExpr{
	pos: position{line: 422, col: 70, offset: 12842},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 422, col: 70, offset: 12842},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 422, col: 74, offset: 12846},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 422, col: 76, offset: 12848},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 422, col: 87, offset: 12859},
	name: "_",
},
&labeledExpr{
	pos: position{line: 422, col: 89, offset: 12861},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 422, col: 94, offset: 12866},
	expr: &seqExpr{
	pos: position{line: 422, col: 95, offset: 12867},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 422, col: 95, offset: 12867},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 422, col: 99, offset: 12871},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 422, col: 101, offset: 12873},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 422, col: 115, offset: 12887},
	name: "_",
},
&litMatcher{
	pos: position{line: 422, col: 117, offset: 12889},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "QuartileAgg",
	pos: position{line: 454, col: 1, offset: 13566},
	expr: &actionExpr{
	pos: position{line: 454, col: 16, offset: 13581},
	run: (*parser).callonQuartileAgg1,
	expr: &seqExpr{
	pos: position{line: 454, col: 16, offset: 13581},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 454, col: 16, offset: 13581},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 454, col: 22, offset: 13587},
	expr: &ruleRefExpr{
	pos: position{line: 454, col: 22, offset: 13587},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 454, col: 32, offset: 13597},
	name: "_",
},
&litMatcher{
	pos: position{line: 454, col: 34, offset: 13599},
	val: "quartile(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 454, col: 46, offset: 13611},
	name: "_",
},
&labeledExpr{
	pos: position{line: 454, col: 48, offset: 13613},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 454, col: 54, offset: 13619},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 454, col: 63, offset: 13628},
	name: "_",
},
&labeledExpr{
	pos: position{line: 454, col: 66, offset: 13631},
	label: "weight",
	expr: &zeroOrOneExpr{
	pos: position{line: 454, col: 73, offset: 13638},
	expr: &seqExpr{
	pos: position{line: 454, col: 74, offset: 13639},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 454, col: 74, offset: 13639},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 454, col: 78, offset: 13643},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 454, col: 80, offset: 13645},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 454, col: 91, offset: 13656},
	name: "_",
},
&litMatcher{
	pos: position{line: 454, col: 93, offset: 13658},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 454, col: 97, offset: 13662},
	name: "_",
},
&labeledExpr{
	pos: position{line: 454, col: 99, offset: 13664},
	label: "qth",
	expr: &ruleRefExpr{
	pos: position{line: 454, col: 103, offset: 13668},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 454, col: 110, offset: 13675},
	name: "_",
},
&labeledExpr{
	pos: position{line: 454, col: 112, offset: 13677},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 454, col: 117, offset: 13682},
	expr: &seqExpr{
	pos: position{line: 454, col: 118, offset: 13683},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 454, col: 118, offset: 13683},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 454, col: 122, offset: 13687},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 454, col: 124, offset: 13689},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 454, col: 138, offset: 13703},
	name: "_",
},
&litMatcher{
	pos: position{line: 454, col: 140, offset: 13705},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CovAgg",
	pos: position{line: 500, col: 1, offset: 14732},
	expr: &actionExpr{
	pos: position{line: 500, col: 11, offset: 14742},
	run: (*parser).callonCovAgg1,
	expr: &seqExpr{
	pos: position{line: 500, col: 11, offset: 14742},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 500, col: 11, offset: 14742},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 500, col: 17, offset: 14748},
	expr: &ruleRefExpr{
	pos: position{line: 500, col: 17, offset: 14748},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 500, col: 27, offset: 14758},
	name: "_",
},
&litMatcher{
	pos: position{line: 500, col: 29, offset: 14760},
	val: "cov(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 500, col: 36, offset: 14767},
	name: "_",
},
&labeledExpr{
	pos: position{line: 500, col: 38, offset: 14769},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 500, col: 45, offset: 14776},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 500, col: 54, offset: 14785},
	name: "_",
},
&litMatcher{
	pos: position{line: 500, col: 56, offset: 14787},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 500, col: 60, offset: 14791},
	name: "_",
},
&labeledExpr{
	pos: position{line: 500, col: 62, offset: 14793},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 500, col: 69, offset: 14800},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 500, col: 78, offset: 14809},
	name: "_",
},
&labeledExpr{
	pos: position{line: 500, col: 80, offset: 14811},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 500, col: 85, offset: 14816},
	expr: &seqExpr{
	pos: position{line: 500, col: 86, offset: 14817},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 500, col: 86, offset: 14817},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 500, col: 90, offset: 14821},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 500, col: 92, offset: 14823},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 500, col: 106, offset: 14837},
	name: "_",
},
&litMatcher{
	pos: position{line: 500, col: 108, offset: 14839},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "CorrAgg",
	pos: position{line: 523, col: 1, offset: 15364},
	expr: &actionExpr{
	pos: position{line: 523, col: 12, offset: 15375},
	run: (*parser).callonCorrAgg1,
	expr: &seqExpr{
	pos: position{line: 523, col: 12, offset: 15375},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 523, col: 12, offset: 15375},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 523, col: 18, offset: 15381},
	expr: &ruleRefExpr{
	pos: position{line: 523, col: 18, offset: 15381},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 523, col: 28, offset: 15391},
	name: "_",
},
&litMatcher{
	pos: position{line: 523, col: 30, offset: 15393},
	val: "correlation(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 523, col: 45, offset: 15408},
	name: "_",
},
&labeledExpr{
	pos: position{line: 523, col: 47, offset: 15410},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 523, col: 54, offset: 15417},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 523, col: 63, offset: 15426},
	name: "_",
},
&litMatcher{
	pos: position{line: 523, col: 65, offset: 15428},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 523, col: 69, offset: 15432},
	name: "_",
},
&labeledExpr{
	pos: position{line: 523, col: 71, offset: 15434},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 523, col: 78, offset: 15441},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 523, col: 87, offset: 15450},
	name: "_",
},
&labeledExpr{
	pos: position{line: 523, col: 89, offset: 15452},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 523, col: 94, offset: 15457},
	expr: &seqExpr{
	pos: position{line: 523, col: 95, offset: 15458},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 523, col: 95, offset: 15458},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 523, col: 99, offset: 15462},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 523, col: 101, offset: 15464},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 523, col: 115, offset: 15478},
	name: "_",
},
&litMatcher{
	pos: position{line: 523, col: 117, offset: 15480},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PValueAgg",
	pos: position{line: 546, col: 1, offset: 16006},
	expr: &actionExpr{
	pos: position{line: 546, col: 14, offset: 16019},
	run: (*parser).callonPValueAgg1,
	expr: &seqExpr{
	pos: position{line: 546, col: 14, offset: 16019},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 546, col: 14, offset: 16019},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 546, col: 20, offset: 16025},
	expr: &ruleRefExpr{
	pos: position{line: 546, col: 20, offset: 16025},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 546, col: 30, offset: 16035},
	name: "_",
},
&litMatcher{
	pos: position{line: 546, col: 32, offset: 16037},
	val: "pvalue(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 546, col: 42, offset: 16047},
	name: "_",
},
&labeledExpr{
	pos: position{line: 546, col: 44, offset: 16049},
	label: "field1",
	expr: &ruleRefExpr{
	pos: position{line: 546, col: 51, offset: 16056},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 546, col: 60, offset: 16065},
	name: "_",
},
&litMatcher{
	pos: position{line: 546, col: 62, offset: 16067},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 546, col: 66, offset: 16071},
	name: "_",
},
&labeledExpr{
	pos: position{line: 546, col: 68, offset: 16073},
	label: "field2",
	expr: &ruleRefExpr{
	pos: position{line: 546, col: 75, offset: 16080},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 546, col: 84, offset: 16089},
	name: "_",
},
&labeledExpr{
	pos: position{line: 546, col: 86, offset: 16091},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 546, col: 91, offset: 16096},
	expr: &seqExpr{
	pos: position{line: 546, col: 92, offset: 16097},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 546, col: 92, offset: 16097},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 546, col: 96, offset: 16101},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 546, col: 98, offset: 16103},
	name: "FilterMulti",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 546, col: 112, offset: 16117},
	name: "_",
},
&litMatcher{
	pos: position{line: 546, col: 114, offset: 16119},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "AggGroupBy",
	pos: position{line: 571, col: 1, offset: 16644},
	expr: &actionExpr{
	pos: position{line: 571, col: 15, offset: 16658},
	run: (*parser).callonAggGroupBy1,
	expr: &seqExpr{
	pos: position{line: 571, col: 15, offset: 16658},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 571, col: 15, offset: 16658},
	name: "_",
},
&litMatcher{
	pos: position{line: 571, col: 17, offset: 16660},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 571, col: 22, offset: 16665},
	name: "_",
},
&labeledExpr{
	pos: position{line: 571, col: 24, offset: 16667},
	label: "param",
	expr: &ruleRefExpr{
	pos: position{line: 571, col: 30, offset: 16673},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 571, col: 39, offset: 16682},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 571, col: 44, offset: 16687},
	expr: &seqExpr{
	pos: position{line: 571, col: 45, offset: 16688},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 571, col: 45, offset: 16688},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 571, col: 49, offset: 16692},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 571, col: 51, offset: 16694},
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
	pos: position{line: 586, col: 1, offset: 17045},
	expr: &actionExpr{
	pos: position{line: 586, col: 11, offset: 17055},
	run: (*parser).callonDoJobG1,
	expr: &labeledExpr{
	pos: position{line: 586, col: 11, offset: 17055},
	label: "job",
	expr: &choiceExpr{
	pos: position{line: 586, col: 16, offset: 17060},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 586, col: 16, offset: 17060},
	name: "FilterDo",
},
&ruleRefExpr{
	pos: position{line: 586, col: 27, offset: 17071},
	name: "BranchDo",
},
&ruleRefExpr{
	pos: position{line: 586, col: 38, offset: 17082},
	name: "SelectDo",
},
&ruleRefExpr{
	pos: position{line: 586, col: 49, offset: 17093},
	name: "PickDo",
},
&ruleRefExpr{
	pos: position{line: 586, col: 58, offset: 17102},
	name: "SortDo",
},
&ruleRefExpr{
	pos: position{line: 586, col: 67, offset: 17111},
	name: "BatchDo",
},
	},
},
},
},
},
{
	name: "BranchDo",
	pos: position{line: 592, col: 1, offset: 17230},
	expr: &actionExpr{
	pos: position{line: 592, col: 13, offset: 17242},
	run: (*parser).callonBranchDo1,
	expr: &seqExpr{
	pos: position{line: 592, col: 13, offset: 17242},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 592, col: 13, offset: 17242},
	val: "branch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 592, col: 23, offset: 17252},
	name: "_",
},
&labeledExpr{
	pos: position{line: 592, col: 25, offset: 17254},
	label: "br",
	expr: &ruleRefExpr{
	pos: position{line: 592, col: 28, offset: 17257},
	name: "BranchPrimary",
},
},
&ruleRefExpr{
	pos: position{line: 592, col: 42, offset: 17271},
	name: "_",
},
&litMatcher{
	pos: position{line: 592, col: 44, offset: 17273},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "BranchPrimary",
	pos: position{line: 599, col: 1, offset: 17468},
	expr: &actionExpr{
	pos: position{line: 599, col: 18, offset: 17485},
	run: (*parser).callonBranchPrimary1,
	expr: &seqExpr{
	pos: position{line: 599, col: 18, offset: 17485},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 599, col: 18, offset: 17485},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 599, col: 23, offset: 17490},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 599, col: 28, offset: 17495},
	name: "_",
},
&labeledExpr{
	pos: position{line: 599, col: 30, offset: 17497},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 599, col: 33, offset: 17500},
	name: "BranchPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 599, col: 45, offset: 17512},
	name: "_",
},
&labeledExpr{
	pos: position{line: 599, col: 47, offset: 17514},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 599, col: 54, offset: 17521},
	name: "Value",
},
},
	},
},
},
},
{
	name: "BranchPriOp",
	pos: position{line: 603, col: 1, offset: 17572},
	expr: &ruleRefExpr{
	pos: position{line: 603, col: 16, offset: 17587},
	name: "FilterPriOp",
},
},
{
	name: "SelectDo",
	pos: position{line: 606, col: 1, offset: 17621},
	expr: &actionExpr{
	pos: position{line: 606, col: 13, offset: 17633},
	run: (*parser).callonSelectDo1,
	expr: &seqExpr{
	pos: position{line: 606, col: 13, offset: 17633},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 606, col: 13, offset: 17633},
	val: "select(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 606, col: 23, offset: 17643},
	name: "_",
},
&labeledExpr{
	pos: position{line: 606, col: 25, offset: 17645},
	label: "field",
	expr: &ruleRefExpr{
	pos: position{line: 606, col: 31, offset: 17651},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 606, col: 40, offset: 17660},
	name: "_",
},
&labeledExpr{
	pos: position{line: 606, col: 42, offset: 17662},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 606, col: 47, offset: 17667},
	expr: &seqExpr{
	pos: position{line: 606, col: 48, offset: 17668},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 606, col: 48, offset: 17668},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 606, col: 52, offset: 17672},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 606, col: 54, offset: 17674},
	name: "Variable",
},
	},
},
},
},
&ruleRefExpr{
	pos: position{line: 606, col: 65, offset: 17685},
	name: "_",
},
&litMatcher{
	pos: position{line: 606, col: 67, offset: 17687},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "PickDo",
	pos: position{line: 621, col: 1, offset: 18024},
	expr: &actionExpr{
	pos: position{line: 621, col: 11, offset: 18034},
	run: (*parser).callonPickDo1,
	expr: &seqExpr{
	pos: position{line: 621, col: 11, offset: 18034},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 621, col: 11, offset: 18034},
	val: "pick",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 621, col: 18, offset: 18041},
	name: "_",
},
&labeledExpr{
	pos: position{line: 621, col: 20, offset: 18043},
	label: "desc",
	expr: &ruleRefExpr{
	pos: position{line: 621, col: 25, offset: 18048},
	name: "Variable",
},
},
&litMatcher{
	pos: position{line: 621, col: 34, offset: 18057},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 621, col: 38, offset: 18061},
	name: "_",
},
&labeledExpr{
	pos: position{line: 621, col: 40, offset: 18063},
	label: "num",
	expr: &ruleRefExpr{
	pos: position{line: 621, col: 44, offset: 18067},
	name: "Number",
},
},
&ruleRefExpr{
	pos: position{line: 621, col: 51, offset: 18074},
	name: "_",
},
&litMatcher{
	pos: position{line: 621, col: 53, offset: 18076},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "SortDo",
	pos: position{line: 633, col: 1, offset: 18298},
	expr: &actionExpr{
	pos: position{line: 633, col: 11, offset: 18308},
	run: (*parser).callonSortDo1,
	expr: &seqExpr{
	pos: position{line: 633, col: 11, offset: 18308},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 633, col: 11, offset: 18308},
	val: "sort()",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 633, col: 20, offset: 18317},
	name: "_",
},
&litMatcher{
	pos: position{line: 633, col: 22, offset: 18319},
	val: "by",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 633, col: 27, offset: 18324},
	name: "_",
},
&labeledExpr{
	pos: position{line: 633, col: 29, offset: 18326},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 633, col: 35, offset: 18332},
	name: "Variable",
},
},
	},
},
},
},
{
	name: "BatchDo",
	pos: position{line: 639, col: 1, offset: 18438},
	expr: &actionExpr{
	pos: position{line: 639, col: 12, offset: 18449},
	run: (*parser).callonBatchDo1,
	expr: &seqExpr{
	pos: position{line: 639, col: 12, offset: 18449},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 639, col: 12, offset: 18449},
	val: "batch(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 639, col: 21, offset: 18458},
	name: "_",
},
&labeledExpr{
	pos: position{line: 639, col: 23, offset: 18460},
	label: "num",
	expr: &zeroOrOneExpr{
	pos: position{line: 639, col: 27, offset: 18464},
	expr: &ruleRefExpr{
	pos: position{line: 639, col: 27, offset: 18464},
	name: "Number",
},
},
},
&ruleRefExpr{
	pos: position{line: 639, col: 35, offset: 18472},
	name: "_",
},
&litMatcher{
	pos: position{line: 639, col: 37, offset: 18474},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterDo",
	pos: position{line: 649, col: 1, offset: 18619},
	expr: &actionExpr{
	pos: position{line: 649, col: 13, offset: 18631},
	run: (*parser).callonFilterDo1,
	expr: &seqExpr{
	pos: position{line: 649, col: 13, offset: 18631},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 649, col: 13, offset: 18631},
	val: "filter(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 649, col: 23, offset: 18641},
	name: "_",
},
&labeledExpr{
	pos: position{line: 649, col: 25, offset: 18643},
	label: "filt",
	expr: &ruleRefExpr{
	pos: position{line: 649, col: 30, offset: 18648},
	name: "FilterMulti",
},
},
&ruleRefExpr{
	pos: position{line: 649, col: 42, offset: 18660},
	name: "_",
},
&litMatcher{
	pos: position{line: 649, col: 44, offset: 18662},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FilterMulti",
	pos: position{line: 659, col: 1, offset: 18950},
	expr: &actionExpr{
	pos: position{line: 659, col: 16, offset: 18965},
	run: (*parser).callonFilterMulti1,
	expr: &seqExpr{
	pos: position{line: 659, col: 16, offset: 18965},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 659, col: 16, offset: 18965},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 659, col: 22, offset: 18971},
	name: "FilterFactor",
},
},
&labeledExpr{
	pos: position{line: 659, col: 35, offset: 18984},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 659, col: 40, offset: 18989},
	expr: &seqExpr{
	pos: position{line: 659, col: 41, offset: 18990},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 659, col: 41, offset: 18990},
	name: "_",
},
&labeledExpr{
	pos: position{line: 659, col: 43, offset: 18992},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 659, col: 46, offset: 18995},
	name: "FilterSecOp",
},
},
&ruleRefExpr{
	pos: position{line: 659, col: 58, offset: 19007},
	name: "_",
},
&labeledExpr{
	pos: position{line: 659, col: 60, offset: 19009},
	label: "f2",
	expr: &ruleRefExpr{
	pos: position{line: 659, col: 63, offset: 19012},
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
	pos: position{line: 664, col: 1, offset: 19170},
	expr: &choiceExpr{
	pos: position{line: 664, col: 17, offset: 19186},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 664, col: 17, offset: 19186},
	run: (*parser).callonFilterFactor2,
	expr: &seqExpr{
	pos: position{line: 664, col: 17, offset: 19186},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 664, col: 17, offset: 19186},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 664, col: 21, offset: 19190},
	label: "sf",
	expr: &ruleRefExpr{
	pos: position{line: 664, col: 24, offset: 19193},
	name: "FilterMulti",
},
},
&litMatcher{
	pos: position{line: 664, col: 36, offset: 19205},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 666, col: 5, offset: 19236},
	run: (*parser).callonFilterFactor8,
	expr: &labeledExpr{
	pos: position{line: 666, col: 5, offset: 19236},
	label: "pf",
	expr: &ruleRefExpr{
	pos: position{line: 666, col: 8, offset: 19239},
	name: "FilterPrimary",
},
},
},
	},
},
},
{
	name: "FilterSecOp",
	pos: position{line: 670, col: 1, offset: 19281},
	expr: &choiceExpr{
	pos: position{line: 670, col: 16, offset: 19296},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 670, col: 16, offset: 19296},
	val: "and",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 670, col: 24, offset: 19304},
	run: (*parser).callonFilterSecOp3,
	expr: &litMatcher{
	pos: position{line: 670, col: 24, offset: 19304},
	val: "or",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "FilterPrimary",
	pos: position{line: 677, col: 1, offset: 19483},
	expr: &actionExpr{
	pos: position{line: 677, col: 18, offset: 19500},
	run: (*parser).callonFilterPrimary1,
	expr: &seqExpr{
	pos: position{line: 677, col: 18, offset: 19500},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 677, col: 18, offset: 19500},
	label: "name",
	expr: &ruleRefExpr{
	pos: position{line: 677, col: 23, offset: 19505},
	name: "Expr",
},
},
&ruleRefExpr{
	pos: position{line: 677, col: 28, offset: 19510},
	name: "_",
},
&labeledExpr{
	pos: position{line: 677, col: 30, offset: 19512},
	label: "op",
	expr: &ruleRefExpr{
	pos: position{line: 677, col: 33, offset: 19515},
	name: "FilterPriOp",
},
},
&ruleRefExpr{
	pos: position{line: 677, col: 45, offset: 19527},
	name: "_",
},
&labeledExpr{
	pos: position{line: 677, col: 47, offset: 19529},
	label: "value",
	expr: &ruleRefExpr{
	pos: position{line: 677, col: 53, offset: 19535},
	name: "Value",
},
},
&ruleRefExpr{
	pos: position{line: 677, col: 59, offset: 19541},
	name: "_",
},
	},
},
},
},
{
	name: "FilterPriOp",
	pos: position{line: 681, col: 1, offset: 19615},
	expr: &choiceExpr{
	pos: position{line: 681, col: 16, offset: 19630},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 681, col: 16, offset: 19630},
	val: "==",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 681, col: 23, offset: 19637},
	val: "!=",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 681, col: 30, offset: 19644},
	val: "<",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 681, col: 36, offset: 19650},
	val: ">",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 681, col: 42, offset: 19656},
	val: "<=",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 681, col: 49, offset: 19663},
	run: (*parser).callonFilterPriOp7,
	expr: &litMatcher{
	pos: position{line: 681, col: 49, offset: 19663},
	val: ">=",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "StdOutG",
	pos: position{line: 690, col: 1, offset: 19773},
	expr: &choiceExpr{
	pos: position{line: 690, col: 12, offset: 19784},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 690, col: 12, offset: 19784},
	run: (*parser).callonStdOutG2,
	expr: &seqExpr{
	pos: position{line: 690, col: 12, offset: 19784},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 690, col: 12, offset: 19784},
	val: "stdout(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 690, col: 22, offset: 19794},
	name: "_",
},
&labeledExpr{
	pos: position{line: 690, col: 24, offset: 19796},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 690, col: 31, offset: 19803},
	expr: &ruleRefExpr{
	pos: position{line: 690, col: 32, offset: 19804},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 690, col: 48, offset: 19820},
	name: "_",
},
&litMatcher{
	pos: position{line: 690, col: 50, offset: 19822},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 702, col: 3, offset: 20098},
	run: (*parser).callonStdOutG11,
	expr: &litMatcher{
	pos: position{line: 702, col: 3, offset: 20098},
	val: "stdout",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "SinkInto",
	pos: position{line: 706, col: 1, offset: 20168},
	expr: &choiceExpr{
	pos: position{line: 706, col: 12, offset: 20179},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 706, col: 12, offset: 20179},
	run: (*parser).callonSinkInto2,
	expr: &seqExpr{
	pos: position{line: 706, col: 12, offset: 20179},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 706, col: 12, offset: 20179},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 706, col: 18, offset: 20185},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 706, col: 27, offset: 20194},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 706, col: 33, offset: 20200},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 706, col: 33, offset: 20200},
	name: "_",
},
&litMatcher{
	pos: position{line: 706, col: 35, offset: 20202},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 706, col: 39, offset: 20206},
	name: "_",
},
&litMatcher{
	pos: position{line: 706, col: 41, offset: 20208},
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
	pos: position{line: 712, col: 7, offset: 20361},
	run: (*parser).callonSinkInto12,
	expr: &seqExpr{
	pos: position{line: 712, col: 7, offset: 20361},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 712, col: 7, offset: 20361},
	val: "into",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 712, col: 14, offset: 20368},
	name: "_",
},
&litMatcher{
	pos: position{line: 712, col: 16, offset: 20370},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 712, col: 20, offset: 20374},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 712, col: 22, offset: 20376},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 712, col: 31, offset: 20385},
	name: "_",
},
&litMatcher{
	pos: position{line: 712, col: 33, offset: 20387},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 714, col: 7, offset: 20483},
	run: (*parser).callonSinkInto21,
	expr: &seqExpr{
	pos: position{line: 714, col: 7, offset: 20483},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 714, col: 7, offset: 20483},
	val: "into",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 714, col: 14, offset: 20490},
	name: "_",
},
&litMatcher{
	pos: position{line: 714, col: 16, offset: 20492},
	val: "(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 714, col: 20, offset: 20496},
	name: "_",
},
&litMatcher{
	pos: position{line: 714, col: 22, offset: 20498},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 716, col: 7, offset: 20637},
	run: (*parser).callonSinkInto28,
	expr: &seqExpr{
	pos: position{line: 716, col: 7, offset: 20637},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 716, col: 7, offset: 20637},
	val: "into",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 716, col: 14, offset: 20644},
	name: "_",
},
	},
},
},
&actionExpr{
	pos: position{line: 718, col: 7, offset: 20737},
	run: (*parser).callonSinkInto32,
	expr: &seqExpr{
	pos: position{line: 718, col: 7, offset: 20737},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 718, col: 7, offset: 20737},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 718, col: 13, offset: 20743},
	name: "Variable",
},
},
&labeledExpr{
	pos: position{line: 718, col: 22, offset: 20752},
	label: "rest",
	expr: &seqExpr{
	pos: position{line: 718, col: 29, offset: 20759},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 718, col: 29, offset: 20759},
	name: "_",
},
&litMatcher{
	pos: position{line: 718, col: 31, offset: 20761},
	val: "=",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 718, col: 35, offset: 20765},
	name: "_",
},
&litMatcher{
	pos: position{line: 718, col: 37, offset: 20767},
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
	pos: position{line: 722, col: 1, offset: 20838},
	expr: &choiceExpr{
	pos: position{line: 722, col: 15, offset: 20852},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 722, col: 15, offset: 20852},
	run: (*parser).callonBlackHoleG2,
	expr: &seqExpr{
	pos: position{line: 722, col: 15, offset: 20852},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 722, col: 15, offset: 20852},
	val: "blackhole(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 722, col: 28, offset: 20865},
	name: "_",
},
&labeledExpr{
	pos: position{line: 722, col: 30, offset: 20867},
	label: "params",
	expr: &zeroOrOneExpr{
	pos: position{line: 722, col: 37, offset: 20874},
	expr: &ruleRefExpr{
	pos: position{line: 722, col: 38, offset: 20875},
	name: "VarParamListG",
},
},
},
&ruleRefExpr{
	pos: position{line: 722, col: 54, offset: 20891},
	name: "_",
},
&litMatcher{
	pos: position{line: 722, col: 56, offset: 20893},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 724, col: 3, offset: 20948},
	run: (*parser).callonBlackHoleG11,
	expr: &seqExpr{
	pos: position{line: 724, col: 3, offset: 20948},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 724, col: 3, offset: 20948},
	val: "blackhole",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 724, col: 15, offset: 20960},
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
	pos: position{line: 728, col: 1, offset: 21045},
	expr: &choiceExpr{
	pos: position{line: 728, col: 10, offset: 21054},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 728, col: 10, offset: 21054},
	run: (*parser).callonPlotG2,
	expr: &seqExpr{
	pos: position{line: 728, col: 10, offset: 21054},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 728, col: 10, offset: 21054},
	val: "plot",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 728, col: 17, offset: 21061},
	name: "_",
},
&labeledExpr{
	pos: position{line: 728, col: 19, offset: 21063},
	label: "widget",
	expr: &ruleRefExpr{
	pos: position{line: 728, col: 27, offset: 21071},
	name: "BarWidget",
},
},
	},
},
},
&actionExpr{
	pos: position{line: 730, col: 3, offset: 21168},
	run: (*parser).callonPlotG8,
	expr: &seqExpr{
	pos: position{line: 730, col: 3, offset: 21168},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 730, col: 3, offset: 21168},
	val: "plot",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 730, col: 10, offset: 21175},
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
	pos: position{line: 734, col: 1, offset: 21235},
	expr: &actionExpr{
	pos: position{line: 734, col: 14, offset: 21248},
	run: (*parser).callonBarWidget1,
	expr: &seqExpr{
	pos: position{line: 734, col: 14, offset: 21248},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 734, col: 14, offset: 21248},
	label: "alias",
	expr: &zeroOrOneExpr{
	pos: position{line: 734, col: 20, offset: 21254},
	expr: &ruleRefExpr{
	pos: position{line: 734, col: 20, offset: 21254},
	name: "EqAliasG",
},
},
},
&ruleRefExpr{
	pos: position{line: 734, col: 30, offset: 21264},
	name: "_",
},
&litMatcher{
	pos: position{line: 734, col: 32, offset: 21266},
	val: "bar(",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 734, col: 39, offset: 21273},
	name: "_",
},
&labeledExpr{
	pos: position{line: 734, col: 41, offset: 21275},
	label: "xField",
	expr: &ruleRefExpr{
	pos: position{line: 734, col: 48, offset: 21282},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 734, col: 57, offset: 21291},
	name: "_",
},
&litMatcher{
	pos: position{line: 734, col: 59, offset: 21293},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 734, col: 63, offset: 21297},
	name: "_",
},
&labeledExpr{
	pos: position{line: 734, col: 65, offset: 21299},
	label: "yField",
	expr: &ruleRefExpr{
	pos: position{line: 734, col: 72, offset: 21306},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 734, col: 81, offset: 21315},
	name: "_",
},
&labeledExpr{
	pos: position{line: 734, col: 83, offset: 21317},
	label: "args",
	expr: &zeroOrOneExpr{
	pos: position{line: 734, col: 88, offset: 21322},
	expr: &seqExpr{
	pos: position{line: 734, col: 89, offset: 21323},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 734, col: 89, offset: 21323},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 734, col: 93, offset: 21327},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 734, col: 95, offset: 21329},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 734, col: 102, offset: 21336},
	name: "_",
},
&zeroOrOneExpr{
	pos: position{line: 734, col: 104, offset: 21338},
	expr: &seqExpr{
	pos: position{line: 734, col: 105, offset: 21339},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 734, col: 105, offset: 21339},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 734, col: 109, offset: 21343},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 734, col: 111, offset: 21345},
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
	pos: position{line: 734, col: 122, offset: 21356},
	name: "_",
},
&litMatcher{
	pos: position{line: 734, col: 124, offset: 21358},
	val: ")",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "FileName",
	pos: position{line: 767, col: 1, offset: 22043},
	expr: &actionExpr{
	pos: position{line: 767, col: 13, offset: 22055},
	run: (*parser).callonFileName1,
	expr: &seqExpr{
	pos: position{line: 767, col: 13, offset: 22055},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 767, col: 13, offset: 22055},
	name: "Variable",
},
&zeroOrMoreExpr{
	pos: position{line: 767, col: 22, offset: 22064},
	expr: &seqExpr{
	pos: position{line: 767, col: 23, offset: 22065},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 767, col: 23, offset: 22065},
	val: ".",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 767, col: 27, offset: 22069},
	label: "ext",
	expr: &ruleRefExpr{
	pos: position{line: 767, col: 31, offset: 22073},
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
	pos: position{line: 772, col: 1, offset: 22172},
	expr: &actionExpr{
	pos: position{line: 772, col: 10, offset: 22181},
	run: (*parser).callonMExpr1,
	expr: &seqExpr{
	pos: position{line: 772, col: 10, offset: 22181},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 772, col: 10, offset: 22181},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 772, col: 16, offset: 22187},
	name: "MValue",
},
},
&labeledExpr{
	pos: position{line: 772, col: 23, offset: 22194},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 772, col: 28, offset: 22199},
	expr: &seqExpr{
	pos: position{line: 772, col: 29, offset: 22200},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 772, col: 29, offset: 22200},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 772, col: 31, offset: 22202},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 772, col: 36, offset: 22207},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 772, col: 38, offset: 22209},
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
	pos: position{line: 776, col: 1, offset: 22274},
	expr: &actionExpr{
	pos: position{line: 776, col: 11, offset: 22284},
	run: (*parser).callonMValue1,
	expr: &labeledExpr{
	pos: position{line: 776, col: 11, offset: 22284},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 776, col: 16, offset: 22289},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 776, col: 16, offset: 22289},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 776, col: 24, offset: 22297},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 776, col: 33, offset: 22306},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 776, col: 48, offset: 22321},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 776, col: 59, offset: 22332},
	name: "MFactor",
},
	},
},
},
},
},
{
	name: "MOps",
	pos: position{line: 780, col: 1, offset: 22374},
	expr: &choiceExpr{
	pos: position{line: 780, col: 9, offset: 22382},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 780, col: 9, offset: 22382},
	val: "%",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 780, col: 15, offset: 22388},
	val: "-",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 780, col: 21, offset: 22394},
	val: "+",
	ignoreCase: false,
},
&litMatcher{
	pos: position{line: 780, col: 27, offset: 22400},
	val: "*",
	ignoreCase: false,
},
&actionExpr{
	pos: position{line: 780, col: 33, offset: 22406},
	run: (*parser).callonMOps6,
	expr: &litMatcher{
	pos: position{line: 780, col: 33, offset: 22406},
	val: "/",
	ignoreCase: false,
},
},
	},
},
},
{
	name: "MFactor",
	pos: position{line: 784, col: 1, offset: 22454},
	expr: &choiceExpr{
	pos: position{line: 784, col: 12, offset: 22465},
	alternatives: []interface{}{
&actionExpr{
	pos: position{line: 784, col: 12, offset: 22465},
	run: (*parser).callonMFactor2,
	expr: &seqExpr{
	pos: position{line: 784, col: 12, offset: 22465},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 784, col: 12, offset: 22465},
	val: "(",
	ignoreCase: false,
},
&labeledExpr{
	pos: position{line: 784, col: 16, offset: 22469},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 784, col: 19, offset: 22472},
	name: "MExpr",
},
},
&litMatcher{
	pos: position{line: 784, col: 25, offset: 22478},
	val: ")",
	ignoreCase: false,
},
	},
},
},
&actionExpr{
	pos: position{line: 788, col: 5, offset: 22554},
	run: (*parser).callonMFactor8,
	expr: &labeledExpr{
	pos: position{line: 788, col: 5, offset: 22554},
	label: "ex",
	expr: &ruleRefExpr{
	pos: position{line: 788, col: 8, offset: 22557},
	name: "MExpr",
},
},
},
	},
},
},
{
	name: "Expr",
	pos: position{line: 795, col: 1, offset: 22634},
	expr: &actionExpr{
	pos: position{line: 795, col: 9, offset: 22642},
	run: (*parser).callonExpr1,
	expr: &seqExpr{
	pos: position{line: 795, col: 9, offset: 22642},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 795, col: 9, offset: 22642},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 795, col: 15, offset: 22648},
	name: "Value",
},
},
&labeledExpr{
	pos: position{line: 795, col: 21, offset: 22654},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 795, col: 26, offset: 22659},
	expr: &seqExpr{
	pos: position{line: 795, col: 27, offset: 22660},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 795, col: 27, offset: 22660},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 795, col: 29, offset: 22662},
	name: "MOps",
},
&ruleRefExpr{
	pos: position{line: 795, col: 34, offset: 22667},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 795, col: 36, offset: 22669},
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
	pos: position{line: 799, col: 1, offset: 22733},
	expr: &actionExpr{
	pos: position{line: 799, col: 10, offset: 22742},
	run: (*parser).callonValue1,
	expr: &labeledExpr{
	pos: position{line: 799, col: 10, offset: 22742},
	label: "val",
	expr: &choiceExpr{
	pos: position{line: 799, col: 15, offset: 22747},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 799, col: 15, offset: 22747},
	name: "Float",
},
&ruleRefExpr{
	pos: position{line: 799, col: 23, offset: 22755},
	name: "Number",
},
&ruleRefExpr{
	pos: position{line: 799, col: 32, offset: 22764},
	name: "String_Type1",
},
&ruleRefExpr{
	pos: position{line: 799, col: 47, offset: 22779},
	name: "Variable",
},
	},
},
},
},
},
{
	name: "Variable",
	pos: position{line: 803, col: 1, offset: 22822},
	expr: &actionExpr{
	pos: position{line: 803, col: 13, offset: 22834},
	run: (*parser).callonVariable1,
	expr: &seqExpr{
	pos: position{line: 803, col: 13, offset: 22834},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 803, col: 13, offset: 22834},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 803, col: 20, offset: 22841},
	expr: &choiceExpr{
	pos: position{line: 803, col: 21, offset: 22842},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 803, col: 21, offset: 22842},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 803, col: 31, offset: 22852},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 803, col: 39, offset: 22860},
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
	pos: position{line: 807, col: 1, offset: 22910},
	expr: &actionExpr{
	pos: position{line: 807, col: 17, offset: 22926},
	run: (*parser).callonString_Type11,
	expr: &seqExpr{
	pos: position{line: 807, col: 17, offset: 22926},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 807, col: 17, offset: 22926},
	val: "\"",
	ignoreCase: false,
},
&zeroOrMoreExpr{
	pos: position{line: 807, col: 21, offset: 22930},
	expr: &choiceExpr{
	pos: position{line: 807, col: 23, offset: 22932},
	alternatives: []interface{}{
&seqExpr{
	pos: position{line: 807, col: 23, offset: 22932},
	exprs: []interface{}{
&notExpr{
	pos: position{line: 807, col: 23, offset: 22932},
	expr: &ruleRefExpr{
	pos: position{line: 807, col: 24, offset: 22933},
	name: "EscapedChar",
},
},
&anyMatcher{
	line: 807, col: 36, offset: 22945,
},
	},
},
&seqExpr{
	pos: position{line: 807, col: 40, offset: 22949},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 807, col: 40, offset: 22949},
	val: "\\",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 807, col: 45, offset: 22954},
	name: "EscapeSequence",
},
	},
},
	},
},
},
&litMatcher{
	pos: position{line: 807, col: 63, offset: 22972},
	val: "\"",
	ignoreCase: false,
},
	},
},
},
},
{
	name: "Ip",
	pos: position{line: 814, col: 1, offset: 23153},
	expr: &actionExpr{
	pos: position{line: 814, col: 7, offset: 23159},
	run: (*parser).callonIp1,
	expr: &oneOrMoreExpr{
	pos: position{line: 814, col: 7, offset: 23159},
	expr: &choiceExpr{
	pos: position{line: 814, col: 8, offset: 23160},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 814, col: 8, offset: 23160},
	name: "DecimalDigit",
},
&ruleRefExpr{
	pos: position{line: 814, col: 21, offset: 23173},
	name: "SpecialChar",
},
	},
},
},
},
},
{
	name: "StreamName",
	pos: position{line: 818, col: 1, offset: 23221},
	expr: &actionExpr{
	pos: position{line: 818, col: 13, offset: 23233},
	run: (*parser).callonStreamName1,
	expr: &oneOrMoreExpr{
	pos: position{line: 818, col: 13, offset: 23233},
	expr: &choiceExpr{
	pos: position{line: 818, col: 14, offset: 23234},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 818, col: 14, offset: 23234},
	name: "Word",
},
&ruleRefExpr{
	pos: position{line: 818, col: 20, offset: 23240},
	name: "SpecialChar",
},
	},
},
},
},
},
{
	name: "StreamSubject",
	pos: position{line: 822, col: 1, offset: 23288},
	expr: &actionExpr{
	pos: position{line: 822, col: 18, offset: 23305},
	run: (*parser).callonStreamSubject1,
	expr: &oneOrMoreExpr{
	pos: position{line: 822, col: 18, offset: 23305},
	expr: &choiceExpr{
	pos: position{line: 822, col: 19, offset: 23306},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 822, col: 19, offset: 23306},
	name: "Word",
},
&ruleRefExpr{
	pos: position{line: 822, col: 26, offset: 23313},
	name: "SpecialChar",
},
	},
},
},
},
},
{
	name: "SpecialChar",
	pos: position{line: 825, col: 1, offset: 23359},
	expr: &charClassMatcher{
	pos: position{line: 825, col: 15, offset: 23373},
	val: "[-._:]",
	chars: []rune{'-','.','_',':',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapedChar",
	pos: position{line: 827, col: 1, offset: 23383},
	expr: &charClassMatcher{
	pos: position{line: 827, col: 16, offset: 23398},
	val: "[\\x00-\\x1f\"\\\\]",
	chars: []rune{'"','\\',},
	ranges: []rune{'\x00','\x1f',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "EscapeSequence",
	pos: position{line: 829, col: 1, offset: 23416},
	expr: &ruleRefExpr{
	pos: position{line: 829, col: 19, offset: 23434},
	name: "SingleCharEscape",
},
},
{
	name: "SingleCharEscape",
	pos: position{line: 831, col: 1, offset: 23454},
	expr: &charClassMatcher{
	pos: position{line: 831, col: 21, offset: 23474},
	val: "[\"\\\\/bfnrt]",
	chars: []rune{'"','\\','/','b','f','n','r','t',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "Word",
	pos: position{line: 833, col: 1, offset: 23489},
	expr: &actionExpr{
	pos: position{line: 833, col: 9, offset: 23497},
	run: (*parser).callonWord1,
	expr: &oneOrMoreExpr{
	pos: position{line: 833, col: 9, offset: 23497},
	expr: &choiceExpr{
	pos: position{line: 833, col: 10, offset: 23498},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 833, col: 10, offset: 23498},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 833, col: 19, offset: 23507},
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
	pos: position{line: 837, col: 1, offset: 23557},
	expr: &choiceExpr{
	pos: position{line: 837, col: 11, offset: 23567},
	alternatives: []interface{}{
&charClassMatcher{
	pos: position{line: 837, col: 11, offset: 23567},
	val: "[a-z]",
	ranges: []rune{'a','z',},
	ignoreCase: false,
	inverted: false,
},
&charClassMatcher{
	pos: position{line: 837, col: 19, offset: 23575},
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
	pos: position{line: 839, col: 1, offset: 23584},
	expr: &actionExpr{
	pos: position{line: 839, col: 11, offset: 23594},
	run: (*parser).callonNumber1,
	expr: &seqExpr{
	pos: position{line: 839, col: 11, offset: 23594},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 839, col: 11, offset: 23594},
	expr: &litMatcher{
	pos: position{line: 839, col: 11, offset: 23594},
	val: "-",
	ignoreCase: false,
},
},
&ruleRefExpr{
	pos: position{line: 839, col: 16, offset: 23599},
	name: "Integer",
},
	},
},
},
},
{
	name: "PositiveNumber",
	pos: position{line: 843, col: 1, offset: 23672},
	expr: &actionExpr{
	pos: position{line: 843, col: 19, offset: 23690},
	run: (*parser).callonPositiveNumber1,
	expr: &ruleRefExpr{
	pos: position{line: 843, col: 19, offset: 23690},
	name: "Integer",
},
},
},
{
	name: "Float",
	pos: position{line: 847, col: 1, offset: 23763},
	expr: &actionExpr{
	pos: position{line: 847, col: 10, offset: 23772},
	run: (*parser).callonFloat1,
	expr: &seqExpr{
	pos: position{line: 847, col: 10, offset: 23772},
	exprs: []interface{}{
&zeroOrOneExpr{
	pos: position{line: 847, col: 10, offset: 23772},
	expr: &litMatcher{
	pos: position{line: 847, col: 10, offset: 23772},
	val: "-",
	ignoreCase: false,
},
},
&zeroOrOneExpr{
	pos: position{line: 847, col: 15, offset: 23777},
	expr: &ruleRefExpr{
	pos: position{line: 847, col: 15, offset: 23777},
	name: "Integer",
},
},
&litMatcher{
	pos: position{line: 847, col: 24, offset: 23786},
	val: ".",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 847, col: 28, offset: 23790},
	name: "Integer",
},
	},
},
},
},
{
	name: "Integer",
	pos: position{line: 851, col: 1, offset: 23857},
	expr: &choiceExpr{
	pos: position{line: 851, col: 12, offset: 23868},
	alternatives: []interface{}{
&litMatcher{
	pos: position{line: 851, col: 12, offset: 23868},
	val: "0",
	ignoreCase: false,
},
&seqExpr{
	pos: position{line: 851, col: 18, offset: 23874},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 851, col: 18, offset: 23874},
	name: "NonZeroDecimalDigit",
},
&zeroOrMoreExpr{
	pos: position{line: 851, col: 38, offset: 23894},
	expr: &ruleRefExpr{
	pos: position{line: 851, col: 38, offset: 23894},
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
	pos: position{line: 853, col: 1, offset: 23911},
	expr: &charClassMatcher{
	pos: position{line: 853, col: 17, offset: 23927},
	val: "[0-9]",
	ranges: []rune{'0','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "NonZeroDecimalDigit",
	pos: position{line: 855, col: 1, offset: 23936},
	expr: &charClassMatcher{
	pos: position{line: 855, col: 24, offset: 23959},
	val: "[1-9]",
	ranges: []rune{'1','9',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "_",
	displayName: "\"whitespace\"",
	pos: position{line: 857, col: 1, offset: 23968},
	expr: &zeroOrMoreExpr{
	pos: position{line: 857, col: 19, offset: 23986},
	expr: &charClassMatcher{
	pos: position{line: 857, col: 19, offset: 23986},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
},
{
	name: "OneWhiteSpace",
	pos: position{line: 859, col: 1, offset: 24000},
	expr: &charClassMatcher{
	pos: position{line: 859, col: 18, offset: 24017},
	val: "[ \\t\\r\\n]",
	chars: []rune{' ','\t','\r','\n',},
	ignoreCase: false,
	inverted: false,
},
},
{
	name: "IDENTIFIER",
	pos: position{line: 860, col: 1, offset: 24028},
	expr: &actionExpr{
	pos: position{line: 860, col: 15, offset: 24042},
	run: (*parser).callonIDENTIFIER1,
	expr: &seqExpr{
	pos: position{line: 860, col: 15, offset: 24042},
	exprs: []interface{}{
&ruleRefExpr{
	pos: position{line: 860, col: 15, offset: 24042},
	name: "Letter",
},
&zeroOrMoreExpr{
	pos: position{line: 860, col: 22, offset: 24049},
	expr: &choiceExpr{
	pos: position{line: 860, col: 23, offset: 24050},
	alternatives: []interface{}{
&ruleRefExpr{
	pos: position{line: 860, col: 23, offset: 24050},
	name: "Integer",
},
&ruleRefExpr{
	pos: position{line: 860, col: 33, offset: 24060},
	name: "Letter",
},
&litMatcher{
	pos: position{line: 860, col: 41, offset: 24068},
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
	pos: position{line: 863, col: 1, offset: 24141},
	expr: &notExpr{
	pos: position{line: 863, col: 8, offset: 24148},
	expr: &anyMatcher{
	line: 863, col: 9, offset: 24149,
},
},
},
{
	name: "TOK_SEMICOLON",
	pos: position{line: 865, col: 1, offset: 24154},
	expr: &litMatcher{
	pos: position{line: 865, col: 17, offset: 24170},
	val: ";",
	ignoreCase: false,
},
},
{
	name: "VarParamListG",
	pos: position{line: 871, col: 1, offset: 24257},
	expr: &actionExpr{
	pos: position{line: 871, col: 18, offset: 24274},
	run: (*parser).callonVarParamListG1,
	expr: &seqExpr{
	pos: position{line: 871, col: 18, offset: 24274},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 871, col: 18, offset: 24274},
	label: "first",
	expr: &ruleRefExpr{
	pos: position{line: 871, col: 24, offset: 24280},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 871, col: 33, offset: 24289},
	name: "_",
},
&labeledExpr{
	pos: position{line: 871, col: 35, offset: 24291},
	label: "rest",
	expr: &zeroOrMoreExpr{
	pos: position{line: 871, col: 40, offset: 24296},
	expr: &seqExpr{
	pos: position{line: 871, col: 41, offset: 24297},
	exprs: []interface{}{
&litMatcher{
	pos: position{line: 871, col: 41, offset: 24297},
	val: ",",
	ignoreCase: false,
},
&ruleRefExpr{
	pos: position{line: 871, col: 45, offset: 24301},
	name: "_",
},
&ruleRefExpr{
	pos: position{line: 871, col: 47, offset: 24303},
	name: "Variable",
},
&ruleRefExpr{
	pos: position{line: 871, col: 56, offset: 24312},
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
	pos: position{line: 874, col: 1, offset: 24365},
	expr: &actionExpr{
	pos: position{line: 874, col: 13, offset: 24377},
	run: (*parser).callonEqAliasG1,
	expr: &seqExpr{
	pos: position{line: 874, col: 13, offset: 24377},
	exprs: []interface{}{
&labeledExpr{
	pos: position{line: 874, col: 13, offset: 24377},
	label: "alias",
	expr: &ruleRefExpr{
	pos: position{line: 874, col: 19, offset: 24383},
	name: "Variable",
},
},
&ruleRefExpr{
	pos: position{line: 874, col: 28, offset: 24392},
	name: "_",
},
&litMatcher{
	pos: position{line: 874, col: 30, offset: 24394},
	val: "=",
	ignoreCase: false,
},
	},
},
},
},
	},
}
func (c *current) onCommand1(first, second, last interface{}) (interface{}, error) {

    commands := []interface{}{(first.([]interface{}))}
    for _,val := range(cast.ToIfaceSlice(second)){
        for _,statement := range cast.ToIfaceSlice(val){
            castedStatement := cast.ToIfaceSlice(statement)
            if len(castedStatement)>0{
                commands = append(commands,castedStatement)
            }
        }
    }
    return commands,nil
}

func (p *parser) callonCommand1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCommand1(stack["first"], stack["second"], stack["last"])
}

func (c *current) onStatement1(source, transform, sink, last interface{}) (interface{}, error) {

     stages := []interface{}{source}
     transformStages := cast.ToIfaceSlice(transform)
     for _,transformStage := range transformStages{
        castedTransformStage := cast.ToIfaceSlice(transformStage)
        stages = append(stages,castedTransformStage[3])
     }
     castedSinkStage := cast.ToIfaceSlice(sink)
     stages = append(stages,castedSinkStage[3])
     return stages,nil
}

func (p *parser) callonStatement1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatement1(stack["source"], stack["transform"], stack["sink"], stack["last"])
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

func (c *current) onLiftBridgeReaderG1(params interface{}) (interface{}, error) {


    return SourceJob{Type:LIFTREADER,OperateOn:LiftReadJob{params:params.(LiftReader)}},nil
}

func (p *parser) callonLiftBridgeReaderG1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLiftBridgeReaderG1(stack["params"])
}

func (c *current) onLiftReaderParams1(ip, streamName, streamSubject interface{}) (interface{}, error) {

    return LiftReader{Ip:ip.(string),StreamName:streamName.(string),StreamSubject:streamSubject.(string)},nil
}

func (p *parser) callonLiftReaderParams1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLiftReaderParams1(stack["ip"], stack["streamName"], stack["streamSubject"])
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

func (c *current) onIp1() (interface{}, error) {

return string(c.text),nil
}

func (p *parser) callonIp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIp1()
}

func (c *current) onStreamName1() (interface{}, error) {

return string(c.text),nil
}

func (p *parser) callonStreamName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStreamName1()
}

func (c *current) onStreamSubject1() (interface{}, error) {

return string(c.text),nil
}

func (p *parser) callonStreamSubject1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStreamSubject1()
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

