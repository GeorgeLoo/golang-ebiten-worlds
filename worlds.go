


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
    kMoveLeft = 1001
    kMoveRight = 1002
    kMoveUp = 1003
    kMoveDown = 1004
    kMoveStop = 1005
    
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
    direction int
    
}

var (
    mousedownState bool
    city []worldtype 
    city2 [][]worldtype
    oldmousex,oldmousey int
    hero baseObject

	keyStates    = map[ebiten.Key]int{
		ebiten.KeyUp:    0,
		ebiten.KeyDown:  0,
		ebiten.KeyLeft:  0,
		ebiten.KeyRight: 0,
		ebiten.KeyA:     0,
		ebiten.KeyS:     0,
		ebiten.KeyW:     0,
		ebiten.KeyD:     0,
	}
)

func controls(screen *ebiten.Image) {


   	for key := range keyStates {
    	if !ebiten.IsKeyPressed(key) {
			keyStates[key] = 0
			continue
		}
		keyStates[key]++
	}

    if keyStates[ebiten.KeyA] == 1 {
        fmt.Println("A key")
        hero.move(kMoveLeft)

    }
    if keyStates[ebiten.KeyW] == 1 {
        fmt.Println("W up key")
        hero.move(kMoveUp)
    }
    if keyStates[ebiten.KeyS] == 1 {
        fmt.Println("S down key")
        hero.move(kMoveDown)
    }
    if keyStates[ebiten.KeyD] == 1 {
        
        hero.move(kMoveRight)
    }


}

func movement(screen *ebiten.Image) {


    hero.draw(100,100, screen)


}

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



func (p *baseObject)  move(direction int) {

    if direction == kMoveUp {
        p.direction = kMoveUp

    } else if direction == kMoveDown {
        p.direction = kMoveDown

    } else if direction == kMoveLeft {
        p.direction = kMoveLeft
    } else if direction == kMoveRight {
        fmt.Println("move right key")
        p.direction = kMoveRight
    }

}


func (p *baseObject) init(x float64, y float64,h int, w int, name string) {

    p.direction = kMoveStop
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
		//fmt.Print(mx,",",my," mouse within \n")	
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

    controls(screen)

    

    movement(screen)

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
