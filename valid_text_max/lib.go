package valid_text_max

import (
	"errors"

	"github.com/reizzao/RzLibs_GO/feedback"
)

func MaxText(text string, length int) (string, error) {
	if len(text) > length {
		return "", errors.New(feedback.MaxText(length))
	}
	return text, nil
}
