/*
go test -v
go test -v -run TestStringToFloat64
*/
package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToFloat64(t *testing.T) {
	t.Parallel()
	// define inputs
	input := "6"

	// define expected result
	expected := 6

	// perform test
	actual := StringToFloat64(input)

	// assert that the actual result is equal to expected
	assert.Equal(t, expected, actual)
}
