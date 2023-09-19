package candle

import (
	"time"
)

type cacheItem struct {
	Timetable Timetable
	Timestamp time.Time
}

var cache map[string]cacheItem

func Get(timetable string) (Timetable, error) {
	if cache == nil {
		cache = map[string]cacheItem{}
	}

	c, exists := cache[timetable]
	if !exists || c.Timestamp.Before(time.Now().Add(-6*time.Hour)) {
		tt, err := Download(timetable)
		if err != nil {
			return Timetable{}, err
		}

		cache[timetable] = cacheItem{
			Timetable: tt,
			Timestamp: time.Now(),
		}
	}

	return cache[timetable].Timetable, nil
}
