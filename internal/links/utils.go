package links

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrInvalidUTMValue = errors.New("invalid utm value")
)

var utmValuePattern = regexp.MustCompile(`^[a-zA-Z0-9\-_\.]+$`)

func validateUTMValue(val string) error {
	if !utmValuePattern.MatchString(val) {
		return fmt.Errorf("%w: %s", ErrInvalidUTMValue, val)
	}

	return nil
}
