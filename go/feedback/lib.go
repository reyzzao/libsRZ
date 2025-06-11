package feedback

import (
	"fmt"

	"github.com/reizzao/RzLibs_GO/message"
)

func MinText(length int) string {
	return fmt.Sprintf("%s %s %d letras", message.StaticMessage.Fail, message.StaticMessage.Min, length)
}

func MaxText(length int) string {
	return fmt.Sprintf("%s %s %d letras", message.StaticMessage.Fail, message.StaticMessage.Max, length)
}
