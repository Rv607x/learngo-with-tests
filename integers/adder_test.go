package integers

import "testing"

func TestAdder(t *testing.T) {
	sum_of_int := sum(2, 2)
	want := 4
	if sum_of_int != want {
		t.Errorf(" expected %q, and got %q", want, sum_of_int)
	}
}
