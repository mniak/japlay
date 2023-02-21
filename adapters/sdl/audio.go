package sdl

import (
	"time"

	"github.com/veandco/go-sdl2/mix"
)

type audioAdapter struct {
	music *mix.Music
}

func AudioAdapter(params AdapterParams) (*audioAdapter, error) {
	return &audioAdapter{}, nil
}

func (ad *audioAdapter) Finish() {
	if ad.music != nil {
		ad.music.Free()
	}
}

func (ad *audioAdapter) LoadAudio(audiofile string) error {
	err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096)
	if err != nil {
		return err
	}

	music, err := mix.LoadMUS(audiofile)
	if err != nil {
		return err
	}

	if ad.music != nil {
		ad.music.Free()
	}
	ad.music = music

	return nil
}

func (ad *audioAdapter) Play() error {
	err := ad.music.Play(0)
	if err != nil {
		return err
	}
	return nil
}

// func (ad *audioAdapter) Seek(time time.Duration) error {
// 	err := mix.SetMusicPosition(int64(time.Seconds()))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (ad *audioAdapter) Wait() {
	for mix.PlayingMusic() {
		time.Sleep(100 * time.Millisecond)
	}
}
