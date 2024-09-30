package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

type Stats struct {
	directions map[string]Coord
	count      int
	start      Coord
	end        Coord
}

type Coord struct {
	posX, posY int
}

func readfile(stats *Stats, m *[][]string, filename string) {

	file, err := os.Open(filename)

	if err != nil {
		return
	}

	defer file.Close()
	var line string
	scanner := bufio.NewScanner(file)

	rows := 0
	var i int
	for scanner.Scan() {
		line := scanner.Text()

		//important
		*m = append(*m, make([]string, len(line)))
		i = 0
		for col, char := range line {
			(*m)[rows][col] = string(char)
			i++

		}

		rows++
	}

	*stats = Stats{
		directions: make(map[string]Coord),                     // Inicializamos el mapa vacío
		count:      0,                                          // Inicializamos el contador en 0
		start:      Coord{posX: 0, posY: 0},                    // Coordenada de inicio (0, 0)
		end:        Coord{posX: rows - 1, posY: len(line) - 1}, // Coordenada de fin (10, 10)
	}

	stats.directions["^"] = Coord{posX: -1, posY: 0}
	stats.directions["v"] = Coord{posX: 1, posY: 0}
	stats.directions[">"] = Coord{posX: 0, posY: 1}
	stats.directions["<"] = Coord{posX: 0, posY: -1}

	fmt.Println(stats)
}

func main() {

	m := make([][]string, 0)
	filename := "input.txt"
	stats := Stats{}
	readfile(&stats, &m, filename)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			fmt.Printf("MATRIX[%d][%d]: %s\n", i, j, m[i][j])
		}

	}
	path := make([]string, 1)
	pathIndex := 1

	var symbol string
	//var posX, posY int
	var block bool

	// minHeat := 10000
	// var suma int

	//algorithm

	for {
		currentPos := Coord{posX: 0, posY: 0}
		//	suma = 0

		for currentPos.posX != stats.end.posX && currentPos.posY != stats.end.posY && !block {
			random(&symbol)
			fmt.Println(symbol)
			switch symbol {

			case "^":
				fmt.Println("SYMBOL: ", symbol)

				if path[pathIndex-1] != "v" && count(path, symbol, pathIndex) < 3 && isValid(currentPos, stats, symbol, m) {
					fmt.Println("ENTRAAAA")
					currentPos.posX += stats.directions[symbol].posX
					currentPos.posY += stats.directions[symbol].posY
					m[currentPos.posX][currentPos.posY] = symbol
					path = append(path, symbol)
					pathIndex++
				}

			case "v":
				fmt.Println("SYMBOL: ", symbol)
				if path[pathIndex-1] != "^" && count(path, symbol, pathIndex) < 3 && isValid(currentPos, stats, symbol, m) {
					fmt.Println("ENTRAAAA")

					currentPos.posX += stats.directions[symbol].posX
					currentPos.posY += stats.directions[symbol].posY
					m[currentPos.posX][currentPos.posY] = symbol
					path = append(path, symbol)
					pathIndex++
				}

			case ">":
				fmt.Println("SYMBOL: ", symbol)

				if path[pathIndex-1] != "<" && count(path, symbol, pathIndex) < 3 && isValid(currentPos, stats, symbol, m) {
					fmt.Println("ENTRAAAA")

					currentPos.posX += stats.directions[symbol].posX
					currentPos.posY += stats.directions[symbol].posY
					m[currentPos.posX][currentPos.posY] = symbol
					path = append(path, symbol)
					pathIndex++
				}
			case "<":
				fmt.Println("SYMBOL: ", symbol)

				if path[pathIndex-1] != ">" && count(path, symbol, pathIndex) < 3 && isValid(currentPos, stats, symbol, m) {
					fmt.Println("ENTRAAAA")

					currentPos.posX += stats.directions[symbol].posX
					currentPos.posY += stats.directions[symbol].posY
					m[currentPos.posX][currentPos.posY] = symbol
					path = append(path, symbol)
					pathIndex++
				}
			default:
				fmt.Println("Símbolo no reconocido") // Caso por defecto si no coincide ningún símbolo
			}

			for i, char := range path {

				fmt.Printf("path[%d]: %s   ", i, char)

			}

			fmt.Println("\n=============================================")

		}
	}

}

func isValid(currentPos Coord, stats Stats, symbol string, m [][]string) bool {
	x := currentPos.posX + stats.directions[symbol].posX
	y := currentPos.posY + stats.directions[symbol].posY
	fmt.Println("[X, Y]: ", x, y)

	if x >= 0 && x <= stats.end.posX && y >= 0 && y <= stats.start.posY {
		if isNumber(m, x, y) {
			fmt.Println("ISVALID")
			fmt.Println("Matrix[x,y]: ", m[x][y])
			return true
		}

	}
	fmt.Println("IS_NOT_VALID")

	return false
}

func isNumber(m [][]string, x, y int) bool {

	if m[x][y] < "0" || m[x][y] > "9" {
		return false
	}

	fmt.Println("isNUMBER")

	return true

}

func count(path []string, symbol string, index int) int {
	count := 0
	for i := index - 1; i > 0; i-- {
		if path[i] != symbol {
			break
		}
		count++
	}

	fmt.Println("count: ", count)
	return count
}

func random(symbol *string) {
	// Generate a random number between 0 and 3
	symbols := [4]string{"<", ">", "^", "v"}
	rangeSize := big.NewInt(4)
	n, err := rand.Int(rand.Reader, rangeSize)
	option := int(n.Int64())
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return
	}

	*symbol = symbols[option]

}
