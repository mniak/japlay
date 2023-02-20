package console

import (
	"fmt"

	"github.com/mniak/louvorja/player"
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
