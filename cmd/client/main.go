package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	gostealth "github.com/MariusVanDerWijden/go-stealth"
	"github.com/MariusVanDerWijden/go-stealth/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	secretKey *ecdsa.PrivateKey
	client    ethclient.Client
)

func main() {

}

func generateStealthAddr() error {
	var secret [32]byte
	n, err := rand.Read(secret[:])
	if err != nil || n != 32 {
		return errors.New("error generating randomness")
	}
	pubKeyX, pubKeyY, addr := gostealth.NewStealthAddress(secret)
	_ = pubKeyX
	_ = pubKeyY
	_ = addr
	return nil
}

func scan(start uint64, addresses []common.Address, callers []common.Address, contract *bindings.ERC5564Announcer) error {
	filterOpts := bind.FilterOpts{Start: start, End: nil, Context: context.Background()}
	schemeIDs := []*big.Int{new(big.Int)} // Secp256k1 has scheme id 0
	it, err := contract.FilterAnnouncement(&filterOpts, schemeIDs, addresses, callers)
	if err != nil {
		return err
	}
	defer it.Close()
	for it.Next() {
		if err := handleEvent(it.Event, client); err != nil {
			return err
		}
	}
	return it.Error()
}

func wait(start uint64, addresses []common.Address, callers []common.Address, contract *bindings.ERC5564Announcer) error {
	schemeIDs := []*big.Int{new(big.Int)} // Secp256k1 has scheme id 0
	sink := make(chan *bindings.ERC5564AnnouncerAnnouncement)
	sub, err := contract.WatchAnnouncement(&bind.WatchOpts{}, sink, schemeIDs, addresses, callers)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	for {
		ev := <-sink
		if err := handleEvent(ev, client); err != nil {
			return err
		}
	}
}

func handleEvent(event *bindings.ERC5564AnnouncerAnnouncement, client ethclient.Client) error {
	fmt.Printf("Found event: scheme: %v stealthAddr: %v caller: %v epheremeralPubKey: %v metadata: %v \n", event.SchemeId, event.StealthAddress, event.Caller, event.EphemeralPubKey, event.Metadata)
	addr := gostealth.ComputeSharedSecret(event.EphemeralPubKey, secretKey)
	// check if address has funds
	bal, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}
	if bal.Cmp(new(big.Int)) != 0 {
		fmt.Printf("Found my stealth address: %v with balance %v\n", addr, bal)
	}
	return nil
}
