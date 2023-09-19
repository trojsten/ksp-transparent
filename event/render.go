package event

import (
	"fmt"
	"sort"
	"time"
)

func SetEventWidths(events []Event) []Event {
	sort.Slice(events, func(i, j int) bool {
		return events[i].Start.Before(events[j].Start)
	})

	currentIndexes := []int{}
	currentEnd := time.Time{}
	for i, event := range events {
		fmt.Println(event.Start, currentEnd)

		if event.Start.After(currentEnd) || currentEnd.IsZero() {
			for j, index := range currentIndexes {
				events[index].Concurrent = len(currentIndexes)
				events[index].Offset = j
			}
			currentIndexes = []int{}
		}

		if event.End.After(currentEnd) || currentEnd.IsZero() {
			currentEnd = event.End
		}

		currentIndexes = append(currentIndexes, i)
	}

	for j, index := range currentIndexes {
		events[index].Concurrent = len(currentIndexes)
		events[index].Offset = j
	}

	return events
}
