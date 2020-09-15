package average

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEverage(t *testing.T) {
	assert := assert.New(t)
	tc1 := []float64{}
	tc2 := []float64{5, 2, 1, 2, 3}
	tc3 := []float64{3, 5, 9, 11}
	tc4 := []float64{71.8, 56.2, 89.5}

	assert.Equal(0.0, Calc(tc1))
	assert.Equal(13.0/5.0, Calc(tc2))
	assert.Equal(28.0/4.0, Calc(tc3))
	assert.Equal(217.5/3.0, Calc(tc4))

}
