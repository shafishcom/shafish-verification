# shafish-verification
Check &amp; verify the ShaFish timestamps


[![Build Status](https://travis-ci.com/shafishcom/shafish-verification.png)](https://travis-ci.com/shafishcom/shafish-verification) [![Go Report Card](https://goreportcard.com/badge/github.com/shafishcom/shafish-verification)](https://goreportcard.com/report/github.com/shafishcom/shafish-verification) [![codecov.io](https://codecov.io/gh/shafishcom/shafish-verification/graphs/badge.svg)](https://codecov.io/gh/shafishcom/shafish-verification/branch/master)

## Requirements

[Golang](http://golang.org)

## Build

```
$ git clone https://github.com/shafishcom/shafish-verification
$ cd shafish-verification
$ make build
```

## Usage

```
$ ./bin/shafish-verification
  Usage: ./bin/shafish-verification [-s] <shafish-seed-txt-file>
```

## Examples

```
$./bin/shafish-verification -s src/verify/testdata/verify/valid.testnet.seed.txt
Seed:        src/verify/testdata/verify/valid.testnet.seed.txt
Result:      ok
Public Key:  026b8d79bc7a175a2e7ddae811ed04b61c2089c717d385cd7050d1bc3c05ee3cde
Address:     mmMgqNamJiYYqpmS4MHYvmpyuqmUq5gpTv
Transaction: ee7fffa05091a9036da678a371769059f610a5a6c26c32841aba485fbdbf5702
Explorer:    https://testnet.blockchain.info/tx/ee7fffa05091a9036da678a371769059f610a5a6c26c32841aba485fbdbf5702
```

## Can I trust this code?
> Don't trust. Verify.

