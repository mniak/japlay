package player

type SongLoader interface {
	AlbumTrack(albumID, track int) (Song, error)
	SongByID(id string) (Song, error)
}
