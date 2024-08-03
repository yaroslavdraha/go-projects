package main

import "fmt"

func main() {
	fmt.Println("Welcome to Weather CLI :)")
	fmt.Println("You can get up to date current weather information based on a location name")

	location := readInput()

	getCurrentWeather(location)
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
