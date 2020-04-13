package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./prime.txt")
	if err != nil {
		log.Fatal(err)
	}

	PrimeFile, err := os.Create("./primeOutput.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		isPrime := 0
		primeInt, err := strconv.Atoi(scanner.Text())
		primeMe := prime(primeInt)

		if err != nil {
			panic(err)
		}

		if primeMe {
			isPrime = 1
		}

		PrimeFile.WriteString(strconv.Itoa(isPrime) + "\n")
	}

	if err := scanner.Err(); err != nil {

		log.Fatal(err)
	}

}

func prime(primeVal int) bool {

	if primeVal < 1 {
		return false
	}

	for i := 2; i < primeVal; i++ {
		value := primeVal % i
		if value == 0 {
			return false
		}
	}

	return true
}
