package mutil

type NoteState string

const (
	NOTE_ON  NoteState = "NOTE_ON"
	NOTE_OFF NoteState = "NOTE_OFF"
)

type MIDI struct {
	State    NoteState
	Channel  int
	Number   int
	Velocity int
}
