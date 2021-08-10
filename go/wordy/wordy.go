package wordy

import (
	//"fmt"
	"strconv"
	"strings"
)

// What is 5?
// What is 5 pluse 5?

func Answer(q string) (int, bool) {
	q = strings.Trim(q, "What is ?")
	q = strings.Replace(q, "by ", "", 2)
	op := strings.Split(q, " ")
	leng := len(op)

	if leng == 1 {
		if op[0] == "" {
			return 0, false
		}
		first, _ := strconv.Atoi(op[0])
		return first, true
	} else if leng == 3 {
		first, _ := strconv.Atoi(op[0])
		second, _ := strconv.Atoi(op[2])
		switch {
		case op[1] == "plus":
			return first + second, true
		case op[1] == "minus":
			return first - second, true
		case op[1] == "multiplied":
			return first * second, true
		case op[1] == "divided":
			return first / second, true
		default:
		}
	} else if leng == 5 {
		first, _ := strconv.Atoi(op[0])
		second, _ := strconv.Atoi(op[2])
		third, _ := strconv.Atoi(op[4])
		firstOp, secondOp := op[1], op[3]
		switch {
		case firstOp == "plus":
			if secondOp == "plus" {
				return first + second + third, true
			} else if secondOp == "minus" {
				return first + second - third, true
			} else if secondOp == "multiplied" {
				return (first + second) * third, true // i not aree witn this
			} else if secondOp == "divided" {
				return first + second/third, true
			}
		case firstOp == "minus":
			if secondOp == "plus" {
				return first - second + third, true
			} else if secondOp == "minus" {
				return first - second - third, true
			} else if secondOp == "multiplied" {
				return first - second*third, true
			} else if secondOp == "divided" {
				return first - second/third, true
			}
		case firstOp == "multiplied":
			if secondOp == "plus" {
				return first*second + third, true
			} else if secondOp == "minus" {
				return first*second - third, true
			} else if secondOp == "multiplied" {
				return first * second * third, true
			} else if secondOp == "divided" {
				return first * second / third, true
			}
		case firstOp == "divided":
			if secondOp == "plus" {
				return first/second + third, true
			} else if secondOp == "minus" {
				return first/second - third, true
			} else if secondOp == "multiplied" {
				return first / second * third, true
			} else if secondOp == "divided" {
				return first / second / third, true
			}
		default:
		}
	}
	return 0, false
}
