package main

// Scene contains objects
type Scene struct {
	Objects []Hitable
}

func (scene Scene) Hit(ray *Ray, tMin float64, tMax float64) (bool, HitRecord) {
	record := HitRecord{}
	isHit := false
	closest := tMax

	for _, object := range scene.Objects {
		hit, temp := object.Hit(ray, tMin, closest)

		if hit {
			isHit = true
			closest = temp.T
			record = temp
		}
	}
	return isHit, record
}
