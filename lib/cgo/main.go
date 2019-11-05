/*
libskycoin shim
*/
package main

import (
	sky "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"

	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
)

// #cgo CFLAGS: -I../../include
import "C"

func main() {
	util.RegisterAltcoin(sky.NewSkyFiberPlugin(sky.SkycoinMainNetParams))
}
