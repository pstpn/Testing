package main

import (
	"fmt"
	"io"
	_ "io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func halsteadMetrics(code string) float64 {
	operators := []string{"+", "-", "*", "/", "=", ">", "<", "&&", "||", "!"}
	operands := strings.Fields(code)

	uniqueOperators := make(map[string]bool)
	uniqueOperands := make(map[string]bool)

	operatorCount := 0
	operandCount := 0

	for _, token := range operands {
		if contains(operators, token) {
			uniqueOperators[token] = true
			operatorCount++
		} else {
			uniqueOperands[token] = true
			operandCount++
		}
	}

	n1 := len(uniqueOperators)
	n2 := len(uniqueOperands)
	N2 := operandCount

	var difficulty float64

	if n1 > 0 && n2 > 0 {
		difficulty = (float64(n1) / 2) * (float64(N2) / float64(n2))
	}

	return difficulty
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func processFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Ошибка открытия файла %s: %v\n", path, err)
		return
	}

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Ошибка чтения файла %s: %v\n", path, err)
		return
	}

	fmt.Printf("%s: %f\n", path, halsteadMetrics(string(content)))
}

func main() {
	root := "."

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			processFile(path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Ошибка обхода директории: %v\n", err)
	}
}
