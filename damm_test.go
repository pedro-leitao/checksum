package checksum

import "testing"

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
