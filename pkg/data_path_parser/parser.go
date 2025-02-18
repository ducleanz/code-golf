package data_path_parser

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"code_golf/pkg/types"
)

func ParsePath(path string) (types.PathSegments, error) {
	var result types.PathSegments
	if path == "." {
		return result, nil
	}

	if !strings.HasPrefix(path, ".") {
		return nil, errors.New("path must start with a leading dot")
	}
	if strings.Contains(path, "..") {
		return nil, errors.New("path contains consecutive dots")
	}

	// Regular expressions to match different parts of the path
	re := regexp.MustCompile(`\.(?:(-?\d+)|"([^"]*)")`)
	matches := re.FindAllStringSubmatch(path, -1)

	if matches == nil {
		return nil, errors.New("invalid path")
	}

	for _, match := range matches {
		if match[1] != "" {
			// Match is a number
			num, err := strconv.Atoi(match[1])
			if err != nil {
				return nil, err
			}
			result = append(result, num)
		} else {
			// Match is a string
			result = append(result, match[2])
		}
	}

	return result, nil
}
