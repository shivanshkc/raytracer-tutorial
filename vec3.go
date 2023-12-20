package main

import (
	"math"
)

// Vec3 represents a 3D vector.
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// NewVector provides a new Vec3 instance.
func NewVector(x, y, z float64) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

// Plus adds this vector with the given vector and returns the result.
func (v Vec3) Plus(vec Vec3) Vec3 {
	return NewVector(v.X+vec.X, v.Y+vec.Y, v.Z+vec.Z)
}

// Minus subtracts the given vector from this vector and returns the result.
func (v Vec3) Minus(vec Vec3) Vec3 {
	return NewVector(v.X-vec.X, v.Y-vec.Y, v.Z-vec.Z)
}

// Multiply this vector with a scalar.
func (v Vec3) Multiply(arg float64) Vec3 {
	return NewVector(v.X*arg, v.Y*arg, v.Z*arg)
}

// Divide this vector by a scalar.
func (v Vec3) Divide(arg float64) Vec3 {
	return v.Multiply(1 / arg)
}

// Dot product of this vector with another vector.
func (v Vec3) Dot(vec Vec3) float64 {
	return v.X*vec.X + v.Y*vec.Y + v.Z*vec.Z
}

// Cross calculates the cross product of this vector with the given vector
// and returns the result.
func (v Vec3) Cross(arg Vec3) Vec3 {
	return NewVector(
		v.Y*arg.Z-v.Z*arg.Y,
		v.Z*arg.X-v.X*arg.Z,
		v.X*arg.Y-v.Y*arg.X,
	)
}

// Magnitude of this vector.
func (v Vec3) Magnitude() float64 {
	return math.Sqrt(v.Dot(v))
}

// Direction of this vector, also called a unit vector.
func (v Vec3) Direction() Vec3 {
	return v.Divide(v.Magnitude())
}

// Reflection returns the reflection of this vector about the given normal.
func (v Vec3) Reflection(normal Vec3) Vec3 {
	return v.Minus(normal.Multiply(v.Dot(normal)).Multiply(2))
}

// Refraction returns the refracted vector according to the given normal and refraction ratio.
func (v *Vec3) Refraction(normal Vec3, refractionRatio float64) Vec3 {
	incidentDir := v.Direction()
	cosTheta := math.Min(incidentDir.Multiply(-1).Dot(normal), 1)

	refPerpendicular := incidentDir.Plus(normal.Multiply(cosTheta)).
		Multiply(refractionRatio)

	refParallel := normal.Multiply(
		-math.Sqrt(math.Abs(1 - refPerpendicular.Dot(refPerpendicular))))

	return refPerpendicular.Plus(refParallel)
}

// IsNearZero returns true if all components of the vector are near zero.
func (v Vec3) IsNearZero() bool {
	limit := 0.00001
	return v.X < limit && v.Y < limit && v.Z < limit
}
