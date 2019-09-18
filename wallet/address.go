package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.org/btcsuite/btcutil/base58"
	"github.org/pkg/errors"
	"go-spg/db/walletdb"
	"golang.org/x/crypto/ripemd160"
	"log"
)

const walletVersion = byte(0x00)
const addressChecksumLen = 4

type Wallet struct {
	//PrivateKey ecdsa.PrivateKey
	//PublicKey  []byte
}

func NewWallet() *Wallet {

	wallet := Wallet{}

	return &wallet
}

func (w Wallet) GetNewAddress() string {
	privateKey, publicKey := newKeyPair()

	pubKeyHash := HashPubKey(publicKey)

	walletVersionedPayload := append([]byte{walletVersion}, pubKeyHash...)
	checksum := checksum(walletVersionedPayload)

	fullPayload := append(walletVersionedPayload, checksum...)
	address := Encode(fullPayload)

	privateKeyStr := privateKey.D.String()
	walletdb.GetInstance().SaveKeyWithAddress(address, privateKeyStr)

	return address
}

func (w Wallet) DumpPrivKey(address string) (string, error){
	r, key := walletdb.GetInstance().GetKeyWithAddress(address)
	if r {
		base58Key := base58.Encode([]byte(key))
		return base58Key, nil
	}
	return "", errors.New("address is not in this wallet")
}

func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}

func ValidateAddress(address string) bool {
	pubKeyHash := base58.Decode(address)
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Compare(actualChecksum, targetChecksum) == 0
}

func TestAddress(address string) bool {
	return true
}