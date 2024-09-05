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

func MessageErrorTestLIB(real_expected any, try_test any) string {
	msgFormat := fmt.Sprintf(`

		***********RESPONSE TEST *************

		REAL_ESPERADO : %s

		TENTADO_NO_TEST : %s

		*********** FINAL RESPONSE TEST *************

		`, real_expected, try_test)

	res := msgFormat

	return res
}
