package event

import (
	"fmt"
	"sort"
	"time"
)

func removeDuplicateValues(intSlice []time.Time) []time.Time {
    keys := make(map[time.Time]bool)
    list := []time.Time{}
 
    // If the key(values of the slice) is not equal
    // to the already present value in new slice (list)
    // then we append it. else we jump on another element.
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func SetEventWidths(events []Event) []Event {
	starts := []time.Time{}

	for i, event := range events {
		starts = append(starts, event.Start)
		events[i].Offset = -1
	}

	starts = removeDuplicateValues(starts)

	sort.Slice(starts, func(i, j int) bool {
		return starts[i].Before(starts[j])
	})
	var max int = 0
	prekryvy := make(map[time.Time][]Event)
	for _, t := range starts {
		for _, event := range events {
			if !(event.End.Before(t) || event.Start.After(t)) {
				prekryvy[t] = append(prekryvy[t], event)
			}
		}
		if max < len(prekryvy[t]) {
			max = len(prekryvy[t])
		}
	}
	fmt.Println("max: ", max)

	for i := range events {
		events[i].Concurrent = max
	}

	for _, t := range starts {
		pozicie := []int{}
		for i := 0; i < 100; i++ {
			pozicie = append(pozicie, 0)
		}
		for _, event := range events {
			if !(event.End.Before(t) || event.Start.After(t)) {
				if event.Offset != -1{
					pozicie[event.Offset] = 1
				}
			}
		}

		for eventid, event := range events {
			if !(event.End.Before(t) || event.Start.After(t)) {
				if event.Offset == -1{
					for i := 0; i < max; i++ {
						if pozicie[i] == 0 {
							pozicie[i] = 1
							events[eventid].Offset = i
							break
						}
					}
				}
			}
		}
		
	}
	
	return events
}
