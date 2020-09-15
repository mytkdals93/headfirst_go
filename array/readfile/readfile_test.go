package readfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFloat(t *testing.T) {
	assert := assert.New(t)
	numbers, err := GetFloats("testdata.txt")

	assert.NoError(err)

	assert.Equal([]float64{32.5, 35.5, 31.5}, numbers)
}
