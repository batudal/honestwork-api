package crypto

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(salt string, address string, signature string) bool {
	sig := hexutil.MustDecode(signature)
	msg := accounts.TextHash([]byte(salt))
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	return address == recoveredAddr.Hex()
}

func GenerateID(salt string, address string) string {
	elements := [][]byte{[]byte(salt), []byte(address)}
	data := bytes.Join(elements, []byte(""))
	hash := crypto.Keccak256Hash(data)
	return hash.Hex()
}
