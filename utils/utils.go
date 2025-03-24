package utils

import (
	"cmp"
	"math"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Max[T cmp.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func Min[T cmp.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Clamp[T cmp.Ordered](x, min, max T) T {
	return Min(Max(min, x), max)
}

func Distance[T Numeric](x1, y1, x2, y2 T) float64 {
	dx := x1 - x2
	dy := y1 - y2
	return math.Sqrt(float64(dx*dx + dy*dy))
}
