package goAba

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func (aba *ABA) Generate() error {
	var transactionSlice []string

	headerStr := aba.GenerateHeader()
	transactionSlice = aba.GenerateTransactions()
	footerStr := aba.GenerateFooter()
	log.Println(headerStr)
	log.Println(transactionSlice)
	log.Println(footerStr)
	return nil
}

func (aba *ABA) GenerateHeader() string {
	userTrunc := aba.Header.User[0:26] // Truncate user name to 26 chars

	userNum := strconv.Itoa(aba.Header.UserNumber)
	if len(userNum) > 6 {
		userNum = userNum[0:6] // Too long, truncate user number
	} else if len(userNum) < 6 { // else shift to the right and prepend 0's
		userNum = fillField(6, userNum, "right", "0")
	}

	descTrunc := aba.Header.Description[0:12] // truncate description

	if aba.Header.Date == "" {
		// Get the current time
		now := time.Now()

		// Format the time as DDMMYY
		aba.Header.Date = now.Format("020106")
	}

	headerStr := fmt.Sprintf("0                 01%s       %s%s%s%s                                        ", aba.Header.Bank, userTrunc, userNum, descTrunc, aba.Header.Date)
	return headerStr
}

func (aba *ABA) GenerateTransactions() []string {
	var transactions []string
	for _, transaction := range aba.Transactions {
		bankNum := buildBankNumber(transaction.BSB)

		amount := strconv.Itoa(int(transaction.Amount))
		amount = fillField(10, amount, "right", "0")

		transactionStr := fmt.Sprintf("1%s   %s %s%s%s%s", bankNum, transaction.Account, transaction.TransactionCode, amount, transaction.AccountTitle, transaction.Reference)
		transactions = append(transactions, transactionStr)
	}
	return transactions
}

func (aba *ABA) GenerateFooter() string {

	netTotalAmt := aba.Footer.NetTotal
	netTotalAmt = fillField(10, netTotalAmt, "right", "0")

	creditTotalAmt := aba.Footer.CreditTotal
	creditTotalAmt = fillField(10, creditTotalAmt, "right", "0")

	debitTotalAmt := aba.Footer.DebitTotal
	debitTotalAmt = fillField(10, debitTotalAmt, "right", "0")

	countOfRecords := strconv.Itoa(len(aba.Transactions))
	countOfRecords = fillField(6, countOfRecords, "right", "0")

	footerStr := fmt.Sprintf("7999-999            %s%s%s                        %s                                        ", netTotalAmt, creditTotalAmt, debitTotalAmt, countOfRecords)
	return footerStr
}
