package package2

import (
	"testing"
)

func TestAdd(t *testing.T) {
	v := Add(2, 3)
	if v != 5 {
		t.Errorf("Expected 5, got %d", v)
	}
}
