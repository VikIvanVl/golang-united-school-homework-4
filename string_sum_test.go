package string_sum

import (
	"testing"
)

func TestStringSum(t *testing.T) {
	data := []struct {
		In       string
		Expected string
		Err      error
	}{
		{In: "-70-9", Expected: "-79", Err: nil},
		{In: " 70 + 5 ", Expected: "75", Err: nil},
		{In: "3+5", Expected: "8", Err: nil},
		{In: "-3+5", Expected: "2", Err: nil},
		{In: "", Expected: "", Err: GetErrorEmptyInput()},
		{In: " ", Expected: "", Err: GetErrorEmptyInput()},
		{In: "3", Expected: "", Err: GetErrorNotTwoOperands()},
		{In: "-3+4+5", Expected: "", Err: GetErrorNotTwoOperands()},
	}
	for _, q := range data {
		got, err := StringSum(q.In)
		if got != q.Expected || q.Err != err {
			if got != q.Expected {
				t.Logf("StringSum expected: %s, got: %s", q.Expected, got)
				t.Fail()
			} else {
				if q.Err.Error() != err.Error() {
					t.Logf("StringSum error expected: %s, got: %s", q.Err, err)
					t.Fail()
				}
			}
		}
	}
}
