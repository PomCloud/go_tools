package go_tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyString(t *testing.T) {
	testStr := []struct {
		inputStr string
		result   bool
	}{
		{
			inputStr: " ",
			result:   false,
		},
		{
			inputStr: "abcd",
			result:   true,
		},
	}
	for _, v := range testStr {
		assert.Equal(t, EmptyString(v.inputStr), v.result)
	}
}

func TestValidInt(t *testing.T) {
	test := []struct {
		input  int
		result bool
	}{
		{
			input:  -1,
			result: false,
		},
		{
			input:  0,
			result: false,
		}, {
			input:  1,
			result: true,
		},
	}
	for _, v := range test {
		assert.Equal(t, ValidInt(v.input), v.result)
	}
}

func TestValidIP4(t *testing.T) {
	test := []struct {
		input  string
		result bool
	}{
		{
			input:  "192.168.13.256",
			result: false,
		},
		{
			input:  "1234567",
			result: false,
		},
		{
			input:  "abcddsf",
			result: false,
		},
		{
			input:  "0.0.0.0",
			result: true,
		},
		{
			input:  "127.0.0.1",
			result: true,
		},
	}
	for _, v := range test {
		assert.Equal(t, ValidIP4(v.input), v.result)
	}
}
