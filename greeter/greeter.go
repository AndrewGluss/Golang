package greeter

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func Greet(name string, hour int) string {
	value := strconv.Itoa(hour) + ":00:00"
	result, err := time.Parse("15:04:05", value)
	if err != nil {
		return "Hour out of range"
	} else {
		var greeting string

		morning, _ := time.Parse("15:04:05", "06:00:00")
		day, _ := time.Parse("15:04:05", "12:00:00")
		evening, _ := time.Parse("15:04:05", "18:00:00")
		night, _ := time.Parse("15:04:05", "22:00:00")

		if result.Hour() >= morning.Hour() && result.Hour() < day.Hour() {
			greeting = "Good morning"
		} else if result.Hour() >= day.Hour() && result.Hour() <= evening.Hour() {
			greeting = "Hello"
		} else if result.Hour() >= evening.Hour() && result.Hour() < night.Hour() {
			greeting = "Good evening"
		} else if (result.Hour() >= night.Hour() && result.Hour() < 24) || (result.Hour() < morning.Hour() && result.Hour() >= 0) {
			greeting = "Good night"
		}
		trimmedName := strings.Trim(name, " ")
		chars := []rune(trimmedName)
		first := unicode.ToUpper(chars[0])
		resultName := string(first) + string(chars[1:])
		return fmt.Sprintf("%s, %s!", greeting, strings.Title(resultName))
	}
}
