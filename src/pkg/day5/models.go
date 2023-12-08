package day5

type AlmanacLine struct {
	SourceStart      int
	DestinationStart int
	RangeLength      int
}

type AlmanacMap []AlmanacLine

type Almanac struct {
	Seeds []int
	Maps  []AlmanacMap
}
