package types

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Header struct {
	ParentHash  string    `json:"parentHash"       gencodec:"required"`
	UncleHash   string    `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    string `json:"miner"            gencodec:"required"`
	Root        string    `json:"stateRoot"        gencodec:"required"`
	TxHash      string    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash string    `json:"receiptsRoot"     gencodec:"required"`
	Bloom       types.Bloom          `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
	Number      *big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Time        uint64         `json:"timestamp"        gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`
	MixDigest   string   `json:"mixHash"`
	Nonce       uint32     `json:"nonce"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`
}
// Transaction LegacyTx is the transaction data of regular Ethereum transactions.
type Transaction struct {
	BlockHash string `json:"blockHash,omitempty"`
	BlockNumber string `json:"blockNumber,omitempty"`
	From string `json:"from,omitempty"`
	To string `json:"to,omitempty"`
	Data string `json:"input,omitempty"`
	TransactionIndex string `json:"transactionIndex,omitempty"`
	Value string `json:"value,omitempty"`
	Type string
	V, R, S  *big.Int        // signature values


}

type Transactions []*Transaction
//Block ETH 交易结构体
type Block struct {
	Header       *Header
	Uncles       []*Header
	//TODO 保留ETH的交易结构类型还是替换为chain33的Transaction
	Transactions  Transactions
	//Transactions chain33Types.Transactions
	// caches
	Hash string
}




// Receipt represents the results of a transaction.
type Receipt struct {
	// Consensus fields: These fields are defined by the Yellow Paper
	Type              string  `json:"type,omitempty"`
	PostState         []byte `json:"root"`
	//0x1 表示交易成功，0x2 表示交易失败
	Status            string `json:"status"`
	CumulativeGasUsed uint64 `json:"cumulativeGasUsed" gencodec:"required"`
	//Bloom             Bloom  `json:"logsBloom"         gencodec:"required"`
	Logs              []string `json:"logs"              gencodec:"required"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	// They are stored in the chain database.
	TxHash          string   `json:"transactionHash" gencodec:"required"`
	ContractAddress string `json:"contractAddress"`
	GasUsed         uint64         `json:"gasUsed" gencodec:"required"`
	To string `json:"to,omitempty"`
	From string `json:"from,omitempty"`
	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        string `json:"blockHash,omitempty"`
	//十六进制
	BlockNumber      string    `json:"blockNumber,omitempty"`
	TransactionIndex string        `json:"transactionIndex"`
}

type CallMsg struct {
	From    string  `json:"from"`
	To       string`json:"to"`
	Gas      *hexutil.Uint64 `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    *hexutil.Big    `json:"value"`
	Nonce    *hexutil.Uint64 `json:"nonce"`
	// We accept "data" and "input" for backwards-compatibility reasons. "input" is the
	// newer name and should be preferred by clients.
	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`
}