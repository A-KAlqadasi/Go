package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteValueToFile(value float64, fileName string) {
	var byteValue = fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(byteValue), 0644)
	if err != nil {
		fmt.Println("Error writing value to file", err)
	}

}

func ReadValueFromFile(fileName string) (float64, error) {
	var value float64
	data, err := os.ReadFile(fileName)
	
	if err != nil {
		err = errors.New("Failed to read value from file")
		return 0, err
	}

	valueText := string(data)
	value, _= strconv.ParseFloat(valueText, 64)
	return  value, nil
}
