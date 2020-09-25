package eosaddr

import (
	"bytes"
	"crypto/ecdsa"
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"golang.org/x/crypto/ripemd160"
)

func GeneratePair(r io.Reader) (string, string, error) {
	if r == nil {
		r = crand.Reader
	}
	priv, err := ecdsa.GenerateKey(secp256k1.S256(), r)
	if err != nil {
		return "", "", err
	}

	finPri, err := privateKey2WIF(fmt.Sprintf("%x", priv.D))
	if err != nil {
		return "", "", err
	}

	pub := privateKey2PublicKey(priv.X.Bytes(), priv.Y.Bytes())
	return finPri, publicKey2Address(pub), nil
}

func privateKey2WIF(pri string) (string, error) {
	by, err := hex.DecodeString("80" + pri)
	if err != nil {
		return "", err
	}
	h0 := sha256.Sum256(by)
	h1 := sha256.Sum256(h0[:])
	return base58.Encode(append(by, h1[:4]...)), nil
}

func WIFToAddress(wif string) (string, error) {
	priv, err := crypto.ToECDSA(getPrivateKeyFromWIF(wif))
	if err != nil {
		return "", err
	}
	pub := privateKey2PublicKey(priv.X.Bytes(), priv.Y.Bytes())
	return publicKey2Address(pub), nil
}

func getPrivateKeyFromWIF(wif string) []byte {
	if CheckWIF(wif) {
		base58DecodeData := base58.Decode(wif)
		return base58DecodeData[1:33]
	}
	return []byte{}

}

func privateKey2PublicKey(x, y []byte) []byte {
	if y[len(y)-1]%2 == 0 {
		return append([]byte{byte(02)}, x...)
	}
	return append([]byte{byte(03)}, x...)
}

// check wif key valid.
func CheckWIF(wif string) bool {
	base58DecodeData := base58.Decode(wif)
	length := len(base58DecodeData)

	if length < 37 {
		return false
	}

	private := base58DecodeData[:(length - 4)]
	h0 := sha256.Sum256(private)
	h1 := sha256.Sum256(h0[:])
	checksum := h1[:4]
	checksum1 := base58DecodeData[(length - 4):]
	if bytes.Compare(checksum, checksum1) == 0 {
		return true
	}
	return false
}

func publicKey2Address(data []byte) string {
	h160 := ripemd(data)
	checkSum := h160[:4]
	return "EOS" + base58.Encode(append(data, checkSum...))

}

func ripemd(data []byte) []byte {
	hasher := ripemd160.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}
