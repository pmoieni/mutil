package mutil

import (
	"math"

	"github.com/pmoieni/mutil/internal/consts"
	"github.com/pmoieni/mutil/internal/math32"
)

type Pitch float32

const (
	bufferSize int = 2048 // this should be dynamic
)

func PitchToNote(p Pitch) *Note {
	num := consts.OctaveLen * (math32.Log(float32(p)/440) / math32.Log(2))
	var mnum MIDINumber = MIDINumber(math32.Round(num) + 69)

	return mnum.ToNote()
}

func AutoCorrelate(buf [bufferSize]float32, sampleRate, thres float64) Pitch {
	size := len(buf)
	var rms float64 = 0
	for _, v := range buf {
		val := float64(v)
		rms += val * val
	}
	rms = math.Sqrt(rms / float64(size))
	if rms < thres {
		return -1
	}

	r1 := 0
	r2 := size - 1

	for i := 0; i < size/2; i++ {
		if math.Abs(float64(buf[i])) < thres {
			r1 = i
			break
		}
	}

	for i := 0; i < size/2; i++ {
		if math.Abs(float64(buf[size-i])) < thres {
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
	for d < size-2 && c[d] > c[d+1] {
		d++
	}

	var mv float64 = -1.0 // max value
	var mp int = -1       // max position

	for i := d; i < size; i++ {
		if float64(c[i]) > mv {
			mv = float64(c[i])
			mp = i
		}
	}

	t0 := mp

	if mp < 1 {
		return -1
	}

	x1 := c[t0-1]
	x2 := c[t0]
	x3 := c[t0+1]

	a := float64((x1 + x3 - 2*x2) / 2)
	b := float64((x3 - x1) / 2)

	return Pitch(sampleRate / (float64(t0) - b/(2*a)))
}

func (p Pitch) CentsOff(mnum MIDINumber) float64 {
	return math.Floor(consts.OctaveLen *
		100 *
		math.Log(float64(p)/float64(*mnum.ToPitch())) /
		math.Log(2),
	)
}
