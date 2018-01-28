


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


type worldtype struct {
	x, y float64
	msg string 
}

var (
  mousedownState bool
  city []worldtype 
  city2 [][]worldtype

)

func InitProg() {

	city = make([]worldtype,5)
	city[0].x = 1
	city[0].msg = "George"
	city[1].x = 23
	city[1].msg = "Linda"
	city[4].msg = "Audrey"
	city[3].msg = "foo"

	i := 0
	for i < 5 {
		fmt.Print(city[i].x,",",city[i].msg," jibai \n")
		i += 1
	}

	city2 = make([][]worldtype,10)
	for i=0;i<10;i++ {
    	city2[i] = make([]worldtype, 10)  // 
	}
	city2[0][0].msg = "gopher"
	city2[9][0].msg = "codemotherfucker"
	city2[5][5].msg = "Singapore"
	city2[6][5].msg = "Screwed"
	city2[6][5].msg = "Saved"
	i = 0

	for i < 10 {
		j := 0
		for j < 10 {
			fmt.Print(i,",",j," = ",city2[i][j].msg,"/")
			j += 1
		}
		fmt.Print("\n")
		i += 1
	}

}

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


	InitProg()

    scale := 1.0
    // Initialize Ebiten, and loop the update() function
    if err := ebiten.Run(update, screenwidth, screenheight, scale, "Animation test 0.0 by George Loo"); err != nil {
      panic(err)
    }
    fmt.Printf("Program ended -----------------\n")

}
