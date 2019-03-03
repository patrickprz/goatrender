package main

import (
	"math"
)

// Sphere object
type Sphere struct {
	Position Vector3
	Radius   float64
}

// Hit sphere test
func (sphere Sphere) Hit(ray Ray, tMin float64, tMax float64) (bool, HitRecord) {
	oc := ray.Origin.Sub(sphere.Position)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - sphere.Radius*sphere.Radius
	discriminant := b*b - 4*a*c

	record := HitRecord{}

	t := (-b - math.Sqrt(discriminant)) / (2.0 * a)

	if discriminant > 0 {
		if t < tMax && t > tMin {
			record.T = t
			record.Position = ray.Point(t)
			record.Normal = record.Position.Sub(sphere.Position).DivideScalar(sphere.Radius)
			return true, record
		}
		if t < tMax && t > tMin {
			t = (-b + math.Sqrt(discriminant)) / (2.0 * a)
			record.T = t
			record.Position = ray.Point(t)
			record.Normal = record.Position.Sub(sphere.Position).DivideScalar(sphere.Radius)
			return false, record
		}
	}

	// if discriminant < 0 {
	// 	record.T = -1
	// 	return true, record
	// }

	// record.T = (-b - math.Sqrt(discriminant)) / (2.0 * a)
	return false, record
}
