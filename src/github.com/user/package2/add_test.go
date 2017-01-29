package package2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	v := Add(2, 3)
	expected := 5
	assert.Equal(t, expected, v, "Expected %v but received: %v", expected, v)

}
