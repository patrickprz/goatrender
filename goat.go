package main

import (
	"fmt"
	"os"
)

func color(ray Ray) Vector3 {

	sphere := Sphere{Position: Vector3{0, 0, -1}, Radius: 0.5}

	_, record := sphere.Hit(ray, 0, 1)

	if record.T > 0 {
		//n := ray.Direction.Sub(sphere.Position)
		n := ray.Point(record.T).Sub(sphere.Position)
		return Vector3{n.X + 1, n.Y + 1, n.Z + 1}.MultiplyScalar(0.5)
	}

	unitDirection := ray.Direction.Normalize()
	record.T = 0.5*unitDirection.Y + 1.0

	white := Vector3{1.0, 1.0, 1.0}
	blue := Vector3{0.5, 0.7, 1.0}

	return white.MultiplyScalar(1.0 - record.T).Add(blue.MultiplyScalar(record.T))
}

func render() {

	nx := 1200
	ny := 600

	file, _ := os.Create("out.ppm")
	defer file.Close()

	fmt.Fprintf(file, "P3\n%d %d\n255\n", nx, ny)

	lowerLeft := Vector3{-2.0, -1.0, -1.0}
	horizontal := Vector3{4.0, 0.0, 0.0}
	vertical := Vector3{0.0, 2.0, 0.0}
	origin := Vector3{0.0, 0.0, 0.0}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			position := horizontal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))
			direction := lowerLeft.Add(position)

			ray := Ray{origin, direction}
			c := color(ray)

			ir := int(255.99 * c.X)
			ig := int(255.99 * c.Y)
			ib := int(255.99 * c.Z)

			fmt.Fprintf(file, "%d %d %d\n", ir, ig, ib)
		}
	}

}

func main() {
	render()
}
