package main

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
