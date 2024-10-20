Implement a function for parsing this simple alternative to JSON Path

# Rules

* a path is a collection of segments separated by .
    * a path ALWAYS starts with a leading .
    * a path CAN be just a single .
        * outside this special case, a path NEVER ends with a .
* a segment can be either:
    * a string enclosed in double quotes, eg. "hello"
        * string segments may be the empty string, ie. ""
        * to keep things simple: you do not need to handle escaped characters
            * ie. there’s no need to handle segments like "this \"is\" harder"
    * a positive int



# Example paths

* Valid
    * `."data"."people".0."name"`  
    * `.`  
    * `.""`  
    * `."".1.2.3."meta".1`  
* Invalid
    * (empty)  
    * `"no"."leading"."dot"`  
    * `..`  

# Expected function signature

The function should accept a string and return an array of any (or an error if the path was invalid)
Each entry in the array should be either an int or a string (string segments should not have their surrounding quotes included)

```golang
// A collection of path segments, each either a string or an int
type PathSegments []any

func ParsePath(path string) (PathSegments, error) {
// your clever stuff here!
}
```

# Example tests

```golang
package main

import (
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
            assert.Equal(t, expect, val)
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
            assertError(t, err)
        })
    }
}
```