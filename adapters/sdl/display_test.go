package sdl

import player "github.com/mniak/japlayer"

var (
	_ player.Display = &displayAdapter{}
	_ player.Display = &sdlAdapter{}
)
