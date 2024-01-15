package main

import (
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"

	"github.com/MariusVanDerWijden/go-stealth/bindings"
	"github.com/MariusVanDerWijden/go-stealth/scanner"
	"github.com/ethereum/go-ethereum/common"
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
		Items: []string{"View events", "Spend events", "Start daemon"},
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
		case 2:
			execDaemon(client, contract)
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
	return scanner.Daemon(client, contract, scanningSK, spendingPK)
}
