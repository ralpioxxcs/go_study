package main

import (
	"fmt"
	"math"
)

func main() {
	for _, input := range inputs {
		triangulate(input)
	}
}

var inputs = []string{"0 0 45\n10 0 0", "0 0 24.0\n11 7 343.4", "10 1 0.0\n2 8 352.82", "0 0 30.9\n10 1 336.42"}

type Point struct{ X, Y, Z int }
type Vector struct{ X, Y, Z float64 }
type Ray struct {
	V Vector
	P Point
	T float64 // theta
}

func (u Vector) Cross(v Vector) Vector {
	return Vector{u.Y*v.Z - u.Z*v.Y, u.X*v.Z - u.Z*v.X, u.X*v.Y - u.Y*v.X}
}
func (u Vector) Mag() float64             { return math.Sqrt(u.X*u.X + u.Y*u.Y + u.Z*u.Z) }
func (u Vector) Divide(v Vector) Vector   { return Vector{u.X / v.X, u.Y / v.Y, u.Z / v.Z} }
func (u Vector) MultMag(n float64) Vector { return Vector{u.X * n, u.Y * n, u.Z * n} }

func (u Point) Add(v Vector) Vector {
	return Vector{float64(u.X) + v.X, float64(u.Y) + v.Y, float64(u.Z) + v.Z}
}
func (u Point) Subtract(v Point) Vector {
	return Vector{float64(u.X - v.X), float64(u.Y - v.Y), float64(u.Z - v.Z)}
}

func triangulate(input string) {
	var r1, r2 Ray
	fmt.Sscanf(input, "%d %d %f\n%d %d %f", &r1.P.X, &r1.P.Y, &r1.T, &r2.P.X, &r2.P.Y, &r2.T)

	// get unit vectors v.x and v.y using the provided angle theta
	r1.V = Vector{math.Sin(r1.T * math.Pi / 180), math.Cos(r1.T * math.Pi / 180), 0}
	r2.V = Vector{math.Sin(r2.T * math.Pi / 180), math.Cos(r2.T * math.Pi / 180), 0}

	// p3 = p1 + a*v1
	// p3 = p2 + b*v2
	// setting equations to equal each other and solving for a:
	// a = mag[ (p2-p1) cross v2 ] / mag[ v1 cross v2 ]
	// where a is magnitude of distance between r1.P and p3 (intersection point)

	a := r2.P.Subtract(r1.P).Cross(r2.V).Mag() / r1.V.Cross(r2.V).Mag()
	p3 := r1.P.Add(r1.V.MultMag(a))

	fmt.Printf("(%4.1f, %4.1f)\n", p3.X, p3.Y)
}
