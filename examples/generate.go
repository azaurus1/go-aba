package main

import goAba "github.com/azaurus1/go-aba/pkg"

func main() {
	Header := goAba.Header{
		Bank:        "ANZ",
		User:        "Allowasa Pertolio Accounting&Tax",
		UserNumber:  1234,
		Description: "Credits Of The Wooloomooloo",
	}
	Transaction := goAba.Transaction{
		BSB:             "061021",
		TransactionCode: "50",
		Account:         "123456",
		Amount:          12.0,
		AccountTitle:    "Georgian Council of New South Wales",
		Reference:       "Invoice # 1234",
		TraceBSB:        "061123",
		TraceAccount:    "1234567",
		Remitter:        "Bank LLC",
	}
	Footer := goAba.Footer{}

	ABA := goAba.ABA{
		Header:       Header,
		Transactions: []goAba.Transaction{Transaction},
		Footer:       Footer,
	}

	ABA.Generate()
}
