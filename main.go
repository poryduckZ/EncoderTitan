package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Encoder Titan!")
	fmt.Println("Please enter the type of encoding you would like to use:")
	fmt.Println("1. Base 64") // TODO: Add more encoding types and use array to store them
	fmt.Print("Enter your choice: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input")
		return
	}
	input = strings.TrimSpace(input)
	option, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error converting input to a number:", err)
		return
	}
	if option == 1 {
		fmt.Print("Please enter the string you would like to encode: ")
		stringToEncode, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			return
		}
		encodedString := base64Encode(stringToEncode)
		fmt.Println("Encoded string:", encodedString)
	}
}

// TODO: write tests for the base64Encode function
func base64Encode(input string) string {
	base64Table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	var encoded string
	var val uint
	var valb int
	inputByteArr := []byte(input)
	for i := 0; i < len(inputByteArr); i++ {
		val = (val << 8) | uint(inputByteArr[i])
		valb += 8

		for valb >= 6 {
			encoded += string(base64Table[(val>>(valb-6))&0x3F])
			valb -= 6
		}
	}

	if valb > 0 {
		encoded += string(base64Table[(val<<(6-valb))&0x3F])
	}

	for len(encoded)%4 != 0 {
		encoded += "="
	}

	return encoded
}
