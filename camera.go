package main

// Camera struct
type Camera struct {
	lowerLeft  Vector3
	horizontal Vector3
	vertical   Vector3
	origin     Vector3
}

// CreateCamera initialize a camera
func CreateCamera() *Camera {
	camera := new(Camera)
	camera.lowerLeft = Vector3{-2.0, -1.0, -1.0}
	camera.horizontal = Vector3{4.0, 0.0, 0.0}
	camera.vertical = Vector3{0.0, 2.0, 0.0}
	camera.origin = Vector3{0.0, 0.0, 0.0}

	return camera
}

// GetRay Return ray from camera
func (camera *Camera) GetRay(u float64, v float64) Ray {
	origin := camera.horizontal.MultiplyScalar(u).Add(camera.vertical.MultiplyScalar(v))
	direction := camera.lowerLeft.Add(origin)

	return Ray{camera.origin, direction}
}
