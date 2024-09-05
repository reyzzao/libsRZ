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

var molduraMsgTestInit = "\n ***********RESPONSE TEST *************\n"
var msgEsperado = "Ops... Esperado: %t"
var msgExpect = "\n Tentativa: %t"
var molduraMsgTestFinal = "\n*********** FINAL RESPONSE TEST *************\n"

// func MessageFN() Message{
// 	return StaticMessage
// }

func MessageErrorTestLIB() string {
	msgFormat := molduraMsgTestInit + msgEsperado + msgExpect + molduraMsgTestFinal

	res := msgFormat

	return res

	// return "\n ***********RESPONSE TEST *************\n Ops... Esperado: %t --> Tentativa: %t \n*********** FINAL RESPONSE TEST ************* "
}

// func MessageErrorTest(compare any, expect any)  {
// 	t.Error("\n ***********RESPONSE TEST *************\n Ops... Esperado: %t --> Tentativa: %t \n*********** FINAL RESPONSE TEST ************* ", expect, compare)

// }
