package main

import (
	"fmt"
	"strings"
    
	//"math/rand"
)

const w, h = 10, 5
// GameBoard defines the boards initial width and height
type GameBoard struct {
	width, height int
	cells         string
}

func main() {
	board := &GameBoard{width: w, height: h, cells: "O"}
	// fmt.Println(board.cells)
	// board.MakeLife()
	// fmt.Println(board.cells)
	board.PrintBoard()
}

// MakeLife is associated with GameBoard and gives life to the cells
func (gb *GameBoard) MakeLife() {
	gb.cells = "X"
}

//NewGameBoard returns an initialized GameBoard
func NewGameBoard(w, h int, c string) *GameBoard {
	return &GameBoard{width: w, height: h, cells: c}
}

func (gb *GameBoard) PrintBoard() {
	for i := 0; i < gb.height; i++{
		x := strings.Repeat(gb.cells, gb.width)
		fmt.Println(x)
	}
}



// func BoardArray(w,h int) [][]string{
// 	returng
// }
