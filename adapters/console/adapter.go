package console

import (
	"fmt"

	player "github.com/mniak/japlayer"
)

type consoleAdapter struct{}

func NewAdapter() *consoleAdapter {
	return &consoleAdapter{}
}

func (ad *consoleAdapter) ShowVerse(verse player.Verse) error {
	fmt.Println(verse.Text)
	fmt.Println("---------------------")
	return nil
}
