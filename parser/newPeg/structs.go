package newPeg

import (
	"time"
	"github.com/Knetic/govaluate"
	"github.com/raralabs/canal/core/transforms/event/poll"
)

var TrueFilter = func(map[string]interface{}) (bool, error) {
	return true, nil
}
var FalseFilter = func(map[string]interface{}) (bool, error) {
	return false, nil
}

type FileJob struct {
	Filename string
}

type FakeJob struct {
	NumData int64
}

type DoNodeJob struct {
	Function interface{}
}

type AggNodeJob struct {
	Trigger   poll.Event
	Functions []interface{}
	GroupBy   []string
}

type JoinNodeJob struct{
	Alias string
	Type string
	Attributes JoinAttributes
	JoinSubType string

}

type LiftReadJob struct{
	params LiftReader
}

//type SinkJob struct {
//	Type string
//	Args interface{}
//	Header []string
//}

// Sink Jobs
type Into struct {
	StreamTo string
}

type Plot struct {
	Type   string
	Widget interface{}
}

type BarPlot struct {
	Title    string
	XField   string
	YField   string
	BarWidth int
	BarGap   int
}

// Do Functions
//type Filter func(map[string]interface{}) (bool, error)

type BranchJob struct{
	BranchName string
}
type Branch struct {
	Condition string
}

type Delay struct {
	Time time.Duration
}

type Pass struct {
}

type Select struct {
	Fields []string
}

type Pick struct {
	Desc string
	Num  uint64
}

type Sort struct {
	Field string
}

type Batch struct {
	Num int64
}

type Enrich struct {
	Field string
	Expr  *govaluate.EvaluableExpression
}

// Aggregator Functions

//type Count struct {
//	Alias  string
//	Filter Filter
//}

type Max struct {
	Alias  string
	Field  string
	Filter Filter
}

type Min struct {
	Alias  string
	Field  string
	Filter Filter
}

type Avg struct {
	Alias  string
	Field  string
	Filter Filter
}

type Sum struct {
	Alias  string
	Field  string
	Filter Filter
}

type Variance struct {
	Alias  string
	Field  string
	Filter Filter
}

type DistinctCount struct {
	Alias  string
	Field  string
	Filter Filter
}

type Mode struct {
	Alias  string
	Field  string
	Filter Filter
}

type Quantile struct {
	Alias  string
	Field  string
	Filter Filter
	Weight string
	Qth    float64
}

type Covariance struct {
	Alias  string
	Field1 string
	Field2 string
	Filter Filter
}

type Correlation struct {
	Alias  string
	Field1 string
	Field2 string
	Filter Filter
}

type PValue struct {
	Alias  string
	Field1 string
	Field2 string
	Filter Filter
}