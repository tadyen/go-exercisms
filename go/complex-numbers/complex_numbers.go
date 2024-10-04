package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	a, b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	var n Number
	n.a = n1.Real() + n2.Real()
	n.b = n1.Imaginary() + n2.Imaginary()
	return n
}

func (n1 Number) Subtract(n2 Number) Number {
	var n Number
	n.a = n1.a - n2.a
	n.b = n1.b - n2.b
	return n
}

func (n1 Number) Multiply(n2 Number) Number {
	var n Number
	n.a = n1.a*n2.a - n1.b*n2.b
	n.b = n1.a*n2.b + n2.a*n1.b
	return n
}

func (n Number) Times(factor float64) Number {
	var o Number
	o.a = n.a * factor
	o.b = n.b * factor
	return o
}

func (n1 Number) Divide(n2 Number) Number {
	var o Number
	o = n1
	o = o.Multiply(n2.Conjugate())
	m := n2.Abs()
	o = o.Times(1 / (m * m))
	return o
}

func (n Number) Conjugate() Number {
	var o Number
	o.a = n.a
	o.b = -n.b
	return o
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a + n.b*n.b)
}

func (n Number) Exp() Number {
	var a float64
	var o Number
	a = math.Exp(n.a)
	o.a = math.Cos(n.b)
	o.b = math.Sin(n.b)
	o = o.Times(a)
	return o
}
