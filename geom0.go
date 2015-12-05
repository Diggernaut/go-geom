package geom

type geom0 struct {
	layout     Layout
	stride     int
	flatCoords []float64
}

func (g *geom0) Bounds() *Bounds {
	return NewBounds().extendFlatCoords(g.flatCoords, 0, len(g.flatCoords), g.stride)
}

func (g *geom0) Coords() []float64 {
	return inflate0(g.flatCoords, 0, len(g.flatCoords), g.stride)
}

func (g *geom0) Ends() []int {
	return nil
}

func (g *geom0) Endss() [][]int {
	return nil
}

func (g *geom0) FlatCoords() []float64 {
	return g.flatCoords
}

func (g *geom0) Layout() Layout {
	return g.layout
}

func (g *geom0) NumCoords() int {
	return 1
}

func (g *geom0) SetCoords(coords0 []float64) error {
	var err error
	if g.flatCoords, err = deflate0(nil, coords0, g.stride); err != nil {
		return err
	}
	return nil
}

func (g *geom0) Stride() int {
	return g.stride
}

func (g *geom0) mustVerify() {
	if g.stride != g.layout.Stride() {
		panic("geom: stride/layout mismatch")
	}
	if len(g.flatCoords) != g.stride {
		panic("geom: length/stride mismatch")
	}

}

func deflate0(flatCoords []float64, c []float64, stride int) ([]float64, error) {
	if len(c) != stride {
		return nil, ErrStrideMismatch{Got: len(c), Want: stride}
	}
	flatCoords = append(flatCoords, c...)
	return flatCoords, nil
}

func inflate0(flatCoords []float64, offset, end, stride int) []float64 {
	if offset+stride != end {
		panic("geom: stride mismatch")
	}
	c := make([]float64, stride)
	copy(c, flatCoords[offset:end])
	return c
}