package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	//"github.com/faiface/pixel/imdraw"
)

func make_window() {

}
func test_print() {
	fmt.Println("I'm in here")
}

func run() { //this effectively is our new main function for this pkg
	cfg := pixelgl.WindowConfig{
		Title:     "Game Of Life!",
		Bounds:    pixel.R(0, 0, 1024, 768),
		VSync:     true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Whitesmoke)
	for !win.Closed() {
		win.Update()
	}

}
func main() {
	pixelgl.Run(run)

}
