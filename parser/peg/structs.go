package peg

import (
	"time"

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

type DoNodeJob struct {
	Function interface{}
}

type AggNodeJob struct {
	Trigger   poll.Event
	Functions []interface{}
	GroupBy   []string
}

type SinkJob struct {
	Type string
	Args interface{}
}

// Sink Jobs
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
type Filter func(map[string]interface{}) (bool, error)

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

// Aggregator Functions

type Count struct {
	Alias  string
	Filter Filter
}

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
