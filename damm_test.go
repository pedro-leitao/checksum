package checksum

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestDammCheck(t *testing.T) {
	var dm Damm

	c, err := dm.Check("5724")
	if c != true || err != nil {
		t.Errorf("Check(5724) != true, got %v (%v)", c, err)
	}
}

func TestDammCompute(t *testing.T) {
	var dm Damm

	c, s, err := dm.Compute("572")
	if c != 4 || err != nil {
		t.Errorf("Compute(572) != 4, got %v, %v (%v)", c, s, err)
	}
}

func TestDammInvalid(t *testing.T) {
	var dm Damm

	c, err := dm.Check("57A24")
	if c != false || err == nil {
		t.Errorf("Check(57A24) != false, got %v (%v)", c, err)
	}
}

func BenchmarkDamm(b *testing.B) {

	var dm Damm

	//b.StopTimer()
	from := rand.Intn(1000000000) + 1000000000
	to := from + b.N
	for i := from; i <= to; i++ {
		var s, ns string
		var checks bool

		s = strconv.Itoa(i)
		//b.StartTimer()
		_, ns, _ = dm.Compute(s)
		if checks, _ = dm.Check(ns); checks != true {
			b.Errorf("%v: failed, which should never have happened...", i)
		}
		//b.StopTimer()
	}
}
