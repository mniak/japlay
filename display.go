package player

type Display interface {
	ShowVerse(lines ...string) error
	ShowTitle(title string) error
	SetBackgroundImage(filename string) error
}
