package main

import "math"

// Vector3 XYZ
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// Normalize a vector3
func (a Vector3) Normalize() Vector3 {
	magnitude := math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
	return Vector3{a.X / magnitude, a.Y / magnitude, a.Z / magnitude}
}

// Dot product between two vector3
func (a Vector3) Dot(b Vector3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Length of a vector3
func (a Vector3) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Add two vector3
func (a Vector3) Add(b Vector3) Vector3 {
	return Vector3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Sub two vector3
func (a Vector3) Sub(b Vector3) Vector3 {
	return Vector3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// MultiplyScalar in a vector3
func (a Vector3) MultiplyScalar(t float64) Vector3 {
	return Vector3{a.X * t, a.Y * t, a.Z * t}
}

// DivideScalar in a vector3
func (a Vector3) DivideScalar(t float64) Vector3 {
	return Vector3{a.X / t, a.Y / t, a.Z / t}
}

// AddScalar in a vector3
func (a Vector3) AddScalar(t float64) Vector3 {
	return Vector3{a.X + t, a.Y + t, a.Z + t}
}
