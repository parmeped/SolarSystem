package position

import (
	"fmt"
	m "math"
)

type Triangle struct {
	AngleA    float32 // opposite to BC
	AngleB    float32 // opposite to AC
	AngleC    float32 // opposite to AB
	LengthAB  float32
	LengthBC  float32
	LengthAC  float32
	Perimeter float32
}

// TODO: Check method.
func CalculatePerimeter(t *Triangle) float32 {
	t.Perimeter = t.LengthAB + t.LengthAC + t.LengthBC
	return t.Perimeter
}

// returns a complete triangle knowing only two sides and one angle.
func TwoSidesOneAngleTStrategy(sideA, sideB float32, angle float32) {
	var triangle = Triangle{
		LengthAB: sideA,
		LengthAC: sideB,
		AngleA:   angle,
	}

	triangle.LengthBC = calculateCosin(sideA, sideB, angle)
	triangle.AngleB = calculateSin()

}

// Simple check to know if a triangle is ok or is all bollocks.
func checkTriangleAngles(t *Triangle) bool {
	if t.AngleA+t.AngleB+t.AngleC != 180 {
		fmt.Printf("Error checking triangle! AngleA: %v, AngleB: %v, AngleC: %v \n", t.AngleA, t.AngleB, t.AngleC)
		return false
	} else {
		return true
	}
}

// Cosin? (Teorema del coseno) this is for calculate the remaining side on a triangle where two sides and an angle is known
func calculateCosin(sideA, sideB, angle float32) float32 {
	// horrible conversion because apparently math pkg works only with float64
	_sideA, _sideB := float64(sideA), float64(sideB)
	_angle := float64(angle)
	return float32(m.Sqrt(m.Pow(_sideA, 2) + m.Pow(_sideB, 2) - ((2 * _sideA * _sideB) * m.Cos(_angle))))
}

func calculateSin() float32 {
	return 0
}
