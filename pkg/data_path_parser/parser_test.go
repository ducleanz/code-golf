package data_path_parser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePathHappy(t *testing.T) {
	tests := map[string][]any{
		`.`:                 {},        // Root reference
		`."hello"`:          {"hello"}, // Basic object references
		`."multi"."path"`:   {"multi", "path"},
		`.1`:                {1}, // Array references
		`.1.234`:            {1, 234},
		`.1."hi".2."there"`: {1, "hi", 2, "there"}, // Mixed
		`."hi".1."there".2`: {"hi", 1, "there", 2},
		`.""`:               {""}, // might look like it's invalid, but you can totally have an empty string
		// as a key in JSON
	}
	for input, expect := range tests {
		t.Run(input, func(t *testing.T) {
			val, err := ParsePath(input)
			assert.NoError(t, err)
			assert.Equalf(t, len(expect), len(val), "expected len: %v, actual len: %v", len(expect), len(val))
			for i := range val {
				assert.Truef(t, reflect.DeepEqual(expect[i], val[i]), "expected element[%d]: %v, actual[%d]: %v", i, expect[i], i, val[i])
			}
		})
	}
}

func TestParsePathSad(t *testing.T) {
	tests := map[string]assert.ErrorAssertionFunc{
		`"no"."leading"."dot"`: assert.Error, // No leading dot
		"":                     assert.Error, // Empty path is invalid (and also has no leading dot)
		`..`:                   assert.Error, // Double dots are not allowed (effectively a `null` missing key)
		`."hi".."broken"`:      assert.Error,
	}
	for input, assertError := range tests {
		t.Run(input, func(t *testing.T) {
			val, err := ParsePath(input)
			assert.Nil(t, val)
			assertError(t, err, "error: %v", err)
		})
	}
}
