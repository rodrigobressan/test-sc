package test_sc

import "testing"

func Test_sum(t *testing.T) {

	tests := []struct {
		name string
		a    int
		b    int
		sum  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.a, tt.b); got != tt.sum {
				t.Errorf("sum() = %v, want %v", got, tt.sum)
			}
		})
	}
}
