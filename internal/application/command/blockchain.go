package command

type TransactionCreate struct {
  SenderBlockchainAddress    *string  `json:"sender_blockchain_address"`
  RecipientBlockchainAddress *string  `json:"recipient_blockchain_address"`
  SenderPublicKey            *string  `json:"sender_public_key"`
	Value                      *float64 `json:"value"`
  Signature                  *string  `json:"signature"`
}
