package adventofcode

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func Day1Part1() int {
	file, err := os.Open("./input/day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fuelRequired := 0

	for scanner.Scan() {
		var moduleMass, err = strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		fuelRequired += int(math.Floor(float64(moduleMass / 3))) - 2
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fuelRequired
}

func Day1Part2() int {
	file, err := os.Open("./input/day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fuelRequired := 0

	for scanner.Scan() {
		var moduleMass, err = strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		fuelRequiredForModuleMass := int(math.Floor(float64(moduleMass / 3))) - 2
		extraFuelRequired := CalculateFuelRequiredByFuel(fuelRequiredForModuleMass)
		fuelRequired += fuelRequiredForModuleMass + extraFuelRequired
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fuelRequired
}

func CalculateFuelRequiredByFuel(fuelRequired int) int {

	result := 0

	extraFuelRequired := int(math.Floor(float64(fuelRequired/3))) - 2

	if extraFuelRequired > 0 {
		result += extraFuelRequired
	}

	for extraFuelRequired > 0 {
		extraFuelRequired = int(math.Floor(float64(extraFuelRequired/3))) - 2

		if extraFuelRequired > 0 {
			result += extraFuelRequired
		}

	}
	return result
}