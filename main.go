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

	minHeat := 10000

	//algorithm

	for {
		currentPos := Coord{posX: 0, posY: 0}
		suma := 0
		
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

			posX, posY := 0, 0
			switch option {
			case 0:

				if !lastDirection.down && count < 3{
					currentDirection.up = true
					posX = -1
					
					
				} 

				break
			case 1:
				currentDirection.down = true
				posX = +1
				break
			case 2:
				currentDirection.right = true
				posY = +1
				break
			case 3:
				currentDirection.left = true
				posY = -1
				break
			}

			//lets verifty if this is possible

			



			if currentPos.posX + posX && currentPos.posY + 1  


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

func isValid(){

	if currentPos.posX + posX && currentPos.posY + 1  {
	
	}

}
