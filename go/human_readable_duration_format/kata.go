package kata

import (
	"math"
	"strconv"
	"strings"
)

func plural(count int, singular string) (result string) {
	if (count == 1) || (count == 0) {
		result = strconv.Itoa(count) + " " + singular
	} else {
		result = strconv.Itoa(count) + " " + singular + "s"
	}
	return
}

func FormatDuration(input int64) (result string) {
	if input == 0 {
		return "now"
	}

	years := math.Floor(float64(input) / 60 / 60 / 24 / 365)
	seconds := input % (60 * 60 * 24 * 365)
	days := math.Floor(float64(seconds) / 60 / 60 / 24)
	seconds = input % (60 * 60 * 24)
	hours := math.Floor(float64(seconds) / 60 / 60)
	seconds = input % (60 * 60)
	minutes := math.Floor(float64(seconds) / 60)
	seconds = input % 60

	var temp []string
	if years > 0 {
		temp = append(temp, plural(int(years), "year"))
	}
	if days > 0 {
		temp = append(temp, plural(int(days), "day"))
	}
	if hours > 0 {
		temp = append(temp, plural(int(hours), "hour"))
	}
	if minutes > 0 {
		temp = append(temp, plural(int(minutes), "minute"))
	}
	if seconds > 0 {
		temp = append(temp, plural(int(seconds), "second"))
	}

	if len(temp) > 1 {
		result = strings.Join(temp[:len(temp)-1], ", ")
		result += " and " + temp[len(temp)-1]
	} else {
		result = temp[0]
	}
	return
}
