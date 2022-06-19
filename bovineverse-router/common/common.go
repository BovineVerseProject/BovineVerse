package common

import (
	"fmt"
	"math"
	"strconv"
)

func MustNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func CalUPrice(amount float64, precision float64, rate float64) float64 {
	if rate <= 0 {
		rate = 1
	}

	value := amount * rate
	if precision > 0 {
		p := math.Pow(10, precision)
		value = math.Ceil(value*p) / p
		value, _ = strconv.ParseFloat(fmt.Sprintf(fmt.Sprint("%.", precision, "f"), value), 64)
	}

	return value
}
