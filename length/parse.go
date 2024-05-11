package length

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	LengthPatternMicrometer = "mkm"
	LengthPatternMillimeter = "mm"
	LengthPatternDecimeter  = "dm"
	LengthPatternCantimeter = "cm"
	LengthPatternMeter      = "m"
	LengthPatternKilometer  = "km"
)

var (
	volumeRegExp = regexp.MustCompile("([+-]?([0-9]*[.])?[0-9]+)([mkm|mm|dm|cm|m|km])")

	Zero = Unit(0)
)

func Parse(s string) (Unit, error) {
	if !volumeRegExp.MatchString(s) {
		return Zero, errors.New("invalid volume description, not match")
	}

	match := volumeRegExp.FindStringSubmatch(s)
	if len(match) != 4 {
		return Zero, fmt.Errorf("invalid volume description, parts len must be 4, got: %d", len(match))
	}

	value, _ := strconv.ParseFloat(match[1], 64)

	unit := match[3]

	switch unit {
	case LengthPatternMicrometer:
		return Unit(value) * Micrometer, nil
	case LengthPatternMillimeter:
		return Unit(value) * Millimeter, nil
	case LengthPatternDecimeter:
		return Unit(value) * Decimeter, nil
	case LengthPatternCantimeter:
		return Unit(value) * Cantimeter, nil
	case LengthPatternMeter:
		return Unit(value) * Meter, nil
	case LengthPatternKilometer:
		return Unit(value) * Kilometer, nil
	default:
		return Zero, fmt.Errorf("unknown length unit type: %s", unit)
	}
}
