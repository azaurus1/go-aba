package goAba

import (
	"log"
	"strconv"
	"time"
)

const (
	Debit           = "13"
	Credit          = "50"
	AusGovSecInt    = "51"
	FamilyAllowance = "52"
	Pay             = "53"
	Pension         = "54"
	Allotment       = "55"
	Dividend        = "56"
	Debenture       = "57"
)

func (aba *ABA) Generate() error {
	var transactionSlice []string

	headerStr := aba.GenerateHeader()
	transactionSlice = aba.GenerateTransactions()
	footerStr := aba.GenerateFooter()

	log.Println(headerStr)
	for _, transaction := range transactionSlice {
		log.Println(transaction)
	}
	log.Println(footerStr)

	return nil
}

func (aba *ABA) GenerateHeader() string {
	aba.Header.User = aba.Header.User[0:26] // Truncate user name to 26 chars

	if len(aba.Header.UserNumber) > 6 {
		aba.Header.UserNumber = aba.Header.UserNumber[0:6] // Too long, truncate user number
	} else if len(aba.Header.UserNumber) < 6 { // else shift to the right and prepend 0's
		aba.Header.UserNumber = fillField(6, aba.Header.UserNumber, "right", "0")
	}

	aba.Header.Description = aba.Header.Description[0:12] // truncate description

	if aba.Header.Date == "" {
		// Get the current time
		now := time.Now()

		// Format the time as DDMMYY
		aba.Header.Date = now.Format("020106")
	}

	return aba.Header.ToString()
}

func (aba *ABA) GenerateTransactions() []string {
	var transactions []string
	var transactionStr string
	for _, transaction := range aba.Transactions {
		transactionStr = transaction.ToString()
		transactions = append(transactions, transactionStr)
	}
	return transactions
}

func (aba *ABA) GenerateFooter() string {
	var netTotalAmt float64
	var creditTotalAmt float64
	var debitTotalAmt float64

	var netTotalStr string
	var creditTotalStr string
	var debitTotalStr string

	// net total: credit - debit, unsigned
	// build from transactions
	for _, transaction := range aba.Transactions {
		if transaction.TransactionCode == Credit {
			// add to credit total
			creditTotalAmt += transaction.Amount
		} else if transaction.TransactionCode == Debit {
			// add to debit total amount
			debitTotalAmt += transaction.Amount
		}
	}
	// build netTotal
	if creditTotalAmt > debitTotalAmt {
		netTotalAmt = creditTotalAmt - debitTotalAmt
	} else if creditTotalAmt < debitTotalAmt {
		netTotalAmt = debitTotalAmt - creditTotalAmt
	} else {
		// throw it away, it cannot be 0
		log.Println("net ammount cannot be 0")
		return ""
	}

	netTotalStr = buildTotal(netTotalAmt)
	creditTotalStr = buildTotal(creditTotalAmt)
	debitTotalStr = buildTotal(debitTotalAmt)

	aba.Footer.NetTotal = netTotalStr
	aba.Footer.CreditTotal = creditTotalStr
	aba.Footer.DebitTotal = debitTotalStr

	countOfRecords := strconv.Itoa(len(aba.Transactions))
	aba.Footer.NumberOfTransactions = fillField(6, countOfRecords, "right", "0")

	return aba.Footer.ToString()
}
