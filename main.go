package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

const maxCount = 3

type Stats struct {
	count int
	start Coord
	end   Coord
}

type Coord struct {
	posX, posY int
}

type Direction struct {
	up, down, right, left bool
}

func readfile(stats *Stats, m *[][]string, filename string) {

	file, err := os.Open(filename)

	if err != nil {
		return
	}

	defer file.Close()
	var line string
	scanner := bufio.NewScanner(file)
	var row int
	for scanner.Scan() {
		line = scanner.Text()
		row = 0
		*m = append(*m, make([]string, len(line)))

		for _, char := range line {
			if row < len(*m) {
				(*m)[row] = append((*m)[row], string(char))
			}
			row++
		}
	}
	stats.count = 0
	stats.end = Coord{posX: row - 1, posY: len(line) - 1}
	stats.start = Coord{posX: 0, posY: 0}

	fmt.Println(stats)
}

func main() {

	m := make([][]string, 0)
	filename := "input.txt"
	stats := Stats{}
	readfile(&stats, &m, filename)
	var posX, posY int
	minHeat := 10000
	var suma int
	//algorithm

	for {
		currentPos := Coord{posX: 0, posY: 0}
		suma = 0

		lastDirection := Direction{
			up:    false,
			down:  false,
			right: false,
			left:  false,
		}

		currentDirection := Direction{
			up:    false,
			down:  false,
			right: false,
			left:  false,
		}

		//algorithm start

		for currentPos.posX != stats.end.posX && currentPos.posY != stats.end.posY {

			//select direction
			var option int
			random(&option)

			posX, posY = 0, 0

			switch option {
			case 0:
				if !lastDirection.down && stats.count < 3 {
					currentDirection.up = true
					posX = -1
				}

				if !lastDirection.up {
					stats.count++
				}

				break
			case 1:
				if !lastDirection.down && stats.count < 3 {
					currentDirection.down = true
					posX = +1
				}

				if !lastDirection.down {
					stats.count++
				}
				break
			case 2:
				if !lastDirection.left && stats.count < 3 {
					currentDirection.right = true
					posY = +1
				}

				if !lastDirection.right {
					stats.count++
				}

				break
			case 3:
				if !lastDirection.right && stats.count < 3 {
					currentDirection.left = true
					posY = -1
				}

				if !lastDirection.left {
					stats.count++
				}
				break
			}

			if inRange(currentPos, posX, posY, stats.end.posX, stats.start.posY, m) {

			}

		}

	}

}

func random(option *int) {
	// Generate a random number between 0 and 3
	rangeSize := big.NewInt(4)
	n, err := rand.Int(rand.Reader, rangeSize)
	*option = int(n.Int64())
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return
	}
}

func inRange(currentPos Coord, posX, posY, row, col int, m [][]string) bool {
	x := currentPos.posX + posX
	y := currentPos.posY + posY
	if x > 0 && x <= row && y > 0 && y <= col && m[x][y] {
		return true
	}
	return false
}
