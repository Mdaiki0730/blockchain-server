package result

type Blockchain struct {
	Blocks []*Block `json:"chain"`
}

type Block struct {
	Timestamp    int64          `json:"timestamp"`
	Nonce        int            `json:"nonce"`
	PreviousHash string         `json:"previous_hash"`
	Transactions []*Transaction `json:"transactions"`
}

type Transaction struct {
	SenderBlockchainAddress    *string  `json:"sender_blockchain_address"`
	RecipientBlockchainAddress *string  `json:"recipient_blockchain_address"`
	Value                      *float64 `json:"value"`
}

type Transactions struct {
	Transactions []*Transaction `json:"transactions"`
	Length       int            `json:"length"`
}
