// shafish-verification
//
// Copyright (c) 2018 ShaFish
// GPL License

package verify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeed(t *testing.T) {
	// Valid.
	{
		path := "testdata/seed/valid.seed.txt"
		_, err := LoadSeed(path)
		assert.Nil(t, err)
	}

	// Path error.
	{
		path := "testdata/seed/invalid.seed"
		_, err := LoadSeed(path)
		assert.NotNil(t, err)
	}

	// Invalid.
	{
		path := "testdata/seed/invalid.seed.txt"
		_, err := LoadSeed(path)
		assert.NotNil(t, err)
	}
}
