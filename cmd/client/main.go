package main

import (
	"context"
	"errors"
	"math/big"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"

	"github.com/MariusVanDerWijden/go-stealth/bindings"
	"github.com/MariusVanDerWijden/go-stealth/scanner"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Sepolia testdata
	// https://rpc2.sepolia.org
	// 0xFe6335f5dc5a469e74fB6a9FDAe116bFD5346365
	// 45D342EB58207CB50824AD8D4E446AAE6C70DAA8C39E08E8F8B20E62CCC3BE31
	// 03e017e9d9dbcb9ce5771acfce74c95bc0eafb5db37ef4b1ac62375f8e7a4c8aef
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
		Items: []string{"View events", "Spend events", "Start daemon"},
	}
	for {
		idx, _, err := promptInit.Run()
		if err != nil {
			color.Red("Goodbye: %v", err)
			return nil
		}
		switch idx {
		case 0:
			if err := execView(client, contract); err != nil {
				color.Red("Error execution view: %v", err)
			}
		case 1:
			if err := execSpend(client, contract); err != nil {
				color.Red("Error execution spend: %v", err)
			}
		case 2:
			if err := execDaemon(client, contract); err != nil {
				color.Red("Error execution daemon: %v", err)
			}
		case 3:
			if err := execAnnounce(client, contract); err != nil {
				color.Red("Error announcing stealth address: %v", err)
			}
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
	color.Yellow("Starting view")
	return scanner.ViewScan(client, contract, scanningSK, spendingPK)
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
	color.Yellow("Starting scan")
	return scanner.Scan(client, contract, scanningSK, spendingSK)
}

func execDaemon(client *ethclient.Client, contract *bindings.ERC5564Announcer) error {
	scanningSK, err := getScanningSK()
	if err != nil {
		return err
	}
	spendingPK, err := getSpendingPK()
	if err != nil {
		return err
	}
	color.Yellow("Starting daemon")
	return scanner.Daemon(client, contract, scanningSK, spendingPK)
}

func execAnnounce(client *ethclient.Client, contract *bindings.ERC5564Announcer) error {
	prompt := promptui.Prompt{Label: "Please provide the stealth address"}
	str, err := prompt.Run()
	if err != nil {
		return err
	}
	address := common.HexToAddress(str)

	prompt = promptui.Prompt{Label: "Please provide the ephemeral pubkey (in hex)"}
	str, err = prompt.Run()
	if err != nil {
		return err
	}
	ephemeralPK := common.Hex2Bytes(str)

	prompt = promptui.Prompt{Label: "Please provide the metadata (in hex)"}
	str, err = prompt.Run()
	if err != nil {
		return err
	}
	metadata := common.Hex2Bytes(str)

	prompt = promptui.Prompt{Label: "Please provide the key for sending the transaction (in hex)"}
	str, err = prompt.Run()
	if err != nil {
		return err
	}
	key, err := crypto.HexToECDSA(str)
	if err != nil {
		return err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(key, chainId)
	if err != nil {
		return err
	}

	tx, err := contract.Announce(opts, new(big.Int), address, ephemeralPK, metadata)
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return errors.New("transaction mined but failed")
	}
	return nil
}
