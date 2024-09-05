package message

// import (
// 	"testing"
// )

// TODO : Organizar e aproveitar types : por enquanto funcionando somnet a funcao de msg test

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

func MessageErrorTestLIB() string {
	msgFormat := "\n ***********RESPONSE TEST *************\n Ops... Esperado: %t --> Tentativa: %t \n*********** FINAL RESPONSE TEST ************* "
	res := msgFormat
	return res

	// return "\n ***********RESPONSE TEST *************\n Ops... Esperado: %t --> Tentativa: %t \n*********** FINAL RESPONSE TEST ************* "
}

// func MessageErrorTest(compare any, expect any)  {
// 	t.Error("\n ***********RESPONSE TEST *************\n Ops... Esperado: %t --> Tentativa: %t \n*********** FINAL RESPONSE TEST ************* ", expect, compare)

// }
