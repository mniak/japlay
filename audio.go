package player

type AudioPlayer interface {
	LoadAudio(audiofile string) error
	Play() error
	// Pause() error
	// Seek(time time.Duration)
}
