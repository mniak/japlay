package sdl

import (
	player "github.com/mniak/japlayer"
)

var (
	_ player.AudioPlayer = &audioAdapter{}
	_ player.AudioPlayer = &sdlAdapter{}
)
