package goAba

import (
	"fmt"
	"strconv"
	"strings"
)

func fillField(maxLen int, field string, justify string, fill string) string {
	if len(field) < maxLen {
		if justify == "right" {
			shifted := maxLen - len(field)
			str := ""
			for i := 0; i < shifted; i++ {
				str += fill
			}
			field = str + field
		} else if justify == "left" {
			shifted := maxLen - len(field)
			str := ""
			for i := 0; i < shifted; i++ {
				str += fill
			}
			field = field + str
		}
		return field
	} else {
		return field[:maxLen]
	}
}

func buildBankNumber(bankStr string) string {
	// string must be numeric with - at pos 5
	firstStr := bankStr[:3]
	secondStr := bankStr[3:]

	bankNumStr := fmt.Sprintf("%s-%s", firstStr, secondStr)
	return bankNumStr
}

func buildTotal(totalAmt float64) string {
	str := strconv.FormatFloat(totalAmt, 'f', 2, 64)
	str = strings.Replace(str, ".", "", -1)

	str = fillField(10, str, "right", "0")
	return str
}

func buildAccNumber(accStr string) string {
	// if more than 9 chars, do not add hyphens
	str := fillField(9, accStr, "right", " ")
	return str
}
