package gostealth

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

var (
	ErrInvalidAddress  = errors.New("invalid Stealth Meta Address")
	ErrViewTagMismatch = errors.New("view tag mismatch")
	ErrAddressMismatch = errors.New("address mismatch")

	curve = secp256k1.S256()
)

type ViewTag int

func StealthMetaAddress(spending, scanning *ecdsa.PublicKey) string {
	return fmt.Sprintf("st:eth:0x%x%x", crypto.CompressPubkey(spending), crypto.CompressPubkey(scanning))
}

func ParseStealthAddress(address string) (*ecdsa.PublicKey, *ecdsa.PublicKey, error) {
	sma := common.FromHex(strings.TrimPrefix(address, "st:eth:0x"))
	if len(sma) != 33*2 {
		return nil, nil, ErrInvalidAddress
	}
	spending, err := crypto.DecompressPubkey(sma[:33])
	if err != nil {
		return nil, nil, ErrInvalidAddress
	}
	scanning, err := crypto.DecompressPubkey(sma[33:])
	if err != nil {
		return nil, nil, ErrInvalidAddress
	}
	return spending, scanning, nil
}

func GenStealthAddress(address string, key *ecdsa.PrivateKey) (*common.Address, *ecdsa.PublicKey, ViewTag, error) {
	spending, scanning, err := ParseStealthAddress(address)
	if err != nil {
		return nil, nil, 0, err
	}

	dhSecret := computeDHSecret(scanning, key.D)
	dhSecretHash := hashPK(dhSecret)
	viewTag := getViewTag(dhSecretHash)
	dhSecretPoint := secretToPoint(dhSecretHash)
	stealthAddressPK := stealthPKFromPoint(dhSecretPoint, spending)
	stealthAddr := crypto.PubkeyToAddress(*stealthAddressPK)
	return &stealthAddr, &key.PublicKey, viewTag, nil
}

// ParseEventView returns the shared secret if the stealth address
// belongs to us. If not an error is returned. This method does not
// need the spending private key.
func ParseEventView(pubkey *ecdsa.PublicKey, stA common.Address, viewTag ViewTag, scanning *ecdsa.PrivateKey, spending *ecdsa.PublicKey) (*big.Int, error) {
	dhSecret := computeDHSecret(pubkey, scanning.D)
	dhSecretHash := hashPK(dhSecret)
	if getViewTag(dhSecretHash) != viewTag {
		return nil, ErrViewTagMismatch
	}
	fmt.Printf("%x\n", dhSecretHash)
	dhSecretPoint := secretToPoint(dhSecretHash)
	stealthAddrPK := stealthPKFromPoint(dhSecretPoint, spending)
	stealthAddr := crypto.PubkeyToAddress(*stealthAddrPK)
	if !strings.EqualFold(stA.Hex(), stealthAddr.Hex()) {
		return nil, ErrAddressMismatch
	}
	return dhSecretHash, nil
}

// ParseEvent returns the private key to spend the funds on the address
// or an error if the address does not belong to our keys.
func ParseEvent(pubkey *ecdsa.PublicKey, stA common.Address, viewTag ViewTag, scanning *ecdsa.PrivateKey, spending *ecdsa.PrivateKey) (*ecdsa.PrivateKey, error) {
	secret, err := ParseEventView(pubkey, stA, viewTag, scanning, &spending.PublicKey)
	if err != nil {
		return nil, err
	}
	_, key, err := stealthAddrPrivKey(secret, spending)
	return key, err
}

func computeDHSecret(pk *ecdsa.PublicKey, secret *big.Int) *ecdsa.PublicKey {
	pubX, pubY := curve.ScalarMult(pk.X, pk.Y, secret.Bytes())
	return &ecdsa.PublicKey{Curve: curve, X: pubX, Y: pubY}
}

func hashPK(pk *ecdsa.PublicKey) *big.Int {
	xBytes := make([]byte, 32)
	yBytes := make([]byte, 32)
	pk.X.FillBytes(xBytes)
	pk.Y.FillBytes(yBytes)
	hash := new(big.Int).SetBytes(crypto.Keccak256(xBytes, yBytes))
	return hash.Mod(hash, secp256k1.S256().N)
}

func getViewTag(hash *big.Int) ViewTag {
	return ViewTag(hash.Bytes()[0])
}

func secretToPoint(secret *big.Int) *ecdsa.PublicKey {
	x, y := curve.ScalarBaseMult(secret.Bytes())
	return &ecdsa.PublicKey{Curve: curve, X: x, Y: y}
}

func stealthPKFromPoint(pk1, pk2 *ecdsa.PublicKey) *ecdsa.PublicKey {
	x, y := secp256k1.S256().Add(pk1.X, pk1.Y, pk2.X, pk2.Y)
	return &ecdsa.PublicKey{Curve: curve, X: x, Y: y}
}

func stealthAddrPrivKey(secret *big.Int, key *ecdsa.PrivateKey) (*big.Int, *ecdsa.PrivateKey, error) {
	sk := new(big.Int).Mod(new(big.Int).Add(secret, key.D), curve.N)
	skEc, err := crypto.ToECDSA(sk.Bytes())
	return sk, skEc, err
}
