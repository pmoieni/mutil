package mutil

import "time"

type Note struct {
	Class
	Accidental
	Octave
	Pitch
	Length time.Duration
}
