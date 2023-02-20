package sqlite

import (
	"database/sql"

	"github.com/mniak/louvorja/player"
)

func (ad *sqliteAdapter) SongByID(songID string) (player.Song, error) {
	song, err := ad.songByID(songID)
	if err != nil {
		return player.Song{}, err
	}

	song.Verses, err = ad.versesBySongID(songID)
	return song, err
}

func (ad *sqliteAdapter) songByID(songID string) (player.Song, error) {
	const script = `
		select NOME, IMAGEM
		from MUSICAS
		where ID = ?
	`
	dbrow := ad.db.QueryRowx(script, songID)

	if dbrow.Err() != nil {
		return player.Song{}, dbrow.Err()
	}

	var row struct {
		Nome   string `db:"NOME"`
		Imagem string `db:"IMAGEM"`
	}
	err := dbrow.StructScan(&row)
	return player.Song{
		Title:      row.Nome,
		CoverImage: row.Imagem,
	}, err
}

func (ad *sqliteAdapter) versesBySongID(songID string) ([]player.Verse, error) {
	const script = `
		select ORDEM, LETRA, TEMPO
		from MUSICAS_LETRA
		where MUSICA = ?
		order by ORDEM asc
	`
	rows, err := ad.db.Queryx(script, songID)
	if err != nil {
		return nil, err
	}

	type Row struct {
		Ordem int            `db:"ORDEM"`
		Letra sql.NullString `db:"LETRA"`
		Tempo int            `db:"TEMPO"`
	}
	verses := make([]player.Verse, 0)
	for rows.Next() {
		var row Row
		err = rows.StructScan(&row)
		if err != nil {
			return nil, err
		}

		var verse player.Verse
		if row.Letra.Valid {
			verse.Text = row.Letra.String
		}
		verse.Time = player.ParseTime(row.Tempo)
		verses = append(verses, verse)
	}
	return verses, nil
}
