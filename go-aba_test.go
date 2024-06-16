package goAba_test

import (
	"testing"

	goAba "github.com/azaurus1/go-aba"
	"gotest.tools/v3/assert"
)

// TODO: use testContainers to generate a fuzzed ABA file with https://github.com/flash-oss/aba-generator/
// then generate using goaba.Generate() and compare

func TestGenerate(t *testing.T) {
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

	ABA := goAba.ABA{
		Header:       Header,
		Transactions: []goAba.Transaction{Transaction},
		Footer:       Footer,
	}

	ABA.Generate()

	assert.Equal(t, ABA.Header.ToString(), "0                 01ANZ       Allowasa Pertolio Accounti001234Credits Of T220524                                        ", "header string does not match")

	for _, transaction := range ABA.Transactions {
		assert.Equal(t, transaction.ToString(), "1061-021   123456 500000001200Georgian Council of New South WaInvoice # 1234    061-123  1234567Bank LLC        00000000", "transaction does not match")
	}

	assert.Equal(t, ABA.Footer.ToString(), "7999-999            000000120000000012000000000000                        000001                                        ", "footer string does not match")
}
