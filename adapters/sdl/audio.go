package sdl

import (
	"github.com/veandco/go-sdl2/mix"
)

func (ad *sdlAdapter) LoadAudio(audiofile string) error {
	err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096)
	if err != nil {
		return err
	}

	ad.music, err = mix.LoadMUS(audiofile)
	if err != nil {
		return err
	}
	return nil
}

func (ad *sdlAdapter) Play() error {
	err := ad.music.Play(0)
	if err != nil {
		return err
	}
	return nil
}

// func (ad *sdlAdapter) Seek(time time.Duration) error {
// 	err := mix.SetMusicPosition(int64(time.Seconds()))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
