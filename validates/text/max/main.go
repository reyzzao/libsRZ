package max

import (
	"errors"

	"github.com/rzjprogramador/Libs_Rzj/Go_Libs/validates/feedback"
)

func MaxText(text string, length int) (string, error) {
	if len(text) > length {
		return "", errors.New(feedback.MaxText(length))
	}
	return text, nil
}
