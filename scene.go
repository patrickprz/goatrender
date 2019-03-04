package main

// Scene contains objects
type Scene struct {
	Objects []Hitable
}

// Hit objects
func (scene Scene) Hit(ray *Ray, tMin float64, tMax float64) (bool, HitRecord) {
	isHit := false
	closest := tMax
	record := HitRecord{}

	for _, object := range scene.Objects {
		hit, tempRecord := object.Hit(ray, tMin, closest)

		if hit {
			isHit = true
			closest = tempRecord.T
			record = tempRecord
		}
	}
	return isHit, record
}
