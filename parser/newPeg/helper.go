package newPeg

import (
	"errors"
	"fmt"
	"strings"
	"github.com/raralabs/canal/utils/cast"
)

var posMap = map[int]string {
	1: "st",
	2: "nd",
	3: "rd",
}

func parseAggArgs(alias interface{}, flt interface{}) (string, Filter, error) {
	name := ""
	filter := TrueFilter
	if alias != nil {
		if nm, ok := cast.TryString(alias); ok {
			name = nm
		} else {
			return name, filter, errors.New("could not decode alias")
		}
	}
	if flt != nil {
		if fl, ok := flt.(Filter); ok {
			filter = fl
		} else {
			return name, filter, errors.New("could not decode filter")
		}
	}
	return name, filter, nil
}

func parseMathExpr(first, rest interface{}) string {
	var str strings.Builder

	if f, ok := cast.TryString(first); ok {
		str.WriteString(f)
	} else {
		str.WriteString(fmt.Sprintf("%v", first))
	}

	rst := cast.ToIfaceSlice(rest)
	for _, r := range rst {
		v := cast.ToIfaceSlice(r)

		str.WriteString(" ")
		op, ok := cast.TryString(v[1])
		if ok {
			str.WriteString(op)
		} else {
			str.WriteString(fmt.Sprintf("%v", v[1]))
		}

		str.WriteString(" ")
		rt, ok := cast.TryString(v[3])
		if ok {
			str.WriteString(rt)
		} else {
			str.WriteString(fmt.Sprintf("%v", v[3]))
		}
	}

	return str.String()
}

func getParamList(first, rest interface{}) []string {
	var strs []string

	f, _ := cast.TryString(first)
	strs = append(strs, f)

	restStrs := cast.ToIfaceSlice(rest)
	for _, v := range restStrs {
		str := cast.ToIfaceSlice(v)
		r, _ := cast.TryString(str[2])

		strs = append(strs, r)
	}

	return strs
}
