package data_getter

import (
	"errors"
	"fmt"

	"code_golf/pkg/types"
)

func LookupPath(p types.PathSegments, target any) (any, error) {
	current := target
	for _, segment := range p {
		switch curr := current.(type) {
		case map[string]any:
			key, ok := segment.(string)
			if !ok {
				return nil, fmt.Errorf("expected string key, got %T", segment)
			}
			value, exists := curr[key]
			if !exists {
				return nil, fmt.Errorf("key %q not found", key)
			}
			current = value
		case []any:
			index, ok := segment.(int)
			if !ok {
				return nil, fmt.Errorf("expected integer index, got %T", segment)
			}
			if index < 0 {
				index = len(curr) + index
			}
			if index < 0 || index >= len(curr) {
				return nil, fmt.Errorf("index %d out of range", index)
			}
			current = curr[index]
		default:
			return nil, errors.New("invalid structure")
		}
	}
	return current, nil
}
