package bitmedia_api

type Block struct {
	JSONRpc string `json:"jsonrpc" bson:"jsonrpc"`
	ID      int    `json:"id" bson:"id"`
	Result  string `json:"result" bson:"_id"`
}

type BlockByNumber struct {
	id      string `bson:"_id"`
	JSONRpc string `json:"jsonrpc" bson:"jsonrpc"`
	ID      int    `json:"id" bson:"id"`
	Result  Result `json:"result" bson:"result"`
}

type Result struct {
	BaseFeePerGas string         `json:"baseFeePerGas" bson:"baseFeePerGas"`
	Timestamp     string         `json:"timestamp" bson:"timestamp"`
	Transactions  []Transactions `json:"transactions" bson:"transactions"`
}

type Transactions struct {
	BlockNumber string `json:"blockNumber" bson:"blockNumber"`
	From        string `json:"from" bson:"from"`
	To          string `json:"to" bson:"to"`
	GasPrice    string `json:"gasPrice" bson:"gasPrice"`
	Hash        string `json:"hash" bson:"hash"`
}
