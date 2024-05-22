# go-aba

A Go library for generating ABA banking files.

# Usage
Install:

`go get github.com/azaurus1/go-aba`

Example:
```go
package main

import goAba "github.com/azaurus1/go-aba/pkg"

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

	ABA := goAba.ABA{
		Header:       Header,
		Transactions: []goAba.Transaction{Transaction},
		Footer:       Footer,
	}

	ABA.Generate()
}

```

The above will return:
```
0                 01ANZ       Allowasa Pertolio Accounti001234Credits Of T220524                                        
1061-021   123456 500000001200Georgian Council of New South WaInvoice # 1234    061-123  1234567Bank LLC        00000000
7999-999            000000120000000012000000000000                        000001
```

# Sources
1. https://github.com/flash-oss/aba-generator
2. http://ddkonline.blogspot.com/2009/01/aba-bank-payment-file-format-australian.html ([archive](https://web.archive.org/web/20240202085445/http://ddkonline.blogspot.com/2009/01/aba-bank-payment-file-format-australian.html))

# Credits
Many thanks to [Flash OSS](https://github.com/flash-oss) for the [aba-generator](https://github.com/flash-oss/aba-generator) which I used for reference in building this library.
