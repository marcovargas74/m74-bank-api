package account

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
)

/*
//INTERFACES----
//Transation is a transation bank interface
type Transation interface {
	GetBalance(nome string) int
	SaveTransfer(nome string)
	GetTransations() Liga


}*/

/*Codigo de apoio UBER

type SMap struct {
	mu sync.Mutex

	data map[string]string
  }

  func NewSMap() *SMap {
	return &SMap{
	  data: make(map[string]string),
	}
  }

  func (m *SMap) Get(k string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data[k]
  }*/

// Transation armazena uma coleção de Transition
type Transation []Transfer

// Encontrar tenta retornar um jogador de uma liga
func (l Transation) FindTransition(nome string) *Client {
	for i, p := range l {
		if p.Nome == nome {
			return &l[i]
		}
	}
	return nil
}

// NovaLiga cria uma liga de um JSON
func NovaLiga(leitor io.Reader) (Liga, error) {
	var liga []Jogador
	err := json.NewDecoder(leitor).Decode(&liga)

	if err != nil {
		err = fmt.Errorf("falha ao analizar a liga, %v", err)
	}

	return liga, err
}

func transfer() {
	fmt.Println("transfer")
}

//TransferBank is A struct to used to make a transfer
type Transfer struct {
	ID                   string  `json:"id"`
	AccountOriginID      string  `json:"acount_origin_id"`
	AccountDestinationID string  `json:"Account_destination_id"`
	Amount               float64 `json:"Amount"`
	CreatedAt            string  `json:"created_at"` //TODO change to date
	Tmutex               sync.Mutex
	//data map[string]string
}

func structAndJSONTransfer() {
	transfer1 := Transfer{"xyz", "abc", "def", 12.00, "17-01-2022"}
	transfJSON, _ := json.Marshal(transfer1)
	fmt.Println(string(transfJSON))
	//Convert Json To struct
	var aTransfFromJSON Transfer
	json.Unmarshal(transfJSON, &aTransfFromJSON)
	fmt.Println(aTransfFromJSON.ID)

}

/*
/transfers
A entidade Transfer possui os seguintes atributos:

id
account_origin_id
account_destination_id
amount
created_at
Espera-se as seguintes ações:

GET /transfers - obtém a lista de transferencias da usuaria autenticada.
POST /transfers - faz transferencia de uma Account para outra.
Regras para esta rota

Quem fizer a transferência precisa estar autenticada.
O account_origin_id deve ser obtido no Token enviado.
Caso Account de origem não tenha saldo, retornar um código de erro apropriado
Atualizar o balance das contas
*/
