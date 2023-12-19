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

// PointAt returns the point at which the ray arrives after traveling the given distance.
func (r Ray) PointAt(distance float64) Vec3 {
	return r.Origin.Plus(r.Direction.Multiply(distance))
}
