package player

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/mniak/japlayer/log"
)

type Player struct {
	SongLoader  SongLoader
	Display     Display
	AudioPlayer AudioPlayer

	ImagesDir string
	MusicDir  string

	stopChan chan struct{}
}

const (
	hasd2022 = 712
)

func (p *Player) Stop() {
	if p.stopChan != nil {
		close(p.stopChan)
		p.stopChan = nil
	}
}

func (p *Player) PresentLyrics(hymnNumber int) error {
	p.stopChan = make(chan struct{})

	song, err := p.SongLoader.AlbumTrack(hasd2022, hymnNumber)
	if err != nil {
		return err
	}

	imgPath := filepath.Join(p.ImagesDir, song.CoverImage)
	err = p.Display.SetBackgroundImage(imgPath)
	if err != nil {
		return err
	}

	err = p.Display.ShowTitle(song.Title)
	if err != nil {
		return err
	}
	musicPath := filepath.Join(p.MusicDir, "Hinário Adventista 2022/Santo, Santo, Santo!.mp3")
	err = p.AudioPlayer.LoadAudio(musicPath)
	if err != nil {
		return err
	}

	err = p.AudioPlayer.Play()
	if err != nil {
		return err
	}
	startTime := time.Now()
	for _, verse := range song.Verses {
		absoluteVerseTime := startTime.Add(verse.Time)
		relativeVerseTime := time.Until(absoluteVerseTime)
		select {
		case <-p.stopChan:
			log.Info("stopping because Stop() was called")
			return nil
		case <-time.After(relativeVerseTime):
			p.Display.ShowVerse(strings.Split(verse.Text, "\n")...)
		}
	}

	p.AudioPlayer.Wait()
	fmt.Println("Waiting last second")
	time.Sleep(1 * time.Second)
	return nil
}
