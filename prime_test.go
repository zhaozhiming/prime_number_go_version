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
