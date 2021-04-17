package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"thinhlu123/crawl-uniswap/src/client"
)

func TestGetPair(t *testing.T) {
	client.InitUniSwapClient()

	var pairArr = []string{"0x94b0a3d511b6ecdb17ebf877278ab030acb0a878", "0x9928e4046d7c6513326ccea028cd3e7a91c7590a", "0xcd7989894bc033581532d2cd88da5db0a4b12859"}
	for _, id := range pairArr {
		_, err := client.GetPairById(id)
		assert.NoError(t, err)
	}
}

func TestGetTransaction(t *testing.T) {
	client.InitUniSwapClient()

	var pairArr = []string{"0x94b0a3d511b6ecdb17ebf877278ab030acb0a878", "0x9928e4046d7c6513326ccea028cd3e7a91c7590a", "0xcd7989894bc033581532d2cd88da5db0a4b12859"}
	for _, id := range pairArr {
		_, err := client.GetTokenTransactionById(id)
		assert.NoError(t, err)
	}

}

func TestGetBundle(t *testing.T) {
	client.InitUniSwapClient()

	a, err := client.GetPriceETH()

	assert.NoError(t, err)
	fmt.Println(a)
}
