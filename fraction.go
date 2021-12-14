package fraction

import (
	"fmt"
)

type Fraction struct {
	Numerator int
	Denominator int
}

func (f *Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}

func Add(a, b Fraction) Fraction {
	n := (a.Numerator * b.Denominator) + (a.Denominator * b.Numerator)
	d := a.Denominator * b.Denominator

	div := GCD(n ,d)
	n /= div
	d /= div

	return Fraction{n, d}
}

func Sub(a, b Fraction) Fraction {
	n := (a.Numerator * b.Denominator) - (a.Denominator * b.Numerator)
	d := a.Denominator * b.Denominator

	div := GCD(n ,d)
	n /= div
	d /= div

	return Fraction{n, d}
}

func Mul(a, b Fraction) Fraction {
	n := a.Numerator * b.Numerator
	d := a.Denominator * b.Denominator

	div := GCD(n ,d)
	n /= div
	d /= div

	return Fraction{n, d}
}

func Div(a, b Fraction) Fraction {
	n := a.Numerator * b.Denominator
	d := a.Denominator * b.Numerator

	div := GCD(n ,d)
	n /= div
	d /= div

	return Fraction{n, d}
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a % b)
}