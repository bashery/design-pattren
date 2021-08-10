/*
package diffsquares

func SquareOfSum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i * i
	}
	return sum
}
func Difference(i int) int {
	return SquareOfSum(i) - SumOfSquares(i)
}
*/
package diffsquares

//SquareOfSum takes as input the first N natural numbers and return the squar root of their sum
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

//SumOfSquares takes as input the first N natural numbers and return the sum of the squart root of each number
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

//Difference should return diference between two numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
