package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"thinhlu123/crawl-uniswap/src/client"
)

func TestGetPair(t *testing.T) {
	client.InitUniSwapClient()

	a, err := client.GetPair()

	assert.NoError(t, err)
	fmt.Println(a)
}
