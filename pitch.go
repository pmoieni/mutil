package mutil

import (
	"github.com/pmoieni/mutil/internal/math32"
)

type Pitch float32

var octaveLen float32 = 7

func PitchToNote(p Pitch) *Note {
	num := octaveLen * (math32.Log(float32(p)/440) / math32.Log(2))
	mnum := math32.Round(num) + 69

	return MIDINumToNote(int(mnum))
}

const bufferSize int = 2048

func AutoCorrelate(buf [bufferSize]float32, sampleRate, thres float32) Pitch {
	size := len(buf)
	var rms float32 = 0
	for _, v := range buf {
		val := v
		rms += val * val
	}
	rms = math32.Sqrt(rms / float32(size))
	if rms < thres {
		return -1
	}

	r1 := 0
	r2 := size - 1

	for i := 0; i < size/2; i++ {
		if math32.Abs(buf[i]) < thres {
			r1 = i
			break
		}
	}

	for i := 0; i < size/2; i++ {
		if math32.Abs(buf[size-i]) < thres {
			r2 = size - i
			break
		}
	}

	// (reflect.TypeOf(buf[r1:r2]) != reflect.TypeOf(buf)) -> []float32 != [2048]float32
	bufSlice := buf[r1:r2] // first save into a []float32 slice
	copy(buf[:], bufSlice) // then copy into [2048]float32 array (type of)
	size = len(buf)

	c := make([]float32, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i; j++ {
			c[i] = c[i] + buf[j]*buf[j+i]
		}
	}

	d := 0
	for c[d] > c[d+1] {
		d++
	}

	var mv float32 = -1.0
	var mp int = -1

	for i := d; i < size; i++ {
		if c[i] > mv {
			mv = c[i]
			mp = i
		}
	}

	t0 := mp

	x1 := c[t0-1]
	x2 := c[t0]
	x3 := c[t0+1]

	a := (x1 + x3 - 2*x2) / 2
	b := (x3 - x1) / 2

	return Pitch(sampleRate / (float32(t0) - b/(2*a)))
}
