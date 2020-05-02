package checksum

import "testing"

func TestVerhoeffCheck(t *testing.T) {
	var vh Verhoeff

	c, err := vh.Check("2363")
	if c != true || err != nil {
		t.Errorf("Check(2363) != true, got %v (%v)", c, err)
	}
}

func TestVerhoeffCompute(t *testing.T) {
	var vh Verhoeff

	c, s, err := vh.Compute("236")
	if c != 3 || err != nil {
		t.Errorf("Compute(236) != 3, got %v, %v (%v)", c, s, err)
	}
}

func TestVerhoeffInvalid(t *testing.T) {
	var vh Verhoeff

	c, err := vh.Check("2A363")
	if c != false || err == nil {
		t.Errorf("Check(23A63) != false, got %v (%v)", c, err)
	}
}
