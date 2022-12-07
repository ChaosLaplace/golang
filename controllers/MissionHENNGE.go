package controllers

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
)

/*
Description
We want you to calculate the sum of squares of given integers, excluding any negatives.
The first line of the input will be an integer N (1 <= N <= 100), indicating the number of test cases to follow.
Each of the test cases will consist of a line with an integer X (0 < X <= 100), followed by another line consisting of X number of space-separated integers Yn (-100 <= Yn <= 100).
For each test case, calculate the sum of squares of the integers, excluding any negatives, and print the calculated sum in the output.
Note: There should be no output until all the input has been received.
Note 2: Do not put blank lines between test cases solutions.
Note 3: Take input from standard input, and output to standard output.
Rules
Write your solution using Go Programming Language or Python Programming Language. Do not submit your solution with both languages at once!

You may only use standard library packages. In addition, extra point is awarded if solution does not declare any global variables.

Specific rules for Go solution
Your source code must be a single file
Do not use any for and goto statement
*/

func MissionHENNGE(c *gin.Context) {
	resultChan := make(chan int, 10)
	var frequency int
	fmt.Println("Frequency")
	fmt.Scanf("%d", &frequency)
	if !validate("frequency", frequency) {
		fmt.Printf("Invalidate frequency=%d (1 <= N <= 100)", frequency)
		return
	}
	// 迴圈次數
	status := loop(frequency, resultChan)
	if status {
		result(frequency, resultChan)
	}
}

// 代替迴圈
func loop(frequency int, resultChan chan int) bool {
	var (
		number int
		result int
	)

	frequency--

	fmt.Scanf("%d", &number)
	if !validate("number", number) {
		fmt.Printf("Invalidate number=%d (0 < X <= 100)", number)
		return false
	}

	inputReader := bufio.NewReader(os.Stdin)
	numberToSquare, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("Invalidate numberToSquare err=%v", err)
		return false
	}

	numberToSquare = strings.Replace(numberToSquare, "\n", "", -1)
	numberToSquareArr := strings.Split(numberToSquare, " ")
	count := len(numberToSquareArr)
	if number != count {
		fmt.Printf("Invalidate numberToSquareArr | number=%d | len(numberToSquareArr)=%d", number, count)
		return false
	}

	square(numberToSquareArr, count, result, resultChan)

	if frequency > 0 {
		loop(frequency, resultChan)
	}
	return true
}

// 計算平方
func square(numberToSquareArr []string, count int, result int, resultChan chan int) {
	count--

	squareInt, err := strconv.Atoi(numberToSquareArr[count])
	if err != nil {
		fmt.Printf("Invalidate square err=%v", err)
		return
	}

	if !validate("numberToSquare", squareInt) {
		fmt.Printf("Invalidate squareInt=%d (-100 <= Yn <= 100) \n", squareInt)
		return
	}

	if squareInt > 0 {
		result += squareInt * squareInt
	}

	if count > 0 {
		square(numberToSquareArr, count, result, resultChan)
		return
	}

	resultChan <- result
	return
}

// 執行結果
func result(frequency int, resultChan chan int) {
	frequency--

	if len(resultChan) > 0 {
		fmt.Printf("%d \n", <-resultChan)
	}

	if frequency > 0 {
		result(frequency, resultChan)
	}
	return
}

// 驗證參數
func validate(tag string, intRange int) bool {
	switch tag {
	case "frequency":
		// 測試次數 1 <= N <= 100
		if 1 <= intRange && intRange <= 100 {
			return true
		}
	case "number":
		// 輸入幾個整數 0 < X <= 100
		if 0 < intRange && intRange <= 100 {
			return true
		}
	case "numberToSquare":
		// 整數範圍 -100 <= Yn <= 100
		if -100 <= intRange && intRange <= 100 {
			return true
		}
	}
	return false
}
