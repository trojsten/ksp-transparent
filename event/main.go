package event

import (
	"github.com/teacat/noire"
	"ksp.sk/transparent/config"
	"time"
)

type Event struct {
	Start      time.Time
	End        time.Time
	Title      string
	Location   string
	Color      string
	People     []config.Person
	Offset     int
	Concurrent int
}

func (e Event) Duration() time.Duration {
	return e.End.Sub(e.Start)
}

func (e Event) BgColor() string {
	color := noire.NewHex(e.Color)
	//return "#fff"
	return color.Lighten(0.25).HTML()
}
