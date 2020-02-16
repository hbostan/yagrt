package yagrt

import (
	"math"
	"strconv"
	"strings"
)

func ParseFloat(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return val, nil
	}

	//Some number may be seperated by comma, for example, 23,120,123, so remove the comma firstly
	str = strings.Replace(str, ",", "", -1)

	//Some number is specifed in scientific notation
	pos := strings.IndexAny(str, "eE")
	if pos < 0 {
		return strconv.ParseFloat(str, 64)
	}

	var baseVal float64
	var expVal int64

	baseStr := str[0:pos]
	baseVal, err = strconv.ParseFloat(baseStr, 64)
	if err != nil {
		return 0, err
	}

	expStr := str[(pos + 1):]
	expVal, err = strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return baseVal * math.Pow10(int(expVal)), nil
}

func Determinant(v1, v2, v3 *Vector) float64 {
	return v1.X*(v2.Y*v3.Z-v3.Y*v2.Z) + v2.X*(v3.Y*v1.Z-v1.Y*v3.Z) + v3.X*(v1.Y*v2.Z-v2.Y*v1.Z)
}
