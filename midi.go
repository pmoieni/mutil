package mutil

import (
	"math"

	"github.com/pmoieni/mutil/internal/consts"
)

type NoteState int

const (
	NOTE_ON NoteState = iota
	NOTE_OFF
)

type MIDINumber int

type MIDI struct {
	State    NoteState
	Channel  int
	Number   MIDINumber
	Velocity int
}

func (num MIDINumber) ToNote() *Note {
	n := Notes[int(num)%len(Notes)]
	no := int(math.Abs(float64(int(num) / len(Notes))))

	return &Note{
		Class:      n.Class,
		Accidental: Accidental(n.Accidental),
		Octave:     Octave(no),
	}
}

func (num MIDINumber) ToPitch() *Pitch {
	p := Pitch(440 * math.Pow(2, float64((num-2)/consts.OctaveLen)))
	return &p
}
