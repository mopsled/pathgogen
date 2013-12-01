package main

type Propagator struct {
	grid *Grid
}

func NewPropagator(grid *Grid) *Propagator {
	return &Propagator{grid}
}

func (p Propagator) Propagate(cell *Cell, coordinate string) {
	p.grid.Set(coordinate, cell)
}
