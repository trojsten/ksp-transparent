package main

import (
	"github.com/gin-gonic/gin"
	"ksp.sk/transparent/candle"
	"ksp.sk/transparent/config"
	"ksp.sk/transparent/event"
	"net/http"
	"time"
)

import _ "gopkg.in/yaml.v2"

func getCandle(c *gin.Context) {
	events, err := candle.Events(time.Now())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	events = event.SetEventWidths(events)

	var ticks []time.Time
	for i := 8; i < 23; i++ {
		ticks = append(ticks, time.Date(0, 0, 0, i, 0, 0, 0, time.Local))
	}

	c.HTML(http.StatusOK, "candle.gohtml", gin.H{
		"events": events,
		"ticks":  ticks,
	})
}

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", nil)
}

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.FuncMap["timeString"] = func(t time.Time) string {
		return t.Format("15:04")
	}
	router.FuncMap["topForTime"] = func(t time.Time) int {
		t = t.Local()
		return 120*(t.Hour()-8) + t.Minute()*2
	}
	router.FuncMap["pxForDuration"] = func(d time.Duration) int {
		return int(d.Seconds()/60) * 2
	}
	router.FuncMap["percentage"] = func(a, b int) float64 {
		return float64(a) / float64(b) * 100
	}
	router.FuncMap["multiply"] = func(a float64, b int) float64 {
		return a * float64(b)
	}
	router.LoadHTMLGlob("templates/*")

	router.GET("/", getIndex)
	router.GET("/candle", getCandle)
	router.Static("/static", "./static")

	err = router.Run()
	if err != nil {
		panic(err)
	}
}
