package math

import (
	"fmt"
	"strconv"
)

// Media é responsável por calcular o que ja sabemos
func Media(numbers ...float64) float64 {
	total := 0.0

	for _, num := range numbers {
		total += num
	}

	media := total / float64(len(numbers))

	roundMedia, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)

	return roundMedia
}
