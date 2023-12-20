package main

import (
	"math/rand"
	"time"
)

// Create a random number generator.
var _randSeed int64 = int64(time.Now().Nanosecond())
var _randomGen = rand.New(rand.NewSource(_randSeed))

func RandomFloat() float64 {
	return _randomGen.Float64()
}

// RandomFloatBetween generates a random float between the given min and max range.
func RandomFloatBetween(min, max float64) float64 {
	return min + (RandomFloat() * (max - min))
}

func RandomVector() Vec3 {
	return NewVector(RandomFloat(), RandomFloat(), RandomFloat())
}

func RandomVectorInUnitSphere() Vec3 {
	for {
		point := RandomVector()
		if point.Dot(point) < 1 {
			return point
		}
	}
}

// RandomVectorInUnitDisk returns a random Vec3 inside a unit disk.
func RandomVectorInUnitDisk() Vec3 {
	for {
		vec := NewVector(RandomFloatBetween(-1, 1), RandomFloatBetween(-1, 1), 0)
		if vec.Dot(vec) < 1 {
			return vec
		}
	}
}
