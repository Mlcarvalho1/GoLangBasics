package files

import (
	"fmt"
	"os"
)

func WriteFloatInFile(fileName string, balance float64) {
	os.WriteFile(fileName, []byte(fmt.Sprint(balance)), 0644)
}

func GetFloatFromFile(fileName string) (balance float64) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0
	}

	fmt.Sscan(string(data), &balance)

	return balance
}
