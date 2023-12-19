package main

import (
	"math"
)

// Material is an abstraction for the properties of a surface.
type Material interface {
	// Scatter simulates the scattering of a ray and returns the scattered ray,
	// the attenuation and a flag to tell whether the ray was scattered at all.
	Scatter(incomingRay Ray, hitInfo HitInfo) (Ray, Color, bool)
}

// Metal implements the Material interface for metallic surfaces.
type Metal struct {
	Attenuation Color
	Fuzz        float64
}

func (m Metal) Scatter(incomingRay Ray, info HitInfo) (Ray, Color, bool) {
	reflected := incomingRay.Direction.Reflection(info.Normal)
	scatterDir := reflected.Direction().
		Plus(RandomVectorInUnitSphere().Multiply(m.Fuzz)).Direction()
	scattered := NewRay(info.Point, scatterDir)

	isScattered := scattered.Direction.Dot(info.Normal) > 0
	return scattered, m.Attenuation, isScattered
}

// Lambertian implements the Material interface for rough materials.
type Lambertian struct {
	Attenuation Color
}

func (l Lambertian) Scatter(incomingRay Ray, info HitInfo) (Ray, Color, bool) {
	scatterDir := info.Normal.Plus(RandomVectorInUnitSphere().Direction()).Direction()

	// Catch degenerate scatter direction
	if scatterDir.IsNearZero() {
		scatterDir = info.Normal
	}

	scattered := NewRay(info.Point, scatterDir)
	return scattered, l.Attenuation, true
}

// Dielectric represents glass-like materials that refract light.
type Dielectric struct {
	RefractiveIndex float64
}

func (d Dielectric) Scatter(incomingRay Ray, info HitInfo) (Ray, Color, bool) {
	refractionRatio := d.RefractiveIndex
	if info.IsNormalOutward {
		refractionRatio = 1 / refractionRatio
	}

	incomingDir := incomingRay.Direction.Direction()

	cosTheta := math.Min(incomingDir.Multiply(-1).Dot(info.Normal), 1)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1

	var scatterDir Vec3
	if cannotRefract || d.schlickAppoximation(cosTheta, refractionRatio) > RandomFloat() {
		scatterDir = incomingDir.Reflection(info.Normal)
	} else {
		scatterDir = incomingDir.Refraction(info.Normal, refractionRatio)
	}

	scattered := NewRay(info.Point, scatterDir.Direction())
	return scattered, NewColor(1, 1, 1), true
}

func (d *Dielectric) schlickAppoximation(cosine, refractionRatio float64) float64 {
	r0 := math.Pow((1-refractionRatio)/(1+refractionRatio), 2)
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
