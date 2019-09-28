package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const x, y = 150, 40 // consts for height and width of board
const lifePct = 0.30 // % of life the board starts off with 

// Board defines the boards initial width and height
type Board struct {
	x, y  int // x = width, y = height
	cells [][]bool
}

// NewBoard returns an initialized Board
// with width w and height h and 2d array of bools is
// populated on each row with a slice of bools
func NewBoard(x, y int) *Board {
	b := &Board{x: x, y: y}
	b.cells = make([][]bool, y)

	for i := 0; i < y; i++ {
		b.cells[i] = make([]bool, x)
	}

	return b

}

// IntializeRandomLife sets indices of the board
// that are to be randomly given life on creation
func (b *Board) IntializeRandomLife() Board {

	rand.Seed(time.Now().UnixNano())
	totalCells := b.x * b.y                                        // total cells in 2d array
	livingCells := int(math.Ceil((lifePct * float64(totalCells)))) //number of cells to make alive

	for i := 0; i < livingCells; i++ {

		randX := rand.Intn(b.x)
		randY := rand.Intn(b.y)

		// while this index remains true, try another
		for b.cells[randY][randX] {
			randX = rand.Intn(b.x)
			randY = rand.Intn(b.y)
		}

		b.cells[randY][randX] = true

	}

	return *b

}

// Tick implements the rules of Conways' game of life
// for each generation by testing against certain conditions
func (b *Board) Tick() {

	cb := NewBoard(b.x, b.y)

	cb.cells = make([][]bool, len(b.cells))
	for i := range b.cells {
		cb.cells[i] = make([]bool, len(b.cells[i]))
		copy(cb.cells[i], b.cells[i])
	}
	for y := 0; y < b.y; y++ {
		for x := 0; x < b.x; x++ {
			switch {
			//less than two live neighbours it dies by underpopulation
			case cb.cells[y][x] && cb.CheckNeighbours(x, y) < 2:
				b.cells[y][x] = false
			//more than three live neighbours it dies of overpopulation
			case cb.cells[y][x] && cb.CheckNeighbours(x, y) > 3:
				b.cells[y][x] = false
			// dead with exactly three live neigbours comes to life by reproduction
			case !cb.cells[y][x] && cb.CheckNeighbours(x, y) == 3:
				b.cells[y][x] = true
			//two or three live neighbours it lives to next generation
			case cb.cells[y][x] && cb.CheckNeighbours(x, y) == 2 || cb.CheckNeighbours(x, y) == 3:
				b.cells[y][x] = true

			default:
				b.cells[y][x] = false
			}
		}
	}
}

// CheckNeighbours tests the 8 possible surrounding neighbours
// of a cell and determines if the index is legal
// and if the cell is living, in which case livingcount is incremented
func (b *Board) CheckNeighbours(x, y int) int {
	livingCount := 0

	if b.CheckLegalRange(y-1, x) {
		livingCount++
	}

	if b.CheckLegalRange(y+1, x) {
		livingCount++
	}

	if b.CheckLegalRange(y, x-1) {
		livingCount++
	}

	if b.CheckLegalRange(y, x+1) {
		livingCount++
	}

	if b.CheckLegalRange(y-1, x-1) {
		livingCount++
	}

	if b.CheckLegalRange(y-1, x+1) {
		livingCount++
	}

	if b.CheckLegalRange(y+1, x-1) {
		livingCount++
	}
	if b.CheckLegalRange(y+1, x+1) {
		livingCount++
	}
	return livingCount
}

// CheckLegalRange makes sure that the index being tested
// is legal on our board
func (b *Board) CheckLegalRange(y, x int) bool {
	return (x >= 0 && x < b.x && y >= 0 && y < b.y && b.cells[y][x] == true)

}
// PrintBoard displays the current state of the board graphically in the CLI
func (b *Board) PrintBoard() {
	fmt.Print("+")
	for x := 1; x <= b.x; x++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	for y := 0; y < b.y; y++ {
		fmt.Print("|")
		for x := 0; x < b.x; x++ {
			if b.cells[y][x] {
				fmt.Print("▓")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}

	fmt.Print("+")
	for x := 1; x <= b.x; x++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func main() {
	n := NewBoard(x, y)
	n.IntializeRandomLife()
	for i := 0; i > -1; i++ {
		
		clrCmd := exec.Command("clear")
		clrCmd.Stdout = os.Stdout
		clrCmd.Run()
		n.Tick()
		n.PrintBoard()
		fmt.Printf("Current Gen:%v",i)
		time.Sleep(time.Second / 20)

	}

}
