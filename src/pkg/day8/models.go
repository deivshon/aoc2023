package day8

const waypointLength = 3

type DesertLookup struct {
	Left  [waypointLength]byte
	Right [waypointLength]byte
}

type DesertMap = map[[waypointLength]byte]DesertLookup

type DesertTravel struct {
	Path string
	Map  DesertMap
}

type Direction = rune

const (
	DirectionLeft  Direction = 'L'
	DirectionRight Direction = 'R'
)
