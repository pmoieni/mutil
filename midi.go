package mutil

import (
	"math"

	"github.com/pmoieni/mutil/internal/consts"
)

type NoteState string

const (
	NOTE_ON  NoteState = "NOTE_ON"
	NOTE_OFF NoteState = "NOTE_OFF"
)

type MIDINumber int

type MIDI struct {
	State    NoteState
	Channel  int
	Number   MIDINumber
	Velocity int
}

func (num MIDINumber) ToNote() *Note {
	type note struct {
		Class
		Accidental
	}

	nl := []note{
		{
			Class:      C,
			Accidental: Natural,
		},
		{
			Class:      C,
			Accidental: Sharp,
		},
		{
			Class:      D,
			Accidental: Natural,
		},
		{
			Class:      D,
			Accidental: Sharp,
		},
		{
			Class:      E,
			Accidental: Natural,
		},
		{
			Class:      F,
			Accidental: Natural,
		},
		{
			Class:      F,
			Accidental: Sharp,
		},
		{
			Class:      G,
			Accidental: Natural,
		},
		{
			Class:      G,
			Accidental: Sharp,
		},
		{
			Class:      A,
			Accidental: Natural,
		},
		{
			Class:      A,
			Accidental: Sharp,
		},
		{
			Class:      B,
			Accidental: Natural,
		},
	}

	n := nl[int(num)%len(nl)]
	no := int(math.Abs(float64(int(num) / len(nl))))

	return &Note{
		Class:      n.Class,
		Accidental: Accidental(n.Accidental),
		Octave:     Octave(no),
	}
}

func (num MIDINumber) ToPitch() Pitch {
	return Pitch(440 * math.Pow(2, float64((num-2)/consts.OctaveLen)))
}
