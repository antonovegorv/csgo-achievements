package util

import (
	"math/big"
	"regexp"
	"strings"

	"github.com/antonovegorv/csgo-achievements/internal/pkg/models"
)

const (
	dictionary string = "ABCDEFGHJKLMNOPQRSTUVWXYZabcdefhijkmnopqrstuvwxyz23456789"
	bitmask64  uint64 = 18446744073709551615
)

// DecodeShareCode ...
func DecodeShareCode(shareCode string) *models.Match {
	re := regexp.MustCompile(`^CSGO|\-`)
	s := re.ReplaceAllString(shareCode, "")
	s = ReverseString(s)

	bigNumber := big.NewInt(0)

	for _, c := range s {
		bigNumber = bigNumber.Mul(bigNumber, big.NewInt(int64(len(dictionary))))
		bigNumber = bigNumber.Add(bigNumber, big.NewInt(int64(strings.Index(dictionary, string(c)))))
	}

	a := SwapEndianness(bigNumber)

	matchID := big.NewInt(0)
	outcomeID := big.NewInt(0)
	tokenID := big.NewInt(0)

	matchID = matchID.And(a, big.NewInt(0).SetUint64(bitmask64))
	outcomeID = outcomeID.Rsh(a, 64)
	outcomeID = outcomeID.And(outcomeID, big.NewInt(0).SetUint64(bitmask64))
	tokenID = tokenID.Rsh(a, 128)
	tokenID = tokenID.And(tokenID, big.NewInt(0xFFFF))

	return &models.Match{
		ShareCode: shareCode,
		MatchID:   matchID.Uint64(),
		OutcomeID: outcomeID.Uint64(),
		TokenID:   tokenID.Uint64(),
	}
}

// SwapEndianness ...
func SwapEndianness(n *big.Int) *big.Int {
	result := big.NewInt(0)

	left := big.NewInt(0)
	rightTemp := big.NewInt(0)
	rightResult := big.NewInt(0)

	for i := 0; i < 144; i += 8 {
		left = left.Lsh(result, 8)
		rightTemp = rightTemp.Rsh(n, uint(i))
		rightResult = rightResult.And(rightTemp, big.NewInt(0xFF))
		result = left.Add(left, rightResult)
	}

	return result
}

// ReverseString ...
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
