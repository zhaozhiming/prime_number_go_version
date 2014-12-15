package prime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_given_1_then_return_empty_list(t *testing.T) {
	assert.Equal(t, []int{}, Prime(1))
}

func Test_given_2_then_return_2(t *testing.T) {
	assert.Equal(t, []int{2}, Prime(2))
}

func Test_given_3_then_return_3(t *testing.T) {
	assert.Equal(t, []int{3}, Prime(3))
}

func Test_given_4_then_return_2_2(t *testing.T) {
	assert.Equal(t, []int{2, 2}, Prime(4))
}

func Test_given_6_then_return_2_3(t *testing.T) {
	assert.Equal(t, []int{2, 3}, Prime(6))
}

func Test_given_8_then_return_2_2_2(t *testing.T) {
	assert.Equal(t, []int{2, 2, 2}, Prime(8))
}

func Test_given_9_then_return_3_3(t *testing.T) {
	assert.Equal(t, []int{3, 3}, Prime(9))
}

func Test_given_20_then_return_2_2_5(t *testing.T) {
	assert.Equal(t, []int{2, 2, 5}, Prime(20))
}

func Test_given_30_then_return_2_3_5(t *testing.T) {
	assert.Equal(t, []int{2, 3, 5}, Prime(30))
}

func Test_given_64_then_return_2_2_2_2_2_2(t *testing.T) {
	assert.Equal(t, []int{2, 2, 2, 2, 2, 2}, Prime(64))
}

func Test_given_10984_then_return_2_2_2_1373(t *testing.T) {
	assert.Equal(t, []int{2, 2, 2, 1373}, Prime(10984))
}
