package geom

import (
	"math"
)

type LineString struct {
	layout     Layout
	stride     int
	flatCoords []float64
}

var _ T = &LineString{}

func NewLineString(layout Layout, coords1 [][]float64) (*LineString, error) {
	ls := &LineString{
		layout:     layout,
		stride:     layout.Stride(),
		flatCoords: nil,
	}
	var err error
	if ls.flatCoords, err = deflate1(ls.flatCoords, coords1, ls.stride); err != nil {
		return nil, err
	}
	return ls, nil
}

func NewLineStringFlat(layout Layout, flatCoords []float64) *LineString {
	return &LineString{
		layout:     layout,
		stride:     layout.Stride(),
		flatCoords: flatCoords,
	}
}

func (ls *LineString) Clone() *LineString {
	flatCoords := make([]float64, len(ls.flatCoords))
	copy(flatCoords, ls.flatCoords)
	return &LineString{
		layout:     ls.layout,
		stride:     ls.stride,
		flatCoords: flatCoords,
	}
}

func (ls *LineString) Coords() interface{} {
	return inflate1(ls.flatCoords, 0, len(ls.flatCoords), ls.stride)
}

func (ls *LineString) Ends() []int {
	return nil
}

func (ls *LineString) Endss() [][]int {
	return nil
}

func (ls *LineString) Envelope() *Envelope {
	return NewEnvelope().extendFlatCoords(ls.flatCoords, 0, len(ls.flatCoords), ls.stride)
}

func (ls *LineString) FlatCoords() []float64 {
	return ls.flatCoords
}

func (ls *LineString) Layout() Layout {
	return ls.layout
}

func (ls *LineString) Length() float64 {
	length := 0.0
	for i := ls.stride; i < len(ls.flatCoords); i += ls.stride {
		dx := ls.flatCoords[i] - ls.flatCoords[i-ls.stride]
		dy := ls.flatCoords[i+1] - ls.flatCoords[i+1-ls.stride]
		length += math.Sqrt(dx*dx + dy*dy)
	}
	return length
}

func (ls *LineString) Stride() int {
	return ls.stride
}