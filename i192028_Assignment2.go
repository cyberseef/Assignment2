package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	Nonce       int
	Transaction []Transaction
	Prevhash    string
	Hash        string
}

// func (blc *Block) CalculateHash1() string {

// 	trans_json, err := json.Marshal(blc.Transaction)
// 	trans := hex.EncodeToString(trans_json[:])

// 	if err != nil {
// 		panic(err)
// 	}
// 	return fmt.Sprintf("%x", sha256.Sum256([]byte(blc.Prevhash+trans+strconv.Itoa(blc.Nonce))))
// }

func (blc *Blocklist) Newblock(n int) *Block {
	s := new(Block)
	s.Nonce = n

	len := len(blc.list)

	if blc.TransactionPool != nil {
		for i := range blc.TransactionPool {
			s.Transaction = append(s.Transaction, blc.TransactionPool[i])
		}

		if len != 0 {
			s.Prevhash = blc.list[len-1].Hash
		} else {
			s.Prevhash = ""
		}

		CalculateHash(blc)
		blc.TransactionPool = nil

		return s
	} else {
		fmt.Printf("No transaction.\n\n")
		return nil
	}
}

type Blocklist struct {
	list            []*Block
	TransactionPool []Transaction
}

type Transaction struct {
	TransactionId              string
	SenderBlockchainAddress    string
	RecipientBlockchainAddress string
	Value                      float32
}

func ListBlocks(obj *Blocklist) {

	var l = len(obj.list)
	for i := l - 1; i >= 0; i-- {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Println("Previous Hash:   ", obj.list[i].Prevhash)
		fmt.Println("Current Hash:   ", obj.list[i].Hash)
		fmt.Println("Nonce:    ", obj.list[i].Nonce)
		fmt.Printf("%s Transactions %s\n", strings.Repeat("-", 22), strings.Repeat("-", 21))

		trans_json, err := json.MarshalIndent(obj.list[i].Transaction, "", "    ")
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n\n", trans_json)

	}

}

func CalculateHash(stud *Blocklist) {
	var l = len(stud.list)
	for i := l - 1; i >= 0; i-- {
		sum := sha256.Sum256([]byte(stud.list[i].GetString()))
		stud.list[i].Hash = fmt.Sprintf("%x", sum)
		if i < len(stud.list)-1 {
			stud.list[i+1].Prevhash = fmt.Sprintf("%x", sum)
		}
	}
}
func (s *Block) GetString() string {

	jsonTransaction, err := json.Marshal(s.Transaction)
	trans := hex.EncodeToString(jsonTransaction[:])

	if err != nil {
		panic(err)
	}

	var r = ""
	r += strconv.Itoa(s.Nonce)
	r += trans + s.Prevhash

	return r
}

func (t *Transaction) SetTransactionId() {
	sum := sha256.Sum256([]byte(t.SenderBlockchainAddress + t.RecipientBlockchainAddress + strconv.FormatFloat(float64(t.Value), 'E', -1, 32)))
	t.TransactionId = hex.EncodeToString(sum[:])
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	trans := new(Transaction)
	trans.SenderBlockchainAddress = sender
	trans.RecipientBlockchainAddress = recipient
	trans.Value = value
	trans.SetTransactionId()
	return trans
}

func (blc *Blocklist) AddTransaction(sender string, recipient string, value float32) {
	trans := NewTransaction(sender, recipient, value)
	blc.TransactionPool = append(blc.TransactionPool, *trans)
}

func main() {
	blc := new(Blocklist)
	blc.AddTransaction("touseef", "gujjar", 1)
	blc.AddTransaction("charlie", "bob", 8)
	blc.AddTransaction("alice", "alex", 99)
	blk := blc.Newblock(999)
	blc.list = append(blc.list, blk)

	blc.AddTransaction("billie", "elish", 3)
	blc.AddTransaction("wiz", "khalifa", 2)
	blc.AddTransaction("Waseem", "Shehzad", 46)
	blk = blc.Newblock(1000)
	blc.list = append(blc.list, blk)
	ListBlocks(blc)

}
