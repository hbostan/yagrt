package yagrt

import "math"

// INF represents infinity
var INF float64 = math.Inf(1)

// EPS represents a generic epsilon value
var EPS float64 = 1e-3

// HitEpsilon is used to check ray intersections
var HitEpsilon float64 = 1e-3

// ShadowEpsilon is used to shift intersection points
// to prevent surface acne
var ShadowEpsilon float64 = 1e-3
