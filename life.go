package main

import (
	"fmt"
	
	//"math/rand"
)

const w, h = 10, 30
// GameBoard defines the boards initial width and height
type GameBoard struct {
	width, height int
	cells         string
}

func main() {
	board := &GameBoard{width: w, height: h, cells: "=="}
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
	fmt.Println("I get called")
	return &GameBoard{width: w, height: h, cells: c}
}

// PrintBoard does a thing
func (gb *GameBoard) PrintBoard() {
	board := make([][]string,gb.width*gb.height)
	for i := range board {
		board[i] = append(board[i],gb.cells)
	}
	

	fmt.Println(board[0][0])



}



// func BoardArray(w,h int) [][]string{
// 	returng
// }
