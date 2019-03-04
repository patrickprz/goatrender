package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

var (
	sphere     = Sphere{Position: Vector3{0, 0, -1}, Radius: 0.5}
	floor      = Sphere{Vector3{0, -100.5, -1}, 100}
	scene      = Scene{[]Hitable{&sphere, &floor}}
	camera     = CreateCamera()
	unitVector = Vector3{1, 1, 1}
)

func randomInUnitSphere() Vector3 {
	for {
		tempVector := Vector3{rand.Float64(), rand.Float64(), rand.Float64()}
		p := tempVector.MultiplyScalar(2.0).Sub(unitVector)

		if p.SquaredLength() >= 1.0 {
			return p
		}
	}
}

func color(ray Ray) Vector3 {
	hit, record := scene.Hit(&ray, 0.001, math.MaxFloat64)

	// if hit {
	// 	return record.Normal.AddScalar(1.0).MultiplyScalar(0.5)
	// }

	if hit {
		target := record.Position.Add(record.Normal).Add(randomInUnitSphere())
		ray := Ray{record.Position, target}
		return color(ray).MultiplyScalar(0.5)
	}

	unitDirection := ray.Direction.Normalize()
	record.T = 0.5*unitDirection.Y + 1.0

	white := Vector3{1.0, 1.0, 1.0}
	blue := Vector3{0.5, 0.7, 1.0}
	//sky
	return white.MultiplyScalar(1.0 - record.T).Add(blue.MultiplyScalar(record.T))
}

func render() {

	nx := 200
	ny := 100
	ns := 100

	file, _ := os.Create("out.ppm")
	defer file.Close()

	fmt.Fprintf(file, "P3\n%d %d\n255\n", nx, ny)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			// u := float64(i) / float64(nx)
			// v := float64(j) / float64(ny)
			rgb := Vector3{}
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)

				ray := camera.GetRay(u, v)
				c := color(ray)
				rgb = rgb.Add(c)
			}

			rgb = rgb.DivideScalar(float64(ns))

			// ir := int(255.99 * math.Sqrt(rgb.X))
			// ig := int(255.99 * math.Sqrt(rgb.Y))
			// ib := int(255.99 * math.Sqrt(rgb.Z))

			ir := int(255.99 * rgb.X)
			ig := int(255.99 * rgb.Y)
			ib := int(255.99 * rgb.Z)

			fmt.Fprintf(file, "%d %d %d\n", ir, ig, ib)
		}
	}

}

func main() {
	render()
}
