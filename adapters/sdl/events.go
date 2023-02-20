package sdl

import (
	"time"

	player "github.com/mniak/japlayer"
	"github.com/veandco/go-sdl2/sdl"
)

func HandleEvents(h player.EventHandler) {
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			// fmt.Println(event)
			switch event.(type) {
			case *sdl.QuitEvent:
				if h.Quit != nil {
					h.Quit()
				}
				return
			}
		}
		time.Sleep(1 * time.Second)
	}
}
