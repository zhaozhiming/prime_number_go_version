package prime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Given_1_then_return_empty_list(t *testing.T) {
	list := Prime(1)
	assert.Equal(t, []int{}, list)
}

func Test_Given_2_then_return_2(t *testing.T) {
	list := Prime(2)
	assert.Equal(t, []int{2}, list)
}
