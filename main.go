package main

import (
	"learn-go/questions"
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math/rand"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	score := 0

	for {
		q := questions.Question{}
		ops := []string{"+", "-", "*", "/", "%"}
		q.Generate_easy(ops[rand.Intn(len(ops))])
		fmt.Printf("What is %d %s %d? ", q.Num1, q.Operation, q.Num2)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if input == "q" {
			fmt.Println("Game ended. Final score:", score)
			break
		}
		userAnswer, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		if userAnswer == q.Answer {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect. Correct answer:", q.Answer)
		}
	}
}


/*
func main() {
	q := questions.Question{}
	q.Generate_hard("-")
	fmt.Println(q.Num1, q.Operation, q.Num2, "= ")
	var answer int
	fmt.Scanf("%d", &answer)
	fmt.Println(answer == q.Answer)
}
*/