package string_sum

import (
	"fmt"
	"testing"
)

func TestStringSum(t *testing.T) {
	data := []struct {
		In       string
		Expected string
		Err      error
	}{
		{In: "3+5", Expected: "8", Err: nil},
		{In: "-3+5", Expected: "2", Err: nil},
		{In: "-3-5", Expected: "-8", Err: nil},
		{In: "", Expected: "0", Err: fmt.Errorf("something went wrong: %w", errorEmptyInput)},
		{In: " ", Expected: "0", Err: fmt.Errorf("something went wrong: %w", errorEmptyInput)},
		{In: "3", Expected: "0", Err: fmt.Errorf("something went wrong: %w", errorNotTwoOperands)},
		{In: "-3+4+5", Expected: "0", Err: fmt.Errorf("something went wrong: %w", errorNotTwoOperands)},
	}
	for _, q := range data {
		got, err := StringSum(q.In)
		if got != q.Expected && q.Err == err {
			t.Logf("StringSum expected: %s, got: %s", q.Expected, got)
			t.Fail()
		}
	}
}
