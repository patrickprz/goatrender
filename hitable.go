package main

// HitRecord from ray hit
type HitRecord struct {
	T        float64
	Position Vector3
	Normal   Vector3
}

// Hitable interface to objects
type Hitable interface {
	Hit(r *Ray, tMin float64, tMax float64) (bool, HitRecord)
}
