package checksum

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestLuhnCheck(t *testing.T) {
	var lh Luhn

	c, err := lh.Check("71767263188128")
	if c != true || err != nil {
		t.Errorf("Check(71767263188128) != true, got %v (%v)", c, err)
	}
}

func TestLuhnCompute(t *testing.T) {
	var lh Luhn

	c, s, err := lh.Compute("7176726318812")
	if c != 8 || err != nil {
		t.Errorf("Compute(7176726318812) != 8, got %v, %v (%v)", c, s, err)
	}
}

func TestLuhnInvalid(t *testing.T) {
	var lh Luhn

	c, err := lh.Check("7A1767263188128")
	if c != false || err == nil {
		t.Errorf("Check(7A1767263188128) != false, got %v (%v)", c, err)
	}
}

func BenchmarkLuhn(b *testing.B) {

	var lh Luhn

	//b.StopTimer()
	from := rand.Intn(1000000000) + 1000000000
	to := from + b.N
	for i := from; i <= to; i++ {
		var s, ns string
		var checks bool

		s = strconv.Itoa(i)
		//b.StartTimer()
		_, ns, _ = lh.Compute(s)
		if checks, _ = lh.Check(ns); checks != true {
			b.Errorf("%v: failed, which should never have happened...", i)
		}
		//b.StopTimer()
	}
}
