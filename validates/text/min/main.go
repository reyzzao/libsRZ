package min

import (
	"errors"

	"github.com/rzjprogramador/Libs_Rzj/Go_Libs/validates/feedback"
)

func MinText(text string, length int) (string, error) {
	if len(text) < length {
		return "", errors.New(feedback.MinText(length))
	}
	return text, nil
}
