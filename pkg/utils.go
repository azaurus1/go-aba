package goAba

import (
	"fmt"
)

func fillField(maxLen int, field string, justify string, fill string) string {
	if justify == "right" {
		if len(field) < maxLen {
			shifted := maxLen - len(field)
			str := ""
			for i := 0; i < shifted; i++ {
				str += fill
			}
			field = str + field
		}
	}
	return field
}

func buildBankNumber(bankStr string) string {
	// string must be numeric with - at pos 5
	firstStr := bankStr[:3]
	secondStr := bankStr[3:]

	bankNumStr := fmt.Sprintf("%s-%s", firstStr, secondStr)
	return bankNumStr
}
