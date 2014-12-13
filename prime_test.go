package prime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Given_1_then_return_empty_list(t *testing.T) {
	assert.Equal(t, []int{}, Prime(1))
}
