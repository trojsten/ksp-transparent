package candle

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

func Download(timetable string) (Timetable, error) {
	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Get(fmt.Sprintf("https://candle.fmph.uniba.sk/rozvrh/%v.xml", timetable))
	if err != nil {
		return Timetable{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Timetable{}, fmt.Errorf("HTTP %v while downloading RozvrhXML: %v", resp.StatusCode, timetable)
	}

	var data Timetable
	err = xml.NewDecoder(resp.Body).Decode(&data)
	return data, err
}
