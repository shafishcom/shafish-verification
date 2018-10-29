// shafish-verification
//
// Copyright (c) 2018 ShaFish
// GPL License

package verify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	// TestNet Valid.
	{
		path := "testdata/verify/valid.testnet.seed.txt"
		verify := NewVerify()
		res, err := verify.Verify(path)
		assert.Nil(t, err)
		t.Logf("%+v", res)
	}

	// MainNet Valid.
	{
		path := "testdata/verify/valid.mainnet.seed.txt"
		verify := NewVerify()
		res, err := verify.Verify(path)
		assert.Nil(t, err)
		t.Logf("%+v", res)
	}

	// Merkle Invalid.
	{
		path := "testdata/verify/invalid.seed.txt"
		verify := NewVerify()
		_, err := verify.Verify(path)
		assert.NotNil(t, err)
	}

	// Seed path broken.
	{
		path := "testdata/verify/invalid.roothex"
		verify := NewVerify()
		_, err := verify.Verify(path)
		assert.NotNil(t, err)
	}

	// Root hex broken.
	{
		path := "testdata/verify/invalid.roothex.seed.txt"
		verify := NewVerify()
		_, err := verify.Verify(path)
		assert.NotNil(t, err)
	}

	// Left hex broken.
	{
		path := "testdata/verify/invalid.lefthex.seed.txt"
		verify := NewVerify()
		_, err := verify.Verify(path)
		assert.NotNil(t, err)
	}

	// Right hex broken.
	{
		path := "testdata/verify/invalid.righthex.seed.txt"
		verify := NewVerify()
		_, err := verify.Verify(path)
		assert.NotNil(t, err)
	}

	// Parent hex broken.
	{
		path := "testdata/verify/invalid.parenthex.seed.txt"
		verify := NewVerify()
		_, err := verify.Verify(path)
		assert.NotNil(t, err)
	}
}
