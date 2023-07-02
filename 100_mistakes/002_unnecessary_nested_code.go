package mistakes

import (
	"errors"
	"strings"
)

type EmptyStringErr error

var (
	ErrEmptyStringS1 = EmptyStringErr(errors.New("s1 is empty"))
	ErrEmptyStringS2 = EmptyStringErr(errors.New("s2 is empty"))
)

// Align the happy path to the left; you should quickly be able to scan down one column to see the expected execution flow
// the happy path is where each if starts.
func Join(s1, s2 string, max uint) (string, error) {
	if strings.TrimSpace(s1) == "" {
		return "", ErrEmptyStringS1
	}

	if strings.TrimSpace(s2) == "" {
		return "", ErrEmptyStringS2
	}

	if joined := s1 + s2; len(joined) > int(max) {
		return joined[:max], nil
	} else {
		return joined, nil
	}
}

func UglyJoin(s1, s2 string, max uint) (string, error) {
	if strings.TrimSpace(s1) == "" {
		return "", ErrEmptyStringS1
	} else {
		if strings.TrimSpace(s2) == "" {
			return "", ErrEmptyStringS2
		} else {
			joined := s1 + s2
			if len(joined) > int(max) {
				return joined[:max], nil
			} else {
				return joined, nil
			}
		}
	}
}
