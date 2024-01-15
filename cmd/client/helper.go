package main

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/manifoldco/promptui"
)

func getScanningSK() (*ecdsa.PrivateKey, error) {
	prompt := promptui.Prompt{Label: "Please provide the SCANNING secret key"}
	str, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	scanningSK, err := crypto.HexToECDSA(str)
	if err != nil {
		return nil, err
	}
	return scanningSK, nil
}

func getSpendingSK() (*ecdsa.PrivateKey, error) {
	prompt := promptui.Prompt{Label: "Please provide the SPENDING secret key"}
	str, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	spendingSK, err := crypto.HexToECDSA(str)
	if err != nil {
		return nil, err
	}
	return spendingSK, nil
}

func getSpendingPK() (*ecdsa.PublicKey, error) {
	prompt := promptui.Prompt{Label: "Please provide the SPENDING public key"}
	str, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	spendingPK, err := crypto.UnmarshalPubkey(common.Hex2Bytes(str))
	if err != nil {
		return nil, err
	}
	return spendingPK, nil
}
