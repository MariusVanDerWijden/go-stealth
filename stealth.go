package gostealth

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func NewStealthAddress(secret [32]byte) (*big.Int, *big.Int, common.Address) {
	curve := secp256k1.S256()
	//  s*G = S
	pubX, pubY := curve.ScalarMult(GX, GY, secret[:])
	//  s*P = q
	qX, qY := curve.ScalarMult(pubX, pubY, secret[:])
	hash := crypto.Keccak256Hash(qX.Bytes(), qX.Bytes()) // TODO fix
	// hash value to pubKey
	qX, qY = curve.ScalarMult(pubX, pubY, hash[:])
	// derive pub key
	qX, qY = curve.Add(pubX, pubY, qX, qY)
	pubkey := ecdsa.PublicKey{Curve: curve, X: qX, Y: qY}
	return pubX, pubY, crypto.PubkeyToAddress(pubkey)
}
