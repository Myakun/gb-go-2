package numbers_test

import (
	"lesson4/numbers"
	"testing"
)

func TestIncrementInvalidInput(t *testing.T) {
	if _, err := numbers.Increment(1, -1); nil == err {
		t.Error("Incorrect return result. Expected error, got nil")
	}
}

func TestIncrementSuccess(t *testing.T) {
	incBy := 1000
	initialValues := [...]int{0, 13, 15, 29, 100, 567, 1045}
	for _, initialValue := range initialValues {
		initialValue = 0
		expected := initialValue + incBy
		got, err := numbers.Increment(initialValue, incBy)
		if nil != err {
			t.Errorf("Incorrect return result. Expected nil, got error %s", err)
		}

		if expected != got {
			t.Errorf("Incorrect result. Expected %d, got %d", expected, got)
		}
	}
}
