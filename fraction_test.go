package fraction_test

import (
	"github.com/jason-ball/fraction"
	"testing"
)

func TestFraction(t *testing.T) {
	n := 1
	d := 2

	f := fraction.Fraction{n, d}

	if f.Numerator != n {
		t.Errorf("f.Numerator = %d; want %d", f.Numerator, n)
	}

	if f.Denominator != d {
		t.Errorf("f.Denominator = %d; want %d", f.Denominator, d)
	}
}