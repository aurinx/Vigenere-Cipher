package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
    "regexp"
)

func main() {

	cipherTextFile := os.Args[1];

	fileReadIn, err := ioutil.ReadFile(string(cipherTextFile)) 
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

    cipherTextLength := len(strUpperCaseAlphaNumeric)

    fmt.Println(cipherTextLength)

    max := 0
   
    var indexOfCoincidenceArray [20]float64

    //length of keylength being tested
    for i := 1; i <= 20; i++ {

    	var letterFrequencies = make([]float64, i)

    	// Split converted ciphertext into i pieces, start at j index, use sum to continuously
    	// go through entire text and jump by i each iteration
    	for j := 0; j < i; j++ {


    		var currLetterSequence = ""
    		sum := j


    		for sum < cipherTextLength {
    			currLetterSequence += string(strUpperCaseAlphaNumeric[sum])
    			sum += i
    		}

    		var letterSequenceLen = len(currLetterSequence)

    		//Only look up upper case letters
    		for k := 0; k < 26; k++ {
    			var letter = string(k + 65)
    			var letterCountInLetterSequence = strings.Count(currLetterSequence, letter)
    			var freqOfLetter = float64(letterCountInLetterSequence) / float64(letterSequenceLen)
    			var freqInd = freqOfLetter * ( (float64((letterCountInLetterSequence - 1))) / (float64(letterSequenceLen - 1)) )
    			letterFrequencies[j] += freqInd
    		}
 
 			indexOfCoincidenceArray[i - 1] += letterFrequencies[j]
    	}

    	indexOfCoincidenceArray[i - 1] /= float64(i)


    	if indexOfCoincidenceArray[i - 1] > 0.06 {
    		max = i - 1
    		break
    	} else if indexOfCoincidenceArray[i - 1] > float64(indexOfCoincidenceArray[max]) {
    		max = i - 1
    	}
 	}


 	fmt.Println(max + 1)

}