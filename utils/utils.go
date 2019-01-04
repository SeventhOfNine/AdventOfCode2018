package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFromFile(f string) []string {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	dataString := strings.Split(string(data), "\n")

	return dataString
}

func ReadStringFromFile(f string) string {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	dataString := strings.Replace(string(data), "\n", "", -1)

	return dataString
}

func FormatInt(val string) int {
	i, err := strconv.Atoi(strings.TrimSpace(val))
	if err != nil {
		fmt.Println("cannot convert to int")
	}

	return i
}
