package mutil

import (
	"math"
)

type Pitch float64

var octaveLen = 7.0

func PitchToNote(p Pitch) *Note {
	num := octaveLen * (math.Log(float64(p)/440) / math.Log(2))
	mnum := math.Round(num) + 69

	return MIDINumToNote(int(mnum))
}

func AutoCorrelate(buf []float64, sampleRate float64, thres float64) Pitch {
	size := len(buf)
	rms := 0.0
	for _, v := range buf {
		val := v
		rms += val * val
	}
	rms = math.Sqrt(rms / float64(size))
	if rms < thres {
		return -1
	}

	r1 := 0
	r2 := size - 1

	for i := 0; i < size/2; i++ {
		if math.Abs(buf[i]) < thres {
			r1 = i
			break
		}
	}

	for i := 0; i < size/2; i++ {
		if math.Abs(buf[size-i]) < thres {
			r2 = size - i
			break
		}
	}

	buf = buf[r1:r2]
	size = len(buf)

	c := make([]float64, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i; j++ {
			c[i] = c[i] + buf[j]*buf[j+i]
		}
	}

	d := 0
	for c[d] > c[d+1] {
		d++
	}

	mv := -1.0
	mp := -1

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

	return Pitch(sampleRate / (float64(t0) - b/(2*a)))
}
