package keygen

import (
	operation "AES/Operations"
	"strconv"
	"strings"
)

// func rotWord(a string) string {

// 	var firstFourBytes []string

// 	for i := 0; i < 8; i += 2 {
// 		oneByte := string(a[i]) + string(a[i+1])
// 		firstFourBytes = append(firstFourBytes, oneByte)
// 	}

// 	firstByte := firstFourBytes[0]
// 	firstFourBytes = firstFourBytes[1:]

// 	firstFourBytes = append(firstFourBytes, firstByte)

// 	fourByteString := strings.Join(firstFourBytes, "")

// 	return fourByteString
// }

func rotWord(a string) string {
	var lastWord []string

	for i := 24; i < 32; i += 2 {
		oneByte := string(a[i]) + string(a[i+1])
		lastWord = append(lastWord, oneByte)
	}

	byteToRotate := lastWord[0]
	lastWord = lastWord[1:]

	lastWord = append(lastWord, byteToRotate)

	wordString := strings.Join(lastWord, "")

	//return only the last word thats already been rotated
	return wordString
}

func rCon(numberOfRountConstant int) []int {

	numberOfRountConstant -= 1
	var roundConstantArray = []int{0x01}

	for i := range numberOfRountConstant {

		nextRountConstant := roundConstantArray[i] * 0x2

		if roundConstantArray[i] >= 0x80 {
			nextRountConstant ^= 0x11B
		}

		roundConstantArray = append(roundConstantArray, nextRountConstant)
	}

	return roundConstantArray
}

func convToHex(a string) []int {

	var hexArray []int

	// for _, element := range a {
	for i := 0; i < len(a); i += 2 {

		stringVal := string(a[i]) + string(a[i+1])

		hexVal, err := strconv.ParseInt(stringVal, 16, 64)
		// hexVal2, err := strconv.ParseInt(string(a[i+1]), 16, 64)

		if err != nil {
			panic(err)
		}
		hexArray = append(hexArray, int(hexVal))

	}

	return hexArray
}

func GenerateRoundKey(password string) [11]string {

	//hard set the amount of keys
	var resultMatrix [11]string

	//128-bit aes encryption needs 10 round keys
	roundConstant := rCon(10)

	cypherKey := GenerateKey(password)
	resultMatrix[0] = cypherKey

	for i, element := range resultMatrix {

		// predefine a variable for the xor multipliyer & temp result matrix house
		var xorMultiplyer []int
		var resultMatrixHouse []string

		rotWord := rotWord(element)
		subWord := operation.Substitude(rotWord, false)

		//converting the subword to int and inserting into xorMultipliyer
		xorMultiplyer = convToHex(subWord)

		//getting the rCon values
		rCon := roundConstant[i]

		//round constant operation
		xorMultiplyer[0] = rCon ^ xorMultiplyer[0]

		//converting element into int
		hexElement := convToHex(element)

		for j := 0; j < len(hexElement); j++ {
			for z := 0; z < 4; z++ {
				xorMultiplyer[z] = xorMultiplyer[z] ^ hexElement[j+z]
				resultMatrixHouse = append(resultMatrixHouse, string(xorMultiplyer[z]))
			}

		}

		fullResultString := strings.Join(resultMatrixHouse, "")
		resultMatrix[i+1] = fullResultString
	}

	return resultMatrix

	// for i, element := range resultMatrix {

	// 	var hexWordMultipliyer []int

	// 	rotBytes := rotWord(element)
	// 	subBytes := operation.Substitude(rotBytes, false)

	// 	hexWordMultipliyer = convToHex(subBytes)

	// 	rConstant := roundConstant[i]
	// 	hexWordMultipliyer[0] = hexWordMultipliyer[0] ^ rConstant

	// 	var tempArray []int

	// 	hexElement := convToHex(element)
	// 	for j := 0; j < 32; j += 4 {

	// 	}
	// }

}
