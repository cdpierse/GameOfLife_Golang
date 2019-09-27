package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/kr/pretty"
	//"math/rand"
)

const w, h = 10, 15
const lifePct = 0.50

// Board defines the boards initial width and height
type Board struct {
	width, height int
	cells         [][]bool
}

// NewBoard returns an initialized Board
// with width w and height h and 2d array of bools is
// populated on each row with a slice of bools
func NewBoard(w, h int) *Board {
	b := &Board{width: w, height: h}
	b.cells = make([][]bool, h)

	for i := 0; i < h; i++ {
		b.cells[i] = make([]bool, w)
	}
	IntializeRandomLife(b)

	return b

}

// IntializeRandomLife sets indices of the board
// that are to be randomly given life on creation
func IntializeRandomLife(b *Board) Board {
	rand.Seed(time.Now().UnixNano())
	totalCells := b.width * b.height                               // total cells in 2d array
	livingCells := int(math.Ceil((lifePct * float64(totalCells)))) //number of cells to make alive

	for i := 0; i < livingCells; i++ {

		randI := rand.Intn(b.height)
		randJ := rand.Intn(b.width)

		// while this index remains true, try another
		for b.cells[randI][randJ] == true { 
			randI = rand.Intn(b.height)
			randJ = rand.Intn(b.width)
		}

		b.cells[randI][randJ] = true

	}

	return *b

}

func main() {
	n := NewBoard(w, h)
	pretty.Println(n.cells)
}

// func main() {
// 	board := &Board{width: w, height: h, cells: "=="}
// 	// fmt.Println(board.cells)
// 	// board.MakeLife()
// 	// fmt.Println(board.cells)
// 	board.PrintBoard()
// }

// // MakeLife is associated with Board and gives life to the cells
// func (gb *Board) MakeLife() {
// 	gb.cells = "X"
// }

// // PrintBoard does a thing
// func (gb *Board) PrintBoard() {
// 	board := make([][]string,gb.width*gb.height)
// 	for i := range board {
// 		board[i] = append(board[i],gb.cells)
// 	}

// 	fmt.Println(board[0][0])

// }

// func BoardArray(w,h int) [][]string{
// 	returng
// }
