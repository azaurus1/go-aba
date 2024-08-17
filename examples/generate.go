package main

import (
	"encoding/json"
	"log"

	goAba "github.com/azaurus1/go-aba"
)

type depayload struct {
	Header       goAba.Header        `json:"header"`
	Transactions []goAba.Transaction `json:"transactions"`
	Footer       goAba.Footer        `json:"footer"`
	FileName     string              `json:"filename"`
}

func main() {
	Header := goAba.Header{
		Bank:        "ANZ",
		User:        "Allowasa Pertolio Accounting&Tax",
		UserNumber:  "1234",
		Description: "Credits Of The Wooloomooloo",
	}
	Transaction := goAba.Transaction{
		BSB:             "061021",
		TransactionCode: goAba.Credit,
		Account:         "123456",
		Amount:          12.0,
		AccountTitle:    "Georgian Council of New South Wales",
		Reference:       "Invoice # 1234",
		TraceBSB:        "061123",
		TraceAccount:    "1234567",
		Remitter:        "Bank LLC",
	}
	Footer := goAba.Footer{}

	// ABA := goAba.ABA{
	// 	Header:       Header,
	// 	Transactions: []goAba.Transaction{Transaction},
	// 	Footer:       Footer,
	// }

	// str, _ := ABA.Generate()

	dePayload := depayload{
		Header:       Header,
		Transactions: []goAba.Transaction{Transaction},
		Footer:       Footer,
		FileName:     "test.aba",
	}

	marshalledStr, _ := json.Marshal(dePayload)

	log.Println(string(marshalledStr))
}
