package main

import (
	"fmt"
	"io"

	"github.com/mniak/louvorja/player"
	"github.com/mniak/louvorja/player/adapters/sdl"
	"github.com/mniak/louvorja/player/adapters/sqlite"
	"github.com/mniak/louvorja/player/songs"
	"github.com/samber/lo"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	cmdRoot.AddCommand(cmdPlay)
}

var cmdPlay = &cobra.Command{
	Use:   "play DATABASE",
	Short: "Play a hymn",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dbFilePath := args[0]
		// hymnName := args[1]

		sqliteAdapter := lo.Must(sqlite.NewAdapter(dbFilePath))
		defer sqliteAdapter.Close()

		lo.Must0(sdl.Init())
		defer sdl.Quit()

		// consoleAdapter := console.NewAdapter()
		sdlAdapter := lo.Must(sdl.NewAdapter(sdl.AdapterParams{
			FontPath: "config/fontes/din-condensed-bold.ttf",
			FontSize: 96,
		}))
		defer sdlAdapter.Finish()

		app := player.Player{
			SongLoader: sqliteAdapter,
			// Display:    consoleAdapter,
			Display:     sdlAdapter,
			AudioPlayer: sdlAdapter,
			ImagesDir:   "config/imagens",
			MusicDir:    "config/musicas",
		}
		lo.Must0(app.PresentLyrics())
	},
}

func playHymn(mdbfile io.ReadSeeker, _ string) error {
	db, err := songs.LoadDatabaseFromMDB(mdbfile)
	if err != nil {
		return err
	}

	albums, err := db.Albums()
	if err != nil {
		return err
	}

	hinario := albums[0]
	songs, err := hinario.Songs()
	if err != nil {
		return err
	}

	song := songs[0]
	lyrics, err := song.Lyrics()
	if err != nil {
		return err
	}

	for _, verse := range lyrics {
		fmt.Print(verse.Text)
	}

	return nil
}
