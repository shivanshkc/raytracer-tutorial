package main

// Ray represents a ray of light.
type Ray struct {
	Origin    Vec3
	Direction Vec3
}

// NewRay returns a new Ray instance.
func NewRay(origin Vec3, direction Vec3) Ray {
	return Ray{Origin: origin, Direction: direction}
}
