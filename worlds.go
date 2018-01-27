


/*

worlds.go by George Loo 27.1.2018

mouse
keyboard
messaging


jj
*/


package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
)

const (
	screenwidth = 800
	screenheight = 400
)

var (
  mousedownState bool
)


func mouseLeftdown(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()
	fmt.Print(mx,",",my," mousedown \n")
}

func mouseLeftup(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()
	fmt.Print(mx,",",my," mouseup \n")
}


func update(screen *ebiten.Image) error {

	if ebiten.IsRunningSlowly() {
		return nil
		//fmt.Print("running slowly! \n")
	}

  if mousedownState {
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			mousedownState = false
			mouseLeftup(screen)
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !mousedownState {
			mousedownState = true
		}
		mouseLeftdown(screen)
	}

  return nil

}

func main() {

    scale := 1.0
    // Initialize Ebiten, and loop the update() function
    if err := ebiten.Run(update, screenwidth, screenheight, scale, "Animation test 0.0 by George Loo"); err != nil {
      panic(err)
    }
    fmt.Printf("Program ended -----------------\n")

}
