package player

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

type Player struct {
	SongLoader  SongLoader
	Display     Display
	AudioPlayer AudioPlayer

	ImagesDir string
	MusicDir  string
}

const santoSantoSantoID = "1728"

func (p *Player) PresentLyrics(hymnNumber int) error {
	song, err := p.SongLoader.SongByID(santoSantoSantoID)
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
		relativeVerseTime := absoluteVerseTime.Sub(time.Now())
		time.Sleep(relativeVerseTime)

		if len(strings.TrimSpace(verse.Text)) == 0 {
			fmt.Println("Empty verse at", verse.Time)
		}

		p.Display.ShowVerse(strings.Split(verse.Text, "\n")...)
	}

	p.AudioPlayer.Wait()
	fmt.Println("Waiting last second")
	time.Sleep(1 * time.Second)
	return nil
}
