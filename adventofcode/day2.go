package adventofcode

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const AddOpcode int = 1
const MultOpcode int = 2
const ExitOpcode = 99

func Day2Part1() int {
	file, err := os.Open("./input/day2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var result []string

	for scanner.Scan() {
		var listOfOpcodes = scanner.Text()
		result = strings.Split(listOfOpcodes, ",")
	}

	var intCodes []int

	//Convert to array of ints
	for _, element := range result {
		var opcode, err = strconv.Atoi(element)

		if err != nil {
			log.Fatal(err)
		}

		intCodes = append(intCodes, opcode)
	}

	var i = 0
	for {
		var opCode = intCodes[i]

		if opCode == AddOpcode {
			var sum = intCodes[intCodes[i+1]] + intCodes[intCodes[i+2]]
			intCodes[intCodes[i+3]] = sum
			i = i + 4
		} else if opCode == MultOpcode {
			var product = intCodes[intCodes[i+1]] * intCodes[intCodes[i+2]]
			intCodes[intCodes[i+3]] = product
			i = i + 4
		} else if opCode == ExitOpcode {
			break
		} else {
		}
	}

	return intCodes[0]
}

func Day2Part2() int {
	file, err := os.Open("./input/day2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var result []string

	for scanner.Scan() {
		var listOfOpcodes = scanner.Text()
		result = strings.Split(listOfOpcodes, ",")
	}

	var intCodes []int

	//Convert to array of ints
	for _, element := range result {
		var opcode, err = strconv.Atoi(element)

		if err != nil {
			log.Fatal(err)
		}

		intCodes = append(intCodes, opcode)
	}

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copyOfIntCodes := make([]int, 0)
			copyOfIntCodes = append(copyOfIntCodes, intCodes...)
			copyOfIntCodes[1] = noun
			copyOfIntCodes[2] = verb

			var i = 0

			for {
				var opCode = copyOfIntCodes[i]

				if opCode == AddOpcode {
					var sum = copyOfIntCodes[copyOfIntCodes[i+1]] + copyOfIntCodes[copyOfIntCodes[i+2]]
					copyOfIntCodes[copyOfIntCodes[i+3]] = sum
					if i+4 <= len(copyOfIntCodes)-1 {
						i = i + 4
					} else {
						break
					}
				} else if opCode == MultOpcode {
					var product = copyOfIntCodes[copyOfIntCodes[i+1]] * copyOfIntCodes[copyOfIntCodes[i+2]]
					copyOfIntCodes[copyOfIntCodes[i+3]] = product

					if i+4 <= len(copyOfIntCodes)-1{
						i = i + 4
					} else {
						break
					}
				} else if opCode == ExitOpcode {
					break
				} else {
					break
				}
			}

			if copyOfIntCodes[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return -1
}
