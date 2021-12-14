// A simple fraction library
package fraction

import (
	"fmt"
)

// A struct for a Fraction.
// Contains a Numerator and a Denominator.
type Fraction struct {
	Numerator int
	Denominator int
}

// Returns fractions as strings in the format:
// "{Numerator}/{Denominator}".
func (f *Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}

// Adds two fractions and simplifies the result.
func Add(a, b Fraction) Fraction {
	n := (a.Numerator * b.Denominator) + (a.Denominator * b.Numerator)
	d := a.Denominator * b.Denominator

	div := GCD(n ,d)
	n /= div
	d /= div

	return Fraction{n, d}
}

// Subtracts two fractions and simplifies the result.
func Sub(a, b Fraction) Fraction {
	n := (a.Numerator * b.Denominator) - (a.Denominator * b.Numerator)
	d := a.Denominator * b.Denominator

	div := GCD(n, d)
	n /= div
	d /= div

	return Fraction{n, d}
}

// Multiplies two fractions and simplifies the result.
func Mul(a, b Fraction) Fraction {
	n := a.Numerator * b.Numerator
	d := a.Denominator * b.Denominator

	div := GCD(n ,d)
	n /= div
	d /= div

	return Fraction{n, d}
}

// Divides two fractions and simplifies the result.
func Div(a, b Fraction) Fraction {
	n := a.Numerator * b.Denominator
	d := a.Denominator * b.Numerator

	div := GCD(n ,d)
	n /= div
	d /= div

	return Fraction{n, d}
}

// Computes the greatest common divisor between two
// integers.
// Algorithm from https://en.wikipedia.org/wiki/Euclidean_algorithm#Implementations.
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a % b)
}
