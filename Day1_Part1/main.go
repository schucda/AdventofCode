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

	j := 0
	NumberofIncrease := 0

	for i := 0; i < len(numbers); i++ {
		if j < len(numbers)-1 {
			j++
		}
		if numbers[j] > numbers[i] {
			NumberofIncrease++
		}

	}
	println("There were", NumberofIncrease, "increases")
}
