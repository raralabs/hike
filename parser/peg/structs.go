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
