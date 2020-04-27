package adventofcode

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	slope      float64
	yIntercept float64
	xIntercept float64
}

//plan of attack: Find all intersections. Calculate Manhattan distance for each.

func Day3() int {
	file, err := os.Open("./input/day3.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	result := make([][]string, 1000)

	index := 0

	for scanner.Scan() {
		result[index] = strings.Split(scanner.Text(), ",")
		index++
	}

	//2d array of directions
	fmt.Println(result)
	
	points := [][]point{{{
		x: 0,
		y: 0,
	}}, {{x: 0, y: 0}}}

	//create points
	//points := make([][]point, len(result))
	//origin := point{x:0, y: 0}
	//points[0][0] = origin
	//points[1][0] = origin

	for i, line := range result {
		for j, move := range line {
			if !(i == 0 && (j == 0 || j == 1)) {
				direction := move[0]
				numOfSteps, err := strconv.Atoi(move[1:])

				if err != nil {
					log.Fatal(err)
				}

				if direction == 'R' {
					//changes y
					points[i][j] = point{x: points[i][j-1].x, y:numOfSteps }
				} else if direction == 'U' {
					//changes x
					points[i][j] = point{x:numOfSteps, y:points[i][j-1].y}
				} else if direction == 'L' {
					//changes y
					points[i][j] = point{x: points[i][j-1].x, y:-numOfSteps }
				} else if direction == 'D' {
					//changes x
					points[i][j] = point{x:-numOfSteps, y:points[i][j-1].y}
				}
			}
		}
	}
	fmt.Println(points)

	//create lines
	//lines := make([][]line, len(result))
	//
	//k := 0
	//for i:=0; i < len(points); i++ {
	//	for j := 0; j < len(points[i]); j = j + 2 {
	//		lines[i][k] = createLine(points[i][j], points[i][j+1])
	//		k++
	//	}
	//}
	//
	//for _, line := range lines {
	//	fmt.Println(line)
	//}

	return 1
}

func createLine(pt1 point, pt2 point) line {

	if pt2.x-pt1.x == 0 { return line{xIntercept: float64(pt2.x)}}

	slope := float64(pt2.y - pt1.y) / float64(pt2.x - pt1.x)

	intercept := -slope*float64(pt1.x) + float64(pt1.y)

	result := line{slope: slope, yIntercept:intercept, xIntercept:0}

	return result
}

func calculateManhattanDistance(pt point) int {
	return 1
}


