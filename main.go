package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func getX(number int) int {
	log.Println("number: ", number)
	X := number * 7
	log.Println("X: ", X)
	log.Printf("X в двоичной системе: %b\n", X)
	return X
}
func getKey(q int, r int) int {
	mainKey := 4096 - 11*q*r
	log.Println("mainKey: ", mainKey)
	log.Printf("mainKey в двоичной системе: %b\n", mainKey)
	mainKeyBinary := fmt.Sprintf("%b", mainKey)
	var firstRoundKey, secondRoundKey, thirdRoundKey string
	firstRoundKeySchemе := [8]int{1, 3, 5, 7, 2, 4, 6, 8}
	secondRoundKeySchemе := [8]int{5, 7, 9, 11, 6, 8, 10, 12}
	thirdRoundKeySchemе := [8]int{12, 10, 4, 2, 1, 3, 9, 11}
	mainKeyBinarySlice := make([]rune, 0)
	for index, value := range mainKeyBinary {
		log.Printf("%#U starts at byte position %d \n", value, index)
		mainKeyBinarySlice = append(mainKeyBinarySlice, value)
	}
	for i := 7; i >= 0; i-- {
		firstRoundKey += string(mainKeyBinarySlice[firstRoundKeySchemе[i]-1])
		secondRoundKey += string(mainKeyBinarySlice[secondRoundKeySchemе[i]-1])
		thirdRoundKey += string(mainKeyBinarySlice[thirdRoundKeySchemе[i]-1])
	}
	log.Println("Раундовые ключи: ")
	log.Println("first: ", firstRoundKey)
	log.Println("second: ", secondRoundKey)
	log.Println("third: ", thirdRoundKey)
	return mainKey
}
func main() {
	log.Info("Let's start")
	encryptedText := ""
	reader := bufio.NewReader(os.Stdin)
	log.Println("Введите Ваш номер в списке")
	n, _, _ := reader.ReadLine()
	numberOfStudent := string(n)
	number, err := strconv.Atoi(numberOfStudent)
	getX(number)
	if err != nil {
		log.Error(err)
	}
	log.Println("Введите количество букв в имени")
	n, _, _ = reader.ReadLine()
	amountOfCharsInName := string(n)
	amountCharsName, err := strconv.Atoi(amountOfCharsInName)
	if err != nil {
		log.Error(err)
	}

	log.Println("Введите количество букв в фамилии")
	n, _, _ = reader.ReadLine()
	amountOfCharsInSurname := string(n)
	amountCharsSurname, err := strconv.Atoi(amountOfCharsInSurname)
	if err != nil {
		log.Error(err)
	}
	getKey(amountCharsName, amountCharsSurname)

	log.Println("Введите текст, который нужно зашифровать")
	t, _, _ := reader.ReadLine()
	text := string(t)
	for index, value := range text {
		log.Printf("%#U starts at byte position %d \n", value, index)
	}
	log.Println("encryptedText: ", encryptedText)
}
