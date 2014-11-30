package prime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Given_1_then_return_empty_list(t *testing.T) {
	assert.Equal(t, []int{}, Prime(1))
}

func Test_Given_2_then_return_2(t *testing.T) {
	assert.Equal(t, []int{2}, Prime(2))
}

func Test_Given_3_then_return_3(t *testing.T) {
	assert.Equal(t, []int{3}, Prime(3))
}

func Test_Given_4_then_return_2_2(t *testing.T) {
	assert.Equal(t, []int{2, 2}, Prime(4))
}

func Test_Given_6_then_return_2_3(t *testing.T) {
	assert.Equal(t, []int{2, 3}, Prime(6))
}

func Test_Given_8_then_return_2_2_2(t *testing.T) {
	assert.Equal(t, []int{2, 2, 2}, Prime(8))
}

func Test_Given_5_then_return_5(t *testing.T) {
	assert.Equal(t, []int{5}, Prime(5))
}

func Test_Given_7_then_return_7(t *testing.T) {
	assert.Equal(t, []int{7}, Prime(7))
}

func Test_Given_9_then_return_3_3(t *testing.T) {
	assert.Equal(t, []int{3, 3}, Prime(9))
}

func Test_Given_10_then_return_2_5(t *testing.T) {
	assert.Equal(t, []int{2, 5}, Prime(10))
}

func Test_Given_12_then_return_2_2_3(t *testing.T) {
	assert.Equal(t, []int{2, 2, 3}, Prime(12))
}

func Test_Given_15_then_return_3_5(t *testing.T) {
	assert.Equal(t, []int{3, 5}, Prime(15))
}

func Test_Given_78654_then_return_3_5(t *testing.T) {
	assert.Equal(t, []int{2, 2, 2, 2, 2, 2, 1229}, Prime(78656))
}
