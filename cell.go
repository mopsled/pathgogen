package main

type CellType int

const (
	nilCell CellType = iota
	aCell
	bCell
	cCell
	wallCell
)

type Cell struct {
	cellType CellType
	team     string
}
