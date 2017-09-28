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
	decipherKey := os.Args[1]
	fileName := os.Args[2]
	var recoveredPlainText = flag.String("recoveredPlainText", "", "plaintext recovered from cipheredText")

	decipherKey = strings.ToUpper(decipherKey)

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

    var decipherKeyArray = []byte(decipherKey)
	var cipherTextArray = []byte(strUpperCaseAlphaNumeric)
	var recoveredPlainTextArray = []byte(*recoveredPlainText)

	var keyBitTracker int = 0

	for i := 0; i < len(strUpperCaseAlphaNumeric); i++ {
		messageBit := cipherTextArray[i]
		keyBit := decipherKeyArray[keyBitTracker]

		var recoveredByte byte = 0

		//Check for negative difference, if so then reverse operation
		if (messageBit < keyBit) {
			recoveredByte = 26 - (keyBit - messageBit)
		} else {
			recoveredByte = messageBit - keyBit
		}
		// add dencrypted bit to the end of the cipherTextArray
		recoveredPlainTextArray = append(recoveredPlainTextArray, (recoveredByte % 26 + 65))
		// loop back around the keyBitTracker if it reaches end
		keyBitTracker++
		if keyBitTracker >= len(decipherKeyArray) {
			keyBitTracker = 0
		}
	}

	fmt.Println("stripped and uppercased ciphertext: ", strUpperCaseAlphaNumeric) 
	fmt.Println("outputted recoveredPlainText: ", string(recoveredPlainTextArray))


}