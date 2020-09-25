package eos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testAddr = "EOS6sbqNFJrmpXxvb2NnSqX19dChgBuDiZ5349bWWmR3jKJACW5ED"
	testWif  = "5KMvN716ckvewapUU1seBYUMbBEcs1K4J22exVbc6nkmnxMCzy4"
)

func TestWIFToAddress(t *testing.T) {
	addr, err := WIFToAddress(testWif)
	assert.Nil(t, err)
	assert.Equal(t, testAddr, addr)
}

func TestCheckWIF(t *testing.T) {
	assert.True(t, CheckWIF(testWif))
}

func TestGeneratePair(t *testing.T) {
	wif, addr, err := GeneratePair(nil)
	assert.Nil(t, err)

	assert.True(t, CheckWIF(wif))

	addr1, err := WIFToAddress(wif)
	assert.Nil(t, err)
	assert.Equal(t, addr, addr1)
}
