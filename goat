package main

import (
	"math"
)

// Geometry interface
type Geometry interface {
	intersect() bool
}

// Vector3 XYZ
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// Color RGB
type Color struct {
	R uint8
	G uint8
	B uint8
}

// Plane object
type Plane struct {
	Position Vector3
	Normal   Vector3
}

// Sphere object
type Sphere struct {
	Position Vector3
	Radius   float64
}

// Ray struct
type Ray struct {
	Origin    Vector3
	Direction Vector3
}

const (
	// WIDTH Screen width
	WIDTH int = 480
	// HEIGHT  Height
	HEIGHT int = 240
	// FOV CONST
	FOV float32 = 51.52
)

func normalize(a Vector3) Vector3 {
	magnitude := math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
	return Vector3{a.X / magnitude, a.Y / magnitude, a.Z / magnitude}
}

func dot(a Vector3, b Vector3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func sub(a Vector3, b Vector3) Vector3 {
	return Vector3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (plane Plane) intersect(ray Ray) bool {
	denom := dot(normalize(plane.Normal), normalize(ray.Direction))
	if denom > 1e-6 {
		t := dot(sub(normalize(plane.Position), normalize(ray.Origin)), normalize(plane.Normal)) / denom
		if t >= 0 {
			return true
		}
	}
	return false
}

func (sphere Sphere) intersect(ray Ray) bool {
	// Analytic solution
	L := sub(sphere.Position, ray.Origin)
	tca := dot(L, ray.Direction)
	d2 := dot(L, L) - tca*tca
	r2 := sphere.Radius * sphere.Radius
	if d2 > r2 {
		return false
	}
	thc := math.Sqrt(r2 - d2)
	t0 := tca - thc
	t1 := tca + thc

	if t0 > t1 {
		t0, t1 = t1, t0
	}

	if t0 < 0 {
		t0 = t1
		if t0 < 0 {
			return false
		}
	}

	return false
}

func render() {

}

func main() {
	// Render loop
	//for y := 0; y < HEIGHT; y++ {
	//for x := 0; x < WIDTH; x++ {

	//}
	//}
}
