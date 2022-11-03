package mutil

import "time"

type Note struct {
	Class
	Accidental
	Octave
	Pitch  float64
	Length time.Duration
}
