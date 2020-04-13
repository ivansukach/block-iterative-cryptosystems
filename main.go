package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"os"
	"strconv"
	"strings"
)

var sTemplate8 = make(map[rune]rune)
var sTemplate4 = make(map[rune]rune)
var hexToBinary = make(map[rune]string)

func init() {
	sTemplate8['0'] = 'B'
	sTemplate8['1'] = 'A'
	sTemplate8['2'] = 'F'
	sTemplate8['3'] = '5'
	sTemplate8['4'] = '0'
	sTemplate8['5'] = 'C'
	sTemplate8['6'] = 'E'
	sTemplate8['7'] = '8'
	sTemplate8['8'] = '6'
	sTemplate8['9'] = '2'
	sTemplate8['A'] = '3'
	sTemplate8['B'] = '9'
	sTemplate8['C'] = '1'
	sTemplate8['D'] = '7'
	sTemplate8['E'] = 'D'
	sTemplate8['F'] = '4'

	sTemplate4['0'] = 'E'
	sTemplate4['1'] = '7'
	sTemplate4['2'] = 'A'
	sTemplate4['3'] = 'C'
	sTemplate4['4'] = 'D'
	sTemplate4['5'] = '1'
	sTemplate4['6'] = '3'
	sTemplate4['7'] = '9'
	sTemplate4['8'] = '0'
	sTemplate4['9'] = '2'
	sTemplate4['A'] = 'B'
	sTemplate4['B'] = '4'
	sTemplate4['C'] = 'F'
	sTemplate4['D'] = '8'
	sTemplate4['E'] = '5'
	sTemplate4['F'] = '6'

	hexToBinary['0'] = "0000"
	hexToBinary['1'] = "0001"
	hexToBinary['2'] = "0010"
	hexToBinary['3'] = "0011"
	hexToBinary['4'] = "0100"
	hexToBinary['5'] = "0101"
	hexToBinary['6'] = "0110"
	hexToBinary['7'] = "0111"
	hexToBinary['8'] = "1000"
	hexToBinary['9'] = "1001"
	hexToBinary['A'] = "1010"
	hexToBinary['B'] = "1011"
	hexToBinary['C'] = "1100"
	hexToBinary['D'] = "1101"
	hexToBinary['E'] = "1110"
	hexToBinary['F'] = "1111"
}
func S8(T1 string) rune {
	var arg rune
	for _, value := range T1 {
		arg = value
	}
	return sTemplate8[arg]
}
func S4(T2 string) rune {
	var arg rune
	for _, value := range T2 {
		arg = value
	}
	return sTemplate4[arg]
}
func P(input string) string {
	runes := []rune(input)
	substring := string(runes[0:3])
	input = string(runes[3:])
	log.Infoln("input after cutting: ", input)
	output := input + substring
	return output
}
func getX(number int) string {
	log.Info("number: ", number)
	X := number * 7
	log.Info("X: ", X)
	log.Infof("X binary: %b\n", X)
	message := fmt.Sprintf("%b", X)
	if len(message) < 8 {
		for i := 0; i < (8 - len(message)); i++ {
			message = "0" + message
		}
	}
	return message
}
func BinaryToHex(arg string) string {
	i := len(arg) - 1
	hexValue := 0
	for _, value := range arg {
		term := int(value - '0')
		if term == 1 {
			hexValue += int(math.Pow(2, float64(i)))
		}
		i--
	}
	return strings.ToUpper(fmt.Sprintf("%x", hexValue))
}
func startEncryption(q int, r int, x string) string {
	mainKey := 4096 - 11*q*r
	log.Info("mainKey: ", mainKey)
	log.Infof("mainKey binary: %b\n", mainKey)
	mainKeyBinary := fmt.Sprintf("%b", mainKey)
	var firstRoundKey, secondRoundKey, thirdRoundKey string
	firstRoundKeyScheme := [8]int{1, 3, 5, 7, 2, 4, 6, 8}
	secondRoundKeyScheme := [8]int{5, 7, 9, 11, 6, 8, 10, 12}
	thirdRoundKeyScheme := [8]int{12, 10, 4, 2, 1, 3, 9, 11}
	mainKeyBinarySlice := make([]rune, 0)
	for _, value := range mainKeyBinary {
		mainKeyBinarySlice = append(mainKeyBinarySlice, value)
	}
	for i := 7; i >= 0; i-- {
		firstRoundKey += string(mainKeyBinarySlice[firstRoundKeyScheme[i]-1])
		secondRoundKey += string(mainKeyBinarySlice[secondRoundKeyScheme[i]-1])
		thirdRoundKey += string(mainKeyBinarySlice[thirdRoundKeyScheme[i]-1])
	}
	log.Info("Round keys: ")
	log.Info("first: ", firstRoundKey)
	log.Info("second: ", secondRoundKey)
	log.Info("third: ", thirdRoundKey)
	log.Info("===round 1===")
	result1Iteration := round(x, firstRoundKey)
	log.Info("===round 2===")
	result2Iteration := round(result1Iteration, secondRoundKey)
	log.Info("===round 3===")
	return round(result2Iteration, secondRoundKey)
}
func round(message string, roundKey string) string {
	messageBinarySlice := make([]rune, 0)
	for _, value := range message {
		messageBinarySlice = append(messageBinarySlice, value)
	}
	roundKeyBinarySlice := make([]rune, 0)
	for _, value := range roundKey {
		roundKeyBinarySlice = append(roundKeyBinarySlice, value)
	}
	result := ""
	for i := 0; i < 8; i++ {
		result += fmt.Sprintf("%b", uint8(messageBinarySlice[i]-'0')^uint8(roundKeyBinarySlice[i]-'0'))
	}
	log.Info("result after XOR: ", result)
	firstPart := ""
	iteration := 0
	secondPart := ""
	for _, value := range result {
		if iteration < 4 {
			firstPart += string(value)
		} else {
			secondPart += string(value)
		}
		iteration++
	}
	log.Info("first part", firstPart)
	log.Info("second part", secondPart)
	firstPartHex := BinaryToHex(firstPart)
	secondPartHex := BinaryToHex(secondPart)
	log.Info("first part in Hex", firstPartHex)
	log.Info("second part in Hex", secondPartHex)
	N1Hex := S8(firstPartHex)
	N2Hex := S4(secondPartHex)
	log.Info("first part in Hex after S", string(N1Hex))
	log.Info("second part in Hex after S", string(N2Hex))
	N1 := hexToBinary[N1Hex]
	N2 := hexToBinary[N2Hex]
	inputDataP := N1 + N2
	log.Infoln("input data to P: ", inputDataP)
	output := P(inputDataP)
	log.Infoln("OUTPUT: ", output)
	return output
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	log.Info("Input your number in group list")
	n, _, _ := reader.ReadLine()
	numberOfStudent := string(n)
	number, err := strconv.Atoi(numberOfStudent)
	if err != nil {
		log.Error(err)
	}
	X := getX(number)
	var X2 string
	withoutLatestRune := []rune(X)[0:7]
	latestRune := []rune(X)[7:8]
	if latestRune[0] == '0' {
		X2 = string(withoutLatestRune) + "1"
	} else {
		X2 = string(withoutLatestRune) + "0"
	}
	log.Info("Input amount of letters in your name")
	n, _, _ = reader.ReadLine()
	amountOfCharsInName := string(n)
	amountCharsName, err := strconv.Atoi(amountOfCharsInName)
	if err != nil {
		log.Error(err)
	}

	log.Info("Input amount of letters in your surname")
	n, _, _ = reader.ReadLine()
	amountOfCharsInSurname := string(n)
	amountCharsSurname, err := strconv.Atoi(amountOfCharsInSurname)
	if err != nil {
		log.Error(err)
	}
	encryptedText := startEncryption(amountCharsName, amountCharsSurname, X)
	log.Info()
	log.Info("RESULT")
	log.Info("encryptedText: ", encryptedText)
	log.Info()
	researchAvalancheEffect := startEncryption(amountCharsName, amountCharsSurname, X2)
	log.Info()
	log.Info("RESULT")
	log.Info("Research for Avalanche Effect(1 bit has been changed): ", researchAvalancheEffect)
}
