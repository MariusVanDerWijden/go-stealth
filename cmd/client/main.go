package main

import (
	"crypto/ecdsa"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"

	"github.com/MariusVanDerWijden/go-stealth/bindings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	if err := mainLoop(); err != nil {
		panic(err)
	}
}

func mainLoop() error {
	color.Green("Welcome to go-stealth")

	prompt := promptui.Prompt{Label: "Please provide your RPC provider"}
	str, err := prompt.Run()
	if err != nil {
		return err
	}
	client, err := ethclient.Dial(str)
	if err != nil {
		return err
	}

	prompt = promptui.Prompt{Label: "Please provide the contract address"}
	str, err = prompt.Run()
	if err != nil {
		return err
	}
	cAddr := common.HexToAddress(str)
	contract, err := bindings.NewERC5564Announcer(cAddr, client)
	if err != nil {
		return err
	}

	promptInit := promptui.Select{
		Label: "Choose one of the following",
		Items: []string{"View events", "Spend events"},
	}
	for {
		idx, _, err := promptInit.Run()
		if err != nil {
			color.Red("Goodbye: %v", err)
		}
		switch idx {
		case 0:
			execView(client, contract)
		case 1:
			execSpend(client, contract)
		}
	}
}

func execView(client *ethclient.Client, contract *bindings.ERC5564Announcer) error {
	scanningSK, err := getScanningSK()
	if err != nil {
		return err
	}
	spendingPK, err := getSpendingPK()
	if err != nil {
		return err
	}
	return ViewScan(client, contract, scanningSK, spendingPK)
}

func execSpend(client *ethclient.Client, contract *bindings.ERC5564Announcer) error {
	scanningSK, err := getScanningSK()
	if err != nil {
		return err
	}
	spendingSK, err := getSpendingSK()
	if err != nil {
		return err
	}
	return Scan(client, contract, scanningSK, spendingSK)
}

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
