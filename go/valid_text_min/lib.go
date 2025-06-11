package valid_text_min

import (
	"errors"

	"github.com/reizzao/RzLibs_GO/feedback"
)

func MinText(text string, length int) (string, error) {
	if len(text) < length {
		return "", errors.New(feedback.MinText(length))
	}
	return text, nil
}
