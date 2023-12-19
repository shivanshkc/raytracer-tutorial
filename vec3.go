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

// Magnitude of this vector.
func (v Vec3) Magnitude() float64 {
	return math.Sqrt(v.Dot(v))
}

// Direction of this vector, also called a unit vector.
func (v Vec3) Direction() Vec3 {
	return v.Divide(v.Magnitude())
}
