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

func Test_given_5_then_return_5(t *testing.T) {
	assert.Equal(t, []int{5}, Prime(5))
}
