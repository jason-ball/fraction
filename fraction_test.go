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

func TestString(t *testing.T) {
	f := fraction.Fraction{1, 2}
	
	if f.String() != "1/2" {
		t.Errorf("f.String = \"%s\"; want \"%s\"", f.String(), "1/2")
	}
}

func TestGCD(t *testing.T) {
	a := 1071
	b := 462
	result := fraction.GCD(a, b)
	expected := 21

	if result != expected {
		t.Errorf("GCD(%d, %d) = %d; want %d", a, b, result, expected)
	}
}
