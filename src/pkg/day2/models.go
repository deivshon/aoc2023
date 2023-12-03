package day2

type Round struct {
	GameId int
	Red    int
	Green  int
	Blue   int
}

type Color string

const (
	ColorRed   Color = "red"
	ColorGreen Color = "green"
	ColorBlue  Color = "blue"
)
