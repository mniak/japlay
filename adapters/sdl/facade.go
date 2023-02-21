package sdl

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var initialized bool

func Init() error {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		return err
	}

	err = ttf.Init()
	if err != nil {
		sdl.Quit()
		return err
	}
	initialized = true
	return nil
}

func Quit() {
	ttf.Quit()
	sdl.Quit()
}

type AdapterParams struct {
	FontPath   string
	FontSize   float32
	LetterCase LetterCase
	Display    int
}

type sdlAdapter struct {
	*displayAdapter
	*audioAdapter
}

func SDLAdapter(params AdapterParams) (*sdlAdapter, error) {
	display, err := DisplayAdapter(params)
	if err != nil {
		return nil, err
	}
	audio, err := AudioAdapter(params)
	if err != nil {
		return nil, err
	}
	return &sdlAdapter{
		displayAdapter: display,
		audioAdapter:   audio,
	}, nil
}

func (ad *sdlAdapter) Finish() {
	ad.displayAdapter.Finish()
	ad.audioAdapter.Finish()
}
