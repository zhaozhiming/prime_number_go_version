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
