package keygen

import "strings"

func RotWord(a string) string {

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
