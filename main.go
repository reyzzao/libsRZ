package main

import (
	"github.com/reizzao/RzLibs_GO/feedback"
	"github.com/reizzao/RzLibs_GO/message"
)

func main() {

	feedback.TesterFeedback()
	message.TesterMessage()
	message.LibMessageErrorTest()
}

// importante: tem que chamar a lib aqui na main - para poderem acessa-la remotamente
