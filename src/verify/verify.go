// shafish-verification
//
// Copyright (c) 2018 ShaFish
// GPL License

package verify

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/tokublock/tokucore/network"
	"github.com/tokublock/tokucore/xcore"
	"github.com/tokublock/tokucore/xcrypto"
)

// Result -- verify result.
type Result struct {
	TxID        string `json:"txid"`
	SeedHash    string `json:"seedhash"`
	PubKey      string `json:"pubkey"`
	BtcAddress  string `json:"btcaddress"`
	BlockHeight int64  `json:"blockheight"`
	URL         string `json:"url"`
}

// Verify --
type Verify struct{}

// NewVerify -- creates new Verify.
func NewVerify() *Verify {
	return &Verify{}
}

// Verify -- verify the seed txt file, if error is nil, result returns.
func (v *Verify) Verify(path string) (*Result, error) {
	// Load seed txt file.
	seed, err := LoadSeed(path)
	if err != nil {
		return nil, err
	}

	// Proof the merkle.
	if !verifyMerkle(seed.DocumentHash, seed.MerkleRoot, seed.MerklePath) {
		return nil, fmt.Errorf("merkle verify returns false")
	}

	// Private key.
	prvhex, err := hex.DecodeString(seed.MerkleRoot)
	if err != nil {
		return nil, err
	}

	// Public key.
	pubkey := xcrypto.PrvKeyFromBytes(prvhex).PubKey()

	// Address.
	address := xcore.NewPayToPubKeyHashAddress(pubkey.Hash160())
	var url string
	var chain *network.Network
	switch seed.ChainNet {
	case "testnet":
		chain = network.TestNet
		url = fmt.Sprintf("https://testnet.blockchain.info/tx/%v", seed.Transaction)
	case "mainnet":
		chain = network.MainNet
		url = fmt.Sprintf("https://blockchain.info/tx/%v", seed.Transaction)
	}
	return &Result{
		TxID:        seed.Transaction,
		SeedHash:    seed.MerkleRoot,
		PubKey:      fmt.Sprintf("%x", pubkey.Serialize()),
		BtcAddress:  address.ToString(chain),
		BlockHeight: seed.BlockHeight,
		URL:         url,
	}, nil
}

// verifyMerkle -- verify the merkle tree with path.
func verifyMerkle(leafHex string, rootHex string, nodes []MerkleNode) bool {
	var err error
	for _, node := range nodes {
		var left, right, parent []byte

		if left, err = hex.DecodeString(node.Left); err != nil {
			return false
		}
		if right, err = hex.DecodeString(node.Right); err != nil {
			return false
		}
		if parent, err = hex.DecodeString(node.Parent); err != nil {
			return false
		}

		merkleHash := merkleNodeHash(left, right)
		if !bytes.Equal(merkleHash, parent) {
			return false
		}
	}
	return true
}

func merkleNodeHash(left []byte, right []byte) []byte {
	const merkleHashSize = 32
	var hash [merkleHashSize * 2]byte

	copy(hash[:merkleHashSize], left[:])
	copy(hash[merkleHashSize:], right[:])
	return xcrypto.DoubleSha256(hash[:])
}
