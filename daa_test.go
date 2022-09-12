package data_structure_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerpangkatan(t *testing.T) {
	angka := PerpangkatanRekursif(3, 3)
	assert.Equal(t, int64(27), angka)

}
