package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(filePath string) (numbers []int) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	var line int
	for {

		_, err := fmt.Fscanf(fd, "%d\n", &line)

		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				return
			}
			panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))

		}
		numbers = append(numbers, line)
	}
}

func main() {
	numbers := readFile("Day1Input.txt")
	fmt.Println(len(numbers))

	A := 0
	B := 0

	NumberofIncrease := 0

	for i := 0; i < len(numbers); i++ {

		if i < len(numbers)-3 {
			A = numbers[i] + numbers[i+1] + numbers[i+2]
			B = numbers[i+1] + numbers[i+2] + numbers[i+3]

			if B > A {
				NumberofIncrease++
				println("B is", B, "A is", A)
			}
		}

	}
	println("There were", NumberofIncrease, "increases")

}
