package data_getter

import (
	parser "code_golf/pkg/data_path_parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGet(t *testing.T) {
	testData := map[string]any{
		"people": map[string]any{
			"alice": map[string]any{"name": "Alice", "tags": []any{"manager", "human"}},
			"bob":   map[string]any{"name": "Bob", "tags": []any{"programmer", "human", map[string]any{"types": "special"}}, "boss": "alice"},
		},
	}

	tests := map[string]struct {
		expect    any
		expectErr assert.ErrorAssertionFunc
	}{
		`."people"`:                         {expect: testData["people"], expectErr: assert.NoError},
		`."people".10`:                      {expectErr: assert.Error},
		`."people"."alice"."name"`:          {expect: "Alice", expectErr: assert.NoError},
		`."people"."alice"."WRONGKEY"`:      {expectErr: assert.Error},
		`."people"."bob"."tags".0`:          {expect: "programmer", expectErr: assert.NoError},
		`."people"."bob"."tags".-1."types"`: {expect: "special", expectErr: assert.NoError},
	}

	for query, test := range tests {
		t.Run(query, func(t *testing.T) {
			pathParts, err := parser.ParsePath(query)
			require.NoError(t, err)

			val, err := LookupPath(pathParts, testData)
			test.expectErr(t, err)
			assert.Equal(t, test.expect, val)
		})
	}
}
