package prime

import (
	"reflect"
	"testing"
)

func Test_Given_1_then_return_empty_list(t *testing.T) {
	list := Prime(1)
	if !reflect.DeepEqual(list, []int{}) {
		t.Errorf("test failed")
	}
}

func Test_Given_2_then_return_2(t *testing.T) {
	list := Prime(2)
	if !reflect.DeepEqual(list, []int{2}) {
		t.Errorf("test failed")
	}
}
