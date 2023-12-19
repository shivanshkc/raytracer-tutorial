package main

import (
	"math"
)

// HitInfo contains information about an object hit.
type HitInfo struct {
	// Point of hit.
	Point                 Vec3
	DistanceFromRayOrigin float64

	// Normal from the surface at the point of hit.
	Normal Vec3
	// Flag to tell whether the normal points inward or outward.
	IsNormalOutward bool
}

// setNormalFacing sets the facing direction of the normal.
func (h *HitInfo) setNormalFacing(ray Ray) {
	h.IsNormalOutward = ray.Direction.Dot(h.Normal) < 0
	if !h.IsNormalOutward {
		h.Normal = h.Normal.Multiply(-1)
	}
}

// Hittable represents an object that can be rendered using Ray Tracing.
type Hittable interface {
	// IsHit checks whether the object is hit by the given ray.
	// If yes, its hit information is returned.
	IsHit(ray Ray, tMin, tMax float64) (HitInfo, bool)
}

// Sphere implements Hittable using a spherical geometry.
type Sphere struct {
	Center Vec3
	Radius float64
}

func (s Sphere) IsHit(ray Ray, tMin, tMax float64) (HitInfo, bool) {
	origin2Center := ray.Origin.Minus(s.Center)

	// Calculations as per the sphere equation.
	a := ray.Direction.Dot(ray.Direction)
	bHalf := origin2Center.Dot(ray.Direction)
	c := origin2Center.Dot(origin2Center) - s.Radius*s.Radius

	discriminant := bHalf*bHalf - a*c
	if discriminant < 0 {
		return HitInfo{}, false
	}

	// Note that root1 is always smaller than root2 because
	// discriminant is positive.
	root1 := (-bHalf - math.Sqrt(discriminant)) / a
	root2 := (-bHalf + math.Sqrt(discriminant)) / a

	closerRoot := root1
	// If root1 is outside of bounds...
	if root1 < tMin || root1 > tMax {
		// If root2 is outside of bounds...
		if root2 < tMin || root2 > tMax {
			return HitInfo{}, false
		}
		// Update the closer root.
		closerRoot = root2
	}

	// Create the hit information object.
	var hitInfo HitInfo
	hitInfo.Point = ray.PointAt(closerRoot)
	hitInfo.DistanceFromRayOrigin = closerRoot

	hitInfo.Normal = hitInfo.Point.Minus(s.Center).Direction()
	hitInfo.setNormalFacing(ray)

	return hitInfo, true
}
