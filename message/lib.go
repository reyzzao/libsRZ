package message

// import (
// 	"testing"
// )

// var t *testing.T

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

func LibMessageErrorTest() string {
	return "\n ***********RESPONSE TEST *************\n Ops... Esperado: %t --> Tentativa: %t \n*********** FINAL RESPONSE TEST ************* "
}

// func MessageErrorTest(compare any, expect any)  {
// 	t.Error("\n ***********RESPONSE TEST *************\n Ops... Esperado: %t --> Tentativa: %t \n*********** FINAL RESPONSE TEST ************* ", expect, compare)

// }
