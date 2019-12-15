package solution

import "testing"

func TestTwoSumSuccess(t *testing.T) {
	execBinaryGap(t, 9, 2)
	execBinaryGap(t, 529, 4)
	execBinaryGap(t, 51272, 4)
	execBinaryGap(t, 15, 0)
	execBinaryGap(t, 32, 0)
	execBinaryGap(t, 0, 0)
	execBinaryGap(t, 6291457, 20)
	execBinaryGap(t, 2147483647, 0)

}

func execBinaryGap(t *testing.T, n int, correct int) {
	result := binaryGap(n)
	if result != correct {
		t.Fatalf("function binaryGap returns incorrect values. n: %v, result: %v, collect: %v", n, result, correct)
	}
}
