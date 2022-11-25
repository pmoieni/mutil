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

/*
  function autoCorrelate(buf, sampleRate) {
    // Implements the ACF2+ algorithm
    let SIZE = buf.length;
    let rms = 0;

    for (let i = 0; i < SIZE; i++) {
      let val = buf[i];
      rms += val * val;
    }
    rms = Math.sqrt(rms / SIZE);
    if (rms < sensitivity)
      // not enough signal
      // the note is ignored
      return -1;

    let r1 = 0,
      r2 = SIZE - 1,
      thres = 0.2;
    for (let i = 0; i < SIZE / 2; i++)
      if (Math.abs(buf[i]) < thres) {
        r1 = i;
        break;
      }
    for (let i = 1; i < SIZE / 2; i++)
      if (Math.abs(buf[SIZE - i]) < thres) {
        r2 = SIZE - i;
        break;
      }

    buf = buf.slice(r1, r2);
    SIZE = buf.length;

    let c = new Array(SIZE).fill(0);
    for (let i = 0; i < SIZE; i++)
      for (let j = 0; j < SIZE - i; j++) c[i] = c[i] + buf[j] * buf[j + i];

    let d = 0;
    while (c[d] > c[d + 1]) d++;
    let maxval = -1,
      maxpos = -1;
    for (let i = d; i < SIZE; i++) {
      if (c[i] > maxval) {
        maxval = c[i];
        maxpos = i;
      }
    }
    let T0 = maxpos;

    let x1 = c[T0 - 1],
      x2 = c[T0],
      x3 = c[T0 + 1];
    let a = (x1 + x3 - 2 * x2) / 2,
      b = (x3 - x1) / 2;
    if (a) T0 = T0 - b / (2 * a);

    return sampleRate / T0;
  }

  // Implements modified ACF2+ algorithm
// Source: https://github.com/cwilso/PitchDetect
export const autoCorrelate = (buf, sampleRate) => {
  // Not enough signal check
  const RMS = Math.sqrt(buf.reduce((acc, el) => acc + el ** 2, 0) / buf.length)
  if (RMS < 0.001) return NaN

  const THRES = 0.2
  let r1 = 0
  let r2 = buf.length - 1
  for (let i = 0; i < buf.length / 2; ++i) {
    if (Math.abs(buf[i]) < THRES) {
      r1 = i
      break
    }
  }
  for (let i = 1; i < buf.length / 2; ++i) {
    if (Math.abs(buf[buf.length - i]) < THRES) {
      r2 = buf.length - i
      break
    }
  }

  const buf2 = buf.slice(r1, r2)
  const c = new Array(buf2.length).fill(0)
  for (let i = 0; i < buf2.length; ++i) {
    for (let j = 0; j < buf2.length - i; ++j) {
      c[i] = c[i] + buf2[j] * buf2[j + i]
    }
  }

  let d = 0
  for (; c[d] > c[d + 1]; ++d);

  let maxval = -1
  let maxpos = -1
  for (let i = d; i < buf2.length; ++i) {
    if (c[i] > maxval) {
      maxval = c[i]
      maxpos = i
    }
  }
  let T0 = maxpos

  let x1 = c[T0 - 1]
  let x2 = c[T0]
  let x3 = c[T0 + 1]
  let a = (x1 + x3 - 2 * x2) / 2
  let b = (x3 - x1) / 2

  return sampleRate / (a ? T0 - b / (2 * a) : T0)
}
*/
