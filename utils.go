package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func CreateAndSaveMatrix(N int) {
	matrix := make([][]int, N)

	// Initialize the matrix
	for i := range matrix {
		matrix[i] = make([]int, N)
		for j := range matrix[i] {
			matrix[i][j] = i*N + j
		}
	}

	// Save the matrix to a file
	file, err := os.Create("matrix.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < N; i++ {
		record := make([]string, N)
		for j := 0; j < N; j++ {
			record[j] = strconv.Itoa(matrix[i][j])
		}
		if err := writer.Write(record); err != nil {
			log.Fatalf("Failed to write record to file: %v", err)
		}
	}

	fmt.Println("Matrix saved to file.")
}

func ReadMatrix(N int) [][]int {
	// Read the matrix back from the file
	file, err := os.Open("matrix.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	matrix := make([][]int, N)
	for i := 0; i < N; i++ {
		record, err := reader.Read()
		if err != nil {
			log.Fatalf("Failed to read record from file: %v", err)
		}
		matrix[i] = make([]int, N)
		for j := 0; j < N; j++ {
			value, err := strconv.Atoi(record[j])
			if err != nil {
				log.Fatalf("Failed to convert string to int: %v", err)
			}
			matrix[i][j] = value
		}
	}
	return matrix
}

func PrintResult(durationColumnMajor time.Duration, durationRowMajor time.Duration, sumColumnMajor int, sumRowMajor int) {
	fmt.Printf("Sum Row-Major...: %d, Time: %v\n\n", sumRowMajor, durationRowMajor)
	fmt.Printf("Sum Column-Major: %d, Time: %v\n\n", sumColumnMajor, durationColumnMajor)

	// Is more faster x times?
	var better string = "Row-Major"
	var slower string = "Column-Major"
	if durationColumnMajor < durationRowMajor {
		better = "Column-Major"
		slower = "Row-Major"
	}
	fmt.Printf("%s is %f times faster than %s\n", better, float64(durationColumnMajor)/float64(durationRowMajor), slower)
}
