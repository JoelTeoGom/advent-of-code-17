package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
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
		directions: make(map[string]Coord),             // Inicializamos el mapa vacío
		count:      0,                                  // Inicializamos el contador en 0
		start:      Coord{posX: 0, posY: 0},            // Coordenada de inicio (0, 0)
		end:        Coord{posX: rows - 1, posY: i - 1}, // Coordenada de fin (10, 10)
	}

	stats.directions["^"] = Coord{posX: -1, posY: 0}
	stats.directions["v"] = Coord{posX: 1, posY: 0}
	stats.directions[">"] = Coord{posX: 0, posY: 1}
	stats.directions["<"] = Coord{posX: 0, posY: -1}

	fmt.Println(stats)
}

func main() {
	filename := "input.txt"
	m := make([][]string, 0)
	stats := Stats{}

	readfile(&stats, &m, filename)

	//clonando la matriz
	temporal := make([][]string, len(m))
	for i := range m {
		temporal[i] = make([]string, len(m[i]))
		copy(temporal[i], m[i])
	}

	path := make([]string, 1)
	pathIndex := 1

	var symbol string
	var block, end bool

	minHeat := 10000000000
	var suma, number int

	//algorithm

	for {

		//reset
		currentPos := Coord{posX: 0, posY: 0}
		block, end = false, false
		suma = 0

		for i := range temporal {
			copy(m[i], temporal[i])
		}

		for !end && !block {
			random(&symbol)
			//fmt.Println(symbol)
			switch symbol {

			case "^":
				//fmt.Println("SYMBOL: ", symbol)

				if path[pathIndex-1] != "v" && count(path, symbol, pathIndex) < 3 && isValid(currentPos, stats, symbol, m) {
					//	fmt.Println("ENTRAAAA")
					currentPos.posX += stats.directions[symbol].posX
					currentPos.posY += stats.directions[symbol].posY
					number, _ = strconv.Atoi(m[currentPos.posX][currentPos.posY])
					suma += number
					m[currentPos.posX][currentPos.posY] = symbol
					path = append(path, symbol)
					pathIndex++
				}

			case "v":
				//fmt.Println("SYMBOL: ", symbol)
				if path[pathIndex-1] != "^" && count(path, symbol, pathIndex) < 3 && isValid(currentPos, stats, symbol, m) {
					//	fmt.Println("ENTRAAAA")

					currentPos.posX += stats.directions[symbol].posX
					currentPos.posY += stats.directions[symbol].posY

					number, _ = strconv.Atoi(m[currentPos.posX][currentPos.posY])
					suma += number
					m[currentPos.posX][currentPos.posY] = symbol
					path = append(path, symbol)
					pathIndex++
				}

			case ">":
				//fmt.Println("SYMBOL: ", symbol)

				if path[pathIndex-1] != "<" && count(path, symbol, pathIndex) < 3 && isValid(currentPos, stats, symbol, m) {
					//fmt.Println("ENTRAAAA")

					currentPos.posX += stats.directions[symbol].posX
					currentPos.posY += stats.directions[symbol].posY

					number, _ = strconv.Atoi(m[currentPos.posX][currentPos.posY])
					suma += number
					m[currentPos.posX][currentPos.posY] = symbol
					path = append(path, symbol)
					pathIndex++
				}
			case "<":
				//fmt.Println("SYMBOL: ", symbol)

				if path[pathIndex-1] != ">" && count(path, symbol, pathIndex) < 3 && isValid(currentPos, stats, symbol, m) {
					//fmt.Println("ENTRAAAA")

					currentPos.posX += stats.directions[symbol].posX
					currentPos.posY += stats.directions[symbol].posY

					number, _ = strconv.Atoi(m[currentPos.posX][currentPos.posY])
					suma += number
					m[currentPos.posX][currentPos.posY] = symbol
					path = append(path, symbol)
					pathIndex++
				}
			default:
				fmt.Println("Símbolo no reconocido") // Caso por defecto si no coincide ningún símbolo
			}

			block = isblock(m, currentPos)

			if currentPos.posX == stats.end.posX && currentPos.posY == stats.end.posY {

				if suma < minHeat {
					minHeat = suma
				}

				end = true
			}

		}
		// for i := 0; i < len(m); i++ {
		// 	for j := 0; j < len(m); j++ {
		// 		fmt.Printf("%s", m[i][j])
		// 	}
		// 	fmt.Println()
		// }

		fmt.Println("SUMA: ", minHeat)

	}
}

func isblock(m [][]string, currentPos Coord) bool {
	x, y := currentPos.posX, currentPos.posY
	rows := len(m)
	cols := len(m[0])

	// Función auxiliar para verificar si una coordenada está dentro de los límites de la matriz
	isValid := func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}

	// Función auxiliar para comprobar si una celda contiene un número
	isNumber := func(s string) bool {
		if len(s) != 1 {
			return false
		}
		return s >= "0" && s <= "9"
	}

	// Verificar las posiciones adyacentes (arriba, abajo, izquierda, derecha)

	// Arriba
	if isValid(x-1, y) && isNumber(m[x-1][y]) {
		return false
	}

	// Abajo
	if isValid(x+1, y) && isNumber(m[x+1][y]) {
		return false
	}

	// Izquierda
	if isValid(x, y-1) && isNumber(m[x][y-1]) {
		return false
	}

	// Derecha
	if isValid(x, y+1) && isNumber(m[x][y+1]) {
		return false
	}

	// Si todas las posiciones adyacentes no contienen números o están fuera de los límites, está bloqueado
	return true
}

func isValid(currentPos Coord, stats Stats, symbol string, m [][]string) bool {
	x := currentPos.posX + stats.directions[symbol].posX
	y := currentPos.posY + stats.directions[symbol].posY
	// fmt.Println("[X, Y]: ", x, y)
	// fmt.Println(stats.end.posX, stats.end.posY)
	if x >= 0 && x <= stats.end.posX && y >= 0 && y <= stats.end.posY {
		if isNumber(m, x, y) {
			// fmt.Println("ISVALID")
			// fmt.Println("Matrix[x,y]: ", m[x][y])
			return true
		}

	}
	//fmt.Println("IS_NOT_VALID")

	return false
}

func isNumber(m [][]string, x, y int) bool {

	if m[x][y] < "0" || m[x][y] > "9" {
		return false
	}

	//fmt.Println("isNUMBER")

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

	//fmt.Println("count: ", count)
	return count
}

func random(symbol *string) {
	// Generate a random number between 0 and 3
	symbols := [4]string{"<", ">", "^", "v"}
	rangeSize := big.NewInt(4)
	n, err := rand.Int(rand.Reader, rangeSize)
	option := int(n.Int64())
	if err != nil {
		//fmt.Println("Error generating random number:", err)
		return
	}

	*symbol = symbols[option]

}
