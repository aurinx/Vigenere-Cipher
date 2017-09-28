package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
    "regexp"
    "strconv"
    "math"
)

func main() {

	cipherTextFile := os.Args[1]
	keyLengthAsString := os.Args[2]
	keyLength, err := strconv.Atoi(keyLengthAsString)
	if err != nil {
		fmt.Print(err)
	}

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
    letterFreqInEnglishLang := [26]float64 {.08167, .01492, .02792, .04253, .12702, .0228, .02015, .06094, .06966, .0153, .0772, .04025, .02406, .06749, .07507, .01929, .0095, .05987, .06327, .09056, .02758, .00978, .02360, .00150, .01974, .0074}

    proposedKey := ""
    //splits cipher into 'keyLength' snippets that start from index i
    for i := 0; i < keyLength; i++ {
    	var currLetterSequence = ""
    	var sum = i;
    	for sum < cipherTextLength{
    		currLetterSequence += string(strUpperCaseAlphaNumeric[sum])
    		sum += keyLength;
    	}
    	//fmt.Println(currLetterSequence)
    	currLetterSequenceLen := len (currLetterSequence)

    	
     	var chiSquareCollection [26]float64

    	for j := 0; j < 26; j++ {

    		
    		shiftedLetterSequence := ""
    		for k := 0; k < currLetterSequenceLen; k++ {
    			currLetter := currLetterSequence[k]

    			currLetter = (currLetter - (byte)(j) - 65) % 26
    			currLetter += 65

    			
    			
    			shiftedLetterSequence += string(currLetter)
    			
    		}
    		var totalChiAmount float64

    		//fmt.Println(shiftedLetterSequence)
    		
    		for l:= 0; l < 26; l++ {
    			letter := string(l + 65)

    			var observed = strings.Count(shiftedLetterSequence, letter)
    			
    			//fmt.Println("obv: ", string(l + 65), observed)
    			var expected = float64(letterFreqInEnglishLang[l]) * float64(currLetterSequenceLen)

    			var differenceSquared = math.Pow(float64(observed) - expected, 2)

    			//fmt.Println("both: ", differenceSquared)

    			var currentChi = float64(differenceSquared) / expected
    			//var totalChiAmount float64 
    			totalChiAmount += currentChi
    		}
    		chiSquareCollection[j] += totalChiAmount
    		//if chiSquareCollection[i]
    		//fmt.Println(string(j+ 65), ": ", totalChiAmount)
    	}
    	
    	//min = chiSquareCollection[0]
    	var min = 0

    	for m := 0; m < 26; m++ {
    		if (chiSquareCollection[m] < chiSquareCollection[min]) {
    			//min = chiSquareCollection[m]
    			min = m
    		}
    	}

    	proposedKey += string(min + 65)

    }
    fmt.Println("Proposed key: ", proposedKey)
}




