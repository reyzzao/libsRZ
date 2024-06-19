package message

type Message struct {
	sucess     string
	fail       string
	processing string
	wait       string
}

var StaticMessage = Message{
	sucess:     "Sucesso!",
	fail:       "Ops!",
	processing: "Processando...",
	wait:       "Aguarde!",
}

// func MessageFN() Message{
// 	return StaticMessage
// }
