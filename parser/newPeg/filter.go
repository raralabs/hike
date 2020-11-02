package newPeg

import (
	"errors"
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/raralabs/canal/utils/cast"
)

// A secondaryOperatorMap holds the mapping of secondary operators like or, and.
var secondaryOperatorMap = map[string]func(l, r bool) (bool, error){
	"or": func(l, r bool) (bool, error) {
		return l || r, nil
	},
	"and": func(l, r bool) (bool, error) {
		return l && r, nil
	},
}

// newPrimaryFilter returns a filter defined by op.
func newPrimaryFilter(name string, op, val interface{}) Filter {

	operator := ""
	switch op := op.(type) {
	case []uint8:
		operator = string(op)
	case string:
		operator = op
	}

	condition := fmt.Sprintf("%s %s %v", name, operator, val)

	expression, err := govaluate.NewEvaluableExpression(condition)
	if err != nil {
		panic(err)
	}

	filter := func(valMap map[string]interface{}) (bool, error) {
		ret := false

		result, err := expression.Evaluate(valMap)
		if err != nil {
			return false, err
		}

		rbool, ok := result.(bool)
		if !ok {
			return false, errors.New("non boolean result in filter condition")
		} else {
			ret = rbool
		}

		return ret, nil
	}

	return filter
}

// newSecondaryFilter returns a filter defined by rest recursively.
func newSecondaryFilter(first, rest interface{}) Filter {

	filter := func(valMap map[string]interface{}) (bool, error) {

		if valMap == nil {
			return false, nil
		}

		l := first.(Filter)
		outl, _ := l(valMap)

		restSl := cast.ToIfaceSlice(rest)
		for _, v := range restSl {
			restExpr := cast.ToIfaceSlice(v)
			r := restExpr[3].(Filter)
			op := restExpr[1]
			operator := ""
			switch op := op.(type) {
			case []uint8:
				operator = string(op)
			case string:
				operator = op
			}

			outr, _ := r(valMap)
			outl, _ = secondaryOperatorMap[operator](outl, outr)
		}
		return outl, nil
	}

	return filter
}
