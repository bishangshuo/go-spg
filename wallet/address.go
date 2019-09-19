package wallet

import (
	"errors"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"go-spg/db/walletdb"
)

type Wallet struct {
	//PrivateKey ecdsa.PrivateKey
	//PublicKey  []byte
}

func NewWallet() *Wallet {

	wallet := Wallet{}

	return &wallet
}

func (w Wallet) GetNewAddress() string {

	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return ""
	}

	privKeyWif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, false)
	if err != nil {
		return ""
	}
	pubKeySerial := privKey.PubKey().SerializeUncompressed()

	pubKeyAddress, err := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.MainNetParams)
	if err != nil {
		return ""
	}

	privKeyString := privKeyWif.String()
	walletdb.GetInstance().SaveKeyWithAddress(pubKeyAddress.EncodeAddress(), privKeyString)

	return pubKeyAddress.EncodeAddress()
}

func (w Wallet) DumpPrivKey(address string) (string, error){
	r, key := walletdb.GetInstance().GetKeyWithAddress(address)
	if r {
		return key, nil
	}
	return "", errors.New("address is not in this wallet")
}

