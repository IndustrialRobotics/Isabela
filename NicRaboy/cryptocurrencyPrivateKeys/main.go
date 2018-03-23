package main

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/btcec"
	"fmt"
	"errors"
)

//Generate Cryptocurrency Private Keys And Public Addresses With Golang - YouTube -
/*
output:
 L2qXeJwTkxB6Tjf9XxLR7J6q5KwPJRZEB9oYNwtpy7coSQDS4JU6
 LSueWvwYFz6a78SksKsiDZNBbsCgafzTXD
------------------------------
T7tE3xfFtfLvx2QYRpaRBRPQP12MJykMqiHm3eawJwsjtVMa5sbb
 1HoXSKDrRDGthaGpQP5eKyWcWubdisdsm9

The WIF key is not valid for the bitcoin network*/

type Network struct {
	name        string
	symbol      string
	xpubkey     byte
	xprivatekey byte
}

var network = map[string]Network{
	"btc": {name: "bitcoin", symbol: "btc", xpubkey: 0x00, xprivatekey: 0x80},
	"ltc": {name: "litecoin", symbol: "ltc", xpubkey: 0x30, xprivatekey: 0xb0},
}

func (network Network) getNetworkParams() *chaincfg.Params {
	networkParams := &chaincfg.MainNetParams
	networkParams.PubKeyHashAddrID = network.xpubkey
	networkParams.PrivateKeyID = network.xprivatekey
	return networkParams
}

//WIF = wallet information format
func (network Network) createWIF() (*btcutil.WIF, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	return btcutil.NewWIF(secret, network.getNetworkParams(), true)
}

func (network Network) getAddress(wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.getNetworkParams())
}

func (network Network) importWIF(wifStr string) (*btcutil.WIF, error) {
	wif, err := btcutil.DecodeWIF(wifStr)
	if err != nil {
		return nil, err
	}
	if !wif.IsForNet(network.getNetworkParams()) {
		return nil, errors.New("\nThis WIF key is not valid on the " + network.name + " network")
	}
	return wif, nil
}

func main() {

	wif, _ := network["btc"].createWIF()
	address, _ := network["ltc"].getAddress(wif)
	fmt.Println("\n", wif.String(), "\n", address.EncodeAddress())

	fmt.Println("------------------------------")

	wif, _ = network["ltc"].createWIF()
	address, _ = network["btc"].getAddress(wif)
	fmt.Println(wif.String(), "\n", address.EncodeAddress())

	wif1, err := network["btc"].importWIF(wif.String())
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(wif1)
	}

}
