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


func TestAdd(t *testing.T) {
	a := fraction.Fraction{3, 6}
	b := fraction.Fraction{25, 100}

	r := fraction.Add(a, b)
	e := fraction.Fraction{3, 4}

	if r != e {
		t.Errorf("%s + %s = %s; want %s", a.String(), b.String(), r.String(), e.String())
	}
}

func TestSub(t *testing.T) {
	a := fraction.Fraction{3, 6}
	b := fraction.Fraction{75, 100}

	r := fraction.Sub(b, a)
	e := fraction.Fraction{1, 4}

	if r != e {
		t.Errorf("%s + %s = %s; want %s", a.String(), b.String(), r.String(), e.String())
	}
}

func TestMul(t *testing.T) {
	a := fraction.Fraction{3, 6}
	b := fraction.Fraction{25, 100}

	r := fraction.Mul(a, b)
	e := fraction.Fraction{1, 8}

	if r != e {
		t.Errorf("%s + %s = %s; want %s", a.String(), b.String(), r.String(), e.String())
	}
}

func TestDiv(t *testing.T) {
	a := fraction.Fraction{3, 6}
	b := fraction.Fraction{25, 100}

	r := fraction.Div(a, b)
	e := fraction.Fraction{2, 1}

	if r != e {
		t.Errorf("%s + %s = %s; want %s", a.String(), b.String(), r.String(), e.String())
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