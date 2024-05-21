package goAba

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
