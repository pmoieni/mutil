package mutil

type Note struct {
	Class
	Accidental
	Octave
	Pitch
}

var Notes = []Note{
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
