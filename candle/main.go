package candle

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Timetable struct {
	XMLName xml.Name `xml:"timetable"`
	Lessons []Lesson `xml:"lesson"`
}

type Lesson struct {
	XMLName xml.Name `xml:"lesson"`
	Id      int      `xml:"id,attr"`
	Type    string   `xml:"type"`
	Room    string   `xml:"room"`
	Subject string   `xml:"subject"`
	Day     Weekday  `xml:"day"`
	Start   Time     `xml:"start"`
	End     Time     `xml:"end"`
	Teacher []string `xml:"teacher"`
	Note    string   `xml:"note"`
}

type Weekday string

func (w Weekday) Weekday() time.Weekday {
	switch w {
	case "Po":
		return time.Monday
	case "Ut":
		return time.Tuesday
	case "St":
		return time.Wednesday
	case "Št":
		return time.Thursday
	case "Pi":
		return time.Friday
	}
	panic(fmt.Errorf("%v is not a valid weekday", w))
}

type Time string

func (t Time) Time() time.Time {
	parsed, err := time.ParseInLocation("15:04", string(t), time.Local)
	if err != nil {
		panic(err)
	}
	return parsed
}
