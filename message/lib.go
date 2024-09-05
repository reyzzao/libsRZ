package message

import "fmt"

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

// var molduraMsgTestInit = "***********RESPONSE TEST *************"
// var msgCompare = "COMPARE : %s"
// var msgExpect = "EXPECTED : %s"
// var molduraMsgTestFinal = "*********** FINAL RESPONSE TEST *************"

// func MessageFN() Message{
// 	return StaticMessage
// }

func MessageErrorTestLIB(compare any, expect any) string {
	msgFormat := fmt.Sprintf(`

		***********RESPONSE TEST *************

		COMPARE : %s

		EXPECTED : %s

		*********** FINAL RESPONSE TEST *************

		`, compare, expect)

	res := msgFormat

	return res
}
