package peg

import (
	"errors"

	"github.com/raralabs/canal/utils/cast"
)

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
