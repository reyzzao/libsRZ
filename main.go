package main

import (
	"github.com/reizzao/RzLibs_GO/feedback"
	"github.com/reizzao/RzLibs_GO/message"
	"github.com/reizzao/RzLibs_GO/message/message_bytest"
)

func main() {

	feedback.TesterFeedback()
	message.TesterMessage()
	message_bytest.LibMessageErrorTest()
}

// importante: tem que chamar a lib aqui na main - para poderem acessa-la remotamente
