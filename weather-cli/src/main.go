package main

import "fmt"

func main() {
	fmt.Println("Welcome to Weather CLI :)")
	fmt.Println("You can get up to date current weather information based on a location name")

	location := readInput()

	weather, err := getCurrentWeather(location)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v°C in %v, %v\n", weather.Temp, weather.Location, weather.Condition)
	}
}

func readInput() string {
	var location string

	fmt.Println("Please enter location:")
	fmt.Scan(&location)

	if location == "" {
		readInput()
	}

	return location
}
