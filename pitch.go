package mutil

import (
	"math"
)

type Pitch float64

func (p Pitch) ToNote() *Note {
	return &Note{}
}

func (p Pitch) AutoCorrelate(buf []float64, sampleRate float64, thres float64) float64 {
	size := len(buf)
	var rms float64 = 0
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

	maxval := -1.0
	maxpos := -1

	for i := d; i < size; i++ {
		if c[i] > maxval {
			maxval = c[i]
			maxpos = i
		}
	}

	t0 := maxpos

	x1 := c[t0-1]
	x2 := c[t0]
	x3 := c[t0+1]

	a := (x1 + x3 - 2*x2) / 2
	b := (x3 - x1) / 2

	return sampleRate / (float64(t0) - b/(2*a))
}
