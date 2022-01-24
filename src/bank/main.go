package main

import (
	"fmt"
	"log"
	"net/http"

	account "github.com/marcovargas74/m74-bank-api/src/account"
	bank "github.com/marcovargas74/m74-bank-api/src/api-bank"
)

var isProduction = false

func init() {
	bank.SetIsProduction(isProduction)
}

func main() {
	fmt.Printf("======== API BANK Version %s isPruduction=%v\n", bank.GetVersion(), bank.GetIsProduction())

	account.StructAndJSON()
	//bank.StartAPI("dev")
	server := bank.NovoServidorJogador(bank.NovoArmazenamentoJogadorEmMemoria())
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}

}
