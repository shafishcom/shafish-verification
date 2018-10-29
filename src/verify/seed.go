// shafish-verification
//
// Copyright (c) 2018 ShaFish
// GPL License

package verify

import (
	"encoding/json"
	"io/ioutil"
)

// MerkleNode --
type MerkleNode struct {
	Left   string `json:"left"`
	Right  string `json:"right"`
	Parent string `json:"parent"`
}

// Seed --
type Seed struct {
	Version       string       `json:"version"`
	DocumentHash  string       `json:"document_hash"`
	MerkleRoot    string       `json:"merkle_root"`
	MerklePath    []MerkleNode `json:"merkle_path"`
	ChainNet      string       `json:"chainnet"`
	BlockHeight   int64        `json:"block_height"`
	Transaction   string       `json:"transaction"`
	Certification string       `json:"certification"`
	Help          string       `json:"help"`
}

// LoadSeed -- load the seed txt file to json.
func LoadSeed(path string) (*Seed, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	seed := &Seed{}
	if err := json.Unmarshal([]byte(data), seed); err != nil {
		return nil, err
	}
	return seed, nil
}
