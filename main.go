package main

import (
	"fmt"
)

const (
	kmMultiplier      float64 = 0.62
	kgMultiplier      float64 = 2.205
	celsiusMultiplier float64 = 1.8
	celsiusOffset     float64 = 32
)

func kmToMiles(km float64) (float64, string) {
	miles := km * kmMultiplier
	return miles, "miles"
}

func kgToPounds(kg float64) (float64, string) {
	pounds := kg * kgMultiplier
	return pounds, "lbs"
}

func celsiusToFahrenheit(c float64) (float64, string) {
	fahrenheit := (c * celsiusMultiplier) + celsiusOffset
	return fahrenheit, "°F"
}

func main() {
	miles, unitM := kmToMiles(23)
	fmt.Printf("23 km = %.2f %s \n", miles, unitM)

	pounds, unitP := kgToPounds(100)
	fmt.Printf("100 kg = %.2f %s \n", pounds, unitP)

	fahrenheit, unitF := celsiusToFahrenheit(30)
	fmt.Printf("30° C = %.2f %s \n", fahrenheit, unitF)
}
