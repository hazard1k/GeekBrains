package statistic

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func TestSum(t *testing.T) {
	var v float64
	v = Sum([]float64{1, 2.1})
	if v != 3.1 {
		t.Error("Expected 3.1, got ", v)
	}
}
