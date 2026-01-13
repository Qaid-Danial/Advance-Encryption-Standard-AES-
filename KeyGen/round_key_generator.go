package keygen

import (
	operation "AES/Operations"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func rotWord(a string) string {
	// Takes a 4-byte hex string (word) and performs a cyclic permutation.
	var bytes []string
	for i := 0; i < len(a); i += 2 {
		bytes = append(bytes, a[i:i+2])
	}

	firstByte := bytes[0]
	bytes = bytes[1:]
	bytes = append(bytes, firstByte)

	return strings.Join(bytes, "")
}

func rCon(num int) []int {
	// Generates the round constants (Rcon) for the key schedule.
	if num <= 0 {
		return []int{}
	}
	roundConstantArray := make([]int, num)
	roundConstantArray[0] = 0x01

	for i := 1; i < num; i++ {
		nextRountConstant := roundConstantArray[i-1] * 2
		if roundConstantArray[i-1] >= 0x80 {
			nextRountConstant ^= 0x11B
		}
		roundConstantArray[i] = nextRountConstant
	}

	return roundConstantArray
}

func convToHex(a string) []int {
	var hexArray []int
	for i := 0; i < len(a); i += 2 {
		stringVal := a[i : i+2]
		hexVal, err := strconv.ParseInt(stringVal, 16, 64)
		if err != nil {
			panic(err)
		}
		hexArray = append(hexArray, int(hexVal))
	}
	return hexArray
}

func intToHexStr(n int) string {
	return fmt.Sprintf("%02x", n)
}

func xorWords(word1, word2 string) string {
	// XORs two 4-byte hex strings (words) and returns the resulting hex string.
	b1 := convToHex(word1)
	b2 := convToHex(word2)
	result := ""
	for i := 0; i < 4; i++ {
		result += intToHexStr(b1[i] ^ b2[i])
	}
	return result
}

func GenerateRoundKey(password string) ([11]string, [][]byte) {
	var strRoundKeys [11]string
	var intRoundKeys [][]byte

	constants := rCon(10)

	// The first round key is the original cipher key.
	strRoundKeys[0] = GenerateKey(password)

	for i := 1; i <= 10; i++ {
		prevKey := strRoundKeys[i-1]
		prevWords := make([]string, 4)
		for j := 0; j < 4; j++ {
			prevWords[j] = prevKey[j*8 : (j+1)*8]
		}

		// g(w) = SubWord(RotWord(w)) ^ Rcon
		g_w3 := rotWord(prevWords[3])
		g_w3, _ = operation.Substitude(g_w3, false)
		g_w3_bytes := convToHex(g_w3)
		g_w3_bytes[0] ^= constants[i-1] // XOR with round constant
		g_w3 = intToHexStr(g_w3_bytes[0]) + intToHexStr(g_w3_bytes[1]) + intToHexStr(g_w3_bytes[2]) + intToHexStr(g_w3_bytes[3])

		// Generate the new words for the new round key.
		newW0 := xorWords(prevWords[0], g_w3)
		newW1 := xorWords(prevWords[1], newW0)
		newW2 := xorWords(prevWords[2], newW1)
		newW3 := xorWords(prevWords[3], newW2)

		currentRoundKey := newW0 + newW1 + newW2 + newW3
		byteForm, _ := hex.DecodeString(currentRoundKey)

		strRoundKeys[i] = currentRoundKey
		intRoundKeys = append(intRoundKeys, byteForm)
	}	

	return strRoundKeys, intRoundKeys
}
