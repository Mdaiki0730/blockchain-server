package model

import (
  "fmt"
  "strings"
  "encoding/json"
)

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float64
}

func NewTransaction(sender string, recipient string, value float64) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address       %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address    %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                           %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
		Value                      float64 `json:"value"`
	}{
		SenderBlockchainAddress:    t.senderBlockchainAddress,
		RecipientBlockchainAddress: t.recipientBlockchainAddress,
		Value:                      t.value,
	})
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
  v := &struct{
    SenderBlockchainAddress    *string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress *string  `json:"recipient_blockchain_address"`
		Value                      *float64 `json:"value"`
  } {
    SenderBlockchainAddress:    &t.senderBlockchainAddress,
		RecipientBlockchainAddress: &t.recipientBlockchainAddress,
		Value:                      &t.value,
  }
  if err := json.Unmarshal(data, &v); err != nil {
    return err
  }
  return nil
}
