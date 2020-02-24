package yagrt

import (
	"math"
	"strconv"
	"strings"
)

// ParseFloat parses floats from a given string by extending
// the functionality of strconv.ParseFloat to include scientific
// notation
func ParseFloat(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return val, nil
	}
	str = strings.Replace(str, ",", "", -1)
	pos := strings.IndexAny(str, "eE")
	if pos < 0 {
		return strconv.ParseFloat(str, 64)
	}
	var base float64
	var exp int64
	baseStr := str[0:pos]
	base, err = strconv.ParseFloat(baseStr, 64)
	if err != nil {
		return 0, err
	}
	expStr := str[(pos + 1):]
	exp, err = strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return base * math.Pow10(int(exp)), nil
}

// Determinant is used to take 3x3 matrix determinant
func Determinant(v1, v2, v3 *Vector) float64 {
	return v1.X*((v2.Y*v3.Z)-(v3.Y*v2.Z)) -
		v2.X*((v1.Y*v3.Z)-(v3.Y*v1.Z)) +
		v3.X*((v1.Y*v2.Z)-(v2.Y*v1.Z))
}
