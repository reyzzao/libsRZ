package message

type Message struct {
	Sucess     string
	Fail       string
	Processing string
	Wait       string
	Min        string
	Max        string
}

var StaticMessage = Message{
	Sucess:     "Sucesso!",
	Fail:       "Ops!",
	Processing: "Processando...",
	Wait:       "Aguarde!",
	Min:        "no mínimo",
	Max:        "no máximo",
}

// func MessageFN() Message{
// 	return StaticMessage
// }
