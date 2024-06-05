package feedback

import "fmt"

func MinText(length int) string {
	return fmt.Sprintf("Ops... o minimo de letras é %d", length)
}

func MaxText(length int) string {
	return fmt.Sprintf("Ops... o maximo de letras é %d", length)
}
