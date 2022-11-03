package mutil

type Class int

const (
	A Class = iota + 1 // 1
	B                  // 2
	C                  // 3
	D                  // 4
	E                  // 5
	F                  // 6
	G                  // 7
)

// up distance between two classes
// moves to next octave if higher classes in current octave don't match "from"
func (c Class) DiffUp(from Class) int {
	if c == from {
		return 0
	} else if from > c {
		return int(from - c)
	} else {
		ol := 7 // octave length
		return ol - int(c-from)
	}
}

// down distance between two classes
// moves to previous octave if lower classes in current octave don't match "from"
func (c Class) DiffDown(from Class) int {
	if c == from {
		return 0
	} else if c > from {
		return int(c - from)
	} else {
		ol := 7
		return ol - int(from-c)
	}
}
