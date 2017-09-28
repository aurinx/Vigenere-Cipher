package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
    "regexp"
    "flag"
)

func main() {

	cipherKey := os.Args[1]
	fileName := os.Args[2]
	var cipherText = flag.String("cipherText", "", "cipherText output of plaintext input")

	cipherKey = strings.ToUpper(cipherKey)

	fileReadIn, err := ioutil.ReadFile(string(fileName)) 
    if err != nil {
        fmt.Print(err)
    }

    str := string(fileReadIn) 
    strUpperCase := strings.ToUpper(str)

    regex, err := regexp.Compile("[^a-zA-Z]+")
    if err != nil {
        fmt.Print(err)
    }

    strUpperCaseAlphaNumeric := regex.ReplaceAllString(strUpperCase, "")

    flag.Parse()

	var cipherKeyArray = []byte(cipherKey)
	var plainTextArray = []byte(strUpperCaseAlphaNumeric)
	var cipherTextArray = []byte(*cipherText)
	//To check when we reach end of the key, loop back around
	var keyBitTracker int = 0
    
    for i := 0; i < len(strUpperCaseAlphaNumeric); i++ {
		messageBit := plainTextArray[i]
		keyBit := cipherKeyArray[keyBitTracker]
		// add encrypted bit to the end of the cipherTextArray, use 65 = 'A' in ASCII 
		cipherTextArray = append(cipherTextArray, ((messageBit + keyBit) - 65 << 1) % 26 + 65)
		// loop back around the keyBitTracker if it reaches end
		keyBitTracker++
		if keyBitTracker >= len(cipherKeyArray) {
			keyBitTracker = 0
		}
	}

	fmt.Println("stripped and uppercased plaintext: ", strUpperCaseAlphaNumeric) 
	fmt.Println("outputted cipherText: ", string(cipherTextArray))

}