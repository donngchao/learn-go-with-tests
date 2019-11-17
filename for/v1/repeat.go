package iteration

// Repeat returns character repeated 5 times
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}

func repeatAgain(subStr string) string {
	var repeatStr string
	for i:=0 ; i< 6; i++{
		repeatStr += subStr
	}
	return repeatStr
}
