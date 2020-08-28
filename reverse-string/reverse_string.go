package reverse



func Reverse(input string) string{
	runeString := []rune(input)
	for i, j := 0, len(runeString)-1; i < j; i, j = i+1, j-1 { 
		runeString[i],runeString[j]=runeString[j],runeString[i]
	}
	return string(runeString)
}