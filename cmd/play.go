package main

import (
	"path/filepath"
	"strconv"

	player "github.com/mniak/japlayer"
	"github.com/mniak/japlayer/adapters/sdl"
	"github.com/mniak/japlayer/adapters/sqlite"
	"github.com/samber/lo"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	cmdRoot.AddCommand(cmdPlay)

	cmdPlay.Flags().StringP("config-dir", "c", "config", "Path to the 'config' directory, where the database, the images and music is stored.")
	cmdPlay.Flags().Float32("font-size", 96, "Specify the font size")
	cmdPlay.Flags().Bool("windowed", false, "Don't use full screen, but open in a window")
}

var cmdPlay = &cobra.Command{
	Use:   "play DATABASE",
	Short: "Play a hymn",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hymnNumber := lo.Must(strconv.Atoi(args[0]))
		fontSize := lo.Must(cmd.Flags().GetFloat32("font-size"))
		configDir := lo.Must(cmd.Flags().GetString("config-dir"))
		windowed := lo.Must(cmd.Flags().GetBool("windowed"))

		imagesDir := filepath.Join(configDir, "imagens")
		musicDir := filepath.Join(configDir, "musicas")

		dbFilename := filepath.Join(configDir, "DB.db")
		fontFilename := filepath.Join(configDir, "fontes/din-condensed-bold.ttf")

		sqliteAdapter := lo.Must(sqlite.NewAdapter(dbFilename))
		defer sqliteAdapter.Close()

		lo.Must0(sdl.Init())
		defer sdl.Quit()

		sdlAdapter := lo.Must(sdl.SDLAdapter(sdl.AdapterParams{
			FontPath:   fontFilename,
			FontSize:   fontSize,
			Fullscreen: !windowed,
		}))
		defer sdlAdapter.Finish()

		app := player.Player{
			SongLoader:  sqliteAdapter,
			Display:     sdlAdapter,
			AudioPlayer: sdlAdapter,
			ImagesDir:   imagesDir,
			MusicDir:    musicDir,
		}
		go sdl.HandleEvents(player.EventHandler{
			Quit: func() {
				app.Stop()
			},
		})
		lo.Must0(app.PresentLyrics(hymnNumber))
	},
}
