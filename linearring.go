package geom

import "math"

// A LinearRing is a linear ring.
type LinearRing struct {
	geom1
}

// NewLinearRing returns a new LinearRing with no coordinates.
func NewLinearRing(layout Layout) *LinearRing {
	return NewLinearRingFlat(layout, nil)
}

// NewLinearRingFlat returns a new LinearRing with the given flat coordinates.
func NewLinearRingFlat(layout Layout, flatCoords []float64) *LinearRing {
	lr := new(LinearRing)
	lr.layout = layout
	lr.stride = layout.Stride()
	lr.flatCoords = flatCoords
	return lr
}

// Area returns the the area.
func (lr *LinearRing) Area() float64 {
	return doubleArea1(lr.flatCoords, 0, len(lr.flatCoords), lr.stride) / 2
}

// Clone returns a deep copy.
func (lr *LinearRing) Clone() *LinearRing {
	return deriveCloneLinearRing(lr)
}

// Empty returns false.
func (lr *LinearRing) Empty() bool {
	return false
}

// Length returns the length of the perimeter.
func (lr *LinearRing) Length() float64 {
	return length1(lr.flatCoords, 0, len(lr.flatCoords), lr.stride)
}

// MustSetCoords sets the coordinates and panics if there is any error.
func (lr *LinearRing) MustSetCoords(coords []Coord) *LinearRing {
	Must(lr.SetCoords(coords))
	return lr
}

// SetCoords sets the coordinates.
func (lr *LinearRing) SetCoords(coords []Coord) (*LinearRing, error) {
	if err := lr.setCoords(coords); err != nil {
		return nil, err
	}
	return lr, nil
}

// SetSRID sets the SRID of lr.
func (lr *LinearRing) SetSRID(srid int) *LinearRing {
	lr.srid = srid
	return lr
}

// Swap swaps the values of lr and lr2.
func (lr *LinearRing) Swap(lr2 *LinearRing) {
	*lr, *lr2 = *lr2, *lr
}

// ContainCoord checks if LinearRing contains Coord.
func (lr *LinearRing) ContainCoord(pt *Coord) bool {
	// Cast ray from pt.X towards the right
	intersections := 0
	c := lr.Coords()
	for i := range c {
		curr := c[i]
		ii := i + 1
		if ii == len(c) {
			continue
		}
		next := c[ii]

		// Is the point out of the edge's bounding box?
		// bottom vertex is inclusive (belongs to edge), top vertex is
		// exclusive (not part of edge) -- i.e. p lies "slightly above
		// the ray"
		bottom, top := curr, next
		if bottom.Y() > top.Y() {
			bottom, top = top, bottom
		}
		if pt.Y() < bottom.Y() || pt.Y() >= top.Y() {
			continue
		}
		// Edge is from curr to next.

		if pt.X() >= math.Max(curr.X(), next.X()) ||
			next.Y() == curr.Y() {
			continue
		}

		// Find where the line intersects...
		xint := (pt.Y()-curr.Y())*(next.X()-curr.X())/(next.Y()-curr.Y()) + curr.X()
		if curr.X() != next.X() && pt.X() > xint {
			continue
		}
		intersections++
	}
	return intersections%2 != 0
}
