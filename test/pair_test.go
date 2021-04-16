package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"thinhlu123/crawl-uniswap/src/client"
)

func TestGetPair(t *testing.T) {
	client.InitUniSwapClient()

	a, err := client.GetPairById("0x94b0a3d511b6ecdb17ebf877278ab030acb0a878")

	assert.NoError(t, err)
	fmt.Println(a)
}
