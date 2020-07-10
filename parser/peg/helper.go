package peg

import (
	"errors"

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
