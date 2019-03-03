package main

// Ray struct
type Ray struct {
	Origin    Vector3
	Direction Vector3
}

// Point Get the parametric position of the ray
func (ray Ray) Point(t float64) Vector3 {
	b := ray.Direction.MultiplyScalar(t)
	a := ray.Origin
	return a.Add(b)
}
