package main

import (
	"math/rand"
	"path"
	"runtime"
	"time"

	"github.com/yuuna-stack/go_doodle_jump/wrapper"

	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

const resourcesDir = "images"

type Point struct {
	x int
	y int
}

func init() { runtime.LockOSThread() }

func fullname(filename string) string {
	return path.Join(resourcesDir, filename)
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	resources := wrapper.Resources{}

	const gameWidth = 400
	const gameHeight = 533

	option := uint(window.SfResize | window.SfClose)
	wnd := wrapper.CreateWindow(gameWidth, gameHeight, "Doodle Game!", option, 60)

	sBackground, err := wrapper.FileToSprite(fullname("background.png"), &resources)
	if err != nil {
		panic("Couldn't load background.png")
	}
	sPlat, err := wrapper.FileToSprite(fullname("platform.png"), &resources)
	if err != nil {
		panic("Couldn't load platform.png")
	}
	sPers, err := wrapper.FileToSprite(fullname("doodle.png"), &resources)
	if err != nil {
		panic("Couldn't load doodle.png")
	}

	var plat [20]Point

	for i := 0; i < 10; i++ {
		plat[i].x = r1.Int() % gameWidth
		plat[i].y = r1.Int() % gameHeight
	}

	x := 100
	y := 100
	h := 200
	dy := 0.0

	for wnd.IsOpen() {
		for wnd.Poll_Event() {
			if wnd.Close_Window() {
				return
			}
			if wnd.Key_Pressed() {
				if wnd.Key_Is(window.SfKeyLeft) {
					x -= 3
				} else if wnd.Key_Is(window.SfKeyRight) {
					x += 3
				}
			}
		}

		dy += 0.2
		y += int(dy)
		if y > gameHeight {
			dy -= 10
		}

		if y < h {
			for i := 0; i < 10; i++ {
				y = h
				plat[i].y -= int(dy)
				if plat[i].y > gameHeight {
					plat[i].y = 0
					plat[i].x = r1.Int() % gameWidth
				}
			}
		}

		for i := 0; i < 10; i++ {
			if x+50 > plat[i].x && x+20 < plat[i].x+68 &&
				y+70 > plat[i].y && y+70 < plat[i].y+14 && dy > 0 {
				dy -= 10
			}
		}

		sPers.SetPosition(float32(x), float32(y))

		wnd.Clear_Window(graphics.GetSfBlack())

		sBackground.Draw(wnd.Get_Window())

		sPers.Draw(wnd.Get_Window())

		for i := 0; i < 10; i++ {
			sPlat.SetPosition(float32(plat[i].x), float32(plat[i].y))
			sPlat.Draw(wnd.Get_Window())
		}

		graphics.SfRenderWindow_display(wnd.Get_Window())
	}

	resources.Clear()
	wnd.Clear()
}
