package fraction

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