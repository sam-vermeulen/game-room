package utils

import "math/rand"

func GenerateCode() string {
	const chars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	const codeLength = 6

	code := make([]byte, codeLength)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}

	return string(code)
}
