package sdl

import (
	"github.com/mniak/louvorja/player"
)

var _ player.AudioPlayer = &sdlAdapter{}
