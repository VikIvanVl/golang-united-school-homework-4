package string_sum

import (
	"testing"
)

func Test_stringSum(t *testing.T) {
	data := []struct {
		In       string
		Expected string
		Err      error
	}{
		{In: " 43 + 40 ", Expected: "83", Err: nil},
		{In: "32+45", Expected: "77", Err: nil},
		{In: "-44-44", Expected: "-88", Err: nil},
		{In: "-3+5", Expected: "2", Err: nil},
		{In: "", Expected: "", Err: GetErrorEmptyInput()},
		{In: " ", Expected: "", Err: GetErrorEmptyInput()},
		{In: "-", Expected: "", Err: GetErrorNotTwoOperands()},
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
