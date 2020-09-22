package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"math"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	colors "golang.org/x/image/colornames"
	"gonum.org/v1/plot/tools/bezier"
	"gonum.org/v1/plot/vg"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	offsetX      = 100
	offsetY      = 100
	scaleX       = 700
	scaleY       = 500
)

func loadPicture(pic string) (pixel.Picture, error) {
	file, err := os.Open(pic)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func run() {
	// Create a new window.
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1280, 720),
		VSync:  true,
		//Undecorated: true,
		//Monitor: pixelgl.PrimaryMonitor(),
	}
	win, err := pixelgl.NewWindow(cfg)
	handleError(err)

	// Create a sprite out of the loaded picture.
	picture, err := loadPicture("suika.jpg")
	handleError(err)
	sprite := pixel.NewSprite(picture, pixel.R(0, 0,
		picture.Bounds().Max.X, picture.Bounds().Max.Y))
	transform := pixel.IM.Moved(pixel.V(100, 100))

	// Create a new transition.
	translation1 := &translationNode{
		passed:      0,
		duration:    5 * time.Second,
		start:       pixel.V(100, 100),
		destination: pixel.V(600, 600),
	}
	translation2 := &translationNode{
		passed:      0,
		duration:    7 * time.Second,
		start:       pixel.V(600, 600),
		destination: pixel.V(1200, 200),
	}
	translation3 := &translationNode{
		passed:      0,
		duration:    3 * time.Second,
		start:       pixel.V(1200, 200),
		destination: pixel.V(100, 100),
	}

	points := []vg.Point{
		{X: 0*scaleX + offsetX, Y: 0*scaleY + offsetY},
		{X: 1.989*scaleX + offsetX, Y: 0.263*scaleY + offsetY},
		{X: 0.172*scaleX + offsetX, Y: 1.24*scaleY + offsetY},
		{X: 1.927*scaleX + offsetX, Y: 1.004*scaleY + offsetY},
	}
	bezier1 := &bezierNode{
		passed:   0,
		duration: 5 * time.Second,
		curve:    bezier.New(points...),
	}

	rotation1 := &rotationNode{
		passed:     0,
		duration:   3 * time.Second,
		startAngle: 0,
		destAngle:  math.Pi / 2,
	}
	rotation2 := &rotationNode{
		passed:     0,
		duration:   7 * time.Second,
		startAngle: math.Pi / 2,
		destAngle:  0,
	}

	transition := &Transition{
		transform: &transform,
		tracks:    []*transitionTrack{},
	}

	err = transition.AddTrack("movement")
	handleError(err)

	err = transition.AddTransitionNode("movement", translation1)
	handleError(err)
	err = transition.AddTransitionNode("movement", translation2)
	handleError(err)
	err = transition.AddTransitionNode("movement", translation3)
	handleError(err)
	err = transition.AddTransitionNode("movement", bezier1)

	err = transition.AddTrack("rotation")
	handleError(err)

	err = transition.AddTransitionNode("rotation", rotation1)
	handleError(err)
	err = transition.AddTransitionNode("rotation", rotation2)
	handleError(err)

	last := time.Now()
	fps := 0
	perSecond := time.Tick(time.Second)

	for !win.Closed() {
		deltaTime := time.Since(last)
		last = time.Now()

		win.Clear(colors.White)

		transition.Update(deltaTime)

		sprite.Draw(win, transform)

		win.Update()

		// Show FPS in the window title.
		fps++

		select {
		case <-perSecond:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d",
				cfg.Title, fps))
			fps = 0

		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
