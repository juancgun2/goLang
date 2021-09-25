package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResult(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // parser result
	}{
		{"TX02AB", true},
		{"NN100987654321", true},
		{"TX04ABCDF", true}, // for me this input it's ok, just will print 4 characters of the value
		{"NN04000A", false}, // type doesn´t match with the input value
		{"NN23AAAA", false}, // lenght out of index
	}
	assert := assert.New(t) // If you assert many times, use this format. LINK:https://pkg.go.dev/github.com/stretchr/testify/assert

	for _, testData := range cases {
		_, err := NewResult(testData.Input)
		assert.Equal(err == nil, testData.Success, "If err == nil, success should be true")
		assert.Equal(err != nil, !testData.Success, "If err != nil, success should be false")
	}

}
