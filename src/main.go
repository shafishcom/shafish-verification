// shafish-verification
//
// Copyright (c) 2018 ShaFish
// GPL License

package main

import (
	"flag"
	"fmt"
	"os"

	"verify"
)

var (
	flagSeed string
)

func init() {
	flag.StringVar(&flagSeed, "s", "", "shafish seed txt file")
}

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " [-s] <shafish-seed-txt-file>")
}

func main() {
	flag.Parse()
	if flagSeed == "" {
		usage()
		os.Exit(0)
	}
	fmt.Printf("Seed:\t%v\n", flagSeed)

	verify := verify.NewVerify()
	result, err := verify.Verify(flagSeed)
	if err != nil {
		fmt.Printf("Result:\tfailed\n")
		fmt.Printf("Error:\t%v\n", err)
	} else {
		fmt.Printf("Result:\tok\n")
		fmt.Printf("Public Key:\t%v\n", result.PubKey)
		fmt.Printf("Address:\t%v\n", result.BtcAddress)
		fmt.Printf("Transaction:\t%v\n", result.TxID)
		fmt.Printf("Explorer:\t%v\n", result.URL)
	}
}
