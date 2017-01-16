package package1

import (
	"testing"
)

func TestAverage(t *testing.T) {
	v := Average(1, 2)
	if v != 1.5 {
		t.Fatalf("expected 1.6 received %f", v)
	}
}
