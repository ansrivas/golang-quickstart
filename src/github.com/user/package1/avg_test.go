package package1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	v := Average(1, 2)
	assert.Equal(t, 1.5, v, "expected 1.5 as an output")
}
