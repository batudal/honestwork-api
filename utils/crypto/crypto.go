package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateSignature(privkey *ecdsa.PrivateKey, data []byte) []byte {
	hash := crypto.Keccak256Hash(data)
	signature, err := crypto.Sign(hash.Bytes(), privkey)
	if err != nil {
		log.Fatal(err)
	}
	return signature
}

func generatePrivkey() (privkey []byte, pubkey []byte, address string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return privateKeyBytes, publicKeyBytes, address
}

func VerifySignatureTest() bool {
	privateKeyBytes, publicKeyBytes, address := generatePrivkey()

	data := []byte("hello" + address)
	hash := crypto.Keccak256Hash(data)

	privateKeyECDSA, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	signature, err := crypto.Sign(hash.Bytes(), privateKeyECDSA)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)

	return matches
}

func VerifySignature(message string, pubkey []byte, signature []byte) bool {
	hash := GenerateID(message, pubkey, signature)
	sigPubKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	matches := bytes.Equal(sigPubKey, pubkey)
	return matches
}

func GenerateID(message string, pubkey []byte, signature []byte) common.Hash {
	elements := [][]byte{[]byte(message), []byte(pubkey)}
	data := bytes.Join(elements, []byte(""))
	hash := crypto.Keccak256Hash(data)
	return hash
}
