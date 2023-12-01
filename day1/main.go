package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fileContent := loadDocument("input.txt")
	calibrationSum := calculateTotalValue(fileContent)
	println("Calibration Sum Value is:", calibrationSum)
}

func loadDocument(documentPath string) *os.File {
	fileContent, err := os.Open(documentPath)
	if err != nil {
		fmt.Println("Err reading file")
	}
	return fileContent
}

func calculateTotalValue(fileContent *os.File) int {
	var total int
	fileScanner := bufio.NewScanner(fileContent)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		total = total + getCalibrationValue(fileScanner.Text())
	}

	return total
}

func getCalibrationValue(line string) int {
	var calibration []string

	for _, char := range line {
		if !unicode.IsDigit(char) {
			continue
		} else {
			calibration = append(calibration, string(char))
		}
	}
	calibrationSize := len(calibration)
	calibrationResult, _ := strconv.Atoi(calibration[0] + calibration[calibrationSize-1])
	println("For line", line, "Calibration Value is:", calibrationResult)
	return calibrationResult
}
