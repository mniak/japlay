package player

type SongLoader interface {
	SongByID(id string) (Song, error)
}
