package diffsquares

func SquareOfSum(input int) int{
	return (input*(input+1)/2)*(input*(input+1)/2)
}
func SumOfSquares(input int) int{
	sumSquare := 0
	for i:=0;i<=input;i++ {
		sumSquare = sumSquare + i*i
	}
	return sumSquare
}
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}