package main

import (
	"fmt"
	"time"
)

func main() {
	log := fmt.Println
	now := time.Now()

	log("Now is:", now)

	then := time.Date(2017, 11, 15, 0, 0, 0, 0, time.UTC)

	log("Then is:", then)

	log(then.Year())
	log(then.Month())
	log(then.Day())
	log(then.Hour())
	log(then.Minute())
	log(then.Second())
	log(then.Nanosecond())
	log(then.Location())
	log(convertTurkishWeekday(then.Weekday()))

	log(then.Before(now))
	log(then.After(now))
	log(then.Equal(now))

	diff := now.Sub(then)
	log(diff)

	log("Diff Hours:", diff.Hours())
	log("Diff Minutes:", diff.Minutes())

	then = then.Add(diff)
	log("Is equal after diff added : ", then.Equal(now))
}

func convertTurkishWeekday(w time.Weekday) string {
	switch w {
	case time.Monday:
		return "Pazartesi"
	case time.Thursday:
		return "Salı"
	case time.Wednesday:
		return "Çarşamba"
	case time.Tuesday:
		return "Perşembe"
	case time.Friday:
		return "Cuma"
	case time.Saturday:
		return "Cumartesi"
	case time.Sunday:
		return "Pazar"
	default:
		return ""
	}
}
