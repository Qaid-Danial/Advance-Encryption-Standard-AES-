package keygen

import "strings"

func rotWord(a string) string {

	var firstFourBytes []string

	for i := 0; i < 8; i += 2 {
		oneByte := string(a[i]) + string(a[i+1])
		firstFourBytes = append(firstFourBytes, oneByte)
	}

	firstByte := firstFourBytes[0]
	firstFourBytes = firstFourBytes[1:]

	firstFourBytes = append(firstFourBytes, firstByte)

	fourByteString := strings.Join(firstFourBytes, "")

	return fourByteString
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

func GenerateRoundKey() {

}
