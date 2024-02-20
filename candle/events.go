package candle

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"time"

	"log"

	"ksp.sk/transparent/config"
	"ksp.sk/transparent/event"
)

func Events(day time.Time) ([]event.Event, error) {
	lessonPeople := map[int][]config.Person{}
	lessons := map[int]Lesson{}

	for _, person := range config.Get().People {
		if person.Candle == "" {
			continue
		}

		tt, err := Get(person.Candle)
		if err != nil {
			log.Printf("error while loading timetable %s: %v", person.Name, err)
			continue
		}

		for _, lesson := range tt.Lessons {
			if lesson.Day.Weekday() != day.Weekday() {
				continue
			}

			_, exists := lessons[lesson.Id]
			if !exists {
				lessons[lesson.Id] = lesson
			}
			lessonPeople[lesson.Id] = append(lessonPeople[lesson.Id], person)
		}
	}

	var events []event.Event
	for _, lesson := range lessons {
		h := sha1.New()
		h.Write([]byte(lesson.Subject))
		color := hex.EncodeToString(h.Sum(nil))[2 : 6+2]
		events = append(events, event.Event{
			Start:    lesson.Start.Time(),
			End:      lesson.End.Time(),
			Title:    lesson.Subject,
			Location: strings.Trim(lesson.Room, " "),
			Color:    fmt.Sprintf("#%s", color),
			People:   lessonPeople[lesson.Id],
		})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Start.Before(events[j].Start)
	})

	return events, nil
}
