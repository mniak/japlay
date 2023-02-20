package player

import "time"

type Song struct {
	Title      string
	CoverImage string
	Verses     []Verse
}

type Verse struct {
	Text string
	Time time.Duration
}
