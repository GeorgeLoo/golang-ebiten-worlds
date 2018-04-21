


/*

worlds.go by George Loo 27.1.2018

mouse
keyboard
messaging


jj
*/
/**/

package main

import (
	"fmt"
	"image/color"
	"github.com/hajimehoshi/ebiten"
)

const (
	screenwidth = 800
	screenheight = 400
)

/*
where in the world
what objects in this sector
locations of these objects
characteristics - can be cover or not
doors
glass
roads
walls
wood


*/
type worldtype struct {
	x, y float64
	msg string 
}

type baseObject struct {
    x,y float64
    name string
    h,w int
	image      *ebiten.Image
    
}

var (
    mousedownState bool
    city []worldtype 
    city2 [][]worldtype
    oldmousex,oldmousey int
    hero baseObject
)

func InitProg() {


    hero.init(10,10,50,20,"hero")

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
	city2[9][0].msg = "python"
	city2[5][5].msg = "golang"
	city2[6][5].msg = "neither"
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


func (p *baseObject) init(x float64, y float64,h int, w int, name string) {

    blue := color.NRGBA{0x00, 0x00, 0xff, 0xff}
    p.image, _ =  ebiten.NewImage(w, h, ebiten.FilterNearest)
    p.image.Fill(blue)

}

func (p *baseObject) draw(x float64, y float64, screen *ebiten.Image) {

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)
   	screen.DrawImage(p.image, opts)
    
}


func mouseLeftdown(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()
	fmt.Print(mx,",",my," mousedown \n")
}

func mouseLeftup(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()
	fmt.Print(mx,",",my," mouseup \n")

}

func mouseWithin(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()

	if mx != oldmousex && my != oldmousey {
		fmt.Print(mx,",",my," mouse within \n")	
	} 
	
	oldmousex,oldmousey = mx, my

}

func mouseRightDown(screen *ebiten.Image) {
	mx, my := ebiten.CursorPosition()
	fmt.Print(mx,",",my," mouse right down \n")

}

func update(screen *ebiten.Image) error {


    backcolor := color.NRGBA{175, 215, 122, 0xff}
    screen.Fill(backcolor)

	if ebiten.IsRunningSlowly() {
		return nil
		//fmt.Print("running slowly! \n")
	}

    hero.draw(100,100, screen)

	mouseWithin(screen)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		mouseRightDown(screen)
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
    if err := ebiten.Run(update, screenwidth, screenheight, scale, "Worlds 0.0 by George Loo"); err != nil {
      panic(err)
    }
    fmt.Printf("Program ended -----------------\n")

}
