package math32

import (
	"math"
	"testing"
	"testing/quick"
)

func TestAbs(t *testing.T) {
	f := func(x float32) bool {
		y := Abs(x)
		return y == float32(math.Abs(float64(x)))
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInf(t *testing.T) {
	if float64(Inf(1)) != math.Inf(1) || float64(Inf(-1)) != math.Inf(-1) {
		t.Error("float32(inf) not infinite")
	}
}

func TestIsInf(t *testing.T) {
	posInf := float32(math.Inf(1))
	negInf := float32(math.Inf(-1))
	if !IsInf(posInf, 0) || !IsInf(negInf, 0) || !IsInf(posInf, 1) || !IsInf(negInf, -1) || IsInf(posInf, -1) || IsInf(negInf, 1) {
		t.Error("unexpected isInf value")
	}
	f := func(x struct {
		F    float32
		Sign int
	}) bool {
		y := IsInf(x.F, x.Sign)
		return y == math.IsInf(float64(x.F), x.Sign)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestIsNaN(t *testing.T) {
	f := func(x float32) bool {
		y := IsNaN(x)
		return y == math.IsNaN(float64(x))
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNaN(t *testing.T) {
	if !math.IsNaN(float64(NaN())) {
		t.Errorf("float32(nan) is a number: %f", NaN())
	}
}

func TestSignbit(t *testing.T) {
	f := func(x float32) bool {
		return Signbit(x) == math.Signbit(float64(x))
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
