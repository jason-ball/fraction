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

// Computes the greatest common divisor between two
// integers.
// Algorithm from https://en.wikipedia.org/wiki/Euclidean_algorithm#Implementations.
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a % b)
}
