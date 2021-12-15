package fraction_test

import (
	"github.com/jason-ball/fraction"
	"testing"
)

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