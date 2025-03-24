package utils_test

import (
	"asteroid/utils"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, utils.Max(1, 2), "int: The max of 1 and 2 should be 2")
	assert.Equal(int8(2), utils.Max(int8(1), int8(2)), "int8: The max of 1 and 2 should be 2")
	assert.Equal(int16(2), utils.Max(int16(1), int16(2)), "int16: The max of 1 and 2 should be 2")
	assert.Equal(int32(2), utils.Max(int32(1), int32(2)), "int32: The max of 1 and 2 should be 2")
	assert.Equal(int64(2), utils.Max(int64(1), int64(2)), "int64: The max of 1 and 2 should be 2")
	assert.Equal(uint(2), utils.Max(uint(1), uint(2)), "uint: The max of 1 and 2 should be 2")
	assert.Equal(uint8(2), utils.Max(uint8(1), uint8(2)), "uint8: The max of 1 and 2 should be 2")
	assert.Equal(uint16(2), utils.Max(uint16(1), uint16(2)), "uint16: The max of 1 and 2 should be 2")
	assert.Equal(uint32(2), utils.Max(uint32(1), uint32(2)), "uint32: The max of 1 and 2 should be 2")
	assert.Equal(uint64(2), utils.Max(uint64(1), uint64(2)), "uint64: The max of 1 and 2 should be 2")
	assert.Equal(float32(2), utils.Max(float32(1), float32(2)), "float32: The max of 1 and 2 should be 2")
	assert.Equal(float64(2), utils.Max(float64(1), float64(2)), "float64: The max of 1 and 2 should be 2")
}

func TestMin(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, utils.Min(1, 2), "int: The min of 1 and 2 should be 1")
	assert.Equal(int8(1), utils.Min(int8(1), int8(2)), "int8: The min of 1 and 2 should be 1")
	assert.Equal(int16(1), utils.Min(int16(1), int16(2)), "int16: The min of 1 and 2 should be 1")
	assert.Equal(int32(1), utils.Min(int32(1), int32(2)), "int32: The min of 1 and 2 should be 1")
	assert.Equal(int64(1), utils.Min(int64(1), int64(2)), "int64: The min of 1 and 2 should be 1")
	assert.Equal(uint(1), utils.Min(uint(1), uint(2)), "uint: The min of 1 and 2 should be 1")
	assert.Equal(uint8(1), utils.Min(uint8(1), uint8(2)), "uint8: The min of 1 and 2 should be 1")
	assert.Equal(uint16(1), utils.Min(uint16(1), uint16(2)), "uint16: The min of 1 and 2 should be 1")
	assert.Equal(uint32(1), utils.Min(uint32(1), uint32(2)), "uint32: The min of 1 and 2 should be 1")
	assert.Equal(uint64(1), utils.Min(uint64(1), uint64(2)), "uint64: The min of 1 and 2 should be 1")
	assert.Equal(float32(1), utils.Min(float32(1), float32(2)), "float32: The min of 1 and 2 should be 1")
	assert.Equal(float64(1), utils.Min(float64(1), float64(2)), "float64: The min of 1 and 2 should be 1")
}

func TestClamp(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, utils.Clamp(0, 1, 10), "int: The clamp[1, 10] of 0 should be 1")
	assert.Equal(5, utils.Clamp(5, 1, 10), "int: The clamp[1, 10] of 5 should be 5")
	assert.Equal(10, utils.Clamp(15, 1, 10), "int: The clamp[1, 10] of 15 should be 10")
	assert.Equal(1.0, utils.Clamp(0.0, 1.0, 10.0), "float64: The clamp[1, 10] of 0.0 should be 1.0")
	assert.Equal(5.0, utils.Clamp(5.0, 1.0, 10.0), "float64: The clamp[1, 10] of 5.0 should be 5.0")
	assert.Equal(10.0, utils.Clamp(15.0, 1.0, 10.0), "float64: The clamp[1, 10] of 15.0 should be 10.0")
}

func TestDistance(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		name           string
		x1, y1, x2, y2 int
		want           float64
	}{
		{"(0, 0) to (0, 0)", 0, 0, 0, 0, 0},
		{"(0, 0) to (3, 0)", 0, 0, 3, 0, 3},
		{"(0, 0) to (0, 4)", 0, 0, 0, 4, 4},
		{"(0, 0) to (3, 4)", 0, 0, 3, 4, 5},
		{"(0, 0) to (-3, 0)", 0, 0, -3, 0, 3},
		{"(0, 0) to (0, -4)", 0, 0, 0, -4, 4},
		{"(0, 0) to (-3, -4)", 0, 0, -3, -4, 5},
		{"(0, 0) to (-3, 4)", 0, 0, -3, 4, 5},
		{"(0, 0) to (3, -4)", 0, 0, 3, -4, 5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(c.want, utils.Distance(c.x1, c.y1, c.x2, c.y2), "The distance between (%d, %d) and (%d, %d) should be %f", c.x1, c.y1, c.x2, c.y2, c.want)
		})
	}
}
