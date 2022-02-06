package bitmedia_api

type Block struct {
	JSONRpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

type BlockByNumber struct {
	JSONRpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Result `json:"result"`
}

type Result struct {
	BaseFeePerGas string         `json:"baseFeePerGas"`
	Timestamp     string         `json:"timestamp"`
	Transactions  []Transactions `json:"transactions,omitempty"`
}

type Transactions struct {
	BlockNumber string `json:"blockNumber" bson:"blockNumber"`
	From        string `json:"from" bson:"from"`
	To          string `json:"to" bson:"to"`
	GasPrice    string `json:"gasPrice" bson:"gasPrice"`
	Hash        string `json:"hash" bson:"hash"`
}

type Transaction struct {
	BaseFeePerGas string `json:"baseFeePerGas" bson:"baseFeePerGas"`
	Timestamp     string `json:"timestamp" bson:"timestamp"`
	BlockNumber   string `json:"blockNumber" bson:"blockNumber"`
	From          string `json:"from" bson:"from"`
	To            string `json:"to" bson:"to"`
	GasPrice      string `json:"gasPrice" bson:"gasPrice"`
	Hash          string `json:"hash" bson:"hash"`
}
