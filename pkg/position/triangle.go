package position

import (
	"fmt"
	m "math"
)

type Triangle struct {
	AngleA    float32
	AngleB    float32
	AngleC    float32
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

	// horrible conversion because apparently math pkg works only with float64
	_sideA, _sideB := float64(sideA), float64(sideB)
	_angle := float64(angle)
	// Cosino theorem (Teorema del coseno)
	triangle.LengthBC = float32(m.Sqrt(m.Pow(_sideA, 2) + m.Pow(_sideB, 2) - ((2 * _sideA * _sideB) * m.Cos(_angle))))

	triangle.AngleB

}

func checkTriangleAngles(t *Triangle) bool {
	if t.AngleA+t.AngleB+t.AngleC != 180 {
		fmt.Printf("Error checking triangle! AngleA: %v, AngleB: %v, AngleC: %v \n", t.AngleA, t.AngleB, t.AngleC)
		return false
	} else {
		return true
	}
}
